package rotation_info

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	pbentity "service/app/banner/api/pbentity"
	v1 "service/app/banner/api/rotation_info/v1"
	"service/app/banner/internal/consts"
	"service/app/banner/internal/dao"
	"service/app/banner/internal/model/entity"
	"service/utility"

	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

type Controller struct {
	v1.UnimplementedRotationInfoServer
}

func Register(s *grpcx.GrpcServer) {
	v1.RegisterRotationInfoServer(s.Server, &Controller{})
}

func (*Controller) GetList(ctx context.Context, req *v1.RotationInfoGetListReq) (res *v1.RotationInfoGetListRes, err error) {
	infoError := consts.InfoError(consts.RotationInfo, consts.GetListFail)
	response := &v1.RotationInfoListResponse{
		List:  make([]*pbentity.RotationInfo, 0),
		Page:  req.Page,
		Size:  req.Size,
		Total: 0,
	}

	total, err := dao.RotationInfo.Ctx(ctx).Count()
	if err != nil {
		g.Log().Errorf(ctx, "%v %v", infoError, err)
		return nil, gerror.WrapCode(gcode.CodeDbOperationError, err, infoError)
	}
	response.Total = uint32(total)

	rotationRecords, err := dao.RotationInfo.Ctx(ctx).
		Order(utility.GetOrderBy(req.Sort)).
		Page(int(req.Page), int(req.Size)).
		All()
	if err != nil {
		g.Log().Errorf(ctx, "%v %v", infoError, err)
		return nil, gerror.WrapCode(gcode.CodeDbOperationError, err, infoError)
	}

	for _, record := range rotationRecords {
		var rotation entity.RotationInfo
		if err := record.Struct(&rotation); err != nil {
			continue
		}

		var pbRotation pbentity.RotationInfo
		if err := gconv.Struct(rotation, &pbRotation); err != nil {
			continue
		}

		pbRotation.CreatedAt = utility.SafeConvertTime(rotation.CreatedAt)
		pbRotation.UpdatedAt = utility.SafeConvertTime(rotation.UpdatedAt)
		pbRotation.DeletedAt = utility.SafeConvertTime(rotation.DeletedAt)
		response.List = append(response.List, &pbRotation)
	}
	return &v1.RotationInfoGetListRes{Data: response}, nil
}

func (*Controller) Create(ctx context.Context, req *v1.RotationInfoCreateReq) (res *v1.RotationInfoCreateRes, err error) {
	infoError := consts.InfoError(consts.RotationInfo, consts.CreateFail)
	id, err := dao.RotationInfo.Ctx(ctx).InsertAndGetId(req)
	if err != nil {
		g.Log().Errorf(ctx, "%v %v", infoError, err)
		return nil, gerror.WrapCode(gcode.CodeDbOperationError, err, infoError)
	}
	return &v1.RotationInfoCreateRes{Id: uint32(id)}, nil
}

func (*Controller) Update(ctx context.Context, req *v1.RotationInfoUpdateReq) (res *v1.RotationInfoUpdateRes, err error) {
	infoError := consts.InfoError(consts.RotationInfo, consts.UpdateFail)
	_, err = dao.RotationInfo.Ctx(ctx).Where("id", req.Id).Update(req)
	if err != nil {
		g.Log().Errorf(ctx, "%v %v", infoError, err)
		return nil, gerror.WrapCode(gcode.CodeDbOperationError, err, infoError)
	}
	return &v1.RotationInfoUpdateRes{Id: req.Id}, nil
}

func (*Controller) Delete(ctx context.Context, req *v1.RotationInfoDeleteReq) (res *v1.RotationInfoDeleteRes, err error) {
	infoError := consts.InfoError(consts.RotationInfo, consts.DeleteFail)
	_, err = dao.RotationInfo.Ctx(ctx).Where("id", req.Id).Delete()
	if err != nil {
		g.Log().Errorf(ctx, "%v %v", infoError, err)
		return nil, gerror.WrapCode(gcode.CodeDbOperationError, err, infoError)
	}
	return &v1.RotationInfoDeleteRes{}, nil
}
