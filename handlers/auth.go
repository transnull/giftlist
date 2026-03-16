package handlers

import (
	"encoding/json"
	"net/http"
	"os"

	"giftlist/middleware"

	"github.com/gorilla/sessions"
)

type loginRequest struct {
	Password string `json:"password"`
}

// Login handles POST /api/login
func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, `{"error":"Method not allowed"}`, http.StatusMethodNotAllowed)
		return
	}

	var req loginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, `{"error":"Invalid request body"}`, http.StatusBadRequest)
		return
	}

	adminPassword := os.Getenv("ADMIN_PASSWORD")
	if adminPassword == "" {
		adminPassword = "admin123"
	}

	if req.Password != adminPassword {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid password"})
		return
	}

	session, _ := middleware.Store.Get(r, "giftlist-session")
	session.Values["authenticated"] = true
	session.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 7,
		HttpOnly: true,
	}
	session.Save(r, w)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Login successful"})
}

// Logout handles POST /api/logout
func Logout(w http.ResponseWriter, r *http.Request) {
	session, _ := middleware.Store.Get(r, "giftlist-session")
	session.Values["authenticated"] = false
	session.Options.MaxAge = -1
	session.Save(r, w)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Logged out"})
}

// CheckAuth handles GET /api/auth/check
func CheckAuth(w http.ResponseWriter, r *http.Request) {
	session, err := middleware.Store.Get(r, "giftlist-session")
	if err != nil || session.Values["authenticated"] != true {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]bool{"authenticated": false})
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]bool{"authenticated": true})
}
