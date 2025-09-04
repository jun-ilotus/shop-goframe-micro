package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"time"
)

type AdminInfoLoginReq struct {
	g.Meta   `path:"/admin/login" tags:"管理员登录" method:"post" summary:"登录"`
	Name     string `json:"name" v:"required#用户名不能为空" dc:"用户名"`
	Password string `json:"password" v:"required#密码不能为空" dc:"密码"`
}

type AdminInfoLoginRes struct {
	Token  string `json:"token"`
	Expire int64  `json:"expire"`
}

type AdminInfoRegisterReq struct {
	g.Meta   `path:"/admin/register" tags:"管理员注册" method:"post" summary:"注册"`
	Name     string `json:"name" v:"required#用户名不能为空" dc:"用户名"`
	Password string `json:"password" v:"required#密码不能为空" dc:"密码"`
}

type AdminInfoRegisterRes struct {
	Id        uint32    `json:"id"`
	Name      string    `json:"name"`
	IsAdmin   uint32    `json:"is_admin"`
	RoleIds   string    `json:"role_ids"`
	CreatedAt time.Time `json:"created_at"`
}
