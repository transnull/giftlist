package models

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB(dataSourceName string) {
	var err error
	DB, err = sql.Open("sqlite3", dataSourceName)
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	createTables()
	insertDefaults()
	log.Println("Database initialized successfully")
}

func createTables() {
	queries := []string{
		`CREATE TABLE IF NOT EXISTS config (
			key TEXT PRIMARY KEY,
			value TEXT NOT NULL
		)`,
		`CREATE TABLE IF NOT EXISTS items (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			category TEXT,
			priority TEXT NOT NULL CHECK(priority IN ('high', 'medium', 'low')),
			estimated_cost REAL,
			image_url TEXT NOT NULL,
			product_link TEXT,
			remark TEXT,
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP
		)`,
	}

	for _, q := range queries {
		if _, err := DB.Exec(q); err != nil {
			log.Fatalf("Failed to create table: %v", err)
		}
	}
}

func insertDefaults() {
	defaults := []struct{ key, value string }{
		{"site_name", "我的礼物愿望清单"},
		{"intro_text", "这是我的礼物愿望清单，记录着我喜欢和需要的东西。如果你需要灵感，希望它能帮到你！为方便寄送，我的顺丰地址ID是：SFADD999999"},
	}

	for _, d := range defaults {
		_, err := DB.Exec(
			`INSERT OR IGNORE INTO config (key, value) VALUES (?, ?)`,
			d.key, d.value,
		)
		if err != nil {
			log.Printf("Failed to insert default config %s: %v", d.key, err)
		}
	}
}
