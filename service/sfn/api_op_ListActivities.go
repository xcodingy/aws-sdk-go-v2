// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sfn

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/states-2016-11-23/ListActivitiesInput
type ListActivitiesInput struct {
	_ struct{} `type:"structure"`

	// The maximum number of results that are returned per call. You can use nextToken
	// to obtain further pages of results. The default is 100 and the maximum allowed
	// page size is 1000. A value of 0 uses the default.
	//
	// This is only an upper limit. The actual number of results returned per call
	// might be fewer than the specified maximum.
	MaxResults *int64 `locationName:"maxResults" type:"integer"`

	// If nextToken is returned, there are more results available. The value of
	// nextToken is a unique pagination token for each page. Make the call again
	// using the returned token to retrieve the next page. Keep all other arguments
	// unchanged. Each pagination token expires after 24 hours. Using an expired
	// pagination token will return an HTTP 400 InvalidToken error.
	NextToken *string `locationName:"nextToken" min:"1" type:"string"`
}

// String returns the string representation
func (s ListActivitiesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListActivitiesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListActivitiesInput"}
	if s.NextToken != nil && len(*s.NextToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NextToken", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/states-2016-11-23/ListActivitiesOutput
type ListActivitiesOutput struct {
	_ struct{} `type:"structure"`

	// The list of activities.
	//
	// Activities is a required field
	Activities []ActivityListItem `locationName:"activities" type:"list" required:"true"`

	// If nextToken is returned, there are more results available. The value of
	// nextToken is a unique pagination token for each page. Make the call again
	// using the returned token to retrieve the next page. Keep all other arguments
	// unchanged. Each pagination token expires after 24 hours. Using an expired
	// pagination token will return an HTTP 400 InvalidToken error.
	NextToken *string `locationName:"nextToken" min:"1" type:"string"`
}

// String returns the string representation
func (s ListActivitiesOutput) String() string {
	return awsutil.Prettify(s)
}

const opListActivities = "ListActivities"

// ListActivitiesRequest returns a request value for making API operation for
// AWS Step Functions.
//
// Lists the existing activities.
//
// If nextToken is returned, there are more results available. The value of
// nextToken is a unique pagination token for each page. Make the call again
// using the returned token to retrieve the next page. Keep all other arguments
// unchanged. Each pagination token expires after 24 hours. Using an expired
// pagination token will return an HTTP 400 InvalidToken error.
//
// This operation is eventually consistent. The results are best effort and
// may not reflect very recent updates and changes.
//
//    // Example sending a request using ListActivitiesRequest.
//    req := client.ListActivitiesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/states-2016-11-23/ListActivities
func (c *Client) ListActivitiesRequest(input *ListActivitiesInput) ListActivitiesRequest {
	op := &aws.Operation{
		Name:       opListActivities,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"nextToken"},
			OutputTokens:    []string{"nextToken"},
			LimitToken:      "maxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListActivitiesInput{}
	}

	req := c.newRequest(op, input, &ListActivitiesOutput{})
	return ListActivitiesRequest{Request: req, Input: input, Copy: c.ListActivitiesRequest}
}

// ListActivitiesRequest is the request type for the
// ListActivities API operation.
type ListActivitiesRequest struct {
	*aws.Request
	Input *ListActivitiesInput
	Copy  func(*ListActivitiesInput) ListActivitiesRequest
}

// Send marshals and sends the ListActivities API request.
func (r ListActivitiesRequest) Send(ctx context.Context) (*ListActivitiesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListActivitiesResponse{
		ListActivitiesOutput: r.Request.Data.(*ListActivitiesOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListActivitiesRequestPaginator returns a paginator for ListActivities.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListActivitiesRequest(input)
//   p := sfn.NewListActivitiesRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListActivitiesPaginator(req ListActivitiesRequest) ListActivitiesPaginator {
	return ListActivitiesPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListActivitiesInput
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

// ListActivitiesPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListActivitiesPaginator struct {
	aws.Pager
}

func (p *ListActivitiesPaginator) CurrentPage() *ListActivitiesOutput {
	return p.Pager.CurrentPage().(*ListActivitiesOutput)
}

// ListActivitiesResponse is the response type for the
// ListActivities API operation.
type ListActivitiesResponse struct {
	*ListActivitiesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListActivities request.
func (r *ListActivitiesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
