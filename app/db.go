package app

import (
	"database/sql"
	"time"

	"github.com/juandaantoniusapakpahan/go-restful-api/helper"
)

func Connect() *sql.DB {
	db, err := sql.Open("mysql", "root:@root123@tcp(localhost:3306)/belajar_golang_database_migration")
	helper.PanicIfError(err)
	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)
	return db

	// migrate create -ext sql -dir db/migration create_first

	// migrate -database "mysql://root:@root123@tcp(localhost:3306)/belajar_golang_database_migration" -path db/migration up

	// migrate -database "mysql://root:@root123@tcp(localhost:3306)/belajar_golang_database_migration" -path db/migration down
	// migrate -database "mysql://root:@root123@tcp(localhost:3306)/belajar_golang_database_migration" -path db/migration version
	// migrate -database "mysql://root:@root123@tcp(localhost:3306)/belajar_golang_database_migration" -path db/migration force version before dirty

}
