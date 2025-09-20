package banner

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	rotation_info "service/app/banner/api/rotation_info/v1"

	"service/app/gateway-admin/api/banner/v1"
)

func (c *ControllerV1) RotationInfoGetList(ctx context.Context, req *v1.RotationInfoGetListReq) (res *v1.RotationInfoGetListRes, err error) {
	grpcReq := &rotation_info.RotationInfoGetListReq{}
	if err = gconv.Struct(req, grpcReq); err != nil {
		return nil, err
	}
	grpcRes, err := c.RotationInfoClient.GetList(ctx, grpcReq)
	if err != nil {
		return nil, err
	}
	res = &v1.RotationInfoGetListRes{
		Page:  grpcRes.Data.Page,
		Size:  grpcRes.Data.Size,
		Total: grpcRes.Data.Total,
	}
	if err = gconv.Scan(grpcRes.Data.List, &res.List); err != nil {
		return nil, err
	}
	return res, nil
}
