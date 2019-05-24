// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package clouddirectory

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/clouddirectory-2017-01-11/DetachPolicyRequest
type DetachPolicyInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) that is associated with the Directory where
	// both objects reside. For more information, see arns.
	//
	// DirectoryArn is a required field
	DirectoryArn *string `location:"header" locationName:"x-amz-data-partition" type:"string" required:"true"`

	// Reference that identifies the object whose policy object will be detached.
	//
	// ObjectReference is a required field
	ObjectReference *ObjectReference `type:"structure" required:"true"`

	// Reference that identifies the policy object.
	//
	// PolicyReference is a required field
	PolicyReference *ObjectReference `type:"structure" required:"true"`
}

// String returns the string representation
func (s DetachPolicyInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DetachPolicyInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DetachPolicyInput"}

	if s.DirectoryArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("DirectoryArn"))
	}

	if s.ObjectReference == nil {
		invalidParams.Add(aws.NewErrParamRequired("ObjectReference"))
	}

	if s.PolicyReference == nil {
		invalidParams.Add(aws.NewErrParamRequired("PolicyReference"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DetachPolicyInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.ObjectReference != nil {
		v := s.ObjectReference

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "ObjectReference", v, metadata)
	}
	if s.PolicyReference != nil {
		v := s.PolicyReference

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "PolicyReference", v, metadata)
	}
	if s.DirectoryArn != nil {
		v := *s.DirectoryArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-data-partition", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/clouddirectory-2017-01-11/DetachPolicyResponse
type DetachPolicyOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DetachPolicyOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DetachPolicyOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opDetachPolicy = "DetachPolicy"

// DetachPolicyRequest returns a request value for making API operation for
// Amazon CloudDirectory.
//
// Detaches a policy from an object.
//
//    // Example sending a request using DetachPolicyRequest.
//    req := client.DetachPolicyRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/clouddirectory-2017-01-11/DetachPolicy
func (c *Client) DetachPolicyRequest(input *DetachPolicyInput) DetachPolicyRequest {
	op := &aws.Operation{
		Name:       opDetachPolicy,
		HTTPMethod: "PUT",
		HTTPPath:   "/amazonclouddirectory/2017-01-11/policy/detach",
	}

	if input == nil {
		input = &DetachPolicyInput{}
	}

	req := c.newRequest(op, input, &DetachPolicyOutput{})
	return DetachPolicyRequest{Request: req, Input: input, Copy: c.DetachPolicyRequest}
}

// DetachPolicyRequest is the request type for the
// DetachPolicy API operation.
type DetachPolicyRequest struct {
	*aws.Request
	Input *DetachPolicyInput
	Copy  func(*DetachPolicyInput) DetachPolicyRequest
}

// Send marshals and sends the DetachPolicy API request.
func (r DetachPolicyRequest) Send(ctx context.Context) (*DetachPolicyResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DetachPolicyResponse{
		DetachPolicyOutput: r.Request.Data.(*DetachPolicyOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DetachPolicyResponse is the response type for the
// DetachPolicy API operation.
type DetachPolicyResponse struct {
	*DetachPolicyOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DetachPolicy request.
func (r *DetachPolicyResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
