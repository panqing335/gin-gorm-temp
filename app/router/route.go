package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"temp/app/controller/login"
	"temp/app/middleware"
	_ "temp/config"
)

var router *gin.Engine

func Setup() {
	router = gin.New()
	router.Use(gin.Logger())
	router.Use(middleware.Recover)
	router.Use(middleware.Cors)
	GroupDefault()
	GroupUser()

	port := viper.GetString("server.port")
	fmt.Println("当前端口：", port)
	err := router.Run(":" + port)
	if err != nil {
		return
	}
}

func GroupDefault() {
	router.POST("/login", login.Login)
	router.POST("/register", login.Register)
	router.GET("/paginator", login.Paginator)
}

func GroupUser() *gin.RouterGroup {
	v1 := router.Group("user")
	{
		v1.Use(middleware.Auth)
		v1.POST("/edit", login.Edit)
		v1.DELETE("/destroy", login.Destroy)
	}
	return v1
}
