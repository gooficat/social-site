package main

import "github.com/google/uuid"

var (
	sessions map[string]int
)

func InitSessions() {
	sessions = make(map[string]int)
}

func CreateSession(user_id int) string {
	session_id := uuid.New().String()
	sessions[session_id] = user_id

	return session_id
}

func GetUserId(session_id string) int {
	if user_id, ok := sessions[session_id]; ok {
		return user_id
	}
	return 0
}

func DeleteSession(session_id string) {
	delete(sessions, session_id)
}
