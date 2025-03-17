package dynamo

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/guregu/dynamo/v2"
)

type Config struct {
	IsLocal         bool
	Region          string
	Endpoint        string
	AccessKeyID     string
	SecretAccessKey string
}

func New(ctx context.Context, cfg *Config) (*dynamo.DB, error) {
	opts := []func(*config.LoadOptions) error{}
	if cfg.IsLocal {
		opts = append(opts, config.WithBaseEndpoint(cfg.Endpoint))
	}
	opts = append(opts, config.WithRegion(cfg.Region))

	conf, err := config.LoadDefaultConfig(
		ctx,
		opts...,
	)
	if err != nil {
		return nil, err
	}

	db := dynamo.New(conf)

	return db, nil
}
