// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package file

import (
	"context"

	"service/app/gateway-admin/api/file/v1"
)

type IFileV1 interface {
	UploadImage(ctx context.Context, req *v1.UploadImageReq) (res *v1.UploadImageRes, err error)
}
