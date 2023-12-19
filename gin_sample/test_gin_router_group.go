package main

import "github.com/gin-gonic/gin"

func F1(c *gin.Context) {}
func F2(c *gin.Context) {}
func F3(c *gin.Context) {}
func F4(c *gin.Context) {}
func F5(c *gin.Context) {}
func F6(c *gin.Context) {}

func main() {
	router := gin.Default()

	// 博客
	// 访问：http://localhost:8080/blog/list
	v1 := router.Group("/blog")
	{
		v1.GET("/list", F1)
		v1.POST("/post", F2)
		v1.POST("/add", F3)
	}

	// 视频
	// 访问：http://localhost:8080/video/list
	v2 := router.Group("/video")
	{
		v2.GET("/list", F4)
		v2.POST("/post", F5)
		v2.POST("/add", F6)
	}

	router.Run(":8080")
}
