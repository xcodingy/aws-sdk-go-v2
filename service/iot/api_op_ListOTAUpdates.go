// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iot

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type ListOTAUpdatesInput struct {
	_ struct{} `type:"structure"`

	// The maximum number of results to return at one time.
	MaxResults *int64 `location:"querystring" locationName:"maxResults" min:"1" type:"integer"`

	// A token used to retrieve the next set of results.
	NextToken *string `location:"querystring" locationName:"nextToken" type:"string"`

	// The OTA update job status.
	OtaUpdateStatus OTAUpdateStatus `location:"querystring" locationName:"otaUpdateStatus" type:"string" enum:"true"`
}

// String returns the string representation
func (s ListOTAUpdatesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListOTAUpdatesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListOTAUpdatesInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListOTAUpdatesInput) MarshalFields(e protocol.FieldEncoder) error {

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
	if len(s.OtaUpdateStatus) > 0 {
		v := s.OtaUpdateStatus

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "otaUpdateStatus", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	return nil
}

type ListOTAUpdatesOutput struct {
	_ struct{} `type:"structure"`

	// A token to use to get the next set of results.
	NextToken *string `locationName:"nextToken" type:"string"`

	// A list of OTA update jobs.
	OtaUpdates []OTAUpdateSummary `locationName:"otaUpdates" type:"list"`
}

// String returns the string representation
func (s ListOTAUpdatesOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListOTAUpdatesOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.OtaUpdates) > 0 {
		v := s.OtaUpdates

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "otaUpdates", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	return nil
}

const opListOTAUpdates = "ListOTAUpdates"

// ListOTAUpdatesRequest returns a request value for making API operation for
// AWS IoT.
//
// Lists OTA updates.
//
//    // Example sending a request using ListOTAUpdatesRequest.
//    req := client.ListOTAUpdatesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) ListOTAUpdatesRequest(input *ListOTAUpdatesInput) ListOTAUpdatesRequest {
	op := &aws.Operation{
		Name:       opListOTAUpdates,
		HTTPMethod: "GET",
		HTTPPath:   "/otaUpdates",
	}

	if input == nil {
		input = &ListOTAUpdatesInput{}
	}

	req := c.newRequest(op, input, &ListOTAUpdatesOutput{})
	return ListOTAUpdatesRequest{Request: req, Input: input, Copy: c.ListOTAUpdatesRequest}
}

// ListOTAUpdatesRequest is the request type for the
// ListOTAUpdates API operation.
type ListOTAUpdatesRequest struct {
	*aws.Request
	Input *ListOTAUpdatesInput
	Copy  func(*ListOTAUpdatesInput) ListOTAUpdatesRequest
}

// Send marshals and sends the ListOTAUpdates API request.
func (r ListOTAUpdatesRequest) Send(ctx context.Context) (*ListOTAUpdatesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListOTAUpdatesResponse{
		ListOTAUpdatesOutput: r.Request.Data.(*ListOTAUpdatesOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListOTAUpdatesResponse is the response type for the
// ListOTAUpdates API operation.
type ListOTAUpdatesResponse struct {
	*ListOTAUpdatesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListOTAUpdates request.
func (r *ListOTAUpdatesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
