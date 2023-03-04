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

func TestEngine_buildForwardChain(t *testing.T) {
	dns := []config.DNS{
		{Addr: "1.1.1.1:53", Network: "udp"},
		{Addr: "2.2.2.2:53", Network: "udp"},
		{Addr: "3.3.3.3:53", Network: "udp"},
	}
	engine := &Engine{dns: dns}

	engine.forwardIndex = 1
	expectedChain := []config.DNS{
		{Addr: "2.2.2.2:53", Network: "udp"},
		{Addr: "3.3.3.3:53", Network: "udp"},
		{Addr: "1.1.1.1:53", Network: "udp"},
	}
	actualChain := engine.buildForwardChain()
	if !reflect.DeepEqual(actualChain, expectedChain) {
		t.Errorf("Test case 1 failed. Expected %v, but got %v",
			expectedChain, actualChain)
	}

	engine.forwardIndex = len(dns)
	expectedChain = dns
	actualChain = engine.buildForwardChain()
	if !reflect.DeepEqual(actualChain, expectedChain) {
		t.Errorf("Test case 2 failed. Expected %v, but got %v",
			expectedChain, actualChain)
	}

	engine.forwardIndex = len(dns) + 1
	expectedChain = dns
	actualChain = engine.buildForwardChain()
	if !reflect.DeepEqual(actualChain, expectedChain) {
		t.Errorf("Test case 3 failed. Expected %v, but got %v",
			expectedChain, actualChain)
	}
}
