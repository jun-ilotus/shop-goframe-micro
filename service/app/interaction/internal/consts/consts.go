package consts

const (
	CollectionInfo = "CollectionInfo"
	CommentInfo    = "CommentInfo"
	PraiseInfo     = "PraiseInfo"
	CreateFail     = "Create 插入失败"
	DeleteFail     = "Delete 删除失败"
	GetListFail    = "GetList 获取列表失败"
)

func InfoError(info string, fail string) string {
	return info + " " + fail
}
