// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package route53domains

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// A request to set the transfer lock for the specified domain.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/route53domains-2014-05-15/EnableDomainTransferLockRequest
type EnableDomainTransferLockInput struct {
	_ struct{} `type:"structure"`

	// The name of the domain that you want to set the transfer lock for.
	//
	// DomainName is a required field
	DomainName *string `type:"string" required:"true"`
}

// String returns the string representation
func (s EnableDomainTransferLockInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *EnableDomainTransferLockInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "EnableDomainTransferLockInput"}

	if s.DomainName == nil {
		invalidParams.Add(aws.NewErrParamRequired("DomainName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// The EnableDomainTransferLock response includes the following elements.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/route53domains-2014-05-15/EnableDomainTransferLockResponse
type EnableDomainTransferLockOutput struct {
	_ struct{} `type:"structure"`

	// Identifier for tracking the progress of the request. To use this ID to query
	// the operation status, use GetOperationDetail.
	//
	// OperationId is a required field
	OperationId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s EnableDomainTransferLockOutput) String() string {
	return awsutil.Prettify(s)
}

const opEnableDomainTransferLock = "EnableDomainTransferLock"

// EnableDomainTransferLockRequest returns a request value for making API operation for
// Amazon Route 53 Domains.
//
// This operation sets the transfer lock on the domain (specifically the clientTransferProhibited
// status) to prevent domain transfers. Successful submission returns an operation
// ID that you can use to track the progress and completion of the action. If
// the request is not completed successfully, the domain registrant will be
// notified by email.
//
//    // Example sending a request using EnableDomainTransferLockRequest.
//    req := client.EnableDomainTransferLockRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/route53domains-2014-05-15/EnableDomainTransferLock
func (c *Client) EnableDomainTransferLockRequest(input *EnableDomainTransferLockInput) EnableDomainTransferLockRequest {
	op := &aws.Operation{
		Name:       opEnableDomainTransferLock,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &EnableDomainTransferLockInput{}
	}

	req := c.newRequest(op, input, &EnableDomainTransferLockOutput{})
	return EnableDomainTransferLockRequest{Request: req, Input: input, Copy: c.EnableDomainTransferLockRequest}
}

// EnableDomainTransferLockRequest is the request type for the
// EnableDomainTransferLock API operation.
type EnableDomainTransferLockRequest struct {
	*aws.Request
	Input *EnableDomainTransferLockInput
	Copy  func(*EnableDomainTransferLockInput) EnableDomainTransferLockRequest
}

// Send marshals and sends the EnableDomainTransferLock API request.
func (r EnableDomainTransferLockRequest) Send(ctx context.Context) (*EnableDomainTransferLockResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &EnableDomainTransferLockResponse{
		EnableDomainTransferLockOutput: r.Request.Data.(*EnableDomainTransferLockOutput),
		response:                       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// EnableDomainTransferLockResponse is the response type for the
// EnableDomainTransferLock API operation.
type EnableDomainTransferLockResponse struct {
	*EnableDomainTransferLockOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// EnableDomainTransferLock request.
func (r *EnableDomainTransferLockResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
