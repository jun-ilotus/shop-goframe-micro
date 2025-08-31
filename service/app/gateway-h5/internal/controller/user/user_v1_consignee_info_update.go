package user

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"service/app/gateway-h5/api/user/v1"
)

func (c *ControllerV1) ConsigneeInfoUpdate(ctx context.Context, req *v1.ConsigneeInfoUpdateReq) (res *v1.ConsigneeInfoUpdateRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
