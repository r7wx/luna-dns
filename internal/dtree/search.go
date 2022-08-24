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

package dtree

import (
	"github.com/r7wx/luna-dns/internal/entry"
)

// Search - Search for an entry in domain tree
func (t *DTree) Search(entry *entry.Entry) string {
	foundTLD, wildcard := searchNode(&t.tlds, entry.TLD)
	if wildcard {
		return foundTLD.ip
	}

	if foundTLD != nil {
		current := foundTLD
		for index, subdomain := range entry.Subdomains {
			foundNode, wildcard := searchNode(&current.childrens, subdomain)
			switch {
			case wildcard:
				return foundNode.ip
			case foundNode == nil:
				return ""
			case index == len(entry.Subdomains)-1:
				return foundNode.ip
			}
			current = foundNode
		}
	}

	return ""
}

func searchNode(nodes *map[string]*node, key string) (*node, bool) {
	wildcardNode, ok := (*nodes)["*"]
	if ok {
		return wildcardNode, true
	}

	node, ok := (*nodes)[key]
	if !ok {
		return nil, false
	}

	return node, false
}
