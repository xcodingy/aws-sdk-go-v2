// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ssm

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ssm-2014-11-06/DescribeSessionsRequest
type DescribeSessionsInput struct {
	_ struct{} `type:"structure"`

	// One or more filters to limit the type of sessions returned by the request.
	Filters []SessionFilter `min:"1" type:"list"`

	// The maximum number of items to return for this call. The call also returns
	// a token that you can specify in a subsequent call to get the next set of
	// results.
	MaxResults *int64 `min:"1" type:"integer"`

	// The token for the next set of items to return. (You received this token from
	// a previous call.)
	NextToken *string `type:"string"`

	// The session status to retrieve a list of sessions for. For example, "Active".
	//
	// State is a required field
	State SessionState `type:"string" required:"true" enum:"true"`
}

// String returns the string representation
func (s DescribeSessionsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeSessionsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeSessionsInput"}
	if s.Filters != nil && len(s.Filters) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Filters", 1))
	}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}
	if len(s.State) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("State"))
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

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ssm-2014-11-06/DescribeSessionsResponse
type DescribeSessionsOutput struct {
	_ struct{} `type:"structure"`

	// The token for the next set of items to return. (You received this token from
	// a previous call.)
	NextToken *string `type:"string"`

	// A list of sessions meeting the request parameters.
	Sessions []Session `type:"list"`
}

// String returns the string representation
func (s DescribeSessionsOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeSessions = "DescribeSessions"

// DescribeSessionsRequest returns a request value for making API operation for
// Amazon Simple Systems Manager (SSM).
//
// Retrieves a list of all active sessions (both connected and disconnected)
// or terminated sessions from the past 30 days.
//
//    // Example sending a request using DescribeSessionsRequest.
//    req := client.DescribeSessionsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ssm-2014-11-06/DescribeSessions
func (c *Client) DescribeSessionsRequest(input *DescribeSessionsInput) DescribeSessionsRequest {
	op := &aws.Operation{
		Name:       opDescribeSessions,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeSessionsInput{}
	}

	req := c.newRequest(op, input, &DescribeSessionsOutput{})
	return DescribeSessionsRequest{Request: req, Input: input, Copy: c.DescribeSessionsRequest}
}

// DescribeSessionsRequest is the request type for the
// DescribeSessions API operation.
type DescribeSessionsRequest struct {
	*aws.Request
	Input *DescribeSessionsInput
	Copy  func(*DescribeSessionsInput) DescribeSessionsRequest
}

// Send marshals and sends the DescribeSessions API request.
func (r DescribeSessionsRequest) Send(ctx context.Context) (*DescribeSessionsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeSessionsResponse{
		DescribeSessionsOutput: r.Request.Data.(*DescribeSessionsOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeSessionsResponse is the response type for the
// DescribeSessions API operation.
type DescribeSessionsResponse struct {
	*DescribeSessionsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeSessions request.
func (r *DescribeSessionsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
