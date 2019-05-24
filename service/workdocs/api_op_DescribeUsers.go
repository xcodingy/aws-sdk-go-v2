// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package workdocs

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/workdocs-2016-05-01/DescribeUsersRequest
type DescribeUsersInput struct {
	_ struct{} `type:"structure"`

	// Amazon WorkDocs authentication token. Do not set this field when using administrative
	// API actions, as in accessing the API using AWS credentials.
	AuthenticationToken *string `location:"header" locationName:"Authentication" min:"1" type:"string"`

	// A comma-separated list of values. Specify "STORAGE_METADATA" to include the
	// user storage quota and utilization information.
	Fields *string `location:"querystring" locationName:"fields" min:"1" type:"string"`

	// The state of the users. Specify "ALL" to include inactive users.
	Include UserFilterType `location:"querystring" locationName:"include" type:"string" enum:"true"`

	// The maximum number of items to return.
	Limit *int64 `location:"querystring" locationName:"limit" min:"1" type:"integer"`

	// The marker for the next set of results. (You received this marker from a
	// previous call.)
	Marker *string `location:"querystring" locationName:"marker" min:"1" type:"string"`

	// The order for the results.
	Order OrderType `location:"querystring" locationName:"order" type:"string" enum:"true"`

	// The ID of the organization.
	OrganizationId *string `location:"querystring" locationName:"organizationId" min:"1" type:"string"`

	// A query to filter users by user name.
	Query *string `location:"querystring" locationName:"query" min:"1" type:"string"`

	// The sorting criteria.
	Sort UserSortType `location:"querystring" locationName:"sort" type:"string" enum:"true"`

	// The IDs of the users.
	UserIds *string `location:"querystring" locationName:"userIds" min:"1" type:"string"`
}

// String returns the string representation
func (s DescribeUsersInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeUsersInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeUsersInput"}
	if s.AuthenticationToken != nil && len(*s.AuthenticationToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("AuthenticationToken", 1))
	}
	if s.Fields != nil && len(*s.Fields) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Fields", 1))
	}
	if s.Limit != nil && *s.Limit < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("Limit", 1))
	}
	if s.Marker != nil && len(*s.Marker) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Marker", 1))
	}
	if s.OrganizationId != nil && len(*s.OrganizationId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("OrganizationId", 1))
	}
	if s.Query != nil && len(*s.Query) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Query", 1))
	}
	if s.UserIds != nil && len(*s.UserIds) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("UserIds", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeUsersInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	if s.AuthenticationToken != nil {
		v := *s.AuthenticationToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "Authentication", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Fields != nil {
		v := *s.Fields

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "fields", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.Include) > 0 {
		v := s.Include

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "include", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.Limit != nil {
		v := *s.Limit

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "limit", protocol.Int64Value(v), metadata)
	}
	if s.Marker != nil {
		v := *s.Marker

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "marker", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.Order) > 0 {
		v := s.Order

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "order", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.OrganizationId != nil {
		v := *s.OrganizationId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "organizationId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Query != nil {
		v := *s.Query

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "query", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.Sort) > 0 {
		v := s.Sort

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "sort", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.UserIds != nil {
		v := *s.UserIds

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "userIds", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/workdocs-2016-05-01/DescribeUsersResponse
type DescribeUsersOutput struct {
	_ struct{} `type:"structure"`

	// The marker to use when requesting the next set of results. If there are no
	// additional results, the string is empty.
	Marker *string `min:"1" type:"string"`

	// The total number of users included in the results.
	TotalNumberOfUsers *int64 `deprecated:"true" type:"long"`

	// The users.
	Users []User `type:"list"`
}

// String returns the string representation
func (s DescribeUsersOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeUsersOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Marker != nil {
		v := *s.Marker

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Marker", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.TotalNumberOfUsers != nil {
		v := *s.TotalNumberOfUsers

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "TotalNumberOfUsers", protocol.Int64Value(v), metadata)
	}
	if len(s.Users) > 0 {
		v := s.Users

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "Users", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	return nil
}

const opDescribeUsers = "DescribeUsers"

// DescribeUsersRequest returns a request value for making API operation for
// Amazon WorkDocs.
//
// Describes the specified users. You can describe all users or filter the results
// (for example, by status or organization).
//
// By default, Amazon WorkDocs returns the first 24 active or pending users.
// If there are more results, the response includes a marker that you can use
// to request the next set of results.
//
//    // Example sending a request using DescribeUsersRequest.
//    req := client.DescribeUsersRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/workdocs-2016-05-01/DescribeUsers
func (c *Client) DescribeUsersRequest(input *DescribeUsersInput) DescribeUsersRequest {
	op := &aws.Operation{
		Name:       opDescribeUsers,
		HTTPMethod: "GET",
		HTTPPath:   "/api/v1/users",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"Marker"},
			OutputTokens:    []string{"Marker"},
			LimitToken:      "Limit",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &DescribeUsersInput{}
	}

	req := c.newRequest(op, input, &DescribeUsersOutput{})
	return DescribeUsersRequest{Request: req, Input: input, Copy: c.DescribeUsersRequest}
}

// DescribeUsersRequest is the request type for the
// DescribeUsers API operation.
type DescribeUsersRequest struct {
	*aws.Request
	Input *DescribeUsersInput
	Copy  func(*DescribeUsersInput) DescribeUsersRequest
}

// Send marshals and sends the DescribeUsers API request.
func (r DescribeUsersRequest) Send(ctx context.Context) (*DescribeUsersResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeUsersResponse{
		DescribeUsersOutput: r.Request.Data.(*DescribeUsersOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewDescribeUsersRequestPaginator returns a paginator for DescribeUsers.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.DescribeUsersRequest(input)
//   p := workdocs.NewDescribeUsersRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewDescribeUsersPaginator(req DescribeUsersRequest) DescribeUsersPaginator {
	return DescribeUsersPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *DescribeUsersInput
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

// DescribeUsersPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type DescribeUsersPaginator struct {
	aws.Pager
}

func (p *DescribeUsersPaginator) CurrentPage() *DescribeUsersOutput {
	return p.Pager.CurrentPage().(*DescribeUsersOutput)
}

// DescribeUsersResponse is the response type for the
// DescribeUsers API operation.
type DescribeUsersResponse struct {
	*DescribeUsersOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeUsers request.
func (r *DescribeUsersResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
