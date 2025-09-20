package banner

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	position_info "service/app/banner/api/position_info/v1"

	"service/app/gateway-admin/api/banner/v1"
)

func (c *ControllerV1) PositionInfoCreate(ctx context.Context, req *v1.PositionInfoCreateReq) (res *v1.PositionInfoCreateRes, err error) {
	grpcReq := &position_info.PositionInfoCreateReq{}
	if err := gconv.Struct(req, grpcReq); err != nil {
		return nil, err
	}
	grpcRes, err := c.PositionInfoClient.Create(ctx, grpcReq)
	if err != nil {
		return nil, err
	}
	return &v1.PositionInfoCreateRes{Id: grpcRes.Id}, nil
}
