package mock

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/aws/aws-sdk-go/service/sqs/sqsiface"
)

type SQSClient struct {
	sqsiface.SQSAPI
}

func (s *SQSClient) SendMessageWithContext(ctx aws.Context, input *sqs.SendMessageInput, options ...request.Option) (*sqs.SendMessageOutput, error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}

	return &sqs.SendMessageOutput{}, nil
}
