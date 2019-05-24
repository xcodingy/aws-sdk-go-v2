// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Contains the parameters for DescribeConversionTasks.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DescribeConversionTasksRequest
type DescribeConversionTasksInput struct {
	_ struct{} `type:"structure"`

	// One or more conversion task IDs.
	ConversionTaskIds []string `locationName:"conversionTaskId" locationNameList:"item" type:"list"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `locationName:"dryRun" type:"boolean"`
}

// String returns the string representation
func (s DescribeConversionTasksInput) String() string {
	return awsutil.Prettify(s)
}

// Contains the output for DescribeConversionTasks.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DescribeConversionTasksResult
type DescribeConversionTasksOutput struct {
	_ struct{} `type:"structure"`

	// Information about the conversion tasks.
	ConversionTasks []ConversionTask `locationName:"conversionTasks" locationNameList:"item" type:"list"`
}

// String returns the string representation
func (s DescribeConversionTasksOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeConversionTasks = "DescribeConversionTasks"

// DescribeConversionTasksRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Describes one or more of your conversion tasks. For more information, see
// the VM Import/Export User Guide (https://docs.aws.amazon.com/vm-import/latest/userguide/).
//
// For information about the import manifest referenced by this API action,
// see VM Import Manifest (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/manifest.html).
//
//    // Example sending a request using DescribeConversionTasksRequest.
//    req := client.DescribeConversionTasksRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DescribeConversionTasks
func (c *Client) DescribeConversionTasksRequest(input *DescribeConversionTasksInput) DescribeConversionTasksRequest {
	op := &aws.Operation{
		Name:       opDescribeConversionTasks,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeConversionTasksInput{}
	}

	req := c.newRequest(op, input, &DescribeConversionTasksOutput{})
	return DescribeConversionTasksRequest{Request: req, Input: input, Copy: c.DescribeConversionTasksRequest}
}

// DescribeConversionTasksRequest is the request type for the
// DescribeConversionTasks API operation.
type DescribeConversionTasksRequest struct {
	*aws.Request
	Input *DescribeConversionTasksInput
	Copy  func(*DescribeConversionTasksInput) DescribeConversionTasksRequest
}

// Send marshals and sends the DescribeConversionTasks API request.
func (r DescribeConversionTasksRequest) Send(ctx context.Context) (*DescribeConversionTasksResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeConversionTasksResponse{
		DescribeConversionTasksOutput: r.Request.Data.(*DescribeConversionTasksOutput),
		response:                      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeConversionTasksResponse is the response type for the
// DescribeConversionTasks API operation.
type DescribeConversionTasksResponse struct {
	*DescribeConversionTasksOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeConversionTasks request.
func (r *DescribeConversionTasksResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
