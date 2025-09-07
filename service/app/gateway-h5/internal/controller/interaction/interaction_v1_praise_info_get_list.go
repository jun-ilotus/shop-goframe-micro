package interaction

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"

	"service/app/gateway-h5/api/interaction/v1"
	praise "service/app/interaction/api/praise_info/v1"
)

func (c *ControllerV1) PraiseInfoGetList(ctx context.Context, req *v1.PraiseInfoGetListReq) (res *v1.PraiseInfoGetListRes, err error) {
	grpcReq := &praise.PraiseInfoGetListReq{}
	if err = gconv.Struct(req, grpcReq); err != nil {
		return nil, err
	}
	grpcRes, err := c.PraiseInfoClient.GetList(ctx, grpcReq)
	if err != nil {
		return nil, err
	}
	res = &v1.PraiseInfoGetListRes{
		Page:  grpcRes.Data.Page,
		Size:  grpcRes.Data.Size,
		Total: grpcRes.Data.Total,
	}
	if err = gconv.Structs(grpcRes.Data.List, &res.List); err != nil {
		return nil, err
	}
	return res, nil
}
