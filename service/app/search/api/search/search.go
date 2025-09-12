// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package search

import (
	"context"

	"service/app/search/api/search/v1"
)

type ISearchV1 interface {
	SearchGoods(ctx context.Context, req *v1.SearchGoodsReq) (res *v1.SearchGoodsRes, err error)
	SyncGoods(ctx context.Context, req *v1.SyncGoodsReq) (res *v1.SyncGoodsRes, err error)
}
