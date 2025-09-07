// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// GoodsImages is the golang structure of table goods_images for DAO operations like Where/Data.
type GoodsImages struct {
	g.Meta  `orm:"table:goods_images, do:true"`
	Id      interface{} //
	Url     interface{} // 七牛云url
	GoodsId interface{} // 商品ID
	FileId  interface{} // 文件ID（关联file_info）
	Sort    interface{} // 排序
}
