// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package lambda

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/lambda-2015-03-31/PublishVersionRequest
type PublishVersionInput struct {
	_ struct{} `type:"structure"`

	// Only publish a version if the hash value matches the value that's specified.
	// Use this option to avoid publishing a version if the function code has changed
	// since you last updated it. You can get the hash for the version that you
	// uploaded from the output of UpdateFunctionCode.
	CodeSha256 *string `type:"string"`

	// A description for the version to override the description in the function
	// configuration.
	Description *string `type:"string"`

	// The name of the Lambda function.
	//
	// Name formats
	//
	//    * Function name - MyFunction.
	//
	//    * Function ARN - arn:aws:lambda:us-west-2:123456789012:function:MyFunction.
	//
	//    * Partial ARN - 123456789012:function:MyFunction.
	//
	// The length constraint applies only to the full ARN. If you specify only the
	// function name, it is limited to 64 characters in length.
	//
	// FunctionName is a required field
	FunctionName *string `location:"uri" locationName:"FunctionName" min:"1" type:"string" required:"true"`

	// Only update the function if the revision ID matches the ID that's specified.
	// Use this option to avoid publishing a version if the function configuration
	// has changed since you last updated it.
	RevisionId *string `type:"string"`
}

// String returns the string representation
func (s PublishVersionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PublishVersionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "PublishVersionInput"}

	if s.FunctionName == nil {
		invalidParams.Add(aws.NewErrParamRequired("FunctionName"))
	}
	if s.FunctionName != nil && len(*s.FunctionName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("FunctionName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s PublishVersionInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.CodeSha256 != nil {
		v := *s.CodeSha256

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "CodeSha256", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Description != nil {
		v := *s.Description

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Description", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.RevisionId != nil {
		v := *s.RevisionId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "RevisionId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.FunctionName != nil {
		v := *s.FunctionName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "FunctionName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Details about a function's configuration.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/lambda-2015-03-31/FunctionConfiguration
type PublishVersionOutput struct {
	_ struct{} `type:"structure"`

	// The SHA256 hash of the function's deployment package.
	CodeSha256 *string `type:"string"`

	// The size of the function's deployment package, in bytes.
	CodeSize *int64 `type:"long"`

	// The function's dead letter queue.
	DeadLetterConfig *DeadLetterConfig `type:"structure"`

	// The function's description.
	Description *string `type:"string"`

	// The function's environment variables.
	Environment *EnvironmentResponse `type:"structure"`

	// The function's Amazon Resource Name (ARN).
	FunctionArn *string `type:"string"`

	// The name of the function.
	FunctionName *string `min:"1" type:"string"`

	// The function that Lambda calls to begin executing your function.
	Handler *string `type:"string"`

	// The KMS key that's used to encrypt the function's environment variables.
	// This key is only returned if you've configured a customer-managed CMK.
	KMSKeyArn *string `type:"string"`

	// The date and time that the function was last updated, in ISO-8601 format
	// (https://www.w3.org/TR/NOTE-datetime) (YYYY-MM-DDThh:mm:ss.sTZD).
	LastModified *string `type:"string"`

	// The function's  layers (https://docs.aws.amazon.com/lambda/latest/dg/configuration-layers.html).
	Layers []Layer `type:"list"`

	// For Lambda@Edge functions, the ARN of the master function.
	MasterArn *string `type:"string"`

	// The memory that's allocated to the function.
	MemorySize *int64 `min:"128" type:"integer"`

	// The latest updated revision of the function or alias.
	RevisionId *string `type:"string"`

	// The function's execution role.
	Role *string `type:"string"`

	// The runtime environment for the Lambda function.
	Runtime Runtime `type:"string" enum:"true"`

	// The amount of time that Lambda allows a function to run before stopping it.
	Timeout *int64 `min:"1" type:"integer"`

	// The function's AWS X-Ray tracing configuration.
	TracingConfig *TracingConfigResponse `type:"structure"`

	// The version of the Lambda function.
	Version *string `min:"1" type:"string"`

	// The function's networking configuration.
	VpcConfig *VpcConfigResponse `type:"structure"`
}

// String returns the string representation
func (s PublishVersionOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s PublishVersionOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.CodeSha256 != nil {
		v := *s.CodeSha256

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "CodeSha256", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.CodeSize != nil {
		v := *s.CodeSize

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "CodeSize", protocol.Int64Value(v), metadata)
	}
	if s.DeadLetterConfig != nil {
		v := s.DeadLetterConfig

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "DeadLetterConfig", v, metadata)
	}
	if s.Description != nil {
		v := *s.Description

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Description", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Environment != nil {
		v := s.Environment

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "Environment", v, metadata)
	}
	if s.FunctionArn != nil {
		v := *s.FunctionArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "FunctionArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.FunctionName != nil {
		v := *s.FunctionName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "FunctionName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Handler != nil {
		v := *s.Handler

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Handler", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.KMSKeyArn != nil {
		v := *s.KMSKeyArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "KMSKeyArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.LastModified != nil {
		v := *s.LastModified

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "LastModified", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.Layers) > 0 {
		v := s.Layers

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "Layers", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.MasterArn != nil {
		v := *s.MasterArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "MasterArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.MemorySize != nil {
		v := *s.MemorySize

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "MemorySize", protocol.Int64Value(v), metadata)
	}
	if s.RevisionId != nil {
		v := *s.RevisionId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "RevisionId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Role != nil {
		v := *s.Role

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Role", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.Runtime) > 0 {
		v := s.Runtime

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Runtime", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.Timeout != nil {
		v := *s.Timeout

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Timeout", protocol.Int64Value(v), metadata)
	}
	if s.TracingConfig != nil {
		v := s.TracingConfig

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "TracingConfig", v, metadata)
	}
	if s.Version != nil {
		v := *s.Version

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Version", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.VpcConfig != nil {
		v := s.VpcConfig

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "VpcConfig", v, metadata)
	}
	return nil
}

const opPublishVersion = "PublishVersion"

// PublishVersionRequest returns a request value for making API operation for
// AWS Lambda.
//
// Creates a version (https://docs.aws.amazon.com/lambda/latest/dg/versioning-aliases.html)
// from the current code and configuration of a function. Use versions to create
// a snapshot of your function code and configuration that doesn't change.
//
// AWS Lambda doesn't publish a version if the function's configuration and
// code haven't changed since the last version. Use UpdateFunctionCode or UpdateFunctionConfiguration
// to update the function before publishing a version.
//
// Clients can invoke versions directly or with an alias. To create an alias,
// use CreateAlias.
//
//    // Example sending a request using PublishVersionRequest.
//    req := client.PublishVersionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/lambda-2015-03-31/PublishVersion
func (c *Client) PublishVersionRequest(input *PublishVersionInput) PublishVersionRequest {
	op := &aws.Operation{
		Name:       opPublishVersion,
		HTTPMethod: "POST",
		HTTPPath:   "/2015-03-31/functions/{FunctionName}/versions",
	}

	if input == nil {
		input = &PublishVersionInput{}
	}

	req := c.newRequest(op, input, &PublishVersionOutput{})
	return PublishVersionRequest{Request: req, Input: input, Copy: c.PublishVersionRequest}
}

// PublishVersionRequest is the request type for the
// PublishVersion API operation.
type PublishVersionRequest struct {
	*aws.Request
	Input *PublishVersionInput
	Copy  func(*PublishVersionInput) PublishVersionRequest
}

// Send marshals and sends the PublishVersion API request.
func (r PublishVersionRequest) Send(ctx context.Context) (*PublishVersionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &PublishVersionResponse{
		PublishVersionOutput: r.Request.Data.(*PublishVersionOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// PublishVersionResponse is the response type for the
// PublishVersion API operation.
type PublishVersionResponse struct {
	*PublishVersionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// PublishVersion request.
func (r *PublishVersionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
