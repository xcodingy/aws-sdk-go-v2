// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package clouddirectory

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/clouddirectory-2017-01-11/TagResourceRequest
type TagResourceInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the resource. Tagging is only supported
	// for directories.
	//
	// ResourceArn is a required field
	ResourceArn *string `type:"string" required:"true"`

	// A list of tag key-value pairs.
	//
	// Tags is a required field
	Tags []Tag `type:"list" required:"true"`
}

// String returns the string representation
func (s TagResourceInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *TagResourceInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "TagResourceInput"}

	if s.ResourceArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("ResourceArn"))
	}

	if s.Tags == nil {
		invalidParams.Add(aws.NewErrParamRequired("Tags"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s TagResourceInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.ResourceArn != nil {
		v := *s.ResourceArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "ResourceArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.Tags) > 0 {
		v := s.Tags

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "Tags", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/clouddirectory-2017-01-11/TagResourceResponse
type TagResourceOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s TagResourceOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s TagResourceOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opTagResource = "TagResource"

// TagResourceRequest returns a request value for making API operation for
// Amazon CloudDirectory.
//
// An API operation for adding tags to a resource.
//
//    // Example sending a request using TagResourceRequest.
//    req := client.TagResourceRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/clouddirectory-2017-01-11/TagResource
func (c *Client) TagResourceRequest(input *TagResourceInput) TagResourceRequest {
	op := &aws.Operation{
		Name:       opTagResource,
		HTTPMethod: "PUT",
		HTTPPath:   "/amazonclouddirectory/2017-01-11/tags/add",
	}

	if input == nil {
		input = &TagResourceInput{}
	}

	req := c.newRequest(op, input, &TagResourceOutput{})
	return TagResourceRequest{Request: req, Input: input, Copy: c.TagResourceRequest}
}

// TagResourceRequest is the request type for the
// TagResource API operation.
type TagResourceRequest struct {
	*aws.Request
	Input *TagResourceInput
	Copy  func(*TagResourceInput) TagResourceRequest
}

// Send marshals and sends the TagResource API request.
func (r TagResourceRequest) Send(ctx context.Context) (*TagResourceResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &TagResourceResponse{
		TagResourceOutput: r.Request.Data.(*TagResourceOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// TagResourceResponse is the response type for the
// TagResource API operation.
type TagResourceResponse struct {
	*TagResourceOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// TagResource request.
func (r *TagResourceResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
