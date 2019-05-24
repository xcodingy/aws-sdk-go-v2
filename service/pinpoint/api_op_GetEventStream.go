// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package pinpoint

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-2016-12-01/GetEventStreamRequest
type GetEventStreamInput struct {
	_ struct{} `type:"structure"`

	// ApplicationId is a required field
	ApplicationId *string `location:"uri" locationName:"application-id" type:"string" required:"true"`
}

// String returns the string representation
func (s GetEventStreamInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetEventStreamInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetEventStreamInput"}

	if s.ApplicationId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ApplicationId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetEventStreamInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	if s.ApplicationId != nil {
		v := *s.ApplicationId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "application-id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-2016-12-01/GetEventStreamResponse
type GetEventStreamOutput struct {
	_ struct{} `type:"structure" payload:"EventStream"`

	// Model for an event publishing subscription export.
	//
	// EventStream is a required field
	EventStream *EventStream `type:"structure" required:"true"`
}

// String returns the string representation
func (s GetEventStreamOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetEventStreamOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.EventStream != nil {
		v := s.EventStream

		metadata := protocol.Metadata{}
		e.SetFields(protocol.PayloadTarget, "EventStream", v, metadata)
	}
	return nil
}

const opGetEventStream = "GetEventStream"

// GetEventStreamRequest returns a request value for making API operation for
// Amazon Pinpoint.
//
// Returns the event stream for an app.
//
//    // Example sending a request using GetEventStreamRequest.
//    req := client.GetEventStreamRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-2016-12-01/GetEventStream
func (c *Client) GetEventStreamRequest(input *GetEventStreamInput) GetEventStreamRequest {
	op := &aws.Operation{
		Name:       opGetEventStream,
		HTTPMethod: "GET",
		HTTPPath:   "/v1/apps/{application-id}/eventstream",
	}

	if input == nil {
		input = &GetEventStreamInput{}
	}

	req := c.newRequest(op, input, &GetEventStreamOutput{})
	return GetEventStreamRequest{Request: req, Input: input, Copy: c.GetEventStreamRequest}
}

// GetEventStreamRequest is the request type for the
// GetEventStream API operation.
type GetEventStreamRequest struct {
	*aws.Request
	Input *GetEventStreamInput
	Copy  func(*GetEventStreamInput) GetEventStreamRequest
}

// Send marshals and sends the GetEventStream API request.
func (r GetEventStreamRequest) Send(ctx context.Context) (*GetEventStreamResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetEventStreamResponse{
		GetEventStreamOutput: r.Request.Data.(*GetEventStreamOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetEventStreamResponse is the response type for the
// GetEventStream API operation.
type GetEventStreamResponse struct {
	*GetEventStreamOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetEventStream request.
func (r *GetEventStreamResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
