package web

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes() *gin.Engine {
	server := gin.Default()
	registerUserRounts(server)

	return server
}

func registerUserRounts(server *gin.Engine) {
	//注册
	u := &UserHandler{}
	ug := server.Group("/users")
	ug.POST("/signup", u.Signup)
	ug.POST("/login", u.Login)
	ug.POST("/edit", u.Edit)
	ug.GET("profile", u.Profile)
	//server.POST("/users/signup", u.Signup)
	////登录
	//server.POST("/users/login", u.Login)
	////编辑
	//server.POST("users/edit", u.Edit)
	////用户信息
	//server.GET("/users/profile", func(ctx *gin.Context) {
	//})
	server.Run(":8080")
}
