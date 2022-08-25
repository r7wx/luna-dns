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
	"time"

	"github.com/miekg/dns"
	"github.com/r7wx/luna-dns/internal/config"
)

func TestNewEngine(t *testing.T) {
	_, err := NewEngine(&config.Config{
		Hosts: []config.Host{
			{
				Domain: "google.com",
				IP:     "127.0.0.1",
			},
		},
	})
	if err != nil {
		t.Fatal()
	}

	_, err = NewEngine(&config.Config{
		Hosts: []config.Host{
			{
				Domain: "x",
			},
		},
	})
	if err == nil {
		t.Fatal()
	}
}

func TestQuery(t *testing.T) {
	dns.HandleFunc(".", func(w dns.ResponseWriter, r *dns.Msg) {
		message := dns.Msg{}
		message.SetReply(r)
		rr, _ := dns.NewRR("test.test. A 127.0.0.1")
		message.Answer = append(message.Answer, rr)
		w.WriteMsg(&message)
	})

	testServer := &dns.Server{Addr: "127.0.0.1:55553", Net: "udp"}
	go testServer.ListenAndServe()
	defer testServer.Shutdown()

	engine, _ := NewEngine(&config.Config{
		DNS: []config.DNS{
			{
				Addr:    "xxxxxxxx",
				Network: "udp",
			},
			{
				Addr:    "127.0.0.1:55553",
				Network: "udp",
			},
		},
		Hosts: []config.Host{
			{
				Domain: "google.com",
				IP:     "127.0.0.1",
			},
		},
	})

	testMessage := new(dns.Msg)
	testMessage.SetQuestion("google.com.", dns.TypeA)
	engine.query(testMessage)
	if len(testMessage.Answer) == 0 {
		t.Fatal()
	}

	testMessage = new(dns.Msg)
	testMessage.SetQuestion("go.dev.", dns.TypeA)
	engine.query(testMessage)
	if len(testMessage.Answer) == 0 {
		t.Fatal()
	}
	time.Sleep(1 * time.Second)
	testMessage.Answer = []dns.RR{}
	engine.query(testMessage)
	if len(testMessage.Answer) == 0 {
		t.Fatal()
	}

	testMessage = new(dns.Msg)
	testMessage.SetQuestion("go.dev.", dns.TypeTXT)
	engine.query(testMessage)
	if len(testMessage.Answer) == 0 {
		t.Fatal()
	}
}
