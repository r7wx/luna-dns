/*
MIT License
Copyright (c) 2022 r7wx
Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:
The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.
THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

package cache

import (
	"fmt"
	"sync"
	"time"

	"github.com/r7wx/luna-dns/internal/logger"
)

type cacheEntry struct {
	createdAt time.Time
	value     string
}

// Cache - DNS cache struct
type Cache struct {
	sync.Mutex
	domains map[string]cacheEntry
	ttl     time.Duration
}

// NewCache - Create a new cache struct
func NewCache(ttl time.Duration) *Cache {
	return &Cache{
		ttl:     ttl,
		domains: map[string]cacheEntry{},
	}
}

// Routine - Starts the cache cleaning routine
func (c *Cache) Routine() {
	for {
		time.Sleep(1 * time.Minute)
		logger.Info("Cleaning old cache entries...")

		deletedDomains := 0
		for domain, entry := range c.domains {
			delta := time.Now().Sub(entry.createdAt)
			if delta > c.ttl {
				delete(c.domains, domain)
				deletedDomains++
			}
		}

		if deletedDomains > 0 {
			logger.Info("Deleted " + fmt.Sprint(deletedDomains) +
				" domains from cache")
		}
	}
}

// SearchDomain - Search for a domain in Cache
func (c *Cache) SearchDomain(domain string) string {
	c.Lock()
	defer c.Unlock()
	return c.domains[domain].value
}

// InsertDomain - Insert a new domain in Cache
func (c *Cache) InsertDomain(domain, ip string) {
	c.Lock()
	defer c.Unlock()
	c.domains[domain] = cacheEntry{
		createdAt: time.Now(),
		value:     ip,
	}
}
