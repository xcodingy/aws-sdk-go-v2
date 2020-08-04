// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package workmail

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type UpdateMailboxQuotaInput struct {
	_ struct{} `type:"structure"`

	// The updated mailbox quota, in MB, for the specified user.
	//
	// MailboxQuota is a required field
	MailboxQuota *int64 `min:"1" type:"integer" required:"true"`

	// The identifier for the organization that contains the user for whom to update
	// the mailbox quota.
	//
	// OrganizationId is a required field
	OrganizationId *string `type:"string" required:"true"`

	// The identifer for the user for whom to update the mailbox quota.
	//
	// UserId is a required field
	UserId *string `min:"12" type:"string" required:"true"`
}

// String returns the string representation
func (s UpdateMailboxQuotaInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateMailboxQuotaInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateMailboxQuotaInput"}

	if s.MailboxQuota == nil {
		invalidParams.Add(aws.NewErrParamRequired("MailboxQuota"))
	}
	if s.MailboxQuota != nil && *s.MailboxQuota < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MailboxQuota", 1))
	}

	if s.OrganizationId == nil {
		invalidParams.Add(aws.NewErrParamRequired("OrganizationId"))
	}

	if s.UserId == nil {
		invalidParams.Add(aws.NewErrParamRequired("UserId"))
	}
	if s.UserId != nil && len(*s.UserId) < 12 {
		invalidParams.Add(aws.NewErrParamMinLen("UserId", 12))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type UpdateMailboxQuotaOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s UpdateMailboxQuotaOutput) String() string {
	return awsutil.Prettify(s)
}

const opUpdateMailboxQuota = "UpdateMailboxQuota"

// UpdateMailboxQuotaRequest returns a request value for making API operation for
// Amazon WorkMail.
//
// Updates a user's current mailbox quota for a specified organization and user.
//
//    // Example sending a request using UpdateMailboxQuotaRequest.
//    req := client.UpdateMailboxQuotaRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/workmail-2017-10-01/UpdateMailboxQuota
func (c *Client) UpdateMailboxQuotaRequest(input *UpdateMailboxQuotaInput) UpdateMailboxQuotaRequest {
	op := &aws.Operation{
		Name:       opUpdateMailboxQuota,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateMailboxQuotaInput{}
	}

	req := c.newRequest(op, input, &UpdateMailboxQuotaOutput{})

	return UpdateMailboxQuotaRequest{Request: req, Input: input, Copy: c.UpdateMailboxQuotaRequest}
}

// UpdateMailboxQuotaRequest is the request type for the
// UpdateMailboxQuota API operation.
type UpdateMailboxQuotaRequest struct {
	*aws.Request
	Input *UpdateMailboxQuotaInput
	Copy  func(*UpdateMailboxQuotaInput) UpdateMailboxQuotaRequest
}

// Send marshals and sends the UpdateMailboxQuota API request.
func (r UpdateMailboxQuotaRequest) Send(ctx context.Context) (*UpdateMailboxQuotaResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateMailboxQuotaResponse{
		UpdateMailboxQuotaOutput: r.Request.Data.(*UpdateMailboxQuotaOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateMailboxQuotaResponse is the response type for the
// UpdateMailboxQuota API operation.
type UpdateMailboxQuotaResponse struct {
	*UpdateMailboxQuotaOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateMailboxQuota request.
func (r *UpdateMailboxQuotaResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}