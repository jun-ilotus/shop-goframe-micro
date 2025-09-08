// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// FileInfo is the golang structure of table file_info for DAO operations like Where/Data.
type FileInfo struct {
	g.Meta       `orm:"table:file_info, do:true"`
	Id           interface{} // 文件ID
	Name         interface{} // 文件名字
	Url          interface{} // 七牛云URL
	UploaderId   interface{} // 上传者ID（根据uploader_type区分是用户ID还是管理员ID）
	UploaderType interface{} // 上传者类型：1-H5用户，2-管理员
	FileType     interface{} // 文件类型：1-图片，2-视频，3-其他
	CreatedAt    *gtime.Time //
	DeletedAt    *gtime.Time //
}
