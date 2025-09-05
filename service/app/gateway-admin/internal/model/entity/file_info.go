// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// FileInfo is the golang structure for table file_info.
type FileInfo struct {
	Id         uint        `json:"id"         orm:"id"          description:"文件ID"`   // 文件ID
	Name       string      `json:"name"       orm:"name"        description:"文件名字"`   // 文件名字
	Url        string      `json:"url"        orm:"url"         description:"七牛云URL"` // 七牛云URL
	UploaderId uint        `json:"uploaderId" orm:"uploader_id" description:"上传者ID"`  // 上传者ID
	CreatedAt  *gtime.Time `json:"createdAt"  orm:"created_at"  description:""`       //
	DeletedAt  *gtime.Time `json:"deletedAt"  orm:"deleted_at"  description:""`       //
}
