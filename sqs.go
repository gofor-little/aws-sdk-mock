package mock

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/aws/aws-sdk-go/service/sqs/sqsiface"
)

// SQSClient is a mock client for SQS to help with unit tests.
type SQSClient struct {
	sqsiface.SQSAPI
}

// SendMessageWithContext sends a message to an SQS queue.
func (s *SQSClient) SendMessageWithContext(ctx aws.Context, input *sqs.SendMessageInput, options ...request.Option) (*sqs.SendMessageOutput, error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}

	return &sqs.SendMessageOutput{}, nil
}
