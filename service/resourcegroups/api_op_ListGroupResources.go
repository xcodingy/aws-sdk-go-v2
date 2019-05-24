// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package resourcegroups

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/resource-groups-2017-11-27/ListGroupResourcesInput
type ListGroupResourcesInput struct {
	_ struct{} `type:"structure"`

	// Filters, formatted as ResourceFilter objects, that you want to apply to a
	// ListGroupResources operation.
	//
	//    * resource-type - Filter resources by their type. Specify up to five resource
	//    types in the format AWS::ServiceCode::ResourceType. For example, AWS::EC2::Instance,
	//    or AWS::S3::Bucket.
	Filters []ResourceFilter `type:"list"`

	// The name of the resource group.
	//
	// GroupName is a required field
	GroupName *string `location:"uri" locationName:"GroupName" min:"1" type:"string" required:"true"`

	// The maximum number of group member ARNs that are returned in a single call
	// by ListGroupResources, in paginated output. By default, this number is 50.
	MaxResults *int64 `location:"querystring" locationName:"maxResults" min:"1" type:"integer"`

	// The NextToken value that is returned in a paginated ListGroupResources request.
	// To get the next page of results, run the call again, add the NextToken parameter,
	// and specify the NextToken value.
	NextToken *string `location:"querystring" locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s ListGroupResourcesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListGroupResourcesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListGroupResourcesInput"}

	if s.GroupName == nil {
		invalidParams.Add(aws.NewErrParamRequired("GroupName"))
	}
	if s.GroupName != nil && len(*s.GroupName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("GroupName", 1))
	}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}
	if s.Filters != nil {
		for i, v := range s.Filters {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Filters", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListGroupResourcesInput) MarshalFields(e protocol.FieldEncoder) error {

	if len(s.Filters) > 0 {
		v := s.Filters

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "Filters", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.GroupName != nil {
		v := *s.GroupName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "GroupName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
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

// Please also see https://docs.aws.amazon.com/goto/WebAPI/resource-groups-2017-11-27/ListGroupResourcesOutput
type ListGroupResourcesOutput struct {
	_ struct{} `type:"structure"`

	// The NextToken value to include in a subsequent ListGroupResources request,
	// to get more results.
	NextToken *string `type:"string"`

	// A list of QueryError objects. Each error is an object that contains ErrorCode
	// and Message structures. Possible values for ErrorCode are CLOUDFORMATION_STACK_INACTIVE
	// and CLOUDFORMATION_STACK_NOT_EXISTING.
	QueryErrors []QueryError `type:"list"`

	// The ARNs and resource types of resources that are members of the group that
	// you specified.
	ResourceIdentifiers []ResourceIdentifier `type:"list"`
}

// String returns the string representation
func (s ListGroupResourcesOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListGroupResourcesOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "NextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.QueryErrors) > 0 {
		v := s.QueryErrors

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "QueryErrors", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if len(s.ResourceIdentifiers) > 0 {
		v := s.ResourceIdentifiers

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "ResourceIdentifiers", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	return nil
}

const opListGroupResources = "ListGroupResources"

// ListGroupResourcesRequest returns a request value for making API operation for
// AWS Resource Groups.
//
// Returns a list of ARNs of resources that are members of a specified resource
// group.
//
//    // Example sending a request using ListGroupResourcesRequest.
//    req := client.ListGroupResourcesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/resource-groups-2017-11-27/ListGroupResources
func (c *Client) ListGroupResourcesRequest(input *ListGroupResourcesInput) ListGroupResourcesRequest {
	op := &aws.Operation{
		Name:       opListGroupResources,
		HTTPMethod: "POST",
		HTTPPath:   "/groups/{GroupName}/resource-identifiers-list",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListGroupResourcesInput{}
	}

	req := c.newRequest(op, input, &ListGroupResourcesOutput{})
	return ListGroupResourcesRequest{Request: req, Input: input, Copy: c.ListGroupResourcesRequest}
}

// ListGroupResourcesRequest is the request type for the
// ListGroupResources API operation.
type ListGroupResourcesRequest struct {
	*aws.Request
	Input *ListGroupResourcesInput
	Copy  func(*ListGroupResourcesInput) ListGroupResourcesRequest
}

// Send marshals and sends the ListGroupResources API request.
func (r ListGroupResourcesRequest) Send(ctx context.Context) (*ListGroupResourcesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListGroupResourcesResponse{
		ListGroupResourcesOutput: r.Request.Data.(*ListGroupResourcesOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListGroupResourcesRequestPaginator returns a paginator for ListGroupResources.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListGroupResourcesRequest(input)
//   p := resourcegroups.NewListGroupResourcesRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListGroupResourcesPaginator(req ListGroupResourcesRequest) ListGroupResourcesPaginator {
	return ListGroupResourcesPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListGroupResourcesInput
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

// ListGroupResourcesPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListGroupResourcesPaginator struct {
	aws.Pager
}

func (p *ListGroupResourcesPaginator) CurrentPage() *ListGroupResourcesOutput {
	return p.Pager.CurrentPage().(*ListGroupResourcesOutput)
}

// ListGroupResourcesResponse is the response type for the
// ListGroupResources API operation.
type ListGroupResourcesResponse struct {
	*ListGroupResourcesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListGroupResources request.
func (r *ListGroupResourcesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
