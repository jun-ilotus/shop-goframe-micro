package banner

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	position_info "service/app/banner/api/position_info/v1"

	"service/app/gateway-admin/api/banner/v1"
)

func (c *ControllerV1) PositionInfoUpdate(ctx context.Context, req *v1.PositionInfoUpdateReq) (res *v1.PositionInfoUpdateRes, err error) {
	grpcReq := &position_info.PositionInfoUpdateReq{}
	if err := gconv.Struct(req, grpcReq); err != nil {
		return nil, err
	}
	grpcRes, err := c.PositionInfoClient.Update(ctx, grpcReq)
	if err != nil {
		return nil, err
	}
	return &v1.PositionInfoUpdateRes{Id: grpcRes.Id}, nil
}
