package db

import (
	"database/sql"
	"fmt"

	"github.com/bayuwidia/echo-rest/config"

	_ "github.com/lib/pq"
)

var db *sql.DB

var err error

func Init() {

	conf := config.GetConfig()

	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s search_path=%s sslmode=disable", conf.DB_HOST, conf.DB_PORT, conf.DB_USERNAME, conf.DB_PASSWORD, conf.DB_NAME, conf.DB_SCHEMA)

	db, err = sql.Open("postgres", connectionString)
	if err != nil {
		panic("connectionString error")
	}

	err = db.Ping()
	if err != nil {
		panic("dsn invalid")
	}
}

func CreateCon() *sql.DB {
	return db
}
