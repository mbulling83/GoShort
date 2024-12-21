package analytics

import (
	"sync"
	"time"
)

// AnalyticsTracker is a thread-safe structure to track and store URL usage statistics
type AnalyticsTracker struct {
	mu         sync.RWMutex
	statistics map[string]*URLStats
}

// URLStats holds analytics data for a specific URL
type URLStats struct {
	ClickCount   int
	LastAccessed time.Time
	Referrers    map[string]int
	Geolocations map[string]int
}

// NewAnalyticsTracker creates and initializes a new AnalyticsTracker
func NewAnalyticsTracker() *AnalyticsTracker {
	return &AnalyticsTracker{
		statistics: make(map[string]*URLStats),
	}
}

// RecordClick increments the click count for the given shortURL and tracks metadata
func (at *AnalyticsTracker) RecordClick(shortURL, referrer, geolocation string) {
	at.mu.Lock()
	defer at.mu.Unlock()

	if _, exists := at.statistics[shortURL]; !exists {
		at.statistics[shortURL] = &URLStats{
			Referrers:    make(map[string]int),
			Geolocations: make(map[string]int),
		}
	}

	stats := at.statistics[shortURL]
	stats.ClickCount++
	stats.LastAccessed = time.Now()
	stats.Referrers[referrer]++
	stats.Geolocations[geolocation]++
}

// GetStats retrieves statistics for the given shortURL
func (at *AnalyticsTracker) GetStats(shortURL string) (*URLStats, bool) {
	at.mu.RLock()
	defer at.mu.RUnlock()

	stats, exists := at.statistics[shortURL]
	return stats, exists
}

// CleanupOldStats removes analytics data for URLs not accessed since the provided cutoff time
func (at *AnalyticsTracker) CleanupOldStats(cutoff time.Time) {
	at.mu.Lock()
	defer at.mu.Unlock()

	for shortURL, stats := range at.statistics {
		if stats.LastAccessed.Before(cutoff) {
			delete(at.statistics, shortURL)
		}
	}
}
