package position_info

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	pbentity "service/app/banner/api/pbentity"
	v1 "service/app/banner/api/position_info/v1"
	"service/app/banner/internal/consts"
	"service/app/banner/internal/dao"
	"service/app/banner/internal/model/entity"
	"service/utility"

	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

type Controller struct {
	v1.UnimplementedPositionInfoServer
}

func Register(s *grpcx.GrpcServer) {
	v1.RegisterPositionInfoServer(s.Server, &Controller{})
}

func (*Controller) GetList(ctx context.Context, req *v1.PositionInfoGetListReq) (res *v1.PositionInfoGetListRes, err error) {
	infoError := consts.InfoError(consts.PositionInfo, consts.GetListFail)
	response := &v1.PositionInfoListResponse{
		List:  make([]*pbentity.PositionInfo, 0),
		Page:  req.Page,
		Size:  req.Size,
		Total: 0,
	}

	total, err := dao.PositionInfo.Ctx(ctx).Count()
	if err != nil {
		g.Log().Errorf(ctx, "%v %v", infoError, err)
		return nil, gerror.WrapCode(gcode.CodeDbOperationError, err, infoError)
	}
	response.Total = uint32(total)

	// 查询当前页数据，sort值越小越靠前
	positionRecords, err := dao.PositionInfo.Ctx(ctx).
		Order(utility.GetOrderBy(req.Sort)).
		Page(int(req.Page), int(req.Size)).
		All()
	if err != nil {
		g.Log().Errorf(ctx, "%v %v", infoError, err)
		return nil, gerror.WrapCode(gcode.CodeDbOperationError, err, infoError)
	}

	for _, record := range positionRecords {
		var positionInfo entity.PositionInfo
		if err := record.Struct(&positionInfo); err != nil {
			continue
		}

		var pbPositionInfo pbentity.PositionInfo
		if err := gconv.Struct(positionInfo, &pbPositionInfo); err != nil {
			continue
		}

		pbPositionInfo.CreatedAt = utility.SafeConvertTime(positionInfo.CreatedAt)
		pbPositionInfo.UpdatedAt = utility.SafeConvertTime(positionInfo.UpdatedAt)
		pbPositionInfo.DeletedAt = utility.SafeConvertTime(positionInfo.DeletedAt)
		response.List = append(response.List, &pbPositionInfo)
	}
	return &v1.PositionInfoGetListRes{Data: response}, nil
}

func (*Controller) Create(ctx context.Context, req *v1.PositionInfoCreateReq) (res *v1.PositionInfoCreateRes, err error) {
	infoError := consts.InfoError(consts.PositionInfo, consts.CreateFail)
	id, err := dao.PositionInfo.Ctx(ctx).InsertAndGetId(req)
	if err != nil {
		g.Log().Errorf(ctx, "%v %v", infoError, err)
		return nil, gerror.WrapCode(gcode.CodeDbOperationError, err, infoError)
	}
	return &v1.PositionInfoCreateRes{Id: uint32(id)}, nil
}

func (*Controller) Update(ctx context.Context, req *v1.PositionInfoUpdateReq) (res *v1.PositionInfoUpdateRes, err error) {
	infoError := consts.InfoError(consts.PositionInfo, consts.UpdateFail)
	_, err = dao.PositionInfo.Ctx(ctx).Where("id", req.Id).Update(req)
	if err != nil {
		g.Log().Errorf(ctx, "%v %v", infoError, err)
		return nil, gerror.WrapCode(gcode.CodeDbOperationError, err, infoError)
	}
	return &v1.PositionInfoUpdateRes{Id: req.Id}, nil
}

func (*Controller) Delete(ctx context.Context, req *v1.PositionInfoDeleteReq) (res *v1.PositionInfoDeleteRes, err error) {
	infoError := consts.InfoError(consts.PositionInfo, consts.DeleteFail)
	_, err = dao.PositionInfo.Ctx(ctx).Where("id", req.Id).Delete()
	if err != nil {
		g.Log().Errorf(ctx, "%v %v", infoError, err)
		return nil, gerror.WrapCode(gcode.CodeDbOperationError, err, infoError)
	}
	return &v1.PositionInfoDeleteRes{}, nil
}
