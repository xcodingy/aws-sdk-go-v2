// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package inspector

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/inspector-2016-02-16/DescribeRulesPackagesRequest
type DescribeRulesPackagesInput struct {
	_ struct{} `type:"structure"`

	// The locale that you want to translate a rules package description into.
	Locale Locale `locationName:"locale" type:"string" enum:"true"`

	// The ARN that specifies the rules package that you want to describe.
	//
	// RulesPackageArns is a required field
	RulesPackageArns []string `locationName:"rulesPackageArns" min:"1" type:"list" required:"true"`
}

// String returns the string representation
func (s DescribeRulesPackagesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeRulesPackagesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeRulesPackagesInput"}

	if s.RulesPackageArns == nil {
		invalidParams.Add(aws.NewErrParamRequired("RulesPackageArns"))
	}
	if s.RulesPackageArns != nil && len(s.RulesPackageArns) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("RulesPackageArns", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/inspector-2016-02-16/DescribeRulesPackagesResponse
type DescribeRulesPackagesOutput struct {
	_ struct{} `type:"structure"`

	// Rules package details that cannot be described. An error code is provided
	// for each failed item.
	//
	// FailedItems is a required field
	FailedItems map[string]FailedItemDetails `locationName:"failedItems" type:"map" required:"true"`

	// Information about the rules package.
	//
	// RulesPackages is a required field
	RulesPackages []RulesPackage `locationName:"rulesPackages" type:"list" required:"true"`
}

// String returns the string representation
func (s DescribeRulesPackagesOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeRulesPackages = "DescribeRulesPackages"

// DescribeRulesPackagesRequest returns a request value for making API operation for
// Amazon Inspector.
//
// Describes the rules packages that are specified by the ARNs of the rules
// packages.
//
//    // Example sending a request using DescribeRulesPackagesRequest.
//    req := client.DescribeRulesPackagesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/inspector-2016-02-16/DescribeRulesPackages
func (c *Client) DescribeRulesPackagesRequest(input *DescribeRulesPackagesInput) DescribeRulesPackagesRequest {
	op := &aws.Operation{
		Name:       opDescribeRulesPackages,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeRulesPackagesInput{}
	}

	req := c.newRequest(op, input, &DescribeRulesPackagesOutput{})
	return DescribeRulesPackagesRequest{Request: req, Input: input, Copy: c.DescribeRulesPackagesRequest}
}

// DescribeRulesPackagesRequest is the request type for the
// DescribeRulesPackages API operation.
type DescribeRulesPackagesRequest struct {
	*aws.Request
	Input *DescribeRulesPackagesInput
	Copy  func(*DescribeRulesPackagesInput) DescribeRulesPackagesRequest
}

// Send marshals and sends the DescribeRulesPackages API request.
func (r DescribeRulesPackagesRequest) Send(ctx context.Context) (*DescribeRulesPackagesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeRulesPackagesResponse{
		DescribeRulesPackagesOutput: r.Request.Data.(*DescribeRulesPackagesOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeRulesPackagesResponse is the response type for the
// DescribeRulesPackages API operation.
type DescribeRulesPackagesResponse struct {
	*DescribeRulesPackagesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeRulesPackages request.
func (r *DescribeRulesPackagesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
