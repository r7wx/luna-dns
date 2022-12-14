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

package cache

import (
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/miekg/dns"
)

func TestCache(t *testing.T) {
	cache := NewCache(1 * time.Millisecond)

	rr, _ := dns.NewRR(fmt.Sprintf("%s A %s", "test", "127.0.0.1"))
	question := dns.Question{
		Name:   "test",
		Qtype:  1,
		Qclass: 1,
	}
	cache.Insert([]dns.Question{question}, []dns.RR{rr})
	found := cache.Search([]dns.Question{question})
	if found == nil {
		t.Fatal()
	}
	if !reflect.DeepEqual(found[0], rr) {
		t.Fatal()
	}

	time.Sleep(1 * time.Millisecond)
	deleted := cache.deleteOldEntries()
	if deleted == 0 {
		t.Fatal()
	}
}
