// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package mediaconnect

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/mediaconnect-2018-11-14/StopFlowRequest
type StopFlowInput struct {
	_ struct{} `type:"structure"`

	// FlowArn is a required field
	FlowArn *string `location:"uri" locationName:"flowArn" type:"string" required:"true"`
}

// String returns the string representation
func (s StopFlowInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *StopFlowInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "StopFlowInput"}

	if s.FlowArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("FlowArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s StopFlowInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	if s.FlowArn != nil {
		v := *s.FlowArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "flowArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// The result of a successful StopFlow request.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/mediaconnect-2018-11-14/StopFlowResponse
type StopFlowOutput struct {
	_ struct{} `type:"structure"`

	// The ARN of the flow that you stopped.
	FlowArn *string `locationName:"flowArn" type:"string"`

	// The status of the flow when the StopFlow process begins.
	Status Status `locationName:"status" type:"string" enum:"true"`
}

// String returns the string representation
func (s StopFlowOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s StopFlowOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.FlowArn != nil {
		v := *s.FlowArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "flowArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.Status) > 0 {
		v := s.Status

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "status", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	return nil
}

const opStopFlow = "StopFlow"

// StopFlowRequest returns a request value for making API operation for
// AWS MediaConnect.
//
// Stops a flow.
//
//    // Example sending a request using StopFlowRequest.
//    req := client.StopFlowRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/mediaconnect-2018-11-14/StopFlow
func (c *Client) StopFlowRequest(input *StopFlowInput) StopFlowRequest {
	op := &aws.Operation{
		Name:       opStopFlow,
		HTTPMethod: "POST",
		HTTPPath:   "/v1/flows/stop/{flowArn}",
	}

	if input == nil {
		input = &StopFlowInput{}
	}

	req := c.newRequest(op, input, &StopFlowOutput{})
	return StopFlowRequest{Request: req, Input: input, Copy: c.StopFlowRequest}
}

// StopFlowRequest is the request type for the
// StopFlow API operation.
type StopFlowRequest struct {
	*aws.Request
	Input *StopFlowInput
	Copy  func(*StopFlowInput) StopFlowRequest
}

// Send marshals and sends the StopFlow API request.
func (r StopFlowRequest) Send(ctx context.Context) (*StopFlowResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &StopFlowResponse{
		StopFlowOutput: r.Request.Data.(*StopFlowOutput),
		response:       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// StopFlowResponse is the response type for the
// StopFlow API operation.
type StopFlowResponse struct {
	*StopFlowOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// StopFlow request.
func (r *StopFlowResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
