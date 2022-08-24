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

func TestTree(t *testing.T) {
	testEntry, _ := entry.NewEntry("google.com", "127.0.0.1")
	tree := NewTree([]entry.Entry{*testEntry})

	testEntry, _ = entry.NewEntry("*.google.com", "8.8.8.8")
	tree.Insert(testEntry)

	found := tree.SearchDomain("google.com")
	if found != "127.0.0.1" {
		t.Fatal()
	}

	found = tree.SearchDomain("test.google.com")
	if found != "8.8.8.8" {
		t.Fatal()
	}

	found = tree.SearchDomain("xxxx")
	if found != "" {
		t.Fatal()
	}
}
