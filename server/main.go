package main

import (
	"fmt"
	"net/http"
	"strings"
	"path"
)

func ApiUserRoute(w http.ResponseWriter, r *http.Request, rt []string) {
	if len(rt) < 1 {
		fmt.Fprintf(w, "Hello from the user api\n")
		return
	}
	switch rt[0] {
	case "login":
		fmt.Fprintf(w, "Login\n")
	case "register":
		fmt.Fprintf(w, "Register\n")
	default:
		fmt.Fprintf(w, "Unknown path\n")
	}
}

func ApiRoute(w http.ResponseWriter, r *http.Request, rt []string) {
	if len(rt) < 1 {
		fmt.Fprintf(w, "Hello from the api\n")
		return
	}
	switch rt[0] {
	case "user":
		ApiUserRoute(w, r, rt[1:])
	default:
		fmt.Fprintf(w, "Unknown path\n")
	}
}
func Routes(w http.ResponseWriter, r *http.Request) {
	rt := strings.Split(r.URL.Path, "/")[1:]
	switch rt[0] {
	case "api":
		ApiRoute(w, r, rt[1:])
	default:
		http.ServeFile(w, r, path.Join("../static", r.URL.Path))
	}
}

func main() {
	fmt.Println("Hello from the server")
	http.HandleFunc("/", Routes)
	http.ListenAndServe(":3000", nil)
}
