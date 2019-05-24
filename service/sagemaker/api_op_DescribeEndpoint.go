// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sagemaker

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/sagemaker-2017-07-24/DescribeEndpointInput
type DescribeEndpointInput struct {
	_ struct{} `type:"structure"`

	// The name of the endpoint.
	//
	// EndpointName is a required field
	EndpointName *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeEndpointInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeEndpointInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeEndpointInput"}

	if s.EndpointName == nil {
		invalidParams.Add(aws.NewErrParamRequired("EndpointName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/sagemaker-2017-07-24/DescribeEndpointOutput
type DescribeEndpointOutput struct {
	_ struct{} `type:"structure"`

	// A timestamp that shows when the endpoint was created.
	//
	// CreationTime is a required field
	CreationTime *time.Time `type:"timestamp" timestampFormat:"unix" required:"true"`

	// The Amazon Resource Name (ARN) of the endpoint.
	//
	// EndpointArn is a required field
	EndpointArn *string `min:"20" type:"string" required:"true"`

	// The name of the endpoint configuration associated with this endpoint.
	//
	// EndpointConfigName is a required field
	EndpointConfigName *string `type:"string" required:"true"`

	// Name of the endpoint.
	//
	// EndpointName is a required field
	EndpointName *string `type:"string" required:"true"`

	// The status of the endpoint.
	//
	//    * OutOfService: Endpoint is not available to take incoming requests.
	//
	//    * Creating: CreateEndpoint is executing.
	//
	//    * Updating: UpdateEndpoint or UpdateEndpointWeightsAndCapacities is executing.
	//
	//    * SystemUpdating: Endpoint is undergoing maintenance and cannot be updated
	//    or deleted or re-scaled until it has completed. This maintenance operation
	//    does not change any customer-specified values such as VPC config, KMS
	//    encryption, model, instance type, or instance count.
	//
	//    * RollingBack: Endpoint fails to scale up or down or change its variant
	//    weight and is in the process of rolling back to its previous configuration.
	//    Once the rollback completes, endpoint returns to an InService status.
	//    This transitional status only applies to an endpoint that has autoscaling
	//    enabled and is undergoing variant weight or capacity changes as part of
	//    an UpdateEndpointWeightsAndCapacities call or when the UpdateEndpointWeightsAndCapacities
	//    operation is called explicitly.
	//
	//    * InService: Endpoint is available to process incoming requests.
	//
	//    * Deleting: DeleteEndpoint is executing.
	//
	//    * Failed: Endpoint could not be created, updated, or re-scaled. Use DescribeEndpointOutput$FailureReason
	//    for information about the failure. DeleteEndpoint is the only operation
	//    that can be performed on a failed endpoint.
	//
	// EndpointStatus is a required field
	EndpointStatus EndpointStatus `type:"string" required:"true" enum:"true"`

	// If the status of the endpoint is Failed, the reason why it failed.
	FailureReason *string `type:"string"`

	// A timestamp that shows when the endpoint was last modified.
	//
	// LastModifiedTime is a required field
	LastModifiedTime *time.Time `type:"timestamp" timestampFormat:"unix" required:"true"`

	// An array of ProductionVariantSummary objects, one for each model hosted behind
	// this endpoint.
	ProductionVariants []ProductionVariantSummary `min:"1" type:"list"`
}

// String returns the string representation
func (s DescribeEndpointOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeEndpoint = "DescribeEndpoint"

// DescribeEndpointRequest returns a request value for making API operation for
// Amazon SageMaker Service.
//
// Returns the description of an endpoint.
//
//    // Example sending a request using DescribeEndpointRequest.
//    req := client.DescribeEndpointRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/sagemaker-2017-07-24/DescribeEndpoint
func (c *Client) DescribeEndpointRequest(input *DescribeEndpointInput) DescribeEndpointRequest {
	op := &aws.Operation{
		Name:       opDescribeEndpoint,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeEndpointInput{}
	}

	req := c.newRequest(op, input, &DescribeEndpointOutput{})
	return DescribeEndpointRequest{Request: req, Input: input, Copy: c.DescribeEndpointRequest}
}

// DescribeEndpointRequest is the request type for the
// DescribeEndpoint API operation.
type DescribeEndpointRequest struct {
	*aws.Request
	Input *DescribeEndpointInput
	Copy  func(*DescribeEndpointInput) DescribeEndpointRequest
}

// Send marshals and sends the DescribeEndpoint API request.
func (r DescribeEndpointRequest) Send(ctx context.Context) (*DescribeEndpointResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeEndpointResponse{
		DescribeEndpointOutput: r.Request.Data.(*DescribeEndpointOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeEndpointResponse is the response type for the
// DescribeEndpoint API operation.
type DescribeEndpointResponse struct {
	*DescribeEndpointOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeEndpoint request.
func (r *DescribeEndpointResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
