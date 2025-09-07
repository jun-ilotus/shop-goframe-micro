package interaction

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"

	"service/app/gateway-h5/api/interaction/v1"
	comment "service/app/interaction/api/comment_info/v1"
)

func (c *ControllerV1) CommentInfoGetList(ctx context.Context, req *v1.CommentInfoGetListReq) (res *v1.CommentInfoGetListRes, err error) {
	grpcReq := &comment.CommentInfoGetListReq{}
	if err = gconv.Struct(req, grpcReq); err != nil {
		return nil, err
	}

	grpcRes, err := c.CommentInfoClient.GetList(ctx, grpcReq)
	if err != nil {
		return nil, err
	}

	res = &v1.CommentInfoGetListRes{
		Page:  grpcRes.Data.Page,
		Size:  grpcRes.Data.Size,
		Total: grpcRes.Data.Total,
	}
	if err = gconv.Structs(grpcRes.Data.List, &res.List); err != nil {
		return nil, err
	}

	return res, nil
}
