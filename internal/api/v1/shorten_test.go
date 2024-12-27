package v1

import (
	"GoShort/internal/utils"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestValidateURL(t *testing.T) {
	validURL := "https://example.com"
	invalidURL := "invalid-url"

	assert.True(t, utils.ValidateURL(validURL), "Expected valid URL to pass validation")
	assert.False(t, utils.ValidateURL(invalidURL), "Expected invalid URL to fail validation")
}

func TestGenerateShortURL(t *testing.T) {
	shortURL := utils.GenerateShortURL()
	assert.NotEmpty(t, shortURL, "Expected generated short URL to be non-empty")
	assert.Len(t, shortURL, 8, "Expected generated short URL to have length 8")
}

func TestParseExpiry(t *testing.T) {
	validExpiry := "2025-01-01T12:00:00Z"
	invalidExpiry := "not-a-date"

	parsedTime, err := time.Parse(time.RFC3339, validExpiry)
	assert.NoError(t, err, "Expected valid expiry to parse without error")
	assert.Equal(t, parsedTime.Format(time.RFC3339), validExpiry, "Expected parsed expiry to match input")

	_, err = time.Parse(time.RFC3339, invalidExpiry)
	assert.Error(t, err, "Expected invalid expiry to return an error")
}
