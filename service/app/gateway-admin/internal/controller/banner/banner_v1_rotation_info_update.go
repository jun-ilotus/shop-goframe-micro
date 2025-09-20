package banner

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	rotation_info "service/app/banner/api/rotation_info/v1"

	"service/app/gateway-admin/api/banner/v1"
)

func (c *ControllerV1) RotationInfoUpdate(ctx context.Context, req *v1.RotationInfoUpdateReq) (res *v1.RotationInfoUpdateRes, err error) {
	grpcReq := &rotation_info.RotationInfoUpdateReq{}
	if err := gconv.Struct(req, grpcReq); err != nil {
		return nil, err
	}
	grpcRes, err := c.RotationInfoClient.Update(ctx, grpcReq)
	if err != nil {
		return nil, err
	}
	return &v1.RotationInfoUpdateRes{Id: grpcRes.Id}, nil
}
