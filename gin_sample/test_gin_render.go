package main

import "github.com/gin-gonic/gin"

func TestJson(c *gin.Context) {
	c.JSON(200, gin.H{
		"name": "多课网",
		"site": "www.duoke360.com",
	})
}

func TestXML(c *gin.Context) {
	c.XML(200, gin.H{
		"name": "多课网",
		"site": "www.duoke360.com",
	})
}

func TestHtml(c *gin.Context) {
	c.HTML(200, "login.html", nil)
}

func TestString(c *gin.Context) {
	c.String(200, "多课网，老郭讲golang")
}

func main() {
	e := gin.Default()

	e.GET("/test_json", TestJson)
	e.GET("/test_xml", TestXML)
	e.LoadHTMLGlob("./templates/*")
	e.GET("/test_html", TestHtml)
	e.GET("/test_string", TestString)

	e.Run()
}
