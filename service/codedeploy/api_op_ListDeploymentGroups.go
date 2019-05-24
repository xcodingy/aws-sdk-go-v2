// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codedeploy

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents the input of a ListDeploymentGroups operation.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codedeploy-2014-10-06/ListDeploymentGroupsInput
type ListDeploymentGroupsInput struct {
	_ struct{} `type:"structure"`

	// The name of an AWS CodeDeploy application associated with the IAM user or
	// AWS account.
	//
	// ApplicationName is a required field
	ApplicationName *string `locationName:"applicationName" min:"1" type:"string" required:"true"`

	// An identifier returned from the previous list deployment groups call. It
	// can be used to return the next set of deployment groups in the list.
	NextToken *string `locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s ListDeploymentGroupsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListDeploymentGroupsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListDeploymentGroupsInput"}

	if s.ApplicationName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ApplicationName"))
	}
	if s.ApplicationName != nil && len(*s.ApplicationName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ApplicationName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents the output of a ListDeploymentGroups operation.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codedeploy-2014-10-06/ListDeploymentGroupsOutput
type ListDeploymentGroupsOutput struct {
	_ struct{} `type:"structure"`

	// The application name.
	ApplicationName *string `locationName:"applicationName" min:"1" type:"string"`

	// A list of deployment group names.
	DeploymentGroups []string `locationName:"deploymentGroups" type:"list"`

	// If a large amount of information is returned, an identifier is also returned.
	// It can be used in a subsequent list deployment groups call to return the
	// next set of deployment groups in the list.
	NextToken *string `locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s ListDeploymentGroupsOutput) String() string {
	return awsutil.Prettify(s)
}

const opListDeploymentGroups = "ListDeploymentGroups"

// ListDeploymentGroupsRequest returns a request value for making API operation for
// AWS CodeDeploy.
//
// Lists the deployment groups for an application registered with the IAM user
// or AWS account.
//
//    // Example sending a request using ListDeploymentGroupsRequest.
//    req := client.ListDeploymentGroupsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codedeploy-2014-10-06/ListDeploymentGroups
func (c *Client) ListDeploymentGroupsRequest(input *ListDeploymentGroupsInput) ListDeploymentGroupsRequest {
	op := &aws.Operation{
		Name:       opListDeploymentGroups,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"nextToken"},
			OutputTokens:    []string{"nextToken"},
			LimitToken:      "",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListDeploymentGroupsInput{}
	}

	req := c.newRequest(op, input, &ListDeploymentGroupsOutput{})
	return ListDeploymentGroupsRequest{Request: req, Input: input, Copy: c.ListDeploymentGroupsRequest}
}

// ListDeploymentGroupsRequest is the request type for the
// ListDeploymentGroups API operation.
type ListDeploymentGroupsRequest struct {
	*aws.Request
	Input *ListDeploymentGroupsInput
	Copy  func(*ListDeploymentGroupsInput) ListDeploymentGroupsRequest
}

// Send marshals and sends the ListDeploymentGroups API request.
func (r ListDeploymentGroupsRequest) Send(ctx context.Context) (*ListDeploymentGroupsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListDeploymentGroupsResponse{
		ListDeploymentGroupsOutput: r.Request.Data.(*ListDeploymentGroupsOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListDeploymentGroupsRequestPaginator returns a paginator for ListDeploymentGroups.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListDeploymentGroupsRequest(input)
//   p := codedeploy.NewListDeploymentGroupsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListDeploymentGroupsPaginator(req ListDeploymentGroupsRequest) ListDeploymentGroupsPaginator {
	return ListDeploymentGroupsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListDeploymentGroupsInput
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

// ListDeploymentGroupsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListDeploymentGroupsPaginator struct {
	aws.Pager
}

func (p *ListDeploymentGroupsPaginator) CurrentPage() *ListDeploymentGroupsOutput {
	return p.Pager.CurrentPage().(*ListDeploymentGroupsOutput)
}

// ListDeploymentGroupsResponse is the response type for the
// ListDeploymentGroups API operation.
type ListDeploymentGroupsResponse struct {
	*ListDeploymentGroupsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListDeploymentGroups request.
func (r *ListDeploymentGroupsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
