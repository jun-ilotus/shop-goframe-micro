package goods

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	goods_images "service/app/goods/api/goods_images/v1"

	"service/app/gateway-admin/api/goods/v1"
)

func (c *ControllerV1) GoodsImagesGetList(ctx context.Context, req *v1.GoodsImagesGetListReq) (res *v1.GoodsImagesGetListRes, err error) {
	grpcReq := &goods_images.GoodsImagesGetListReq{}
	if err = gconv.Struct(req, grpcReq); err != nil {
		return nil, err
	}

	grpcRes, err := c.GoodsImagesClient.GetList(ctx, grpcReq)
	if err != nil {
		return nil, err
	}

	res = &v1.GoodsImagesGetListRes{
		Page:  grpcRes.Data.Page,
		Size:  grpcRes.Data.Size,
		Total: grpcRes.Data.Total,
	}

	if err := gconv.Structs(grpcRes.Data.List, &res.List); err != nil {
		return nil, err
	}
	return res, nil
}
