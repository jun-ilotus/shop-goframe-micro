package user

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"service/app/gateway-h5/api/user/v1"
)

func (c *ControllerV1) ConsigneeInfoGetList(ctx context.Context, req *v1.ConsigneeInfoGetListReq) (res *v1.ConsigneeInfoGetListRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
