package consts

const (
	ConsigneeInfo = "ConsigneeInfo"
	GetListFail   = "GetList 查询失败"
	CreateFail    = "Create 插入失败"
	UpdateFail    = "Update 更新失败"
	DeleteFail    = "Delete 删除失败"

	UserInfo           = "UserInfo"
	RegisterFail       = "Register 注册失败"
	LoginFail          = "Login 登录失败"
	UpdatePasswordFail = "UpdatePassword 修改密码失败"
	GetUserInfoFail    = "GetUserInfo 获取用户信息失败"
)

func InfoError(info string, fail string) string {
	return info + " " + fail
}
