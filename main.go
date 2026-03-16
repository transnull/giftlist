package main

import (
	"crypto/rand"
	"embed"
	"encoding/hex"
	"io/fs"
	"log"
	"net/http"
	"os"
	"strings"

	"giftlist/handlers"
	"giftlist/middleware"
	"giftlist/models"

	"github.com/gorilla/sessions"
)

//go:embed static/*
var staticFiles embed.FS

func main() {
	// Initialize session store
	secret := os.Getenv("SESSION_SECRET")
	if secret == "" {
		b := make([]byte, 32)
		rand.Read(b)
		secret = hex.EncodeToString(b)
		log.Println("WARNING: SESSION_SECRET not set, using random secret (sessions won't persist across restarts)")
	}
	middleware.Store = sessions.NewCookieStore([]byte(secret))

	// Check admin password
	if os.Getenv("ADMIN_PASSWORD") == "" {
		log.Println("WARNING: ADMIN_PASSWORD not set, defaulting to 'admin123'")
	}

	// Initialize database
	dbPath := os.Getenv("DB_PATH")
	if dbPath == "" {
		dbPath = "giftlist.db"
	}
	models.InitDB(dbPath)

	// Setup routes
	mux := http.NewServeMux()

	// API routes
	mux.HandleFunc("/api/login", handlers.Login)
	mux.HandleFunc("/api/logout", handlers.Logout)
	mux.HandleFunc("/api/auth/check", handlers.CheckAuth)
	mux.HandleFunc("/api/config", configRouter)
	mux.HandleFunc("/api/items", itemsRouter)
	mux.HandleFunc("/api/items/", itemRouter)

	// Static files (embedded)
	staticFS, err := fs.Sub(staticFiles, "static")
	if err != nil {
		log.Fatalf("Failed to create static sub-filesystem: %v", err)
	}
	fileServer := http.FileServer(http.FS(staticFS))

	// SPA fallback: serve index.html for non-API, non-asset routes
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		// Try to serve static file
		if path != "/" && !strings.HasPrefix(path, "/assets/") {
			// For SPA routes, serve index.html
			_, err := staticFiles.Open("static" + path)
			if err != nil {
				// File not found, serve index.html for SPA routing
				indexContent, err := staticFiles.ReadFile("static/index.html")
				if err != nil {
					http.Error(w, "index.html not found", http.StatusInternalServerError)
					return
				}
				w.Header().Set("Content-Type", "text/html; charset=utf-8")
				w.Write(indexContent)
				return
			}
		}
		fileServer.ServeHTTP(w, r)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("🎁 Gift Wishlist server starting on http://localhost:%s", port)
	log.Printf("📋 Admin panel: http://localhost:%s/admin", port)
	if err := http.ListenAndServe(":"+port, mux); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}

func configRouter(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		handlers.GetConfig(w, r)
	case http.MethodPut:
		middleware.RequireAuth(handlers.UpdateConfig)(w, r)
	default:
		http.Error(w, `{"error":"Method not allowed"}`, http.StatusMethodNotAllowed)
	}
}

func itemsRouter(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		handlers.GetItems(w, r)
	case http.MethodPost:
		middleware.RequireAuth(handlers.CreateItem)(w, r)
	default:
		http.Error(w, `{"error":"Method not allowed"}`, http.StatusMethodNotAllowed)
	}
}

func itemRouter(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		handlers.GetItem(w, r)
	case http.MethodPut:
		middleware.RequireAuth(handlers.UpdateItem)(w, r)
	case http.MethodDelete:
		middleware.RequireAuth(handlers.DeleteItem)(w, r)
	default:
		http.Error(w, `{"error":"Method not allowed"}`, http.StatusMethodNotAllowed)
	}
}
