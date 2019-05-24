// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ses

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/query"
)

// Represents a request to delete an existing custom verification email template.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/email-2010-12-01/DeleteCustomVerificationEmailTemplateRequest
type DeleteCustomVerificationEmailTemplateInput struct {
	_ struct{} `type:"structure"`

	// The name of the custom verification email template that you want to delete.
	//
	// TemplateName is a required field
	TemplateName *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteCustomVerificationEmailTemplateInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteCustomVerificationEmailTemplateInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteCustomVerificationEmailTemplateInput"}

	if s.TemplateName == nil {
		invalidParams.Add(aws.NewErrParamRequired("TemplateName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/email-2010-12-01/DeleteCustomVerificationEmailTemplateOutput
type DeleteCustomVerificationEmailTemplateOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteCustomVerificationEmailTemplateOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteCustomVerificationEmailTemplate = "DeleteCustomVerificationEmailTemplate"

// DeleteCustomVerificationEmailTemplateRequest returns a request value for making API operation for
// Amazon Simple Email Service.
//
// Deletes an existing custom verification email template.
//
// For more information about custom verification email templates, see Using
// Custom Verification Email Templates (http://docs.aws.amazon.com/ses/latest/DeveloperGuide/custom-verification-emails.html)
// in the Amazon SES Developer Guide.
//
// You can execute this operation no more than once per second.
//
//    // Example sending a request using DeleteCustomVerificationEmailTemplateRequest.
//    req := client.DeleteCustomVerificationEmailTemplateRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/email-2010-12-01/DeleteCustomVerificationEmailTemplate
func (c *Client) DeleteCustomVerificationEmailTemplateRequest(input *DeleteCustomVerificationEmailTemplateInput) DeleteCustomVerificationEmailTemplateRequest {
	op := &aws.Operation{
		Name:       opDeleteCustomVerificationEmailTemplate,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteCustomVerificationEmailTemplateInput{}
	}

	req := c.newRequest(op, input, &DeleteCustomVerificationEmailTemplateOutput{})
	req.Handlers.Unmarshal.Remove(query.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return DeleteCustomVerificationEmailTemplateRequest{Request: req, Input: input, Copy: c.DeleteCustomVerificationEmailTemplateRequest}
}

// DeleteCustomVerificationEmailTemplateRequest is the request type for the
// DeleteCustomVerificationEmailTemplate API operation.
type DeleteCustomVerificationEmailTemplateRequest struct {
	*aws.Request
	Input *DeleteCustomVerificationEmailTemplateInput
	Copy  func(*DeleteCustomVerificationEmailTemplateInput) DeleteCustomVerificationEmailTemplateRequest
}

// Send marshals and sends the DeleteCustomVerificationEmailTemplate API request.
func (r DeleteCustomVerificationEmailTemplateRequest) Send(ctx context.Context) (*DeleteCustomVerificationEmailTemplateResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteCustomVerificationEmailTemplateResponse{
		DeleteCustomVerificationEmailTemplateOutput: r.Request.Data.(*DeleteCustomVerificationEmailTemplateOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteCustomVerificationEmailTemplateResponse is the response type for the
// DeleteCustomVerificationEmailTemplate API operation.
type DeleteCustomVerificationEmailTemplateResponse struct {
	*DeleteCustomVerificationEmailTemplateOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteCustomVerificationEmailTemplate request.
func (r *DeleteCustomVerificationEmailTemplateResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
