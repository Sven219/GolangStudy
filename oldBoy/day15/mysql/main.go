package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	//数据库信息
	dsn := "root:root@tcp(10.158.1.3:3306)/sven_test"
	// 连接数据库
	db, err := sql.Open("mysql", dsn) // 不会校验用户名和密码是否正确
	if err != nil {
		fmt.Printf("dsn: %s inbvalid failed,err:%v\n", dsn, err)
		return
	}
	err = db.Ping()
	if err != nil {
		fmt.Printf("open %s failed,err:%v\n", dsn, err)
		return
	}
	db.Close()
	fmt.Println("successful！")
}
