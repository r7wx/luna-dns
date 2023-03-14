package blocklists

import (
	"log"
	"strings"
	"time"

	"github.com/r7wx/luna-dns/internal/tree"
)

const (
	updateTime = 12 * time.Hour
)

// Blocklists - Blocklists strutct
type Blocklists struct {
	hosts      *tree.Tree
	blocklists []string
}

// NewBlocklists - Create a new Blocklists
func NewBlocklists(blocklists []string) *Blocklists {
	return &Blocklists{
		hosts:      tree.NewTree(),
		blocklists: blocklists,
	}
}

// Routine - Start Blocklists update routine
func (b *Blocklists) Routine() {
	if len(b.blocklists) == 0 {
		return
	}

	for {
		log.Println("Updating blocklists...")

		newHosts := tree.NewTree()
		for _, blocklist := range b.blocklists {
			if strings.HasPrefix(blocklist, "file://") {
				b.processFile(blocklist, newHosts)
				continue
			}
			b.processRemote(blocklist, newHosts)
		}
		b.hosts = newHosts
		log.Printf("Blocklists updated, next update in %s\n", updateTime)

		time.Sleep(updateTime)
	}
}
