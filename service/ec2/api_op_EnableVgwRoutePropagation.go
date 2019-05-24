// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/ec2query"
)

// Contains the parameters for EnableVgwRoutePropagation.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/EnableVgwRoutePropagationRequest
type EnableVgwRoutePropagationInput struct {
	_ struct{} `type:"structure"`

	// The ID of the virtual private gateway.
	//
	// GatewayId is a required field
	GatewayId *string `type:"string" required:"true"`

	// The ID of the route table.
	//
	// RouteTableId is a required field
	RouteTableId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s EnableVgwRoutePropagationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *EnableVgwRoutePropagationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "EnableVgwRoutePropagationInput"}

	if s.GatewayId == nil {
		invalidParams.Add(aws.NewErrParamRequired("GatewayId"))
	}

	if s.RouteTableId == nil {
		invalidParams.Add(aws.NewErrParamRequired("RouteTableId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/EnableVgwRoutePropagationOutput
type EnableVgwRoutePropagationOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s EnableVgwRoutePropagationOutput) String() string {
	return awsutil.Prettify(s)
}

const opEnableVgwRoutePropagation = "EnableVgwRoutePropagation"

// EnableVgwRoutePropagationRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Enables a virtual private gateway (VGW) to propagate routes to the specified
// route table of a VPC.
//
//    // Example sending a request using EnableVgwRoutePropagationRequest.
//    req := client.EnableVgwRoutePropagationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/EnableVgwRoutePropagation
func (c *Client) EnableVgwRoutePropagationRequest(input *EnableVgwRoutePropagationInput) EnableVgwRoutePropagationRequest {
	op := &aws.Operation{
		Name:       opEnableVgwRoutePropagation,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &EnableVgwRoutePropagationInput{}
	}

	req := c.newRequest(op, input, &EnableVgwRoutePropagationOutput{})
	req.Handlers.Unmarshal.Remove(ec2query.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return EnableVgwRoutePropagationRequest{Request: req, Input: input, Copy: c.EnableVgwRoutePropagationRequest}
}

// EnableVgwRoutePropagationRequest is the request type for the
// EnableVgwRoutePropagation API operation.
type EnableVgwRoutePropagationRequest struct {
	*aws.Request
	Input *EnableVgwRoutePropagationInput
	Copy  func(*EnableVgwRoutePropagationInput) EnableVgwRoutePropagationRequest
}

// Send marshals and sends the EnableVgwRoutePropagation API request.
func (r EnableVgwRoutePropagationRequest) Send(ctx context.Context) (*EnableVgwRoutePropagationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &EnableVgwRoutePropagationResponse{
		EnableVgwRoutePropagationOutput: r.Request.Data.(*EnableVgwRoutePropagationOutput),
		response:                        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// EnableVgwRoutePropagationResponse is the response type for the
// EnableVgwRoutePropagation API operation.
type EnableVgwRoutePropagationResponse struct {
	*EnableVgwRoutePropagationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// EnableVgwRoutePropagation request.
func (r *EnableVgwRoutePropagationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
