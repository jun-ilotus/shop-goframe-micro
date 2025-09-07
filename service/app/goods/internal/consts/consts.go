package consts

const (
	GetDetailFail = "GetDetail 查询失败"
	GoodsImages   = "GoodsImages"
	GoodsInfo     = "GoodsInfo"
	GetListFail   = "GetList 查询失败"
	CreateFail    = "Create 插入失败"
	UpdateFail    = "Update 更新失败"
	DeleteFail    = "Delete 删除失败"
)

func InfoError(info string, fail string) string {
	return info + " " + fail
}
