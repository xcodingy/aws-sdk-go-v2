// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package imagebuilder

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type DeleteImageInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the image to delete.
	//
	// ImageBuildVersionArn is a required field
	ImageBuildVersionArn *string `location:"querystring" locationName:"imageBuildVersionArn" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteImageInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteImageInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteImageInput"}

	if s.ImageBuildVersionArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("ImageBuildVersionArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteImageInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.ImageBuildVersionArn != nil {
		v := *s.ImageBuildVersionArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "imageBuildVersionArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type DeleteImageOutput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the image that was deleted.
	ImageBuildVersionArn *string `locationName:"imageBuildVersionArn" type:"string"`

	// The request ID that uniquely identifies this request.
	RequestId *string `locationName:"requestId" min:"1" type:"string"`
}

// String returns the string representation
func (s DeleteImageOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteImageOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.ImageBuildVersionArn != nil {
		v := *s.ImageBuildVersionArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "imageBuildVersionArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.RequestId != nil {
		v := *s.RequestId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "requestId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opDeleteImage = "DeleteImage"

// DeleteImageRequest returns a request value for making API operation for
// EC2 Image Builder.
//
// Deletes an image.
//
//    // Example sending a request using DeleteImageRequest.
//    req := client.DeleteImageRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/imagebuilder-2019-12-02/DeleteImage
func (c *Client) DeleteImageRequest(input *DeleteImageInput) DeleteImageRequest {
	op := &aws.Operation{
		Name:       opDeleteImage,
		HTTPMethod: "DELETE",
		HTTPPath:   "/DeleteImage",
	}

	if input == nil {
		input = &DeleteImageInput{}
	}

	req := c.newRequest(op, input, &DeleteImageOutput{})

	return DeleteImageRequest{Request: req, Input: input, Copy: c.DeleteImageRequest}
}

// DeleteImageRequest is the request type for the
// DeleteImage API operation.
type DeleteImageRequest struct {
	*aws.Request
	Input *DeleteImageInput
	Copy  func(*DeleteImageInput) DeleteImageRequest
}

// Send marshals and sends the DeleteImage API request.
func (r DeleteImageRequest) Send(ctx context.Context) (*DeleteImageResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteImageResponse{
		DeleteImageOutput: r.Request.Data.(*DeleteImageOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteImageResponse is the response type for the
// DeleteImage API operation.
type DeleteImageResponse struct {
	*DeleteImageOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteImage request.
func (r *DeleteImageResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
