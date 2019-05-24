// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package mq

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/mq-2017-11-27/DeleteTagsRequest
type DeleteTagsInput struct {
	_ struct{} `type:"structure"`

	// ResourceArn is a required field
	ResourceArn *string `location:"uri" locationName:"resource-arn" type:"string" required:"true"`

	// TagKeys is a required field
	TagKeys []string `location:"querystring" locationName:"tagKeys" type:"list" required:"true"`
}

// String returns the string representation
func (s DeleteTagsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteTagsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteTagsInput"}

	if s.ResourceArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("ResourceArn"))
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
func (s DeleteTagsInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	if s.ResourceArn != nil {
		v := *s.ResourceArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "resource-arn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.TagKeys) > 0 {
		v := s.TagKeys

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.QueryTarget, "tagKeys", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/mq-2017-11-27/DeleteTagsOutput
type DeleteTagsOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteTagsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteTagsOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opDeleteTags = "DeleteTags"

// DeleteTagsRequest returns a request value for making API operation for
// AmazonMQ.
//
// Removes a tag from a resource.
//
//    // Example sending a request using DeleteTagsRequest.
//    req := client.DeleteTagsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/mq-2017-11-27/DeleteTags
func (c *Client) DeleteTagsRequest(input *DeleteTagsInput) DeleteTagsRequest {
	op := &aws.Operation{
		Name:       opDeleteTags,
		HTTPMethod: "DELETE",
		HTTPPath:   "/v1/tags/{resource-arn}",
	}

	if input == nil {
		input = &DeleteTagsInput{}
	}

	req := c.newRequest(op, input, &DeleteTagsOutput{})
	req.Handlers.Unmarshal.Remove(restjson.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return DeleteTagsRequest{Request: req, Input: input, Copy: c.DeleteTagsRequest}
}

// DeleteTagsRequest is the request type for the
// DeleteTags API operation.
type DeleteTagsRequest struct {
	*aws.Request
	Input *DeleteTagsInput
	Copy  func(*DeleteTagsInput) DeleteTagsRequest
}

// Send marshals and sends the DeleteTags API request.
func (r DeleteTagsRequest) Send(ctx context.Context) (*DeleteTagsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteTagsResponse{
		DeleteTagsOutput: r.Request.Data.(*DeleteTagsOutput),
		response:         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteTagsResponse is the response type for the
// DeleteTags API operation.
type DeleteTagsResponse struct {
	*DeleteTagsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteTags request.
func (r *DeleteTagsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
