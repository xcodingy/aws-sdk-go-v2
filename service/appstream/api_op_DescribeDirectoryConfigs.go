// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package appstream

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/appstream-2016-12-01/DescribeDirectoryConfigsRequest
type DescribeDirectoryConfigsInput struct {
	_ struct{} `type:"structure"`

	// The directory names.
	DirectoryNames []string `type:"list"`

	// The maximum size of each page of results.
	MaxResults *int64 `type:"integer"`

	// The pagination token to use to retrieve the next page of results for this
	// operation. If this value is null, it retrieves the first page.
	NextToken *string `min:"1" type:"string"`
}

// String returns the string representation
func (s DescribeDirectoryConfigsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeDirectoryConfigsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeDirectoryConfigsInput"}
	if s.NextToken != nil && len(*s.NextToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NextToken", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/appstream-2016-12-01/DescribeDirectoryConfigsResult
type DescribeDirectoryConfigsOutput struct {
	_ struct{} `type:"structure"`

	// Information about the directory configurations. Note that although the response
	// syntax in this topic includes the account password, this password is not
	// returned in the actual response.
	DirectoryConfigs []DirectoryConfig `type:"list"`

	// The pagination token to use to retrieve the next page of results for this
	// operation. If there are no more pages, this value is null.
	NextToken *string `min:"1" type:"string"`
}

// String returns the string representation
func (s DescribeDirectoryConfigsOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeDirectoryConfigs = "DescribeDirectoryConfigs"

// DescribeDirectoryConfigsRequest returns a request value for making API operation for
// Amazon AppStream.
//
// Retrieves a list that describes one or more specified Directory Config objects
// for AppStream 2.0, if the names for these objects are provided. Otherwise,
// all Directory Config objects in the account are described. These objects
// include the information required to join streaming instances to an Active
// Directory domain.
//
// Although the response syntax in this topic includes the account password,
// this password is not returned in the actual response.
//
//    // Example sending a request using DescribeDirectoryConfigsRequest.
//    req := client.DescribeDirectoryConfigsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/appstream-2016-12-01/DescribeDirectoryConfigs
func (c *Client) DescribeDirectoryConfigsRequest(input *DescribeDirectoryConfigsInput) DescribeDirectoryConfigsRequest {
	op := &aws.Operation{
		Name:       opDescribeDirectoryConfigs,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeDirectoryConfigsInput{}
	}

	req := c.newRequest(op, input, &DescribeDirectoryConfigsOutput{})
	return DescribeDirectoryConfigsRequest{Request: req, Input: input, Copy: c.DescribeDirectoryConfigsRequest}
}

// DescribeDirectoryConfigsRequest is the request type for the
// DescribeDirectoryConfigs API operation.
type DescribeDirectoryConfigsRequest struct {
	*aws.Request
	Input *DescribeDirectoryConfigsInput
	Copy  func(*DescribeDirectoryConfigsInput) DescribeDirectoryConfigsRequest
}

// Send marshals and sends the DescribeDirectoryConfigs API request.
func (r DescribeDirectoryConfigsRequest) Send(ctx context.Context) (*DescribeDirectoryConfigsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeDirectoryConfigsResponse{
		DescribeDirectoryConfigsOutput: r.Request.Data.(*DescribeDirectoryConfigsOutput),
		response:                       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeDirectoryConfigsResponse is the response type for the
// DescribeDirectoryConfigs API operation.
type DescribeDirectoryConfigsResponse struct {
	*DescribeDirectoryConfigsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeDirectoryConfigs request.
func (r *DescribeDirectoryConfigsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
