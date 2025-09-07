package goods

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	goods_info "service/app/goods/api/goods_info/v1"

	"service/app/gateway-admin/api/goods/v1"
)

func (c *ControllerV1) GoodsInfoCreate(ctx context.Context, req *v1.GoodsInfoCreateReq) (res *v1.GoodsInfoCreateRes, err error) {
	grpcReq := &goods_info.GoodsInfoCreateReq{}
	if err = gconv.Struct(req, grpcReq); err != nil {
		return nil, err
	}
	grpcRes, err := c.GoodsInfoClient.Create(ctx, grpcReq)
	if err != nil {
		return nil, err
	}
	return &v1.GoodsInfoCreateRes{Id: grpcRes.Id}, nil
}
