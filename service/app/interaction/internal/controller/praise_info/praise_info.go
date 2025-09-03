package praise_info

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	pbentity "service/app/interaction/api/pbentity"
	v1 "service/app/interaction/api/praise_info/v1"
	"service/app/interaction/internal/consts"
	"service/app/interaction/internal/dao"
	"service/app/interaction/internal/model/entity"
	"service/utility"

	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

type Controller struct {
	v1.UnimplementedPraiseInfoServer
}

func Register(s *grpcx.GrpcServer) {
	v1.RegisterPraiseInfoServer(s.Server, &Controller{})
}

func (*Controller) GetList(ctx context.Context, req *v1.PraiseInfoGetListReq) (res *v1.PraiseInfoGetListRes, err error) {
	response := &v1.PraiseInfoGetListResponse{
		List:  make([]*pbentity.PraiseInfo, 0),
		Page:  req.Page,
		Size:  req.Size,
		Total: 0,
	}
	infoError := consts.InfoError(consts.PraiseInfo, consts.GetListFail)
	total, err := dao.PraiseInfo.Ctx(ctx).Count()
	if err != nil {
		g.Log().Errorf(ctx, "%v %v", infoError, err)
		return nil, gerror.WrapCode(gcode.CodeDbOperationError, err, infoError)
	}
	response.Total = uint32(total)

	PraiseRecords, err := dao.PraiseInfo.Ctx(ctx).Page(int(req.Page), int(req.Size)).All()
	if err != nil {
		g.Log().Errorf(ctx, "%v %v", infoError, err)
		return nil, gerror.WrapCode(gcode.CodeDbOperationError, err, infoError)
	}

	for _, record := range PraiseRecords {
		var praise entity.PraiseInfo
		if err = record.Struct(&praise); err != nil {
			continue
		}

		var pbPraiseInfo pbentity.PraiseInfo
		if err = gconv.Struct(praise, &pbPraiseInfo); err != nil {
			continue
		}

		pbPraiseInfo.CreatedAt = utility.SafeConvertTime(praise.CreatedAt)
		pbPraiseInfo.UpdatedAt = utility.SafeConvertTime(praise.UpdatedAt)

		response.List = append(response.List, &pbPraiseInfo)
	}

	return &v1.PraiseInfoGetListRes{Data: response}, nil
}

func (*Controller) Create(ctx context.Context, req *v1.PraiseInfoCreateReq) (res *v1.PraiseInfoCreateRes, err error) {
	infoError := consts.InfoError(consts.CollectionInfo, consts.CreateFail)
	var praiseInfo *entity.PraiseInfo
	if err = gconv.Struct(req, &praiseInfo); err != nil {
		return nil, err
	}

	result, err := dao.PraiseInfo.Ctx(ctx).InsertAndGetId(req)
	if err != nil {
		g.Log().Errorf(ctx, "%v %v", infoError, err)
		return nil, gerror.WrapCode(gcode.CodeDbOperationError, err, infoError)
	}

	return &v1.PraiseInfoCreateRes{Id: uint32(result)}, nil
}

func (*Controller) Delete(ctx context.Context, req *v1.PraiseInfoDeleteReq) (res *v1.PraiseInfoDeleteRes, err error) {
	infoError := consts.InfoError(consts.CollectionInfo, consts.DeleteFail)
	_, err = dao.PraiseInfo.Ctx(ctx).Where("id", req.Id).Delete()
	if err != nil {
		g.Log().Errorf(ctx, "%v %v", infoError, err)
		return nil, gerror.WrapCode(gcode.CodeDbOperationError, err, infoError)
	}
	return &v1.PraiseInfoDeleteRes{}, nil
}
