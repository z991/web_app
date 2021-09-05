package routes

import (
	"net/http"
	"web_app/controller"
	"web_app/logger"
	"web_app/middlewares"
	"web_app/settings"

	"github.com/gin-gonic/gin"
)

func Setup(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	// 用户注册
	r.POST("/signup", controller.SignUpHandler)
	// 用户登录
	r.POST("/login", controller.LoginHandler)
	r.GET("/ping", middlewares.JWTAuthMiddleware(), func(c *gin.Context) {
		c.Request.Header.Get("Authorization")
		c.String(http.StatusOK, "pong")
	})
	r.GET("/version", func(c *gin.Context) {
		c.String(http.StatusOK, settings.Conf.Version)
	})
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})
	return r
}
