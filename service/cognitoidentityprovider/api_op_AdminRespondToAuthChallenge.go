// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cognitoidentityprovider

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// The request to respond to the authentication challenge, as an administrator.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cognito-idp-2016-04-18/AdminRespondToAuthChallengeRequest
type AdminRespondToAuthChallengeInput struct {
	_ struct{} `type:"structure"`

	// The analytics metadata for collecting Amazon Pinpoint metrics for AdminRespondToAuthChallenge
	// calls.
	AnalyticsMetadata *AnalyticsMetadataType `type:"structure"`

	// The challenge name. For more information, see .
	//
	// ChallengeName is a required field
	ChallengeName ChallengeNameType `type:"string" required:"true" enum:"true"`

	// The challenge responses. These are inputs corresponding to the value of ChallengeName,
	// for example:
	//
	//    * SMS_MFA: SMS_MFA_CODE, USERNAME, SECRET_HASH (if app client is configured
	//    with client secret).
	//
	//    * PASSWORD_VERIFIER: PASSWORD_CLAIM_SIGNATURE, PASSWORD_CLAIM_SECRET_BLOCK,
	//    TIMESTAMP, USERNAME, SECRET_HASH (if app client is configured with client
	//    secret).
	//
	//    * ADMIN_NO_SRP_AUTH: PASSWORD, USERNAME, SECRET_HASH (if app client is
	//    configured with client secret).
	//
	//    * NEW_PASSWORD_REQUIRED: NEW_PASSWORD, any other required attributes,
	//    USERNAME, SECRET_HASH (if app client is configured with client secret).
	//
	//
	// The value of the USERNAME attribute must be the user's actual username, not
	// an alias (such as email address or phone number). To make this easier, the
	// AdminInitiateAuth response includes the actual username value in the USERNAMEUSER_ID_FOR_SRP
	// attribute, even if you specified an alias in your call to AdminInitiateAuth.
	ChallengeResponses map[string]string `type:"map"`

	// The app client ID.
	//
	// ClientId is a required field
	ClientId *string `min:"1" type:"string" required:"true"`

	// Contextual data such as the user's device fingerprint, IP address, or location
	// used for evaluating the risk of an unexpected event by Amazon Cognito advanced
	// security.
	ContextData *ContextDataType `type:"structure"`

	// The session which should be passed both ways in challenge-response calls
	// to the service. If InitiateAuth or RespondToAuthChallenge API call determines
	// that the caller needs to go through another challenge, they return a session
	// with other challenge parameters. This session should be passed as it is to
	// the next RespondToAuthChallenge API call.
	Session *string `min:"20" type:"string"`

	// The ID of the Amazon Cognito user pool.
	//
	// UserPoolId is a required field
	UserPoolId *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s AdminRespondToAuthChallengeInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AdminRespondToAuthChallengeInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "AdminRespondToAuthChallengeInput"}
	if len(s.ChallengeName) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("ChallengeName"))
	}

	if s.ClientId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ClientId"))
	}
	if s.ClientId != nil && len(*s.ClientId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ClientId", 1))
	}
	if s.Session != nil && len(*s.Session) < 20 {
		invalidParams.Add(aws.NewErrParamMinLen("Session", 20))
	}

	if s.UserPoolId == nil {
		invalidParams.Add(aws.NewErrParamRequired("UserPoolId"))
	}
	if s.UserPoolId != nil && len(*s.UserPoolId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("UserPoolId", 1))
	}
	if s.ContextData != nil {
		if err := s.ContextData.Validate(); err != nil {
			invalidParams.AddNested("ContextData", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Responds to the authentication challenge, as an administrator.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cognito-idp-2016-04-18/AdminRespondToAuthChallengeResponse
type AdminRespondToAuthChallengeOutput struct {
	_ struct{} `type:"structure"`

	// The result returned by the server in response to the authentication request.
	AuthenticationResult *AuthenticationResultType `type:"structure"`

	// The name of the challenge. For more information, see .
	ChallengeName ChallengeNameType `type:"string" enum:"true"`

	// The challenge parameters. For more information, see .
	ChallengeParameters map[string]string `type:"map"`

	// The session which should be passed both ways in challenge-response calls
	// to the service. If the or API call determines that the caller needs to go
	// through another challenge, they return a session with other challenge parameters.
	// This session should be passed as it is to the next RespondToAuthChallenge
	// API call.
	Session *string `min:"20" type:"string"`
}

// String returns the string representation
func (s AdminRespondToAuthChallengeOutput) String() string {
	return awsutil.Prettify(s)
}

const opAdminRespondToAuthChallenge = "AdminRespondToAuthChallenge"

// AdminRespondToAuthChallengeRequest returns a request value for making API operation for
// Amazon Cognito Identity Provider.
//
// Responds to an authentication challenge, as an administrator.
//
// Requires developer credentials.
//
//    // Example sending a request using AdminRespondToAuthChallengeRequest.
//    req := client.AdminRespondToAuthChallengeRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cognito-idp-2016-04-18/AdminRespondToAuthChallenge
func (c *Client) AdminRespondToAuthChallengeRequest(input *AdminRespondToAuthChallengeInput) AdminRespondToAuthChallengeRequest {
	op := &aws.Operation{
		Name:       opAdminRespondToAuthChallenge,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &AdminRespondToAuthChallengeInput{}
	}

	req := c.newRequest(op, input, &AdminRespondToAuthChallengeOutput{})
	return AdminRespondToAuthChallengeRequest{Request: req, Input: input, Copy: c.AdminRespondToAuthChallengeRequest}
}

// AdminRespondToAuthChallengeRequest is the request type for the
// AdminRespondToAuthChallenge API operation.
type AdminRespondToAuthChallengeRequest struct {
	*aws.Request
	Input *AdminRespondToAuthChallengeInput
	Copy  func(*AdminRespondToAuthChallengeInput) AdminRespondToAuthChallengeRequest
}

// Send marshals and sends the AdminRespondToAuthChallenge API request.
func (r AdminRespondToAuthChallengeRequest) Send(ctx context.Context) (*AdminRespondToAuthChallengeResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &AdminRespondToAuthChallengeResponse{
		AdminRespondToAuthChallengeOutput: r.Request.Data.(*AdminRespondToAuthChallengeOutput),
		response:                          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// AdminRespondToAuthChallengeResponse is the response type for the
// AdminRespondToAuthChallenge API operation.
type AdminRespondToAuthChallengeResponse struct {
	*AdminRespondToAuthChallengeOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// AdminRespondToAuthChallenge request.
func (r *AdminRespondToAuthChallengeResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
