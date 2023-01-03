package bot

import (
	"gopkg.in/yaml.v3"
)

// Config contains the configuration about the bot module
type Config struct {
	TokenId  string `yaml:"tokenId"`
	Tcp_port string `yaml:"tcp-port"`
}

func NewConfig(tokenId string, tcp_port string) *Config {
	return &Config{
		TokenId:  tokenId,
		Tcp_port: tcp_port,
	}
}

func DefaultConfig() *Config {
	return &Config{
		TokenId:  "discord-bot-token-id",
		Tcp_port: "tcp-port",
	}
}

func ParseConfig(bz []byte) (*Config, error) {
	type T struct {
		Config *Config `yaml:"bot"`
	}
	var cfg T
	err := yaml.Unmarshal(bz, &cfg)

	if cfg.Config == nil {
		return DefaultConfig(), nil
	}

	return cfg.Config, err
}
