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
	conf, err := config.LoadDefaultConfig(
		ctx,
		// config.WithBaseEndpoint(cfg.Endpoint),
		config.WithRegion(cfg.Region),
	)
	if err != nil {
		return nil, err
	}

	db := dynamo.New(conf)

	return db, nil
}
