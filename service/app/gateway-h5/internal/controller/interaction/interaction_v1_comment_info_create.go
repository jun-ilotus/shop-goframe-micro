package interaction

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"

	"service/app/gateway-h5/api/interaction/v1"
	comment "service/app/interaction/api/comment_info/v1"
)

func (c *ControllerV1) CommentInfoCreate(ctx context.Context, req *v1.CommentInfoCreateReq) (res *v1.CommentInfoCreateRes, err error) {
	grpcReq := &comment.CommentInfoCreateReq{}
	if err = gconv.Struct(req, grpcReq); err != nil {
		return nil, err
	}

	grpcRes, err := c.CommentInfoClient.Create(ctx, grpcReq)
	if err != nil {
		return nil, err
	}

	return &v1.CommentInfoCreateRes{Id: grpcRes.Id}, nil
}
