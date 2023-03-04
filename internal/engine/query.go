package engine

import (
	"fmt"

	"github.com/miekg/dns"
)

func (e *Engine) query(message *dns.Msg) {
	for _, q := range message.Question {
		switch q.Qtype {
		case dns.TypeA:
			ip, err := e.Hosts.Search(q.Name[:len(q.Name)-1])
			if ip == "" || err != nil {
				e.forward(message)
				return
			}

			rr, err := dns.NewRR(fmt.Sprintf("%s A %s", q.Name, ip))
			if err == nil {
				message.Answer = append(message.Answer, rr)
			}
		default:
			e.forward(message)
		}
	}
}
