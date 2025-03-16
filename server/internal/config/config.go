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
	Dynamo *DynamoConfig
}

type DynamoConfig struct {
	Region   string `envconfig:"REGION" default:"us-east-1"`
	Endpoint string `envconfig:"DYNAMO_ENDPOINT" default:"http://localhost:8000"`
}

func GetConfig() (*Config, error) {
	var err error
	once.Do(func() {
		if _err := envconfig.Process("", &cfg.Dynamo); _err != nil {
			err = _err
			return
		}

	})

	return &cfg, err
}
