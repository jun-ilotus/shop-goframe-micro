package collection_info

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	v1 "service/app/interaction/api/collection_info/v1"
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
	v1.UnimplementedCollectionInfoServer
}

func Register(s *grpcx.GrpcServer) {
	v1.RegisterCollectionInfoServer(s.Server, &Controller{})
}

func (*Controller) GetList(ctx context.Context, req *v1.CollectionInfoGetListReq) (res *v1.CollectionInfoGetListRes, err error) {
	response := &v1.CollectionInfoGetListResponse{
		List:  make([]*pbentity.CollectionInfo, 0),
		Page:  req.Page,
		Size:  req.Size,
		Total: 0,
	}
	infoError := consts.InfoError(consts.CollectionInfo, consts.GetListFail)
	total, err := dao.CollectionInfo.Ctx(ctx).Count()
	if err != nil {
		g.Log().Errorf(ctx, "%v %v", infoError, err)
		return nil, gerror.WrapCode(gcode.CodeDbOperationError, err, infoError)
	}
	response.Total = uint32(total)

	collectionRecords, err := dao.CollectionInfo.Ctx(ctx).Page(int(req.Page), int(req.Size)).All()
	if err != nil {
		g.Log().Errorf(ctx, "%v %v", infoError, err)
		return nil, gerror.WrapCode(gcode.CodeDbOperationError, err, infoError)
	}

	for _, record := range collectionRecords {
		var collection entity.CollectionInfo
		if err := record.Struct(&collection); err != nil {
			continue
		}

		var pbCollection pbentity.CollectionInfo
		if err := gconv.Struct(collection, &pbCollection); err != nil {
			continue
		}

		pbCollection.CreatedAt = utility.SafeConvertTime(collection.CreatedAt)
		pbCollection.UpdatedAt = utility.SafeConvertTime(collection.UpdatedAt)

		response.List = append(response.List, &pbCollection)
	}
	return &v1.CollectionInfoGetListRes{Data: response}, nil
}

func (*Controller) Create(ctx context.Context, req *v1.CollectionInfoCreateReq) (res *v1.CollectionInfoCreateRes, err error) {
	var collectionInfo *entity.CollectionInfo
	if err := gconv.Struct(req, &collectionInfo); err != nil {
		return nil, err
	}

	infoError := consts.InfoError(consts.CollectionInfo, consts.CreateFail)

	result, err := dao.CollectionInfo.Ctx(ctx).InsertAndGetId(req)
	if err != nil {
		g.Log().Errorf(ctx, "%v %v", infoError, err)
		return nil, gerror.WrapCode(gcode.CodeDbOperationError, err, infoError)
	}

	return &v1.CollectionInfoCreateRes{Id: uint32(result)}, nil
}

func (*Controller) Delete(ctx context.Context, req *v1.CollectionInfoDeleteReq) (res *v1.CollectionInfoDeleteRes, err error) {
	_, err = dao.CollectionInfo.Ctx(ctx).Where("id", req.Id).Delete()
	infoError := consts.InfoError(consts.CollectionInfo, consts.DeleteFail)
	if err != nil {
		g.Log().Errorf(ctx, "%v %v", infoError, err)
		return nil, gerror.WrapCode(gcode.CodeDbOperationError, err, infoError)
	}
	return &v1.CollectionInfoDeleteRes{}, nil
}
