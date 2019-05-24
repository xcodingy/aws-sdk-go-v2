// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package greengrass

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/greengrass-2017-06-07/CreateGroupVersionRequest
type CreateGroupVersionInput struct {
	_ struct{} `type:"structure"`

	AmznClientToken *string `location:"header" locationName:"X-Amzn-Client-Token" type:"string"`

	ConnectorDefinitionVersionArn *string `type:"string"`

	CoreDefinitionVersionArn *string `type:"string"`

	DeviceDefinitionVersionArn *string `type:"string"`

	FunctionDefinitionVersionArn *string `type:"string"`

	// GroupId is a required field
	GroupId *string `location:"uri" locationName:"GroupId" type:"string" required:"true"`

	LoggerDefinitionVersionArn *string `type:"string"`

	ResourceDefinitionVersionArn *string `type:"string"`

	SubscriptionDefinitionVersionArn *string `type:"string"`
}

// String returns the string representation
func (s CreateGroupVersionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateGroupVersionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateGroupVersionInput"}

	if s.GroupId == nil {
		invalidParams.Add(aws.NewErrParamRequired("GroupId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateGroupVersionInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	if s.ConnectorDefinitionVersionArn != nil {
		v := *s.ConnectorDefinitionVersionArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "ConnectorDefinitionVersionArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.CoreDefinitionVersionArn != nil {
		v := *s.CoreDefinitionVersionArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "CoreDefinitionVersionArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.DeviceDefinitionVersionArn != nil {
		v := *s.DeviceDefinitionVersionArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "DeviceDefinitionVersionArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.FunctionDefinitionVersionArn != nil {
		v := *s.FunctionDefinitionVersionArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "FunctionDefinitionVersionArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.LoggerDefinitionVersionArn != nil {
		v := *s.LoggerDefinitionVersionArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "LoggerDefinitionVersionArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ResourceDefinitionVersionArn != nil {
		v := *s.ResourceDefinitionVersionArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "ResourceDefinitionVersionArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.SubscriptionDefinitionVersionArn != nil {
		v := *s.SubscriptionDefinitionVersionArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "SubscriptionDefinitionVersionArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.AmznClientToken != nil {
		v := *s.AmznClientToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "X-Amzn-Client-Token", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.GroupId != nil {
		v := *s.GroupId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "GroupId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/greengrass-2017-06-07/CreateGroupVersionResponse
type CreateGroupVersionOutput struct {
	_ struct{} `type:"structure"`

	Arn *string `type:"string"`

	CreationTimestamp *string `type:"string"`

	Id *string `type:"string"`

	Version *string `type:"string"`
}

// String returns the string representation
func (s CreateGroupVersionOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateGroupVersionOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Arn != nil {
		v := *s.Arn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Arn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.CreationTimestamp != nil {
		v := *s.CreationTimestamp

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "CreationTimestamp", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Id != nil {
		v := *s.Id

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Version != nil {
		v := *s.Version

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Version", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opCreateGroupVersion = "CreateGroupVersion"

// CreateGroupVersionRequest returns a request value for making API operation for
// AWS Greengrass.
//
// Creates a version of a group which has already been defined.
//
//    // Example sending a request using CreateGroupVersionRequest.
//    req := client.CreateGroupVersionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/greengrass-2017-06-07/CreateGroupVersion
func (c *Client) CreateGroupVersionRequest(input *CreateGroupVersionInput) CreateGroupVersionRequest {
	op := &aws.Operation{
		Name:       opCreateGroupVersion,
		HTTPMethod: "POST",
		HTTPPath:   "/greengrass/groups/{GroupId}/versions",
	}

	if input == nil {
		input = &CreateGroupVersionInput{}
	}

	req := c.newRequest(op, input, &CreateGroupVersionOutput{})
	return CreateGroupVersionRequest{Request: req, Input: input, Copy: c.CreateGroupVersionRequest}
}

// CreateGroupVersionRequest is the request type for the
// CreateGroupVersion API operation.
type CreateGroupVersionRequest struct {
	*aws.Request
	Input *CreateGroupVersionInput
	Copy  func(*CreateGroupVersionInput) CreateGroupVersionRequest
}

// Send marshals and sends the CreateGroupVersion API request.
func (r CreateGroupVersionRequest) Send(ctx context.Context) (*CreateGroupVersionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateGroupVersionResponse{
		CreateGroupVersionOutput: r.Request.Data.(*CreateGroupVersionOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateGroupVersionResponse is the response type for the
// CreateGroupVersion API operation.
type CreateGroupVersionResponse struct {
	*CreateGroupVersionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateGroupVersion request.
func (r *CreateGroupVersionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
