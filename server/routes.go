package main

import (
	"fmt"
	"net/http"
	"strings"
	"path"
	"encoding/json"
)

func ApiUser(w http.ResponseWriter, r *http.Request, rt []string) {
	if len(rt) < 1 {
		fmt.Fprintf(w, "Hello from the user api\n")
		return
	}
	switch rt[0] {
	case "login":
		ApiUserLogin(w, r, rt[1:])
	case "register":
		ApiUserRegister(w, r, rt[1:])
	default:
		fmt.Fprintf(w, "Unknown path\n")
	}
}

func Api(w http.ResponseWriter, r *http.Request, rt []string) {
	if len(rt) < 1 {
		fmt.Fprintf(w, "Hello from the api\n")
		return
	}
	switch rt[0] {
	case "user":
		ApiUser(w, r, rt[1:])
	default:
		fmt.Fprintf(w, "Unknown path\n")
	}
}
func Routes(w http.ResponseWriter, r *http.Request) {
	rt := strings.Split(r.URL.Path, "/")[1:]
	switch rt[0] {
	case "api":
		Api(w, r, rt[1:])
	default:
		http.ServeFile(w, r, path.Join("../static", r.URL.Path))
	}
}

func WriteJson(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func ApiUserLogin(w http.ResponseWriter, r *http.Request, rt []string) {
	var body struct {
		handle string `json:"handle"`
		password string `json:"password"`
	}

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		WriteJson(w, http.StatusBadRequest, map[string]any{
			"message": "malformed body",
		})
		return
	}

	user := GetUserByHandle(body.handle)
	if user == nil {
		WriteJson(w, http.StatusBadRequest, map[string]any{
			"message": "user does not exist",
		})
		return
	}

	WriteJson(w, http.StatusOK, map[string]any{
		"session_id": CreateSession(user.id),
	})
}

func ApiUserRegister(w http.ResponseWriter, r *http.Request, rt []string) {
	var body struct {
		name string `json:"name"`
		handle string `json:"handle"`
		email string `json:"email"`

		password string `json:"password"`
	}

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		WriteJson(w, http.StatusBadRequest, map[string]any{
			"message": "malformed body",
		})
		return
	}

	user := GetUserByHandle(body.handle)
	if user != nil {
		WriteJson(w, http.StatusBadRequest, map[string]any{
			"message": "user with that handle already exists",
		})
		return
	}
	user = GetUserByEmail(body.email)
	if user != nil {
		WriteJson(w, http.StatusBadRequest, map[string]any{
			"message": "user with that email already exists",
		})
		return
	}

	newUser := CreateUser(body.name, body.handle, body.email, body.password)
	WriteJson(w, http.StatusOK, map[string]any{
		"session_id": CreateSession(newUser.id),
	})
}
