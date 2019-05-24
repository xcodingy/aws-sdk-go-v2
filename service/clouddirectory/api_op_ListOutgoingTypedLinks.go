// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package clouddirectory

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/clouddirectory-2017-01-11/ListOutgoingTypedLinksRequest
type ListOutgoingTypedLinksInput struct {
	_ struct{} `type:"structure"`

	// The consistency level to execute the request at.
	ConsistencyLevel ConsistencyLevel `type:"string" enum:"true"`

	// The Amazon Resource Name (ARN) of the directory where you want to list the
	// typed links.
	//
	// DirectoryArn is a required field
	DirectoryArn *string `location:"header" locationName:"x-amz-data-partition" type:"string" required:"true"`

	// Provides range filters for multiple attributes. When providing ranges to
	// typed link selection, any inexact ranges must be specified at the end. Any
	// attributes that do not have a range specified are presumed to match the entire
	// range.
	FilterAttributeRanges []TypedLinkAttributeRange `type:"list"`

	// Filters are interpreted in the order of the attributes defined on the typed
	// link facet, not the order they are supplied to any API calls.
	FilterTypedLink *TypedLinkSchemaAndFacetName `type:"structure"`

	// The maximum number of results to retrieve.
	MaxResults *int64 `min:"1" type:"integer"`

	// The pagination token.
	NextToken *string `type:"string"`

	// A reference that identifies the object whose attributes will be listed.
	//
	// ObjectReference is a required field
	ObjectReference *ObjectReference `type:"structure" required:"true"`
}

// String returns the string representation
func (s ListOutgoingTypedLinksInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListOutgoingTypedLinksInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListOutgoingTypedLinksInput"}

	if s.DirectoryArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("DirectoryArn"))
	}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if s.ObjectReference == nil {
		invalidParams.Add(aws.NewErrParamRequired("ObjectReference"))
	}
	if s.FilterAttributeRanges != nil {
		for i, v := range s.FilterAttributeRanges {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "FilterAttributeRanges", i), err.(aws.ErrInvalidParams))
			}
		}
	}
	if s.FilterTypedLink != nil {
		if err := s.FilterTypedLink.Validate(); err != nil {
			invalidParams.AddNested("FilterTypedLink", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListOutgoingTypedLinksInput) MarshalFields(e protocol.FieldEncoder) error {

	if len(s.ConsistencyLevel) > 0 {
		v := s.ConsistencyLevel

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "ConsistencyLevel", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if len(s.FilterAttributeRanges) > 0 {
		v := s.FilterAttributeRanges

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "FilterAttributeRanges", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.FilterTypedLink != nil {
		v := s.FilterTypedLink

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "FilterTypedLink", v, metadata)
	}
	if s.MaxResults != nil {
		v := *s.MaxResults

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "MaxResults", protocol.Int64Value(v), metadata)
	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "NextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
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

// Please also see https://docs.aws.amazon.com/goto/WebAPI/clouddirectory-2017-01-11/ListOutgoingTypedLinksResponse
type ListOutgoingTypedLinksOutput struct {
	_ struct{} `type:"structure"`

	// The pagination token.
	NextToken *string `type:"string"`

	// Returns a typed link specifier as output.
	TypedLinkSpecifiers []TypedLinkSpecifier `type:"list"`
}

// String returns the string representation
func (s ListOutgoingTypedLinksOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListOutgoingTypedLinksOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "NextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.TypedLinkSpecifiers) > 0 {
		v := s.TypedLinkSpecifiers

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "TypedLinkSpecifiers", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	return nil
}

const opListOutgoingTypedLinks = "ListOutgoingTypedLinks"

// ListOutgoingTypedLinksRequest returns a request value for making API operation for
// Amazon CloudDirectory.
//
// Returns a paginated list of all the outgoing TypedLinkSpecifier information
// for an object. It also supports filtering by typed link facet and identity
// attributes. For more information, see Typed Links (https://docs.aws.amazon.com/clouddirectory/latest/developerguide/directory_objects_links.html#directory_objects_links_typedlink).
//
//    // Example sending a request using ListOutgoingTypedLinksRequest.
//    req := client.ListOutgoingTypedLinksRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/clouddirectory-2017-01-11/ListOutgoingTypedLinks
func (c *Client) ListOutgoingTypedLinksRequest(input *ListOutgoingTypedLinksInput) ListOutgoingTypedLinksRequest {
	op := &aws.Operation{
		Name:       opListOutgoingTypedLinks,
		HTTPMethod: "POST",
		HTTPPath:   "/amazonclouddirectory/2017-01-11/typedlink/outgoing",
	}

	if input == nil {
		input = &ListOutgoingTypedLinksInput{}
	}

	req := c.newRequest(op, input, &ListOutgoingTypedLinksOutput{})
	return ListOutgoingTypedLinksRequest{Request: req, Input: input, Copy: c.ListOutgoingTypedLinksRequest}
}

// ListOutgoingTypedLinksRequest is the request type for the
// ListOutgoingTypedLinks API operation.
type ListOutgoingTypedLinksRequest struct {
	*aws.Request
	Input *ListOutgoingTypedLinksInput
	Copy  func(*ListOutgoingTypedLinksInput) ListOutgoingTypedLinksRequest
}

// Send marshals and sends the ListOutgoingTypedLinks API request.
func (r ListOutgoingTypedLinksRequest) Send(ctx context.Context) (*ListOutgoingTypedLinksResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListOutgoingTypedLinksResponse{
		ListOutgoingTypedLinksOutput: r.Request.Data.(*ListOutgoingTypedLinksOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListOutgoingTypedLinksResponse is the response type for the
// ListOutgoingTypedLinks API operation.
type ListOutgoingTypedLinksResponse struct {
	*ListOutgoingTypedLinksOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListOutgoingTypedLinks request.
func (r *ListOutgoingTypedLinksResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
