// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iot

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type ListThingRegistrationTaskReportsInput struct {
	_ struct{} `type:"structure"`

	// The maximum number of results to return per request.
	MaxResults *int64 `location:"querystring" locationName:"maxResults" min:"1" type:"integer"`

	// The token to retrieve the next set of results.
	NextToken *string `location:"querystring" locationName:"nextToken" type:"string"`

	// The type of task report.
	//
	// ReportType is a required field
	ReportType ReportType `location:"querystring" locationName:"reportType" type:"string" required:"true" enum:"true"`

	// The id of the task.
	//
	// TaskId is a required field
	TaskId *string `location:"uri" locationName:"taskId" type:"string" required:"true"`
}

// String returns the string representation
func (s ListThingRegistrationTaskReportsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListThingRegistrationTaskReportsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListThingRegistrationTaskReportsInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}
	if len(s.ReportType) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("ReportType"))
	}

	if s.TaskId == nil {
		invalidParams.Add(aws.NewErrParamRequired("TaskId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListThingRegistrationTaskReportsInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.TaskId != nil {
		v := *s.TaskId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "taskId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.MaxResults != nil {
		v := *s.MaxResults

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "maxResults", protocol.Int64Value(v), metadata)
	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.ReportType) > 0 {
		v := s.ReportType

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "reportType", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	return nil
}

type ListThingRegistrationTaskReportsOutput struct {
	_ struct{} `type:"structure"`

	// The token used to get the next set of results, or null if there are no additional
	// results.
	NextToken *string `locationName:"nextToken" type:"string"`

	// The type of task report.
	ReportType ReportType `locationName:"reportType" type:"string" enum:"true"`

	// Links to the task resources.
	ResourceLinks []string `locationName:"resourceLinks" type:"list"`
}

// String returns the string representation
func (s ListThingRegistrationTaskReportsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListThingRegistrationTaskReportsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.ReportType) > 0 {
		v := s.ReportType

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "reportType", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if len(s.ResourceLinks) > 0 {
		v := s.ResourceLinks

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "resourceLinks", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	return nil
}

const opListThingRegistrationTaskReports = "ListThingRegistrationTaskReports"

// ListThingRegistrationTaskReportsRequest returns a request value for making API operation for
// AWS IoT.
//
// Information about the thing registration tasks.
//
//    // Example sending a request using ListThingRegistrationTaskReportsRequest.
//    req := client.ListThingRegistrationTaskReportsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) ListThingRegistrationTaskReportsRequest(input *ListThingRegistrationTaskReportsInput) ListThingRegistrationTaskReportsRequest {
	op := &aws.Operation{
		Name:       opListThingRegistrationTaskReports,
		HTTPMethod: "GET",
		HTTPPath:   "/thing-registration-tasks/{taskId}/reports",
	}

	if input == nil {
		input = &ListThingRegistrationTaskReportsInput{}
	}

	req := c.newRequest(op, input, &ListThingRegistrationTaskReportsOutput{})
	return ListThingRegistrationTaskReportsRequest{Request: req, Input: input, Copy: c.ListThingRegistrationTaskReportsRequest}
}

// ListThingRegistrationTaskReportsRequest is the request type for the
// ListThingRegistrationTaskReports API operation.
type ListThingRegistrationTaskReportsRequest struct {
	*aws.Request
	Input *ListThingRegistrationTaskReportsInput
	Copy  func(*ListThingRegistrationTaskReportsInput) ListThingRegistrationTaskReportsRequest
}

// Send marshals and sends the ListThingRegistrationTaskReports API request.
func (r ListThingRegistrationTaskReportsRequest) Send(ctx context.Context) (*ListThingRegistrationTaskReportsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListThingRegistrationTaskReportsResponse{
		ListThingRegistrationTaskReportsOutput: r.Request.Data.(*ListThingRegistrationTaskReportsOutput),
		response:                               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListThingRegistrationTaskReportsResponse is the response type for the
// ListThingRegistrationTaskReports API operation.
type ListThingRegistrationTaskReportsResponse struct {
	*ListThingRegistrationTaskReportsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListThingRegistrationTaskReports request.
func (r *ListThingRegistrationTaskReportsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
