// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ram

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ram-2018-01-04/AssociateResourceShareRequest
type AssociateResourceShareInput struct {
	_ struct{} `type:"structure"`

	// A unique, case-sensitive identifier that you provide to ensure the idempotency
	// of the request.
	ClientToken *string `locationName:"clientToken" type:"string"`

	// The principals.
	Principals []string `locationName:"principals" type:"list"`

	// The Amazon Resource Names (ARN) of the resources.
	ResourceArns []string `locationName:"resourceArns" type:"list"`

	// The Amazon Resource Name (ARN) of the resource share.
	//
	// ResourceShareArn is a required field
	ResourceShareArn *string `locationName:"resourceShareArn" type:"string" required:"true"`
}

// String returns the string representation
func (s AssociateResourceShareInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AssociateResourceShareInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "AssociateResourceShareInput"}

	if s.ResourceShareArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("ResourceShareArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s AssociateResourceShareInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	if s.ClientToken != nil {
		v := *s.ClientToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "clientToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
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
	if s.ResourceShareArn != nil {
		v := *s.ResourceShareArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "resourceShareArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ram-2018-01-04/AssociateResourceShareResponse
type AssociateResourceShareOutput struct {
	_ struct{} `type:"structure"`

	// A unique, case-sensitive identifier that you provide to ensure the idempotency
	// of the request.
	ClientToken *string `locationName:"clientToken" type:"string"`

	// Information about the associations.
	ResourceShareAssociations []ResourceShareAssociation `locationName:"resourceShareAssociations" type:"list"`
}

// String returns the string representation
func (s AssociateResourceShareOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s AssociateResourceShareOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.ClientToken != nil {
		v := *s.ClientToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "clientToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.ResourceShareAssociations) > 0 {
		v := s.ResourceShareAssociations

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "resourceShareAssociations", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	return nil
}

const opAssociateResourceShare = "AssociateResourceShare"

// AssociateResourceShareRequest returns a request value for making API operation for
// AWS Resource Access Manager.
//
// Associates the specified resource share with the specified principals and
// resources.
//
//    // Example sending a request using AssociateResourceShareRequest.
//    req := client.AssociateResourceShareRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ram-2018-01-04/AssociateResourceShare
func (c *Client) AssociateResourceShareRequest(input *AssociateResourceShareInput) AssociateResourceShareRequest {
	op := &aws.Operation{
		Name:       opAssociateResourceShare,
		HTTPMethod: "POST",
		HTTPPath:   "/associateresourceshare",
	}

	if input == nil {
		input = &AssociateResourceShareInput{}
	}

	req := c.newRequest(op, input, &AssociateResourceShareOutput{})
	return AssociateResourceShareRequest{Request: req, Input: input, Copy: c.AssociateResourceShareRequest}
}

// AssociateResourceShareRequest is the request type for the
// AssociateResourceShare API operation.
type AssociateResourceShareRequest struct {
	*aws.Request
	Input *AssociateResourceShareInput
	Copy  func(*AssociateResourceShareInput) AssociateResourceShareRequest
}

// Send marshals and sends the AssociateResourceShare API request.
func (r AssociateResourceShareRequest) Send(ctx context.Context) (*AssociateResourceShareResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &AssociateResourceShareResponse{
		AssociateResourceShareOutput: r.Request.Data.(*AssociateResourceShareOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// AssociateResourceShareResponse is the response type for the
// AssociateResourceShare API operation.
type AssociateResourceShareResponse struct {
	*AssociateResourceShareOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// AssociateResourceShare request.
func (r *AssociateResourceShareResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
