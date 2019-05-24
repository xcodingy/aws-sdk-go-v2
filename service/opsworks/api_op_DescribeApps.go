// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package opsworks

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/opsworks-2013-02-18/DescribeAppsRequest
type DescribeAppsInput struct {
	_ struct{} `type:"structure"`

	// An array of app IDs for the apps to be described. If you use this parameter,
	// DescribeApps returns a description of the specified apps. Otherwise, it returns
	// a description of every app.
	AppIds []string `type:"list"`

	// The app stack ID. If you use this parameter, DescribeApps returns a description
	// of the apps in the specified stack.
	StackId *string `type:"string"`
}

// String returns the string representation
func (s DescribeAppsInput) String() string {
	return awsutil.Prettify(s)
}

// Contains the response to a DescribeApps request.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/opsworks-2013-02-18/DescribeAppsResult
type DescribeAppsOutput struct {
	_ struct{} `type:"structure"`

	// An array of App objects that describe the specified apps.
	Apps []App `type:"list"`
}

// String returns the string representation
func (s DescribeAppsOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeApps = "DescribeApps"

// DescribeAppsRequest returns a request value for making API operation for
// AWS OpsWorks.
//
// Requests a description of a specified set of apps.
//
// This call accepts only one resource-identifying parameter.
//
// Required Permissions: To use this action, an IAM user must have a Show, Deploy,
// or Manage permissions level for the stack, or an attached policy that explicitly
// grants permissions. For more information about user permissions, see Managing
// User Permissions (http://docs.aws.amazon.com/opsworks/latest/userguide/opsworks-security-users.html).
//
//    // Example sending a request using DescribeAppsRequest.
//    req := client.DescribeAppsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/opsworks-2013-02-18/DescribeApps
func (c *Client) DescribeAppsRequest(input *DescribeAppsInput) DescribeAppsRequest {
	op := &aws.Operation{
		Name:       opDescribeApps,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeAppsInput{}
	}

	req := c.newRequest(op, input, &DescribeAppsOutput{})
	return DescribeAppsRequest{Request: req, Input: input, Copy: c.DescribeAppsRequest}
}

// DescribeAppsRequest is the request type for the
// DescribeApps API operation.
type DescribeAppsRequest struct {
	*aws.Request
	Input *DescribeAppsInput
	Copy  func(*DescribeAppsInput) DescribeAppsRequest
}

// Send marshals and sends the DescribeApps API request.
func (r DescribeAppsRequest) Send(ctx context.Context) (*DescribeAppsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeAppsResponse{
		DescribeAppsOutput: r.Request.Data.(*DescribeAppsOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeAppsResponse is the response type for the
// DescribeApps API operation.
type DescribeAppsResponse struct {
	*DescribeAppsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeApps request.
func (r *DescribeAppsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
