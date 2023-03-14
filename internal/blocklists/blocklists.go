package blocklists

import (
	"bufio"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/r7wx/luna-dns/internal/entry"
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
			log.Printf("Downloading %s...\n", blocklist)
			resp, err := http.Get(blocklist)
			if err != nil {
				log.Printf("Error downloading blocklist %s: %s",
					blocklist, err)
				continue
			}
			defer resp.Body.Close()

			scanner := bufio.NewScanner(resp.Body)
			for scanner.Scan() {
				line := scanner.Text()
				line = strings.TrimSpace(line)
				entry, err := entry.NewEntry(line, "0.0.0.0")
				if err != nil {
					continue
				}
				newHosts.Insert(entry)
			}

			if err := scanner.Err(); err != nil {
				log.Printf("Error processing blocklist %s: %s",
					blocklist, err)
				return
			}
		}
		b.hosts = newHosts

		log.Printf("Blocklists updated, next update in %s\n", updateTime)

		time.Sleep(updateTime)
	}
}
