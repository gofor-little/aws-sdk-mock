package mock

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3iface"
)

// S3Client is a mock client for S3 to help with unit tests.
type S3Client struct {
	s3iface.S3API
}

// GetObjectWithContext fetches an object from a bucket.
func (m *S3Client) GetObjectWithContext(ctx aws.Context, input *s3.GetObjectInput, options ...request.Option) (*s3.GetObjectOutput, error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}

	return &s3.GetObjectOutput{}, nil
}
