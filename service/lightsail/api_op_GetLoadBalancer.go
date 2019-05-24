// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package lightsail

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/lightsail-2016-11-28/GetLoadBalancerRequest
type GetLoadBalancerInput struct {
	_ struct{} `type:"structure"`

	// The name of the load balancer.
	//
	// LoadBalancerName is a required field
	LoadBalancerName *string `locationName:"loadBalancerName" type:"string" required:"true"`
}

// String returns the string representation
func (s GetLoadBalancerInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetLoadBalancerInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetLoadBalancerInput"}

	if s.LoadBalancerName == nil {
		invalidParams.Add(aws.NewErrParamRequired("LoadBalancerName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/lightsail-2016-11-28/GetLoadBalancerResult
type GetLoadBalancerOutput struct {
	_ struct{} `type:"structure"`

	// An object containing information about your load balancer.
	LoadBalancer *LoadBalancer `locationName:"loadBalancer" type:"structure"`
}

// String returns the string representation
func (s GetLoadBalancerOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetLoadBalancer = "GetLoadBalancer"

// GetLoadBalancerRequest returns a request value for making API operation for
// Amazon Lightsail.
//
// Returns information about the specified Lightsail load balancer.
//
//    // Example sending a request using GetLoadBalancerRequest.
//    req := client.GetLoadBalancerRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/lightsail-2016-11-28/GetLoadBalancer
func (c *Client) GetLoadBalancerRequest(input *GetLoadBalancerInput) GetLoadBalancerRequest {
	op := &aws.Operation{
		Name:       opGetLoadBalancer,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetLoadBalancerInput{}
	}

	req := c.newRequest(op, input, &GetLoadBalancerOutput{})
	return GetLoadBalancerRequest{Request: req, Input: input, Copy: c.GetLoadBalancerRequest}
}

// GetLoadBalancerRequest is the request type for the
// GetLoadBalancer API operation.
type GetLoadBalancerRequest struct {
	*aws.Request
	Input *GetLoadBalancerInput
	Copy  func(*GetLoadBalancerInput) GetLoadBalancerRequest
}

// Send marshals and sends the GetLoadBalancer API request.
func (r GetLoadBalancerRequest) Send(ctx context.Context) (*GetLoadBalancerResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetLoadBalancerResponse{
		GetLoadBalancerOutput: r.Request.Data.(*GetLoadBalancerOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetLoadBalancerResponse is the response type for the
// GetLoadBalancer API operation.
type GetLoadBalancerResponse struct {
	*GetLoadBalancerOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetLoadBalancer request.
func (r *GetLoadBalancerResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
