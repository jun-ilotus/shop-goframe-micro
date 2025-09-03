package interaction

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"

	collection "service/app/interaction/api/collection_info/v1"

	"service/app/gateway-h5/api/interaction/v1"
)

func (c *ControllerV1) CollectionInfoGetList(ctx context.Context, req *v1.CollectionInfoGetListReq) (res *v1.CollectionInfoGetListRes, err error) {
	grpcReq := &collection.CollectionInfoGetListReq{}
	if err := gconv.Struct(req, grpcReq); err != nil {
		return nil, err
	}

	grpcRes, err := c.CollectionInfoClient.GetList(ctx, grpcReq)
	if err != nil {
		return nil, err
	}

	res = &v1.CollectionInfoGetListRes{
		Page:  grpcRes.Data.Page,
		Size:  grpcRes.Data.Size,
		Total: grpcRes.Data.Total,
	}

	if err := gconv.Struct(grpcRes.Data.List, &res.List); err != nil {
		return nil, err
	}

	return res, nil
}
