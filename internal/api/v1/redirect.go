package v1

import (
	"GoShort/internal/db"
	"GoShort/internal/models"
	"net/http"
	"time"
)

// RedirectURL handles redirecting a short URL to its original URL
func RedirectURL(w http.ResponseWriter, r *http.Request) {
	shortURL := r.URL.Path[1:] // Extract short URL from path

	var url models.URL
	if err := db.DB.Where("short_url = ?", shortURL).First(&url).Error; err != nil {
		http.NotFound(w, r)
		return
	}

	// Check for expiration
	if url.Expiry != "" {
		expiryTime, err := time.Parse(time.RFC3339, url.Expiry)
		if err == nil && time.Now().After(expiryTime) {
			http.Error(w, "URL has expired", http.StatusGone)
			return
		}
	}

	// Redirect to the original URL
	http.Redirect(w, r, url.LongURL, http.StatusFound)
}
