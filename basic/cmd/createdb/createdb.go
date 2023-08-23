package main

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

var schema = `
DROP TABLE IF EXISTS user;
CREATE TABLE offers (
	id       TEXT     PRIMARY KEY,
    products TEXT     NULL,
    duration INTEGER  NULL,
	email    TEXT     NULL,
	price    INTEGER  NULL
);`

func main() {
	db, err := sqlx.Connect("sqlite3", "data.db")
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	db.MustExec(schema)

}
