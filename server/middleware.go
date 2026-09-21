package main

import (
	"context"
	"net/http"
)

type contextKey string

const (
	sessionUserIDKey contextKey = "session_user_id"
	hardcodedUserID  string     = "fake-user-123"
)

// fakeSessionMiddleware injects a hardcoded user ID into the request context to simulate an authenticated session.
func fakeSessionMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), sessionUserIDKey, hardcodedUserID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func getUserIDFromSession(r *http.Request) string {
	if val, ok := r.Context().Value(sessionUserIDKey).(string); ok {
		return val
	}
	return ""
}
