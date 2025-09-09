package goods

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	category_info "service/app/goods/api/category_info/v1"

	"service/app/gateway-admin/api/goods/v1"
)

func (c *ControllerV1) CategoryInfoGetAll(ctx context.Context, req *v1.CategoryInfoGetAllReq) (res *v1.CategoryInfoGetAllRes, err error) {
	grpcReq := &category_info.CategoryInfoGetAllReq{}
	if err = gconv.Struct(req, grpcReq); err != nil {
		return nil, err
	}
	grpcRes, err := c.CategoryInfoClient.GetAll(ctx, grpcReq)
	if err != nil {
		return nil, err
	}

	res = &v1.CategoryInfoGetAllRes{
		List:  make([]*v1.CategoryInfoItem, len(grpcRes.List)),
		Total: grpcRes.Total,
	}
	if err = gconv.Structs(grpcRes.List, &res.List); err != nil {
		return nil, err
	}

	return res, nil
}
