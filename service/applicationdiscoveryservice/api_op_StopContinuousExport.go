// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package applicationdiscoveryservice

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/discovery-2015-11-01/StopContinuousExportRequest
type StopContinuousExportInput struct {
	_ struct{} `type:"structure"`

	// The unique ID assigned to this export.
	//
	// ExportId is a required field
	ExportId *string `locationName:"exportId" type:"string" required:"true"`
}

// String returns the string representation
func (s StopContinuousExportInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *StopContinuousExportInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "StopContinuousExportInput"}

	if s.ExportId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ExportId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/discovery-2015-11-01/StopContinuousExportResponse
type StopContinuousExportOutput struct {
	_ struct{} `type:"structure"`

	// Timestamp that represents when this continuous export started collecting
	// data.
	StartTime *time.Time `locationName:"startTime" type:"timestamp" timestampFormat:"unix"`

	// Timestamp that represents when this continuous export was stopped.
	StopTime *time.Time `locationName:"stopTime" type:"timestamp" timestampFormat:"unix"`
}

// String returns the string representation
func (s StopContinuousExportOutput) String() string {
	return awsutil.Prettify(s)
}

const opStopContinuousExport = "StopContinuousExport"

// StopContinuousExportRequest returns a request value for making API operation for
// AWS Application Discovery Service.
//
// Stop the continuous flow of agent's discovered data into Amazon Athena.
//
//    // Example sending a request using StopContinuousExportRequest.
//    req := client.StopContinuousExportRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/discovery-2015-11-01/StopContinuousExport
func (c *Client) StopContinuousExportRequest(input *StopContinuousExportInput) StopContinuousExportRequest {
	op := &aws.Operation{
		Name:       opStopContinuousExport,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &StopContinuousExportInput{}
	}

	req := c.newRequest(op, input, &StopContinuousExportOutput{})
	return StopContinuousExportRequest{Request: req, Input: input, Copy: c.StopContinuousExportRequest}
}

// StopContinuousExportRequest is the request type for the
// StopContinuousExport API operation.
type StopContinuousExportRequest struct {
	*aws.Request
	Input *StopContinuousExportInput
	Copy  func(*StopContinuousExportInput) StopContinuousExportRequest
}

// Send marshals and sends the StopContinuousExport API request.
func (r StopContinuousExportRequest) Send(ctx context.Context) (*StopContinuousExportResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &StopContinuousExportResponse{
		StopContinuousExportOutput: r.Request.Data.(*StopContinuousExportOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// StopContinuousExportResponse is the response type for the
// StopContinuousExport API operation.
type StopContinuousExportResponse struct {
	*StopContinuousExportOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// StopContinuousExport request.
func (r *StopContinuousExportResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
