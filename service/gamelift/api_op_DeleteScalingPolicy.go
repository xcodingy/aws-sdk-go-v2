// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package gamelift

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
)

// Represents the input for a request action.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/gamelift-2015-10-01/DeleteScalingPolicyInput
type DeleteScalingPolicyInput struct {
	_ struct{} `type:"structure"`

	// Unique identifier for a fleet to be deleted.
	//
	// FleetId is a required field
	FleetId *string `type:"string" required:"true"`

	// Descriptive label that is associated with a scaling policy. Policy names
	// do not need to be unique.
	//
	// Name is a required field
	Name *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteScalingPolicyInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteScalingPolicyInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteScalingPolicyInput"}

	if s.FleetId == nil {
		invalidParams.Add(aws.NewErrParamRequired("FleetId"))
	}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}
	if s.Name != nil && len(*s.Name) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Name", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/gamelift-2015-10-01/DeleteScalingPolicyOutput
type DeleteScalingPolicyOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteScalingPolicyOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteScalingPolicy = "DeleteScalingPolicy"

// DeleteScalingPolicyRequest returns a request value for making API operation for
// Amazon GameLift.
//
// Deletes a fleet scaling policy. This action means that the policy is no longer
// in force and removes all record of it. To delete a scaling policy, specify
// both the scaling policy name and the fleet ID it is associated with.
//
// To temporarily suspend scaling policies, call StopFleetActions. This operation
// suspends all policies for the fleet.
//
//    * DescribeFleetCapacity
//
//    * UpdateFleetCapacity
//
//    * DescribeEC2InstanceLimits
//
//    * Manage scaling policies:
//
// PutScalingPolicy (auto-scaling)
//
// DescribeScalingPolicies (auto-scaling)
//
// DeleteScalingPolicy (auto-scaling)
//
//    * Manage fleet actions:
//
// StartFleetActions
//
// StopFleetActions
//
//    // Example sending a request using DeleteScalingPolicyRequest.
//    req := client.DeleteScalingPolicyRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/gamelift-2015-10-01/DeleteScalingPolicy
func (c *Client) DeleteScalingPolicyRequest(input *DeleteScalingPolicyInput) DeleteScalingPolicyRequest {
	op := &aws.Operation{
		Name:       opDeleteScalingPolicy,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteScalingPolicyInput{}
	}

	req := c.newRequest(op, input, &DeleteScalingPolicyOutput{})
	req.Handlers.Unmarshal.Remove(jsonrpc.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return DeleteScalingPolicyRequest{Request: req, Input: input, Copy: c.DeleteScalingPolicyRequest}
}

// DeleteScalingPolicyRequest is the request type for the
// DeleteScalingPolicy API operation.
type DeleteScalingPolicyRequest struct {
	*aws.Request
	Input *DeleteScalingPolicyInput
	Copy  func(*DeleteScalingPolicyInput) DeleteScalingPolicyRequest
}

// Send marshals and sends the DeleteScalingPolicy API request.
func (r DeleteScalingPolicyRequest) Send(ctx context.Context) (*DeleteScalingPolicyResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteScalingPolicyResponse{
		DeleteScalingPolicyOutput: r.Request.Data.(*DeleteScalingPolicyOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteScalingPolicyResponse is the response type for the
// DeleteScalingPolicy API operation.
type DeleteScalingPolicyResponse struct {
	*DeleteScalingPolicyOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteScalingPolicy request.
func (r *DeleteScalingPolicyResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
