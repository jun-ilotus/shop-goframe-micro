package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type PraiseInfoCreateReq struct {
	g.Meta   `path:"/praise" method:"post" tags:"点赞管理" summary:"创建点赞"`
	ObjectId uint32 `json:"objectId" v:"required" dc:"对象ID"`
	Type     uint32 `json:"type" v:"required|in:1,2" dc:"点赞类型：1商品2文章"`
}

type PraiseInfoCreateRes struct {
	Id uint32 `json:"id" dc:"点赞ID"`
}

type PraiseInfoDeleteReq struct {
	g.Meta   `path:"/praise" method:"delete" tags:"点赞管理" summary:"删除点赞"`
	Id       uint32 `json:"id" v:"required" dc:"点赞ID"`
	Type     uint32 `json:"type" v:"required|in:1,2" dc:"点赞类型：1商品2文章"`
	ObjectId uint32 `json:"ObjectId" v:"required" dc:"对象ID"`
}

type PraiseInfoDeleteRes struct {
	Id uint32 `json:"id" dc:"被删除的点赞ID"`
}

type PraiseInfoGetListReq struct {
	g.Meta `path:"/praise" method:"get" tags:"点赞管理" summary:"获取点赞列表"`
	Type   uint32 `json:"type" v:"required|in:1,2" dc:"点赞类型：1商品2文章"`
	Page   uint32 `json:"page" v:"min:1" dc:"页码" d:"1"`
	Size   uint32 `json:"size" v:"max:100" dc:"每页数量" d:"10"`
}

type PraiseInfoGetListRes struct {
	List  []*PraiseInfoItem `json:"list" dc:"点赞列表"`
	Page  uint32            `json:"page" dc:"当前页码"`
	Size  uint32            `json:"size" dc:"每页数量"`
	Total uint32            `json:"total" dc:"总数"`
}

type PraiseInfoItem struct {
	Id        uint32                 `json:"id" dc:"点赞ID"`
	UserId    uint32                 `json:"userId" dc:"用户ID"`
	Type      uint32                 `json:"type" dc:"点赞类型：1商品2文章"`
	ObjectId  uint32                 `json:"objectId" dc:"点赞对象ID"`
	CreatedAt *timestamppb.Timestamp `json:"createdAt" dc:"创建时间"`
	UpdatedAt *timestamppb.Timestamp `json:"updatedAt" dc:"更行时间"`
}
