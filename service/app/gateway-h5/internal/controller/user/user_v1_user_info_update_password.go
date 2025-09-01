package user

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"

	"service/app/gateway-h5/api/user/v1"
	userinfo "service/app/user/api/user_info/v1"
)

func (c *ControllerV1) UserInfoUpdatePassword(ctx context.Context, req *v1.UserInfoUpdatePasswordReq) (res *v1.UserInfoUpdatePasswordRes, err error) {
	grpcReq := &userinfo.UserInfoUpdatePasswordReq{}
	if err = gconv.Struct(req, grpcReq); err != nil {
		return nil, err
	}

	grpcRes, err := c.UserInfoClient.UpdatePassword(ctx, grpcReq)
	if err != nil {
		return nil, err
	}

	return &v1.UserInfoUpdatePasswordRes{Id: grpcRes.Id}, nil
}
