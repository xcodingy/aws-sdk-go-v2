// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package chime

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/chime-2018-05-01/ListAccountsRequest
type ListAccountsInput struct {
	_ struct{} `type:"structure"`

	// The maximum number of results to return in a single call. Defaults to 100.
	MaxResults *int64 `location:"querystring" locationName:"max-results" min:"1" type:"integer"`

	// Amazon Chime account name prefix with which to filter results.
	Name *string `location:"querystring" locationName:"name" min:"1" type:"string"`

	// The token to use to retrieve the next page of results.
	NextToken *string `location:"querystring" locationName:"next-token" type:"string"`

	// User email address with which to filter results.
	UserEmail *string `location:"querystring" locationName:"user-email" type:"string"`
}

// String returns the string representation
func (s ListAccountsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListAccountsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListAccountsInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}
	if s.Name != nil && len(*s.Name) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Name", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListAccountsInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.MaxResults != nil {
		v := *s.MaxResults

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "max-results", protocol.Int64Value(v), metadata)
	}
	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "next-token", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.UserEmail != nil {
		v := *s.UserEmail

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "user-email", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/chime-2018-05-01/ListAccountsResponse
type ListAccountsOutput struct {
	_ struct{} `type:"structure"`

	// List of Amazon Chime accounts and account details.
	Accounts []Account `type:"list"`

	// The token to use to retrieve the next page of results.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s ListAccountsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListAccountsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if len(s.Accounts) > 0 {
		v := s.Accounts

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "Accounts", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "NextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opListAccounts = "ListAccounts"

// ListAccountsRequest returns a request value for making API operation for
// Amazon Chime.
//
// Lists the Amazon Chime accounts under the administrator's AWS account. You
// can filter accounts by account name prefix. To find out which Amazon Chime
// account a user belongs to, you can filter by the user's email address, which
// returns one account result.
//
//    // Example sending a request using ListAccountsRequest.
//    req := client.ListAccountsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/chime-2018-05-01/ListAccounts
func (c *Client) ListAccountsRequest(input *ListAccountsInput) ListAccountsRequest {
	op := &aws.Operation{
		Name:       opListAccounts,
		HTTPMethod: "GET",
		HTTPPath:   "/accounts",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListAccountsInput{}
	}

	req := c.newRequest(op, input, &ListAccountsOutput{})
	return ListAccountsRequest{Request: req, Input: input, Copy: c.ListAccountsRequest}
}

// ListAccountsRequest is the request type for the
// ListAccounts API operation.
type ListAccountsRequest struct {
	*aws.Request
	Input *ListAccountsInput
	Copy  func(*ListAccountsInput) ListAccountsRequest
}

// Send marshals and sends the ListAccounts API request.
func (r ListAccountsRequest) Send(ctx context.Context) (*ListAccountsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListAccountsResponse{
		ListAccountsOutput: r.Request.Data.(*ListAccountsOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListAccountsRequestPaginator returns a paginator for ListAccounts.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListAccountsRequest(input)
//   p := chime.NewListAccountsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListAccountsPaginator(req ListAccountsRequest) ListAccountsPaginator {
	return ListAccountsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListAccountsInput
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

// ListAccountsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListAccountsPaginator struct {
	aws.Pager
}

func (p *ListAccountsPaginator) CurrentPage() *ListAccountsOutput {
	return p.Pager.CurrentPage().(*ListAccountsOutput)
}

// ListAccountsResponse is the response type for the
// ListAccounts API operation.
type ListAccountsResponse struct {
	*ListAccountsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListAccounts request.
func (r *ListAccountsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
