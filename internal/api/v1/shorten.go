package v1

import (
	"GoShort/internal/db"
	"GoShort/internal/models"
	"GoShort/internal/utils"
	"encoding/json"
	"net/http"
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

	// Save the URL to the database
	url := models.URL{
		LongURL:  req.LongURL,
		ShortURL: shortURL,
		Expiry:   req.Expiry,
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
