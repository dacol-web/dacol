package controllers

import (
	"database/sql"
	"fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func Conn() *sql.DB {
	conn, err := sql.Open("mysql", "root:21345678@/dacol")
	if err != nil {
		panic(err)
	}
	return conn
}

type Ctx *gin.Context

func ErrCtrl(ctrl string, err error) {
	panic(fmt.Sprintf("%s controller err : %s", ctrl, err.Error()))
}
