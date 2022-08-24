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
	"github.com/r7wx/luna-dns/internal/dtree"
	"github.com/r7wx/luna-dns/internal/entry"
	"github.com/r7wx/luna-dns/internal/logger"
)

// Tree - Tree struct
type Tree struct {
	domainTree *dtree.DTree
}

// NewTree - Create a new tree
func NewTree(entries []entry.Entry) *Tree {
	tree := &Tree{
		domainTree: dtree.NewDTree(),
	}

	for _, entry := range entries {
		tree.Insert(&entry)
	}

	return tree
}

// Insert - Insert an entry in Tree
func (t *Tree) Insert(entry *entry.Entry) {
	t.domainTree.Insert(entry)
}

// SearchDomain - Search for a domain in Tree
func (t *Tree) SearchDomain(domain string) string {
	entry, err := entry.NewEntry(domain, "")
	if err != nil {
		logger.Error(err)
		return ""
	}
	return t.domainTree.Search(entry)
}
