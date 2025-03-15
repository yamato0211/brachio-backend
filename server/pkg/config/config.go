package config

import (
	"cmp"
	"os"
)

type DynamoConfig struct {
	Region   string
	Endpoint string
}

func NewDynamoConfig() *DynamoConfig {
	return &DynamoConfig{
		Region:   cmp.Or(os.Getenv("REGION"), "us-east-1"),
		Endpoint: cmp.Or(os.Getenv("DYNAMO_ENDPOINT"), "http://localhost:8000"),
	}
}
