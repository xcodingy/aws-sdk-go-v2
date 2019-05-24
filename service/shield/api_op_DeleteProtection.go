// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package shield

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/shield-2016-06-02/DeleteProtectionRequest
type DeleteProtectionInput struct {
	_ struct{} `type:"structure"`

	// The unique identifier (ID) for the Protection object to be deleted.
	//
	// ProtectionId is a required field
	ProtectionId *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteProtectionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteProtectionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteProtectionInput"}

	if s.ProtectionId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ProtectionId"))
	}
	if s.ProtectionId != nil && len(*s.ProtectionId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ProtectionId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/shield-2016-06-02/DeleteProtectionResponse
type DeleteProtectionOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteProtectionOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteProtection = "DeleteProtection"

// DeleteProtectionRequest returns a request value for making API operation for
// AWS Shield.
//
// Deletes an AWS Shield Advanced Protection.
//
//    // Example sending a request using DeleteProtectionRequest.
//    req := client.DeleteProtectionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/shield-2016-06-02/DeleteProtection
func (c *Client) DeleteProtectionRequest(input *DeleteProtectionInput) DeleteProtectionRequest {
	op := &aws.Operation{
		Name:       opDeleteProtection,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteProtectionInput{}
	}

	req := c.newRequest(op, input, &DeleteProtectionOutput{})
	return DeleteProtectionRequest{Request: req, Input: input, Copy: c.DeleteProtectionRequest}
}

// DeleteProtectionRequest is the request type for the
// DeleteProtection API operation.
type DeleteProtectionRequest struct {
	*aws.Request
	Input *DeleteProtectionInput
	Copy  func(*DeleteProtectionInput) DeleteProtectionRequest
}

// Send marshals and sends the DeleteProtection API request.
func (r DeleteProtectionRequest) Send(ctx context.Context) (*DeleteProtectionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteProtectionResponse{
		DeleteProtectionOutput: r.Request.Data.(*DeleteProtectionOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteProtectionResponse is the response type for the
// DeleteProtection API operation.
type DeleteProtectionResponse struct {
	*DeleteProtectionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteProtection request.
func (r *DeleteProtectionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
