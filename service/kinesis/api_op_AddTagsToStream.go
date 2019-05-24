// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package kinesis

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
)

// Represents the input for AddTagsToStream.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/kinesis-2013-12-02/AddTagsToStreamInput
type AddTagsToStreamInput struct {
	_ struct{} `type:"structure"`

	// The name of the stream.
	//
	// StreamName is a required field
	StreamName *string `min:"1" type:"string" required:"true"`

	// A set of up to 10 key-value pairs to use to create the tags.
	//
	// Tags is a required field
	Tags map[string]string `min:"1" type:"map" required:"true"`
}

// String returns the string representation
func (s AddTagsToStreamInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AddTagsToStreamInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "AddTagsToStreamInput"}

	if s.StreamName == nil {
		invalidParams.Add(aws.NewErrParamRequired("StreamName"))
	}
	if s.StreamName != nil && len(*s.StreamName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("StreamName", 1))
	}

	if s.Tags == nil {
		invalidParams.Add(aws.NewErrParamRequired("Tags"))
	}
	if s.Tags != nil && len(s.Tags) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Tags", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/kinesis-2013-12-02/AddTagsToStreamOutput
type AddTagsToStreamOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s AddTagsToStreamOutput) String() string {
	return awsutil.Prettify(s)
}

const opAddTagsToStream = "AddTagsToStream"

// AddTagsToStreamRequest returns a request value for making API operation for
// Amazon Kinesis.
//
// Adds or updates tags for the specified Kinesis data stream. Each time you
// invoke this operation, you can specify up to 10 tags. If you want to add
// more than 10 tags to your stream, you can invoke this operation multiple
// times. In total, each stream can have up to 50 tags.
//
// If tags have already been assigned to the stream, AddTagsToStream overwrites
// any existing tags that correspond to the specified tag keys.
//
// AddTagsToStream has a limit of five transactions per second per account.
//
//    // Example sending a request using AddTagsToStreamRequest.
//    req := client.AddTagsToStreamRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/kinesis-2013-12-02/AddTagsToStream
func (c *Client) AddTagsToStreamRequest(input *AddTagsToStreamInput) AddTagsToStreamRequest {
	op := &aws.Operation{
		Name:       opAddTagsToStream,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &AddTagsToStreamInput{}
	}

	req := c.newRequest(op, input, &AddTagsToStreamOutput{})
	req.Handlers.Unmarshal.Remove(jsonrpc.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return AddTagsToStreamRequest{Request: req, Input: input, Copy: c.AddTagsToStreamRequest}
}

// AddTagsToStreamRequest is the request type for the
// AddTagsToStream API operation.
type AddTagsToStreamRequest struct {
	*aws.Request
	Input *AddTagsToStreamInput
	Copy  func(*AddTagsToStreamInput) AddTagsToStreamRequest
}

// Send marshals and sends the AddTagsToStream API request.
func (r AddTagsToStreamRequest) Send(ctx context.Context) (*AddTagsToStreamResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &AddTagsToStreamResponse{
		AddTagsToStreamOutput: r.Request.Data.(*AddTagsToStreamOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// AddTagsToStreamResponse is the response type for the
// AddTagsToStream API operation.
type AddTagsToStreamResponse struct {
	*AddTagsToStreamOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// AddTagsToStream request.
func (r *AddTagsToStreamResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
