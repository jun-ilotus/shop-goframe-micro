package user

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	consignee "service/app/user/api/consignee_info/v1"

	"service/app/gateway-h5/api/user/v1"
)

func (c *ControllerV1) ConsigneeInfoUpdate(ctx context.Context, req *v1.ConsigneeInfoUpdateReq) (res *v1.ConsigneeInfoUpdateRes, err error) {
	grpcReq := &consignee.ConsigneeInfoUpdateReq{}
	if err := gconv.Struct(req, grpcReq); err != nil {
		return nil, err
	}

	grpcRes, err := c.ConsigneeInfoClient.Update(ctx, grpcReq)
	if err != nil {
		return nil, err
	}

	return &v1.ConsigneeInfoUpdateRes{Id: grpcRes.Id}, nil
}
