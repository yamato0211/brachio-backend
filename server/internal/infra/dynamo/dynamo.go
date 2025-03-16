package dynamo

import (
	"context"

	"github.com/guregu/dynamo/v2"
	"github.com/yamato0211/brachio-backend/internal/config"
	dynamopkg "github.com/yamato0211/brachio-backend/pkg/dynamo"
)

func New(ctx context.Context, cfg config.DynamoConfig) (*dynamo.DB, error) {
	return dynamopkg.New(ctx, &dynamopkg.Config{
		IsLocal:         false,
		Region:          cfg.Region,
		Endpoint:        cfg.Endpoint,
		AccessKeyID:     "",
		SecretAccessKey: "",
	})
}
