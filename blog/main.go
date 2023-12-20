package main

import (
	"blog/controller"
	"blog/dao"
)

func main() {
	dao.TestDao()
	controller.TestController()
}
