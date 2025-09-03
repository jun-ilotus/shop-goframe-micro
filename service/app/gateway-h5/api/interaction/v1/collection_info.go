package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type CollectionInfoCreateReq struct {
	g.Meta   `path:"/collection" method:"post" tags:"收藏管理" summary:"创建收藏"`
	ObjectId uint32 `json:"objectId" v:"required" dc:"收藏ID"`
	Type     uint32 `json:"type" v:"required|in:1,2" dc:"收藏类型:1商品2文章"`
}

type CollectionInfoCreateRes struct {
	Id uint32 `json:"id" dc:"收藏ID"`
}

type CollectionInfoDeleteReq struct {
	g.Meta   `path:"/collection" method:"delete" tags:"收藏管理" summary:"删除收藏"`
	Id       uint32 `json:"id" v:"required" dc:"收藏ID"`
	Type     uint32 `json:"type" v:"required|in:1,2" dc:"收藏类型:1商品2文章"`
	ObjectId uint32 `json:"objectId" v:"required" dc:"对象ID"`
}

type CollectionInfoDeleteRes struct {
	Id uint32 `json:"id" dc:"被删除的收藏ID"`
}

type CollectionInfoGetListReq struct {
	g.Meta `path:"/collection" method:"get" tags:"收藏管理" summary:"获取收藏列表"`
	Type   uint32 `json:"type" v:"required|in:1,2" dc:"收藏类型：1商品 2文章"`
	Page   uint32 `json:"page" v:"min:1" dc:"页码" d:"1"`
	Size   uint32 `json:"size" v:"max:100" dc:"每页数量" d:"10"`
}

type CollectionInfoGetListRes struct {
	List  []*CollectionInfoItem `json:"list" dc:"收藏列表"`
	Page  uint32                `json:"page" dc:"当前页码"`
	Size  uint32                `json:"size" dc:"每页数量"`
	Total uint32                `json:"total" dc:"总数"`
}

type CollectionInfoItem struct {
	Id        uint32                 `json:"id" dc:"收藏ID"`
	UserId    uint32                 `json:"userId" dc:"用户ID"`
	Type      uint32                 `json:"type" dc:"收藏类型：1商品2文章"`
	ObjectId  uint32                 `json:"objectId" dc:"收藏对象ID"`
	CreatedAt *timestamppb.Timestamp `json:"createdAt" dc:"创建时间"`
	UpdatedAt *timestamppb.Timestamp `json:"updatedAt" dc:"更新时间"`
}
