// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package goods

import (
	"context"

	"service/app/gateway-admin/api/goods/v1"
)

type IGoodsV1 interface {
	CategoryInfoGetList(ctx context.Context, req *v1.CategoryInfoGetListReq) (res *v1.CategoryInfoGetListRes, err error)
	CategoryInfoGetAll(ctx context.Context, req *v1.CategoryInfoGetAllReq) (res *v1.CategoryInfoGetAllRes, err error)
	CategoryInfoCreate(ctx context.Context, req *v1.CategoryInfoCreateReq) (res *v1.CategoryInfoCreateRes, err error)
	CategoryInfoUpdate(ctx context.Context, req *v1.CategoryInfoUpdateReq) (res *v1.CategoryInfoUpdateRes, err error)
	CategoryInfoDelete(ctx context.Context, req *v1.CategoryInfoDeleteReq) (res *v1.CategoryInfoDeleteRes, err error)
	GoodsImagesGetList(ctx context.Context, req *v1.GoodsImagesGetListReq) (res *v1.GoodsImagesGetListRes, err error)
	GoodsImagesCreate(ctx context.Context, req *v1.GoodsImagesCreateReq) (res *v1.GoodsImagesCreateRes, err error)
	GoodsImagesDelete(ctx context.Context, req *v1.GoodsImagesDeleteReq) (res *v1.GoodsImagesDeleteRes, err error)
	GoodsInfoGetDetail(ctx context.Context, req *v1.GoodsInfoGetDetailReq) (res *v1.GoodsInfoGetDetailRes, err error)
	GoodsInfoGetList(ctx context.Context, req *v1.GoodsInfoGetListReq) (res *v1.GoodsInfoGetListRes, err error)
	GoodsInfoCreate(ctx context.Context, req *v1.GoodsInfoCreateReq) (res *v1.GoodsInfoCreateRes, err error)
	GoodsInfoUpdate(ctx context.Context, req *v1.GoodsInfoUpdateReq) (res *v1.GoodsInfoUpdateRes, err error)
	GoodsInfoDelete(ctx context.Context, req *v1.GoodsInfoDeleteReq) (res *v1.GoodsInfoDeleteRes, err error)
}
