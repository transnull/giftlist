package handlers

import (
	"encoding/json"
	"net/http"

	"giftlist/models"
)

type configUpdateRequest struct {
	SiteName  *string `json:"site_name"`
	IntroText *string `json:"intro_text"`
}

// GetConfig handles GET /api/config
func GetConfig(w http.ResponseWriter, r *http.Request) {
	config, err := models.GetAllConfig()
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(config)
}

// UpdateConfig handles PUT /api/config
func UpdateConfig(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, `{"error":"Method not allowed"}`, http.StatusMethodNotAllowed)
		return
	}

	var req configUpdateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request body"})
		return
	}

	if req.SiteName == nil && req.IntroText == nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "At least one field required"})
		return
	}

	if req.SiteName != nil {
		if err := models.SetConfig("site_name", *req.SiteName); err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
			return
		}
	}

	if req.IntroText != nil {
		if err := models.SetConfig("intro_text", *req.IntroText); err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
			return
		}
	}

	config, _ := models.GetAllConfig()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(config)
}
