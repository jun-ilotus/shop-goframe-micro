package order

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	order_info "service/app/order/api/order_info/v1"

	"github.com/gogf/gf/v2/errors/gerror"

	"service/app/gateway-h5/api/order/v1"
)

func (c *ControllerV1) OrderInfoCreate(ctx context.Context, req *v1.OrderInfoCreateReq) (res *v1.OrderInfoCreateRes, err error) {
	if len(req.OrderGoodsInfo) == 0 {
		return nil, gerror.New("订单必须包含至少一件商品")
	}

	grpcReq := &order_info.OrderInfoCreateReq{}
	if err := gconv.Struct(req, grpcReq); err != nil {
		return nil, err
	}

	grpcReq.OrderGoodsInfo = make([]*order_info.OrderGoodsItem, len(req.OrderGoodsInfo))
	for i, goods := range req.OrderGoodsInfo {
		grpcReq.OrderGoodsInfo[i] = &order_info.OrderGoodsItem{}
		if err := gconv.Struct(goods, grpcReq.OrderGoodsInfo[i]); err != nil {
			return nil, err
		}
	}

	grpcRes, err := c.OrderInfoClient.Create(ctx, grpcReq)
	if err != nil {
		return nil, err
	}

	res = &v1.OrderInfoCreateRes{}
	if err := gconv.Struct(grpcRes, res); err != nil {
		return nil, err
	}
	return res, nil
}
