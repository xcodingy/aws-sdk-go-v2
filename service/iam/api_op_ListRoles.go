// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iam

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/iam-2010-05-08/ListRolesRequest
type ListRolesInput struct {
	_ struct{} `type:"structure"`

	// Use this parameter only when paginating results and only after you receive
	// a response indicating that the results are truncated. Set it to the value
	// of the Marker element in the response that you received to indicate where
	// the next call should start.
	Marker *string `min:"1" type:"string"`

	// Use this only when paginating results to indicate the maximum number of items
	// you want in the response. If additional items exist beyond the maximum you
	// specify, the IsTruncated response element is true.
	//
	// If you do not include this parameter, the number of items defaults to 100.
	// Note that IAM might return fewer results, even when there are more results
	// available. In that case, the IsTruncated response element returns true, and
	// Marker contains a value to include in the subsequent call that tells the
	// service where to continue from.
	MaxItems *int64 `min:"1" type:"integer"`

	// The path prefix for filtering the results. For example, the prefix /application_abc/component_xyz/
	// gets all roles whose path starts with /application_abc/component_xyz/.
	//
	// This parameter is optional. If it is not included, it defaults to a slash
	// (/), listing all roles. This parameter allows (through its regex pattern
	// (http://wikipedia.org/wiki/regex)) a string of characters consisting of either
	// a forward slash (/) by itself or a string that must begin and end with forward
	// slashes. In addition, it can contain any ASCII character from the ! (\u0021)
	// through the DEL character (\u007F), including most punctuation characters,
	// digits, and upper and lowercased letters.
	PathPrefix *string `min:"1" type:"string"`
}

// String returns the string representation
func (s ListRolesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListRolesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListRolesInput"}
	if s.Marker != nil && len(*s.Marker) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Marker", 1))
	}
	if s.MaxItems != nil && *s.MaxItems < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxItems", 1))
	}
	if s.PathPrefix != nil && len(*s.PathPrefix) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("PathPrefix", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Contains the response to a successful ListRoles request.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iam-2010-05-08/ListRolesResponse
type ListRolesOutput struct {
	_ struct{} `type:"structure"`

	// A flag that indicates whether there are more items to return. If your results
	// were truncated, you can make a subsequent pagination request using the Marker
	// request parameter to retrieve more items. Note that IAM might return fewer
	// than the MaxItems number of results even when there are more results available.
	// We recommend that you check IsTruncated after every call to ensure that you
	// receive all your results.
	IsTruncated *bool `type:"boolean"`

	// When IsTruncated is true, this element is present and contains the value
	// to use for the Marker parameter in a subsequent pagination request.
	Marker *string `min:"1" type:"string"`

	// A list of roles.
	//
	// Roles is a required field
	Roles []Role `type:"list" required:"true"`
}

// String returns the string representation
func (s ListRolesOutput) String() string {
	return awsutil.Prettify(s)
}

const opListRoles = "ListRoles"

// ListRolesRequest returns a request value for making API operation for
// AWS Identity and Access Management.
//
// Lists the IAM roles that have the specified path prefix. If there are none,
// the operation returns an empty list. For more information about roles, go
// to Working with Roles (https://docs.aws.amazon.com/IAM/latest/UserGuide/WorkingWithRoles.html).
//
// You can paginate the results using the MaxItems and Marker parameters.
//
//    // Example sending a request using ListRolesRequest.
//    req := client.ListRolesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iam-2010-05-08/ListRoles
func (c *Client) ListRolesRequest(input *ListRolesInput) ListRolesRequest {
	op := &aws.Operation{
		Name:       opListRoles,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"Marker"},
			OutputTokens:    []string{"Marker"},
			LimitToken:      "MaxItems",
			TruncationToken: "IsTruncated",
		},
	}

	if input == nil {
		input = &ListRolesInput{}
	}

	req := c.newRequest(op, input, &ListRolesOutput{})
	return ListRolesRequest{Request: req, Input: input, Copy: c.ListRolesRequest}
}

// ListRolesRequest is the request type for the
// ListRoles API operation.
type ListRolesRequest struct {
	*aws.Request
	Input *ListRolesInput
	Copy  func(*ListRolesInput) ListRolesRequest
}

// Send marshals and sends the ListRoles API request.
func (r ListRolesRequest) Send(ctx context.Context) (*ListRolesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListRolesResponse{
		ListRolesOutput: r.Request.Data.(*ListRolesOutput),
		response:        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListRolesRequestPaginator returns a paginator for ListRoles.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListRolesRequest(input)
//   p := iam.NewListRolesRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListRolesPaginator(req ListRolesRequest) ListRolesPaginator {
	return ListRolesPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListRolesInput
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

// ListRolesPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListRolesPaginator struct {
	aws.Pager
}

func (p *ListRolesPaginator) CurrentPage() *ListRolesOutput {
	return p.Pager.CurrentPage().(*ListRolesOutput)
}

// ListRolesResponse is the response type for the
// ListRoles API operation.
type ListRolesResponse struct {
	*ListRolesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListRoles request.
func (r *ListRolesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
