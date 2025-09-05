package file_info

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"service/app/gateway-admin/internal/dao"
	"service/app/gateway-admin/internal/model/do"
)

func UploadImage(ctx context.Context, url, fileName string, userId int) (err error) {
	fileRecord := &do.FileInfo{
		Name:       fileName,
		Url:        url,
		UploaderId: userId,
	}

	_, err = dao.FileInfo.Ctx(ctx).Insert(fileRecord)
	if err != nil {
		return gerror.Wrap(err, "创建文件记录失败")
	}

	return nil
}
