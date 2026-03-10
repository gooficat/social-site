package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Hello, World!")

	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("../static"))))

	http.ListenAndServe(":3000", nil)
}