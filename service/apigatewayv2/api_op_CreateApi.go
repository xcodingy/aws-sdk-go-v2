// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package apigatewayv2

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/apigatewayv2-2018-11-29/CreateApiRequest
type CreateApiInput struct {
	_ struct{} `type:"structure"`

	// An expression used to extract information at runtime. See Selection Expressions
	// (https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-apikey-selection-expressions)
	// for more information.
	ApiKeySelectionExpression *string `locationName:"apiKeySelectionExpression" type:"string"`

	// A string with a length between [0-1024].
	Description *string `locationName:"description" type:"string"`

	DisableSchemaValidation *bool `locationName:"disableSchemaValidation" type:"boolean"`

	// A string with a length between [1-128].
	//
	// Name is a required field
	Name *string `locationName:"name" type:"string" required:"true"`

	// ProtocolType is a required field
	ProtocolType ProtocolType `locationName:"protocolType" type:"string" required:"true" enum:"true"`

	// An expression used to extract information at runtime. See Selection Expressions
	// (https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-apikey-selection-expressions)
	// for more information.
	//
	// RouteSelectionExpression is a required field
	RouteSelectionExpression *string `locationName:"routeSelectionExpression" type:"string" required:"true"`

	// A string with a length between [1-64].
	Version *string `locationName:"version" type:"string"`
}

// String returns the string representation
func (s CreateApiInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateApiInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateApiInput"}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}
	if len(s.ProtocolType) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("ProtocolType"))
	}

	if s.RouteSelectionExpression == nil {
		invalidParams.Add(aws.NewErrParamRequired("RouteSelectionExpression"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateApiInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	if s.ApiKeySelectionExpression != nil {
		v := *s.ApiKeySelectionExpression

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "apiKeySelectionExpression", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Description != nil {
		v := *s.Description

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "description", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.DisableSchemaValidation != nil {
		v := *s.DisableSchemaValidation

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "disableSchemaValidation", protocol.BoolValue(v), metadata)
	}
	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.ProtocolType) > 0 {
		v := s.ProtocolType

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "protocolType", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.RouteSelectionExpression != nil {
		v := *s.RouteSelectionExpression

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "routeSelectionExpression", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Version != nil {
		v := *s.Version

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "version", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/apigatewayv2-2018-11-29/CreateApiResponse
type CreateApiOutput struct {
	_ struct{} `type:"structure"`

	ApiEndpoint *string `locationName:"apiEndpoint" type:"string"`

	// The identifier.
	ApiId *string `locationName:"apiId" type:"string"`

	// An expression used to extract information at runtime. See Selection Expressions
	// (https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-apikey-selection-expressions)
	// for more information.
	ApiKeySelectionExpression *string `locationName:"apiKeySelectionExpression" type:"string"`

	CreatedDate *time.Time `locationName:"createdDate" type:"timestamp" timestampFormat:"unix"`

	// A string with a length between [0-1024].
	Description *string `locationName:"description" type:"string"`

	DisableSchemaValidation *bool `locationName:"disableSchemaValidation" type:"boolean"`

	// A string with a length between [1-128].
	Name *string `locationName:"name" type:"string"`

	ProtocolType ProtocolType `locationName:"protocolType" type:"string" enum:"true"`

	// An expression used to extract information at runtime. See Selection Expressions
	// (https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-apikey-selection-expressions)
	// for more information.
	RouteSelectionExpression *string `locationName:"routeSelectionExpression" type:"string"`

	// A string with a length between [1-64].
	Version *string `locationName:"version" type:"string"`

	Warnings []string `locationName:"warnings" type:"list"`
}

// String returns the string representation
func (s CreateApiOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateApiOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.ApiEndpoint != nil {
		v := *s.ApiEndpoint

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "apiEndpoint", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ApiId != nil {
		v := *s.ApiId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "apiId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ApiKeySelectionExpression != nil {
		v := *s.ApiKeySelectionExpression

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "apiKeySelectionExpression", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.CreatedDate != nil {
		v := *s.CreatedDate

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "createdDate", protocol.TimeValue{V: v, Format: protocol.UnixTimeFormat}, metadata)
	}
	if s.Description != nil {
		v := *s.Description

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "description", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.DisableSchemaValidation != nil {
		v := *s.DisableSchemaValidation

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "disableSchemaValidation", protocol.BoolValue(v), metadata)
	}
	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.ProtocolType) > 0 {
		v := s.ProtocolType

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "protocolType", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.RouteSelectionExpression != nil {
		v := *s.RouteSelectionExpression

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "routeSelectionExpression", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Version != nil {
		v := *s.Version

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "version", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.Warnings) > 0 {
		v := s.Warnings

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "warnings", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	return nil
}

const opCreateApi = "CreateApi"

// CreateApiRequest returns a request value for making API operation for
// AmazonApiGatewayV2.
//
// Creates an Api resource.
//
//    // Example sending a request using CreateApiRequest.
//    req := client.CreateApiRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/apigatewayv2-2018-11-29/CreateApi
func (c *Client) CreateApiRequest(input *CreateApiInput) CreateApiRequest {
	op := &aws.Operation{
		Name:       opCreateApi,
		HTTPMethod: "POST",
		HTTPPath:   "/v2/apis",
	}

	if input == nil {
		input = &CreateApiInput{}
	}

	req := c.newRequest(op, input, &CreateApiOutput{})
	return CreateApiRequest{Request: req, Input: input, Copy: c.CreateApiRequest}
}

// CreateApiRequest is the request type for the
// CreateApi API operation.
type CreateApiRequest struct {
	*aws.Request
	Input *CreateApiInput
	Copy  func(*CreateApiInput) CreateApiRequest
}

// Send marshals and sends the CreateApi API request.
func (r CreateApiRequest) Send(ctx context.Context) (*CreateApiResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateApiResponse{
		CreateApiOutput: r.Request.Data.(*CreateApiOutput),
		response:        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateApiResponse is the response type for the
// CreateApi API operation.
type CreateApiResponse struct {
	*CreateApiOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateApi request.
func (r *CreateApiResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
