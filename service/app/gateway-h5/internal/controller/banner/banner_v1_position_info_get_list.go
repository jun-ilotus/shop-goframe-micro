package banner

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	position_info "service/app/banner/api/position_info/v1"

	"service/app/gateway-h5/api/banner/v1"
)

func (c *ControllerV1) PositionInfoGetList(ctx context.Context, req *v1.PositionInfoGetListReq) (res *v1.PositionInfoGetListRes, err error) {
	grpcReq := &position_info.PositionInfoGetListReq{}
	if err := gconv.Struct(req, grpcReq); err != nil {
		return nil, err
	}

	grpcRes, err := c.PositionInfoClient.GetList(ctx, grpcReq)

	if err != nil {
		return nil, err
	}

	res = &v1.PositionInfoGetListRes{
		Page:  grpcRes.Data.Page,
		Size:  grpcRes.Data.Size,
		Total: grpcRes.Data.Total,
	}

	if err := gconv.Structs(grpcRes.Data.List, &res.List); err != nil {
		return nil, err
	}

	return res, nil
}
