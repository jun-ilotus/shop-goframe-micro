package order

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	refund_info "service/app/order/api/refund_info/v1"

	"service/app/gateway-h5/api/order/v1"
)

func (c *ControllerV1) RefundInfoGetList(ctx context.Context, req *v1.RefundInfoGetListReq) (res *v1.RefundInfoGetListRes, err error) {
	grpcReq := &refund_info.RefundInfoGetListReq{}
	if err := gconv.Struct(req, grpcReq); err != nil {
		return nil, err
	}
	grpcRes, err := c.RefundInfoClient.GetList(ctx, grpcReq)
	if err != nil {
		return nil, err
	}
	res = &v1.RefundInfoGetListRes{
		Page:  grpcRes.Data.Page,
		Size:  grpcRes.Data.Size,
		Total: grpcRes.Data.Total,
	}
	if err := gconv.Structs(grpcRes.Data.List, &res.List); err != nil {
		return nil, err
	}
	return res, nil
}
