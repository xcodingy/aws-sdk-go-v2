// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package clouddirectory

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/clouddirectory-2017-01-11/DetachFromIndexRequest
type DetachFromIndexInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the directory the index and object exist
	// in.
	//
	// DirectoryArn is a required field
	DirectoryArn *string `location:"header" locationName:"x-amz-data-partition" type:"string" required:"true"`

	// A reference to the index object.
	//
	// IndexReference is a required field
	IndexReference *ObjectReference `type:"structure" required:"true"`

	// A reference to the object being detached from the index.
	//
	// TargetReference is a required field
	TargetReference *ObjectReference `type:"structure" required:"true"`
}

// String returns the string representation
func (s DetachFromIndexInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DetachFromIndexInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DetachFromIndexInput"}

	if s.DirectoryArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("DirectoryArn"))
	}

	if s.IndexReference == nil {
		invalidParams.Add(aws.NewErrParamRequired("IndexReference"))
	}

	if s.TargetReference == nil {
		invalidParams.Add(aws.NewErrParamRequired("TargetReference"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DetachFromIndexInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.IndexReference != nil {
		v := s.IndexReference

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "IndexReference", v, metadata)
	}
	if s.TargetReference != nil {
		v := s.TargetReference

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "TargetReference", v, metadata)
	}
	if s.DirectoryArn != nil {
		v := *s.DirectoryArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-data-partition", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/clouddirectory-2017-01-11/DetachFromIndexResponse
type DetachFromIndexOutput struct {
	_ struct{} `type:"structure"`

	// The ObjectIdentifier of the object that was detached from the index.
	DetachedObjectIdentifier *string `type:"string"`
}

// String returns the string representation
func (s DetachFromIndexOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DetachFromIndexOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.DetachedObjectIdentifier != nil {
		v := *s.DetachedObjectIdentifier

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "DetachedObjectIdentifier", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opDetachFromIndex = "DetachFromIndex"

// DetachFromIndexRequest returns a request value for making API operation for
// Amazon CloudDirectory.
//
// Detaches the specified object from the specified index.
//
//    // Example sending a request using DetachFromIndexRequest.
//    req := client.DetachFromIndexRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/clouddirectory-2017-01-11/DetachFromIndex
func (c *Client) DetachFromIndexRequest(input *DetachFromIndexInput) DetachFromIndexRequest {
	op := &aws.Operation{
		Name:       opDetachFromIndex,
		HTTPMethod: "PUT",
		HTTPPath:   "/amazonclouddirectory/2017-01-11/index/detach",
	}

	if input == nil {
		input = &DetachFromIndexInput{}
	}

	req := c.newRequest(op, input, &DetachFromIndexOutput{})
	return DetachFromIndexRequest{Request: req, Input: input, Copy: c.DetachFromIndexRequest}
}

// DetachFromIndexRequest is the request type for the
// DetachFromIndex API operation.
type DetachFromIndexRequest struct {
	*aws.Request
	Input *DetachFromIndexInput
	Copy  func(*DetachFromIndexInput) DetachFromIndexRequest
}

// Send marshals and sends the DetachFromIndex API request.
func (r DetachFromIndexRequest) Send(ctx context.Context) (*DetachFromIndexResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DetachFromIndexResponse{
		DetachFromIndexOutput: r.Request.Data.(*DetachFromIndexOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DetachFromIndexResponse is the response type for the
// DetachFromIndex API operation.
type DetachFromIndexResponse struct {
	*DetachFromIndexOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DetachFromIndex request.
func (r *DetachFromIndexResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
