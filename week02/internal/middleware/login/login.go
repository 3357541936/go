package login

import (
	"encoding/gob"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type MiddlewareBuilder struct {
}

func (m *MiddlewareBuilder) CheckLogin() gin.HandlerFunc {
	gob.Register(time.Now())
	return func(context *gin.Context) {
		path := context.Request.URL.Path
		if path == "/user/signup" || path == "/user/login" {
			return
		}
		sess := sessions.Default(context)
		userId := sess.Get("userId")
		if userId == nil {
			// 表示session里面并不存在当前用户登录状态,需要终止执行
			context.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// 重新设定 过期时间
		const updateTimeKey = "update_time"
		now := time.Now()

		val := sess.Get(updateTimeKey)
		prevUpdateTime, ok := val.(time.Time)
		if val == nil || !ok || now.Sub(prevUpdateTime) > time.Minute*14 {
			// redis 中的时间 与 go 中的时间有区别, 需要使用 gob 注册
			sess.Set(updateTimeKey, now)
			sess.Set("userId", userId)
			err := sess.Save()
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}
