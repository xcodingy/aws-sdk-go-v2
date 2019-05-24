// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ram

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ram-2018-01-04/CreateResourceShareRequest
type CreateResourceShareInput struct {
	_ struct{} `type:"structure"`

	// Indicates whether principals outside your organization can be associated
	// with a resource share.
	AllowExternalPrincipals *bool `locationName:"allowExternalPrincipals" type:"boolean"`

	// A unique, case-sensitive identifier that you provide to ensure the idempotency
	// of the request.
	ClientToken *string `locationName:"clientToken" type:"string"`

	// The name of the resource share.
	//
	// Name is a required field
	Name *string `locationName:"name" type:"string" required:"true"`

	// The principals to associate with the resource share. The possible values
	// are IDs of AWS accounts, the ARN of an OU or organization from AWS Organizations.
	Principals []string `locationName:"principals" type:"list"`

	// The Amazon Resource Names (ARN) of the resources to associate with the resource
	// share.
	ResourceArns []string `locationName:"resourceArns" type:"list"`

	// One or more tags.
	Tags []Tag `locationName:"tags" type:"list"`
}

// String returns the string representation
func (s CreateResourceShareInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateResourceShareInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateResourceShareInput"}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateResourceShareInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	if s.AllowExternalPrincipals != nil {
		v := *s.AllowExternalPrincipals

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "allowExternalPrincipals", protocol.BoolValue(v), metadata)
	}
	if s.ClientToken != nil {
		v := *s.ClientToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "clientToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.Principals) > 0 {
		v := s.Principals

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "principals", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	if len(s.ResourceArns) > 0 {
		v := s.ResourceArns

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "resourceArns", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	if len(s.Tags) > 0 {
		v := s.Tags

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "tags", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ram-2018-01-04/CreateResourceShareResponse
type CreateResourceShareOutput struct {
	_ struct{} `type:"structure"`

	// A unique, case-sensitive identifier that you provide to ensure the idempotency
	// of the request.
	ClientToken *string `locationName:"clientToken" type:"string"`

	// Information about the resource share.
	ResourceShare *ResourceShare `locationName:"resourceShare" type:"structure"`
}

// String returns the string representation
func (s CreateResourceShareOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateResourceShareOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.ClientToken != nil {
		v := *s.ClientToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "clientToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ResourceShare != nil {
		v := s.ResourceShare

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "resourceShare", v, metadata)
	}
	return nil
}

const opCreateResourceShare = "CreateResourceShare"

// CreateResourceShareRequest returns a request value for making API operation for
// AWS Resource Access Manager.
//
// Creates a resource share.
//
//    // Example sending a request using CreateResourceShareRequest.
//    req := client.CreateResourceShareRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ram-2018-01-04/CreateResourceShare
func (c *Client) CreateResourceShareRequest(input *CreateResourceShareInput) CreateResourceShareRequest {
	op := &aws.Operation{
		Name:       opCreateResourceShare,
		HTTPMethod: "POST",
		HTTPPath:   "/createresourceshare",
	}

	if input == nil {
		input = &CreateResourceShareInput{}
	}

	req := c.newRequest(op, input, &CreateResourceShareOutput{})
	return CreateResourceShareRequest{Request: req, Input: input, Copy: c.CreateResourceShareRequest}
}

// CreateResourceShareRequest is the request type for the
// CreateResourceShare API operation.
type CreateResourceShareRequest struct {
	*aws.Request
	Input *CreateResourceShareInput
	Copy  func(*CreateResourceShareInput) CreateResourceShareRequest
}

// Send marshals and sends the CreateResourceShare API request.
func (r CreateResourceShareRequest) Send(ctx context.Context) (*CreateResourceShareResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateResourceShareResponse{
		CreateResourceShareOutput: r.Request.Data.(*CreateResourceShareOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateResourceShareResponse is the response type for the
// CreateResourceShare API operation.
type CreateResourceShareResponse struct {
	*CreateResourceShareOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateResourceShare request.
func (r *CreateResourceShareResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
