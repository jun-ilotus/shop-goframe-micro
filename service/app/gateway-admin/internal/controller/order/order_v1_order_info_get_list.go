package order

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	order_info "service/app/order/api/order_info/v1"

	"service/app/gateway-admin/api/order/v1"
)

func (c *ControllerV1) OrderInfoGetList(ctx context.Context, req *v1.OrderInfoGetListReq) (res *v1.OrderInfoGetListRes, err error) {
	grpcReq := &order_info.OrderInfoGetListReq{}
	if err = gconv.Struct(req, grpcReq); err != nil {
		return nil, err
	}
	grpcRes, err := c.OrderInfoClient.GetList(ctx, grpcReq)
	if err != nil {
		return nil, err
	}
	res = &v1.OrderInfoGetListRes{
		Page:  grpcRes.Data.Page,
		Size:  grpcRes.Data.Size,
		Total: grpcRes.Data.Total,
	}
	if err := gconv.Structs(grpcRes.Data.List, &res.List); err != nil {
		return nil, err
	}
	return res, nil
}
