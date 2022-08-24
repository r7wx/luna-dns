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
	"crypto/rand"
	"encoding/binary"
	"fmt"

	"github.com/miekg/dns"
	"github.com/r7wx/luna-dns/internal/logger"
)

func (e *Engine) remoteFallback(message *dns.Msg) {
	cachedAnswer := e.cache.Search(message.Question)
	if cachedAnswer != nil {
		logger.Debug(fmt.Sprintf("Entry found in cache: %v",
			cachedAnswer))
		message.Answer = cachedAnswer
		return
	}

	for _, server := range e.dns {
		conn, err := dns.Dial(server.Protocol, server.Addr)
		if err != nil {
			logger.Debug(fmt.Sprintf("%s (%s): %s",
				server.Addr, server.Protocol, err))
			continue
		}
		defer conn.Close()

		request := formatMessage(message)
		err = conn.WriteMsg(request)
		if err != nil {
			logger.Debug(fmt.Sprintf("%s (%s): %s",
				server.Addr, server.Protocol, err))
			continue
		}

		response, err := conn.ReadMsg()
		if err != nil {
			logger.Debug(fmt.Sprintf("%s (%s): %s",
				server.Addr, server.Protocol, err))
			continue
		}
		if response == nil || response.Rcode != dns.RcodeSuccess {
			logger.Debug(fmt.Sprintf("%s (%s): %s",
				server.Addr, server.Protocol,
				"failed to get a valid response"))
			continue
		}

		if len(response.Answer) > 0 {
			message.Answer = response.Answer
			logger.Debug(fmt.Sprintf("%s (%s) -> %s", server.Addr,
				server.Protocol, response.Answer))
			go e.cache.Insert(message.Question, response.Answer)
			return
		}
	}
}

func formatMessage(original *dns.Msg) *dns.Msg {
	id := func() uint16 {
		var output uint16
		binary.Read(rand.Reader, binary.BigEndian, &output)
		return output
	}

	message := &dns.Msg{}
	original.CopyTo(message)

	message.MsgHdr = dns.MsgHdr{}
	message.Id = id()
	message.RecursionDesired = true

	return message
}
