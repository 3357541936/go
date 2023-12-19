package login

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"strings"
	"time"
	"week02.com/internal/web"
)

type MiddlewareJWTBuilder struct {
}

func (m *MiddlewareJWTBuilder) CheckLogin() gin.HandlerFunc {
	return func(context *gin.Context) {
		path := context.Request.URL.Path
		if path == "/user/signup" || path == "/user/login" {
			return
		}

		authCode := context.GetHeader("Authorization")
		if authCode == "" {
			// 表示不存在当前用户登录状态,需要终止执行
			context.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		seg := strings.Split(authCode, " ")
		if len(seg) != 2 {
			// 表示不存在当前用户登录状态,Authorization是乱写的
			context.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		tokenStr := seg[1]
		var uc web.UserClaims
		token, err := jwt.ParseWithClaims(tokenStr, &uc, func(token *jwt.Token) (interface{}, error) {
			return web.SignedKey, nil
		})
		if err != nil {
			// 解析 token 出错, token 存在伪造问题
			context.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		if !token.Valid {
			// token 解析成功, 但存在非法或过期问题
			context.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// 登录状态保持
		expireTime := uc.ExpiresAt
		// token 有效期小于指定时长, 刷新token
		if expireTime.Sub(time.Now()) < time.Minute {
			uc.ExpiresAt = jwt.NewNumericDate(time.Now().Add(time.Minute * 15))
			tokenStr, err = token.SignedString(web.SignedKey)
			context.Header("x-jwt-token", tokenStr)
			if err != nil {
				// 日志
			}
		}
		context.Set("user", uc)
	}
}
