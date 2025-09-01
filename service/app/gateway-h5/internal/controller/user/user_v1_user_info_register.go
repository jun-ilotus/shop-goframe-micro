package user

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"

	"service/app/gateway-h5/api/user/v1"

	userinfo "service/app/user/api/user_info/v1"
)

func (c *ControllerV1) UserInfoRegister(ctx context.Context, req *v1.UserInfoRegisterReq) (res *v1.UserInfoRegisterRes, err error) {
	grpcReq := &userinfo.UserInfoRegisterReq{}
	if err = gconv.Struct(req, grpcReq); err != nil {
		return nil, err
	}

	grpcRes, err := c.UserInfoClient.Register(ctx, grpcReq)

	if err != nil {
		return nil, err
	}
	return &v1.UserInfoRegisterRes{Id: grpcRes.Id}, nil
}
