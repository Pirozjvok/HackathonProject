package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/pkg/errors"
)

type (
	Config struct {
		App  `yaml:"app"`
		HTTP `yaml:"http"`
		API  `yaml:"api"`
		PG   `yaml:"pg"`
	}

	App struct {
		Name string `yaml:"name" env:"APP_NAME"`
	}
	HTTP struct {
		Port string `yaml:"port" env:"HTTP_PORT"`
	}

	API struct {
		UrlAPI     string `yaml:"url_api" env:"API_URL"`
		ServiceSHD string `yaml:"service_shd" env:"API_SERVICE_SHD"`
	}

	PG struct {
		ConnURI string `yaml:"conn_uri" env:"PG_URI"`
	}
)

func NewConfig() (*Config, error) {
	cfg := &Config{}

	err := cleanenv.ReadConfig("./config/config.yaml", cfg)
	if err != nil {
		return nil, errors.Wrap(err, "NewConfig: fail to read config.yaml file")
	}

	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, errors.Wrap(err, "NewConfig: fail to read env")
	}
	return cfg, err
}
