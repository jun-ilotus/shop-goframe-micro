package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type SearchGoodsReq struct {
	g.Meta   `path:"/search/goods" method:"get" tags:"搜索" sm:"商品搜索"`
	Keyword  string `json:"keyword" dc:"搜索关键词"`
	Brand    string `json:"brand" dc:"品牌名称"`
	MinPrice uint64 `json:"min_price" dc:"最低价格（分）"`
	MaxPrice uint64 `json:"max_price" dc:"最高价格（分）"`
	Sort     string `json:"sort" dc:"排序方式：default默认 price_asc价格升序 price_desc价格降序 sale销量"`
	Page     uint32 `json:"page" d:"1" v:"min:1" dc:"页码"`
	Size     uint32 `json:"size" d:"10" v:"max:100" dc:"每页数量"`
}

type SearchGoodsRes struct {
	List  []*GoodsInfoItem `json:"list" dc:"商品列表"`
	Page  uint32           `json:"page" dc:"当前页面"`
	Size  uint32           `json:"size" dc:"每页数量"`
	Total uint32           `json:"total" dc:"总数"`
}

type GoodsInfoItem struct {
	Id               uint32                 `json:"id" dc:"商品ID"`
	Name             string                 `json:"name" dc:"商品名称"`
	PicUrl           string                 `json:"pic_url" dc:"主图"`
	Images           string                 `json:"images" dc:"支持单图,多图"`
	Price            uint64                 `json:"price" dc:"价格(分)"`
	Level1CategoryId uint32                 `json:"level1_category_id" dc:"一级分类ID"`
	Level2CategoryId uint32                 `json:"level2_category_id" dc:"二级分类ID"`
	Level3CategoryId uint32                 `json:"level3_category_id" dc:"三级分类ID"`
	Brand            string                 `json:"brand" dc:"品牌"`
	Stock            uint32                 `json:"stock" dc:"库存"`
	Sale             uint32                 `json:"sale" dc:"销量"`
	Tags             string                 `json:"tags" dc:"标签"`
	DetailInfo       string                 `json:"detail_info" dc:"详情"`
	CreatedAt        *timestamppb.Timestamp `json:"created_at" dc:"创建时间"` // 改为 string
	UpdatedAt        *timestamppb.Timestamp `json:"updated_at" dc:"更新时间"` // 改为 string
	Highlight        string                 `json:"highlight" dc:"高亮名称"`
}
