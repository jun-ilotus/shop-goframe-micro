// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package banner

import (
	"context"

	"service/app/gateway-h5/api/banner/v1"
)

type IBannerV1 interface {
	PositionInfoGetList(ctx context.Context, req *v1.PositionInfoGetListReq) (res *v1.PositionInfoGetListRes, err error)
	RotationInfoGetList(ctx context.Context, req *v1.RotationInfoGetListReq) (res *v1.RotationInfoGetListRes, err error)
}
