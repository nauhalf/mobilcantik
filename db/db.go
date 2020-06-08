package db

import (
	"fmt"
	"os"
	"strconv"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var DBCon *sql.DB

func Open() {
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPw := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbPort, _ := strconv.Atoi(os.Getenv("DB_PORT"))

	conString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=UTC", dbUser, dbPw, dbHost, dbPort, dbName)

	database, err := sql.Open("mysql", conString)
	if err != nil {
		panic(err.Error())
	}

	DBCon = database
}
