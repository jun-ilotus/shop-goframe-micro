package order_info

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	v1 "service/app/order/api/order_info/v1"
	"service/app/order/api/pbentity"
	"service/app/order/internal/dao"
	"service/app/order/internal/model/entity"
	"service/utility"
)

func Create(ctx context.Context, req *v1.OrderInfoCreateReq) (int32, error) {
	db := g.DB()
	tx, err := db.Begin(ctx)
	if err != nil {
		return 0, fmt.Errorf("开启事务失败:%v", err)
	}

	var success bool
	defer func() {
		if !success {
			if rollbackErr := tx.Rollback(); rollbackErr != nil {
				g.Log().Errorf(ctx, "事务回滚失败: %v", rollbackErr)
			}
		}
	}()

	var order entity.OrderInfo
	if err := gconv.Struct(req, &order); err != nil {
		return 0, fmt.Errorf("订单数据转换失败:%v", err)
	}

	order.Number = utility.GenerateOrderNumber()
	order.Status = 1
	order.CreatedAt = gtime.Now()
	order.UpdatedAt = gtime.Now()

	result, err := dao.OrderInfo.Ctx(ctx).TX(tx).InsertAndGetId(order)
	if err != nil {
		return 0, fmt.Errorf("插入订单失败: %v", err)
	}
	orderId := int32(result)

	var orderGoodsList []entity.OrderGoodsInfo
	if err := gconv.Struct(req.OrderGoodsInfo, &orderGoodsList); err != nil {
		return 0, fmt.Errorf("订单商品数据转换失败: %v", err)
	}

	for i := range orderGoodsList {
		orderGoodsList[i].OrderId = int(orderId)
		orderGoodsList[i].CreatedAt = gtime.Now()
		orderGoodsList[i].UpdatedAt = gtime.Now()
	}

	if len(orderGoodsList) > 0 {
		_, err = dao.OrderGoodsInfo.Ctx(ctx).TX(tx).Insert(orderGoodsList)
		if err != nil {
			return 0, fmt.Errorf("插入订单商品失败: %v", err)
		}
	}

	if err := tx.Commit(); err != nil {
		return 0, fmt.Errorf("提交事务失败: %v", err)
	}

	success = true
	return orderId, nil
}

func GetDetail(ctx context.Context, orderId uint32) (*pbentity.OrderInfo, []*pbentity.OrderGoodsInfo, error) {
	var order entity.OrderInfo
	err := dao.OrderInfo.Ctx(ctx).WherePri(orderId).Scan(&order)
	if err != nil {
		return nil, nil, fmt.Errorf("查询订单商品失败: %v", err)
	}

	var goodsList []*entity.OrderGoodsInfo
	err = dao.OrderGoodsInfo.Ctx(ctx).Where("order_id = ?", orderId).Scan(&goodsList)
	if err != nil {
		return nil, nil, fmt.Errorf("查询订单商品失败: %v", err)
	}

	var pbOrder pbentity.OrderInfo
	if err := gconv.Struct(order, &pbOrder); err != nil {
		return nil, nil, fmt.Errorf("转换订单数据失败: %v", err)
	}
	pbOrder.CreatedAt = utility.SafeConvertTime(order.CreatedAt)
	pbOrder.UpdatedAt = utility.SafeConvertTime(order.UpdatedAt)

	var pbGoodsList []*pbentity.OrderGoodsInfo
	for _, goods := range goodsList {
		var pbGoods pbentity.OrderGoodsInfo
		if err := gconv.Struct(goods, &pbGoods); err != nil {
			continue
		}
		pbGoods.CreatedAt = utility.SafeConvertTime(goods.CreatedAt)
		pbGoods.UpdatedAt = utility.SafeConvertTime(goods.UpdatedAt)
		pbGoodsList = append(pbGoodsList, &pbGoods)
	}
	return &pbOrder, pbGoodsList, nil
}

func GetList(ctx context.Context, req *v1.OrderInfoGetListReq) ([]*pbentity.OrderInfo, int, error) {
	model := dao.OrderInfo.Ctx(ctx)

	if req.Number != "" {
		model = model.Where("number", req.Number)
	}

	if req.UserId != 0 {
		model = model.Where("user_id", req.UserId)
	}

	if req.PayType != 0 {
		model = model.Where("pay_type", req.PayType)
	}

	if req.Status != 0 {
		model = model.Where("status", req.Status)
	}

	if req.ConsigneePhone != "" {
		model = model.Where("consignee_phone", req.ConsigneePhone)
	}

	if req.PriceGte != 0 {
		model = model.Where("price >= ?", req.PriceGte)
	}

	if req.PriceLte != 0 {
		model = model.Where("price <= ?", req.PriceLte)
	}

	if req.PayAtLte != nil {
		model = model.Where("pay_at <= ?", req.PayAtLte.AsTime())
	}

	if req.DateGte != nil {
		model = model.Where("created_at >= ?", req.DateGte.AsTime())
	}

	if req.DateLte != nil {
		model = model.Where("created_at <= ?", req.DateLte.AsTime())
	}

	total, err := model.Count()
	if err != nil {
		return nil, 0, err
	}

	orderRecords, err := model.Page(int(req.Page), int(req.Size)).All()
	if err != nil {
		return nil, 0, err
	}

	var pbOrders []*pbentity.OrderInfo
	for _, record := range orderRecords {
		var order entity.OrderInfo
		if err := record.Struct(&order); err != nil {
			continue
		}

		var pbOrder pbentity.OrderInfo
		if err := gconv.Struct(order, &pbOrder); err != nil {
			continue
		}

		pbOrder.CreatedAt = utility.SafeConvertTime(order.CreatedAt)
		pbOrder.UpdatedAt = utility.SafeConvertTime(order.UpdatedAt)
		pbOrders = append(pbOrders, &pbOrder)
	}
	return pbOrders, total, nil
}
