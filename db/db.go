package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)


var globalSession *sql.DB

func InitDB(dbUrl string) {
	db, err := sql.Open("mysql", dbUrl)
	if err == nil {
		globalSession = db
	} else {
		panic("db open failed: " + err.Error())
	}

	err = db.Ping()
	if err != nil {
		panic("db ping failed" + err.Error())
	}
}

func QueryForInt(query string) (int, error) {
	var res int
	err := globalSession.QueryRow(query).Scan(&res)

	return res, err
}
