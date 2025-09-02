package middleware

import (
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gctx"
	"service/utility"
)

const (
	CtxUserId gctx.StrKey = "userId"
)

func JWTAuth(r *ghttp.Request) {
	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		r.SetError(gerror.NewCode(gcode.CodeNotAuthorized, "未提供Token"))
		return
	}

	// 移除Bearer前缀
	if len(tokenString) > 7 && tokenString[:7] == "Bearer" {
		tokenString = tokenString[7:]
	}

	claims, err := utility.ParseToken(tokenString)
	if err != nil || claims == nil {
		r.SetError(gerror.NewCode(gcode.CodeNotAuthorized, "无效Token"))
		return
	}

	// 将用户 ID 存入上下文
	r.SetCtxVar(CtxUserId, claims.UserId)
	r.Middleware.Next()
}
