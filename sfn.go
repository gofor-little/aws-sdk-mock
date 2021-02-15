package mock

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/sfn"
	"github.com/aws/aws-sdk-go/service/sfn/sfniface"
)

// SFNClient is a mock client for StepFunctions to help with unit tests.
type SFNClient struct {
	sfniface.SFNAPI
}

// StartExecutionWithContext starts execution of a StepFunctions state machine.
func (s *SFNClient) StartExecutionWithContext(ctx aws.Context, input *sfn.StartExecutionInput, options ...request.Option) (*sfn.StartExecutionOutput, error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}

	return &sfn.StartExecutionOutput{}, nil
}
