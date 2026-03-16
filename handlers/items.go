package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"giftlist/models"
)

type itemRequest struct {
	Name          string   `json:"name"`
	Category      *string  `json:"category"`
	Priority      string   `json:"priority"`
	EstimatedCost *float64 `json:"estimated_cost"`
	ImageURL      string   `json:"image_url"`
	ProductLink   *string  `json:"product_link"`
	Remark        *string  `json:"remark"`
}

func writeJSON(w http.ResponseWriter, status int, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}

// GetItems handles GET /api/items
func GetItems(w http.ResponseWriter, r *http.Request) {
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
	if limit == 0 {
		limit = 10
	}

	filter := models.ItemFilter{
		Category: r.URL.Query().Get("category"),
		Priority: r.URL.Query().Get("priority"),
		Sort:     r.URL.Query().Get("sort"),
		Order:    r.URL.Query().Get("order"),
		Page:     page,
		Limit:    limit,
	}

	result, err := models.GetItems(filter)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	writeJSON(w, http.StatusOK, result)
}

// GetItem handles GET /api/items/:id
func GetItem(w http.ResponseWriter, r *http.Request) {
	id, err := extractID(r.URL.Path, "/api/items/")
	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid ID"})
		return
	}

	item, err := models.GetItemByID(id)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	if item == nil {
		writeJSON(w, http.StatusNotFound, map[string]string{"error": "Item not found"})
		return
	}

	writeJSON(w, http.StatusOK, item)
}

// CreateItem handles POST /api/items
func CreateItem(w http.ResponseWriter, r *http.Request) {
	var req itemRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
		return
	}

	if req.Name == "" || req.ImageURL == "" {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "name and image_url are required"})
		return
	}
	if req.Priority != "high" && req.Priority != "medium" && req.Priority != "low" {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "priority must be high, medium, or low"})
		return
	}

	item := &models.Item{
		Name:          req.Name,
		Category:      req.Category,
		Priority:      req.Priority,
		EstimatedCost: req.EstimatedCost,
		ImageURL:      req.ImageURL,
		ProductLink:   req.ProductLink,
		Remark:        req.Remark,
	}

	created, err := models.CreateItem(item)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	writeJSON(w, http.StatusCreated, created)
}

// UpdateItem handles PUT /api/items/:id
func UpdateItem(w http.ResponseWriter, r *http.Request) {
	id, err := extractID(r.URL.Path, "/api/items/")
	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid ID"})
		return
	}

	var req itemRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
		return
	}

	if req.Name == "" || req.ImageURL == "" {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "name and image_url are required"})
		return
	}
	if req.Priority != "high" && req.Priority != "medium" && req.Priority != "low" {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "priority must be high, medium, or low"})
		return
	}

	existing, _ := models.GetItemByID(id)
	if existing == nil {
		writeJSON(w, http.StatusNotFound, map[string]string{"error": "Item not found"})
		return
	}

	item := &models.Item{
		Name:          req.Name,
		Category:      req.Category,
		Priority:      req.Priority,
		EstimatedCost: req.EstimatedCost,
		ImageURL:      req.ImageURL,
		ProductLink:   req.ProductLink,
		Remark:        req.Remark,
	}

	updated, err := models.UpdateItem(id, item)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	writeJSON(w, http.StatusOK, updated)
}

// DeleteItem handles DELETE /api/items/:id
func DeleteItem(w http.ResponseWriter, r *http.Request) {
	id, err := extractID(r.URL.Path, "/api/items/")
	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid ID"})
		return
	}

	existing, _ := models.GetItemByID(id)
	if existing == nil {
		writeJSON(w, http.StatusNotFound, map[string]string{"error": "Item not found"})
		return
	}

	if err := models.DeleteItem(id); err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	writeJSON(w, http.StatusOK, map[string]string{"message": "Deleted"})
}

func extractID(path, prefix string) (int64, error) {
	idStr := strings.TrimPrefix(path, prefix)
	idStr = strings.TrimSuffix(idStr, "/")
	return strconv.ParseInt(idStr, 10, 64)
}
