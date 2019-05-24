// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iot

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type AssociateTargetsWithJobInput struct {
	_ struct{} `type:"structure"`

	// An optional comment string describing why the job was associated with the
	// targets.
	Comment *string `locationName:"comment" type:"string"`

	// The unique identifier you assigned to this job when it was created.
	//
	// JobId is a required field
	JobId *string `location:"uri" locationName:"jobId" min:"1" type:"string" required:"true"`

	// A list of thing group ARNs that define the targets of the job.
	//
	// Targets is a required field
	Targets []string `locationName:"targets" min:"1" type:"list" required:"true"`
}

// String returns the string representation
func (s AssociateTargetsWithJobInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AssociateTargetsWithJobInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "AssociateTargetsWithJobInput"}

	if s.JobId == nil {
		invalidParams.Add(aws.NewErrParamRequired("JobId"))
	}
	if s.JobId != nil && len(*s.JobId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("JobId", 1))
	}

	if s.Targets == nil {
		invalidParams.Add(aws.NewErrParamRequired("Targets"))
	}
	if s.Targets != nil && len(s.Targets) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Targets", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s AssociateTargetsWithJobInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.Comment != nil {
		v := *s.Comment

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "comment", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.Targets) > 0 {
		v := s.Targets

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "targets", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	if s.JobId != nil {
		v := *s.JobId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "jobId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type AssociateTargetsWithJobOutput struct {
	_ struct{} `type:"structure"`

	// A short text description of the job.
	Description *string `locationName:"description" type:"string"`

	// An ARN identifying the job.
	JobArn *string `locationName:"jobArn" type:"string"`

	// The unique identifier you assigned to this job when it was created.
	JobId *string `locationName:"jobId" min:"1" type:"string"`
}

// String returns the string representation
func (s AssociateTargetsWithJobOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s AssociateTargetsWithJobOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Description != nil {
		v := *s.Description

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "description", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.JobArn != nil {
		v := *s.JobArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "jobArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.JobId != nil {
		v := *s.JobId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "jobId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opAssociateTargetsWithJob = "AssociateTargetsWithJob"

// AssociateTargetsWithJobRequest returns a request value for making API operation for
// AWS IoT.
//
// Associates a group with a continuous job. The following criteria must be
// met:
//
//    * The job must have been created with the targetSelection field set to
//    "CONTINUOUS".
//
//    * The job status must currently be "IN_PROGRESS".
//
//    * The total number of targets associated with a job must not exceed 100.
//
//    // Example sending a request using AssociateTargetsWithJobRequest.
//    req := client.AssociateTargetsWithJobRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) AssociateTargetsWithJobRequest(input *AssociateTargetsWithJobInput) AssociateTargetsWithJobRequest {
	op := &aws.Operation{
		Name:       opAssociateTargetsWithJob,
		HTTPMethod: "POST",
		HTTPPath:   "/jobs/{jobId}/targets",
	}

	if input == nil {
		input = &AssociateTargetsWithJobInput{}
	}

	req := c.newRequest(op, input, &AssociateTargetsWithJobOutput{})
	return AssociateTargetsWithJobRequest{Request: req, Input: input, Copy: c.AssociateTargetsWithJobRequest}
}

// AssociateTargetsWithJobRequest is the request type for the
// AssociateTargetsWithJob API operation.
type AssociateTargetsWithJobRequest struct {
	*aws.Request
	Input *AssociateTargetsWithJobInput
	Copy  func(*AssociateTargetsWithJobInput) AssociateTargetsWithJobRequest
}

// Send marshals and sends the AssociateTargetsWithJob API request.
func (r AssociateTargetsWithJobRequest) Send(ctx context.Context) (*AssociateTargetsWithJobResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &AssociateTargetsWithJobResponse{
		AssociateTargetsWithJobOutput: r.Request.Data.(*AssociateTargetsWithJobOutput),
		response:                      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// AssociateTargetsWithJobResponse is the response type for the
// AssociateTargetsWithJob API operation.
type AssociateTargetsWithJobResponse struct {
	*AssociateTargetsWithJobOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// AssociateTargetsWithJob request.
func (r *AssociateTargetsWithJobResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
