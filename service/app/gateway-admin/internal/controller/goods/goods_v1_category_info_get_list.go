package goods

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	category_info "service/app/goods/api/category_info/v1"

	"service/app/gateway-admin/api/goods/v1"
)

func (c *ControllerV1) CategoryInfoGetList(ctx context.Context, req *v1.CategoryInfoGetListReq) (res *v1.CategoryInfoGetListRes, err error) {
	grpcReq := &category_info.CategoryInfoGetListReq{}
	if err = gconv.Struct(req, grpcReq); err != nil {
		return nil, err
	}
	grpcRes, err := c.CategoryInfoClient.GetList(ctx, grpcReq)
	if err != nil {
		return nil, err
	}
	res = &v1.CategoryInfoGetListRes{
		Page:  grpcRes.Data.Page,
		Size:  grpcRes.Data.Size,
		Total: grpcRes.Data.Total,
	}

	if err = gconv.Structs(grpcRes.Data.List, &res.List); err != nil {
		return nil, err
	}

	return res, nil
}
