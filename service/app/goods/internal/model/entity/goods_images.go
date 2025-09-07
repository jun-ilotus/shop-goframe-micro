// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// GoodsImages is the golang structure for table goods_images.
type GoodsImages struct {
	Id      uint   `json:"id"      orm:"id"       description:""`                  //
	Url     string `json:"url"     orm:"url"      description:"七牛云url"`            // 七牛云url
	GoodsId int    `json:"goodsId" orm:"goods_id" description:"商品ID"`              // 商品ID
	FileId  int    `json:"fileId"  orm:"file_id"  description:"文件ID（关联file_info）"` // 文件ID（关联file_info）
	Sort    int    `json:"sort"    orm:"sort"     description:"排序"`                // 排序
}
