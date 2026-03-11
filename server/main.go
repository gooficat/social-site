package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Hello from the server")
	InitDb()
	InitSessions()
	CronJobs()

	http.HandleFunc("/", Routes)
	http.ListenAndServe(":3000", nil)
}
