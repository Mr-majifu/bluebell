package routes

import (
	"bluebell/controllers"
	"bluebell/logger"
	"bluebell/settings"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
	// 使下面的信息也不会出现在终端，也就是开启gin框架的发布模式，默认是debug模式
	//[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
	//- using env:   export GIN_MODE=release
	//- using code:  gin.SetMode(gin.ReleaseMode)
	//
	//[GIN-debug] GET    /hello                    --> bluebell/routes.Setup.func1 (3 handlers)
	//[GIN-debug] POST   /register                 --> bluebell/controllers.Register (3 handlers)
	//[GIN-debug] [WARNING] You trusted all proxies, this is NOT safe. We recommend you to set a value.
	//	Please check https://pkg.go.dev/github.com/gin-gonic/gin#readme-don-t-trust-all-proxies for details.
	//[GIN-debug] Environment variable PORT is undefined. Using port :8080 by default
	//[GIN-debug] Listening and serving HTTP on :8080

	//gin.SetMode(gin.ReleaseMode)

	r := gin.New()

	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	v1 := r.Group("/api/v1")

	v1.GET("/hello", func(c *gin.Context) {
		c.String(http.StatusOK, settings.Conf.Version)
		fmt.Println("versin:", settings.Conf.Version)
	})

	// 注册
	v1.POST("/register", controllers.Register)

	// 登录
	v1.POST("/login", controllers.Login)

	//// test
	//r.GET("/test", func(c *gin.Context) {
	//	if
	//})

	return r
}
