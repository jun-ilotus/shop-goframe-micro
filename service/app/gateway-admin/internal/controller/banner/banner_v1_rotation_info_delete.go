package banner

import (
	"context"
	rotation_info "service/app/banner/api/rotation_info/v1"

	"service/app/gateway-admin/api/banner/v1"
)

func (c *ControllerV1) RotationInfoDelete(ctx context.Context, req *v1.RotationInfoDeleteReq) (res *v1.RotationInfoDeleteRes, err error) {
	_, err = c.RotationInfoClient.Delete(ctx, &rotation_info.RotationInfoDeleteReq{Id: req.Id})
	if err != nil {
		return nil, err
	}
	return &v1.RotationInfoDeleteRes{}, nil
}
