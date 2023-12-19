package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func TestMW(c *gin.Context) {
	c.String(200, "hello world", nil)
}

func MyMiddleware1(c *gin.Context) {
	fmt.Println("my first middle ware")
}

func MyMiddleware2(c *gin.Context) {
	fmt.Println("my second middle ware")
}

func main() {

	// func Default() *Engine {
	// 	debugPrintWARNINGDefault()
	// 	engine := New()
	// 	engine.Use(Logger(), Recovery())
	// 	return engine
	// }

	e := gin.Default()
	e.Use(MyMiddleware1, MyMiddleware2)
	e.GET("testmw", TestMW)
	e.Run()
}
