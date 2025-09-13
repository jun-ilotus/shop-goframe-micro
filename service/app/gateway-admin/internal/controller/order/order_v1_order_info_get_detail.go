package order

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	order_info "service/app/order/api/order_info/v1"

	"service/app/gateway-admin/api/order/v1"
)

func (c *ControllerV1) OrderInfoGetDetail(ctx context.Context, req *v1.OrderInfoGetDetailReq) (res *v1.OrderInfoGetDetailRes, err error) {
	grpcReq := &order_info.OrderInfoGetDetailReq{}
	if err := gconv.Struct(req, grpcReq); err != nil {
		return nil, err
	}

	grpcRes, err := c.OrderInfoClient.GetDetail(ctx, grpcReq)
	if err != nil {
		return nil, err
	}
	res = &v1.OrderInfoGetDetailRes{}

	if grpcRes.OrderInfo != nil {
		res.OrderInfo = &v1.OrderInfoItem{}
		if err := gconv.Struct(grpcRes.OrderInfo, res.OrderInfo); err != nil {
			return nil, err
		}
	}

	if len(grpcRes.OrderGoodsInfos) > 0 {
		res.OrderGoodsInfo = make([]*v1.OrderGoodsDetail, len(grpcRes.OrderGoodsInfos))
		for i, goods := range grpcRes.OrderGoodsInfos {
			res.OrderGoodsInfo[i] = &v1.OrderGoodsDetail{}
			if err := gconv.Struct(goods, res.OrderGoodsInfo[i]); err != nil {
				return nil, err
			}
		}
	}

	return res, nil
}
