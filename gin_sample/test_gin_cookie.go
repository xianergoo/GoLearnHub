package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	// 1.设置cookie
	r.GET("/setCookie", func(c *gin.Context) {
		// 设置cookie
		c.SetCookie("cookie_name", "cookie_value", 3600, "/", "localhost", false, true)
		/*
			name cookie的名称
			value cookie的值
			maxAge int, 单位为秒
			path cookie所在目录
			domain string,域名
			secure 是否智能通过https访问
			httpOnly bool  是否允许通过js获取自己的cookie
		*/
		c.JSON(http.StatusOK, gin.H{
			"message": "set cookie success.",
		})
	})
	// 2.获取cookie
	r.GET("/getCookie", func(c *gin.Context) {
		// 读取cookie,根据cookie名读取
		cookie, err := c.Cookie("cookie_name")
		if err != nil {
			// 直接返回cookie值
			c.JSON(http.StatusOK, gin.H{"message": "get cookie fail"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"cookie_name": cookie})
	})
	// 3.删除cookie
	r.GET("/delCookie", func(c *gin.Context) {
		// 删除cookie, 设置cookie MaxAge设置为-1，表示删除cookie
		c.SetCookie("cookie_name", "cookie_value", -1, "/", "localhost", false, true)
		c.JSON(http.StatusOK, gin.H{
			"message": "delete cookie success.",
		})
	})

	r.Run(":8080")
}
