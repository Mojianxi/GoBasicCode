package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/cost?charset=utf8")
	if err != nil {
		panic(err)
	}
	//插入
	// stmt, err := db.Prepare("insert test set t_id=?,t_name=?,t_time=?")
	// res, err := stmt.Exec(998, "zhangsan", "2019-01-02")
	// id, err := res.LastInsertId()
	//修改
	// stmt, err := db.Prepare("update test set t_name=? where t_id=?")
	// res, err := stmt.Exec("lisi", 999)
	// id, err := res.RowsAffected()
	//数据库一主多从,从库多是用来查询的
	//查询
	rows, err := db.Query("select * from test where t_id=999")
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		var t_id int
		var t_name string
		var t_time string
		err = rows.Scan(&t_id, &t_name, &t_time)
		fmt.Println(t_id, t_name, t_time)
	}
	//在go中创建的变量必须调用一次
	// fmt.Println(id)
}
