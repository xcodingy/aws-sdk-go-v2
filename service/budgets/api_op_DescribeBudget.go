// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package budgets

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Request of DescribeBudget
type DescribeBudgetInput struct {
	_ struct{} `type:"structure"`

	// The accountId that is associated with the budget that you want a description
	// of.
	//
	// AccountId is a required field
	AccountId *string `min:"12" type:"string" required:"true"`

	// The name of the budget that you want a description of.
	//
	// BudgetName is a required field
	BudgetName *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeBudgetInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeBudgetInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeBudgetInput"}

	if s.AccountId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AccountId"))
	}
	if s.AccountId != nil && len(*s.AccountId) < 12 {
		invalidParams.Add(aws.NewErrParamMinLen("AccountId", 12))
	}

	if s.BudgetName == nil {
		invalidParams.Add(aws.NewErrParamRequired("BudgetName"))
	}
	if s.BudgetName != nil && len(*s.BudgetName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("BudgetName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Response of DescribeBudget
type DescribeBudgetOutput struct {
	_ struct{} `type:"structure"`

	// The description of the budget.
	Budget *Budget `type:"structure"`
}

// String returns the string representation
func (s DescribeBudgetOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeBudget = "DescribeBudget"

// DescribeBudgetRequest returns a request value for making API operation for
// AWS Budgets.
//
// Describes a budget.
//
//    // Example sending a request using DescribeBudgetRequest.
//    req := client.DescribeBudgetRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) DescribeBudgetRequest(input *DescribeBudgetInput) DescribeBudgetRequest {
	op := &aws.Operation{
		Name:       opDescribeBudget,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeBudgetInput{}
	}

	req := c.newRequest(op, input, &DescribeBudgetOutput{})
	return DescribeBudgetRequest{Request: req, Input: input, Copy: c.DescribeBudgetRequest}
}

// DescribeBudgetRequest is the request type for the
// DescribeBudget API operation.
type DescribeBudgetRequest struct {
	*aws.Request
	Input *DescribeBudgetInput
	Copy  func(*DescribeBudgetInput) DescribeBudgetRequest
}

// Send marshals and sends the DescribeBudget API request.
func (r DescribeBudgetRequest) Send(ctx context.Context) (*DescribeBudgetResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeBudgetResponse{
		DescribeBudgetOutput: r.Request.Data.(*DescribeBudgetOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeBudgetResponse is the response type for the
// DescribeBudget API operation.
type DescribeBudgetResponse struct {
	*DescribeBudgetOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeBudget request.
func (r *DescribeBudgetResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}