// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package comprehendmedical

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type StopEntitiesDetectionV2JobInput struct {
	_ struct{} `type:"structure"`

	// The identifier of the medical entities job to stop.
	//
	// JobId is a required field
	JobId *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s StopEntitiesDetectionV2JobInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *StopEntitiesDetectionV2JobInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "StopEntitiesDetectionV2JobInput"}

	if s.JobId == nil {
		invalidParams.Add(aws.NewErrParamRequired("JobId"))
	}
	if s.JobId != nil && len(*s.JobId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("JobId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type StopEntitiesDetectionV2JobOutput struct {
	_ struct{} `type:"structure"`

	// The identifier of the medical entities detection job that was stopped.
	JobId *string `min:"1" type:"string"`
}

// String returns the string representation
func (s StopEntitiesDetectionV2JobOutput) String() string {
	return awsutil.Prettify(s)
}

const opStopEntitiesDetectionV2Job = "StopEntitiesDetectionV2Job"

// StopEntitiesDetectionV2JobRequest returns a request value for making API operation for
// AWS Comprehend Medical.
//
// Stops a medical entities detection job in progress.
//
//    // Example sending a request using StopEntitiesDetectionV2JobRequest.
//    req := client.StopEntitiesDetectionV2JobRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/comprehendmedical-2018-10-30/StopEntitiesDetectionV2Job
func (c *Client) StopEntitiesDetectionV2JobRequest(input *StopEntitiesDetectionV2JobInput) StopEntitiesDetectionV2JobRequest {
	op := &aws.Operation{
		Name:       opStopEntitiesDetectionV2Job,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &StopEntitiesDetectionV2JobInput{}
	}

	req := c.newRequest(op, input, &StopEntitiesDetectionV2JobOutput{})

	return StopEntitiesDetectionV2JobRequest{Request: req, Input: input, Copy: c.StopEntitiesDetectionV2JobRequest}
}

// StopEntitiesDetectionV2JobRequest is the request type for the
// StopEntitiesDetectionV2Job API operation.
type StopEntitiesDetectionV2JobRequest struct {
	*aws.Request
	Input *StopEntitiesDetectionV2JobInput
	Copy  func(*StopEntitiesDetectionV2JobInput) StopEntitiesDetectionV2JobRequest
}

// Send marshals and sends the StopEntitiesDetectionV2Job API request.
func (r StopEntitiesDetectionV2JobRequest) Send(ctx context.Context) (*StopEntitiesDetectionV2JobResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &StopEntitiesDetectionV2JobResponse{
		StopEntitiesDetectionV2JobOutput: r.Request.Data.(*StopEntitiesDetectionV2JobOutput),
		response:                         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// StopEntitiesDetectionV2JobResponse is the response type for the
// StopEntitiesDetectionV2Job API operation.
type StopEntitiesDetectionV2JobResponse struct {
	*StopEntitiesDetectionV2JobOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// StopEntitiesDetectionV2Job request.
func (r *StopEntitiesDetectionV2JobResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
