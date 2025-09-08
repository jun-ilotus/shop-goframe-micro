package goods_info

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	v1 "service/app/goods/api/goods_info/v1"
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
	v1.UnimplementedGoodsInfoServer
}

func Register(s *grpcx.GrpcServer) {
	v1.RegisterGoodsInfoServer(s.Server, &Controller{})
}

func (*Controller) GetList(ctx context.Context, req *v1.GoodsInfoGetListReq) (res *v1.GoodsInfoGetListRes, err error) {
	response := &v1.GoodsInfoListResponse{
		List:  make([]*pbentity.GoodsInfo, 0),
		Page:  req.Page,
		Size:  req.Size,
		Total: 0,
	}
	infoError := consts.InfoError(consts.GoodsInfo, consts.GetListFail)
	total, err := dao.GoodsInfo.Ctx(ctx).Count()
	if err != nil {
		g.Log().Errorf(ctx, "%v %v", consts.GoodsInfo, consts.GetListFail)
		return nil, gerror.WrapCode(gcode.CodeDbOperationError, err, infoError)
	}
	response.Total = uint32(total)

	goodsRecord, err := dao.GoodsInfo.Ctx(ctx).Page(int(req.Page), int(req.Size)).All()
	if err != nil {
		g.Log().Errorf(ctx, "%v %v", consts.GoodsInfo, consts.GetListFail)
		return nil, gerror.WrapCode(gcode.CodeDbOperationError, err, infoError)
	}

	for _, record := range goodsRecord {
		var goods entity.GoodsInfo
		if err := record.Struct(&goods); err != nil {
			continue
		}

		var pbGoods pbentity.GoodsInfo
		if err := gconv.Struct(goods, &pbGoods); err != nil {
			continue
		}

		pbGoods.CreatedAt = utility.SafeConvertTime(goods.CreatedAt)
		pbGoods.UpdatedAt = utility.SafeConvertTime(goods.UpdatedAt)

		response.List = append(response.List, &pbGoods)
	}
	return &v1.GoodsInfoGetListRes{Data: response}, nil
}

func (*Controller) GetDetail(ctx context.Context, req *v1.GoodsInfoGetDetailReq) (res *v1.GoodsInfoGetDetailRes, err error) {
	infoError := consts.InfoError(consts.GoodsInfo, consts.GetDetailFail)

	detail, err := goodsRedis.GetGoodsDetail(ctx, req.Id)
	if err != nil {
		g.Log().Infof(ctx, "Redis查询失败：%v", err)
	} else if !detail.IsNil() {
		// 检查是否为空缓存标记
		if detail.String() == "__EMPTY__" {
			g.Log().Info(ctx, "空缓存命中，防止缓存穿透")
			return nil, gerror.New("商品不存在")
		}
		// 缓存命中，反序列化数据
		var cacheRes v1.GoodsInfoGetDetailRes
		if err := detail.Struct(&cacheRes); err != nil {
			g.Log().Errorf(ctx, "goods detail 缓存命中")
			return &cacheRes, nil
		}
	}
	// 缓存未命中，查询数据库

	record, err := dao.GoodsInfo.Ctx(ctx).Where("id", req.Id).One()
	if err != nil {
		g.Log().Errorf(ctx, "%v %v", infoError, err)
		return nil, gerror.WrapCode(gcode.CodeDbOperationError, err, infoError)
	}
	if record.IsEmpty() {
		g.Log().Errorf(ctx, "%v %v", infoError+"查询商品不存在", err)
		// 设置空缓存防止缓存穿透
		_ = goodsRedis.SetEmptyGoodsDetail(ctx, req.Id)
		return nil, gerror.WrapCode(gcode.CodeDbOperationError, err, infoError+"查询商品不存在")
	}

	var goods entity.GoodsInfo
	if err := record.Struct(&goods); err != nil {
		return nil, gerror.WrapCode(gcode.CodeInternalError, err, "数据转换失败")
	}

	var pbGoods pbentity.GoodsInfo
	if err := gconv.Struct(goods, &pbGoods); err != nil {
		return nil, gerror.WrapCode(gcode.CodeInternalError, err, "数据转换失败")
	}

	pbGoods.CreatedAt = utility.SafeConvertTime(goods.CreatedAt)
	pbGoods.UpdatedAt = utility.SafeConvertTime(goods.UpdatedAt)

	res = &v1.GoodsInfoGetDetailRes{Data: &pbGoods}

	// 同步设置缓存（使用较短的超时时间避免阻塞）
	ctxWithTimeout, cancel := context.WithTimeout(ctx, 100*time.Millisecond)
	defer cancel()

	// 缓存未命中，添加缓存
	if err := goodsRedis.SetGoodsDetail(ctxWithTimeout, pbGoods.Id, res); err != nil {
		g.Log().Warningf(ctx, "设置商品详情缓存失败：%v", err)
		// 不返回错误，因为主业务已成功
	}

	return res, nil
}

func (*Controller) Create(ctx context.Context, req *v1.GoodsInfoCreateReq) (res *v1.GoodsInfoCreateRes, err error) {
	var goods *entity.GoodsInfo
	if err := gconv.Struct(req, &goods); err != nil {
		return nil, err
	}
	infoError := consts.InfoError(consts.GoodsInfo, consts.CreateFail)
	result, err := dao.GoodsInfo.Ctx(ctx).InsertAndGetId(req)
	if err != nil {
		g.Log().Errorf(ctx, "%v %v", infoError, err)
		return nil, gerror.WrapCode(gcode.CodeDbOperationError, err, infoError)
	}
	return &v1.GoodsInfoCreateRes{Id: uint32(int(result))}, nil
}

func (*Controller) Update(ctx context.Context, req *v1.GoodsInfoUpdateReq) (res *v1.GoodsInfoUpdateRes, err error) {
	var goods *entity.GoodsInfo
	if err := gconv.Struct(req, &goods); err != nil {
		return nil, err
	}
	infoError := consts.InfoError(consts.GoodsInfo, consts.UpdateFail)
	_, err = dao.GoodsInfo.Ctx(ctx).Where("id", req.Id).Update(req)
	if err != nil {
		g.Log().Errorf(ctx, "%v %v", infoError, err)
		return nil, gerror.WrapCode(gcode.CodeDbOperationError, err, infoError)
	}
	return &v1.GoodsInfoUpdateRes{Id: req.Id}, nil
}

func (*Controller) Delete(ctx context.Context, req *v1.GoodsInfoDeleteReq) (res *v1.GoodsInfoDeleteRes, err error) {
	_, err = dao.GoodsInfo.Ctx(ctx).Where("id", req.Id).Delete()
	infoError := consts.InfoError(consts.GoodsInfo, consts.DeleteFail)
	if err != nil {
		g.Log().Errorf(ctx, "%v %v", infoError, err)
		return nil, gerror.WrapCode(gcode.CodeDbOperationError, err, infoError)
	}
	return &v1.GoodsInfoDeleteRes{}, nil
}
