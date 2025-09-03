package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// 评论分页查询
type CommentInfoGetListReq struct {
	g.Meta `path:"/comment" method:"get" tags:"评论管理" sm:"评论分页列表"`
	Page   uint32 `json:"page" d:"1"  v:"min:1" dc:"页码"`
	Size   uint32 `json:"size" d:"10" v:"max:100" dc:"每页数量"`
}

type CommentInfoGetListRes struct {
	List  []*CommentInfoItem `json:"list" dc:"评论列表"`
	Page  uint32             `json:"page" dc:"当前页码"`
	Size  uint32             `json:"size" dc:"每页数量"`
	Total uint32             `json:"total" dc:"总数"`
}

type CommentInfoItem struct {
	Id        uint32                 `json:"id" dc:"评论ID"`
	UserId    uint32                 `json:"user_id" dc:"用户ID"`
	ObjectId  uint32                 `json:"object_id" dc:"对象ID"`
	Type      uint32                 `json:"type" dc:"评论类型：1商品 2文章"`
	ParentId  uint32                 `json:"parent_id" dc:"父级评论ID"`
	Content   string                 `json:"content" dc:"评论内容"`
	CreatedAt *timestamppb.Timestamp `json:"created_at" dc:"创建时间"`
	UpdatedAt *timestamppb.Timestamp `json:"updated_at" dc:"更新时间"`
}

// 创建评论
type CommentInfoCreateReq struct {
	g.Meta   `path:"/comment" method:"post" tags:"评论管理" sm:"创建评论"`
	ObjectId uint32 `json:"object_id" v:"required" dc:"对象ID"`
	Type     uint32 `json:"type" v:"required|in:1,2" dc:"评论类型：1商品 2文章"`
	ParentId uint32 `json:"parent_id" dc:"父级评论ID"`
	Content  string `json:"content" v:"required" dc:"评论内容"`
}

type CommentInfoCreateRes struct {
	Id uint32 `json:"id" dc:"评论ID"`
}

// 删除评论
type CommentInfoDeleteReq struct {
	g.Meta `path:"/comment" method:"delete" tags:"评论管理" sm:"删除评论"`
	Id     uint32 `json:"id" v:"required" dc:"评论ID"`
}

type CommentInfoDeleteRes struct {
	Id uint32 `json:"id" dc:"被删除的评论ID"`
}
