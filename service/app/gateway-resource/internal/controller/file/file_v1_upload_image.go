package file

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"io"
	"service/app/gateway-resource/internal/consts"
	"service/app/gateway-resource/internal/logic/file_info"
	"service/app/gateway-resource/utility"
	"service/utility/middleware"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"service/app/gateway-resource/api/file/v1"
)

func (c *ControllerV1) UploadImage(ctx context.Context, req *v1.UploadImageReq) (res *v1.UploadImageRes, err error) {
	if req.File == nil {
		return nil, gerror.NewCode(gcode.CodeMissingParameter, "upload image file is empty")
	}

	file, err := req.File.Open()
	if err != nil {
		return nil, gerror.NewCode(gcode.CodeInternalError, "open file error")
	}
	defer file.Close()

	fileContent, err := io.ReadAll(file)
	if err != nil {
		return nil, gerror.NewCode(gcode.CodeInternalError, "read file content error")
	}

	url, fileName, err := utility.UploadToQiniu(ctx, fileContent, req.File.Filename)
	if err != nil {
		return nil, gerror.NewCodef(gcode.CodeInternalError, "upload file to qiniu error: %v", err)
	}

	userId := ctx.Value(middleware.CtxUserId)
	if userId == nil {
		return nil, gerror.NewCode(gcode.CodeNotAuthorized, "not authorized")
	}

	infoError := consts.InfoError(consts.FileInfo, consts.UploadImageFail)
	err = file_info.UploadImage(ctx, url, fileName, int(userId.(uint32)))
	if err != nil {
		g.Log().Errorf(ctx, "%v %v", infoError, err)
		return nil, gerror.NewCode(gcode.CodeDbOperationError, infoError)
	}

	return &v1.UploadImageRes{Url: url}, nil
}
