package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type CategoryInfoGetListReq struct {
	g.Meta `path:"/category" method:"get" tags:"商品分类管理" summary:"商品分类分页列表"`
	Sort   uint32 `json:"sort" dc:"排序"`
	Page   uint32 `json:"page" d:"1" v:"min:1" dc:"页码"`
	Size   uint32 `json:"size" d:"10" v:"max:100" dc:"每页数量"`
}

type CategoryInfoGetListRes struct {
	List  []*CategoryInfoItem `json:"list" dc:"商品分类列表"`
	Page  uint32              `json:"page" dc:"当前页码"`
	Size  uint32              `json:"size" dc:"每页数量"`
	Total uint32              `json:"total" dc:"总数"`
}

type CategoryInfoGetAllReq struct {
	g.Meta `path:"/category/all" method:"get" tags:"商品分类管理" summary:"获取所有商品分类"`
}

type CategoryInfoGetAllRes struct {
	List  []*CategoryInfoItem `json:"list" dc:"商品分类列表"`
	Total uint32              `json:"total" dc:"总数"`
}

type CategoryInfoCreateReq struct {
	g.Meta   `path:"/category" method:"post" tags:"商品分类管理" summary:"创建商品分类"`
	ParentId uint32 `json:"parent_id" dc:"父级id"`
	Name     string `json:"name" v:"required" dc:""`
	PicUrl   string `json:"pic_url" dc:"icon"`
	Level    uint32 `json:"level" dc:"等级 默认1级分类"`
	Sort     uint32 `json:"sort" dc:"排序"`
}

type CategoryInfoCreateRes struct {
	Id uint32 `json:"id" dc:"商品分类ID"`
}

type CategoryInfoUpdateReq struct {
	g.Meta   `path:"/category" method:"put" tags:"商品分类管理" summary:"更新商品分类"`
	Id       uint32 `json:"id" v:"required" dc:"商品分类ID"`
	ParentId uint32 `json:"parent_id" dc:"分类id"`
	Name     string `json:"name" dc:""`
	PicUrl   string `json:"pic_url" dc:"icon"`
	Level    uint32 `json:"level" dc:"等级 默认1级分类"`
	Sort     uint32 `json:"sort" dc:"排序"`
}

type CategoryInfoUpdateRes struct {
	Id uint32 `json:"id" dc:"商品分类ID"`
}

type CategoryInfoDeleteReq struct {
	g.Meta `path:"/category" method:"delete" tags:"商品分类管理" summary:"删除商品分类"`
	Id     uint32 `json:"id" v:"required" dc:"商品分类ID"`
}

type CategoryInfoDeleteRes struct {
}

type CategoryInfoItem struct {
	Id        uint32                 `json:"id" v:"required" dc:"商品分类ID"`
	ParentId  uint32                 `json:"parent_id" dc:"分类id"`
	Name      string                 `json:"name" dc:""`
	PicUrl    string                 `json:"pic_url" dc:"icon"`
	Level     uint32                 `json:"level" dc:"等级 默认1级分类"`
	Sort      uint32                 `json:"sort" dc:"排序"`
	CreatedAt *timestamppb.Timestamp `json:"created_at" dc:"创建时间"`
	UpdatedAt *timestamppb.Timestamp `json:"updated_at" dc:"更新时间"`
}
