package main

import (
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	c.HTML(200, "login.html", nil)
}

func Dologin(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	c.HTML(200, "welcome.html", gin.H{
		"username": username,
		"password": password,
	})
}

func main() {
	e := gin.Default()
	e.LoadHTMLGlob("./templates/*")
	e.GET("/login", Login)
	e.POST("/login", Dologin)

	e.Run()
}
