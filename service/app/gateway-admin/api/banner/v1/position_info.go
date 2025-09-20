package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type PositionInfoGetListReq struct {
	g.Meta `path:"/position" method:"get" tags:"手工位图管理" summary:"手工位图分页列表"`
	Sort   uint32 `json:"sort" dc:"排序方式"`
	Page   uint32 `json:"page" d:"1" v:"min:1" dc:"页码"`
	Size   uint32 `json:"size" d:"10" v:"max:100" dc:"每页数量"`
}

type PositionInfoGetListRes struct {
	List  []*PositionInfoItem `json:"list" dc:"手工位图列表"`
	Page  uint32              `json:"page" dc:"当前页码"`
	Size  uint32              `json:"size" dc:"每页数量"`
	Total uint32              `json:"total" dc:"总数"`
}

type PositionInfoItem struct {
	Id        uint32                 `json:"id" dc:"ID"`
	PicUrl    string                 `json:"pic_url" dc:"图片链接"`
	GoodsName string                 `json:"goods_name" dc:"商品名称"`
	Link      string                 `json:"link" dc:"跳转链接"`
	Sort      uint32                 `json:"sort" dc:"排序字段"`
	GoodsId   uint32                 `json:"goods_id" dc:"商品ID"`
	CreatedAt *timestamppb.Timestamp `json:"created_at" dc:"创建时间"`
	UpdatedAt *timestamppb.Timestamp `json:"updated_at" dc:"更新时间"`
}

type PositionInfoCreateReq struct {
	g.Meta    `path:"/position" method:"post" tags:"手工位管理" summary:"创建手工位图"`
	PicUrl    string `json:"pic_url" v:"required|url" dc:"图片URL"`
	GoodsName string `json:"goods_name" v:"required" dc:"商品名称"`
	Link      string `json:"link" v:"required|url" dc:"跳转链接"`
	Sort      uint32 `json:"sort" v:"min:0" dc:"排序字段"`
	GoodsId   uint32 `json:"goods_id" v:"required" dc:"商品ID"`
}

type PositionInfoCreateRes struct {
	Id uint32 `json:"id" dc:"手工位图ID"`
}

type PositionInfoUpdateReq struct {
	g.Meta    `path:"/position" method:"put" tags:"手工位图管理" summary:"更新手工位图"`
	Id        uint32 `json:"id" v:"required" dc:"手工位图ID"`
	PicUrl    string `json:"pic_url" v:"url" dc:"图片URL"`
	GoodsName string `json:"goods_name" dc:"商品名称"`
	Link      string `json:"link" v:"url" dc:"跳转链接"`
	Sort      uint32 `json:"sort" v:"min:0" dc:"排序字段"`
	GoodsId   uint32 `json:"goods_id" dc:"商品ID"`
}

type PositionInfoUpdateRes struct {
	Id uint32 `json:"id" dc:"手工位图ID"`
}

type PositionInfoDeleteReq struct {
	g.Meta `path:"/position" method:"delete" tags:"手工位图管理" summary:"删除手工位图"`
	Id     uint32 `json:"id" v:"required" dc:"手工位图ID"`
}

type PositionInfoDeleteRes struct{}
