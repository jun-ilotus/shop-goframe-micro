package v1

import "github.com/gogf/gf/v2/frame/g"

type GoodsImagesGetListReq struct {
	g.Meta `path:"/goods/images" method:"get" tags:"商品图片管理" sm:"商品图片列表"`
	Page   uint32 `json:"page" d:"1" v:"min:1" dc:"页码"`
	Size   uint32 `json:"size" d:"10" v:"max:100" dc:"每页数量"`
}

type GoodsImagesGetListRes struct {
	List  []*GoodsImagesItem `json:"list" dc:"商品图片列表"`
	Page  uint32             `json:"page" dc:"当前页码"`
	Size  uint32             `json:"size" dc:"每页数量"`
	Total uint32             `json:"total" dc:"总数"`
}

type GoodsImagesItem struct {
	Id      uint32 `json:"id" dc:"商品图片ID"`
	GoodsId uint32 `json:"goods_id" dc:"商品"`
	FileId  uint32 `json:"file_id" dc:"文件ID"`
	Sort    uint32 `json:"sort" dc:"排序"`
}
