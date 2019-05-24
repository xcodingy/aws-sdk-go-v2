// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package securityhub

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/securityhub-2018-10-26/DeleteInvitationsRequest
type DeleteInvitationsInput struct {
	_ struct{} `type:"structure"`

	// A list of account IDs specifying accounts whose invitations to Security Hub
	// you want to delete.
	AccountIds []string `type:"list"`
}

// String returns the string representation
func (s DeleteInvitationsInput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteInvitationsInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	if len(s.AccountIds) > 0 {
		v := s.AccountIds

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "AccountIds", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/securityhub-2018-10-26/DeleteInvitationsResponse
type DeleteInvitationsOutput struct {
	_ struct{} `type:"structure"`

	// A list of account ID and email address pairs of the AWS accounts that could
	// not be processed.
	UnprocessedAccounts []Result `type:"list"`
}

// String returns the string representation
func (s DeleteInvitationsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteInvitationsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if len(s.UnprocessedAccounts) > 0 {
		v := s.UnprocessedAccounts

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "UnprocessedAccounts", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	return nil
}

const opDeleteInvitations = "DeleteInvitations"

// DeleteInvitationsRequest returns a request value for making API operation for
// AWS SecurityHub.
//
// Deletes invitations that are sent to this AWS account (invitee) by the AWS
// accounts (inviters) that are specified by their account IDs.
//
//    // Example sending a request using DeleteInvitationsRequest.
//    req := client.DeleteInvitationsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/securityhub-2018-10-26/DeleteInvitations
func (c *Client) DeleteInvitationsRequest(input *DeleteInvitationsInput) DeleteInvitationsRequest {
	op := &aws.Operation{
		Name:       opDeleteInvitations,
		HTTPMethod: "POST",
		HTTPPath:   "/invitations/delete",
	}

	if input == nil {
		input = &DeleteInvitationsInput{}
	}

	req := c.newRequest(op, input, &DeleteInvitationsOutput{})
	return DeleteInvitationsRequest{Request: req, Input: input, Copy: c.DeleteInvitationsRequest}
}

// DeleteInvitationsRequest is the request type for the
// DeleteInvitations API operation.
type DeleteInvitationsRequest struct {
	*aws.Request
	Input *DeleteInvitationsInput
	Copy  func(*DeleteInvitationsInput) DeleteInvitationsRequest
}

// Send marshals and sends the DeleteInvitations API request.
func (r DeleteInvitationsRequest) Send(ctx context.Context) (*DeleteInvitationsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteInvitationsResponse{
		DeleteInvitationsOutput: r.Request.Data.(*DeleteInvitationsOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteInvitationsResponse is the response type for the
// DeleteInvitations API operation.
type DeleteInvitationsResponse struct {
	*DeleteInvitationsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteInvitations request.
func (r *DeleteInvitationsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
