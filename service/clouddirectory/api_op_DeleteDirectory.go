// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package clouddirectory

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/clouddirectory-2017-01-11/DeleteDirectoryRequest
type DeleteDirectoryInput struct {
	_ struct{} `type:"structure"`

	// The ARN of the directory to delete.
	//
	// DirectoryArn is a required field
	DirectoryArn *string `location:"header" locationName:"x-amz-data-partition" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteDirectoryInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteDirectoryInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteDirectoryInput"}

	if s.DirectoryArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("DirectoryArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteDirectoryInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.DirectoryArn != nil {
		v := *s.DirectoryArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-data-partition", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/clouddirectory-2017-01-11/DeleteDirectoryResponse
type DeleteDirectoryOutput struct {
	_ struct{} `type:"structure"`

	// The ARN of the deleted directory.
	//
	// DirectoryArn is a required field
	DirectoryArn *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteDirectoryOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteDirectoryOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.DirectoryArn != nil {
		v := *s.DirectoryArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "DirectoryArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opDeleteDirectory = "DeleteDirectory"

// DeleteDirectoryRequest returns a request value for making API operation for
// Amazon CloudDirectory.
//
// Deletes a directory. Only disabled directories can be deleted. A deleted
// directory cannot be undone. Exercise extreme caution when deleting directories.
//
//    // Example sending a request using DeleteDirectoryRequest.
//    req := client.DeleteDirectoryRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/clouddirectory-2017-01-11/DeleteDirectory
func (c *Client) DeleteDirectoryRequest(input *DeleteDirectoryInput) DeleteDirectoryRequest {
	op := &aws.Operation{
		Name:       opDeleteDirectory,
		HTTPMethod: "PUT",
		HTTPPath:   "/amazonclouddirectory/2017-01-11/directory",
	}

	if input == nil {
		input = &DeleteDirectoryInput{}
	}

	req := c.newRequest(op, input, &DeleteDirectoryOutput{})
	return DeleteDirectoryRequest{Request: req, Input: input, Copy: c.DeleteDirectoryRequest}
}

// DeleteDirectoryRequest is the request type for the
// DeleteDirectory API operation.
type DeleteDirectoryRequest struct {
	*aws.Request
	Input *DeleteDirectoryInput
	Copy  func(*DeleteDirectoryInput) DeleteDirectoryRequest
}

// Send marshals and sends the DeleteDirectory API request.
func (r DeleteDirectoryRequest) Send(ctx context.Context) (*DeleteDirectoryResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteDirectoryResponse{
		DeleteDirectoryOutput: r.Request.Data.(*DeleteDirectoryOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteDirectoryResponse is the response type for the
// DeleteDirectory API operation.
type DeleteDirectoryResponse struct {
	*DeleteDirectoryOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteDirectory request.
func (r *DeleteDirectoryResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
