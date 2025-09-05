package consts

const (
	UploadImageFail = "UploadImage 图片保存失败"
	FileInfo        = "FileInfo"
)

func InfoError(info string, fail string) string {
	return info + " " + fail
}
