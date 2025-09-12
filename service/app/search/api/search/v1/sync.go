package v1

import "github.com/gogf/gf/v2/frame/g"

type SyncGoodsReq struct {
	g.Meta           `path:"/sync/goods" method:"post" tags:"同步" sm:"同步商品数据"`
	Id               uint32 `json:"id"`
	Name             uint32 `json:"name"`
	Images           string `json:"images"`
	Price            uint32 `json:"price"`
	Level1CategoryId uint32 `json:"level1_category_id"`
	Level2CategoryId uint32 `json:"level2_category_id"`
	Level3CategoryId uint32 `json:"level3_category_id"`
	Brand            string `json:"brand"`
	Stock            uint32 `json:"stock"`
	Sale             uint32 `json:"sale"`
	Tags             string `json:"tags"`
	DetailInfo       string `json:"detail_info"`
	Operation        string `json:"operation" v:"required#操作类型不能为空" dc:"操作类型：create, update, delete"`
	CreatedAt        string `json:"created_at" dc:"创建时间"`
	UpdatedAt        string `json:"updated_at" dc:"更新时间"`
	DeletedAt        string `json:"deleted_at" dc:"删除时间"`
}

type SyncGoodsRes struct {
	Success bool `json:"success" dc:"是否成功"`
}
