// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package admin

import (
	"context"

	"service/app/gateway-admin/api/admin/v1"
)

type IAdminV1 interface {
	AdminInfoLogin(ctx context.Context, req *v1.AdminInfoLoginReq) (res *v1.AdminInfoLoginRes, err error)
	AdminInfoRegister(ctx context.Context, req *v1.AdminInfoRegisterReq) (res *v1.AdminInfoRegisterRes, err error)
}
