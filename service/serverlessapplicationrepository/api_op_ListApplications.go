// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package serverlessapplicationrepository

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/serverlessrepo-2017-09-08/ListApplicationsRequest
type ListApplicationsInput struct {
	_ struct{} `type:"structure"`

	MaxItems *int64 `location:"querystring" locationName:"maxItems" min:"1" type:"integer"`

	NextToken *string `location:"querystring" locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s ListApplicationsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListApplicationsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListApplicationsInput"}
	if s.MaxItems != nil && *s.MaxItems < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxItems", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListApplicationsInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	if s.MaxItems != nil {
		v := *s.MaxItems

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "maxItems", protocol.Int64Value(v), metadata)
	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/serverlessrepo-2017-09-08/ListApplicationsResponse
type ListApplicationsOutput struct {
	_ struct{} `type:"structure"`

	Applications []ApplicationSummary `locationName:"applications" type:"list"`

	NextToken *string `locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s ListApplicationsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListApplicationsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if len(s.Applications) > 0 {
		v := s.Applications

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "applications", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opListApplications = "ListApplications"

// ListApplicationsRequest returns a request value for making API operation for
// AWSServerlessApplicationRepository.
//
// Lists applications owned by the requester.
//
//    // Example sending a request using ListApplicationsRequest.
//    req := client.ListApplicationsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/serverlessrepo-2017-09-08/ListApplications
func (c *Client) ListApplicationsRequest(input *ListApplicationsInput) ListApplicationsRequest {
	op := &aws.Operation{
		Name:       opListApplications,
		HTTPMethod: "GET",
		HTTPPath:   "/applications",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxItems",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListApplicationsInput{}
	}

	req := c.newRequest(op, input, &ListApplicationsOutput{})
	return ListApplicationsRequest{Request: req, Input: input, Copy: c.ListApplicationsRequest}
}

// ListApplicationsRequest is the request type for the
// ListApplications API operation.
type ListApplicationsRequest struct {
	*aws.Request
	Input *ListApplicationsInput
	Copy  func(*ListApplicationsInput) ListApplicationsRequest
}

// Send marshals and sends the ListApplications API request.
func (r ListApplicationsRequest) Send(ctx context.Context) (*ListApplicationsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListApplicationsResponse{
		ListApplicationsOutput: r.Request.Data.(*ListApplicationsOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListApplicationsRequestPaginator returns a paginator for ListApplications.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListApplicationsRequest(input)
//   p := serverlessapplicationrepository.NewListApplicationsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListApplicationsPaginator(req ListApplicationsRequest) ListApplicationsPaginator {
	return ListApplicationsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListApplicationsInput
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

// ListApplicationsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListApplicationsPaginator struct {
	aws.Pager
}

func (p *ListApplicationsPaginator) CurrentPage() *ListApplicationsOutput {
	return p.Pager.CurrentPage().(*ListApplicationsOutput)
}

// ListApplicationsResponse is the response type for the
// ListApplications API operation.
type ListApplicationsResponse struct {
	*ListApplicationsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListApplications request.
func (r *ListApplicationsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
