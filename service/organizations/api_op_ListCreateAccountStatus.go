// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package organizations

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/organizations-2016-11-28/ListCreateAccountStatusRequest
type ListCreateAccountStatusInput struct {
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

	// A list of one or more states that you want included in the response. If this
	// parameter is not present, then all requests are included in the response.
	States []CreateAccountState `type:"list"`
}

// String returns the string representation
func (s ListCreateAccountStatusInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListCreateAccountStatusInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListCreateAccountStatusInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/organizations-2016-11-28/ListCreateAccountStatusResponse
type ListCreateAccountStatusOutput struct {
	_ struct{} `type:"structure"`

	// A list of objects with details about the requests. Certain elements, such
	// as the accountId number, are present in the output only after the account
	// has been successfully created.
	CreateAccountStatuses []CreateAccountStatus `type:"list"`

	// If present, this value indicates that there is more output available than
	// is included in the current response. Use this value in the NextToken request
	// parameter in a subsequent call to the operation to get the next part of the
	// output. You should repeat this until the NextToken response element comes
	// back as null.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s ListCreateAccountStatusOutput) String() string {
	return awsutil.Prettify(s)
}

const opListCreateAccountStatus = "ListCreateAccountStatus"

// ListCreateAccountStatusRequest returns a request value for making API operation for
// AWS Organizations.
//
// Lists the account creation requests that match the specified status that
// is currently being tracked for the organization.
//
// Always check the NextToken response parameter for a null value when calling
// a List* operation. These operations can occasionally return an empty set
// of results even when there are more results available. The NextToken response
// parameter value is nullonly when there are no more results to display.
//
// This operation can be called only from the organization's master account.
//
//    // Example sending a request using ListCreateAccountStatusRequest.
//    req := client.ListCreateAccountStatusRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/organizations-2016-11-28/ListCreateAccountStatus
func (c *Client) ListCreateAccountStatusRequest(input *ListCreateAccountStatusInput) ListCreateAccountStatusRequest {
	op := &aws.Operation{
		Name:       opListCreateAccountStatus,
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
		input = &ListCreateAccountStatusInput{}
	}

	req := c.newRequest(op, input, &ListCreateAccountStatusOutput{})
	return ListCreateAccountStatusRequest{Request: req, Input: input, Copy: c.ListCreateAccountStatusRequest}
}

// ListCreateAccountStatusRequest is the request type for the
// ListCreateAccountStatus API operation.
type ListCreateAccountStatusRequest struct {
	*aws.Request
	Input *ListCreateAccountStatusInput
	Copy  func(*ListCreateAccountStatusInput) ListCreateAccountStatusRequest
}

// Send marshals and sends the ListCreateAccountStatus API request.
func (r ListCreateAccountStatusRequest) Send(ctx context.Context) (*ListCreateAccountStatusResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListCreateAccountStatusResponse{
		ListCreateAccountStatusOutput: r.Request.Data.(*ListCreateAccountStatusOutput),
		response:                      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListCreateAccountStatusRequestPaginator returns a paginator for ListCreateAccountStatus.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListCreateAccountStatusRequest(input)
//   p := organizations.NewListCreateAccountStatusRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListCreateAccountStatusPaginator(req ListCreateAccountStatusRequest) ListCreateAccountStatusPaginator {
	return ListCreateAccountStatusPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListCreateAccountStatusInput
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

// ListCreateAccountStatusPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListCreateAccountStatusPaginator struct {
	aws.Pager
}

func (p *ListCreateAccountStatusPaginator) CurrentPage() *ListCreateAccountStatusOutput {
	return p.Pager.CurrentPage().(*ListCreateAccountStatusOutput)
}

// ListCreateAccountStatusResponse is the response type for the
// ListCreateAccountStatus API operation.
type ListCreateAccountStatusResponse struct {
	*ListCreateAccountStatusOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListCreateAccountStatus request.
func (r *ListCreateAccountStatusResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
