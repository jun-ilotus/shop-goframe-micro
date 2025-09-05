package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

type UploadImageReq struct {
	g.Meta `path:"/upload/image" method:"post" tags:"文件上传" summary:"上传图片"`
	File   *ghttp.UploadFile `json:"-" dc:"上传的文件" v:"required#请选择上传文件"`
}

type UploadImageRes struct {
	Url string `json:"url" dc:"图片访问URL"`
}
