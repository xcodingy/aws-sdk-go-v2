// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package batch

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/batch-2016-08-10/CancelJobRequest
type CancelJobInput struct {
	_ struct{} `type:"structure"`

	// The AWS Batch job ID of the job to cancel.
	//
	// JobId is a required field
	JobId *string `locationName:"jobId" type:"string" required:"true"`

	// A message to attach to the job that explains the reason for canceling it.
	// This message is returned by future DescribeJobs operations on the job. This
	// message is also recorded in the AWS Batch activity logs.
	//
	// Reason is a required field
	Reason *string `locationName:"reason" type:"string" required:"true"`
}

// String returns the string representation
func (s CancelJobInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CancelJobInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CancelJobInput"}

	if s.JobId == nil {
		invalidParams.Add(aws.NewErrParamRequired("JobId"))
	}

	if s.Reason == nil {
		invalidParams.Add(aws.NewErrParamRequired("Reason"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CancelJobInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	if s.JobId != nil {
		v := *s.JobId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "jobId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Reason != nil {
		v := *s.Reason

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "reason", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/batch-2016-08-10/CancelJobResponse
type CancelJobOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s CancelJobOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CancelJobOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opCancelJob = "CancelJob"

// CancelJobRequest returns a request value for making API operation for
// AWS Batch.
//
// Cancels a job in an AWS Batch job queue. Jobs that are in the SUBMITTED,
// PENDING, or RUNNABLE state are cancelled. Jobs that have progressed to STARTING
// or RUNNING are not cancelled (but the API operation still succeeds, even
// if no job is cancelled); these jobs must be terminated with the TerminateJob
// operation.
//
//    // Example sending a request using CancelJobRequest.
//    req := client.CancelJobRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/batch-2016-08-10/CancelJob
func (c *Client) CancelJobRequest(input *CancelJobInput) CancelJobRequest {
	op := &aws.Operation{
		Name:       opCancelJob,
		HTTPMethod: "POST",
		HTTPPath:   "/v1/canceljob",
	}

	if input == nil {
		input = &CancelJobInput{}
	}

	req := c.newRequest(op, input, &CancelJobOutput{})
	return CancelJobRequest{Request: req, Input: input, Copy: c.CancelJobRequest}
}

// CancelJobRequest is the request type for the
// CancelJob API operation.
type CancelJobRequest struct {
	*aws.Request
	Input *CancelJobInput
	Copy  func(*CancelJobInput) CancelJobRequest
}

// Send marshals and sends the CancelJob API request.
func (r CancelJobRequest) Send(ctx context.Context) (*CancelJobResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CancelJobResponse{
		CancelJobOutput: r.Request.Data.(*CancelJobOutput),
		response:        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CancelJobResponse is the response type for the
// CancelJob API operation.
type CancelJobResponse struct {
	*CancelJobOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CancelJob request.
func (r *CancelJobResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
