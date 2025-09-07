package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type GoodsInfoGetDetailReq struct {
	g.Meta `path:"/goods/detail" method:"get" tags:"商品" sm:"商品详情"`
	Id     uint32 `json:"id" v:"required" dc:"商品ID"`
}

type GoodsInfoGetDetailRes struct {
	*GoodsInfoItem
}

type GoodsInfoItem struct {
	Id               uint32                 `json:"id" dc:"商品"`
	Name             uint32                 `json:"name" dc:"商品名称"`
	Price            uint32                 `json:"price" dc:"价格"`
	Level1CategoryId uint32                 `json:"level1_category_id" dc:"一级分类ID"`
	Level2CategoryId uint32                 `json:"level2_category_id" dc:"二级分类ID"`
	Level3CategoryId uint32                 `json:"level3_category_id" dc:"三级分类ID"`
	Brand            string                 `json:"brand" dc:"品牌"`
	Stock            uint32                 `json:"stock" dc:"库存"`
	Sale             uint32                 `json:"sale" dc:"销量"`
	Tags             string                 `json:"tags" dc:"标签"`
	DetailInfo       string                 `json:"detail_info" dc:"详情"`
	CreateAt         *timestamppb.Timestamp `json:"create_at" dc:"创建时间"`
	UpdateAt         *timestamppb.Timestamp `json:"update_at" dc:"更新时间"`
}

type GoodsInfoGetListReq struct {
	g.Meta `path:"/goods" method:"get" tags:"商品管理" sm:"商品分页列表"`
	Page   uint32 `json:"page" d:"1" v:"min:1" dc:"页码"`
	Size   uint32 `json:"size" d:"10" v:"max:100" dc:"每页数量"`
}

type GoodsInfoGetListRes struct {
	List  []*GoodsInfoItem `json:"list" dc:"商品列表"`
	Page  uint32           `json:"page" dc:"当前页码"`
	Size  uint32           `json:"size" dc:"每页数量"`
	Total uint32           `json:"total" dc:"总数"`
}
