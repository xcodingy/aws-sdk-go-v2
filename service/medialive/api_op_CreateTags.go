// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package medialive

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/medialive-2017-10-14/CreateTagsRequest
type CreateTagsInput struct {
	_ struct{} `type:"structure"`

	// ResourceArn is a required field
	ResourceArn *string `location:"uri" locationName:"resource-arn" type:"string" required:"true"`

	Tags map[string]string `locationName:"tags" type:"map"`
}

// String returns the string representation
func (s CreateTagsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateTagsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateTagsInput"}

	if s.ResourceArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("ResourceArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateTagsInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	if len(s.Tags) > 0 {
		v := s.Tags

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "tags", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetValue(k1, protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ms0.End()

	}
	if s.ResourceArn != nil {
		v := *s.ResourceArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "resource-arn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/medialive-2017-10-14/CreateTagsOutput
type CreateTagsOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s CreateTagsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateTagsOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opCreateTags = "CreateTags"

// CreateTagsRequest returns a request value for making API operation for
// AWS Elemental MediaLive.
//
// Create tags for a resource
//
//    // Example sending a request using CreateTagsRequest.
//    req := client.CreateTagsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/medialive-2017-10-14/CreateTags
func (c *Client) CreateTagsRequest(input *CreateTagsInput) CreateTagsRequest {
	op := &aws.Operation{
		Name:       opCreateTags,
		HTTPMethod: "POST",
		HTTPPath:   "/prod/tags/{resource-arn}",
	}

	if input == nil {
		input = &CreateTagsInput{}
	}

	req := c.newRequest(op, input, &CreateTagsOutput{})
	req.Handlers.Unmarshal.Remove(restjson.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return CreateTagsRequest{Request: req, Input: input, Copy: c.CreateTagsRequest}
}

// CreateTagsRequest is the request type for the
// CreateTags API operation.
type CreateTagsRequest struct {
	*aws.Request
	Input *CreateTagsInput
	Copy  func(*CreateTagsInput) CreateTagsRequest
}

// Send marshals and sends the CreateTags API request.
func (r CreateTagsRequest) Send(ctx context.Context) (*CreateTagsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateTagsResponse{
		CreateTagsOutput: r.Request.Data.(*CreateTagsOutput),
		response:         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateTagsResponse is the response type for the
// CreateTags API operation.
type CreateTagsResponse struct {
	*CreateTagsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateTags request.
func (r *CreateTagsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
