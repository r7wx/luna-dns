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
	"testing"
)

func TestBasics(t *testing.T) {
	testEntries := func(tree *domainTree, entries []map[string]string, t *testing.T) {
		for _, entry := range entries {
			testEntry, _ := newEntry(entry["domain"], entry["ip"])
			tree.insert(testEntry)

			targetIP := testEntry.ip
			testEntry.ip = ""

			found := tree.search(testEntry)
			if found == "" {
				t.Fatal()
				continue
			}

			if found != targetIP {
				t.Fatal()
			}
		}
	}

	tree := &domainTree{}
	testEntries(tree, []map[string]string{
		{
			"domain": "test.com",
			"ip":     "127.0.0.1",
		},
		{
			"domain": "a.test.com",
			"ip":     "127.0.0.1",
		},
		{
			"domain": "test.a",
			"ip":     "127.0.0.1",
		},
	}, t)
}

func TestOthers(t *testing.T) {
	insertEntries := func(tree *domainTree, entries []map[string]string) {
		for _, entry := range entries {
			testEntry, _ := newEntry(entry["domain"], entry["ip"])
			tree.insert(testEntry)
		}
	}

	searchDomains := func(tree *domainTree, entries []map[string]any,
		t *testing.T) {
		for _, entry := range entries {
			domain := entry["domain"].(string)
			expected := entry["expected"].(bool)

			testEntry, _ := newEntry(domain, "")
			found := tree.search(testEntry)

			if found == "" && expected {
				t.Fatal()
			}
			if found != "" && !expected {
				t.Fatal()
			}
		}
	}

	tree := &domainTree{}
	insertEntries(tree, []map[string]string{
		{
			"domain": "*.test.com",
			"ip":     "127.0.0.1",
		},
		{
			"domain": "*.tld",
			"ip":     "127.0.0.1",
		},
	})
	searchDomains(tree, []map[string]any{
		{
			"domain":   "unk.com",
			"expected": false,
		},
		{
			"domain":   "aaa.test.com",
			"expected": true,
		},
		{
			"domain":   "test.tld",
			"expected": true,
		},
		{
			"domain":   "test.xxx",
			"expected": false,
		},
	}, t)

	insertEntries(tree, []map[string]string{
		{
			"domain": "*",
			"ip":     "127.0.0.1",
		},
	})
	searchDomains(tree, []map[string]any{
		{
			"domain":   "google.com",
			"expected": true,
		},
	}, t)
}
