package router

import (
	"bluebell/controller"
	"bluebell/logger"
	"bluebell/middlewares"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
)

func Setup() *gin.Engine {
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	// 注册业务路由
	r.POST("/signup", controller.SignUpHandler)
	// 登录业务路由
	r.POST("/login", controller.LoginHandler)

	// 加入JWT验证中间件
	r.GET("/ping", middlewares.JWTAuthMiddleware(), func(c *gin.Context) {
		// token认证, 判断是否登录, 判断请求头中是否有有效的jwt
		// 从请求头中取出token
		c.String(http.StatusOK, "pong")
	})

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, fmt.Sprintf("项目名称:%v\n", viper.GetString("app.name")))
	})
	// 检测脚手架版本
	r.GET("/version", func(c *gin.Context) {
		c.String(http.StatusOK, fmt.Sprintf("版本号:%s", viper.GetString("app.version")))
	})

	return r
}
