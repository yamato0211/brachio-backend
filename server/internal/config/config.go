package config

import (
	"sync"

	"github.com/kelseyhightower/envconfig"
)

var (
	cfg  Config
	once sync.Once
)

type Config struct {
	Server ServerConfig
	Dynamo DynamoConfig
	Common CommonConfig
}

type ServerConfig struct {
	Port int `envconfig:"PORT" default:"8080"`
}

type DynamoConfig struct {
	Region   string `envconfig:"REGION" default:"ap-northeast-1"`
	Endpoint string `envconfig:"DYNAMO_ENDPOINT" default:"http://localhost:8000"`
}

type CommonConfig struct {
	IsLocal bool `envconfig:"IS_LOCAL" default:"false"`
}

func GetConfig() (*Config, error) {
	var err error
	once.Do(func() {
		if _err := envconfig.Process("", &cfg.Server); _err != nil {
			err = _err
			return
		}

		if _err := envconfig.Process("", &cfg.Dynamo); _err != nil {
			err = _err
			return
		}

		if _err := envconfig.Process("", &cfg.Common); _err != nil {
			err = _err
			return
		}
	})

	return &cfg, err
}
