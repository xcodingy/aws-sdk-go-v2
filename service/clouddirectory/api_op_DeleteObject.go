// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package clouddirectory

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/clouddirectory-2017-01-11/DeleteObjectRequest
type DeleteObjectInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) that is associated with the Directory where
	// the object resides. For more information, see arns.
	//
	// DirectoryArn is a required field
	DirectoryArn *string `location:"header" locationName:"x-amz-data-partition" type:"string" required:"true"`

	// A reference that identifies the object.
	//
	// ObjectReference is a required field
	ObjectReference *ObjectReference `type:"structure" required:"true"`
}

// String returns the string representation
func (s DeleteObjectInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteObjectInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteObjectInput"}

	if s.DirectoryArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("DirectoryArn"))
	}

	if s.ObjectReference == nil {
		invalidParams.Add(aws.NewErrParamRequired("ObjectReference"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteObjectInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.ObjectReference != nil {
		v := s.ObjectReference

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "ObjectReference", v, metadata)
	}
	if s.DirectoryArn != nil {
		v := *s.DirectoryArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-data-partition", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/clouddirectory-2017-01-11/DeleteObjectResponse
type DeleteObjectOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteObjectOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteObjectOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opDeleteObject = "DeleteObject"

// DeleteObjectRequest returns a request value for making API operation for
// Amazon CloudDirectory.
//
// Deletes an object and its associated attributes. Only objects with no children
// and no parents can be deleted. The maximum number of attributes that can
// be deleted during an object deletion is 30. For more information, see Amazon
// Cloud Directory Limits (https://docs.aws.amazon.com/clouddirectory/latest/developerguide/limits.html).
//
//    // Example sending a request using DeleteObjectRequest.
//    req := client.DeleteObjectRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/clouddirectory-2017-01-11/DeleteObject
func (c *Client) DeleteObjectRequest(input *DeleteObjectInput) DeleteObjectRequest {
	op := &aws.Operation{
		Name:       opDeleteObject,
		HTTPMethod: "PUT",
		HTTPPath:   "/amazonclouddirectory/2017-01-11/object/delete",
	}

	if input == nil {
		input = &DeleteObjectInput{}
	}

	req := c.newRequest(op, input, &DeleteObjectOutput{})
	return DeleteObjectRequest{Request: req, Input: input, Copy: c.DeleteObjectRequest}
}

// DeleteObjectRequest is the request type for the
// DeleteObject API operation.
type DeleteObjectRequest struct {
	*aws.Request
	Input *DeleteObjectInput
	Copy  func(*DeleteObjectInput) DeleteObjectRequest
}

// Send marshals and sends the DeleteObject API request.
func (r DeleteObjectRequest) Send(ctx context.Context) (*DeleteObjectResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteObjectResponse{
		DeleteObjectOutput: r.Request.Data.(*DeleteObjectOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteObjectResponse is the response type for the
// DeleteObject API operation.
type DeleteObjectResponse struct {
	*DeleteObjectOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteObject request.
func (r *DeleteObjectResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
