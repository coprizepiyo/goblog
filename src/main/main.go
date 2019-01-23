package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	query()
}

func check(err error) {
	if err != nil{
		fmt.Println(err)
		panic(err)
	}
}

type MyTable struct {
	tData []map[string]string
}

var mt = MyTable{}

func query() {
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/blog")
	check(err)

	rows, err := db.Query("SELECT * FROM blog.blog_user")
	check(err)

	// var table = make([]interface{}, 0)

	for rows.Next() {
		columns, _ := rows.Columns()

		scanArgs := make([]interface{}, len(columns))
		values := make([]interface{}, len(columns))

		for i := range values {
			scanArgs[i] = &values[i]
		}

		//将数据保存到 record 字典
		err = rows.Scan(scanArgs...)
		record := make(map[string]string)
		for i, col := range values {
			if col != nil {
				record[columns[i]] = string(col.([]byte))
			}
		}
		fmt.Println(record)
		// fmt.Println(record["user_name"])
		mt.tData = append(mt.tData, record)
	}
	rows.Close()

	// fmt.Println(table)
	fmt.Println(mt.tData[0]["user_name"])
}

