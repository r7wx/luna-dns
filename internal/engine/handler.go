package engine

import (
	"strings"

	"github.com/miekg/dns"
	"github.com/r7wx/luna-dns/internal/logger"
)

func (e *Engine) handler(w dns.ResponseWriter, r *dns.Msg) {
	message := dns.Msg{}
	message.SetReply(r)

	logger.Debug(strings.ReplaceAll(message.String(),
		"\n", ""))

	switch r.Opcode {
	case dns.OpcodeQuery:
		e.query(&message)
	default:
		e.forward(&message)
	}

	w.WriteMsg(&message)
}
