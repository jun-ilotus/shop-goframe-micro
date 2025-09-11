package elasticsearch

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/olivere/elastic/v7"
)

var client *elastic.Client

func Init(ctx context.Context) error {
	esAddress := g.Cfg().MustGet(ctx, "elasticsearch.address").String()
	sniff := g.Cfg().MustGet(ctx, "elasticsearch.sniff").Bool()
	healthcheck := g.Cfg().MustGet(ctx, "elasticsearch.healthcheck").Bool()

	options := []elastic.ClientOptionFunc{
		elastic.SetURL(esAddress),
		elastic.SetSniff(sniff),
		elastic.SetHealthcheck(healthcheck),
	}

	var err error
	client, err = elastic.NewClient(options...)
	if err != nil {
		return fmt.Errorf("elastic NewClient error: %v", err)
	}

	_, _, err = client.Ping(esAddress).Do(ctx)
	if err != nil {
		return fmt.Errorf("elastic NewClient ping error: %v", err)
	}

	if err := createGoodsIndex(ctx); err != nil {
		return fmt.Errorf("elastic createGoodsIndex error: %v", err)
	}
	g.Log().Info(ctx, "elastic createGoodsIndex success")
	return nil
}

func GetClient() *elastic.Client {
	return client
}

func createGoodsIndex(ctx context.Context) error {
	esIndexGoods := g.Cfg().MustGet(ctx, "elasticsearch.indices.goods").String()
	// 检查索引是否存在
	exists, err := client.IndexExists(esIndexGoods).Do(ctx)
	if err != nil {
		return err
	}

	if exists {
		g.Log().Info(ctx, "elasticsearch.indices.goods already exists")
	}

	mapping := `{
		"mappings": {
			"properties": {
				"id": {"type": "long"},
				"name": {
					"type": "text",
					"analyzer": "ik_max_word",
					"search_analyzer": "ik_smart"
				},
				"pic_url": {"type": "keyword"},
				"images": {"type": "keyword"},
				"price": {"type": "long"},
				"level1_category_id": {"type": "long"},
				"level2_category_id": {"type": "long"},
				"level3_category_id": {"type": "long"},
				"brand": {
					"type": "keyword",
					"fields": {
						"text": {"type": "text"}
					}
				},
				"stock": {"type": "long"},
				"sale": {"type": "long"},
				"tags": {"type": "keyword"},
				"detail_info": {"type": "text"},
				"created_at": {"type": "date"},
				"updated_at": {"type": "date"}
			}
		}
	}`

	createIndex, err := client.CreateIndex(esIndexGoods).Body(mapping).Do(ctx)
	if err != nil {
		g.Log().Errorf(ctx, "elastic CreateIndex error: %v", err)
		return err
	}

	if !createIndex.Acknowledged {
		return fmt.Errorf("elastic CreateIndex error: %v", createIndex.Acknowledged)
	}

	g.Log().Info(ctx, "elastic CreateIndex success")
	return nil
}
