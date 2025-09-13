package order

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	order_info "service/app/order/api/order_info/v1"

	"service/app/gateway-h5/api/order/v1"
)

func (c *ControllerV1) OrderInfoGetList(ctx context.Context, req *v1.OrderInfoGetListReq) (res *v1.OrderInfoGetListRes, err error) {
	// 使用 gconv 自动转换结构体
	grpcReq := &order_info.OrderInfoGetListReq{}
	if err := gconv.Struct(req, grpcReq); err != nil {
		return nil, err
	}

	// 调用gRPC服务
	grpcRes, err := c.OrderInfoClient.GetList(ctx, grpcReq)

	if err != nil {
		return nil, err
	}

	// 转换响应
	res = &v1.OrderInfoGetListRes{
		Page:  grpcRes.Data.Page,
		Size:  grpcRes.Data.Size,
		Total: grpcRes.Data.Total,
	}

	// 批量转换列表项
	if err := gconv.Structs(grpcRes.Data.List, &res.List); err != nil {
		return nil, err
	}

	return res, nil
}
