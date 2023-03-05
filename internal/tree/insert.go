package tree

import (
	"github.com/r7wx/luna-dns/internal/entry"
)

// Insert - Insert new entry in DNS tree
func (t *Tree) Insert(entry *entry.Entry) {
	foundTLD, _ := searchNode(&t.tlds, entry.TLD)
	if foundTLD == nil {
		foundTLD = t.insertNode(&t.tlds, entry.TLD, entry.IP)
	}

	current := foundTLD
	for _, subdomain := range entry.Subdomains {
		foundNode, _ := searchNode(&current.childrens, subdomain)
		if foundNode == nil {
			foundNode = t.insertNode(&current.childrens, subdomain,
				entry.IP)
		}
		current = foundNode
	}
}

func (t *Tree) insertNode(nodes *map[string]*node, host string, ip string) *node {
	if host == "*" {
		for k := range *nodes {
			delete(*nodes, k)
		}
		(*nodes)["*"] = &node{
			ip:        ip,
			childrens: map[string]*node{},
		}

		insertedNode, _ := searchNode(nodes, host)
		return insertedNode
	}

	(*nodes)[host] = &node{
		ip:        ip,
		childrens: map[string]*node{},
	}

	insertedNode, _ := searchNode(nodes, host)
	return insertedNode
}
