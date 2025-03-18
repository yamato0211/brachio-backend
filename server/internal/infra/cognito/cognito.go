package cognito

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/yamato0211/brachio-backend/internal/config"
	cognitopkg "github.com/yamato0211/brachio-backend/pkg/cognito"
)

func New(ctx context.Context, cfg config.CognitoConfig) (*cognitoidentityprovider.Client, error) {
	client, err := cognitopkg.New(ctx, &cognitopkg.Config{
		Region: cfg.Region,
	})
	if err != nil {
		return nil, err
	}
	client.VerifySoftwareToken(context.Background(), &cognitoidentityprovider.VerifySoftwareTokenInput{})
	return client, nil
}
