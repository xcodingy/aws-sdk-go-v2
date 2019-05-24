// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package appsync

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/appsync-2017-07-25/ListApiKeysRequest
type ListApiKeysInput struct {
	_ struct{} `type:"structure"`

	// The API ID.
	//
	// ApiId is a required field
	ApiId *string `location:"uri" locationName:"apiId" type:"string" required:"true"`

	// The maximum number of results you want the request to return.
	MaxResults *int64 `location:"querystring" locationName:"maxResults" type:"integer"`

	// An identifier that was returned from the previous call to this operation,
	// which can be used to return the next set of items in the list.
	NextToken *string `location:"querystring" locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s ListApiKeysInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListApiKeysInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListApiKeysInput"}

	if s.ApiId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ApiId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListApiKeysInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	if s.ApiId != nil {
		v := *s.ApiId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "apiId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
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
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/appsync-2017-07-25/ListApiKeysResponse
type ListApiKeysOutput struct {
	_ struct{} `type:"structure"`

	// The ApiKey objects.
	ApiKeys []ApiKey `locationName:"apiKeys" type:"list"`

	// An identifier to be passed in the next request to this operation to return
	// the next set of items in the list.
	NextToken *string `locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s ListApiKeysOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListApiKeysOutput) MarshalFields(e protocol.FieldEncoder) error {
	if len(s.ApiKeys) > 0 {
		v := s.ApiKeys

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "apiKeys", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
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

const opListApiKeys = "ListApiKeys"

// ListApiKeysRequest returns a request value for making API operation for
// AWS AppSync.
//
// Lists the API keys for a given API.
//
// API keys are deleted automatically sometime after they expire. However, they
// may still be included in the response until they have actually been deleted.
// You can safely call DeleteApiKey to manually delete a key before it's automatically
// deleted.
//
//    // Example sending a request using ListApiKeysRequest.
//    req := client.ListApiKeysRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/appsync-2017-07-25/ListApiKeys
func (c *Client) ListApiKeysRequest(input *ListApiKeysInput) ListApiKeysRequest {
	op := &aws.Operation{
		Name:       opListApiKeys,
		HTTPMethod: "GET",
		HTTPPath:   "/v1/apis/{apiId}/apikeys",
	}

	if input == nil {
		input = &ListApiKeysInput{}
	}

	req := c.newRequest(op, input, &ListApiKeysOutput{})
	return ListApiKeysRequest{Request: req, Input: input, Copy: c.ListApiKeysRequest}
}

// ListApiKeysRequest is the request type for the
// ListApiKeys API operation.
type ListApiKeysRequest struct {
	*aws.Request
	Input *ListApiKeysInput
	Copy  func(*ListApiKeysInput) ListApiKeysRequest
}

// Send marshals and sends the ListApiKeys API request.
func (r ListApiKeysRequest) Send(ctx context.Context) (*ListApiKeysResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListApiKeysResponse{
		ListApiKeysOutput: r.Request.Data.(*ListApiKeysOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListApiKeysResponse is the response type for the
// ListApiKeys API operation.
type ListApiKeysResponse struct {
	*ListApiKeysOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListApiKeys request.
func (r *ListApiKeysResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
