// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package clouddirectory

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/clouddirectory-2017-01-11/ListAppliedSchemaArnsRequest
type ListAppliedSchemaArnsInput struct {
	_ struct{} `type:"structure"`

	// The ARN of the directory you are listing.
	//
	// DirectoryArn is a required field
	DirectoryArn *string `type:"string" required:"true"`

	// The maximum number of results to retrieve.
	MaxResults *int64 `min:"1" type:"integer"`

	// The pagination token.
	NextToken *string `type:"string"`

	// The response for ListAppliedSchemaArns when this parameter is used will list
	// all minor version ARNs for a major version.
	SchemaArn *string `type:"string"`
}

// String returns the string representation
func (s ListAppliedSchemaArnsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListAppliedSchemaArnsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListAppliedSchemaArnsInput"}

	if s.DirectoryArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("DirectoryArn"))
	}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListAppliedSchemaArnsInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.DirectoryArn != nil {
		v := *s.DirectoryArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "DirectoryArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
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
	if s.SchemaArn != nil {
		v := *s.SchemaArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "SchemaArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/clouddirectory-2017-01-11/ListAppliedSchemaArnsResponse
type ListAppliedSchemaArnsOutput struct {
	_ struct{} `type:"structure"`

	// The pagination token.
	NextToken *string `type:"string"`

	// The ARNs of schemas that are applied to the directory.
	SchemaArns []string `type:"list"`
}

// String returns the string representation
func (s ListAppliedSchemaArnsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListAppliedSchemaArnsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "NextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.SchemaArns) > 0 {
		v := s.SchemaArns

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "SchemaArns", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	return nil
}

const opListAppliedSchemaArns = "ListAppliedSchemaArns"

// ListAppliedSchemaArnsRequest returns a request value for making API operation for
// Amazon CloudDirectory.
//
// Lists schema major versions applied to a directory. If SchemaArn is provided,
// lists the minor version.
//
//    // Example sending a request using ListAppliedSchemaArnsRequest.
//    req := client.ListAppliedSchemaArnsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/clouddirectory-2017-01-11/ListAppliedSchemaArns
func (c *Client) ListAppliedSchemaArnsRequest(input *ListAppliedSchemaArnsInput) ListAppliedSchemaArnsRequest {
	op := &aws.Operation{
		Name:       opListAppliedSchemaArns,
		HTTPMethod: "POST",
		HTTPPath:   "/amazonclouddirectory/2017-01-11/schema/applied",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListAppliedSchemaArnsInput{}
	}

	req := c.newRequest(op, input, &ListAppliedSchemaArnsOutput{})
	return ListAppliedSchemaArnsRequest{Request: req, Input: input, Copy: c.ListAppliedSchemaArnsRequest}
}

// ListAppliedSchemaArnsRequest is the request type for the
// ListAppliedSchemaArns API operation.
type ListAppliedSchemaArnsRequest struct {
	*aws.Request
	Input *ListAppliedSchemaArnsInput
	Copy  func(*ListAppliedSchemaArnsInput) ListAppliedSchemaArnsRequest
}

// Send marshals and sends the ListAppliedSchemaArns API request.
func (r ListAppliedSchemaArnsRequest) Send(ctx context.Context) (*ListAppliedSchemaArnsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListAppliedSchemaArnsResponse{
		ListAppliedSchemaArnsOutput: r.Request.Data.(*ListAppliedSchemaArnsOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListAppliedSchemaArnsRequestPaginator returns a paginator for ListAppliedSchemaArns.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListAppliedSchemaArnsRequest(input)
//   p := clouddirectory.NewListAppliedSchemaArnsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListAppliedSchemaArnsPaginator(req ListAppliedSchemaArnsRequest) ListAppliedSchemaArnsPaginator {
	return ListAppliedSchemaArnsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListAppliedSchemaArnsInput
				if req.Input != nil {
					tmp := *req.Input
					inCpy = &tmp
				}

				newReq := req.Copy(inCpy)
				newReq.SetContext(ctx)
				return newReq.Request, nil
			},
		},
	}
}

// ListAppliedSchemaArnsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListAppliedSchemaArnsPaginator struct {
	aws.Pager
}

func (p *ListAppliedSchemaArnsPaginator) CurrentPage() *ListAppliedSchemaArnsOutput {
	return p.Pager.CurrentPage().(*ListAppliedSchemaArnsOutput)
}

// ListAppliedSchemaArnsResponse is the response type for the
// ListAppliedSchemaArns API operation.
type ListAppliedSchemaArnsResponse struct {
	*ListAppliedSchemaArnsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListAppliedSchemaArns request.
func (r *ListAppliedSchemaArnsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
