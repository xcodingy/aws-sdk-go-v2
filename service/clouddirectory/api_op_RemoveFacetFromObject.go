// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package clouddirectory

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/clouddirectory-2017-01-11/RemoveFacetFromObjectRequest
type RemoveFacetFromObjectInput struct {
	_ struct{} `type:"structure"`

	// The ARN of the directory in which the object resides.
	//
	// DirectoryArn is a required field
	DirectoryArn *string `location:"header" locationName:"x-amz-data-partition" type:"string" required:"true"`

	// A reference to the object to remove the facet from.
	//
	// ObjectReference is a required field
	ObjectReference *ObjectReference `type:"structure" required:"true"`

	// The facet to remove. See SchemaFacet for details.
	//
	// SchemaFacet is a required field
	SchemaFacet *SchemaFacet `type:"structure" required:"true"`
}

// String returns the string representation
func (s RemoveFacetFromObjectInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *RemoveFacetFromObjectInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "RemoveFacetFromObjectInput"}

	if s.DirectoryArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("DirectoryArn"))
	}

	if s.ObjectReference == nil {
		invalidParams.Add(aws.NewErrParamRequired("ObjectReference"))
	}

	if s.SchemaFacet == nil {
		invalidParams.Add(aws.NewErrParamRequired("SchemaFacet"))
	}
	if s.SchemaFacet != nil {
		if err := s.SchemaFacet.Validate(); err != nil {
			invalidParams.AddNested("SchemaFacet", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s RemoveFacetFromObjectInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.ObjectReference != nil {
		v := s.ObjectReference

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "ObjectReference", v, metadata)
	}
	if s.SchemaFacet != nil {
		v := s.SchemaFacet

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "SchemaFacet", v, metadata)
	}
	if s.DirectoryArn != nil {
		v := *s.DirectoryArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-data-partition", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/clouddirectory-2017-01-11/RemoveFacetFromObjectResponse
type RemoveFacetFromObjectOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s RemoveFacetFromObjectOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s RemoveFacetFromObjectOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opRemoveFacetFromObject = "RemoveFacetFromObject"

// RemoveFacetFromObjectRequest returns a request value for making API operation for
// Amazon CloudDirectory.
//
// Removes the specified facet from the specified object.
//
//    // Example sending a request using RemoveFacetFromObjectRequest.
//    req := client.RemoveFacetFromObjectRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/clouddirectory-2017-01-11/RemoveFacetFromObject
func (c *Client) RemoveFacetFromObjectRequest(input *RemoveFacetFromObjectInput) RemoveFacetFromObjectRequest {
	op := &aws.Operation{
		Name:       opRemoveFacetFromObject,
		HTTPMethod: "PUT",
		HTTPPath:   "/amazonclouddirectory/2017-01-11/object/facets/delete",
	}

	if input == nil {
		input = &RemoveFacetFromObjectInput{}
	}

	req := c.newRequest(op, input, &RemoveFacetFromObjectOutput{})
	return RemoveFacetFromObjectRequest{Request: req, Input: input, Copy: c.RemoveFacetFromObjectRequest}
}

// RemoveFacetFromObjectRequest is the request type for the
// RemoveFacetFromObject API operation.
type RemoveFacetFromObjectRequest struct {
	*aws.Request
	Input *RemoveFacetFromObjectInput
	Copy  func(*RemoveFacetFromObjectInput) RemoveFacetFromObjectRequest
}

// Send marshals and sends the RemoveFacetFromObject API request.
func (r RemoveFacetFromObjectRequest) Send(ctx context.Context) (*RemoveFacetFromObjectResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &RemoveFacetFromObjectResponse{
		RemoveFacetFromObjectOutput: r.Request.Data.(*RemoveFacetFromObjectOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// RemoveFacetFromObjectResponse is the response type for the
// RemoveFacetFromObject API operation.
type RemoveFacetFromObjectResponse struct {
	*RemoveFacetFromObjectOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// RemoveFacetFromObject request.
func (r *RemoveFacetFromObjectResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
