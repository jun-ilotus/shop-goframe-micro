package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type RotationInfoGetListReq struct {
	g.Meta `path:"/rotation" method:"get" tags:"轮播图管理" summary:"轮播图分页列表"`
	Sort   uint32 `json:"sort" dc:"排列方式"`
	Page   uint32 `json:"page" d:"1" v:"min:1" dc:"页码"`
	Size   uint32 `json:"size" d:"10" v:"max:100" dc:"每页数量"`
}

type RotationInfoGetListRes struct {
	List  []*RotationInfoItem `json:"list" dc:"手工位图列表"`
	Page  uint32              `json:"page" dc:"当前页码"`
	Size  uint32              `json:"size" dc:"每页数量"`
	Total uint32              `json:"total" dc:"总数"`
}

type RotationInfoItem struct {
	Id        uint32                 `json:"id" dc:"ID"`
	PicUrl    string                 `json:"pic_url" dc:"图片链接"`
	Link      string                 `json:"link" dc:"跳转链接"`
	Sort      uint32                 `json:"sort" dc:"排序字段"`
	CreatedAt *timestamppb.Timestamp `json:"created_at" dc:"创建时间"`
	UpdatedAt *timestamppb.Timestamp `json:"updated_at" dc:"更新时间"`
}
