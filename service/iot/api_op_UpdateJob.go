// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iot

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
)

type UpdateJobInput struct {
	_ struct{} `type:"structure"`

	// Allows you to create criteria to abort a job.
	AbortConfig *AbortConfig `locationName:"abortConfig" type:"structure"`

	// A short text description of the job.
	Description *string `locationName:"description" type:"string"`

	// Allows you to create a staged rollout of the job.
	JobExecutionsRolloutConfig *JobExecutionsRolloutConfig `locationName:"jobExecutionsRolloutConfig" type:"structure"`

	// The ID of the job to be updated.
	//
	// JobId is a required field
	JobId *string `location:"uri" locationName:"jobId" min:"1" type:"string" required:"true"`

	// Configuration information for pre-signed S3 URLs.
	PresignedUrlConfig *PresignedUrlConfig `locationName:"presignedUrlConfig" type:"structure"`

	// Specifies the amount of time each device has to finish its execution of the
	// job. The timer is started when the job execution status is set to IN_PROGRESS.
	// If the job execution status is not set to another terminal state before the
	// time expires, it will be automatically set to TIMED_OUT.
	TimeoutConfig *TimeoutConfig `locationName:"timeoutConfig" type:"structure"`
}

// String returns the string representation
func (s UpdateJobInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateJobInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateJobInput"}

	if s.JobId == nil {
		invalidParams.Add(aws.NewErrParamRequired("JobId"))
	}
	if s.JobId != nil && len(*s.JobId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("JobId", 1))
	}
	if s.AbortConfig != nil {
		if err := s.AbortConfig.Validate(); err != nil {
			invalidParams.AddNested("AbortConfig", err.(aws.ErrInvalidParams))
		}
	}
	if s.JobExecutionsRolloutConfig != nil {
		if err := s.JobExecutionsRolloutConfig.Validate(); err != nil {
			invalidParams.AddNested("JobExecutionsRolloutConfig", err.(aws.ErrInvalidParams))
		}
	}
	if s.PresignedUrlConfig != nil {
		if err := s.PresignedUrlConfig.Validate(); err != nil {
			invalidParams.AddNested("PresignedUrlConfig", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateJobInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.AbortConfig != nil {
		v := s.AbortConfig

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "abortConfig", v, metadata)
	}
	if s.Description != nil {
		v := *s.Description

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "description", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.JobExecutionsRolloutConfig != nil {
		v := s.JobExecutionsRolloutConfig

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "jobExecutionsRolloutConfig", v, metadata)
	}
	if s.PresignedUrlConfig != nil {
		v := s.PresignedUrlConfig

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "presignedUrlConfig", v, metadata)
	}
	if s.TimeoutConfig != nil {
		v := s.TimeoutConfig

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "timeoutConfig", v, metadata)
	}
	if s.JobId != nil {
		v := *s.JobId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "jobId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type UpdateJobOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s UpdateJobOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateJobOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opUpdateJob = "UpdateJob"

// UpdateJobRequest returns a request value for making API operation for
// AWS IoT.
//
// Updates supported fields of the specified job.
//
//    // Example sending a request using UpdateJobRequest.
//    req := client.UpdateJobRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) UpdateJobRequest(input *UpdateJobInput) UpdateJobRequest {
	op := &aws.Operation{
		Name:       opUpdateJob,
		HTTPMethod: "PATCH",
		HTTPPath:   "/jobs/{jobId}",
	}

	if input == nil {
		input = &UpdateJobInput{}
	}

	req := c.newRequest(op, input, &UpdateJobOutput{})
	req.Handlers.Unmarshal.Remove(restjson.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return UpdateJobRequest{Request: req, Input: input, Copy: c.UpdateJobRequest}
}

// UpdateJobRequest is the request type for the
// UpdateJob API operation.
type UpdateJobRequest struct {
	*aws.Request
	Input *UpdateJobInput
	Copy  func(*UpdateJobInput) UpdateJobRequest
}

// Send marshals and sends the UpdateJob API request.
func (r UpdateJobRequest) Send(ctx context.Context) (*UpdateJobResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateJobResponse{
		UpdateJobOutput: r.Request.Data.(*UpdateJobOutput),
		response:        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateJobResponse is the response type for the
// UpdateJob API operation.
type UpdateJobResponse struct {
	*UpdateJobOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateJob request.
func (r *UpdateJobResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
