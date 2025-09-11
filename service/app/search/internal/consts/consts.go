package consts

const (
	SearchFail  = "SearchFail  搜索失败 "
	SearchGoods = "SearchGoods"
)

func InfoError(info string, fail string) string {
	return info + " " + fail
}
