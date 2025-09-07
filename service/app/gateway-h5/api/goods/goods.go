// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package goods

import (
	"context"

	"service/app/gateway-h5/api/goods/v1"
)

type IGoodsV1 interface {
	GoodsImagesGetList(ctx context.Context, req *v1.GoodsImagesGetListReq) (res *v1.GoodsImagesGetListRes, err error)
	GoodsInfoGetDetail(ctx context.Context, req *v1.GoodsInfoGetDetailReq) (res *v1.GoodsInfoGetDetailRes, err error)
	GoodsInfoGetList(ctx context.Context, req *v1.GoodsInfoGetListReq) (res *v1.GoodsInfoGetListRes, err error)
}
