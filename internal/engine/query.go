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

	"github.com/miekg/dns"
	"github.com/r7wx/luna-dns/internal/logger"
)

func (e *Engine) query(m *dns.Msg) {
	for _, q := range m.Question {
		switch q.Qtype {
		case dns.TypeA:
			ip := e.queryA(q.Name[:len(q.Name)-1])
			if ip != "" {
				rr, err := dns.NewRR(fmt.Sprintf("%s A %s", q.Name, ip))
				if err == nil {
					m.Answer = append(m.Answer, rr)
				}
			}

		default:
			fmt.Println("TODO: Redirect query")
		}
	}
}

func (e *Engine) queryA(domain string) string {
	logger.Debug("Searching for domain: " + domain)

	ip := e.overrides.SearchDomain(domain)
	if ip != "" {
		logger.Debug("Override match for " + domain +
			": " + ip)
		return ip
	}
	logger.Debug("No override match for: " + domain)

	ip = e.cache.SearchDomain(domain)
	if ip != "" {
		logger.Debug("Cache match for " + domain +
			": " + ip)
		return ip
	}
	logger.Debug("No cache match for: " + domain)

	// TODO: Fallback request to DNS servers
	// TODO: Add to cache (goroutine)

	return ""
}
