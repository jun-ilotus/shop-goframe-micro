package order

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/util/gconv"
	refund_info "service/app/order/api/refund_info/v1"
	"service/utility/middleware"

	"service/app/gateway-h5/api/order/v1"
)

func (c *ControllerV1) RefundInfoCreate(ctx context.Context, req *v1.RefundInfoCreateReq) (res *v1.RefundInfoCreateRes, err error) {
	grpcReq := &refund_info.RefundInfoCreateReq{}
	if err = gconv.Struct(req, grpcReq); err != nil {
		return nil, err
	}
	value := ctx.Value(middleware.CtxUserId)
	userId, ok := value.(uint32)
	if !ok {
		panic("用户ID类型错误或不存在")
	}
	grpcReq.UserId = userId
	fmt.Println(grpcReq.UserId)
	grpcRes, err := c.RefundInfoClient.Create(ctx, grpcReq)
	if err != nil {
		return nil, err
	}
	return &v1.RefundInfoCreateRes{Id: grpcRes.Id}, nil
}
