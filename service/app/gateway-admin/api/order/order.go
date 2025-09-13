// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package order

import (
	"context"

	"service/app/gateway-admin/api/order/v1"
)

type IOrderV1 interface {
	OrderInfoGetList(ctx context.Context, req *v1.OrderInfoGetListReq) (res *v1.OrderInfoGetListRes, err error)
	OrderInfoGetDetail(ctx context.Context, req *v1.OrderInfoGetDetailReq) (res *v1.OrderInfoGetDetailRes, err error)
}
