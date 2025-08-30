package user_info

import (
	"context"
	v1 "service/app/user/api/user_info/v1"

	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

type Controller struct {
	v1.UnimplementedUserInfoServer
}

func Register(s *grpcx.GrpcServer) {
	v1.RegisterUserInfoServer(s.Server, &Controller{})
}

func (*Controller) Login(ctx context.Context, req *v1.UserInfoLoginReq) (res *v1.UserInfoLoginRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) UpdatePassword(ctx context.Context, req *v1.UserInfoUpdatePasswordReq) (res *v1.UserInfoUpdatePasswordRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) GetUserInfo(ctx context.Context, req *v1.UserInfoReq) (res *v1.UserInfoRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
