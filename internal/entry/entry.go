package entry

import (
	"fmt"
	"strings"
)

// Entry - Entry struct
type Entry struct {
	Host       string
	IP         string
	TLD        string
	Subdomains []string
}

// NewEntry - Create a new entry
func NewEntry(host, ip string) (*Entry, error) {
	elements := strings.Split(host, ".")
	if len(elements) == 1 && elements[0] != "*" {
		return nil,
			fmt.Errorf("invalid host: %s", host)
	}

	subdomains := []string{}
	for i := len(elements) - 2; i >= 0; i-- {
		subdomains = append(subdomains, elements[i])
	}

	return &Entry{
		Host:       host,
		IP:         ip,
		TLD:        elements[len(elements)-1],
		Subdomains: subdomains,
	}, nil
}
