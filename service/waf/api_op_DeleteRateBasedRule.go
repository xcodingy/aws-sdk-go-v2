// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package waf

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/waf-2015-08-24/DeleteRateBasedRuleRequest
type DeleteRateBasedRuleInput struct {
	_ struct{} `type:"structure"`

	// The value returned by the most recent call to GetChangeToken.
	//
	// ChangeToken is a required field
	ChangeToken *string `min:"1" type:"string" required:"true"`

	// The RuleId of the RateBasedRule that you want to delete. RuleId is returned
	// by CreateRateBasedRule and by ListRateBasedRules.
	//
	// RuleId is a required field
	RuleId *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteRateBasedRuleInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteRateBasedRuleInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteRateBasedRuleInput"}

	if s.ChangeToken == nil {
		invalidParams.Add(aws.NewErrParamRequired("ChangeToken"))
	}
	if s.ChangeToken != nil && len(*s.ChangeToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ChangeToken", 1))
	}

	if s.RuleId == nil {
		invalidParams.Add(aws.NewErrParamRequired("RuleId"))
	}
	if s.RuleId != nil && len(*s.RuleId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("RuleId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/waf-2015-08-24/DeleteRateBasedRuleResponse
type DeleteRateBasedRuleOutput struct {
	_ struct{} `type:"structure"`

	// The ChangeToken that you used to submit the DeleteRateBasedRule request.
	// You can also use this value to query the status of the request. For more
	// information, see GetChangeTokenStatus.
	ChangeToken *string `min:"1" type:"string"`
}

// String returns the string representation
func (s DeleteRateBasedRuleOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteRateBasedRule = "DeleteRateBasedRule"

// DeleteRateBasedRuleRequest returns a request value for making API operation for
// AWS WAF.
//
// Permanently deletes a RateBasedRule. You can't delete a rule if it's still
// used in any WebACL objects or if it still includes any predicates, such as
// ByteMatchSet objects.
//
// If you just want to remove a rule from a WebACL, use UpdateWebACL.
//
// To permanently delete a RateBasedRule from AWS WAF, perform the following
// steps:
//
// Update the RateBasedRule to remove predicates, if any. For more information,
// see UpdateRateBasedRule.
//
// Use GetChangeToken to get the change token that you provide in the ChangeToken
// parameter of a DeleteRateBasedRule request.
//
// Submit a DeleteRateBasedRule request.
//
//    // Example sending a request using DeleteRateBasedRuleRequest.
//    req := client.DeleteRateBasedRuleRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/waf-2015-08-24/DeleteRateBasedRule
func (c *Client) DeleteRateBasedRuleRequest(input *DeleteRateBasedRuleInput) DeleteRateBasedRuleRequest {
	op := &aws.Operation{
		Name:       opDeleteRateBasedRule,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteRateBasedRuleInput{}
	}

	req := c.newRequest(op, input, &DeleteRateBasedRuleOutput{})
	return DeleteRateBasedRuleRequest{Request: req, Input: input, Copy: c.DeleteRateBasedRuleRequest}
}

// DeleteRateBasedRuleRequest is the request type for the
// DeleteRateBasedRule API operation.
type DeleteRateBasedRuleRequest struct {
	*aws.Request
	Input *DeleteRateBasedRuleInput
	Copy  func(*DeleteRateBasedRuleInput) DeleteRateBasedRuleRequest
}

// Send marshals and sends the DeleteRateBasedRule API request.
func (r DeleteRateBasedRuleRequest) Send(ctx context.Context) (*DeleteRateBasedRuleResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteRateBasedRuleResponse{
		DeleteRateBasedRuleOutput: r.Request.Data.(*DeleteRateBasedRuleOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteRateBasedRuleResponse is the response type for the
// DeleteRateBasedRule API operation.
type DeleteRateBasedRuleResponse struct {
	*DeleteRateBasedRuleOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteRateBasedRule request.
func (r *DeleteRateBasedRuleResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
