package cache

import (
	"sync"
	"time"

	"github.com/miekg/dns"
	"github.com/r7wx/luna-dns/internal/logger"
)

// Cache - DNS cache struct
type Cache struct {
	sync.Mutex
	entries map[string]entry
	ttl     time.Duration
}
type entry struct {
	createdAt time.Time
	answer    []dns.RR
}

// NewCache - Create a new cache struct
func NewCache(ttl time.Duration) *Cache {
	return &Cache{
		ttl:     ttl,
		entries: map[string]entry{},
	}
}

// Search - Search for DNS answer in cache
func (c *Cache) Search(question []dns.Question) []dns.RR {
	c.Lock()
	defer c.Unlock()
	return c.entries[hashQuestion(question)].answer
}

// Insert - Insert a new entry in Cache
func (c *Cache) Insert(question []dns.Question, answer []dns.RR) {
	c.Lock()
	defer c.Unlock()

	hash := hashQuestion(question)
	c.entries[hash] = entry{
		createdAt: time.Now(),
		answer:    answer,
	}
	logger.Debug("New entry in cache: " +
		hash)
}
