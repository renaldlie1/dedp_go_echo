package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/mychengel/depd_go_echo/config"
)

var db *sql.DB
var err error

func Init() {
	conf := config.GetConfig()
	conn := conf.DB_USERNAME + ":" + conf.DB_PASSWORD + "@tcp(" + conf.DB_HOST + ":" + conf.DB_PORT + ")/" + conf.DB_NAME

	db, err = sql.Open("mysql", conn)
	if err != nil {
		panic("Connection failed/error!")
	}

	err = db.Ping()
	if err != nil {
		panic("DSN Invalid!")
	}
}

func CreateCon() *sql.DB {
	return db
}
