// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package configservice

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/config-2014-11-12/DescribeConfigurationAggregatorsRequest
type DescribeConfigurationAggregatorsInput struct {
	_ struct{} `type:"structure"`

	// The name of the configuration aggregators.
	ConfigurationAggregatorNames []string `type:"list"`

	// The maximum number of configuration aggregators returned on each page. The
	// default is maximum. If you specify 0, AWS Config uses the default.
	Limit *int64 `type:"integer"`

	// The nextToken string returned on a previous page that you use to get the
	// next page of results in a paginated response.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s DescribeConfigurationAggregatorsInput) String() string {
	return awsutil.Prettify(s)
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/config-2014-11-12/DescribeConfigurationAggregatorsResponse
type DescribeConfigurationAggregatorsOutput struct {
	_ struct{} `type:"structure"`

	// Returns a ConfigurationAggregators object.
	ConfigurationAggregators []ConfigurationAggregator `type:"list"`

	// The nextToken string returned on a previous page that you use to get the
	// next page of results in a paginated response.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s DescribeConfigurationAggregatorsOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeConfigurationAggregators = "DescribeConfigurationAggregators"

// DescribeConfigurationAggregatorsRequest returns a request value for making API operation for
// AWS Config.
//
// Returns the details of one or more configuration aggregators. If the configuration
// aggregator is not specified, this action returns the details for all the
// configuration aggregators associated with the account.
//
//    // Example sending a request using DescribeConfigurationAggregatorsRequest.
//    req := client.DescribeConfigurationAggregatorsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/config-2014-11-12/DescribeConfigurationAggregators
func (c *Client) DescribeConfigurationAggregatorsRequest(input *DescribeConfigurationAggregatorsInput) DescribeConfigurationAggregatorsRequest {
	op := &aws.Operation{
		Name:       opDescribeConfigurationAggregators,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeConfigurationAggregatorsInput{}
	}

	req := c.newRequest(op, input, &DescribeConfigurationAggregatorsOutput{})
	return DescribeConfigurationAggregatorsRequest{Request: req, Input: input, Copy: c.DescribeConfigurationAggregatorsRequest}
}

// DescribeConfigurationAggregatorsRequest is the request type for the
// DescribeConfigurationAggregators API operation.
type DescribeConfigurationAggregatorsRequest struct {
	*aws.Request
	Input *DescribeConfigurationAggregatorsInput
	Copy  func(*DescribeConfigurationAggregatorsInput) DescribeConfigurationAggregatorsRequest
}

// Send marshals and sends the DescribeConfigurationAggregators API request.
func (r DescribeConfigurationAggregatorsRequest) Send(ctx context.Context) (*DescribeConfigurationAggregatorsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeConfigurationAggregatorsResponse{
		DescribeConfigurationAggregatorsOutput: r.Request.Data.(*DescribeConfigurationAggregatorsOutput),
		response:                               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeConfigurationAggregatorsResponse is the response type for the
// DescribeConfigurationAggregators API operation.
type DescribeConfigurationAggregatorsResponse struct {
	*DescribeConfigurationAggregatorsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeConfigurationAggregators request.
func (r *DescribeConfigurationAggregatorsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
