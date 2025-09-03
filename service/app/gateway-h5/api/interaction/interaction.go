// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package interaction

import (
	"context"

	"service/app/gateway-h5/api/interaction/v1"
)

type IInteractionV1 interface {
	CollectionInfoCreate(ctx context.Context, req *v1.CollectionInfoCreateReq) (res *v1.CollectionInfoCreateRes, err error)
	CollectionInfoDelete(ctx context.Context, req *v1.CollectionInfoDeleteReq) (res *v1.CollectionInfoDeleteRes, err error)
	CollectionInfoGetList(ctx context.Context, req *v1.CollectionInfoGetListReq) (res *v1.CollectionInfoGetListRes, err error)
	CommentInfoGetList(ctx context.Context, req *v1.CommentInfoGetListReq) (res *v1.CommentInfoGetListRes, err error)
	CommentInfoCreate(ctx context.Context, req *v1.CommentInfoCreateReq) (res *v1.CommentInfoCreateRes, err error)
	CommentInfoDelete(ctx context.Context, req *v1.CommentInfoDeleteReq) (res *v1.CommentInfoDeleteRes, err error)
	PraiseInfoCreate(ctx context.Context, req *v1.PraiseInfoCreateReq) (res *v1.PraiseInfoCreateRes, err error)
	PraiseInfoDelete(ctx context.Context, req *v1.PraiseInfoDeleteReq) (res *v1.PraiseInfoDeleteRes, err error)
	PraiseInfoGetList(ctx context.Context, req *v1.PraiseInfoGetListReq) (res *v1.PraiseInfoGetListRes, err error)
}
