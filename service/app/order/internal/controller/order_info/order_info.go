package order_info

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	v1 "service/app/order/api/order_info/v1"
	"service/app/order/api/pbentity"
	"service/app/order/internal/consts"

	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	order_info "service/app/order/internal/logic/order_info"
)

type Controller struct {
	v1.UnimplementedOrderInfoServer
}

func Register(s *grpcx.GrpcServer) {
	v1.RegisterOrderInfoServer(s.Server, &Controller{})
}

func (*Controller) Create(ctx context.Context, req *v1.OrderInfoCreateReq) (res *v1.OrderInfoCreateRes, err error) {
	infoError := consts.InfoError(consts.OrderInfo, consts.CreateFail)
	orderId, err := order_info.Create(ctx, req)
	if err != nil {
		g.Log().Errorf(ctx, "%v %v", infoError, err)
		return nil, gerror.WrapCode(gcode.CodeDbOperationError, err, infoError)
	}
	return &v1.OrderInfoCreateRes{Id: uint32(orderId)}, nil
}

func (*Controller) GetDetail(ctx context.Context, req *v1.OrderInfoGetDetailReq) (res *v1.OrderInfoGetDetailRes, err error) {
	infoError := consts.InfoError(consts.OrderInfo, consts.GetDetailFile)
	pbOrder, pbGoodsList, err := order_info.GetDetail(ctx, req.Id)
	if err != nil {
		g.Log().Errorf(ctx, "%v %v", infoError, err)
		return nil, gerror.WrapCode(gcode.CodeDbOperationError, err, infoError)
	}
	return &v1.OrderInfoGetDetailRes{
		OrderInfo:       pbOrder,
		OrderGoodsInfos: pbGoodsList,
	}, nil
}

func (*Controller) GetList(ctx context.Context, req *v1.OrderInfoGetListReq) (res *v1.OrderInfoGetListRes, err error) {
	response := &v1.OrderInfoListResponse{
		List:  make([]*pbentity.OrderInfo, 0),
		Page:  req.Page,
		Size:  req.Size,
		Total: 0,
	}
	infoError := consts.InfoError(consts.OrderInfo, consts.GetListFail)
	if req.Page <= 0 {
		req.Page = 1
	}
	if req.Size <= 0 || req.Size > 50 {
		req.Size = 10
	}
	pbOrders, total, err := order_info.GetList(ctx, req)
	if err != nil {
		g.Log().Errorf(ctx, "%v %v", infoError, err)
		return nil, gerror.WrapCode(gcode.CodeDbOperationError, err, infoError)
	}
	response.List = pbOrders
	response.Total = uint32(total)
	response.Page = req.Page
	response.Size = req.Size
	return &v1.OrderInfoGetListRes{Data: response}, nil
}
