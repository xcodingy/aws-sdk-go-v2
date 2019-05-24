// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package neptune

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/neptune-2014-10-31/DescribeDBParametersMessage
type DescribeDBParametersInput struct {
	_ struct{} `type:"structure"`

	// The name of a specific DB parameter group to return details for.
	//
	// Constraints:
	//
	//    * If supplied, must match the name of an existing DBParameterGroup.
	//
	// DBParameterGroupName is a required field
	DBParameterGroupName *string `type:"string" required:"true"`

	// This parameter is not currently supported.
	Filters []Filter `locationNameList:"Filter" type:"list"`

	// An optional pagination token provided by a previous DescribeDBParameters
	// request. If this parameter is specified, the response includes only records
	// beyond the marker, up to the value specified by MaxRecords.
	Marker *string `type:"string"`

	// The maximum number of records to include in the response. If more records
	// exist than the specified MaxRecords value, a pagination token called a marker
	// is included in the response so that the remaining results can be retrieved.
	//
	// Default: 100
	//
	// Constraints: Minimum 20, maximum 100.
	MaxRecords *int64 `type:"integer"`

	// The parameter types to return.
	//
	// Default: All parameter types returned
	//
	// Valid Values: user | system | engine-default
	Source *string `type:"string"`
}

// String returns the string representation
func (s DescribeDBParametersInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeDBParametersInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeDBParametersInput"}

	if s.DBParameterGroupName == nil {
		invalidParams.Add(aws.NewErrParamRequired("DBParameterGroupName"))
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

// Contains the result of a successful invocation of the DescribeDBParameters
// action.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/neptune-2014-10-31/DBParameterGroupDetails
type DescribeDBParametersOutput struct {
	_ struct{} `type:"structure"`

	// An optional pagination token provided by a previous request. If this parameter
	// is specified, the response includes only records beyond the marker, up to
	// the value specified by MaxRecords.
	Marker *string `type:"string"`

	// A list of Parameter values.
	Parameters []Parameter `locationNameList:"Parameter" type:"list"`
}

// String returns the string representation
func (s DescribeDBParametersOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeDBParameters = "DescribeDBParameters"

// DescribeDBParametersRequest returns a request value for making API operation for
// Amazon Neptune.
//
// Returns the detailed parameter list for a particular DB parameter group.
//
//    // Example sending a request using DescribeDBParametersRequest.
//    req := client.DescribeDBParametersRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/neptune-2014-10-31/DescribeDBParameters
func (c *Client) DescribeDBParametersRequest(input *DescribeDBParametersInput) DescribeDBParametersRequest {
	op := &aws.Operation{
		Name:       opDescribeDBParameters,
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
		input = &DescribeDBParametersInput{}
	}

	req := c.newRequest(op, input, &DescribeDBParametersOutput{})
	return DescribeDBParametersRequest{Request: req, Input: input, Copy: c.DescribeDBParametersRequest}
}

// DescribeDBParametersRequest is the request type for the
// DescribeDBParameters API operation.
type DescribeDBParametersRequest struct {
	*aws.Request
	Input *DescribeDBParametersInput
	Copy  func(*DescribeDBParametersInput) DescribeDBParametersRequest
}

// Send marshals and sends the DescribeDBParameters API request.
func (r DescribeDBParametersRequest) Send(ctx context.Context) (*DescribeDBParametersResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeDBParametersResponse{
		DescribeDBParametersOutput: r.Request.Data.(*DescribeDBParametersOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewDescribeDBParametersRequestPaginator returns a paginator for DescribeDBParameters.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.DescribeDBParametersRequest(input)
//   p := neptune.NewDescribeDBParametersRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewDescribeDBParametersPaginator(req DescribeDBParametersRequest) DescribeDBParametersPaginator {
	return DescribeDBParametersPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *DescribeDBParametersInput
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

// DescribeDBParametersPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type DescribeDBParametersPaginator struct {
	aws.Pager
}

func (p *DescribeDBParametersPaginator) CurrentPage() *DescribeDBParametersOutput {
	return p.Pager.CurrentPage().(*DescribeDBParametersOutput)
}

// DescribeDBParametersResponse is the response type for the
// DescribeDBParameters API operation.
type DescribeDBParametersResponse struct {
	*DescribeDBParametersOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeDBParameters request.
func (r *DescribeDBParametersResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
