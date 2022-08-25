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

	"github.com/r7wx/luna-dns/internal/logger"
	"gopkg.in/yaml.v3"
)

// Host - Host configuration struct
type Host struct {
	Host string `yaml:"host"`
	IP   string `yaml:"ip"`
}

// DNS - DNS server configuration struct
type DNS struct {
	Addr    string `yaml:"addr"`
	Network string `yaml:"network"`
}

// Config - Main configuration struct
type Config struct {
	Addr     string `yaml:"addr"`
	Network  string `yaml:"network"`
	DNS      []DNS  `yaml:"dns"`
	Hosts    []Host `yaml:"hosts"`
	Debug    bool   `yaml:"debug"`
	CacheTTL int64  `yaml:"cache_ttl"`
}

// Load - Load configuration from file
func Load(filepath string) (*Config, error) {
	confBytes, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	config := &Config{}
	err = yaml.Unmarshal(confBytes, config)

	if config.Debug {
		logger.DebugEnabled = true
	}

	return config, err
}
