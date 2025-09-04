package consts

const (
	GetListFail        = "GetList 查询失败"
	GetDetailFile      = "GetDetail 查询失败"
	CreateFail         = "Create 插入失败"
	UpdateFail         = "Update 更新失败"
	DeleteFail         = "Delete 删除失败"
	RegisterFail       = "Register 注册失败"
	LoginFail          = "Login 登录失败"
	UpdatePasswordFail = "UpdatePassword 修改密码失败"
	GetUserInfoFail    = "GetUserInfo 获取用户信息失败"
	UploadImageFail    = "UploadImage 图片保存失败"
	SearchFail         = "SearchFail  搜索失败 "
	OrderInfo          = "OrderInfo"
	SearchGoods        = "SearchGoods"
	CategoryInfo       = "CategoryInfo"
	GoodsImages        = "GoodsImages"
	GoodsInfo          = "GoodsInfo"
	FileInfo           = "FileInfo"
	ConsigneeInfo      = "ConsigneeInfo"
	UserInfo           = "UserInfo"
	CollectionInfo     = "CollectionInfo"
	CommentInfo        = "CommentInfo"
	PraiseInfo         = "PraiseInfo"
	AdminInfo          = "AdminInfo"
)

func InfoError(info string, fail string) string {
	return info + " " + fail
}
