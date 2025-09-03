package interaction

import (
	"context"
	"service/app/gateway-h5/api/interaction/v1"
	praise "service/app/interaction/api/praise_info/v1"
)

func (c *ControllerV1) PraiseInfoDelete(ctx context.Context, req *v1.PraiseInfoDeleteReq) (res *v1.PraiseInfoDeleteRes, err error) {
	_, err = c.PraiseInfoClient.Delete(ctx, &praise.PraiseInfoDeleteReq{Id: req.Id})
	if err != nil {
		return nil, err
	}
	return &v1.PraiseInfoDeleteRes{}, nil
}
