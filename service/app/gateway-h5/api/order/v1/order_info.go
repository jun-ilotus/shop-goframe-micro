package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type OrderInfoGetListReq struct {
	g.Meta         `path:"/order" method:"get" tags:"订单管理" sm:"订单分页列表"`
	Page           uint32                 `json:"page" d:"1" v:"min:1" dc:"页码"`
	Size           uint32                 `json:"size" d:"10" v:"max:50" dc:"每页数量"`
	Number         string                 `json:"number" dc:"订单编号"`
	UserId         uint32                 `json:"user_id" dc:"用户ID"`
	PayType        uint32                 `json:"pay_type" dc:"支付方式 1微信 2支付宝 3云闪付"`
	Status         uint32                 `json:"status" dc:"订单状态 1待支付 2已支付待发货 3已发货 4已收货待评价"`
	ConsigneePhone string                 `json:"consignee_phone" dc:"收货人手机号"`
	PriceGte       uint32                 `json:"price_gte" dc:"订单金额 >= 单位分"`
	PriceLte       uint32                 `json:"price_lte" dc:"订单金额 <= 单位分"`
	PayAtGte       *timestamppb.Timestamp `json:"pay_at_gte" dc:"支付时间 >="`
	PayAtLte       *timestamppb.Timestamp `json:"pay_at_lte" dc:"支付时间 <="`
	DateGte        *timestamppb.Timestamp `json:"date_gte" dc:"创建时间 >="`
	DateLte        *timestamppb.Timestamp `json:"date_lte" dc:"创建时间 <="`
}

type OrderInfoGetListRes struct {
	List  []*OrderInfoItem `json:"list" dc:"订单列表"`
	Page  uint32           `json:"page" dc:"当前页码"`
	Size  uint32           `json:"size" dc:"每页数量"`
	Total uint32           `json:"total" dc:"总数"`
}

type OrderInfoItem struct {
	Id               uint32                 `json:"id" dc:"订单ID"`
	Number           string                 `json:"number" dc:"订单编号"`
	Price            uint32                 `json:"price" dc:"订单金额"`
	CouponPrice      uint32                 `json:"coupon_price" dc:"优惠券金额"`
	ActualPrice      uint32                 `json:"actual_price" dc:"实际支付金额"`
	ConsigneeName    string                 `json:"consignee_name" dc:"收货人姓名"`
	ConsigneePhone   string                 `json:"consignee_phone" dc:"收货人手机号"`
	ConsigneeAddress string                 `json:"consignee_address" dc:"收货人详细地址"`
	Remark           string                 `json:"remark" dc:"备注"`
	Status           uint32                 `json:"status" dc:"订单状态"`
	CreatedAt        *timestamppb.Timestamp `json:"created_at" dc:"创建时间"`
	UpdatedAt        *timestamppb.Timestamp `json:"updated_at" dc:"更新时间"`
}

type OrderInfoCreateReq struct {
	g.Meta           `path:"/order" method:"post" tags:"订单管理" sm:"创建订单"`
	Price            uint32            `json:"price" v:"required|min:0" dc:"订单金额"`
	CouponPrice      uint32            `json:"coupon_price" d:"0" dc:"优惠券金额"`
	ActualPrice      uint32            `json:"actual_price" v:"required|min:1" dc:"实际支付金额"`
	ConsigneeName    string            `json:"consignee_name" v:"required" dc:"收货人姓名"`
	ConsigneePhone   string            `json:"consignee_phone" v:"required" dc:"收货人手机"`
	ConsigneeAddress string            `json:"consignee_address" v:"required" dc:"收货人详细地址"`
	Remark           string            `json:"remark" dc:"备注"`
	OrderGoodsInfo   []*OrderGoodsItem `json:"order_goods_info" v:"required" dc:"订单商品信息"`
}

type OrderInfoCreateRes struct {
	Id uint32 `json:"id" dc:"订单ID"`
}

type OrderGoodsItem struct {
	GoodsId        uint32 `json:"goods_id" dc:"商品ID"`
	GoodsOptionsId uint32 `json:"goods_options_id" dc:"商品规格ID"`
	Count          uint32 `json:"count" dc:"商品数量"`
	Remark         string `json:"remark" dc:"备注"`
	Price          uint32 `json:"price" dc:"商品金额"`
	CouponPrice    uint32 `json:"coupon_price" dc:"商品优惠券金额"`
	ActualPrice    uint32 `json:"actual_price" dc:"商品实际支付金额"`
}
