// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package mobile

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Request structure used to request projects list in AWS Mobile Hub.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/mobile-2017-07-01/ListProjectsRequest
type ListProjectsInput struct {
	_ struct{} `type:"structure"`

	// Maximum number of records to list in a single response.
	MaxResults *int64 `location:"querystring" locationName:"maxResults" type:"integer"`

	// Pagination token. Set to null to start listing projects from start. If non-null
	// pagination token is returned in a result, then pass its value in here in
	// another request to list more projects.
	NextToken *string `location:"querystring" locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s ListProjectsInput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListProjectsInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	if s.MaxResults != nil {
		v := *s.MaxResults

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "maxResults", protocol.Int64Value(v), metadata)
	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Result structure used for requests to list projects in AWS Mobile Hub.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/mobile-2017-07-01/ListProjectsResult
type ListProjectsOutput struct {
	_ struct{} `type:"structure"`

	// Pagination token. Set to null to start listing records from start. If non-null
	// pagination token is returned in a result, then pass its value in here in
	// another request to list more entries.
	NextToken *string `locationName:"nextToken" type:"string"`

	// List of projects.
	Projects []ProjectSummary `locationName:"projects" type:"list"`
}

// String returns the string representation
func (s ListProjectsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListProjectsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.Projects) > 0 {
		v := s.Projects

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "projects", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	return nil
}

const opListProjects = "ListProjects"

// ListProjectsRequest returns a request value for making API operation for
// AWS Mobile.
//
// Lists projects in AWS Mobile Hub.
//
//    // Example sending a request using ListProjectsRequest.
//    req := client.ListProjectsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/mobile-2017-07-01/ListProjects
func (c *Client) ListProjectsRequest(input *ListProjectsInput) ListProjectsRequest {
	op := &aws.Operation{
		Name:       opListProjects,
		HTTPMethod: "GET",
		HTTPPath:   "/projects",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"nextToken"},
			OutputTokens:    []string{"nextToken"},
			LimitToken:      "maxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListProjectsInput{}
	}

	req := c.newRequest(op, input, &ListProjectsOutput{})
	return ListProjectsRequest{Request: req, Input: input, Copy: c.ListProjectsRequest}
}

// ListProjectsRequest is the request type for the
// ListProjects API operation.
type ListProjectsRequest struct {
	*aws.Request
	Input *ListProjectsInput
	Copy  func(*ListProjectsInput) ListProjectsRequest
}

// Send marshals and sends the ListProjects API request.
func (r ListProjectsRequest) Send(ctx context.Context) (*ListProjectsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListProjectsResponse{
		ListProjectsOutput: r.Request.Data.(*ListProjectsOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListProjectsRequestPaginator returns a paginator for ListProjects.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListProjectsRequest(input)
//   p := mobile.NewListProjectsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListProjectsPaginator(req ListProjectsRequest) ListProjectsPaginator {
	return ListProjectsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListProjectsInput
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

// ListProjectsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListProjectsPaginator struct {
	aws.Pager
}

func (p *ListProjectsPaginator) CurrentPage() *ListProjectsOutput {
	return p.Pager.CurrentPage().(*ListProjectsOutput)
}

// ListProjectsResponse is the response type for the
// ListProjects API operation.
type ListProjectsResponse struct {
	*ListProjectsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListProjects request.
func (r *ListProjectsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
