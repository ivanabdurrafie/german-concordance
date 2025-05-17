package config

import (
	"github.com/spf13/viper"
)

type ServerConfig struct {
	Port string `mapstructure:"PORT"`
}

type Config struct {
	Server ServerConfig `mapstructure:"server"`
}

func Load(configPath string) (*Config, error) {
	viper.SetConfigFile(configPath)
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	if cfg.Server.Port == "" {
		cfg.Server.Port = "8080"
	}

	return &cfg, nil
}
