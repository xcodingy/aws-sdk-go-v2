// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package securityhub

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/securityhub-2018-10-26/AcceptInvitationRequest
type AcceptInvitationInput struct {
	_ struct{} `type:"structure"`

	// The ID of the invitation that is sent to the AWS account by the Security
	// Hub master account.
	InvitationId *string `type:"string"`

	// The account ID of the master Security Hub account whose invitation you're
	// accepting.
	MasterId *string `type:"string"`
}

// String returns the string representation
func (s AcceptInvitationInput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s AcceptInvitationInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	if s.InvitationId != nil {
		v := *s.InvitationId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "InvitationId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.MasterId != nil {
		v := *s.MasterId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "MasterId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/securityhub-2018-10-26/AcceptInvitationResponse
type AcceptInvitationOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s AcceptInvitationOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s AcceptInvitationOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opAcceptInvitation = "AcceptInvitation"

// AcceptInvitationRequest returns a request value for making API operation for
// AWS SecurityHub.
//
// Accepts the invitation to be monitored by a master SecurityHub account.
//
//    // Example sending a request using AcceptInvitationRequest.
//    req := client.AcceptInvitationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/securityhub-2018-10-26/AcceptInvitation
func (c *Client) AcceptInvitationRequest(input *AcceptInvitationInput) AcceptInvitationRequest {
	op := &aws.Operation{
		Name:       opAcceptInvitation,
		HTTPMethod: "POST",
		HTTPPath:   "/master",
	}

	if input == nil {
		input = &AcceptInvitationInput{}
	}

	req := c.newRequest(op, input, &AcceptInvitationOutput{})
	return AcceptInvitationRequest{Request: req, Input: input, Copy: c.AcceptInvitationRequest}
}

// AcceptInvitationRequest is the request type for the
// AcceptInvitation API operation.
type AcceptInvitationRequest struct {
	*aws.Request
	Input *AcceptInvitationInput
	Copy  func(*AcceptInvitationInput) AcceptInvitationRequest
}

// Send marshals and sends the AcceptInvitation API request.
func (r AcceptInvitationRequest) Send(ctx context.Context) (*AcceptInvitationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &AcceptInvitationResponse{
		AcceptInvitationOutput: r.Request.Data.(*AcceptInvitationOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// AcceptInvitationResponse is the response type for the
// AcceptInvitation API operation.
type AcceptInvitationResponse struct {
	*AcceptInvitationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// AcceptInvitation request.
func (r *AcceptInvitationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
