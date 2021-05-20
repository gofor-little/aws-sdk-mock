package mock

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider/cognitoidentityprovideriface"
)

// CognitoClient is a mock client for Cognito to help with unit tests.
type CognitoClient struct {
	cognitoidentityprovideriface.CognitoIdentityProviderAPI
}

// ConfirmSignUpWithContext confirms a newly signed up Cognito user.
func (c *CognitoClient) ConfirmSignUpWithContext(ctx aws.Context, input *cognitoidentityprovider.ConfirmSignUpInput, options ...request.Option) (*cognitoidentityprovider.ConfirmSignUpOutput, error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}

	return &cognitoidentityprovider.ConfirmSignUpOutput{}, nil
}

// GlobalSignOutWithContext signs a Cognito user out of all their devices.
func (c *CognitoClient) GlobalSignOutWithContext(ctx aws.Context, input *cognitoidentityprovider.GlobalSignOutInput, options ...request.Option) (*cognitoidentityprovider.GlobalSignOutOutput, error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}

	return &cognitoidentityprovider.GlobalSignOutOutput{}, nil
}

// InitiateAuthWithContext signs a Cognito user in.
func (c *CognitoClient) InitiateAuthWithContext(ctx aws.Context, input *cognitoidentityprovider.InitiateAuthInput, options ...request.Option) (*cognitoidentityprovider.InitiateAuthOutput, error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}

	return &cognitoidentityprovider.InitiateAuthOutput{
		AuthenticationResult: &cognitoidentityprovider.AuthenticationResultType{
			AccessToken:  aws.String("access-token"),
			ExpiresIn:    aws.Int64(3600),
			IdToken:      aws.String("id-token"),
			RefreshToken: aws.String("refresh-token"),
		},
	}, nil
}

// SignUpWithContext creates a new Cognito user.
func (c *CognitoClient) SignUpWithContext(ctx aws.Context, input *cognitoidentityprovider.SignUpInput, options ...request.Option) (*cognitoidentityprovider.SignUpOutput, error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}

	return &cognitoidentityprovider.SignUpOutput{
		UserSub: aws.String("user-uuid"),
	}, nil
}
