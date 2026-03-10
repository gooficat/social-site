package main

import (
	"fmt"
	"net/http"
	"strings"
	"path"
)

func ApiRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from the api\n")
}
func Routes(w http.ResponseWriter, r *http.Request) {
	rt := strings.Split(r.URL.Path, "/")[1:]


	switch rt[0] {
	case "api":
		ApiRoute(w, r)
	default:
		http.ServeFile(w, r, path.Join("../static", r.URL.Path))
	}
}

func main() {
	fmt.Println("Hello from the server")
	http.HandleFunc("/", Routes)
	http.ListenAndServe(":3000", nil)
}
