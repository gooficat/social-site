package main

import (
	"fmt"
	"net/http"
)


func main() {
	fmt.Println("Hello from the server")
	InitDb()
	http.HandleFunc("/", Routes)
	http.ListenAndServe(":3000", nil)
}
