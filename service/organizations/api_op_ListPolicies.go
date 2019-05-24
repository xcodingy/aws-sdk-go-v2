// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package organizations

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/organizations-2016-11-28/ListPoliciesRequest
type ListPoliciesInput struct {
	_ struct{} `type:"structure"`

	// Specifies the type of policy that you want to include in the response.
	//
	// Filter is a required field
	Filter PolicyType `type:"string" required:"true" enum:"true"`

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
}

// String returns the string representation
func (s ListPoliciesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListPoliciesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListPoliciesInput"}
	if len(s.Filter) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("Filter"))
	}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/organizations-2016-11-28/ListPoliciesResponse
type ListPoliciesOutput struct {
	_ struct{} `type:"structure"`

	// If present, this value indicates that there is more output available than
	// is included in the current response. Use this value in the NextToken request
	// parameter in a subsequent call to the operation to get the next part of the
	// output. You should repeat this until the NextToken response element comes
	// back as null.
	NextToken *string `type:"string"`

	// A list of policies that match the filter criteria in the request. The output
	// list does not include the policy contents. To see the content for a policy,
	// see DescribePolicy.
	Policies []PolicySummary `type:"list"`
}

// String returns the string representation
func (s ListPoliciesOutput) String() string {
	return awsutil.Prettify(s)
}

const opListPolicies = "ListPolicies"

// ListPoliciesRequest returns a request value for making API operation for
// AWS Organizations.
//
// Retrieves the list of all policies in an organization of a specified type.
//
// Always check the NextToken response parameter for a null value when calling
// a List* operation. These operations can occasionally return an empty set
// of results even when there are more results available. The NextToken response
// parameter value is nullonly when there are no more results to display.
//
// This operation can be called only from the organization's master account.
//
//    // Example sending a request using ListPoliciesRequest.
//    req := client.ListPoliciesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/organizations-2016-11-28/ListPolicies
func (c *Client) ListPoliciesRequest(input *ListPoliciesInput) ListPoliciesRequest {
	op := &aws.Operation{
		Name:       opListPolicies,
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
		input = &ListPoliciesInput{}
	}

	req := c.newRequest(op, input, &ListPoliciesOutput{})
	return ListPoliciesRequest{Request: req, Input: input, Copy: c.ListPoliciesRequest}
}

// ListPoliciesRequest is the request type for the
// ListPolicies API operation.
type ListPoliciesRequest struct {
	*aws.Request
	Input *ListPoliciesInput
	Copy  func(*ListPoliciesInput) ListPoliciesRequest
}

// Send marshals and sends the ListPolicies API request.
func (r ListPoliciesRequest) Send(ctx context.Context) (*ListPoliciesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListPoliciesResponse{
		ListPoliciesOutput: r.Request.Data.(*ListPoliciesOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListPoliciesRequestPaginator returns a paginator for ListPolicies.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListPoliciesRequest(input)
//   p := organizations.NewListPoliciesRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListPoliciesPaginator(req ListPoliciesRequest) ListPoliciesPaginator {
	return ListPoliciesPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListPoliciesInput
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

// ListPoliciesPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListPoliciesPaginator struct {
	aws.Pager
}

func (p *ListPoliciesPaginator) CurrentPage() *ListPoliciesOutput {
	return p.Pager.CurrentPage().(*ListPoliciesOutput)
}

// ListPoliciesResponse is the response type for the
// ListPolicies API operation.
type ListPoliciesResponse struct {
	*ListPoliciesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListPolicies request.
func (r *ListPoliciesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
