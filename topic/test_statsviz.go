package main

import (
	"log"
	"net/http"

	"github.com/arl/statsviz"
)

func main() {
	statsviz.RegisterDefault()

	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	// fmt.Println("访问地址： http://localhost:8085/debug/statsviz/\n\n")

	// router := gin.New()
	// router.GET("/debug/statsviz/*filepath", func(context *gin.Context) {
	// 	if context.Param("filepath") == "/ws" {
	// 		statsviz.Ws(context.Writer, context.Request)
	// 		return
	// 	}
	// 	statsviz.IndexAtRoot("/debug/statsviz").ServeHTTP(context.Writer, context.Request)
	// })
	// router.Run(":8085")
}
