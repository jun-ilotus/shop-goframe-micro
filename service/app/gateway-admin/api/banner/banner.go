// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package banner

import (
	"context"

	"service/app/gateway-admin/api/banner/v1"
)

type IBannerV1 interface {
	PositionInfoGetList(ctx context.Context, req *v1.PositionInfoGetListReq) (res *v1.PositionInfoGetListRes, err error)
	PositionInfoCreate(ctx context.Context, req *v1.PositionInfoCreateReq) (res *v1.PositionInfoCreateRes, err error)
	PositionInfoUpdate(ctx context.Context, req *v1.PositionInfoUpdateReq) (res *v1.PositionInfoUpdateRes, err error)
	PositionInfoDelete(ctx context.Context, req *v1.PositionInfoDeleteReq) (res *v1.PositionInfoDeleteRes, err error)
	RotationInfoGetList(ctx context.Context, req *v1.RotationInfoGetListReq) (res *v1.RotationInfoGetListRes, err error)
	RotationInfoCreate(ctx context.Context, req *v1.RotationInfoCreateReq) (res *v1.RotationInfoCreateRes, err error)
	RotationInfoUpdate(ctx context.Context, req *v1.RotationInfoUpdateReq) (res *v1.RotationInfoUpdateRes, err error)
	RotationInfoDelete(ctx context.Context, req *v1.RotationInfoDeleteReq) (res *v1.RotationInfoDeleteRes, err error)
}
