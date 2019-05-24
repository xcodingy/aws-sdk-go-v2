// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Contains the parameters for PurchaseScheduledInstances.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/PurchaseScheduledInstancesRequest
type PurchaseScheduledInstancesInput struct {
	_ struct{} `type:"structure"`

	// Unique, case-sensitive identifier that ensures the idempotency of the request.
	// For more information, see Ensuring Idempotency (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/Run_Instance_Idempotency.html).
	ClientToken *string `type:"string" idempotencyToken:"true"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `type:"boolean"`

	// One or more purchase requests.
	//
	// PurchaseRequests is a required field
	PurchaseRequests []PurchaseRequest `locationName:"PurchaseRequest" locationNameList:"PurchaseRequest" min:"1" type:"list" required:"true"`
}

// String returns the string representation
func (s PurchaseScheduledInstancesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PurchaseScheduledInstancesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "PurchaseScheduledInstancesInput"}

	if s.PurchaseRequests == nil {
		invalidParams.Add(aws.NewErrParamRequired("PurchaseRequests"))
	}
	if s.PurchaseRequests != nil && len(s.PurchaseRequests) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("PurchaseRequests", 1))
	}
	if s.PurchaseRequests != nil {
		for i, v := range s.PurchaseRequests {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "PurchaseRequests", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Contains the output of PurchaseScheduledInstances.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/PurchaseScheduledInstancesResult
type PurchaseScheduledInstancesOutput struct {
	_ struct{} `type:"structure"`

	// Information about the Scheduled Instances.
	ScheduledInstanceSet []ScheduledInstance `locationName:"scheduledInstanceSet" locationNameList:"item" type:"list"`
}

// String returns the string representation
func (s PurchaseScheduledInstancesOutput) String() string {
	return awsutil.Prettify(s)
}

const opPurchaseScheduledInstances = "PurchaseScheduledInstances"

// PurchaseScheduledInstancesRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Purchases one or more Scheduled Instances with the specified schedule.
//
// Scheduled Instances enable you to purchase Amazon EC2 compute capacity by
// the hour for a one-year term. Before you can purchase a Scheduled Instance,
// you must call DescribeScheduledInstanceAvailability to check for available
// schedules and obtain a purchase token. After you purchase a Scheduled Instance,
// you must call RunScheduledInstances during each scheduled time period.
//
// After you purchase a Scheduled Instance, you can't cancel, modify, or resell
// your purchase.
//
//    // Example sending a request using PurchaseScheduledInstancesRequest.
//    req := client.PurchaseScheduledInstancesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/PurchaseScheduledInstances
func (c *Client) PurchaseScheduledInstancesRequest(input *PurchaseScheduledInstancesInput) PurchaseScheduledInstancesRequest {
	op := &aws.Operation{
		Name:       opPurchaseScheduledInstances,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &PurchaseScheduledInstancesInput{}
	}

	req := c.newRequest(op, input, &PurchaseScheduledInstancesOutput{})
	return PurchaseScheduledInstancesRequest{Request: req, Input: input, Copy: c.PurchaseScheduledInstancesRequest}
}

// PurchaseScheduledInstancesRequest is the request type for the
// PurchaseScheduledInstances API operation.
type PurchaseScheduledInstancesRequest struct {
	*aws.Request
	Input *PurchaseScheduledInstancesInput
	Copy  func(*PurchaseScheduledInstancesInput) PurchaseScheduledInstancesRequest
}

// Send marshals and sends the PurchaseScheduledInstances API request.
func (r PurchaseScheduledInstancesRequest) Send(ctx context.Context) (*PurchaseScheduledInstancesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &PurchaseScheduledInstancesResponse{
		PurchaseScheduledInstancesOutput: r.Request.Data.(*PurchaseScheduledInstancesOutput),
		response:                         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// PurchaseScheduledInstancesResponse is the response type for the
// PurchaseScheduledInstances API operation.
type PurchaseScheduledInstancesResponse struct {
	*PurchaseScheduledInstancesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// PurchaseScheduledInstances request.
func (r *PurchaseScheduledInstancesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
