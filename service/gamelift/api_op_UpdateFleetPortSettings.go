// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package gamelift

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents the input for a request action.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/gamelift-2015-10-01/UpdateFleetPortSettingsInput
type UpdateFleetPortSettingsInput struct {
	_ struct{} `type:"structure"`

	// Unique identifier for a fleet to update port settings for.
	//
	// FleetId is a required field
	FleetId *string `type:"string" required:"true"`

	// Collection of port settings to be added to the fleet record.
	InboundPermissionAuthorizations []IpPermission `type:"list"`

	// Collection of port settings to be removed from the fleet record.
	InboundPermissionRevocations []IpPermission `type:"list"`
}

// String returns the string representation
func (s UpdateFleetPortSettingsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateFleetPortSettingsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateFleetPortSettingsInput"}

	if s.FleetId == nil {
		invalidParams.Add(aws.NewErrParamRequired("FleetId"))
	}
	if s.InboundPermissionAuthorizations != nil {
		for i, v := range s.InboundPermissionAuthorizations {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "InboundPermissionAuthorizations", i), err.(aws.ErrInvalidParams))
			}
		}
	}
	if s.InboundPermissionRevocations != nil {
		for i, v := range s.InboundPermissionRevocations {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "InboundPermissionRevocations", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents the returned data in response to a request action.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/gamelift-2015-10-01/UpdateFleetPortSettingsOutput
type UpdateFleetPortSettingsOutput struct {
	_ struct{} `type:"structure"`

	// Unique identifier for a fleet that was updated.
	FleetId *string `type:"string"`
}

// String returns the string representation
func (s UpdateFleetPortSettingsOutput) String() string {
	return awsutil.Prettify(s)
}

const opUpdateFleetPortSettings = "UpdateFleetPortSettings"

// UpdateFleetPortSettingsRequest returns a request value for making API operation for
// Amazon GameLift.
//
// Updates port settings for a fleet. To update settings, specify the fleet
// ID to be updated and list the permissions you want to update. List the permissions
// you want to add in InboundPermissionAuthorizations, and permissions you want
// to remove in InboundPermissionRevocations. Permissions to be removed must
// match existing fleet permissions. If successful, the fleet ID for the updated
// fleet is returned.
//
// Learn more
//
// Working with Fleets (https://docs.aws.amazon.com/gamelift/latest/developerguide/fleets-intro.html).
//
// Related operations
//
//    * CreateFleet
//
//    * ListFleets
//
//    * DeleteFleet
//
//    * Describe fleets:
//
// DescribeFleetAttributes
//
// DescribeFleetCapacity
//
// DescribeFleetPortSettings
//
// DescribeFleetUtilization
//
// DescribeRuntimeConfiguration
//
// DescribeEC2InstanceLimits
//
// DescribeFleetEvents
//
//    * Update fleets:
//
// UpdateFleetAttributes
//
// UpdateFleetCapacity
//
// UpdateFleetPortSettings
//
// UpdateRuntimeConfiguration
//
//    * Manage fleet actions:
//
// StartFleetActions
//
// StopFleetActions
//
//    // Example sending a request using UpdateFleetPortSettingsRequest.
//    req := client.UpdateFleetPortSettingsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/gamelift-2015-10-01/UpdateFleetPortSettings
func (c *Client) UpdateFleetPortSettingsRequest(input *UpdateFleetPortSettingsInput) UpdateFleetPortSettingsRequest {
	op := &aws.Operation{
		Name:       opUpdateFleetPortSettings,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateFleetPortSettingsInput{}
	}

	req := c.newRequest(op, input, &UpdateFleetPortSettingsOutput{})
	return UpdateFleetPortSettingsRequest{Request: req, Input: input, Copy: c.UpdateFleetPortSettingsRequest}
}

// UpdateFleetPortSettingsRequest is the request type for the
// UpdateFleetPortSettings API operation.
type UpdateFleetPortSettingsRequest struct {
	*aws.Request
	Input *UpdateFleetPortSettingsInput
	Copy  func(*UpdateFleetPortSettingsInput) UpdateFleetPortSettingsRequest
}

// Send marshals and sends the UpdateFleetPortSettings API request.
func (r UpdateFleetPortSettingsRequest) Send(ctx context.Context) (*UpdateFleetPortSettingsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateFleetPortSettingsResponse{
		UpdateFleetPortSettingsOutput: r.Request.Data.(*UpdateFleetPortSettingsOutput),
		response:                      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateFleetPortSettingsResponse is the response type for the
// UpdateFleetPortSettings API operation.
type UpdateFleetPortSettingsResponse struct {
	*UpdateFleetPortSettingsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateFleetPortSettings request.
func (r *UpdateFleetPortSettingsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
