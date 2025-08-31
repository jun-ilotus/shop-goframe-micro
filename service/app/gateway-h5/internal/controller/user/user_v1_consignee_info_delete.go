package user

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"service/app/gateway-h5/api/user/v1"
)

func (c *ControllerV1) ConsigneeInfoDelete(ctx context.Context, req *v1.ConsigneeInfoDeleteReq) (res *v1.ConsigneeInfoDeleteRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
