// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package user

import (
	"context"

	"service/app/gateway-h5/api/user/v1"
)

type IUserV1 interface {
	ConsigneeInfoCreate(ctx context.Context, req *v1.ConsigneeInfoCreateReq) (res *v1.ConsigneeInfoCreateRes, err error)
	ConsigneeInfoGetList(ctx context.Context, req *v1.ConsigneeInfoGetListReq) (res *v1.ConsigneeInfoGetListRes, err error)
	ConsigneeInfoUpdate(ctx context.Context, req *v1.ConsigneeInfoUpdateReq) (res *v1.ConsigneeInfoUpdateRes, err error)
	ConsigneeInfoDelete(ctx context.Context, req *v1.ConsigneeInfoDeleteReq) (res *v1.ConsigneeInfoDeleteRes, err error)
	UserInfoLogin(ctx context.Context, req *v1.UserInfoLoginReq) (res *v1.UserInfoLoginRes, err error)
	UserInfoRegister(ctx context.Context, req *v1.UserInfoRegisterReq) (res *v1.UserInfoRegisterRes, err error)
	UserInfo(ctx context.Context, req *v1.UserInfoReq) (res *v1.UserInfoRes, err error)
	UserInfoUpdatePassword(ctx context.Context, req *v1.UserInfoUpdatePasswordReq) (res *v1.UserInfoUpdatePasswordRes, err error)
}
