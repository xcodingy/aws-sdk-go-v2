// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package guardduty

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/guardduty-2017-11-28/GetInvitationsCountRequest
type GetInvitationsCountInput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s GetInvitationsCountInput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetInvitationsCountInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	return nil
}

// GetInvitationsCount response object.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/guardduty-2017-11-28/GetInvitationsCountResponse
type GetInvitationsCountOutput struct {
	_ struct{} `type:"structure"`

	// The number of received invitations.
	InvitationsCount *int64 `locationName:"invitationsCount" type:"integer"`
}

// String returns the string representation
func (s GetInvitationsCountOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetInvitationsCountOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.InvitationsCount != nil {
		v := *s.InvitationsCount

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "invitationsCount", protocol.Int64Value(v), metadata)
	}
	return nil
}

const opGetInvitationsCount = "GetInvitationsCount"

// GetInvitationsCountRequest returns a request value for making API operation for
// Amazon GuardDuty.
//
// Returns the count of all GuardDuty membership invitations that were sent
// to the current member account except the currently accepted invitation.
//
//    // Example sending a request using GetInvitationsCountRequest.
//    req := client.GetInvitationsCountRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/guardduty-2017-11-28/GetInvitationsCount
func (c *Client) GetInvitationsCountRequest(input *GetInvitationsCountInput) GetInvitationsCountRequest {
	op := &aws.Operation{
		Name:       opGetInvitationsCount,
		HTTPMethod: "GET",
		HTTPPath:   "/invitation/count",
	}

	if input == nil {
		input = &GetInvitationsCountInput{}
	}

	req := c.newRequest(op, input, &GetInvitationsCountOutput{})
	return GetInvitationsCountRequest{Request: req, Input: input, Copy: c.GetInvitationsCountRequest}
}

// GetInvitationsCountRequest is the request type for the
// GetInvitationsCount API operation.
type GetInvitationsCountRequest struct {
	*aws.Request
	Input *GetInvitationsCountInput
	Copy  func(*GetInvitationsCountInput) GetInvitationsCountRequest
}

// Send marshals and sends the GetInvitationsCount API request.
func (r GetInvitationsCountRequest) Send(ctx context.Context) (*GetInvitationsCountResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetInvitationsCountResponse{
		GetInvitationsCountOutput: r.Request.Data.(*GetInvitationsCountOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetInvitationsCountResponse is the response type for the
// GetInvitationsCount API operation.
type GetInvitationsCountResponse struct {
	*GetInvitationsCountOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetInvitationsCount request.
func (r *GetInvitationsCountResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
