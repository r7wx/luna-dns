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
	"fmt"
	"net"
	"reflect"
	"testing"

	"github.com/miekg/dns"
	"github.com/r7wx/luna-dns/internal/config"
)

type testResponseWrtiter struct {
	outMessage *dns.Msg
}

func (w *testResponseWrtiter) LocalAddr() net.Addr {
	return &net.UDPAddr{}
}

func (w *testResponseWrtiter) RemoteAddr() net.Addr {
	return &net.UDPAddr{}
}

func (w *testResponseWrtiter) WriteMsg(m *dns.Msg) error {
	w.outMessage = m
	return nil
}

func (w *testResponseWrtiter) Write([]byte) (int, error) {
	return 0, nil
}

func (w *testResponseWrtiter) Close() error {
	return nil
}

func (w *testResponseWrtiter) TsigStatus() error {
	return nil
}

func (w *testResponseWrtiter) TsigTimersOnly(bool) {}

func (w *testResponseWrtiter) Hijack() {}

func TestNewEngine(t *testing.T) {
	_, err := NewEngine(&config.Config{
		Hosts: []config.Host{
			{
				Host: "google.com",
				IP:   "127.0.0.1",
			},
		},
	})
	if err != nil {
		t.Fatal()
	}

	_, err = NewEngine(&config.Config{
		Hosts: []config.Host{
			{
				Host: "x",
			},
		},
	})
	if err == nil {
		t.Fatal()
	}
}

func TestHandler(t *testing.T) {
	engine, _ := NewEngine(&config.Config{
		Addr:    "127.0.0.1:53555",
		Network: "udp",
		Hosts: []config.Host{
			{
				Host: "google.com",
				IP:   "127.0.0.1",
			},
		},
		Debug: true,
	})

	testW := testResponseWrtiter{}
	engine.handler(&testW, &dns.Msg{
		MsgHdr: dns.MsgHdr{
			Opcode: dns.OpcodeQuery,
		},
		Question: []dns.Question{{
			Name:  "google.com.",
			Qtype: dns.TypeA,
		}},
	})

	response := testW.outMessage.Answer[0]
	expected, _ := dns.NewRR(fmt.Sprintf("%s A %s",
		"google.com", "127.0.0.1"))
	if response.String() != expected.String() {
		t.Fail()
	}
}

func TestFormatMessage(t *testing.T) {
	originalHeader := dns.MsgHdr{
		Id:       100,
		Response: false,
		Opcode:   500,
	}
	original := dns.Msg{
		MsgHdr: originalHeader,
	}

	out := formatMessage(&original)
	if reflect.DeepEqual(originalHeader, out.MsgHdr) {
		t.Fail()
	}
}
