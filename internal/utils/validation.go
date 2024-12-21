package utils

import (
	"net/url"
	"regexp"
)

// ValidateURL checks if a given string is a valid URL
func ValidateURL(u string) bool {
	parsedURL, err := url.ParseRequestURI(u)
	if err != nil || parsedURL.Scheme == "" || parsedURL.Host == "" {
		return false
	}
	return true
}

// ValidateCustomShortURL checks if the custom short URL path is alphanumeric
func ValidateCustomShortURL(shortURL string) bool {
	re := regexp.MustCompile("^[a-zA-Z0-9_-]+$")
	return re.MatchString(shortURL)
}
