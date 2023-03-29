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

	v1 := r.Group("/api/v1")

	// 注册业务路由
	v1.POST("/signup", controller.SignUpHandler)
	// 登录业务路由
	v1.POST("/login", controller.LoginHandler)

	// 应用jwt应用中间件
	v1.Use(middlewares.JWTAuthMiddleware())

	{
		v1.GET("/community", controller.CommunityHandler)
		v1.GET("/community/:id", controller.CommunityDetailHandler)
	}

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, fmt.Sprintf("项目名称:%v\n", viper.GetString("app.name")))
	})
	// 检测脚手架版本
	r.GET("/version", func(c *gin.Context) {
		c.String(http.StatusOK, fmt.Sprintf("版本号:%s", viper.GetString("app.version")))
	})

	return r
}
