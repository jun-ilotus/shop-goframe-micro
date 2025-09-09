package goods

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	category_info "service/app/goods/api/category_info/v1"

	"service/app/gateway-admin/api/goods/v1"
)

func (c *ControllerV1) CategoryInfoCreate(ctx context.Context, req *v1.CategoryInfoCreateReq) (res *v1.CategoryInfoCreateRes, err error) {
	grpcReq := &category_info.CategoryInfoCreateReq{}
	if err = gconv.Struct(req, grpcReq); err != nil {
		return nil, err
	}
	grpcRes, err := c.CategoryInfoClient.Create(ctx, grpcReq)
	if err != nil {
		return nil, err
	}
	return &v1.CategoryInfoCreateRes{Id: grpcRes.Id}, nil
}
