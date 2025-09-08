// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// FileInfo is the golang structure for table file_info.
type FileInfo struct {
	Id           uint        `json:"id"           orm:"id"            description:"文件ID"`                                 // 文件ID
	Name         string      `json:"name"         orm:"name"          description:"文件名字"`                                 // 文件名字
	Url          string      `json:"url"          orm:"url"           description:"七牛云URL"`                               // 七牛云URL
	UploaderId   uint        `json:"uploaderId"   orm:"uploader_id"   description:"上传者ID（根据uploader_type区分是用户ID还是管理员ID）"` // 上传者ID（根据uploader_type区分是用户ID还是管理员ID）
	UploaderType uint        `json:"uploaderType" orm:"uploader_type" description:"上传者类型：1-H5用户，2-管理员"`                   // 上传者类型：1-H5用户，2-管理员
	FileType     uint        `json:"fileType"     orm:"file_type"     description:"文件类型：1-图片，2-视频，3-其他"`                  // 文件类型：1-图片，2-视频，3-其他
	CreatedAt    *gtime.Time `json:"createdAt"    orm:"created_at"    description:""`                                     //
	DeletedAt    *gtime.Time `json:"deletedAt"    orm:"deleted_at"    description:""`                                     //
}
