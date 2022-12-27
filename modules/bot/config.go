package bot

import (
	"gopkg.in/yaml.v3"
)

// Config contains the configuration about the bot module
type Config struct {
	TokenId string `yaml:"tokenId"`
}

func NewConfig(tokenId string) *Config {
	return &Config{
		TokenId: tokenId,
	}
}

func DefaultConfig() *Config {
	return &Config{
		TokenId: "discord-bot-token-id",
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
