package goodsRedis

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gogf/gf/v2/database/gredis"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcache"
	"time"
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
	key := fmt.Sprintf("goods:detail:%d", productId)

	emptyValue := "__EMPTY__"
	return goodsCache.Set(ctx, key, emptyValue, 1*time.Minute)
}

func SetGoodsDetail(ctx context.Context, productId uint32, data interface{}) error {
	key := fmt.Sprintf("goods:detail:%d", productId)

	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return goodsCache.Set(ctx, key, jsonData, time.Minute)
}

func GetGoodsDetail(ctx context.Context, productId uint32) (*g.Var, error) {
	key := fmt.Sprintf("goods:detail:%d", productId)
	result, err := goodsCache.Get(ctx, key)
	if err != nil {
		return nil, err
	}

	if result.IsEmpty() || result.String() == "null" {
		return g.NewVar(nil), nil
	}

	return result, nil
}
