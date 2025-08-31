package user

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"

	"service/app/gateway-h5/api/user/v1"

	consignee "service/app/user/api/consignee_info/v1"
)

func (c *ControllerV1) ConsigneeInfoCreate(ctx context.Context, req *v1.ConsigneeInfoCreateReq) (res *v1.ConsigneeInfoCreateRes, err error) {
	// 使用 gconv 自动转换结构体
	grpcReq := &consignee.ConsigneeInfoCreateReq{}
	if err := gconv.Struct(req, grpcReq); err != nil {
		return nil, err
	}
	// 调用 gRPC 服务
	grpcRes, err := c.ConsigneeInfoClient.Create(ctx, grpcReq)
	if err != nil {
		return nil, err
	}
	return &v1.ConsigneeInfoCreateRes{Id: grpcRes.Id}, nil
}
