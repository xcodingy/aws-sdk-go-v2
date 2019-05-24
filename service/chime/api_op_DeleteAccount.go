// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package chime

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/chime-2018-05-01/DeleteAccountRequest
type DeleteAccountInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Chime account ID.
	//
	// AccountId is a required field
	AccountId *string `location:"uri" locationName:"accountId" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteAccountInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteAccountInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteAccountInput"}

	if s.AccountId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AccountId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteAccountInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.AccountId != nil {
		v := *s.AccountId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "accountId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/chime-2018-05-01/DeleteAccountResponse
type DeleteAccountOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteAccountOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteAccountOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opDeleteAccount = "DeleteAccount"

// DeleteAccountRequest returns a request value for making API operation for
// Amazon Chime.
//
// Deletes the specified Amazon Chime account. You must suspend all users before
// deleting a Team account. You can use the BatchSuspendUser action to do so.
//
// For EnterpriseLWA and EnterpriseAD accounts, you must release the claimed
// domains for your Amazon Chime account before deletion. As soon as you release
// the domain, all users under that account are suspended.
//
// Deleted accounts appear in your Disabled accounts list for 90 days. To restore
// a deleted account from your Disabled accounts list, you must contact AWS
// Support.
//
// After 90 days, deleted accounts are permanently removed from your Disabled
// accounts list.
//
//    // Example sending a request using DeleteAccountRequest.
//    req := client.DeleteAccountRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/chime-2018-05-01/DeleteAccount
func (c *Client) DeleteAccountRequest(input *DeleteAccountInput) DeleteAccountRequest {
	op := &aws.Operation{
		Name:       opDeleteAccount,
		HTTPMethod: "DELETE",
		HTTPPath:   "/accounts/{accountId}",
	}

	if input == nil {
		input = &DeleteAccountInput{}
	}

	req := c.newRequest(op, input, &DeleteAccountOutput{})
	return DeleteAccountRequest{Request: req, Input: input, Copy: c.DeleteAccountRequest}
}

// DeleteAccountRequest is the request type for the
// DeleteAccount API operation.
type DeleteAccountRequest struct {
	*aws.Request
	Input *DeleteAccountInput
	Copy  func(*DeleteAccountInput) DeleteAccountRequest
}

// Send marshals and sends the DeleteAccount API request.
func (r DeleteAccountRequest) Send(ctx context.Context) (*DeleteAccountResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteAccountResponse{
		DeleteAccountOutput: r.Request.Data.(*DeleteAccountOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteAccountResponse is the response type for the
// DeleteAccount API operation.
type DeleteAccountResponse struct {
	*DeleteAccountOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteAccount request.
func (r *DeleteAccountResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
