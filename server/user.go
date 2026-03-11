package main

import (
	"errors"
	"log"
	"database/sql"
)

type User struct {
	id int
	name string
	email string
	handle string
	password string
}

func GetUserByID(id int) *User {
	var user User
	err := db.QueryRow("SELECT id, name, email, handle, password FROM USERS WHERE id = $1", id).Scan(&user.id, &user.name, &user.email, &user.handle, &user.password)
	
	if errors.Is(err, sql.ErrNoRows) {
		return nil
	}
	if err != nil {
		log.Fatal(err)
	}

	return &user
}

func GetUserByHandle(handle string) *User {
	var user User
	err := db.QueryRow("SELECT id, name, email, handle, password FROM USERS WHERE handle = $1", handle).Scan(&user.id, &user.name, &user.email, &user.handle, &user.password)
	
	if errors.Is(err, sql.ErrNoRows) {
		return nil
	}
	if err != nil {
		log.Fatal(err)
	}
	
	return &user
}

func CreateUser(name string, handle string, email string, password string) {
	_, err := db.Exec("INSERT INTO users (name, email, handle, password) VALUES ($1, $2, $3, $4)", name, email, handle, password)

	if err != nil {
		log.Fatal(err)
	}
}