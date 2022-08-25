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
