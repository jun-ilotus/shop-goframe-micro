package goods_images

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	v1 "service/app/goods/api/goods_images/v1"
	pbentity "service/app/goods/api/pbentity"
	"service/app/goods/internal/consts"
	"service/app/goods/internal/dao"
	"service/app/goods/internal/model/entity"

	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

type Controller struct {
	v1.UnimplementedGoodsImagesServer
}

func Register(s *grpcx.GrpcServer) {
	v1.RegisterGoodsImagesServer(s.Server, &Controller{})
}

func (*Controller) GetList(ctx context.Context, req *v1.GoodsImagesGetListReq) (res *v1.GoodsImagesGetListRes, err error) {
	reponse := &v1.GoodsImagesListResponse{
		List:  make([]*pbentity.GoodsImages, 0),
		Page:  req.Page,
		Size:  req.Size,
		Total: 0,
	}
	infoError := consts.InfoError(consts.GoodsImages, consts.GetListFail)
	total, err := dao.GoodsImages.Ctx(ctx).Count()
	if err != nil {
		g.Log().Errorf(ctx, "%v %v", infoError, err)
		return nil, gerror.WrapCode(gcode.CodeDbOperationError, err, infoError)
	}
	reponse.Total = uint32(total)

	goodsRecords, err := dao.GoodsImages.Ctx(ctx).Page(int(req.Page), int(req.Size)).All()
	if err != nil {
		g.Log().Errorf(ctx, "%v %v", infoError, err)
		return nil, gerror.WrapCode(gcode.CodeDbOperationError, err, infoError)
	}

	for _, record := range goodsRecords {
		var goods entity.GoodsImages
		if err := record.Struct(&goods); err != nil {
			continue
		}
		var pbGoods pbentity.GoodsImages
		if err := gconv.Struct(goods, &pbGoods); err != nil {
			continue
		}
		reponse.List = append(reponse.List, &pbGoods)
	}
	return &v1.GoodsImagesGetListRes{Data: reponse}, nil
}

func (*Controller) Create(ctx context.Context, req *v1.GoodsImagesCreateReq) (res *v1.GoodsImagesCreateRes, err error) {
	var goodsImages *entity.GoodsImages
	if err := gconv.Struct(req, &goodsImages); err != nil {
		return nil, err
	}
	infoError := consts.InfoError(consts.GoodsImages, consts.CreateFail)
	result, err := dao.GoodsImages.Ctx(ctx).InsertAndGetId(req)
	if err != nil {
		g.Log().Errorf(ctx, "%v %v", infoError, err)
		return nil, gerror.WrapCode(gcode.CodeDbOperationError, err, infoError)
	}
	return &v1.GoodsImagesCreateRes{Id: uint32(result)}, nil
}

func (*Controller) Delete(ctx context.Context, req *v1.GoodsImagesDeleteReq) (res *v1.GoodsImagesDeleteRes, err error) {
	infoError := consts.InfoError(consts.GoodsImages, consts.DeleteFail)
	_, err = dao.GoodsImages.Ctx(ctx).Where("id", req.Id).Delete()
	if err != nil {
		g.Log().Errorf(ctx, "%v %v", infoError, err)
		return nil, gerror.WrapCode(gcode.CodeDbOperationError, err, infoError)
	}
	return &v1.GoodsImagesDeleteRes{}, nil
}
