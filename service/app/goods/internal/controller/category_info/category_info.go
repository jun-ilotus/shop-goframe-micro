package category_info

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	v1 "service/app/goods/api/category_info/v1"
	pbentity "service/app/goods/api/pbentity"
	"service/app/goods/internal/consts"
	"service/app/goods/internal/dao"
	"service/app/goods/internal/model/entity"
	"service/app/goods/utility/goodsRedis"
	"service/utility"
	"time"

	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

type Controller struct {
	v1.UnimplementedCategoryInfoServer
}

func Register(s *grpcx.GrpcServer) {
	v1.RegisterCategoryInfoServer(s.Server, &Controller{})
}

func (*Controller) GetList(ctx context.Context, req *v1.CategoryInfoGetListReq) (res *v1.CategoryInfoGetListRes, err error) {
	response := &v1.CategoryInfoListResponse{
		List:  make([]*pbentity.CategoryInfo, 0),
		Page:  req.Page,
		Size:  req.Size,
		Total: 0,
	}

	countQuery := dao.CategoryInfo.Ctx(ctx)
	if req.Sort != 0 {
		countQuery = countQuery.Where("sort", req.Sort)
	}
	total, err := countQuery.Count()
	if err != nil {
		return &v1.CategoryInfoGetListRes{Data: response}, nil
	}
	response.Total = uint32(total)

	query := dao.CategoryInfo.Ctx(ctx).Page(int(req.Page), int(req.Size))
	if req.Size != 0 {
		query = query.Where("sort", req.Sort)
	}

	categoryRecords, err := query.All()
	if err != nil {
		return &v1.CategoryInfoGetListRes{Data: response}, nil
	}

	for _, record := range categoryRecords {
		var category entity.CategoryInfo
		if err := record.Struct(&category); err != nil {
			continue
		}
		var pbCategoryInfo pbentity.CategoryInfo
		if err := gconv.Struct(category, &pbCategoryInfo); err != nil {
			continue
		}
		pbCategoryInfo.CreatedAt = utility.SafeConvertTime(category.CreatedAt)
		pbCategoryInfo.UpdatedAt = utility.SafeConvertTime(category.UpdatedAt)
		pbCategoryInfo.DeletedAt = utility.SafeConvertTime(category.DeletedAt)

		response.List = append(response.List, &pbCategoryInfo)
	}

	return &v1.CategoryInfoGetListRes{Data: response}, nil
}

func (*Controller) GetAll(ctx context.Context, req *v1.CategoryInfoGetAllReq) (res *v1.CategoryInfoGetAllRes, err error) {
	// 尝试从 redis 中获取
	cachedData, err := goodsRedis.GetCategoryAll(ctx)
	if err != nil {
		g.Log().Infof(ctx, "Redis查询失败: %v", err)
	} else if !cachedData.IsNil() && !cachedData.IsEmpty() {
		var cachedRes v1.CategoryInfoGetAllRes
		if err := cachedData.Struct(&cachedRes); err != nil {
			g.Log().Errorf(ctx, "缓存数据反序列化失败：%v", err)
		} else {
			g.Log().Info(ctx, "category all data 缓存命中")
			return &cachedRes, nil
		}
	}

	response := &v1.CategoryInfoGetAllRes{
		List:  make([]*pbentity.CategoryInfo, 0),
		Total: 0,
	}
	total, err := dao.CategoryInfo.Ctx(ctx).Count()
	if err != nil {
		return response, err
	}
	response.Total = uint32(total)

	categoryRecords, err := dao.CategoryInfo.Ctx(ctx).All()
	if err != nil {
		return response, err
	}

	if total > 0 {
		response.List = make([]*pbentity.CategoryInfo, 0, total)
	}
	for _, record := range categoryRecords {
		var category entity.CategoryInfo
		if err := record.Struct(&category); err != nil {
			continue
		}
		var pbCategoryInfo pbentity.CategoryInfo
		if err := gconv.Struct(category, &pbCategoryInfo); err != nil {
			continue
		}

		pbCategoryInfo.CreatedAt = utility.SafeConvertTime(category.CreatedAt)
		pbCategoryInfo.UpdatedAt = utility.SafeConvertTime(category.UpdatedAt)
		pbCategoryInfo.DeletedAt = utility.SafeConvertTime(category.DeletedAt)

		response.List = append(response.List, &pbCategoryInfo)
	}

	// 为缓存设置操作设置100毫秒的超时时间 避免Redis操作耗时过长影响主业务流程
	ctxWithTimeout, cancel := context.WithTimeout(ctx, 100*time.Millisecond)
	defer cancel()

	// 设置缓存（使用一周的缓存时间）
	if err := goodsRedis.SetCategoryAll(ctxWithTimeout, response); err != nil {
		g.Log().Warningf(ctx, "设置分类全量数据缓存失败：%v", err)
	}

	return response, nil
}

func (*Controller) Create(ctx context.Context, req *v1.CategoryInfoCreateReq) (res *v1.CategoryInfoCreateRes, err error) {
	infoError := consts.InfoError(consts.CategoryInfo, consts.CreateFail)

	result, err := dao.CategoryInfo.Ctx(ctx).InsertAndGetId(req)
	if err != nil {
		g.Log().Errorf(ctx, "%v %v", infoError, err)
		return nil, gerror.WrapCode(gcode.CodeDbOperationError, err, infoError)
	}

	return &v1.CategoryInfoCreateRes{Id: uint32(result)}, nil
}

func (*Controller) Update(ctx context.Context, req *v1.CategoryInfoUpdateReq) (res *v1.CategoryInfoUpdateRes, err error) {
	infoError := consts.InfoError(consts.CategoryInfo, consts.UpdateFail)

	_, err = dao.CategoryInfo.Ctx(ctx).Where("id", req.Id).Update(req)
	if err != nil {
		g.Log().Errorf(ctx, "%v %v", infoError, err)
		return nil, gerror.WrapCode(gcode.CodeDbOperationError, err, infoError)
	}
	ctxWithTimeout, cancel := context.WithTimeout(ctx, 100*time.Millisecond)
	defer cancel()

	if err := goodsRedis.DeleteCategoryAll(ctxWithTimeout); err != nil {
		g.Log().Warningf(ctx, "删除分类全量数据缓存失败：%v", err)
	}

	return &v1.CategoryInfoUpdateRes{Id: req.Id}, nil
}

func (*Controller) Delete(ctx context.Context, req *v1.CategoryInfoDeleteReq) (res *v1.CategoryInfoDeleteRes, err error) {
	_, err = dao.CategoryInfo.Ctx(ctx).Where("id", req.Id).Delete()
	infoError := consts.InfoError(consts.CategoryInfo, consts.DeleteFail)
	if err != nil {
		g.Log().Errorf(ctx, "%v %v", infoError, err)
		return nil, gerror.WrapCode(gcode.CodeDbOperationError, err, infoError)
	}
	return &v1.CategoryInfoDeleteRes{}, nil
}
