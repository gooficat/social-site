package main

import (
	"fmt"
	// "net/http"
)


func main() {
	fmt.Println("Hello from the server")
	InitDb()

	CreateUser("Hasan Sabri", "gooficat", "hsabri2007@gmail.com", "..1134..")

	user := GetUserByHandle("gooficat")

	if user != nil {
		fmt.Printf("User %s (%s) email %s, password %s\n", user.name, user.handle, user.email, user.password)
	}

	// http.HandleFunc("/", Routes)
	// http.ListenAndServe(":3000", nil)
}
