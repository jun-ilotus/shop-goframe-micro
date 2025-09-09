package goods

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"

	category_info "service/app/goods/api/category_info/v1"

	"service/app/gateway-admin/api/goods/v1"
)

func (c *ControllerV1) CategoryInfoUpdate(ctx context.Context, req *v1.CategoryInfoUpdateReq) (res *v1.CategoryInfoUpdateRes, err error) {
	grpcReq := &category_info.CategoryInfoUpdateReq{}
	if err = gconv.Struct(req, grpcReq); err != nil {
		return nil, err
	}
	grpcRes, err := c.CategoryInfoClient.Update(ctx, grpcReq)
	if err != nil {
		return nil, err
	}
	return &v1.CategoryInfoUpdateRes{Id: grpcRes.Id}, nil
}
