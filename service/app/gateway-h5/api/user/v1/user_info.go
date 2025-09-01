package v1

import "github.com/gogf/gf/v2/frame/g"

type UserInfoLoginReq struct {
	g.Meta   `path:"/user/login" tags:"用户管理" method:"post" summary:"登录"`
	Name     string `json:"name" v:"required#用户名不能为空" dc:"用户名"`
	Password string `json:"password" v:"required#密码不能为空" dc:"密码"`
}

type UserInfoLoginRes struct {
	Type     string        `json:"type" dc:"token类型"`
	Token    string        `json:"token" dc:"token字符串"`
	ExpireIn uint32        `json:"expire_in" dc:"过期时间（秒）"`
	UserInfo *UserInfoBase `json:"user_info" dc:"用户基础信息"`
}

type UserInfoRegisterReq struct {
	g.Meta       `path:"/user/register" tags:"用户管理" method:"post" summary:"注册"`
	Name         string `json:"name" v:"required#用户名不能为空" dc:"用户名"`
	Password     string `json:"password" v:"required#密码不能为空" dc:"密码"`
	Avatar       string `json:"avatar" dc:"头像"`
	Sex          uint32 `json:"sex" dc:"1男 2女"`
	Sign         string `json:"sign" dc:"个性签名"`
	SecretAnswer string `json:"secret_answer" dc:"密保问题的答案"`
}

type UserInfoRegisterRes struct {
	Id uint32 `json:"id" dc:"用户ID"`
}

type UserInfoReq struct {
	g.Meta `path:"/user/info" tags:"用户管理" method:"get" summary:"获取用户信息"`
	Id     uint32 `json:"id" v:"required#用户ID不能为空" dc:"用户ID"`
}

type UserInfoRes struct {
	UserInfo *UserInfoBase `json:"user_info" dc:"用户基础信息"`
}

type UserInfoUpdatePasswordReq struct {
	g.Meta       `path:"user/update/password" tags:"用户管理" method:"put" summary:"修改用户密码"`
	Id           uint32 `json:"id" v:"required#用户ID不能为空" dc:"用户ID"`
	Password     string `json:"password" v:"required#新密码不能为空" dc:"新密码"`
	SecretAnswer string `json:"secret_answer" v:"required#密保问题的答案不能为空" dc:"密保问题的答案"`
}

type UserInfoUpdatePasswordRes struct {
	Id uint32 `json:"id" dc:"用户ID"`
}

type UserInfoBase struct {
	Id     uint32 `json:"id" dc:"用户ID"`
	Name   string `json:"name" dc:"用户名"`
	Avatar string `json:"avatar" dc:"头像"`
	Sex    uint32 `json:"sex" dc:"性别"`
	Sign   string `json:"sign" dc:"个性签名"`
	Status uint32 `json:"status" dc:"状态"`
}
