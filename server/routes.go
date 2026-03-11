package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"path"
	"strings"
)

func ApiUser(w http.ResponseWriter, r *http.Request, rt []string) {
	if len(rt) < 1 {
		fmt.Fprintf(w, "Hello from the user api\n")
		return
	}
	switch rt[0] {
	case "login":
		ApiUserLogin(w, r)
	case "register":
		ApiUserRegister(w, r)
	case "logout":
		ApiUserLogout(w, r)
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

func ApiUserLogin(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Handle   string `json:"handle"`
		Password string `json:"password"`
	}

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		WriteJson(w, http.StatusBadRequest, map[string]any{
			"message": "malformed body",
		})
		return
	}

	user := GetUserByHandle(body.Handle)
	if user == nil {
		WriteJson(w, http.StatusBadRequest, map[string]any{
			"message": "user does not exist",
		})
		return
	}

	if user.password != body.Password {
		WriteJson(w, http.StatusBadRequest, map[string]any{
			"message": "incorrect password",
		})
		return
	}

	WriteJson(w, http.StatusOK, map[string]any{
		"session_id": CreateSession(user.id),
	})
}

func ApiUserLogout(w http.ResponseWriter, r *http.Request) {
	var body struct {
		SessionId string `json:"session_id"`
	}

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil || body.SessionId == "" {
		WriteJson(w, http.StatusBadRequest, map[string]any{
			"message": "malformed body",
		})
		return
	}
	err2 := DeleteSession(body.SessionId)
	if err2 != 0 {
		WriteJson(w, http.StatusBadRequest, map[string]any{
			"message": "session does not exist",
		})
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func ApiUserRegister(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Name     string `json:"name"`
		Handle   string `json:"handle"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		WriteJson(w, http.StatusBadRequest, map[string]any{
			"message": "malformed body",
		})
		return
	}

	if body.Name == "" || body.Handle == "" || body.Email == "" || body.Password == "" {
		WriteJson(w, http.StatusBadRequest, map[string]any{
			"message": "All fields are required",
		})
		return
	}

	user := GetUserByHandle(body.Handle)
	if user != nil {
		WriteJson(w, http.StatusBadRequest, map[string]any{
			"message": "user with that handle already exists",
		})
		return
	}
	user = GetUserByEmail(body.Email)
	if user != nil {
		WriteJson(w, http.StatusBadRequest, map[string]any{
			"message": "user with that email already exists",
		})
		return
	}

	CreateUser(body.Name, body.Handle, body.Email, body.Password)

	newUser := GetUserByHandle(body.Handle)

	WriteJson(w, http.StatusOK, map[string]any{
		"session_id": CreateSession(newUser.id),
	})
}
