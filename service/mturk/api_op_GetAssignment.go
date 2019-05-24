// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package mturk

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/mturk-requester-2017-01-17/GetAssignmentRequest
type GetAssignmentInput struct {
	_ struct{} `type:"structure"`

	// The ID of the Assignment to be retrieved.
	//
	// AssignmentId is a required field
	AssignmentId *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s GetAssignmentInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetAssignmentInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetAssignmentInput"}

	if s.AssignmentId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AssignmentId"))
	}
	if s.AssignmentId != nil && len(*s.AssignmentId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("AssignmentId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/mturk-requester-2017-01-17/GetAssignmentResponse
type GetAssignmentOutput struct {
	_ struct{} `type:"structure"`

	// The assignment. The response includes one Assignment element.
	Assignment *Assignment `type:"structure"`

	// The HIT associated with this assignment. The response includes one HIT element.
	HIT *HIT `type:"structure"`
}

// String returns the string representation
func (s GetAssignmentOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetAssignment = "GetAssignment"

// GetAssignmentRequest returns a request value for making API operation for
// Amazon Mechanical Turk.
//
// The GetAssignment operation retrieves the details of the specified Assignment.
//
//    // Example sending a request using GetAssignmentRequest.
//    req := client.GetAssignmentRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/mturk-requester-2017-01-17/GetAssignment
func (c *Client) GetAssignmentRequest(input *GetAssignmentInput) GetAssignmentRequest {
	op := &aws.Operation{
		Name:       opGetAssignment,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetAssignmentInput{}
	}

	req := c.newRequest(op, input, &GetAssignmentOutput{})
	return GetAssignmentRequest{Request: req, Input: input, Copy: c.GetAssignmentRequest}
}

// GetAssignmentRequest is the request type for the
// GetAssignment API operation.
type GetAssignmentRequest struct {
	*aws.Request
	Input *GetAssignmentInput
	Copy  func(*GetAssignmentInput) GetAssignmentRequest
}

// Send marshals and sends the GetAssignment API request.
func (r GetAssignmentRequest) Send(ctx context.Context) (*GetAssignmentResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetAssignmentResponse{
		GetAssignmentOutput: r.Request.Data.(*GetAssignmentOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetAssignmentResponse is the response type for the
// GetAssignment API operation.
type GetAssignmentResponse struct {
	*GetAssignmentOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetAssignment request.
func (r *GetAssignmentResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
