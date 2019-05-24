// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package apigateway

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Request to add a new Authorizer to an existing RestApi resource.
type CreateAuthorizerInput struct {
	_ struct{} `type:"structure"`

	// Optional customer-defined field, used in OpenAPI imports and exports without
	// functional impact.
	AuthType *string `locationName:"authType" type:"string"`

	// Specifies the required credentials as an IAM role for API Gateway to invoke
	// the authorizer. To specify an IAM role for API Gateway to assume, use the
	// role's Amazon Resource Name (ARN). To use resource-based permissions on the
	// Lambda function, specify null.
	AuthorizerCredentials *string `locationName:"authorizerCredentials" type:"string"`

	// The TTL in seconds of cached authorizer results. If it equals 0, authorization
	// caching is disabled. If it is greater than 0, API Gateway will cache authorizer
	// responses. If this field is not set, the default value is 300. The maximum
	// value is 3600, or 1 hour.
	AuthorizerResultTtlInSeconds *int64 `locationName:"authorizerResultTtlInSeconds" type:"integer"`

	// Specifies the authorizer's Uniform Resource Identifier (URI). For TOKEN or
	// REQUEST authorizers, this must be a well-formed Lambda function URI, for
	// example, arn:aws:apigateway:us-west-2:lambda:path/2015-03-31/functions/arn:aws:lambda:us-west-2:{account_id}:function:{lambda_function_name}/invocations.
	// In general, the URI has this form arn:aws:apigateway:{region}:lambda:path/{service_api},
	// where {region} is the same as the region hosting the Lambda function, path
	// indicates that the remaining substring in the URI should be treated as the
	// path to the resource, including the initial /. For Lambda functions, this
	// is usually of the form /2015-03-31/functions/[FunctionARN]/invocations.
	AuthorizerUri *string `locationName:"authorizerUri" type:"string"`

	// The identity source for which authorization is requested. For a TOKEN or
	// COGNITO_USER_POOLS authorizer, this is required and specifies the request
	// header mapping expression for the custom header holding the authorization
	// token submitted by the client. For example, if the token header name is Auth,
	// the header mapping expression is method.request.header.Auth.
	// For the REQUEST authorizer, this is required when authorization caching is
	// enabled. The value is a comma-separated string of one or more mapping expressions
	// of the specified request parameters. For example, if an Auth header, a Name
	// query string parameter are defined as identity sources, this value is method.request.header.Auth,
	// method.request.querystring.Name. These parameters will be used to derive
	// the authorization caching key and to perform runtime validation of the REQUEST
	// authorizer by verifying all of the identity-related request parameters are
	// present, not null and non-empty. Only when this is true does the authorizer
	// invoke the authorizer Lambda function, otherwise, it returns a 401 Unauthorized
	// response without calling the Lambda function. The valid value is a string
	// of comma-separated mapping expressions of the specified request parameters.
	// When the authorization caching is not enabled, this property is optional.
	IdentitySource *string `locationName:"identitySource" type:"string"`

	// A validation expression for the incoming identity token. For TOKEN authorizers,
	// this value is a regular expression. API Gateway will match the aud field
	// of the incoming token from the client against the specified regular expression.
	// It will invoke the authorizer's Lambda function when there is a match. Otherwise,
	// it will return a 401 Unauthorized response without calling the Lambda function.
	// The validation expression does not apply to the REQUEST authorizer.
	IdentityValidationExpression *string `locationName:"identityValidationExpression" type:"string"`

	// [Required] The name of the authorizer.
	//
	// Name is a required field
	Name *string `locationName:"name" type:"string" required:"true"`

	// A list of the Amazon Cognito user pool ARNs for the COGNITO_USER_POOLS authorizer.
	// Each element is of this format: arn:aws:cognito-idp:{region}:{account_id}:userpool/{user_pool_id}.
	// For a TOKEN or REQUEST authorizer, this is not defined.
	ProviderARNs []string `locationName:"providerARNs" type:"list"`

	// [Required] The string identifier of the associated RestApi.
	//
	// RestApiId is a required field
	RestApiId *string `location:"uri" locationName:"restapi_id" type:"string" required:"true"`

	// [Required] The authorizer type. Valid values are TOKEN for a Lambda function
	// using a single authorization token submitted in a custom header, REQUEST
	// for a Lambda function using incoming request parameters, and COGNITO_USER_POOLS
	// for using an Amazon Cognito user pool.
	//
	// Type is a required field
	Type AuthorizerType `locationName:"type" type:"string" required:"true" enum:"true"`
}

// String returns the string representation
func (s CreateAuthorizerInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateAuthorizerInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateAuthorizerInput"}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}

	if s.RestApiId == nil {
		invalidParams.Add(aws.NewErrParamRequired("RestApiId"))
	}
	if len(s.Type) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("Type"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateAuthorizerInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.AuthType != nil {
		v := *s.AuthType

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "authType", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.AuthorizerCredentials != nil {
		v := *s.AuthorizerCredentials

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "authorizerCredentials", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.AuthorizerResultTtlInSeconds != nil {
		v := *s.AuthorizerResultTtlInSeconds

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "authorizerResultTtlInSeconds", protocol.Int64Value(v), metadata)
	}
	if s.AuthorizerUri != nil {
		v := *s.AuthorizerUri

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "authorizerUri", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.IdentitySource != nil {
		v := *s.IdentitySource

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "identitySource", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.IdentityValidationExpression != nil {
		v := *s.IdentityValidationExpression

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "identityValidationExpression", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.ProviderARNs) > 0 {
		v := s.ProviderARNs

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "providerARNs", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	if len(s.Type) > 0 {
		v := s.Type

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "type", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.RestApiId != nil {
		v := *s.RestApiId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "restapi_id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Represents an authorization layer for methods. If enabled on a method, API
// Gateway will activate the authorizer when a client calls the method.
//
// Enable custom authorization (https://docs.aws.amazon.com/apigateway/latest/developerguide/use-custom-authorizer.html)
type CreateAuthorizerOutput struct {
	_ struct{} `type:"structure"`

	// Optional customer-defined field, used in OpenAPI imports and exports without
	// functional impact.
	AuthType *string `locationName:"authType" type:"string"`

	// Specifies the required credentials as an IAM role for API Gateway to invoke
	// the authorizer. To specify an IAM role for API Gateway to assume, use the
	// role's Amazon Resource Name (ARN). To use resource-based permissions on the
	// Lambda function, specify null.
	AuthorizerCredentials *string `locationName:"authorizerCredentials" type:"string"`

	// The TTL in seconds of cached authorizer results. If it equals 0, authorization
	// caching is disabled. If it is greater than 0, API Gateway will cache authorizer
	// responses. If this field is not set, the default value is 300. The maximum
	// value is 3600, or 1 hour.
	AuthorizerResultTtlInSeconds *int64 `locationName:"authorizerResultTtlInSeconds" type:"integer"`

	// Specifies the authorizer's Uniform Resource Identifier (URI). For TOKEN or
	// REQUEST authorizers, this must be a well-formed Lambda function URI, for
	// example, arn:aws:apigateway:us-west-2:lambda:path/2015-03-31/functions/arn:aws:lambda:us-west-2:{account_id}:function:{lambda_function_name}/invocations.
	// In general, the URI has this form arn:aws:apigateway:{region}:lambda:path/{service_api},
	// where {region} is the same as the region hosting the Lambda function, path
	// indicates that the remaining substring in the URI should be treated as the
	// path to the resource, including the initial /. For Lambda functions, this
	// is usually of the form /2015-03-31/functions/[FunctionARN]/invocations.
	AuthorizerUri *string `locationName:"authorizerUri" type:"string"`

	// The identifier for the authorizer resource.
	Id *string `locationName:"id" type:"string"`

	// The identity source for which authorization is requested. For a TOKEN or
	// COGNITO_USER_POOLS authorizer, this is required and specifies the request
	// header mapping expression for the custom header holding the authorization
	// token submitted by the client. For example, if the token header name is Auth,
	// the header mapping expression is method.request.header.Auth.
	// For the REQUEST authorizer, this is required when authorization caching is
	// enabled. The value is a comma-separated string of one or more mapping expressions
	// of the specified request parameters. For example, if an Auth header, a Name
	// query string parameter are defined as identity sources, this value is method.request.header.Auth,
	// method.request.querystring.Name. These parameters will be used to derive
	// the authorization caching key and to perform runtime validation of the REQUEST
	// authorizer by verifying all of the identity-related request parameters are
	// present, not null and non-empty. Only when this is true does the authorizer
	// invoke the authorizer Lambda function, otherwise, it returns a 401 Unauthorized
	// response without calling the Lambda function. The valid value is a string
	// of comma-separated mapping expressions of the specified request parameters.
	// When the authorization caching is not enabled, this property is optional.
	IdentitySource *string `locationName:"identitySource" type:"string"`

	// A validation expression for the incoming identity token. For TOKEN authorizers,
	// this value is a regular expression. API Gateway will match the aud field
	// of the incoming token from the client against the specified regular expression.
	// It will invoke the authorizer's Lambda function when there is a match. Otherwise,
	// it will return a 401 Unauthorized response without calling the Lambda function.
	// The validation expression does not apply to the REQUEST authorizer.
	IdentityValidationExpression *string `locationName:"identityValidationExpression" type:"string"`

	// [Required] The name of the authorizer.
	Name *string `locationName:"name" type:"string"`

	// A list of the Amazon Cognito user pool ARNs for the COGNITO_USER_POOLS authorizer.
	// Each element is of this format: arn:aws:cognito-idp:{region}:{account_id}:userpool/{user_pool_id}.
	// For a TOKEN or REQUEST authorizer, this is not defined.
	ProviderARNs []string `locationName:"providerARNs" type:"list"`

	// The authorizer type. Valid values are TOKEN for a Lambda function using a
	// single authorization token submitted in a custom header, REQUEST for a Lambda
	// function using incoming request parameters, and COGNITO_USER_POOLS for using
	// an Amazon Cognito user pool.
	Type AuthorizerType `locationName:"type" type:"string" enum:"true"`
}

// String returns the string representation
func (s CreateAuthorizerOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateAuthorizerOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.AuthType != nil {
		v := *s.AuthType

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "authType", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.AuthorizerCredentials != nil {
		v := *s.AuthorizerCredentials

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "authorizerCredentials", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.AuthorizerResultTtlInSeconds != nil {
		v := *s.AuthorizerResultTtlInSeconds

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "authorizerResultTtlInSeconds", protocol.Int64Value(v), metadata)
	}
	if s.AuthorizerUri != nil {
		v := *s.AuthorizerUri

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "authorizerUri", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Id != nil {
		v := *s.Id

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.IdentitySource != nil {
		v := *s.IdentitySource

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "identitySource", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.IdentityValidationExpression != nil {
		v := *s.IdentityValidationExpression

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "identityValidationExpression", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.ProviderARNs) > 0 {
		v := s.ProviderARNs

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "providerARNs", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	if len(s.Type) > 0 {
		v := s.Type

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "type", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	return nil
}

const opCreateAuthorizer = "CreateAuthorizer"

// CreateAuthorizerRequest returns a request value for making API operation for
// Amazon API Gateway.
//
// Adds a new Authorizer resource to an existing RestApi resource.
//
// AWS CLI (https://docs.aws.amazon.com/cli/latest/reference/apigateway/create-authorizer.html)
//
//    // Example sending a request using CreateAuthorizerRequest.
//    req := client.CreateAuthorizerRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) CreateAuthorizerRequest(input *CreateAuthorizerInput) CreateAuthorizerRequest {
	op := &aws.Operation{
		Name:       opCreateAuthorizer,
		HTTPMethod: "POST",
		HTTPPath:   "/restapis/{restapi_id}/authorizers",
	}

	if input == nil {
		input = &CreateAuthorizerInput{}
	}

	req := c.newRequest(op, input, &CreateAuthorizerOutput{})
	return CreateAuthorizerRequest{Request: req, Input: input, Copy: c.CreateAuthorizerRequest}
}

// CreateAuthorizerRequest is the request type for the
// CreateAuthorizer API operation.
type CreateAuthorizerRequest struct {
	*aws.Request
	Input *CreateAuthorizerInput
	Copy  func(*CreateAuthorizerInput) CreateAuthorizerRequest
}

// Send marshals and sends the CreateAuthorizer API request.
func (r CreateAuthorizerRequest) Send(ctx context.Context) (*CreateAuthorizerResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateAuthorizerResponse{
		CreateAuthorizerOutput: r.Request.Data.(*CreateAuthorizerOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateAuthorizerResponse is the response type for the
// CreateAuthorizer API operation.
type CreateAuthorizerResponse struct {
	*CreateAuthorizerOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateAuthorizer request.
func (r *CreateAuthorizerResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
