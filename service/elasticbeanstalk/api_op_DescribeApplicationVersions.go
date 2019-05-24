// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package elasticbeanstalk

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Request to describe application versions.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticbeanstalk-2010-12-01/DescribeApplicationVersionsMessage
type DescribeApplicationVersionsInput struct {
	_ struct{} `type:"structure"`

	// Specify an application name to show only application versions for that application.
	ApplicationName *string `min:"1" type:"string"`

	// For a paginated request. Specify a maximum number of application versions
	// to include in each response.
	//
	// If no MaxRecords is specified, all available application versions are retrieved
	// in a single response.
	MaxRecords *int64 `min:"1" type:"integer"`

	// For a paginated request. Specify a token from a previous response page to
	// retrieve the next response page. All other parameter values must be identical
	// to the ones specified in the initial request.
	//
	// If no NextToken is specified, the first page is retrieved.
	NextToken *string `type:"string"`

	// Specify a version label to show a specific application version.
	VersionLabels []string `type:"list"`
}

// String returns the string representation
func (s DescribeApplicationVersionsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeApplicationVersionsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeApplicationVersionsInput"}
	if s.ApplicationName != nil && len(*s.ApplicationName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ApplicationName", 1))
	}
	if s.MaxRecords != nil && *s.MaxRecords < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxRecords", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Result message wrapping a list of application version descriptions.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticbeanstalk-2010-12-01/ApplicationVersionDescriptionsMessage
type DescribeApplicationVersionsOutput struct {
	_ struct{} `type:"structure"`

	// List of ApplicationVersionDescription objects sorted in order of creation.
	ApplicationVersions []ApplicationVersionDescription `type:"list"`

	// In a paginated request, the token that you can pass in a subsequent request
	// to get the next response page.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s DescribeApplicationVersionsOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeApplicationVersions = "DescribeApplicationVersions"

// DescribeApplicationVersionsRequest returns a request value for making API operation for
// AWS Elastic Beanstalk.
//
// Retrieve a list of application versions.
//
//    // Example sending a request using DescribeApplicationVersionsRequest.
//    req := client.DescribeApplicationVersionsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticbeanstalk-2010-12-01/DescribeApplicationVersions
func (c *Client) DescribeApplicationVersionsRequest(input *DescribeApplicationVersionsInput) DescribeApplicationVersionsRequest {
	op := &aws.Operation{
		Name:       opDescribeApplicationVersions,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeApplicationVersionsInput{}
	}

	req := c.newRequest(op, input, &DescribeApplicationVersionsOutput{})
	return DescribeApplicationVersionsRequest{Request: req, Input: input, Copy: c.DescribeApplicationVersionsRequest}
}

// DescribeApplicationVersionsRequest is the request type for the
// DescribeApplicationVersions API operation.
type DescribeApplicationVersionsRequest struct {
	*aws.Request
	Input *DescribeApplicationVersionsInput
	Copy  func(*DescribeApplicationVersionsInput) DescribeApplicationVersionsRequest
}

// Send marshals and sends the DescribeApplicationVersions API request.
func (r DescribeApplicationVersionsRequest) Send(ctx context.Context) (*DescribeApplicationVersionsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeApplicationVersionsResponse{
		DescribeApplicationVersionsOutput: r.Request.Data.(*DescribeApplicationVersionsOutput),
		response:                          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeApplicationVersionsResponse is the response type for the
// DescribeApplicationVersions API operation.
type DescribeApplicationVersionsResponse struct {
	*DescribeApplicationVersionsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeApplicationVersions request.
func (r *DescribeApplicationVersionsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
