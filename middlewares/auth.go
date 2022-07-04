package middlewares

import (
	"bluebell/controllers"
	"bluebell/pkg/jwt"
	"strings"

	"github.com/gin-gonic/gin"
)

func JWTAuthmiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		// Authorization: Bearer token
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			controllers.ResponseError(c, controllers.CodeNeedLogin)
			c.Abort()
			return
		}
		// 按空格分割
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			controllers.ResponseError(c, controllers.CodeInvalidToken)
			c.Abort()
			return
		}
		// 解析token
		mc, err := jwt.ParseToken(parts[1])
		if err != nil {
			controllers.ResponseError(c, controllers.CodeInvalidToken)
			c.Abort()
			return
		}
		// 将当前请求的 userID 信息保存到请求的上下文 c 中
		// 后续就可以使用 c.GET("userID") 拿到
		c.Set(controllers.CtxUserIDKey, mc.UserID)
		c.Next()
	}
}
