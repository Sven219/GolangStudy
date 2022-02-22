package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB // 是一个连接池对象

func initDB() (err error) {
	//数据库信息
	dsn := "root:root@tcp(10.158.1.3:3306)/sven_test"
	// 连接数据库
	db, err = sql.Open("mysql", dsn) // 不会校验用户名和密码是否正确
	if err != nil {
		fmt.Printf("dsn: %s inbvalid failed,err:%v\n", dsn, err)
		return
	}
	err = db.Ping()
	if err != nil {
		fmt.Printf("open %s failed,err:%v\n", dsn, err)
		return
	}
	// 设置数据库连接池最大连接数
	db.SetMaxOpenConns(10)
	return
}

type user struct {
	id   int
	name string
	age  int
}

// 查询
func query() {

}

func insert() {}

func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("init DB failed ,err :%v\n", err)
	}
	fmt.Println("successful！")
	var u1 user
	// 1. 写查询单条记录的 sql 语句
	sqlStr := "select id,name,age from user where id =?;"
	// 2. 执行
	rowObj := db.QueryRow(sqlStr, 1) // 从连接池里拿一个连接去数据库查询单条记录
	rowObj.Scan(&u1.id, &u1.name, &u1.age)
	// 打印结果
	fmt.Printf("u1:%#v\n", u1)

}
