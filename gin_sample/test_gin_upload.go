package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func Upload(c *gin.Context) {
	file, _ := c.FormFile("file")
	log.Println(file.Filename)

	c.SaveUploadedFile(file, file.Filename)
	c.String(200, fmt.Sprintf("%s uploaded!", file.Filename))

}

func GoUpload(c *gin.Context) {
	c.HTML(200, "upload.html", nil)
}

func main() {
	router := gin.Default()
	// 为 multipart forms 设置较低的内存限制 (默认是 32 MiB)
	router.MaxMultipartMemory = 8 << 20 // 8 MiB
	router.LoadHTMLGlob("./templates/*")
	router.GET("/upload", GoUpload)
	router.POST("/upload", Upload)
	router.Run()
}
