package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// 定义一个全局对象db
var db *sql.DB

// 定义一个初始化数据库的函数
func initDB() (err error) {
	dsn := "payroll:123@tcp(127.0.0.1:3306)/go_db"
	// 不会校验账号密码是否正确
	// 注意！！！这里不要使用:=，我们是给全局变量赋值，然后在main函数中使用全局变量db
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	// 尝试与数据库建立连接（校验dsn是否正确）
	err = db.Ping()
	if err != nil {
		return err
	}
	return nil
}

type user struct {
	id       int
	username string
	password string
}

// 查询一条用户数据
func queryRowDemo() {
	sqlStr := "select id, username, password from user_tbl where id=?"
	var u user
	// 确保QueryRow之后调用Scan方法，否则持有的数据库链接不会被释放
	err := db.QueryRow(sqlStr, 1).Scan(&u.id, &u.username, &u.password)
	if err != nil {
		fmt.Printf("scan failed, err:%v\n", err)
		return
	}
	fmt.Printf("id:%d name:%s age:%s\n", u.id, u.username, u.password)
}

// 查询多条数据示例
func queryMultiRow() {
	sqlStr := "select id, username, password from user_tbl where id > ?"
	rows, err := db.Query(sqlStr, 0)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return
	}
	// 非常重要：关闭rows释放持有的数据库链接
	defer rows.Close()

	// 循环读取结果集中的数据
	for rows.Next() {
		var u user
		err := rows.Scan(&u.id, &u.username, &u.password)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return
		}
		fmt.Printf("id:%d username:%s password:%s\n", u.id, u.username, u.password)
	}
}

func insertData() {
	sqlstr := "insert into user_tbl(username, password) values(?,?)"
	ret, err := db.Exec(sqlstr, "张三", "zs")
	if err != nil {
		log.Printf("insert faild, err: %v", err)
		return
	}

	theID, err := ret.LastInsertId()
	if err != nil {
		log.Printf("insert error, err: %v\n", err)
	}
	log.Printf("insert success, the id is %d.\n", theID)

}

func delData() {
	sql := "delete from user_tbl where id =?"
	ret, err := db.Exec(sql, "1")
	if err != nil {
		fmt.Printf("删除失败, err:%v\n", err)
		return
	}
	rows, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("删除行失败, err:%v\n", err)
		return
	}
	fmt.Printf("删除成功, 删除的行数： %d.\n", rows)
}

func updateData() {
	sql := "update user_tbl set username=?, password=? where id=?"
	ret, err := db.Exec(sql, "kite2", "kite123", "2")
	if err != nil {
		fmt.Printf("更新失败, err:%v\n", err)
		return
	}
	rows, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("更新行失败, err:%v\n", err)
		return
	}
	fmt.Printf("更新成功, 更新的行数： %d.\n", rows)
}

func main() {
	err := initDB() // 调用输出化数据库的函数
	if err != nil {
		fmt.Printf("初始化失败！,err:%v\n", err)
		return
	} else {
		fmt.Printf("初始化成功 \n")
	}

	// queryRowDemo()
	queryMultiRow()
	insertData()
	delData()
	updateData()
}
