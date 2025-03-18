package cognito

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
)

type Config struct {
	Region string
}

func New(ctx context.Context, cfg *Config) (*cognitoidentityprovider.Client, error) {
	awsCfg, err := config.LoadDefaultConfig(ctx, config.WithRegion(cfg.Region))
	if err != nil {
		return nil, err
	}
	client := cognitoidentityprovider.NewFromConfig(awsCfg)
	return client, nil
}
