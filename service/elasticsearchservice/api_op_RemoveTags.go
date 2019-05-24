// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package elasticsearchservice

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
)

// Container for the parameters to the RemoveTags operation. Specify the ARN
// for the Elasticsearch domain from which you want to remove the specified
// TagKey.
type RemoveTagsInput struct {
	_ struct{} `type:"structure"`

	// Specifies the ARN for the Elasticsearch domain from which you want to delete
	// the specified tags.
	//
	// ARN is a required field
	ARN *string `type:"string" required:"true"`

	// Specifies the TagKey list which you want to remove from the Elasticsearch
	// domain.
	//
	// TagKeys is a required field
	TagKeys []string `type:"list" required:"true"`
}

// String returns the string representation
func (s RemoveTagsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *RemoveTagsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "RemoveTagsInput"}

	if s.ARN == nil {
		invalidParams.Add(aws.NewErrParamRequired("ARN"))
	}

	if s.TagKeys == nil {
		invalidParams.Add(aws.NewErrParamRequired("TagKeys"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s RemoveTagsInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.ARN != nil {
		v := *s.ARN

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "ARN", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.TagKeys) > 0 {
		v := s.TagKeys

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "TagKeys", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	return nil
}

type RemoveTagsOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s RemoveTagsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s RemoveTagsOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opRemoveTags = "RemoveTags"

// RemoveTagsRequest returns a request value for making API operation for
// Amazon Elasticsearch Service.
//
// Removes the specified set of tags from the specified Elasticsearch domain.
//
//    // Example sending a request using RemoveTagsRequest.
//    req := client.RemoveTagsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) RemoveTagsRequest(input *RemoveTagsInput) RemoveTagsRequest {
	op := &aws.Operation{
		Name:       opRemoveTags,
		HTTPMethod: "POST",
		HTTPPath:   "/2015-01-01/tags-removal",
	}

	if input == nil {
		input = &RemoveTagsInput{}
	}

	req := c.newRequest(op, input, &RemoveTagsOutput{})
	req.Handlers.Unmarshal.Remove(restjson.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return RemoveTagsRequest{Request: req, Input: input, Copy: c.RemoveTagsRequest}
}

// RemoveTagsRequest is the request type for the
// RemoveTags API operation.
type RemoveTagsRequest struct {
	*aws.Request
	Input *RemoveTagsInput
	Copy  func(*RemoveTagsInput) RemoveTagsRequest
}

// Send marshals and sends the RemoveTags API request.
func (r RemoveTagsRequest) Send(ctx context.Context) (*RemoveTagsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &RemoveTagsResponse{
		RemoveTagsOutput: r.Request.Data.(*RemoveTagsOutput),
		response:         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// RemoveTagsResponse is the response type for the
// RemoveTags API operation.
type RemoveTagsResponse struct {
	*RemoveTagsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// RemoveTags request.
func (r *RemoveTagsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
