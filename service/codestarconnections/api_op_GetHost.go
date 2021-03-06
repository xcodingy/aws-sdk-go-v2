// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codestarconnections

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GetHostInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the requested host.
	//
	// HostArn is a required field
	HostArn *string `type:"string" required:"true"`
}

// String returns the string representation
func (s GetHostInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetHostInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetHostInput"}

	if s.HostArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("HostArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type GetHostOutput struct {
	_ struct{} `type:"structure"`

	// The name of the requested host.
	Name *string `min:"1" type:"string"`

	// The endpoint of the infrastructure represented by the requested host.
	ProviderEndpoint *string `min:"1" type:"string"`

	// The provider type of the requested host, such as GitHub Enterprise Server.
	ProviderType ProviderType `type:"string" enum:"true"`

	// The status of the requested host.
	Status *string `type:"string"`

	// The VPC configuration of the requested host.
	VpcConfiguration *VpcConfiguration `type:"structure"`
}

// String returns the string representation
func (s GetHostOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetHost = "GetHost"

// GetHostRequest returns a request value for making API operation for
// AWS CodeStar connections.
//
// Returns the host ARN and details such as status, provider type, endpoint,
// and, if applicable, the VPC configuration.
//
//    // Example sending a request using GetHostRequest.
//    req := client.GetHostRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codestar-connections-2019-12-01/GetHost
func (c *Client) GetHostRequest(input *GetHostInput) GetHostRequest {
	op := &aws.Operation{
		Name:       opGetHost,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetHostInput{}
	}

	req := c.newRequest(op, input, &GetHostOutput{})

	return GetHostRequest{Request: req, Input: input, Copy: c.GetHostRequest}
}

// GetHostRequest is the request type for the
// GetHost API operation.
type GetHostRequest struct {
	*aws.Request
	Input *GetHostInput
	Copy  func(*GetHostInput) GetHostRequest
}

// Send marshals and sends the GetHost API request.
func (r GetHostRequest) Send(ctx context.Context) (*GetHostResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetHostResponse{
		GetHostOutput: r.Request.Data.(*GetHostOutput),
		response:      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetHostResponse is the response type for the
// GetHost API operation.
type GetHostResponse struct {
	*GetHostOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetHost request.
func (r *GetHostResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
