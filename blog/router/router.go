package router

import (
	"blog/controller"

	"github.com/gin-gonic/gin"
)

func Start() {
	e := gin.Default()
	e.LoadHTMLGlob("templates/*")
	e.Static("/assets", "./assets")

	e.GET("/", controller.Index)

	e.GET("/login", controller.GoLogin)
	e.POST("/login", controller.Login)
	e.POST("/register", controller.RegisterUser)
	e.GET("/register", controller.GoRegister)

	e.GET("/post_index", controller.GetPostIndex)
	e.GET("/post", controller.GoAddPost)
	e.POST("/post", controller.AddPost)

	e.GET("/post_detail", controller.GoToPost)

	e.Run()
}
