package analytics

import (
	"log"
	"time"
)

// StartCleanupTask starts a periodic cleanup task to remove old analytics data
func StartCleanupTask(tracker *AnalyticsTracker, interval time.Duration, cutoff time.Duration) {
	go func() {
		for {
			time.Sleep(interval)

			log.Println("Running analytics cleanup task...")
			tracker.CleanupOldStats(time.Now().Add(-cutoff))
			log.Println("Analytics cleanup task completed.")
		}
	}()
}
