package search

import (
	"context"
	"encoding/json"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/olivere/elastic/v7"
	"service/app/search/internal/consts"
	"service/app/search/utility/elasticsearch"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"service/app/search/api/search/v1"
)

func (c *ControllerV1) SearchGoods(ctx context.Context, req *v1.SearchGoodsReq) (res *v1.SearchGoodsRes, err error) {
	response := &v1.SearchGoodsRes{
		List:  make([]*v1.GoodsInfoItem, 0),
		Page:  req.Page,
		Size:  req.Size,
		Total: 0,
	}
	infoError := consts.InfoError(consts.SearchGoods, consts.SearchFail)

	// 获取 ES 客户端
	client := elasticsearch.GetClient()
	if client == nil {
		g.Log().Errorf(ctx, "%v ES客户端未初始化", infoError)
		return nil, gerror.NewCode(gcode.CodeInternalError, "搜索服务不可用")
	}

	// 构建查询服务
	boolQuery := elastic.NewBoolQuery()

	// 软删除过滤
	boolQuery.MustNot(elastic.NewWildcardQuery("delete_at", "*?*"))

	// 关键词搜索（只搜索商品名称）
	if req.Keyword != "" {
		matchQuery := elastic.NewMatchQuery("name", req.Keyword)
		boolQuery.Must(matchQuery)
	}

	// 品牌筛选
	if req.Brand != "" {
		matchQuery := elastic.NewMatchQuery("brand", req.Brand)
		boolQuery.Must(matchQuery)
	}

	// 价格区间筛选
	if req.MinPrice > 0 || req.MaxPrice > 0 {
		rangeQuery := elastic.NewRangeQuery("price")
		if req.MinPrice > 0 {
			rangeQuery.Gte(req.MinPrice)
		}
		if req.MaxPrice > 0 {
			rangeQuery.Lte(req.MaxPrice)
		}
		boolQuery.Filter(rangeQuery)
	}

	// 构建搜索请求
	esIndexGoods := g.Cfg().MustGet(ctx, "elasticsearch.indices.goods").String()
	g.Log().Info(ctx, esIndexGoods)
	searchService := client.Search().Index(esIndexGoods).Query(boolQuery)

	// 分页
	searchService.From(int((req.Page - 1) * req.Size)).Size(int(req.Size))

	// 排序
	switch req.Sort {
	case "price_asc":
		searchService.Sort("price", true)
	case "price_desc":
		searchService.Sort("price", false)
	case "sale":
		searchService.Sort("sale", false)
	default:
		searchService.Sort("_score", false)
	}

	// 高亮显示
	highlight := elastic.NewHighlight().Field("name").PreTags("<em>").PostTags("</em>")
	searchService.Highlight(highlight)

	// 执行搜索
	searchResult, err := searchService.Do(ctx)
	if err != nil {
		g.Log().Errorf(ctx, "%v %v", infoError, err)
		return nil, gerror.WrapCode(gcode.CodeInternalError, err, infoError)
	}

	// 设置总数
	response.Total = uint32(searchResult.TotalHits())

	// 处理商品列表数据
	for _, hit := range searchResult.Hits.Hits {
		// 解析商品信息
		var goods v1.GoodsInfoItem
		if err := json.Unmarshal(hit.Source, &goods); err != nil {
			continue
		}

		// 处理高亮
		if highlight, ok := hit.Highlight["name"]; ok && len(highlight) > 0 {
			goods.Highlight = highlight[0]
		} else {
			goods.Highlight = goods.Name
		}
		response.List = append(response.List, &goods)
	}

	return response, nil
}
