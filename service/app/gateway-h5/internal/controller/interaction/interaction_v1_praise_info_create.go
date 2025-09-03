package interaction

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"

	"service/app/gateway-h5/api/interaction/v1"
	praise "service/app/interaction/api/praise_info/v1"
)

func (c *ControllerV1) PraiseInfoCreate(ctx context.Context, req *v1.PraiseInfoCreateReq) (res *v1.PraiseInfoCreateRes, err error) {
	grpcReq := &praise.PraiseInfoCreateReq{}
	if err = gconv.Struct(req, grpcReq); err != nil {
		return nil, err
	}
	grpcRes, err := c.PraiseInfoClient.Create(ctx, grpcReq)
	if err != nil {
		return nil, err
	}
	return &v1.PraiseInfoCreateRes{Id: grpcRes.Id}, nil
}
