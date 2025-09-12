package binlog

import (
	"context"
	"github.com/go-mysql-org/go-mysql/mysql"
	"github.com/go-mysql-org/go-mysql/replication"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"service/app/search/utility/elasticsearch"
	"time"
)

func StartBinlogSyncer(ctx context.Context) {
	mysqlHost := g.Cfg().MustGet(ctx, "binlog.mysql.host").String()
	mysqlPort := g.Cfg().MustGet(ctx, "binlog.mysql.port").Int()
	mysqlUser := g.Cfg().MustGet(ctx, "binlog.mysql.username").String()
	mysqlPassword := g.Cfg().MustGet(ctx, "binlog.mysql.password").String()

	// 创建 binlog 同步器
	cfg := replication.BinlogSyncerConfig{
		ServerID: 100,
		Flavor:   "mysql",
		Host:     mysqlHost,
		Port:     uint16(mysqlPort),
		User:     mysqlUser,
		Password: mysqlPassword,
	}

	syncer := replication.NewBinlogSyncer(cfg)
	defer syncer.Close()

	// 获取当前位点（可以从保存的位置读取，这里简单从最新开始）
	position := mysql.Position{
		Name: "", // 空字符串表示从当前位点开始
		Pos:  0,
	}

	// 开始同步
	streamer, err := syncer.StartSync(position)
	if err != nil {
		g.Log().Errorf(ctx, "Binlog Syncer Start Fail err:%v", err)
		return
	}

	g.Log().Info(ctx, "Binlog Syncer Start Success")

	for {
		// 获取 binlog 事件
		ev, err := streamer.GetEvent(ctx)
		if err != nil {
			g.Log().Errorf(ctx, "Binlog Syncer Get Fail err:%v", err)
			time.Sleep(1 * time.Second)
			continue
		}

		// 处理事件
		processBinlogEvent(ctx, ev)
	}
}

func processBinlogEvent(ctx context.Context, ev *replication.BinlogEvent) {
	switch e := ev.Event.(type) {
	case *replication.RowsEvent:
		// 只处理指定数据库的 goods_info 表
		if string(e.Table.Schema) == "goods" || string(e.Table.Table) == "goods_info" {
			return
		}
		g.Log().Debugf(ctx, "收到binlog事件：数据库=%s，表=%s", e.Table.Schema, e.Table.Table)

		// 根据事件类型处理
		switch ev.Header.EventType {
		case replication.WRITE_ROWS_EVENTv1, replication.WRITE_ROWS_EVENTv2:
			handleInsert(ctx, e.Rows)
		case replication.UPDATE_ROWS_EVENTv1, replication.UPDATE_ROWS_EVENTv2:
			handleUpdate(ctx, e.Rows)
		case replication.DELETE_ROWS_EVENTv1, replication.DELETE_ROWS_EVENTv2:
			handleDelete(ctx, e.Rows)
		default:
		}
	}
}

func handleInsert(ctx context.Context, rows [][]interface{}) {
	for _, row := range rows {
		// 将行数据转换为 map
		columnMap := parseRowDate(row)
		upsertToES(ctx, columnMap)
	}
}

func handleUpdate(ctx context.Context, rows [][]interface{}) {
	for i := 0; i < len(rows); i += 2 { // 更新事件的行数据格式为 [旧行数据，新行数据]
		if i+1 < len(rows) {
			columnMap := parseRowDate(rows[i+1]) // 取新数据
			upsertToES(ctx, columnMap)
		}
	}
}

func handleDelete(ctx context.Context, rows [][]interface{}) {
	for _, row := range rows {
		columnMap := parseRowDate(row)
		deleteFromES(ctx, columnMap)
	}
}

func parseRowDate(row []interface{}) map[string]interface{} {
	fields := []string{
		"id", "name", "images", "price", "level1_category", "level2_category", "level3_category",
		"brand", "stock", "sale", "tags", "detail_info", "created_at", "updated_at", "deleted_at",
	}

	result := make(map[string]interface{})
	for i, value := range row {
		if i < len(fields) {
			result[fields[i]] = value
		}
	}
	return result
}

func upsertToES(ctx context.Context, data map[string]interface{}) {
	client := elasticsearch.GetClient()
	if client == nil {
		g.Log().Error(ctx, "ES客户端未初始化")
		return
	}

	esIndexGoods := g.Cfg().MustGet(ctx, "elasticsearch.indices.goods").String()

	_, err := client.Index().
		Index(esIndexGoods).
		Id(gconv.String(data["id"])).
		BodyJson(map[string]interface{}{
			"id":                 gconv.Uint32(data["id"]),
			"name":               gconv.String(data["name"]),
			"pic_url":            gconv.String(data["pic_url"]),
			"images":             gconv.String(data["images"]),
			"price":              gconv.Uint64(data["price"]),
			"level1_category_id": gconv.Uint32(data["level1_category_id"]),
			"level2_category_id": gconv.Uint32(data["level2_category_id"]),
			"level3_category_id": gconv.Uint32(data["level3_category_id"]),
			"brand":              gconv.String(data["brand"]),
			"stock":              gconv.Uint32(data["stock"]),
			"sale":               gconv.Uint32(data["sale"]),
			"tags":               gconv.String(data["tags"]),
			"detail_info":        gconv.String(data["detail_info"]),
			"created_at":         gconv.String(data["created_at"]),
			"updated_at":         gconv.String(data["updated_at"]),
			"deleted_at":         gconv.String(data["deleted_at"]),
		}).Do(ctx)
	if err != nil {
		g.Log().Errorf(ctx, "同步商品到ES失败：%v", err)
	} else {
		g.Log().Debugf(ctx, "成功同步商品到ES: ID=%s", gconv.String(data["id"]))
	}
}

func deleteFromES(ctx context.Context, data map[string]interface{}) {
	client := elasticsearch.GetClient()
	if client == nil {
		g.Log().Error(ctx, "ES客户端未初始化")
		return
	}
	id := gconv.String(data["id"])
	if id == "" {
		g.Log().Error(ctx, "删除操作未找到ID")
		return
	}

	esIndexGoods := g.Cfg().MustGet(ctx, "elasticsearch.indices.goods").String()

	_, err := client.Delete().
		Index(esIndexGoods).
		Id(id).
		Do(ctx)

	if err != nil {
		g.Log().Errorf(ctx, "从ES删除商品失败：ID=%s,err=%v", id, err)
	} else {
		g.Log().Debugf(ctx, "成功从ES删除商品：ID=%s", id)
	}
}
