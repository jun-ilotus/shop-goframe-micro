package interaction

import (
	"context"

	"service/app/gateway-h5/api/interaction/v1"
	comment "service/app/interaction/api/comment_info/v1"
)

func (c *ControllerV1) CommentInfoDelete(ctx context.Context, req *v1.CommentInfoDeleteReq) (res *v1.CommentInfoDeleteRes, err error) {
	_, err = c.CommentInfoClient.Delete(ctx, &comment.CommentInfoDeleteReq{Id: req.Id})
	if err != nil {
		return nil, err
	}
	return &v1.CommentInfoDeleteRes{}, nil
}
