// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package organizations

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/organizations-2016-11-28/ListAccountsForParentRequest
type ListAccountsForParentInput struct {
	_ struct{} `type:"structure"`

	// (Optional) Use this to limit the number of results you want included per
	// page in the response. If you do not include this parameter, it defaults to
	// a value that is specific to the operation. If additional items exist beyond
	// the maximum you specify, the NextToken response element is present and has
	// a value (is not null). Include that value as the NextToken request parameter
	// in the next call to the operation to get the next part of the results. Note
	// that Organizations might return fewer results than the maximum even when
	// there are more results available. You should check NextToken after every
	// operation to ensure that you receive all of the results.
	MaxResults *int64 `min:"1" type:"integer"`

	// Use this parameter if you receive a NextToken response in a previous request
	// that indicates that there is more output available. Set it to the value of
	// the previous call's NextToken response to indicate where the output should
	// continue from.
	NextToken *string `type:"string"`

	// The unique identifier (ID) for the parent root or organization unit (OU)
	// whose accounts you want to list.
	//
	// ParentId is a required field
	ParentId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s ListAccountsForParentInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListAccountsForParentInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListAccountsForParentInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if s.ParentId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ParentId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/organizations-2016-11-28/ListAccountsForParentResponse
type ListAccountsForParentOutput struct {
	_ struct{} `type:"structure"`

	// A list of the accounts in the specified root or OU.
	Accounts []Account `type:"list"`

	// If present, this value indicates that there is more output available than
	// is included in the current response. Use this value in the NextToken request
	// parameter in a subsequent call to the operation to get the next part of the
	// output. You should repeat this until the NextToken response element comes
	// back as null.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s ListAccountsForParentOutput) String() string {
	return awsutil.Prettify(s)
}

const opListAccountsForParent = "ListAccountsForParent"

// ListAccountsForParentRequest returns a request value for making API operation for
// AWS Organizations.
//
// Lists the accounts in an organization that are contained by the specified
// target root or organizational unit (OU). If you specify the root, you get
// a list of all the accounts that are not in any OU. If you specify an OU,
// you get a list of all the accounts in only that OU, and not in any child
// OUs. To get a list of all accounts in the organization, use the ListAccounts
// operation.
//
// Always check the NextToken response parameter for a null value when calling
// a List* operation. These operations can occasionally return an empty set
// of results even when there are more results available. The NextToken response
// parameter value is nullonly when there are no more results to display.
//
// This operation can be called only from the organization's master account.
//
//    // Example sending a request using ListAccountsForParentRequest.
//    req := client.ListAccountsForParentRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/organizations-2016-11-28/ListAccountsForParent
func (c *Client) ListAccountsForParentRequest(input *ListAccountsForParentInput) ListAccountsForParentRequest {
	op := &aws.Operation{
		Name:       opListAccountsForParent,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListAccountsForParentInput{}
	}

	req := c.newRequest(op, input, &ListAccountsForParentOutput{})
	return ListAccountsForParentRequest{Request: req, Input: input, Copy: c.ListAccountsForParentRequest}
}

// ListAccountsForParentRequest is the request type for the
// ListAccountsForParent API operation.
type ListAccountsForParentRequest struct {
	*aws.Request
	Input *ListAccountsForParentInput
	Copy  func(*ListAccountsForParentInput) ListAccountsForParentRequest
}

// Send marshals and sends the ListAccountsForParent API request.
func (r ListAccountsForParentRequest) Send(ctx context.Context) (*ListAccountsForParentResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListAccountsForParentResponse{
		ListAccountsForParentOutput: r.Request.Data.(*ListAccountsForParentOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListAccountsForParentRequestPaginator returns a paginator for ListAccountsForParent.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListAccountsForParentRequest(input)
//   p := organizations.NewListAccountsForParentRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListAccountsForParentPaginator(req ListAccountsForParentRequest) ListAccountsForParentPaginator {
	return ListAccountsForParentPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListAccountsForParentInput
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

// ListAccountsForParentPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListAccountsForParentPaginator struct {
	aws.Pager
}

func (p *ListAccountsForParentPaginator) CurrentPage() *ListAccountsForParentOutput {
	return p.Pager.CurrentPage().(*ListAccountsForParentOutput)
}

// ListAccountsForParentResponse is the response type for the
// ListAccountsForParent API operation.
type ListAccountsForParentResponse struct {
	*ListAccountsForParentOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListAccountsForParent request.
func (r *ListAccountsForParentResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
