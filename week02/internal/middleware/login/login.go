package login

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

type MiddlewareBuilder struct {
}

func (m *MiddlewareBuilder) CheckLogin() gin.HandlerFunc {
	return func(context *gin.Context) {
		path := context.Request.URL.Path
		if path == "/user/signup" || path == "/user/login" {
			return
		}
		sess := sessions.Default(context)
		if sess.Get("userId") == nil {
			// 表示session里面并不存在当前用户登录状态,需要终止执行
			context.AbortWithStatus(http.StatusUnauthorized)
			return
		}
	}
}
