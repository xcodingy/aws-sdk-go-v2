// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iotjobsdataplane

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/iot-jobs-data-2017-09-29/UpdateJobExecutionRequest
type UpdateJobExecutionInput struct {
	_ struct{} `type:"structure"`

	// Optional. A number that identifies a particular job execution on a particular
	// device.
	ExecutionNumber *int64 `locationName:"executionNumber" type:"long"`

	// Optional. The expected current version of the job execution. Each time you
	// update the job execution, its version is incremented. If the version of the
	// job execution stored in Jobs does not match, the update is rejected with
	// a VersionMismatch error, and an ErrorResponse that contains the current job
	// execution status data is returned. (This makes it unnecessary to perform
	// a separate DescribeJobExecution request in order to obtain the job execution
	// status data.)
	ExpectedVersion *int64 `locationName:"expectedVersion" type:"long"`

	// Optional. When set to true, the response contains the job document. The default
	// is false.
	IncludeJobDocument *bool `locationName:"includeJobDocument" type:"boolean"`

	// Optional. When included and set to true, the response contains the JobExecutionState
	// data. The default is false.
	IncludeJobExecutionState *bool `locationName:"includeJobExecutionState" type:"boolean"`

	// The unique identifier assigned to this job when it was created.
	//
	// JobId is a required field
	JobId *string `location:"uri" locationName:"jobId" min:"1" type:"string" required:"true"`

	// The new status for the job execution (IN_PROGRESS, FAILED, SUCCESS, or REJECTED).
	// This must be specified on every update.
	//
	// Status is a required field
	Status JobExecutionStatus `locationName:"status" type:"string" required:"true" enum:"true"`

	// Optional. A collection of name/value pairs that describe the status of the
	// job execution. If not specified, the statusDetails are unchanged.
	StatusDetails map[string]string `locationName:"statusDetails" type:"map"`

	// Specifies the amount of time this device has to finish execution of this
	// job. If the job execution status is not set to a terminal state before this
	// timer expires, or before the timer is reset (by again calling UpdateJobExecution,
	// setting the status to IN_PROGRESS and specifying a new timeout value in this
	// field) the job execution status will be automatically set to TIMED_OUT. Note
	// that setting or resetting this timeout has no effect on that job execution
	// timeout which may have been specified when the job was created (CreateJob
	// using field timeoutConfig).
	StepTimeoutInMinutes *int64 `locationName:"stepTimeoutInMinutes" type:"long"`

	// The name of the thing associated with the device.
	//
	// ThingName is a required field
	ThingName *string `location:"uri" locationName:"thingName" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s UpdateJobExecutionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateJobExecutionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateJobExecutionInput"}

	if s.JobId == nil {
		invalidParams.Add(aws.NewErrParamRequired("JobId"))
	}
	if s.JobId != nil && len(*s.JobId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("JobId", 1))
	}
	if len(s.Status) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("Status"))
	}

	if s.ThingName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ThingName"))
	}
	if s.ThingName != nil && len(*s.ThingName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ThingName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateJobExecutionInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.ExecutionNumber != nil {
		v := *s.ExecutionNumber

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "executionNumber", protocol.Int64Value(v), metadata)
	}
	if s.ExpectedVersion != nil {
		v := *s.ExpectedVersion

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "expectedVersion", protocol.Int64Value(v), metadata)
	}
	if s.IncludeJobDocument != nil {
		v := *s.IncludeJobDocument

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "includeJobDocument", protocol.BoolValue(v), metadata)
	}
	if s.IncludeJobExecutionState != nil {
		v := *s.IncludeJobExecutionState

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "includeJobExecutionState", protocol.BoolValue(v), metadata)
	}
	if len(s.Status) > 0 {
		v := s.Status

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "status", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if len(s.StatusDetails) > 0 {
		v := s.StatusDetails

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "statusDetails", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetValue(k1, protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ms0.End()

	}
	if s.StepTimeoutInMinutes != nil {
		v := *s.StepTimeoutInMinutes

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "stepTimeoutInMinutes", protocol.Int64Value(v), metadata)
	}
	if s.JobId != nil {
		v := *s.JobId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "jobId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ThingName != nil {
		v := *s.ThingName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "thingName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/iot-jobs-data-2017-09-29/UpdateJobExecutionResponse
type UpdateJobExecutionOutput struct {
	_ struct{} `type:"structure"`

	// A JobExecutionState object.
	ExecutionState *JobExecutionState `locationName:"executionState" type:"structure"`

	// The contents of the Job Documents.
	JobDocument *string `locationName:"jobDocument" type:"string"`
}

// String returns the string representation
func (s UpdateJobExecutionOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateJobExecutionOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.ExecutionState != nil {
		v := s.ExecutionState

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "executionState", v, metadata)
	}
	if s.JobDocument != nil {
		v := *s.JobDocument

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "jobDocument", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opUpdateJobExecution = "UpdateJobExecution"

// UpdateJobExecutionRequest returns a request value for making API operation for
// AWS IoT Jobs Data Plane.
//
// Updates the status of a job execution.
//
//    // Example sending a request using UpdateJobExecutionRequest.
//    req := client.UpdateJobExecutionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iot-jobs-data-2017-09-29/UpdateJobExecution
func (c *Client) UpdateJobExecutionRequest(input *UpdateJobExecutionInput) UpdateJobExecutionRequest {
	op := &aws.Operation{
		Name:       opUpdateJobExecution,
		HTTPMethod: "POST",
		HTTPPath:   "/things/{thingName}/jobs/{jobId}",
	}

	if input == nil {
		input = &UpdateJobExecutionInput{}
	}

	req := c.newRequest(op, input, &UpdateJobExecutionOutput{})
	return UpdateJobExecutionRequest{Request: req, Input: input, Copy: c.UpdateJobExecutionRequest}
}

// UpdateJobExecutionRequest is the request type for the
// UpdateJobExecution API operation.
type UpdateJobExecutionRequest struct {
	*aws.Request
	Input *UpdateJobExecutionInput
	Copy  func(*UpdateJobExecutionInput) UpdateJobExecutionRequest
}

// Send marshals and sends the UpdateJobExecution API request.
func (r UpdateJobExecutionRequest) Send(ctx context.Context) (*UpdateJobExecutionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateJobExecutionResponse{
		UpdateJobExecutionOutput: r.Request.Data.(*UpdateJobExecutionOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateJobExecutionResponse is the response type for the
// UpdateJobExecution API operation.
type UpdateJobExecutionResponse struct {
	*UpdateJobExecutionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateJobExecution request.
func (r *UpdateJobExecutionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
