// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package pinpoint

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-2016-12-01/DeleteEmailChannelRequest
type DeleteEmailChannelInput struct {
	_ struct{} `type:"structure"`

	// ApplicationId is a required field
	ApplicationId *string `location:"uri" locationName:"application-id" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteEmailChannelInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteEmailChannelInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteEmailChannelInput"}

	if s.ApplicationId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ApplicationId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteEmailChannelInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	if s.ApplicationId != nil {
		v := *s.ApplicationId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "application-id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-2016-12-01/DeleteEmailChannelResponse
type DeleteEmailChannelOutput struct {
	_ struct{} `type:"structure" payload:"EmailChannelResponse"`

	// Email Channel Response.
	//
	// EmailChannelResponse is a required field
	EmailChannelResponse *EmailChannelResponse `type:"structure" required:"true"`
}

// String returns the string representation
func (s DeleteEmailChannelOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteEmailChannelOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.EmailChannelResponse != nil {
		v := s.EmailChannelResponse

		metadata := protocol.Metadata{}
		e.SetFields(protocol.PayloadTarget, "EmailChannelResponse", v, metadata)
	}
	return nil
}

const opDeleteEmailChannel = "DeleteEmailChannel"

// DeleteEmailChannelRequest returns a request value for making API operation for
// Amazon Pinpoint.
//
// Delete an email channel.
//
//    // Example sending a request using DeleteEmailChannelRequest.
//    req := client.DeleteEmailChannelRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-2016-12-01/DeleteEmailChannel
func (c *Client) DeleteEmailChannelRequest(input *DeleteEmailChannelInput) DeleteEmailChannelRequest {
	op := &aws.Operation{
		Name:       opDeleteEmailChannel,
		HTTPMethod: "DELETE",
		HTTPPath:   "/v1/apps/{application-id}/channels/email",
	}

	if input == nil {
		input = &DeleteEmailChannelInput{}
	}

	req := c.newRequest(op, input, &DeleteEmailChannelOutput{})
	return DeleteEmailChannelRequest{Request: req, Input: input, Copy: c.DeleteEmailChannelRequest}
}

// DeleteEmailChannelRequest is the request type for the
// DeleteEmailChannel API operation.
type DeleteEmailChannelRequest struct {
	*aws.Request
	Input *DeleteEmailChannelInput
	Copy  func(*DeleteEmailChannelInput) DeleteEmailChannelRequest
}

// Send marshals and sends the DeleteEmailChannel API request.
func (r DeleteEmailChannelRequest) Send(ctx context.Context) (*DeleteEmailChannelResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteEmailChannelResponse{
		DeleteEmailChannelOutput: r.Request.Data.(*DeleteEmailChannelOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteEmailChannelResponse is the response type for the
// DeleteEmailChannel API operation.
type DeleteEmailChannelResponse struct {
	*DeleteEmailChannelOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteEmailChannel request.
func (r *DeleteEmailChannelResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
