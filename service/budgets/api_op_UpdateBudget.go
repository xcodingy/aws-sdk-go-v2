// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package budgets

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Request of UpdateBudget
type UpdateBudgetInput struct {
	_ struct{} `type:"structure"`

	// The accountId that is associated with the budget that you want to update.
	//
	// AccountId is a required field
	AccountId *string `min:"12" type:"string" required:"true"`

	// The budget that you want to update your budget to.
	//
	// NewBudget is a required field
	NewBudget *Budget `type:"structure" required:"true"`
}

// String returns the string representation
func (s UpdateBudgetInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateBudgetInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateBudgetInput"}

	if s.AccountId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AccountId"))
	}
	if s.AccountId != nil && len(*s.AccountId) < 12 {
		invalidParams.Add(aws.NewErrParamMinLen("AccountId", 12))
	}

	if s.NewBudget == nil {
		invalidParams.Add(aws.NewErrParamRequired("NewBudget"))
	}
	if s.NewBudget != nil {
		if err := s.NewBudget.Validate(); err != nil {
			invalidParams.AddNested("NewBudget", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Response of UpdateBudget
type UpdateBudgetOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s UpdateBudgetOutput) String() string {
	return awsutil.Prettify(s)
}

const opUpdateBudget = "UpdateBudget"

// UpdateBudgetRequest returns a request value for making API operation for
// AWS Budgets.
//
// Updates a budget. You can change every part of a budget except for the budgetName
// and the calculatedSpend. When you modify a budget, the calculatedSpend drops
// to zero until AWS has new usage data to use for forecasting.
//
//    // Example sending a request using UpdateBudgetRequest.
//    req := client.UpdateBudgetRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) UpdateBudgetRequest(input *UpdateBudgetInput) UpdateBudgetRequest {
	op := &aws.Operation{
		Name:       opUpdateBudget,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateBudgetInput{}
	}

	req := c.newRequest(op, input, &UpdateBudgetOutput{})
	return UpdateBudgetRequest{Request: req, Input: input, Copy: c.UpdateBudgetRequest}
}

// UpdateBudgetRequest is the request type for the
// UpdateBudget API operation.
type UpdateBudgetRequest struct {
	*aws.Request
	Input *UpdateBudgetInput
	Copy  func(*UpdateBudgetInput) UpdateBudgetRequest
}

// Send marshals and sends the UpdateBudget API request.
func (r UpdateBudgetRequest) Send(ctx context.Context) (*UpdateBudgetResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateBudgetResponse{
		UpdateBudgetOutput: r.Request.Data.(*UpdateBudgetOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateBudgetResponse is the response type for the
// UpdateBudget API operation.
type UpdateBudgetResponse struct {
	*UpdateBudgetOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateBudget request.
func (r *UpdateBudgetResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
