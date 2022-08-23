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

package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

// Protocol enum
const (
	TCP = "tcp"
	UDP = "udp"
)

// Protocol enum type
type Protocol string

// Override - Override configuration struct
type Override struct {
	Domain string `yaml:"domain"`
	IP     string `yaml:"ip"`
}

// Config - Main configuration struct
type Config struct {
	Addr      string     `yaml:"addr"`
	Protocol  Protocol   `yaml:"protocol"`
	DNS       []string   `yaml:"dns"`
	Overrides []Override `yaml:"overrides"`
}

// Load - Load configuration from file
func Load(filepath string) (*Config, error) {
	confBytes, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	config := &Config{}
	err = yaml.Unmarshal(confBytes, config)
	return config, err
}
