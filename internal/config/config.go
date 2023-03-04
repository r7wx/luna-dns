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
