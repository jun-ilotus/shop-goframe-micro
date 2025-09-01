package user

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"service/app/gateway-h5/api/user/v1"
	userinfo "service/app/user/api/user_info/v1"
)

func (c *ControllerV1) UserInfoLogin(ctx context.Context, req *v1.UserInfoLoginReq) (res *v1.UserInfoLoginRes, err error) {
	grpcReq := &userinfo.UserInfoLoginReq{}
	if err := gconv.Struct(req, grpcReq); err != nil {
		return nil, err
	}

	grpcRes, err := c.UserInfoClient.Login(ctx, grpcReq)
	if err != nil {
		// 这里可以根据gRPC返回的错误码转换成本地错误码
		// 例如，如果gRPC返回的是用户不存在，可以转换为CodeNotFound
		return nil, gerror.WrapCode(gcode.CodeInternalError, err, "登录失败")
	}

	res = &v1.UserInfoLoginRes{}
	if err := gconv.Struct(grpcRes, res); err != nil {
		return nil, err
	}

	return res, nil
}
