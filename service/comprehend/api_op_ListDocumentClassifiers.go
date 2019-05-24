// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package comprehend

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/comprehend-2017-11-27/ListDocumentClassifiersRequest
type ListDocumentClassifiersInput struct {
	_ struct{} `type:"structure"`

	// Filters the jobs that are returned. You can filter jobs on their name, status,
	// or the date and time that they were submitted. You can only set one filter
	// at a time.
	Filter *DocumentClassifierFilter `type:"structure"`

	// The maximum number of results to return in each page. The default is 100.
	MaxResults *int64 `min:"1" type:"integer"`

	// Identifies the next page of results to return.
	NextToken *string `min:"1" type:"string"`
}

// String returns the string representation
func (s ListDocumentClassifiersInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListDocumentClassifiersInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListDocumentClassifiersInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}
	if s.NextToken != nil && len(*s.NextToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NextToken", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/comprehend-2017-11-27/ListDocumentClassifiersResponse
type ListDocumentClassifiersOutput struct {
	_ struct{} `type:"structure"`

	// A list containing the properties of each job returned.
	DocumentClassifierPropertiesList []DocumentClassifierProperties `type:"list"`

	// Identifies the next page of results to return.
	NextToken *string `min:"1" type:"string"`
}

// String returns the string representation
func (s ListDocumentClassifiersOutput) String() string {
	return awsutil.Prettify(s)
}

const opListDocumentClassifiers = "ListDocumentClassifiers"

// ListDocumentClassifiersRequest returns a request value for making API operation for
// Amazon Comprehend.
//
// Gets a list of the document classifiers that you have created.
//
//    // Example sending a request using ListDocumentClassifiersRequest.
//    req := client.ListDocumentClassifiersRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/comprehend-2017-11-27/ListDocumentClassifiers
func (c *Client) ListDocumentClassifiersRequest(input *ListDocumentClassifiersInput) ListDocumentClassifiersRequest {
	op := &aws.Operation{
		Name:       opListDocumentClassifiers,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListDocumentClassifiersInput{}
	}

	req := c.newRequest(op, input, &ListDocumentClassifiersOutput{})
	return ListDocumentClassifiersRequest{Request: req, Input: input, Copy: c.ListDocumentClassifiersRequest}
}

// ListDocumentClassifiersRequest is the request type for the
// ListDocumentClassifiers API operation.
type ListDocumentClassifiersRequest struct {
	*aws.Request
	Input *ListDocumentClassifiersInput
	Copy  func(*ListDocumentClassifiersInput) ListDocumentClassifiersRequest
}

// Send marshals and sends the ListDocumentClassifiers API request.
func (r ListDocumentClassifiersRequest) Send(ctx context.Context) (*ListDocumentClassifiersResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListDocumentClassifiersResponse{
		ListDocumentClassifiersOutput: r.Request.Data.(*ListDocumentClassifiersOutput),
		response:                      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListDocumentClassifiersResponse is the response type for the
// ListDocumentClassifiers API operation.
type ListDocumentClassifiersResponse struct {
	*ListDocumentClassifiersOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListDocumentClassifiers request.
func (r *ListDocumentClassifiersResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
