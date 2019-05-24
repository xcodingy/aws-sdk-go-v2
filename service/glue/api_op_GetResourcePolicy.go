// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package glue

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/glue-2017-03-31/GetResourcePolicyRequest
type GetResourcePolicyInput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s GetResourcePolicyInput) String() string {
	return awsutil.Prettify(s)
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/glue-2017-03-31/GetResourcePolicyResponse
type GetResourcePolicyOutput struct {
	_ struct{} `type:"structure"`

	// The date and time at which the policy was created.
	CreateTime *time.Time `type:"timestamp" timestampFormat:"unix"`

	// Contains the hash value associated with this policy.
	PolicyHash *string `min:"1" type:"string"`

	// Contains the requested policy document, in JSON format.
	PolicyInJson *string `min:"2" type:"string"`

	// The date and time at which the policy was last updated.
	UpdateTime *time.Time `type:"timestamp" timestampFormat:"unix"`
}

// String returns the string representation
func (s GetResourcePolicyOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetResourcePolicy = "GetResourcePolicy"

// GetResourcePolicyRequest returns a request value for making API operation for
// AWS Glue.
//
// Retrieves a specified resource policy.
//
//    // Example sending a request using GetResourcePolicyRequest.
//    req := client.GetResourcePolicyRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/glue-2017-03-31/GetResourcePolicy
func (c *Client) GetResourcePolicyRequest(input *GetResourcePolicyInput) GetResourcePolicyRequest {
	op := &aws.Operation{
		Name:       opGetResourcePolicy,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetResourcePolicyInput{}
	}

	req := c.newRequest(op, input, &GetResourcePolicyOutput{})
	return GetResourcePolicyRequest{Request: req, Input: input, Copy: c.GetResourcePolicyRequest}
}

// GetResourcePolicyRequest is the request type for the
// GetResourcePolicy API operation.
type GetResourcePolicyRequest struct {
	*aws.Request
	Input *GetResourcePolicyInput
	Copy  func(*GetResourcePolicyInput) GetResourcePolicyRequest
}

// Send marshals and sends the GetResourcePolicy API request.
func (r GetResourcePolicyRequest) Send(ctx context.Context) (*GetResourcePolicyResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetResourcePolicyResponse{
		GetResourcePolicyOutput: r.Request.Data.(*GetResourcePolicyOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetResourcePolicyResponse is the response type for the
// GetResourcePolicy API operation.
type GetResourcePolicyResponse struct {
	*GetResourcePolicyOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetResourcePolicy request.
func (r *GetResourcePolicyResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
