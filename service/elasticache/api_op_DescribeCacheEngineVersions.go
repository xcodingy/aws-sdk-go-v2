// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package elasticache

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents the input of a DescribeCacheEngineVersions operation.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticache-2015-02-02/DescribeCacheEngineVersionsMessage
type DescribeCacheEngineVersionsInput struct {
	_ struct{} `type:"structure"`

	// The name of a specific cache parameter group family to return details for.
	//
	// Valid values are: memcached1.4 | redis2.6 | redis2.8 | redis3.2 | redis4.0
	//
	// Constraints:
	//
	//    * Must be 1 to 255 alphanumeric characters
	//
	//    * First character must be a letter
	//
	//    * Cannot end with a hyphen or contain two consecutive hyphens
	CacheParameterGroupFamily *string `type:"string"`

	// If true, specifies that only the default version of the specified engine
	// or engine and major version combination is to be returned.
	DefaultOnly *bool `type:"boolean"`

	// The cache engine to return. Valid values: memcached | redis
	Engine *string `type:"string"`

	// The cache engine version to return.
	//
	// Example: 1.4.14
	EngineVersion *string `type:"string"`

	// An optional marker returned from a prior request. Use this marker for pagination
	// of results from this operation. If this parameter is specified, the response
	// includes only records beyond the marker, up to the value specified by MaxRecords.
	Marker *string `type:"string"`

	// The maximum number of records to include in the response. If more records
	// exist than the specified MaxRecords value, a marker is included in the response
	// so that the remaining results can be retrieved.
	//
	// Default: 100
	//
	// Constraints: minimum 20; maximum 100.
	MaxRecords *int64 `type:"integer"`
}

// String returns the string representation
func (s DescribeCacheEngineVersionsInput) String() string {
	return awsutil.Prettify(s)
}

// Represents the output of a DescribeCacheEngineVersions operation.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticache-2015-02-02/CacheEngineVersionMessage
type DescribeCacheEngineVersionsOutput struct {
	_ struct{} `type:"structure"`

	// A list of cache engine version details. Each element in the list contains
	// detailed information about one cache engine version.
	CacheEngineVersions []CacheEngineVersion `locationNameList:"CacheEngineVersion" type:"list"`

	// Provides an identifier to allow retrieval of paginated results.
	Marker *string `type:"string"`
}

// String returns the string representation
func (s DescribeCacheEngineVersionsOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeCacheEngineVersions = "DescribeCacheEngineVersions"

// DescribeCacheEngineVersionsRequest returns a request value for making API operation for
// Amazon ElastiCache.
//
// Returns a list of the available cache engines and their versions.
//
//    // Example sending a request using DescribeCacheEngineVersionsRequest.
//    req := client.DescribeCacheEngineVersionsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticache-2015-02-02/DescribeCacheEngineVersions
func (c *Client) DescribeCacheEngineVersionsRequest(input *DescribeCacheEngineVersionsInput) DescribeCacheEngineVersionsRequest {
	op := &aws.Operation{
		Name:       opDescribeCacheEngineVersions,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"Marker"},
			OutputTokens:    []string{"Marker"},
			LimitToken:      "MaxRecords",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &DescribeCacheEngineVersionsInput{}
	}

	req := c.newRequest(op, input, &DescribeCacheEngineVersionsOutput{})
	return DescribeCacheEngineVersionsRequest{Request: req, Input: input, Copy: c.DescribeCacheEngineVersionsRequest}
}

// DescribeCacheEngineVersionsRequest is the request type for the
// DescribeCacheEngineVersions API operation.
type DescribeCacheEngineVersionsRequest struct {
	*aws.Request
	Input *DescribeCacheEngineVersionsInput
	Copy  func(*DescribeCacheEngineVersionsInput) DescribeCacheEngineVersionsRequest
}

// Send marshals and sends the DescribeCacheEngineVersions API request.
func (r DescribeCacheEngineVersionsRequest) Send(ctx context.Context) (*DescribeCacheEngineVersionsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeCacheEngineVersionsResponse{
		DescribeCacheEngineVersionsOutput: r.Request.Data.(*DescribeCacheEngineVersionsOutput),
		response:                          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewDescribeCacheEngineVersionsRequestPaginator returns a paginator for DescribeCacheEngineVersions.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.DescribeCacheEngineVersionsRequest(input)
//   p := elasticache.NewDescribeCacheEngineVersionsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewDescribeCacheEngineVersionsPaginator(req DescribeCacheEngineVersionsRequest) DescribeCacheEngineVersionsPaginator {
	return DescribeCacheEngineVersionsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *DescribeCacheEngineVersionsInput
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

// DescribeCacheEngineVersionsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type DescribeCacheEngineVersionsPaginator struct {
	aws.Pager
}

func (p *DescribeCacheEngineVersionsPaginator) CurrentPage() *DescribeCacheEngineVersionsOutput {
	return p.Pager.CurrentPage().(*DescribeCacheEngineVersionsOutput)
}

// DescribeCacheEngineVersionsResponse is the response type for the
// DescribeCacheEngineVersions API operation.
type DescribeCacheEngineVersionsResponse struct {
	*DescribeCacheEngineVersionsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeCacheEngineVersions request.
func (r *DescribeCacheEngineVersionsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
