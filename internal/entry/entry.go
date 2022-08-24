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

package entry

import (
	"fmt"
	"strings"
)

// Entry - Entry struct
type Entry struct {
	Domain     string
	IP         string
	TLD        string
	Subdomains []string
}

// NewEntry - Create a new entry
func NewEntry(domain, ip string) (*Entry, error) {
	elements := strings.Split(domain, ".")
	if len(elements) == 1 && elements[0] != "*" {
		return nil,
			fmt.Errorf("invalid domain: %s", domain)
	}

	subdomains := []string{}
	for i := len(elements) - 2; i >= 0; i-- {
		subdomains = append(subdomains, elements[i])
	}

	return &Entry{
		Domain:     domain,
		IP:         ip,
		TLD:        elements[len(elements)-1],
		Subdomains: subdomains,
	}, nil
}
