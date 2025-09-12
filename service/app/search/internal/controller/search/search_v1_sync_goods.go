package search

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"service/app/search/utility/elasticsearch"
	"time"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"service/app/search/api/search/v1"
)

func (c *ControllerV1) SyncGoods(ctx context.Context, req *v1.SyncGoodsReq) (res *v1.SyncGoodsRes, err error) {
	client := elasticsearch.GetClient()
	if client == nil {
		return nil, gerror.NewCode(gcode.CodeInternalError, "ES客户端未初始化")
	}

	switch req.Operation {
	case "create", "update":
		_, err = client.Index().Index("mall_goods").Id(gconv.String(req.Id)).BodyJson(map[string]interface{}{
			"id":                 req.Id,
			"name":               req.Name,
			"images":             req.Images,
			"price":              req.Price,
			"level1_category_id": req.Level1CategoryId,
			"level2_category_id": req.Level2CategoryId,
			"level3_category_id": req.Level3CategoryId,
			"brand":              req.Brand,
			"stock":              req.Stock,
			"sale":               req.Sale,
			"tags":               req.Tags,
			"detail_info":        req.DetailInfo,
			"created_at":         time.Now().Format(time.RFC3339),
			"updated_at":         time.Now().Format(time.RFC3339),
		}).Do(ctx)
	case "delete":
		_, err = client.Delete().
			Index("mall_goods").
			Id(gconv.String(req.Id)).Do(ctx)
	}
	if err != nil {
		g.Log().Errorf(ctx, "同步商品到ES失败：%v", err)
		return &v1.SyncGoodsRes{Success: false}, err
	}

	return &v1.SyncGoodsRes{Success: true}, nil
}
