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
	"testing"

	"github.com/r7wx/luna-dns/internal/entry"
)

func TestBasics(t *testing.T) {
	testEntries := func(tree *Tree, entries []map[string]string, t *testing.T) {
		for _, e := range entries {
			testEntry, _ := entry.NewEntry(e["host"], e["ip"])
			tree.Insert(testEntry)

			found, _ := tree.Search(testEntry.Host)
			if found == "" {
				t.Fatal()
				continue
			}

			if found != testEntry.IP {
				t.Fatal()
			}
		}
	}

	tree := NewTree()
	testEntries(tree, []map[string]string{
		{
			"host": "test.com",
			"ip":   "127.0.0.1",
		},
		{
			"host": "a.test.com",
			"ip":   "127.0.0.1",
		},
		{
			"host": "test.a",
			"ip":   "127.0.0.1",
		},
	}, t)
}

func TestOthers(t *testing.T) {
	insertEntries := func(tree *Tree, entries []map[string]string) {
		for _, e := range entries {
			testEntry, _ := entry.NewEntry(e["host"], e["ip"])
			tree.Insert(testEntry)
		}
	}

	searchDomains := func(tree *Tree, entries []map[string]any,
		t *testing.T) {
		for _, e := range entries {
			host := e["host"].(string)
			expected := e["expected"].(bool)

			found, _ := tree.Search(host)
			if found == "" && expected {
				t.Fatal()
			}
			if found != "" && !expected {
				t.Fatal()
			}
		}
	}

	tree := NewTree()
	insertEntries(tree, []map[string]string{
		{
			"host": "*.test.com",
			"ip":   "127.0.0.1",
		},
		{
			"host": "*.tld",
			"ip":   "127.0.0.1",
		},
	})
	searchDomains(tree, []map[string]any{
		{
			"host":     "unk.com",
			"expected": false,
		},
		{
			"host":     "aaa.test.com",
			"expected": true,
		},
		{
			"host":     "test.tld",
			"expected": true,
		},
		{
			"host":     "test.xxx",
			"expected": false,
		},
	}, t)

	insertEntries(tree, []map[string]string{
		{
			"host": "*",
			"ip":   "127.0.0.1",
		},
	})
	searchDomains(tree, []map[string]any{
		{
			"host":     "google.com",
			"expected": true,
		},
	}, t)

	_, err := tree.Search("xxx")
	if err == nil {
		t.Fatal()
	}
}
