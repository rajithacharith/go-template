package config

import (
	"github.com/caarlos0/env/v11"
	"github.com/sirupsen/logrus"
)

var config *Config

func LoadConfigs() error {
	config = &Config{}
	if err := env.Parse(config); err != nil {
		return err
	}
	logrus.Info("config loaded successfully")
	return nil
}

func GetConfig() *Config {
	return config
}
