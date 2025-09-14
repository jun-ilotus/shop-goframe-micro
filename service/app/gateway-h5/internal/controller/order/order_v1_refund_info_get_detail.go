package order

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	refund_info "service/app/order/api/refund_info/v1"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"service/app/gateway-h5/api/order/v1"
)

func (c *ControllerV1) RefundInfoGetDetail(ctx context.Context, req *v1.RefundInfoGetDetailReq) (res *v1.RefundInfoGetDetailRes, err error) {
	grpcReq := &refund_info.RefundInfoGetDetailReq{}
	if err := gconv.Struct(req, grpcReq); err != nil {
		return nil, err
	}
	grpcRes, err := c.RefundInfoClient.GetDetail(ctx, grpcReq)
	if err != nil {
		return nil, err
	}
	if grpcRes.Data == nil {
		return nil, gerror.NewCode(gcode.CodeNotFound, "退款记录不存在")
	}
	res = &v1.RefundInfoGetDetailRes{}
	if err := gconv.Struct(grpcRes.Data, res); err != nil {
		return nil, err
	}
	return res, nil
}
