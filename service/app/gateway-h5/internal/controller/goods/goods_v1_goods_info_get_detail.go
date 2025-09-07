package goods

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	goods_info "service/app/goods/api/goods_info/v1"

	"service/app/gateway-h5/api/goods/v1"
)

func (c *ControllerV1) GoodsInfoGetDetail(ctx context.Context, req *v1.GoodsInfoGetDetailReq) (res *v1.GoodsInfoGetDetailRes, err error) {
	grpcReq := &goods_info.GoodsInfoGetDetailReq{}
	if err = gconv.Struct(req, grpcReq); err != nil {
		return nil, err
	}
	grpcRes, err := c.GoodsInfoClient.GetDetail(ctx, grpcReq)
	if err != nil {
		return nil, err
	}
	res = &v1.GoodsInfoGetDetailRes{}
	if err = gconv.Struct(grpcRes.Data, res); err != nil {
		return nil, err
	}
	return res, nil
}
