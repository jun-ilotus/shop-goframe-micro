package goods

import (
	"context"
	goods_images "service/app/goods/api/goods_images/v1"

	"service/app/gateway-admin/api/goods/v1"
)

func (c *ControllerV1) GoodsImagesDelete(ctx context.Context, req *v1.GoodsImagesDeleteReq) (res *v1.GoodsImagesDeleteRes, err error) {
	_, err = c.GoodsImagesClient.Delete(ctx, &goods_images.GoodsImagesDeleteReq{Id: req.Id})
	if err != nil {
		return nil, err
	}
	return &v1.GoodsImagesDeleteRes{}, nil
}
