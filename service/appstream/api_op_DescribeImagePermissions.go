// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package appstream

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/appstream-2016-12-01/DescribeImagePermissionsRequest
type DescribeImagePermissionsInput struct {
	_ struct{} `type:"structure"`

	// The maximum size of each page of results.
	MaxResults *int64 `type:"integer"`

	// The name of the private image for which to describe permissions. The image
	// must be one that you own.
	//
	// Name is a required field
	Name *string `type:"string" required:"true"`

	// The pagination token to use to retrieve the next page of results for this
	// operation. If this value is null, it retrieves the first page.
	NextToken *string `min:"1" type:"string"`

	// The 12-digit identifier of one or more AWS accounts with which the image
	// is shared.
	SharedAwsAccountIds []string `min:"1" type:"list"`
}

// String returns the string representation
func (s DescribeImagePermissionsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeImagePermissionsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeImagePermissionsInput"}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}
	if s.NextToken != nil && len(*s.NextToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NextToken", 1))
	}
	if s.SharedAwsAccountIds != nil && len(s.SharedAwsAccountIds) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("SharedAwsAccountIds", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/appstream-2016-12-01/DescribeImagePermissionsResult
type DescribeImagePermissionsOutput struct {
	_ struct{} `type:"structure"`

	// The name of the private image.
	Name *string `type:"string"`

	// The pagination token to use to retrieve the next page of results for this
	// operation. If there are no more pages, this value is null.
	NextToken *string `min:"1" type:"string"`

	// The permissions for a private image that you own.
	SharedImagePermissionsList []SharedImagePermissions `type:"list"`
}

// String returns the string representation
func (s DescribeImagePermissionsOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeImagePermissions = "DescribeImagePermissions"

// DescribeImagePermissionsRequest returns a request value for making API operation for
// Amazon AppStream.
//
// Retrieves a list that describes the permissions for shared AWS account IDs
// on a private image that you own.
//
//    // Example sending a request using DescribeImagePermissionsRequest.
//    req := client.DescribeImagePermissionsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/appstream-2016-12-01/DescribeImagePermissions
func (c *Client) DescribeImagePermissionsRequest(input *DescribeImagePermissionsInput) DescribeImagePermissionsRequest {
	op := &aws.Operation{
		Name:       opDescribeImagePermissions,
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
		input = &DescribeImagePermissionsInput{}
	}

	req := c.newRequest(op, input, &DescribeImagePermissionsOutput{})
	return DescribeImagePermissionsRequest{Request: req, Input: input, Copy: c.DescribeImagePermissionsRequest}
}

// DescribeImagePermissionsRequest is the request type for the
// DescribeImagePermissions API operation.
type DescribeImagePermissionsRequest struct {
	*aws.Request
	Input *DescribeImagePermissionsInput
	Copy  func(*DescribeImagePermissionsInput) DescribeImagePermissionsRequest
}

// Send marshals and sends the DescribeImagePermissions API request.
func (r DescribeImagePermissionsRequest) Send(ctx context.Context) (*DescribeImagePermissionsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeImagePermissionsResponse{
		DescribeImagePermissionsOutput: r.Request.Data.(*DescribeImagePermissionsOutput),
		response:                       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewDescribeImagePermissionsRequestPaginator returns a paginator for DescribeImagePermissions.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.DescribeImagePermissionsRequest(input)
//   p := appstream.NewDescribeImagePermissionsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewDescribeImagePermissionsPaginator(req DescribeImagePermissionsRequest) DescribeImagePermissionsPaginator {
	return DescribeImagePermissionsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *DescribeImagePermissionsInput
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

// DescribeImagePermissionsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type DescribeImagePermissionsPaginator struct {
	aws.Pager
}

func (p *DescribeImagePermissionsPaginator) CurrentPage() *DescribeImagePermissionsOutput {
	return p.Pager.CurrentPage().(*DescribeImagePermissionsOutput)
}

// DescribeImagePermissionsResponse is the response type for the
// DescribeImagePermissions API operation.
type DescribeImagePermissionsResponse struct {
	*DescribeImagePermissionsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeImagePermissions request.
func (r *DescribeImagePermissionsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
