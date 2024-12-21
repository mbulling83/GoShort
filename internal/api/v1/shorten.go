package v1

import (
	"GoShort/internal/db"
	"GoShort/internal/models"
	"GoShort/internal/utils"
	"encoding/json"
	"net/http"
	"time"
)

// ShortenRequest represents the request payload for URL shortening
type ShortenRequest struct {
	LongURL   string `json:"long_url"`
	CustomURL string `json:"custom_url,omitempty"`
	Expiry    string `json:"expiry,omitempty"` // Optional expiry date
}

// ShortenResponse represents the response payload for URL shortening
type ShortenResponse struct {
	ShortURL string `json:"short_url"`
}

// ShortenURL handles the URL shortening request
func ShortenURL(w http.ResponseWriter, r *http.Request) {
	var req ShortenRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Validate the long URL
	if !utils.ValidateURL(req.LongURL) {
		http.Error(w, "Invalid URL format", http.StatusBadRequest)
		return
	}

	// Generate a shortened URL
	shortURL := req.CustomURL
	if shortURL == "" {
		shortURL = utils.GenerateShortURL()
	}

	// Parse expiry if provided
	var expiry *time.Time
	if req.Expiry != "" {
		parsedExpiry, err := time.Parse(time.RFC3339, req.Expiry)
		if err != nil {
			http.Error(w, "Invalid expiry format", http.StatusBadRequest)
			return
		}
		expiry = &parsedExpiry
	}

	// Save the URL to the database
	url := models.URL{
		LongURL:  req.LongURL,
		ShortURL: shortURL,
		Expiry:   expiry,
	}
	if err := db.DB.Create(&url).Error; err != nil {
		http.Error(w, "Failed to save URL", http.StatusInternalServerError)
		return
	}

	// Return the shortened URL
	resp := ShortenResponse{
		ShortURL: shortURL,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}