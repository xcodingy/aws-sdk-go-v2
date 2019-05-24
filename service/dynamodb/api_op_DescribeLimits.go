// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package dynamodb

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents the input of a DescribeLimits operation. Has no content.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/dynamodb-2012-08-10/DescribeLimitsInput
type DescribeLimitsInput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DescribeLimitsInput) String() string {
	return awsutil.Prettify(s)
}

// Represents the output of a DescribeLimits operation.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/dynamodb-2012-08-10/DescribeLimitsOutput
type DescribeLimitsOutput struct {
	_ struct{} `type:"structure"`

	// The maximum total read capacity units that your account allows you to provision
	// across all of your tables in this region.
	AccountMaxReadCapacityUnits *int64 `min:"1" type:"long"`

	// The maximum total write capacity units that your account allows you to provision
	// across all of your tables in this region.
	AccountMaxWriteCapacityUnits *int64 `min:"1" type:"long"`

	// The maximum read capacity units that your account allows you to provision
	// for a new table that you are creating in this region, including the read
	// capacity units provisioned for its global secondary indexes (GSIs).
	TableMaxReadCapacityUnits *int64 `min:"1" type:"long"`

	// The maximum write capacity units that your account allows you to provision
	// for a new table that you are creating in this region, including the write
	// capacity units provisioned for its global secondary indexes (GSIs).
	TableMaxWriteCapacityUnits *int64 `min:"1" type:"long"`
}

// String returns the string representation
func (s DescribeLimitsOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeLimits = "DescribeLimits"

// DescribeLimitsRequest returns a request value for making API operation for
// Amazon DynamoDB.
//
// Returns the current provisioned-capacity limits for your AWS account in a
// region, both for the region as a whole and for any one DynamoDB table that
// you create there.
//
// When you establish an AWS account, the account has initial limits on the
// maximum read capacity units and write capacity units that you can provision
// across all of your DynamoDB tables in a given region. Also, there are per-table
// limits that apply when you create a table there. For more information, see
// Limits (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Limits.html)
// page in the Amazon DynamoDB Developer Guide.
//
// Although you can increase these limits by filing a case at AWS Support Center
// (https://console.aws.amazon.com/support/home#/), obtaining the increase is
// not instantaneous. The DescribeLimits action lets you write code to compare
// the capacity you are currently using to those limits imposed by your account
// so that you have enough time to apply for an increase before you hit a limit.
//
// For example, you could use one of the AWS SDKs to do the following:
//
// Call DescribeLimits for a particular region to obtain your current account
// limits on provisioned capacity there.
//
// Create a variable to hold the aggregate read capacity units provisioned for
// all your tables in that region, and one to hold the aggregate write capacity
// units. Zero them both.
//
// Call ListTables to obtain a list of all your DynamoDB tables.
//
// For each table name listed by ListTables, do the following:
//
// Call DescribeTable with the table name.
//
// Use the data returned by DescribeTable to add the read capacity units and
// write capacity units provisioned for the table itself to your variables.
//
// If the table has one or more global secondary indexes (GSIs), loop over these
// GSIs and add their provisioned capacity values to your variables as well.
//
// Report the account limits for that region returned by DescribeLimits, along
// with the total current provisioned capacity levels you have calculated.
//
// This will let you see whether you are getting close to your account-level
// limits.
//
// The per-table limits apply only when you are creating a new table. They restrict
// the sum of the provisioned capacity of the new table itself and all its global
// secondary indexes.
//
// For existing tables and their GSIs, DynamoDB will not let you increase provisioned
// capacity extremely rapidly, but the only upper limit that applies is that
// the aggregate provisioned capacity over all your tables and GSIs cannot exceed
// either of the per-account limits.
//
// DescribeLimits should only be called periodically. You can expect throttling
// errors if you call it more than once in a minute.
//
// The DescribeLimits Request element has no content.
//
//    // Example sending a request using DescribeLimitsRequest.
//    req := client.DescribeLimitsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/dynamodb-2012-08-10/DescribeLimits
func (c *Client) DescribeLimitsRequest(input *DescribeLimitsInput) DescribeLimitsRequest {
	op := &aws.Operation{
		Name:       opDescribeLimits,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeLimitsInput{}
	}

	req := c.newRequest(op, input, &DescribeLimitsOutput{})
	return DescribeLimitsRequest{Request: req, Input: input, Copy: c.DescribeLimitsRequest}
}

// DescribeLimitsRequest is the request type for the
// DescribeLimits API operation.
type DescribeLimitsRequest struct {
	*aws.Request
	Input *DescribeLimitsInput
	Copy  func(*DescribeLimitsInput) DescribeLimitsRequest
}

// Send marshals and sends the DescribeLimits API request.
func (r DescribeLimitsRequest) Send(ctx context.Context) (*DescribeLimitsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeLimitsResponse{
		DescribeLimitsOutput: r.Request.Data.(*DescribeLimitsOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeLimitsResponse is the response type for the
// DescribeLimits API operation.
type DescribeLimitsResponse struct {
	*DescribeLimitsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeLimits request.
func (r *DescribeLimitsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
