package user

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	consignee "service/app/user/api/consignee_info/v1"

	"service/app/gateway-h5/api/user/v1"
)

func (c *ControllerV1) ConsigneeInfoGetList(ctx context.Context, req *v1.ConsigneeInfoGetListReq) (res *v1.ConsigneeInfoGetListRes, err error) {
	grpcReq := &consignee.ConsigneeInfoGetListReq{}
	if err := gconv.Struct(req, grpcReq); err != nil {
		return nil, err
	}

	grpcRes, err := c.ConsigneeInfoClient.GetList(ctx, grpcReq)
	if err != nil {
		return nil, err
	}

	// 转换响应
	res = &v1.ConsigneeInfoGetListRes{
		Page:  grpcRes.Data.Page,
		Size:  grpcRes.Data.Size,
		Total: grpcRes.Data.Total,
	}
	// 批量转换列表项
	if err := gconv.Struct(grpcRes.Data.List, &res.List); err != nil {
		return nil, err
	}
	return res, nil
}
