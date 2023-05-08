package database

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/juandaantoniusapakpahan/go-restful-api/helper"
)

func Connect() *sql.DB {
	db, err := sql.Open("mysql", "root:@root123@tcp(localhost:3306)/belajar_golang_result_api")
	helper.ErrorHandle(err)
	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)
	return db
}
