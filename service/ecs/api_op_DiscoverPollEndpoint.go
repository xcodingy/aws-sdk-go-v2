// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ecs

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ecs-2014-11-13/DiscoverPollEndpointRequest
type DiscoverPollEndpointInput struct {
	_ struct{} `type:"structure"`

	// The short name or full Amazon Resource Name (ARN) of the cluster to which
	// the container instance belongs.
	Cluster *string `locationName:"cluster" type:"string"`

	// The container instance ID or full ARN of the container instance. The ARN
	// contains the arn:aws:ecs namespace, followed by the Region of the container
	// instance, the AWS account ID of the container instance owner, the container-instance
	// namespace, and then the container instance ID. For example, arn:aws:ecs:region:aws_account_id:container-instance/container_instance_ID.
	ContainerInstance *string `locationName:"containerInstance" type:"string"`
}

// String returns the string representation
func (s DiscoverPollEndpointInput) String() string {
	return awsutil.Prettify(s)
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ecs-2014-11-13/DiscoverPollEndpointResponse
type DiscoverPollEndpointOutput struct {
	_ struct{} `type:"structure"`

	// The endpoint for the Amazon ECS agent to poll.
	Endpoint *string `locationName:"endpoint" type:"string"`

	// The telemetry endpoint for the Amazon ECS agent.
	TelemetryEndpoint *string `locationName:"telemetryEndpoint" type:"string"`
}

// String returns the string representation
func (s DiscoverPollEndpointOutput) String() string {
	return awsutil.Prettify(s)
}

const opDiscoverPollEndpoint = "DiscoverPollEndpoint"

// DiscoverPollEndpointRequest returns a request value for making API operation for
// Amazon EC2 Container Service.
//
// This action is only used by the Amazon ECS agent, and it is not intended
// for use outside of the agent.
//
// Returns an endpoint for the Amazon ECS agent to poll for updates.
//
//    // Example sending a request using DiscoverPollEndpointRequest.
//    req := client.DiscoverPollEndpointRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ecs-2014-11-13/DiscoverPollEndpoint
func (c *Client) DiscoverPollEndpointRequest(input *DiscoverPollEndpointInput) DiscoverPollEndpointRequest {
	op := &aws.Operation{
		Name:       opDiscoverPollEndpoint,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DiscoverPollEndpointInput{}
	}

	req := c.newRequest(op, input, &DiscoverPollEndpointOutput{})
	return DiscoverPollEndpointRequest{Request: req, Input: input, Copy: c.DiscoverPollEndpointRequest}
}

// DiscoverPollEndpointRequest is the request type for the
// DiscoverPollEndpoint API operation.
type DiscoverPollEndpointRequest struct {
	*aws.Request
	Input *DiscoverPollEndpointInput
	Copy  func(*DiscoverPollEndpointInput) DiscoverPollEndpointRequest
}

// Send marshals and sends the DiscoverPollEndpoint API request.
func (r DiscoverPollEndpointRequest) Send(ctx context.Context) (*DiscoverPollEndpointResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DiscoverPollEndpointResponse{
		DiscoverPollEndpointOutput: r.Request.Data.(*DiscoverPollEndpointOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DiscoverPollEndpointResponse is the response type for the
// DiscoverPollEndpoint API operation.
type DiscoverPollEndpointResponse struct {
	*DiscoverPollEndpointOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DiscoverPollEndpoint request.
func (r *DiscoverPollEndpointResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
