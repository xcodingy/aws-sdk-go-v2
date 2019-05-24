// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ssm

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ssm-2014-11-06/DescribeMaintenanceWindowTargetsRequest
type DescribeMaintenanceWindowTargetsInput struct {
	_ struct{} `type:"structure"`

	// Optional filters that can be used to narrow down the scope of the returned
	// window targets. The supported filter keys are Type, WindowTargetId and OwnerInformation.
	Filters []MaintenanceWindowFilter `type:"list"`

	// The maximum number of items to return for this call. The call also returns
	// a token that you can specify in a subsequent call to get the next set of
	// results.
	MaxResults *int64 `min:"10" type:"integer"`

	// The token for the next set of items to return. (You received this token from
	// a previous call.)
	NextToken *string `type:"string"`

	// The ID of the Maintenance Window whose targets should be retrieved.
	//
	// WindowId is a required field
	WindowId *string `min:"20" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeMaintenanceWindowTargetsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeMaintenanceWindowTargetsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeMaintenanceWindowTargetsInput"}
	if s.MaxResults != nil && *s.MaxResults < 10 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 10))
	}

	if s.WindowId == nil {
		invalidParams.Add(aws.NewErrParamRequired("WindowId"))
	}
	if s.WindowId != nil && len(*s.WindowId) < 20 {
		invalidParams.Add(aws.NewErrParamMinLen("WindowId", 20))
	}
	if s.Filters != nil {
		for i, v := range s.Filters {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Filters", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ssm-2014-11-06/DescribeMaintenanceWindowTargetsResult
type DescribeMaintenanceWindowTargetsOutput struct {
	_ struct{} `type:"structure"`

	// The token to use when requesting the next set of items. If there are no additional
	// items to return, the string is empty.
	NextToken *string `type:"string"`

	// Information about the targets in the Maintenance Window.
	Targets []MaintenanceWindowTarget `type:"list"`
}

// String returns the string representation
func (s DescribeMaintenanceWindowTargetsOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeMaintenanceWindowTargets = "DescribeMaintenanceWindowTargets"

// DescribeMaintenanceWindowTargetsRequest returns a request value for making API operation for
// Amazon Simple Systems Manager (SSM).
//
// Lists the targets registered with the Maintenance Window.
//
//    // Example sending a request using DescribeMaintenanceWindowTargetsRequest.
//    req := client.DescribeMaintenanceWindowTargetsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ssm-2014-11-06/DescribeMaintenanceWindowTargets
func (c *Client) DescribeMaintenanceWindowTargetsRequest(input *DescribeMaintenanceWindowTargetsInput) DescribeMaintenanceWindowTargetsRequest {
	op := &aws.Operation{
		Name:       opDescribeMaintenanceWindowTargets,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeMaintenanceWindowTargetsInput{}
	}

	req := c.newRequest(op, input, &DescribeMaintenanceWindowTargetsOutput{})
	return DescribeMaintenanceWindowTargetsRequest{Request: req, Input: input, Copy: c.DescribeMaintenanceWindowTargetsRequest}
}

// DescribeMaintenanceWindowTargetsRequest is the request type for the
// DescribeMaintenanceWindowTargets API operation.
type DescribeMaintenanceWindowTargetsRequest struct {
	*aws.Request
	Input *DescribeMaintenanceWindowTargetsInput
	Copy  func(*DescribeMaintenanceWindowTargetsInput) DescribeMaintenanceWindowTargetsRequest
}

// Send marshals and sends the DescribeMaintenanceWindowTargets API request.
func (r DescribeMaintenanceWindowTargetsRequest) Send(ctx context.Context) (*DescribeMaintenanceWindowTargetsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeMaintenanceWindowTargetsResponse{
		DescribeMaintenanceWindowTargetsOutput: r.Request.Data.(*DescribeMaintenanceWindowTargetsOutput),
		response:                               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeMaintenanceWindowTargetsResponse is the response type for the
// DescribeMaintenanceWindowTargets API operation.
type DescribeMaintenanceWindowTargetsResponse struct {
	*DescribeMaintenanceWindowTargetsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeMaintenanceWindowTargets request.
func (r *DescribeMaintenanceWindowTargetsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
