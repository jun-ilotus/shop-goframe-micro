package consts

const (
	GetListFail   = "GetList 查询失败"
	GetDetailFile = "GetDetail 查询失败"
	CreateFail    = "Create 插入失败"
	OrderInfo     = "OrderInfo"
)

func InfoError(info string, fail string) string {
	return info + " " + fail
}
