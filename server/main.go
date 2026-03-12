package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Hello from the server")
	InitDb()
	CronJobs()

	http.HandleFunc("/", Routes)
	http.ListenAndServe(":8080", nil)
}
