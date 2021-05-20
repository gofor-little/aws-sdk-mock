package mock

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/ses"
	"github.com/aws/aws-sdk-go/service/ses/sesiface"
)

// SESClient is a mock client for SES to help with unit tests.
type SESClient struct {
	sesiface.SESAPI
}

// SendRawEmailWithContext sends an email via SES.
func (s *SESClient) SendRawEmailWithContext(ctx aws.Context, input *ses.SendRawEmailInput, options ...request.Option) (*ses.SendRawEmailOutput, error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}

	return &ses.SendRawEmailOutput{
		MessageId: aws.String("message-id"),
	}, nil
}
