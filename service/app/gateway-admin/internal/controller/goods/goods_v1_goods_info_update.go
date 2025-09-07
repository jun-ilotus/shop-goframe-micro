package goods

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	goods_info "service/app/goods/api/goods_info/v1"

	"service/app/gateway-admin/api/goods/v1"
)

func (c *ControllerV1) GoodsInfoUpdate(ctx context.Context, req *v1.GoodsInfoUpdateReq) (res *v1.GoodsInfoUpdateRes, err error) {
	grpcReq := &goods_info.GoodsInfoUpdateReq{}
	if err = gconv.Struct(req, grpcReq); err != nil {
		return nil, err
	}
	grpcRes, err := c.GoodsInfoClient.Update(ctx, grpcReq)
	if err != nil {
		return nil, err
	}

	return &v1.GoodsInfoUpdateRes{Id: grpcRes.Id}, nil
}
