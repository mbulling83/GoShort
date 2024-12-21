package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"time"
)

// URLScanResponse represents the response from a URL scanning API
type URLScanResponse struct {
	Safe bool `json:"safe"`
}

// CheckMaliciousURL uses an external API to check if a URL is malicious
func CheckMaliciousURL(apiEndpoint, apiKey, urlToScan string) (bool, error) {
	client := &http.Client{Timeout: 10 * time.Second}

	// Prepare JSON payload
	payload := map[string]string{"url": urlToScan}
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return false, err
	}

	// Create HTTP request
	req, err := http.NewRequest("POST", apiEndpoint, bytes.NewBuffer(payloadBytes))
	if err != nil {
		return false, err
	}
	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content-Type", "application/json")

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return false, errors.New("failed to scan URL: non-200 response")
	}

	// Parse the response
	var result URLScanResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	return result.Safe, nil
}
