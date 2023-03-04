package engine

import (
	"log"
	"time"

	"github.com/miekg/dns"
	"github.com/r7wx/luna-dns/internal/cache"
	"github.com/r7wx/luna-dns/internal/config"
	"github.com/r7wx/luna-dns/internal/entry"
	"github.com/r7wx/luna-dns/internal/tree"
)

// Engine - DNS Engine
type Engine struct {
	Hosts        *tree.Tree
	cache        *cache.Cache
	addr         string
	network      string
	dns          []config.DNS
	forwardIndex int
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
		addr:         config.Addr,
		network:      config.Network,
		dns:          config.DNS,
		forwardIndex: 0,
	}, nil
}

// Start - Start Engine DNS server
func (e *Engine) Start() error {
	go e.cache.Routine()

	log.Printf("Listening on %s (%s)\n", e.addr,
		e.network)

	dns.HandleFunc(".", e.handler)
	server := &dns.Server{Addr: e.addr, Net: e.network}
	err := server.ListenAndServe()
	if err != nil {
		return err
	}
	defer server.Shutdown()

	return nil
}
