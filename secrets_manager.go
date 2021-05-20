package mock

import (
	"encoding/json"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
	"github.com/aws/aws-sdk-go/service/secretsmanager/secretsmanageriface"
)

// SecretsManagerClient is a mock client for secrets manager to help with unit tests.
type SecretsManagerClient struct {
	secretsmanageriface.SecretsManagerAPI
}

// GetSecretValueWithContext fetches a secret value.
func (s *SecretsManagerClient) GetSecretValueWithContext(ctx aws.Context, input *secretsmanager.GetSecretValueInput, options ...request.Option) (*secretsmanager.GetSecretValueOutput, error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}

	data, err := json.Marshal(&struct {
		Key1 string `json:"key-1"`
		Key2 string `json:"key-2"`
	}{"Value-1", "Value-2"})
	if err != nil {
		return nil, err
	}

	return &secretsmanager.GetSecretValueOutput{
		SecretString: aws.String(string(data)),
	}, nil
}
