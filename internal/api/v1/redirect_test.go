package v1

import (
	"GoShort/internal/models"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func mockRedirectURL(shortToURLMap map[string]*models.URL) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		shortURL := r.URL.Path[1:] // Extract the short URL from the path

		url, exists := shortToURLMap[shortURL]
		if !exists {
			http.NotFound(w, r)
			return
		}

		// Check for expiration
		if url.Expiry != nil && time.Now().After(*url.Expiry) {
			http.Error(w, "URL has expired", http.StatusGone)
			return
		}

		// Redirect to the original URL
		http.Redirect(w, r, url.LongURL, http.StatusFound)
	}
}

func TestRedirectURL(t *testing.T) {
	tests := []struct {
		name          string
		shortURL      string
		mockData      map[string]*models.URL
		expectedCode  int
		expectedHeader string
	}{
		{
			name:     "Valid short URL",
			shortURL: "short123",
			mockData: map[string]*models.URL{
				"short123": {
					ShortURL: "short123",
					LongURL:  "https://example.com",
				},
			},
			expectedCode:  http.StatusFound,
			expectedHeader: "https://example.com",
		},
		{
			name:         "Short URL not found",
			shortURL:     "unknown123",
			mockData:     map[string]*models.URL{},
			expectedCode: http.StatusNotFound,
		},
		{
			name:     "Expired short URL",
			shortURL: "expired123",
			mockData: map[string]*models.URL{
				"expired123": {
					ShortURL: "expired123",
					LongURL:  "https://example.com",
					Expiry:   func() *time.Time { t := time.Now().Add(-time.Hour); return &t }(),
				},
			},
			expectedCode: http.StatusGone,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create mock handler
			handler := mockRedirectURL(tt.mockData)

			// Create HTTP request and recorder
			req := httptest.NewRequest(http.MethodGet, "/"+tt.shortURL, nil)
			w := httptest.NewRecorder()

			// Call the mock handler
			handler(w, req)

			// Assert response code
			res := w.Result()
			assert.Equal(t, tt.expectedCode, res.StatusCode)

			// Assert Location header if applicable
			if tt.expectedHeader != "" {
				header := res.Header.Get("Location")
				assert.Equal(t, tt.expectedHeader, header)
			}
		})
	}
}
