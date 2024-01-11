package main

import (
	"fmt"

	"github.com/casbin/casbin"
	gormadapter "github.com/casbin/gorm-adapter"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	// 使用 MySQL 数据库初始化一个 gorm 适配器
	a := gormadapter.NewAdapter("mysql", "root:123456@tcp(127.0.0.1:3306)/test_db", true)

	e := casbin.NewEnforcer("rbac_model.conf", a)

	sub := "zhou"  // 想要访问资源的用户。
	obj := "data1" // 将被访问的资源。
	act := "read"  // 用户对资源执行的操作。

	// e.AddPolicy(sub, obj, act)

	// e.ClearPolicy()

	ok := e.Enforce(sub, obj, act)

	if ok {
		// 允许ghz读取data1
		fmt.Println("yes")
	} else {
		// 拒绝请求，抛出异常
		fmt.Println("no")
	}
}
