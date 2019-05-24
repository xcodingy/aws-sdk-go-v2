// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iot

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type ListIndicesInput struct {
	_ struct{} `type:"structure"`

	// The maximum number of results to return at one time.
	MaxResults *int64 `location:"querystring" locationName:"maxResults" min:"1" type:"integer"`

	// The token used to get the next set of results, or null if there are no additional
	// results.
	NextToken *string `location:"querystring" locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s ListIndicesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListIndicesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListIndicesInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListIndicesInput) MarshalFields(e protocol.FieldEncoder) error {

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
	return nil
}

type ListIndicesOutput struct {
	_ struct{} `type:"structure"`

	// The index names.
	IndexNames []string `locationName:"indexNames" type:"list"`

	// The token used to get the next set of results, or null if there are no additional
	// results.
	NextToken *string `locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s ListIndicesOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListIndicesOutput) MarshalFields(e protocol.FieldEncoder) error {
	if len(s.IndexNames) > 0 {
		v := s.IndexNames

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "indexNames", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opListIndices = "ListIndices"

// ListIndicesRequest returns a request value for making API operation for
// AWS IoT.
//
// Lists the search indices.
//
//    // Example sending a request using ListIndicesRequest.
//    req := client.ListIndicesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) ListIndicesRequest(input *ListIndicesInput) ListIndicesRequest {
	op := &aws.Operation{
		Name:       opListIndices,
		HTTPMethod: "GET",
		HTTPPath:   "/indices",
	}

	if input == nil {
		input = &ListIndicesInput{}
	}

	req := c.newRequest(op, input, &ListIndicesOutput{})
	return ListIndicesRequest{Request: req, Input: input, Copy: c.ListIndicesRequest}
}

// ListIndicesRequest is the request type for the
// ListIndices API operation.
type ListIndicesRequest struct {
	*aws.Request
	Input *ListIndicesInput
	Copy  func(*ListIndicesInput) ListIndicesRequest
}

// Send marshals and sends the ListIndices API request.
func (r ListIndicesRequest) Send(ctx context.Context) (*ListIndicesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListIndicesResponse{
		ListIndicesOutput: r.Request.Data.(*ListIndicesOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListIndicesResponse is the response type for the
// ListIndices API operation.
type ListIndicesResponse struct {
	*ListIndicesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListIndices request.
func (r *ListIndicesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
