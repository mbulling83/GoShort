package utils

import (
	"math/rand"
	"time"
)

const shortURLEncoding = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// GenerateShortURL generates a random short URL path
func GenerateShortURL() string {
	length := 8 // Default length of the short URL
	random := rand.New(rand.NewSource(time.Now().UnixNano())) // Create a local random generator
	result := make([]byte, length)
	for i := range result {
		result[i] = shortURLEncoding[random.Intn(len(shortURLEncoding))]
	}
	return string(result)
}
