// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package servicecatalog

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/servicecatalog-2015-12-10/SearchProductsInput
type SearchProductsInput struct {
	_ struct{} `type:"structure"`

	// The language code.
	//
	//    * en - English (default)
	//
	//    * jp - Japanese
	//
	//    * zh - Chinese
	AcceptLanguage *string `type:"string"`

	// The search filters. If no search filters are specified, the output includes
	// all products to which the caller has access.
	Filters map[string][]string `type:"map"`

	// The maximum number of items to return with this call.
	PageSize *int64 `type:"integer"`

	// The page token for the next set of results. To retrieve the first set of
	// results, use null.
	PageToken *string `type:"string"`

	// The sort field. If no value is specified, the results are not sorted.
	SortBy ProductViewSortBy `type:"string" enum:"true"`

	// The sort order. If no value is specified, the results are not sorted.
	SortOrder SortOrder `type:"string" enum:"true"`
}

// String returns the string representation
func (s SearchProductsInput) String() string {
	return awsutil.Prettify(s)
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/servicecatalog-2015-12-10/SearchProductsOutput
type SearchProductsOutput struct {
	_ struct{} `type:"structure"`

	// The page token to use to retrieve the next set of results. If there are no
	// additional results, this value is null.
	NextPageToken *string `type:"string"`

	// The product view aggregations.
	ProductViewAggregations map[string][]ProductViewAggregationValue `type:"map"`

	// Information about the product views.
	ProductViewSummaries []ProductViewSummary `type:"list"`
}

// String returns the string representation
func (s SearchProductsOutput) String() string {
	return awsutil.Prettify(s)
}

const opSearchProducts = "SearchProducts"

// SearchProductsRequest returns a request value for making API operation for
// AWS Service Catalog.
//
// Gets information about the products to which the caller has access.
//
//    // Example sending a request using SearchProductsRequest.
//    req := client.SearchProductsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/servicecatalog-2015-12-10/SearchProducts
func (c *Client) SearchProductsRequest(input *SearchProductsInput) SearchProductsRequest {
	op := &aws.Operation{
		Name:       opSearchProducts,
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
		input = &SearchProductsInput{}
	}

	req := c.newRequest(op, input, &SearchProductsOutput{})
	return SearchProductsRequest{Request: req, Input: input, Copy: c.SearchProductsRequest}
}

// SearchProductsRequest is the request type for the
// SearchProducts API operation.
type SearchProductsRequest struct {
	*aws.Request
	Input *SearchProductsInput
	Copy  func(*SearchProductsInput) SearchProductsRequest
}

// Send marshals and sends the SearchProducts API request.
func (r SearchProductsRequest) Send(ctx context.Context) (*SearchProductsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &SearchProductsResponse{
		SearchProductsOutput: r.Request.Data.(*SearchProductsOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewSearchProductsRequestPaginator returns a paginator for SearchProducts.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.SearchProductsRequest(input)
//   p := servicecatalog.NewSearchProductsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewSearchProductsPaginator(req SearchProductsRequest) SearchProductsPaginator {
	return SearchProductsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *SearchProductsInput
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

// SearchProductsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type SearchProductsPaginator struct {
	aws.Pager
}

func (p *SearchProductsPaginator) CurrentPage() *SearchProductsOutput {
	return p.Pager.CurrentPage().(*SearchProductsOutput)
}

// SearchProductsResponse is the response type for the
// SearchProducts API operation.
type SearchProductsResponse struct {
	*SearchProductsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// SearchProducts request.
func (r *SearchProductsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
