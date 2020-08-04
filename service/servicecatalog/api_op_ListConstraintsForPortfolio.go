// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package servicecatalog

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ListConstraintsForPortfolioInput struct {
	_ struct{} `type:"structure"`

	// The language code.
	//
	//    * en - English (default)
	//
	//    * jp - Japanese
	//
	//    * zh - Chinese
	AcceptLanguage *string `type:"string"`

	// The maximum number of items to return with this call.
	PageSize *int64 `type:"integer"`

	// The page token for the next set of results. To retrieve the first set of
	// results, use null.
	PageToken *string `type:"string"`

	// The portfolio identifier.
	//
	// PortfolioId is a required field
	PortfolioId *string `min:"1" type:"string" required:"true"`

	// The product identifier.
	ProductId *string `min:"1" type:"string"`
}

// String returns the string representation
func (s ListConstraintsForPortfolioInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListConstraintsForPortfolioInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListConstraintsForPortfolioInput"}

	if s.PortfolioId == nil {
		invalidParams.Add(aws.NewErrParamRequired("PortfolioId"))
	}
	if s.PortfolioId != nil && len(*s.PortfolioId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("PortfolioId", 1))
	}
	if s.ProductId != nil && len(*s.ProductId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ProductId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ListConstraintsForPortfolioOutput struct {
	_ struct{} `type:"structure"`

	// Information about the constraints.
	ConstraintDetails []ConstraintDetail `type:"list"`

	// The page token to use to retrieve the next set of results. If there are no
	// additional results, this value is null.
	NextPageToken *string `type:"string"`
}

// String returns the string representation
func (s ListConstraintsForPortfolioOutput) String() string {
	return awsutil.Prettify(s)
}

const opListConstraintsForPortfolio = "ListConstraintsForPortfolio"

// ListConstraintsForPortfolioRequest returns a request value for making API operation for
// AWS Service Catalog.
//
// Lists the constraints for the specified portfolio and product.
//
//    // Example sending a request using ListConstraintsForPortfolioRequest.
//    req := client.ListConstraintsForPortfolioRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/servicecatalog-2015-12-10/ListConstraintsForPortfolio
func (c *Client) ListConstraintsForPortfolioRequest(input *ListConstraintsForPortfolioInput) ListConstraintsForPortfolioRequest {
	op := &aws.Operation{
		Name:       opListConstraintsForPortfolio,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"PageToken"},
			OutputTokens:    []string{"NextPageToken"},
			LimitToken:      "PageSize",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListConstraintsForPortfolioInput{}
	}

	req := c.newRequest(op, input, &ListConstraintsForPortfolioOutput{})

	return ListConstraintsForPortfolioRequest{Request: req, Input: input, Copy: c.ListConstraintsForPortfolioRequest}
}

// ListConstraintsForPortfolioRequest is the request type for the
// ListConstraintsForPortfolio API operation.
type ListConstraintsForPortfolioRequest struct {
	*aws.Request
	Input *ListConstraintsForPortfolioInput
	Copy  func(*ListConstraintsForPortfolioInput) ListConstraintsForPortfolioRequest
}

// Send marshals and sends the ListConstraintsForPortfolio API request.
func (r ListConstraintsForPortfolioRequest) Send(ctx context.Context) (*ListConstraintsForPortfolioResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListConstraintsForPortfolioResponse{
		ListConstraintsForPortfolioOutput: r.Request.Data.(*ListConstraintsForPortfolioOutput),
		response:                          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListConstraintsForPortfolioRequestPaginator returns a paginator for ListConstraintsForPortfolio.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListConstraintsForPortfolioRequest(input)
//   p := servicecatalog.NewListConstraintsForPortfolioRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListConstraintsForPortfolioPaginator(req ListConstraintsForPortfolioRequest) ListConstraintsForPortfolioPaginator {
	return ListConstraintsForPortfolioPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListConstraintsForPortfolioInput
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

// ListConstraintsForPortfolioPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListConstraintsForPortfolioPaginator struct {
	aws.Pager
}

func (p *ListConstraintsForPortfolioPaginator) CurrentPage() *ListConstraintsForPortfolioOutput {
	return p.Pager.CurrentPage().(*ListConstraintsForPortfolioOutput)
}

// ListConstraintsForPortfolioResponse is the response type for the
// ListConstraintsForPortfolio API operation.
type ListConstraintsForPortfolioResponse struct {
	*ListConstraintsForPortfolioOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListConstraintsForPortfolio request.
func (r *ListConstraintsForPortfolioResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}