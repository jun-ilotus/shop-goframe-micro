package goods

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	goods_images "service/app/goods/api/goods_images/v1"

	"service/app/gateway-admin/api/goods/v1"
)

func (c *ControllerV1) GoodsImagesCreate(ctx context.Context, req *v1.GoodsImagesCreateReq) (res *v1.GoodsImagesCreateRes, err error) {
	grpcReq := &goods_images.GoodsImagesCreateReq{}
	if err = gconv.Struct(req, grpcReq); err != nil {
		return nil, err
	}
	grpcRes, err := c.GoodsImagesClient.Create(ctx, grpcReq)
	if err != nil {
		return nil, err
	}
	return &v1.GoodsImagesCreateRes{Id: grpcRes.Id}, nil
}
