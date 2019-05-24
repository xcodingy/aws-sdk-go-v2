// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package chime

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/chime-2018-05-01/BatchUpdateUserRequest
type BatchUpdateUserInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Chime account ID.
	//
	// AccountId is a required field
	AccountId *string `location:"uri" locationName:"accountId" type:"string" required:"true"`

	// The request containing the user IDs and details to update.
	//
	// UpdateUserRequestItems is a required field
	UpdateUserRequestItems []UpdateUserRequestItem `type:"list" required:"true"`
}

// String returns the string representation
func (s BatchUpdateUserInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *BatchUpdateUserInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "BatchUpdateUserInput"}

	if s.AccountId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AccountId"))
	}

	if s.UpdateUserRequestItems == nil {
		invalidParams.Add(aws.NewErrParamRequired("UpdateUserRequestItems"))
	}
	if s.UpdateUserRequestItems != nil {
		for i, v := range s.UpdateUserRequestItems {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "UpdateUserRequestItems", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s BatchUpdateUserInput) MarshalFields(e protocol.FieldEncoder) error {

	if len(s.UpdateUserRequestItems) > 0 {
		v := s.UpdateUserRequestItems

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "UpdateUserRequestItems", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.AccountId != nil {
		v := *s.AccountId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "accountId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/chime-2018-05-01/BatchUpdateUserResponse
type BatchUpdateUserOutput struct {
	_ struct{} `type:"structure"`

	// If the BatchUpdateUser action fails for one or more of the user IDs in the
	// request, a list of the user IDs is returned, along with error codes and error
	// messages.
	UserErrors []UserError `type:"list"`
}

// String returns the string representation
func (s BatchUpdateUserOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s BatchUpdateUserOutput) MarshalFields(e protocol.FieldEncoder) error {
	if len(s.UserErrors) > 0 {
		v := s.UserErrors

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "UserErrors", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	return nil
}

const opBatchUpdateUser = "BatchUpdateUser"

// BatchUpdateUserRequest returns a request value for making API operation for
// Amazon Chime.
//
// Updates user details within the UpdateUserRequestItem object for up to 20
// users for the specified Amazon Chime account. Currently, only LicenseType
// updates are supported for this action.
//
//    // Example sending a request using BatchUpdateUserRequest.
//    req := client.BatchUpdateUserRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/chime-2018-05-01/BatchUpdateUser
func (c *Client) BatchUpdateUserRequest(input *BatchUpdateUserInput) BatchUpdateUserRequest {
	op := &aws.Operation{
		Name:       opBatchUpdateUser,
		HTTPMethod: "POST",
		HTTPPath:   "/accounts/{accountId}/users",
	}

	if input == nil {
		input = &BatchUpdateUserInput{}
	}

	req := c.newRequest(op, input, &BatchUpdateUserOutput{})
	return BatchUpdateUserRequest{Request: req, Input: input, Copy: c.BatchUpdateUserRequest}
}

// BatchUpdateUserRequest is the request type for the
// BatchUpdateUser API operation.
type BatchUpdateUserRequest struct {
	*aws.Request
	Input *BatchUpdateUserInput
	Copy  func(*BatchUpdateUserInput) BatchUpdateUserRequest
}

// Send marshals and sends the BatchUpdateUser API request.
func (r BatchUpdateUserRequest) Send(ctx context.Context) (*BatchUpdateUserResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &BatchUpdateUserResponse{
		BatchUpdateUserOutput: r.Request.Data.(*BatchUpdateUserOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// BatchUpdateUserResponse is the response type for the
// BatchUpdateUser API operation.
type BatchUpdateUserResponse struct {
	*BatchUpdateUserOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// BatchUpdateUser request.
func (r *BatchUpdateUserResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
