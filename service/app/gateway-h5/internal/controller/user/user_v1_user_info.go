package user

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"

	"service/app/gateway-h5/api/user/v1"

	userinfo "service/app/user/api/user_info/v1"
)

func (c *ControllerV1) UserInfo(ctx context.Context, req *v1.UserInfoReq) (res *v1.UserInfoRes, err error) {
	grpcReq := &userinfo.UserInfoReq{}
	if err = gconv.Struct(req, grpcReq); err != nil {
		return nil, err
	}

	grpcRes, err := c.UserInfoClient.GetUserInfo(ctx, grpcReq)
	if err != nil {
		return nil, err
	}

	res = &v1.UserInfoRes{}
	if err = gconv.Struct(grpcRes, res); err != nil {
		return nil, err
	}
	return res, nil
}
