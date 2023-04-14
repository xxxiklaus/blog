package router

import (
	"blog/controller"

	"github.com/gin-gonic/gin"
)

func Start() {
	e := gin.Default()
	e.LoadHTMLGlob("templates/*")
	e.Static("/assets", "./assets")

	// e.GET("/index", controller.ListUser)
	e.GET("/login", controller.GoLogin)
	e.POST("/login", controller.Login)

	e.POST("/register", controller.Register)
	e.GET("/register", controller.GoRegister)

	e.GET("/post_index", controller.GetPostIndex) //博客列表
	e.POST("/post", controller.AddPost)           // 添加博客
	e.GET("/post", controller.GoAddPost)          //跳转到添加博客页面

	e.GET("/", controller.Index)
	e.Run()
}
