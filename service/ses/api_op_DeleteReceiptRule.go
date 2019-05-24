// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ses

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents a request to delete a receipt rule. You use receipt rules to receive
// email with Amazon SES. For more information, see the Amazon SES Developer
// Guide (http://docs.aws.amazon.com/ses/latest/DeveloperGuide/receiving-email-concepts.html).
// Please also see https://docs.aws.amazon.com/goto/WebAPI/email-2010-12-01/DeleteReceiptRuleRequest
type DeleteReceiptRuleInput struct {
	_ struct{} `type:"structure"`

	// The name of the receipt rule to delete.
	//
	// RuleName is a required field
	RuleName *string `type:"string" required:"true"`

	// The name of the receipt rule set that contains the receipt rule to delete.
	//
	// RuleSetName is a required field
	RuleSetName *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteReceiptRuleInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteReceiptRuleInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteReceiptRuleInput"}

	if s.RuleName == nil {
		invalidParams.Add(aws.NewErrParamRequired("RuleName"))
	}

	if s.RuleSetName == nil {
		invalidParams.Add(aws.NewErrParamRequired("RuleSetName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// An empty element returned on a successful request.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/email-2010-12-01/DeleteReceiptRuleResponse
type DeleteReceiptRuleOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteReceiptRuleOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteReceiptRule = "DeleteReceiptRule"

// DeleteReceiptRuleRequest returns a request value for making API operation for
// Amazon Simple Email Service.
//
// Deletes the specified receipt rule.
//
// For information about managing receipt rules, see the Amazon SES Developer
// Guide (http://docs.aws.amazon.com/ses/latest/DeveloperGuide/receiving-email-managing-receipt-rules.html).
//
// You can execute this operation no more than once per second.
//
//    // Example sending a request using DeleteReceiptRuleRequest.
//    req := client.DeleteReceiptRuleRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/email-2010-12-01/DeleteReceiptRule
func (c *Client) DeleteReceiptRuleRequest(input *DeleteReceiptRuleInput) DeleteReceiptRuleRequest {
	op := &aws.Operation{
		Name:       opDeleteReceiptRule,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteReceiptRuleInput{}
	}

	req := c.newRequest(op, input, &DeleteReceiptRuleOutput{})
	return DeleteReceiptRuleRequest{Request: req, Input: input, Copy: c.DeleteReceiptRuleRequest}
}

// DeleteReceiptRuleRequest is the request type for the
// DeleteReceiptRule API operation.
type DeleteReceiptRuleRequest struct {
	*aws.Request
	Input *DeleteReceiptRuleInput
	Copy  func(*DeleteReceiptRuleInput) DeleteReceiptRuleRequest
}

// Send marshals and sends the DeleteReceiptRule API request.
func (r DeleteReceiptRuleRequest) Send(ctx context.Context) (*DeleteReceiptRuleResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteReceiptRuleResponse{
		DeleteReceiptRuleOutput: r.Request.Data.(*DeleteReceiptRuleOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteReceiptRuleResponse is the response type for the
// DeleteReceiptRule API operation.
type DeleteReceiptRuleResponse struct {
	*DeleteReceiptRuleOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteReceiptRule request.
func (r *DeleteReceiptRuleResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
