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

package engine

import (
	"time"

	"github.com/miekg/dns"
	"github.com/r7wx/luna-dns/internal/cache"
	"github.com/r7wx/luna-dns/internal/config"
	"github.com/r7wx/luna-dns/internal/entry"
	"github.com/r7wx/luna-dns/internal/logger"
	"github.com/r7wx/luna-dns/internal/tree"
)

// Engine - DNS Engine
type Engine struct {
	Hosts   *tree.Tree
	cache   *cache.Cache
	addr    string
	Network string
	dns     []config.DNS
}

// NewEngine - Create a new engine
func NewEngine(config *config.Config) (*Engine, error) {
	Hosts := tree.NewTree()
	for _, host := range config.Hosts {
		entry, err := entry.NewEntry(host.Host, host.IP)
		if err != nil {
			return nil, err
		}
		Hosts.Insert(entry)
	}

	return &Engine{
		Hosts: Hosts,
		cache: cache.NewCache(time.Duration(config.CacheTTL) *
			time.Second),
		addr:    config.Addr,
		Network: config.Network,
		dns:     config.DNS,
	}, nil
}

// Start - Start Engine DNS server
func (e *Engine) Start() error {
	go e.cache.Routine()

	logger.Info("Engine listening on: " + e.addr + " (" +
		e.Network + ")")

	dns.HandleFunc(".", e.handler)
	server := &dns.Server{Addr: e.addr, Net: e.Network}
	err := server.ListenAndServe()
	if err != nil {
		return err
	}
	defer server.Shutdown()

	return nil
}
