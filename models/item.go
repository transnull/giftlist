package models

import (
	"database/sql"
	"fmt"
	"strings"
	"time"
)

// Item represents a wishlist item
type Item struct {
	ID            int64      `json:"id"`
	Name          string     `json:"name"`
	Category      *string    `json:"category"`
	Priority      string     `json:"priority"`
	EstimatedCost *float64   `json:"estimated_cost"`
	ImageURL      string     `json:"image_url"`
	ProductLink   *string    `json:"product_link"`
	Remark        *string    `json:"remark"`
	CreatedAt     time.Time  `json:"created_at"`
}

// ItemListResult holds paginated items
type ItemListResult struct {
	Total int    `json:"total"`
	Items []Item `json:"items"`
}

// ItemFilter holds query filter params
type ItemFilter struct {
	Category string
	Priority string
	Sort     string // created_at | estimated_cost
	Order    string // asc | desc
	Page     int
	Limit    int
}

// GetItems returns paginated, filtered, sorted items
func GetItems(f ItemFilter) (*ItemListResult, error) {
	where := []string{"1=1"}
	args := []interface{}{}

	if f.Category != "" {
		where = append(where, "category LIKE ?")
		args = append(args, "%"+f.Category+"%")
	}
	if f.Priority != "" {
		where = append(where, "priority = ?")
		args = append(args, f.Priority)
	}

	whereClause := strings.Join(where, " AND ")

	// Count total
	var total int
	err := DB.QueryRow(fmt.Sprintf(`SELECT COUNT(*) FROM items WHERE %s`, whereClause), args...).Scan(&total)
	if err != nil {
		return nil, err
	}

	// Validate sort column
	sortCol := "created_at"
	if f.Sort == "estimated_cost" {
		sortCol = "estimated_cost"
	}
	sortOrder := "DESC"
	if strings.ToLower(f.Order) == "asc" {
		sortOrder = "ASC"
	}

	if f.Limit <= 0 {
		f.Limit = 10
	}
	if f.Page <= 0 {
		f.Page = 1
	}
	offset := (f.Page - 1) * f.Limit

	query := fmt.Sprintf(
		`SELECT id, name, category, priority, estimated_cost, image_url, product_link, remark, created_at
		 FROM items WHERE %s ORDER BY %s %s LIMIT ? OFFSET ?`,
		whereClause, sortCol, sortOrder,
	)
	args = append(args, f.Limit, offset)

	rows, err := DB.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	items := []Item{}
	for rows.Next() {
		var item Item
		var createdAt string
		err := rows.Scan(
			&item.ID, &item.Name, &item.Category, &item.Priority,
			&item.EstimatedCost, &item.ImageURL, &item.ProductLink,
			&item.Remark, &createdAt,
		)
		if err != nil {
			return nil, err
		}
		item.CreatedAt, _ = time.Parse("2006-01-02 15:04:05", createdAt)
		items = append(items, item)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &ItemListResult{Total: total, Items: items}, nil
}

// GetItemByID returns a single item
func GetItemByID(id int64) (*Item, error) {
	var item Item
	var createdAt string
	err := DB.QueryRow(
		`SELECT id, name, category, priority, estimated_cost, image_url, product_link, remark, created_at
		 FROM items WHERE id = ?`, id,
	).Scan(
		&item.ID, &item.Name, &item.Category, &item.Priority,
		&item.EstimatedCost, &item.ImageURL, &item.ProductLink,
		&item.Remark, &createdAt,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	item.CreatedAt, _ = time.Parse("2006-01-02 15:04:05", createdAt)
	return &item, nil
}

// CreateItem inserts a new item
func CreateItem(item *Item) (*Item, error) {
	result, err := DB.Exec(
		`INSERT INTO items (name, category, priority, estimated_cost, image_url, product_link, remark)
		 VALUES (?, ?, ?, ?, ?, ?, ?)`,
		item.Name, item.Category, item.Priority, item.EstimatedCost,
		item.ImageURL, item.ProductLink, item.Remark,
	)
	if err != nil {
		return nil, err
	}
	id, _ := result.LastInsertId()
	return GetItemByID(id)
}

// UpdateItem updates an existing item
func UpdateItem(id int64, item *Item) (*Item, error) {
	_, err := DB.Exec(
		`UPDATE items SET name=?, category=?, priority=?, estimated_cost=?,
		 image_url=?, product_link=?, remark=? WHERE id=?`,
		item.Name, item.Category, item.Priority, item.EstimatedCost,
		item.ImageURL, item.ProductLink, item.Remark, id,
	)
	if err != nil {
		return nil, err
	}
	return GetItemByID(id)
}

// DeleteItem removes an item by ID
func DeleteItem(id int64) error {
	_, err := DB.Exec(`DELETE FROM items WHERE id = ?`, id)
	return err
}
