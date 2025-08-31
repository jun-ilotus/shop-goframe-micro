package user

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"service/app/gateway-h5/api/user/v1"
)

func (c *ControllerV1) ConsigneeInfoCreate(ctx context.Context, req *v1.ConsigneeInfoCreateReq) (res *v1.ConsigneeInfoCreateRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
