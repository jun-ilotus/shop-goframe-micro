package admin

import (
	"context"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/util/gconv"
	admin_info "service/app/admin/api/admin_info/v1"

	"service/app/gateway-admin/api/admin/v1"
)

func (c *ControllerV1) AdminInfoLogin(ctx context.Context, req *v1.AdminInfoLoginReq) (res *v1.AdminInfoLoginRes, err error) {
	grpcReq := &admin_info.AdminInfoLoginReq{}
	if err = gconv.Struct(req, grpcReq); err != nil {
		return nil, err
	}

	grpcRes, err := c.AdminInfoClient.Login(ctx, grpcReq)
	if err != nil {
		return nil, gerror.WrapCode(gcode.CodeInternalError, err, "登录失败")
	}

	res = &v1.AdminInfoLoginRes{}
	if err = gconv.Struct(grpcRes, res); err != nil {
		return nil, err
	}
	return res, nil
}
