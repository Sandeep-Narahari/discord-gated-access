package config

import (
	"github.com/AutonomyNetwork/iam/modules/bot"
	initcmd "github.com/forbole/juno/v3/cmd/init"
	junoconfig "github.com/forbole/juno/v3/types/config"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"

	"github.com/AutonomyNetwork/iam/modules/actions"
)

// Config represents the BDJuno configuration
type Config struct {
	JunoConfig    junoconfig.Config `yaml:"-,inline"`
	ActionsConfig *actions.Config   `yaml:"actions"`
	BotConfig     *bot.Config       `yaml:"bot"`
}

// NewConfig returns a new Config instance
func NewConfig(junoCfg junoconfig.Config, actionsCfg *actions.Config, botCfg *bot.Config) Config {
	return Config{
		JunoConfig:    junoCfg,
		ActionsConfig: actionsCfg,
		BotConfig:     botCfg,
	}
}

// GetBytes implements WritableConfig
func (c Config) GetBytes() ([]byte, error) {
	return yaml.Marshal(&c)
}

// Creator represents a configuration creator
func Creator(_ *cobra.Command) initcmd.WritableConfig {
	return NewConfig(junoconfig.DefaultConfig(), actions.DefaultConfig(), bot.DefaultConfig())
}
