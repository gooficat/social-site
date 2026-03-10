package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"sync"
)

var (
	db *sql.DB
	once sync.Once
)

func InitDb() {
	var err error
	db, err = sql.Open("sqlite3", "./db.sqlite")
	if err != nil {
		panic(err)
	}
	db.Exec(`CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY_KEY AUTO_INCREMENT,
		name TEXT NOT NULL,
		handle TEXT UNIQUE NOT NULL,
		password TEXT NOT NULL
	)`)
	db.Exec(`CREATE TABLE IF NOT EXISTS connections
		id INTEGER PRIMARY_KEY AUTO_INCREMENT,
		follower INTEGER NOT NULL,
		following INTEGER NOT NULL,
		FOREIGN KEY (follower) REFERENCES users(follower),
		FOREIGN KEY (following) REFERENCES users(following),
	`)
}