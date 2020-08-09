package ch9

import (
	"database/sql"
	"fmt"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestMysqlOpen(t *testing.T) {
	db, err := sql.Open("mysql", "admin:kingdee8349@tcp(mysql-dev.development:3306)/test_db")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
}

func TestMysqlPing(t *testing.T) {
	db, err := sql.Open("mysql", "admin:kingdee8349@tcp(mysql-dev.development:3306)/test_db")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}
}

func TestMysqlQuery(t *testing.T) {
	db, err := sql.Open("mysql", "admin:kingdee8349@tcp(mysql-dev.development:3306)/shc_user")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()
	rows, err := db.Query("select id,title from user_msg")
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		var id string
		var tit string
		err = rows.Scan(&id, &tit)
		if err != nil {
			panic(err)
		}
		fmt.Println("id:", id)
		fmt.Println("title:", tit)
		fmt.Println("------------------------------------------")
	}
}

func TestMysqlQuery2(t *testing.T) {
	db, err := sql.Open("mysql", "admin:kingdee8349@tcp(mysql-dev.development:3306)/shc_user")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()
	rows, err := db.Query("SELECT * FROM user_msg")
	if err != nil {
		panic(err)
	}
	columns, err := rows.Columns()
	if err != nil {
		panic(err)
	}
	values := make([]sql.RawBytes, len(columns))
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}
	for rows.Next() {
		err = rows.Scan(scanArgs...)
		if err != nil {
			panic(err)
		}
		var value string
		for i, col := range values {
			if col == nil {
				value = "NULL"
			} else {
				value = string(col)
			}
			fmt.Println(columns[i], ":", value)
		}
		fmt.Println("---------------------------------------")
	}
	if err = rows.Err(); err != nil {
		panic(err)
	}
}
