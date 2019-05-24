// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sagemaker

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/sagemaker-2017-07-24/ListModelsInput
type ListModelsInput struct {
	_ struct{} `type:"structure"`

	// A filter that returns only models created after the specified time (timestamp).
	CreationTimeAfter *time.Time `type:"timestamp" timestampFormat:"unix"`

	// A filter that returns only models created before the specified time (timestamp).
	CreationTimeBefore *time.Time `type:"timestamp" timestampFormat:"unix"`

	// The maximum number of models to return in the response.
	MaxResults *int64 `min:"1" type:"integer"`

	// A string in the training job name. This filter returns only models in the
	// training job whose name contains the specified string.
	NameContains *string `type:"string"`

	// If the response to a previous ListModels request was truncated, the response
	// includes a NextToken. To retrieve the next set of models, use the token in
	// the next request.
	NextToken *string `type:"string"`

	// Sorts the list of results. The default is CreationTime.
	SortBy ModelSortKey `type:"string" enum:"true"`

	// The sort order for results. The default is Descending.
	SortOrder OrderKey `type:"string" enum:"true"`
}

// String returns the string representation
func (s ListModelsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListModelsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListModelsInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/sagemaker-2017-07-24/ListModelsOutput
type ListModelsOutput struct {
	_ struct{} `type:"structure"`

	// An array of ModelSummary objects, each of which lists a model.
	//
	// Models is a required field
	Models []ModelSummary `type:"list" required:"true"`

	// If the response is truncated, Amazon SageMaker returns this token. To retrieve
	// the next set of models, use it in the subsequent request.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s ListModelsOutput) String() string {
	return awsutil.Prettify(s)
}

const opListModels = "ListModels"

// ListModelsRequest returns a request value for making API operation for
// Amazon SageMaker Service.
//
// Lists models created with the CreateModel (https://docs.aws.amazon.com/sagemaker/latest/dg/API_CreateModel.html)
// API.
//
//    // Example sending a request using ListModelsRequest.
//    req := client.ListModelsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/sagemaker-2017-07-24/ListModels
func (c *Client) ListModelsRequest(input *ListModelsInput) ListModelsRequest {
	op := &aws.Operation{
		Name:       opListModels,
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
		input = &ListModelsInput{}
	}

	req := c.newRequest(op, input, &ListModelsOutput{})
	return ListModelsRequest{Request: req, Input: input, Copy: c.ListModelsRequest}
}

// ListModelsRequest is the request type for the
// ListModels API operation.
type ListModelsRequest struct {
	*aws.Request
	Input *ListModelsInput
	Copy  func(*ListModelsInput) ListModelsRequest
}

// Send marshals and sends the ListModels API request.
func (r ListModelsRequest) Send(ctx context.Context) (*ListModelsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListModelsResponse{
		ListModelsOutput: r.Request.Data.(*ListModelsOutput),
		response:         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListModelsRequestPaginator returns a paginator for ListModels.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListModelsRequest(input)
//   p := sagemaker.NewListModelsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListModelsPaginator(req ListModelsRequest) ListModelsPaginator {
	return ListModelsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListModelsInput
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

// ListModelsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListModelsPaginator struct {
	aws.Pager
}

func (p *ListModelsPaginator) CurrentPage() *ListModelsOutput {
	return p.Pager.CurrentPage().(*ListModelsOutput)
}

// ListModelsResponse is the response type for the
// ListModels API operation.
type ListModelsResponse struct {
	*ListModelsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListModels request.
func (r *ListModelsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
