// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ram

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ram-2018-01-04/RejectResourceShareInvitationRequest
type RejectResourceShareInvitationInput struct {
	_ struct{} `type:"structure"`

	// A unique, case-sensitive identifier that you provide to ensure the idempotency
	// of the request.
	ClientToken *string `locationName:"clientToken" type:"string"`

	// The Amazon Resource Name (ARN) of the invitation.
	//
	// ResourceShareInvitationArn is a required field
	ResourceShareInvitationArn *string `locationName:"resourceShareInvitationArn" type:"string" required:"true"`
}

// String returns the string representation
func (s RejectResourceShareInvitationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *RejectResourceShareInvitationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "RejectResourceShareInvitationInput"}

	if s.ResourceShareInvitationArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("ResourceShareInvitationArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s RejectResourceShareInvitationInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	if s.ClientToken != nil {
		v := *s.ClientToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "clientToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ResourceShareInvitationArn != nil {
		v := *s.ResourceShareInvitationArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "resourceShareInvitationArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ram-2018-01-04/RejectResourceShareInvitationResponse
type RejectResourceShareInvitationOutput struct {
	_ struct{} `type:"structure"`

	// A unique, case-sensitive identifier that you provide to ensure the idempotency
	// of the request.
	ClientToken *string `locationName:"clientToken" type:"string"`

	// Information about the invitation.
	ResourceShareInvitation *ResourceShareInvitation `locationName:"resourceShareInvitation" type:"structure"`
}

// String returns the string representation
func (s RejectResourceShareInvitationOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s RejectResourceShareInvitationOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.ClientToken != nil {
		v := *s.ClientToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "clientToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ResourceShareInvitation != nil {
		v := s.ResourceShareInvitation

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "resourceShareInvitation", v, metadata)
	}
	return nil
}

const opRejectResourceShareInvitation = "RejectResourceShareInvitation"

// RejectResourceShareInvitationRequest returns a request value for making API operation for
// AWS Resource Access Manager.
//
// Rejects an invitation to a resource share from another AWS account.
//
//    // Example sending a request using RejectResourceShareInvitationRequest.
//    req := client.RejectResourceShareInvitationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ram-2018-01-04/RejectResourceShareInvitation
func (c *Client) RejectResourceShareInvitationRequest(input *RejectResourceShareInvitationInput) RejectResourceShareInvitationRequest {
	op := &aws.Operation{
		Name:       opRejectResourceShareInvitation,
		HTTPMethod: "POST",
		HTTPPath:   "/rejectresourceshareinvitation",
	}

	if input == nil {
		input = &RejectResourceShareInvitationInput{}
	}

	req := c.newRequest(op, input, &RejectResourceShareInvitationOutput{})
	return RejectResourceShareInvitationRequest{Request: req, Input: input, Copy: c.RejectResourceShareInvitationRequest}
}

// RejectResourceShareInvitationRequest is the request type for the
// RejectResourceShareInvitation API operation.
type RejectResourceShareInvitationRequest struct {
	*aws.Request
	Input *RejectResourceShareInvitationInput
	Copy  func(*RejectResourceShareInvitationInput) RejectResourceShareInvitationRequest
}

// Send marshals and sends the RejectResourceShareInvitation API request.
func (r RejectResourceShareInvitationRequest) Send(ctx context.Context) (*RejectResourceShareInvitationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &RejectResourceShareInvitationResponse{
		RejectResourceShareInvitationOutput: r.Request.Data.(*RejectResourceShareInvitationOutput),
		response:                            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// RejectResourceShareInvitationResponse is the response type for the
// RejectResourceShareInvitation API operation.
type RejectResourceShareInvitationResponse struct {
	*RejectResourceShareInvitationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// RejectResourceShareInvitation request.
func (r *RejectResourceShareInvitationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
