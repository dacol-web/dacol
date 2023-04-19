package controllers

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func Conn() *sql.DB {
	conn, err := sql.Open("mysql", "root:21345678@/dacol")
	if err != nil {
		panic(err)
	}
	return conn
}
