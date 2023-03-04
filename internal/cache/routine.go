package cache

import (
	"fmt"
	"time"

	"github.com/r7wx/luna-dns/internal/logger"
)

// Routine - Starts the cache cleaning routine
func (c *Cache) Routine() {
	for {
		time.Sleep(c.ttl + (1 * time.Second))
		logger.Info("Cleaning old cache entries...")

		deletedEntries := c.deleteOldEntries()
		if deletedEntries > 0 {
			logger.Info("Deleted " + fmt.Sprint(deletedEntries) +
				" entries from cache")
		}
	}
}

func (c *Cache) deleteOldEntries() int {
	deletedEntries := 0

	for hash, entry := range c.entries {
		delta := time.Now().Sub(entry.createdAt)
		if delta > c.ttl {
			delete(c.entries, hash)
			deletedEntries++
		}
	}

	return deletedEntries
}
