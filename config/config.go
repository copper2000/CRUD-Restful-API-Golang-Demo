package config

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func GetDatabase() (db *sql.DB, err error) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "root"
	dbName := "golang_demo"
	db, err = sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	return
}
