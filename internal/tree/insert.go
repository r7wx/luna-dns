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
