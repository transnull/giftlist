package middleware

import (
	"net/http"

	"github.com/gorilla/sessions"
)

var Store *sessions.CookieStore

// RequireAuth is a middleware that checks for valid session
func RequireAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, err := Store.Get(r, "giftlist-session")
		if err != nil || session.Values["authenticated"] != true {
			http.Error(w, `{"error":"Unauthorized"}`, http.StatusUnauthorized)
			return
		}
		next(w, r)
	}
}
