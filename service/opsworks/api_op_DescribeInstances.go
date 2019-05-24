// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package opsworks

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/opsworks-2013-02-18/DescribeInstancesRequest
type DescribeInstancesInput struct {
	_ struct{} `type:"structure"`

	// An array of instance IDs to be described. If you use this parameter, DescribeInstances
	// returns a description of the specified instances. Otherwise, it returns a
	// description of every instance.
	InstanceIds []string `type:"list"`

	// A layer ID. If you use this parameter, DescribeInstances returns descriptions
	// of the instances associated with the specified layer.
	LayerId *string `type:"string"`

	// A stack ID. If you use this parameter, DescribeInstances returns descriptions
	// of the instances associated with the specified stack.
	StackId *string `type:"string"`
}

// String returns the string representation
func (s DescribeInstancesInput) String() string {
	return awsutil.Prettify(s)
}

// Contains the response to a DescribeInstances request.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/opsworks-2013-02-18/DescribeInstancesResult
type DescribeInstancesOutput struct {
	_ struct{} `type:"structure"`

	// An array of Instance objects that describe the instances.
	Instances []Instance `type:"list"`
}

// String returns the string representation
func (s DescribeInstancesOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeInstances = "DescribeInstances"

// DescribeInstancesRequest returns a request value for making API operation for
// AWS OpsWorks.
//
// Requests a description of a set of instances.
//
// This call accepts only one resource-identifying parameter.
//
// Required Permissions: To use this action, an IAM user must have a Show, Deploy,
// or Manage permissions level for the stack, or an attached policy that explicitly
// grants permissions. For more information about user permissions, see Managing
// User Permissions (http://docs.aws.amazon.com/opsworks/latest/userguide/opsworks-security-users.html).
//
//    // Example sending a request using DescribeInstancesRequest.
//    req := client.DescribeInstancesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/opsworks-2013-02-18/DescribeInstances
func (c *Client) DescribeInstancesRequest(input *DescribeInstancesInput) DescribeInstancesRequest {
	op := &aws.Operation{
		Name:       opDescribeInstances,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeInstancesInput{}
	}

	req := c.newRequest(op, input, &DescribeInstancesOutput{})
	return DescribeInstancesRequest{Request: req, Input: input, Copy: c.DescribeInstancesRequest}
}

// DescribeInstancesRequest is the request type for the
// DescribeInstances API operation.
type DescribeInstancesRequest struct {
	*aws.Request
	Input *DescribeInstancesInput
	Copy  func(*DescribeInstancesInput) DescribeInstancesRequest
}

// Send marshals and sends the DescribeInstances API request.
func (r DescribeInstancesRequest) Send(ctx context.Context) (*DescribeInstancesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeInstancesResponse{
		DescribeInstancesOutput: r.Request.Data.(*DescribeInstancesOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeInstancesResponse is the response type for the
// DescribeInstances API operation.
type DescribeInstancesResponse struct {
	*DescribeInstancesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeInstances request.
func (r *DescribeInstancesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
