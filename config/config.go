package config

import (
	"github.com/caarlos0/env/v6"
	"github.com/pkg/errors"
)

type (
	Config struct {
		PostgresConfig PostgresConfig `envPrefix:"PG_"`
		HttpConfig     HttpConfig     `envPrefix:"HTTP_"`
	}

	PostgresConfig struct {
		Opts string `env:"OPTS,required,notEmpty" env-default:""`
	}

	HttpConfig struct {
		Port int `env:"PORT" env-default:"1212"`
	}
)

func GetConfig() (*Config, error) {
	cfg := Config{}

	if err := env.Parse(&cfg); err != nil {
		return nil, errors.Wrap(err, "Can`t read config")
	}
	return &cfg, nil
}
