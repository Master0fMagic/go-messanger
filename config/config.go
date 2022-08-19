package config

import (
	"github.com/caarlos0/env/v6"
	"github.com/pkg/errors"
)

type (
	Config struct {
		PostgresConfig PostgresConfig `envPrefix:"PG_"`
	}

	PostgresConfig struct {
		Options string `env:"OPTS,required,notEmpty" env-default:""`
	}
)

func GetConfig() (*Config, error) {
	cfg := Config{}

	if err := env.Parse(cfg); err != nil {
		return nil, errors.Wrap(err, "Can`t read config")
	}
	return &cfg, nil
}
