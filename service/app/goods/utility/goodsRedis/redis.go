package goodsRedis

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/database/gredis"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcache"
	"time"
)

const (
	categoryAllKey = "category:all:data"
	GoodsDetailKey = "goods:detail:"
	EmptyValue     = "__EMPTY__"
)

var goodsCache *gcache.Cache

func InitGoodsRedis(ctx context.Context) error {
	redisConfig, err := g.Cfg().Get(ctx, "redis.goods")
	if err != nil {
		return fmt.Errorf("get redis config err: %v", err)
	}

	config := &gredis.Config{}
	if err := redisConfig.Scan(&config); err != nil {
		return fmt.Errorf("redis config err: %v", err)
	}

	redis, err := gredis.New(config)
	if err != nil {
		return fmt.Errorf("redis config err: %v", err)
	}

	goodsCache = gcache.New()
	goodsCache.SetAdapter(gcache.NewAdapterRedis(redis))

	if _, err := redis.Do(ctx, "PING"); err != nil {
		return fmt.Errorf("redis ping err: %v", err)
	}

	g.Log().Info(ctx, "redis init success")
	return nil
}

func GetGoodsCache() *gcache.Cache {
	return goodsCache
}

func SetEmptyGoodsDetail(ctx context.Context, productId uint32) error {
	key := fmt.Sprintf("%s%d", GoodsDetailKey, productId)

	// 设置一个短时间的空值，防止缓存穿透
	return goodsCache.Set(ctx, key, EmptyValue, 1*time.Minute)
}

func SetGoodsDetail(ctx context.Context, productId uint32, data interface{}) error {
	key := fmt.Sprintf("%s%d", GoodsDetailKey, productId)

	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return goodsCache.Set(ctx, key, jsonData, time.Minute)
}

func GetGoodsDetail(ctx context.Context, productId uint32) (*g.Var, error) {
	key := fmt.Sprintf("%s%d", GoodsDetailKey, productId)
	result, err := goodsCache.Get(ctx, key)
	if err != nil {
		return nil, err
	}

	if result.IsEmpty() || result.String() == "null" {
		return g.NewVar(nil), nil
	}

	return result, nil
}

func DeleteGoodsDetail(ctx context.Context, productId uint32) error {
	key := fmt.Sprintf("%s%d", GoodsDetailKey, productId)
	_, err := goodsCache.Remove(ctx, key)
	return err
}

func GetCategoryAll(ctx context.Context) (*gvar.Var, error) {
	result, err := goodsCache.Get(ctx, categoryAllKey)
	if err != nil {
		return nil, err
	}
	if result.IsEmpty() || result.String() == "null" {
		return gvar.New(nil), nil
	}
	return result, nil
}

func SetCategoryAll(ctx context.Context, data interface{}) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return goodsCache.Set(ctx, categoryAllKey, jsonData, 7*24*time.Hour)
}

func DeleteCategoryAll(ctx context.Context) error {
	_, err := goodsCache.Remove(ctx, categoryAllKey)
	return err
}
