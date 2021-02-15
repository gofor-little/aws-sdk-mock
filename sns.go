package mock

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/aws/aws-sdk-go/service/sns/snsiface"
)

// SNSClient is a mock client for SNS to help with unit tests.
type SNSClient struct {
	snsiface.SNSAPI
}

// PublishWithContext publishes an SNS notification.
func (m *SNSClient) PublishWithContext(ctx aws.Context, input *sns.PublishInput, options ...request.Option) (*sns.PublishOutput, error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}

	return &sns.PublishOutput{
		MessageId: aws.String("message-id"),
	}, nil
}
