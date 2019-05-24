// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ram

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ram-2018-01-04/ListPrincipalsRequest
type ListPrincipalsInput struct {
	_ struct{} `type:"structure"`

	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	MaxResults *int64 `locationName:"maxResults" min:"1" type:"integer"`

	// The token for the next page of results.
	NextToken *string `locationName:"nextToken" type:"string"`

	// The principals.
	Principals []string `locationName:"principals" type:"list"`

	// The Amazon Resource Name (ARN) of the resource.
	ResourceArn *string `locationName:"resourceArn" type:"string"`

	// The type of owner.
	//
	// ResourceOwner is a required field
	ResourceOwner ResourceOwner `locationName:"resourceOwner" type:"string" required:"true" enum:"true"`

	// The Amazon Resource Names (ARN) of the resource shares.
	ResourceShareArns []string `locationName:"resourceShareArns" type:"list"`

	// The resource type.
	ResourceType *string `locationName:"resourceType" type:"string"`
}

// String returns the string representation
func (s ListPrincipalsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListPrincipalsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListPrincipalsInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}
	if len(s.ResourceOwner) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("ResourceOwner"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListPrincipalsInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	if s.MaxResults != nil {
		v := *s.MaxResults

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "maxResults", protocol.Int64Value(v), metadata)
	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.Principals) > 0 {
		v := s.Principals

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "principals", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	if s.ResourceArn != nil {
		v := *s.ResourceArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "resourceArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.ResourceOwner) > 0 {
		v := s.ResourceOwner

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "resourceOwner", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if len(s.ResourceShareArns) > 0 {
		v := s.ResourceShareArns

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "resourceShareArns", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	if s.ResourceType != nil {
		v := *s.ResourceType

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "resourceType", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ram-2018-01-04/ListPrincipalsResponse
type ListPrincipalsOutput struct {
	_ struct{} `type:"structure"`

	// The token to use to retrieve the next page of results. This value is null
	// when there are no more results to return.
	NextToken *string `locationName:"nextToken" type:"string"`

	// The principals.
	Principals []Principal `locationName:"principals" type:"list"`
}

// String returns the string representation
func (s ListPrincipalsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListPrincipalsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.Principals) > 0 {
		v := s.Principals

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "principals", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	return nil
}

const opListPrincipals = "ListPrincipals"

// ListPrincipalsRequest returns a request value for making API operation for
// AWS Resource Access Manager.
//
// Lists the principals with access to the specified resource.
//
//    // Example sending a request using ListPrincipalsRequest.
//    req := client.ListPrincipalsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ram-2018-01-04/ListPrincipals
func (c *Client) ListPrincipalsRequest(input *ListPrincipalsInput) ListPrincipalsRequest {
	op := &aws.Operation{
		Name:       opListPrincipals,
		HTTPMethod: "POST",
		HTTPPath:   "/listprincipals",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"nextToken"},
			OutputTokens:    []string{"nextToken"},
			LimitToken:      "maxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListPrincipalsInput{}
	}

	req := c.newRequest(op, input, &ListPrincipalsOutput{})
	return ListPrincipalsRequest{Request: req, Input: input, Copy: c.ListPrincipalsRequest}
}

// ListPrincipalsRequest is the request type for the
// ListPrincipals API operation.
type ListPrincipalsRequest struct {
	*aws.Request
	Input *ListPrincipalsInput
	Copy  func(*ListPrincipalsInput) ListPrincipalsRequest
}

// Send marshals and sends the ListPrincipals API request.
func (r ListPrincipalsRequest) Send(ctx context.Context) (*ListPrincipalsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListPrincipalsResponse{
		ListPrincipalsOutput: r.Request.Data.(*ListPrincipalsOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListPrincipalsRequestPaginator returns a paginator for ListPrincipals.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListPrincipalsRequest(input)
//   p := ram.NewListPrincipalsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListPrincipalsPaginator(req ListPrincipalsRequest) ListPrincipalsPaginator {
	return ListPrincipalsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListPrincipalsInput
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

// ListPrincipalsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListPrincipalsPaginator struct {
	aws.Pager
}

func (p *ListPrincipalsPaginator) CurrentPage() *ListPrincipalsOutput {
	return p.Pager.CurrentPage().(*ListPrincipalsOutput)
}

// ListPrincipalsResponse is the response type for the
// ListPrincipals API operation.
type ListPrincipalsResponse struct {
	*ListPrincipalsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListPrincipals request.
func (r *ListPrincipalsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
