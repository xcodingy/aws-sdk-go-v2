// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ses

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents a request to create an empty receipt rule set. You use receipt
// rule sets to receive email with Amazon SES. For more information, see the
// Amazon SES Developer Guide (http://docs.aws.amazon.com/ses/latest/DeveloperGuide/receiving-email-concepts.html).
// Please also see https://docs.aws.amazon.com/goto/WebAPI/email-2010-12-01/CreateReceiptRuleSetRequest
type CreateReceiptRuleSetInput struct {
	_ struct{} `type:"structure"`

	// The name of the rule set to create. The name must:
	//
	//    * This value can only contain ASCII letters (a-z, A-Z), numbers (0-9),
	//    underscores (_), or dashes (-).
	//
	//    * Start and end with a letter or number.
	//
	//    * Contain less than 64 characters.
	//
	// RuleSetName is a required field
	RuleSetName *string `type:"string" required:"true"`
}

// String returns the string representation
func (s CreateReceiptRuleSetInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateReceiptRuleSetInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateReceiptRuleSetInput"}

	if s.RuleSetName == nil {
		invalidParams.Add(aws.NewErrParamRequired("RuleSetName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// An empty element returned on a successful request.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/email-2010-12-01/CreateReceiptRuleSetResponse
type CreateReceiptRuleSetOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s CreateReceiptRuleSetOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateReceiptRuleSet = "CreateReceiptRuleSet"

// CreateReceiptRuleSetRequest returns a request value for making API operation for
// Amazon Simple Email Service.
//
// Creates an empty receipt rule set.
//
// For information about setting up receipt rule sets, see the Amazon SES Developer
// Guide (http://docs.aws.amazon.com/ses/latest/DeveloperGuide/receiving-email-receipt-rule-set.html).
//
// You can execute this operation no more than once per second.
//
//    // Example sending a request using CreateReceiptRuleSetRequest.
//    req := client.CreateReceiptRuleSetRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/email-2010-12-01/CreateReceiptRuleSet
func (c *Client) CreateReceiptRuleSetRequest(input *CreateReceiptRuleSetInput) CreateReceiptRuleSetRequest {
	op := &aws.Operation{
		Name:       opCreateReceiptRuleSet,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateReceiptRuleSetInput{}
	}

	req := c.newRequest(op, input, &CreateReceiptRuleSetOutput{})
	return CreateReceiptRuleSetRequest{Request: req, Input: input, Copy: c.CreateReceiptRuleSetRequest}
}

// CreateReceiptRuleSetRequest is the request type for the
// CreateReceiptRuleSet API operation.
type CreateReceiptRuleSetRequest struct {
	*aws.Request
	Input *CreateReceiptRuleSetInput
	Copy  func(*CreateReceiptRuleSetInput) CreateReceiptRuleSetRequest
}

// Send marshals and sends the CreateReceiptRuleSet API request.
func (r CreateReceiptRuleSetRequest) Send(ctx context.Context) (*CreateReceiptRuleSetResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateReceiptRuleSetResponse{
		CreateReceiptRuleSetOutput: r.Request.Data.(*CreateReceiptRuleSetOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateReceiptRuleSetResponse is the response type for the
// CreateReceiptRuleSet API operation.
type CreateReceiptRuleSetResponse struct {
	*CreateReceiptRuleSetOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateReceiptRuleSet request.
func (r *CreateReceiptRuleSetResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
