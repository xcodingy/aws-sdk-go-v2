// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package appsync

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/appsync-2017-07-25/CreateFunctionRequest
type CreateFunctionInput struct {
	_ struct{} `type:"structure"`

	// The GraphQL API ID.
	//
	// ApiId is a required field
	ApiId *string `location:"uri" locationName:"apiId" type:"string" required:"true"`

	// The FunctionDataSource name.
	//
	// DataSourceName is a required field
	DataSourceName *string `locationName:"dataSourceName" type:"string" required:"true"`

	// The Function description.
	Description *string `locationName:"description" type:"string"`

	// The version of the request mapping template. Currently the supported value
	// is 2018-05-29.
	//
	// FunctionVersion is a required field
	FunctionVersion *string `locationName:"functionVersion" type:"string" required:"true"`

	// The Function name. The function name does not have to be unique.
	//
	// Name is a required field
	Name *string `locationName:"name" type:"string" required:"true"`

	// The Function request mapping template. Functions support only the 2018-05-29
	// version of the request mapping template.
	//
	// RequestMappingTemplate is a required field
	RequestMappingTemplate *string `locationName:"requestMappingTemplate" min:"1" type:"string" required:"true"`

	// The Function response mapping template.
	ResponseMappingTemplate *string `locationName:"responseMappingTemplate" min:"1" type:"string"`
}

// String returns the string representation
func (s CreateFunctionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateFunctionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateFunctionInput"}

	if s.ApiId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ApiId"))
	}

	if s.DataSourceName == nil {
		invalidParams.Add(aws.NewErrParamRequired("DataSourceName"))
	}

	if s.FunctionVersion == nil {
		invalidParams.Add(aws.NewErrParamRequired("FunctionVersion"))
	}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}

	if s.RequestMappingTemplate == nil {
		invalidParams.Add(aws.NewErrParamRequired("RequestMappingTemplate"))
	}
	if s.RequestMappingTemplate != nil && len(*s.RequestMappingTemplate) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("RequestMappingTemplate", 1))
	}
	if s.ResponseMappingTemplate != nil && len(*s.ResponseMappingTemplate) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ResponseMappingTemplate", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateFunctionInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	if s.DataSourceName != nil {
		v := *s.DataSourceName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "dataSourceName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Description != nil {
		v := *s.Description

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "description", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.FunctionVersion != nil {
		v := *s.FunctionVersion

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "functionVersion", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.RequestMappingTemplate != nil {
		v := *s.RequestMappingTemplate

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "requestMappingTemplate", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ResponseMappingTemplate != nil {
		v := *s.ResponseMappingTemplate

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "responseMappingTemplate", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ApiId != nil {
		v := *s.ApiId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "apiId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/appsync-2017-07-25/CreateFunctionResponse
type CreateFunctionOutput struct {
	_ struct{} `type:"structure"`

	// The Function object.
	FunctionConfiguration *FunctionConfiguration `locationName:"functionConfiguration" type:"structure"`
}

// String returns the string representation
func (s CreateFunctionOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateFunctionOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.FunctionConfiguration != nil {
		v := s.FunctionConfiguration

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "functionConfiguration", v, metadata)
	}
	return nil
}

const opCreateFunction = "CreateFunction"

// CreateFunctionRequest returns a request value for making API operation for
// AWS AppSync.
//
// Creates a Function object.
//
// A function is a reusable entity. Multiple functions can be used to compose
// the resolver logic.
//
//    // Example sending a request using CreateFunctionRequest.
//    req := client.CreateFunctionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/appsync-2017-07-25/CreateFunction
func (c *Client) CreateFunctionRequest(input *CreateFunctionInput) CreateFunctionRequest {
	op := &aws.Operation{
		Name:       opCreateFunction,
		HTTPMethod: "POST",
		HTTPPath:   "/v1/apis/{apiId}/functions",
	}

	if input == nil {
		input = &CreateFunctionInput{}
	}

	req := c.newRequest(op, input, &CreateFunctionOutput{})
	return CreateFunctionRequest{Request: req, Input: input, Copy: c.CreateFunctionRequest}
}

// CreateFunctionRequest is the request type for the
// CreateFunction API operation.
type CreateFunctionRequest struct {
	*aws.Request
	Input *CreateFunctionInput
	Copy  func(*CreateFunctionInput) CreateFunctionRequest
}

// Send marshals and sends the CreateFunction API request.
func (r CreateFunctionRequest) Send(ctx context.Context) (*CreateFunctionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateFunctionResponse{
		CreateFunctionOutput: r.Request.Data.(*CreateFunctionOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateFunctionResponse is the response type for the
// CreateFunction API operation.
type CreateFunctionResponse struct {
	*CreateFunctionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateFunction request.
func (r *CreateFunctionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
