package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func InitDb() {
	var err error
	db, err = sql.Open("sqlite3", "./db.sqlite")
	if err != nil {
		panic(err)
	}
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY,
		name TEXT NOT NULL,
		email TEXT NOT NULL,
		handle TEXT UNIQUE NOT NULL,
		password TEXT NOT NULL
	)`)
	if err != nil {
		panic(err)
	}
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS connections (
		id INTEGER PRIMARY KEY,
		follower INTEGER NOT NULL,
		following INTEGER NOT NULL,
		FOREIGN KEY (follower) REFERENCES users(id),
		FOREIGN KEY (following) REFERENCES users(id)
	)`)
	if err != nil {
		panic(err)
	}
}
