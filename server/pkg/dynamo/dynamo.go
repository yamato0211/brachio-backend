package dynamo

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
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
	var creds aws.CredentialsProvider
	if cfg.IsLocal {
		creds = credentials.NewStaticCredentialsProvider(cfg.AccessKeyID, cfg.SecretAccessKey, "")
	}

	conf, err := config.LoadDefaultConfig(
		ctx,
		config.WithCredentialsProvider(creds), // creds is nil if not local
		config.WithRegion(cfg.Region),
	)
	if err != nil {
		return nil, err
	}

	db := dynamo.New(conf)

	return db, nil
}
