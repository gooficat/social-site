package main

import (
	"time"

	"github.com/google/uuid"
)

type Session struct {
	UserID    int
	ExpiresAt time.Time
}

var (
	sessions map[string]Session
)

func InitSessions() {
	sessions = make(map[string]Session)
}

func CreateSession(user_id int) string {
	session_id := uuid.New().String()
	sessions[session_id] = Session{
		UserID:    user_id,
		ExpiresAt: time.Now().Add(time.Hour * 24 * 30),
	}

	return session_id
}

func GetUserId(session_id string) int {
	if session, ok := sessions[session_id]; ok {
		return session.UserID
	}
	return 0
}

func DeleteSession(session_id string) int8 {
	if _, ok := sessions[session_id]; !ok {
		return 1
	}
	delete(sessions, session_id)
	return 0
}

func SweepSessions() {
	now := time.Now()
	for session_id, session := range sessions {
		if session.ExpiresAt.Before(now) {
			delete(sessions, session_id)
		}
	}
}
