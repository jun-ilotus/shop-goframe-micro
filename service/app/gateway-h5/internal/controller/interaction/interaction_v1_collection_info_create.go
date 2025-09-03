package interaction

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"

	"service/app/gateway-h5/api/interaction/v1"
	collection "service/app/interaction/api/collection_info/v1"
)

func (c *ControllerV1) CollectionInfoCreate(ctx context.Context, req *v1.CollectionInfoCreateReq) (res *v1.CollectionInfoCreateRes, err error) {
	grpcReq := &collection.CollectionInfoCreateReq{}
	if err = gconv.Struct(req, grpcReq); err != nil {
		return nil, err
	}

	grpcRes, err := c.CollectionInfoClient.Create(ctx, grpcReq)
	if err != nil {
		return nil, err
	}

	return &v1.CollectionInfoCreateRes{Id: grpcRes.Id}, nil
}
