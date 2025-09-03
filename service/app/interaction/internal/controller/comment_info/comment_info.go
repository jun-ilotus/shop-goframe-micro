package comment_info

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	v1 "service/app/interaction/api/comment_info/v1"
	pbentity "service/app/interaction/api/pbentity"
	"service/app/interaction/internal/consts"
	"service/app/interaction/internal/dao"
	"service/app/interaction/internal/model/entity"
	"service/utility"

	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

type Controller struct {
	v1.UnimplementedCommentInfoServer
}

func Register(s *grpcx.GrpcServer) {
	v1.RegisterCommentInfoServer(s.Server, &Controller{})
}

func (*Controller) GetList(ctx context.Context, req *v1.CommentInfoGetListReq) (res *v1.CommentInfoGetListRes, err error) {
	response := &v1.CommentInfoGetListResponse{
		List:  make([]*pbentity.CommentInfo, 0),
		Page:  req.Page,
		Size:  req.Size,
		Total: 0,
	}
	infoError := consts.InfoError(consts.CommentInfo, consts.GetListFail)
	total, err := dao.CommentInfo.Ctx(ctx).Count()
	if err != nil {
		g.Log().Errorf(ctx, "%v %v", infoError, err)
		return nil, gerror.WrapCode(gcode.CodeDbOperationError, err, infoError)
	}
	response.Total = uint32(total)

	CommentRecords, err := dao.CommentInfo.Ctx(ctx).Page(int(req.Page), int(req.Size)).All()
	if err != nil {
		g.Log().Errorf(ctx, "%v %v", infoError, err)
		return nil, gerror.WrapCode(gcode.CodeDbOperationError, err, infoError)
	}

	for _, record := range CommentRecords {
		var comment entity.CommentInfo
		if err := record.Struct(&comment); err != nil {
			continue
		}
		var pbComment pbentity.CommentInfo
		if err := gconv.Struct(comment, &pbComment); err != nil {
			continue
		}
		pbComment.CreatedAt = utility.SafeConvertTime(comment.CreatedAt)
		pbComment.UpdatedAt = utility.SafeConvertTime(comment.UpdatedAt)

		response.List = append(response.List, &pbComment)
	}
	return &v1.CommentInfoGetListRes{Data: response}, nil
}

func (*Controller) Create(ctx context.Context, req *v1.CommentInfoCreateReq) (res *v1.CommentInfoCreateRes, err error) {
	infoError := consts.InfoError(consts.CommentInfo, consts.CreateFail)

	var commentInfo *entity.CommentInfo
	if err = gconv.Struct(req, &commentInfo); err != nil {
		return nil, err
	}
	result, err := dao.CommentInfo.Ctx(ctx).InsertAndGetId(req)
	if err != nil {
		g.Log().Errorf(ctx, "%v %v", infoError, err)
		return nil, gerror.WrapCode(gcode.CodeDbOperationError, err, infoError)
	}

	return &v1.CommentInfoCreateRes{Id: uint32(result)}, nil
}

func (*Controller) Delete(ctx context.Context, req *v1.CommentInfoDeleteReq) (res *v1.CommentInfoDeleteRes, err error) {
	infoError := consts.InfoError(consts.CommentInfo, consts.DeleteFail)
	_, err = dao.CommentInfo.Ctx(ctx).Where("id", req.Id).Delete()
	if err != nil {
		g.Log().Errorf(ctx, "%v %v", infoError, err)
		return nil, gerror.WrapCode(gcode.CodeDbOperationError, err, infoError)
	}
	return &v1.CommentInfoDeleteRes{}, nil
}
