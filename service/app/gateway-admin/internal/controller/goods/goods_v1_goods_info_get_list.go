package goods

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	goods_info "service/app/goods/api/goods_info/v1"

	"service/app/gateway-admin/api/goods/v1"
)

func (c *ControllerV1) GoodsInfoGetList(ctx context.Context, req *v1.GoodsInfoGetListReq) (res *v1.GoodsInfoGetListRes, err error) {
	grpcReq := &goods_info.GoodsInfoGetListReq{}
	if err = gconv.Struct(req, grpcReq); err != nil {
		return nil, err
	}

	grpcRes, err := c.GoodsInfoClient.GetList(ctx, grpcReq)
	if err != nil {
		return nil, err
	}

	res = &v1.GoodsInfoGetListRes{
		Page:  grpcRes.Data.Page,
		Size:  grpcRes.Data.Size,
		Total: grpcRes.Data.Total,
	}

	if err = gconv.Structs(grpcRes.Data.List, &res.List); err != nil {
		return nil, err
	}
	return res, nil
}
