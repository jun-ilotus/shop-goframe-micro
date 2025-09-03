package interaction

import (
	"context"

	collection "service/app/interaction/api/collection_info/v1"

	"service/app/gateway-h5/api/interaction/v1"
)

func (c *ControllerV1) CollectionInfoDelete(ctx context.Context, req *v1.CollectionInfoDeleteReq) (res *v1.CollectionInfoDeleteRes, err error) {
	_, err = c.CollectionInfoClient.Delete(ctx, &collection.CollectionInfoDeleteReq{Id: req.Id})
	if err != nil {
		return nil, err
	}
	return &v1.CollectionInfoDeleteRes{}, nil
}
