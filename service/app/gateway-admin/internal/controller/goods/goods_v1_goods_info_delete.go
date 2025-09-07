package goods

import (
	"context"
	goods_info "service/app/goods/api/goods_info/v1"

	"service/app/gateway-admin/api/goods/v1"
)

func (c *ControllerV1) GoodsInfoDelete(ctx context.Context, req *v1.GoodsInfoDeleteReq) (res *v1.GoodsInfoDeleteRes, err error) {
	_, err = c.GoodsInfoClient.Delete(ctx, &goods_info.GoodsInfoDeleteReq{Id: req.Id})
	if err != nil {
		return nil, err
	}
	return &v1.GoodsInfoDeleteRes{}, nil
}
