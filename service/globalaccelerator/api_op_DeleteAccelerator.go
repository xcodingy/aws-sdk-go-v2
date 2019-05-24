// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package globalaccelerator

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/globalaccelerator-2018-08-08/DeleteAcceleratorRequest
type DeleteAcceleratorInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of an accelerator.
	//
	// AcceleratorArn is a required field
	AcceleratorArn *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteAcceleratorInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteAcceleratorInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteAcceleratorInput"}

	if s.AcceleratorArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("AcceleratorArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/globalaccelerator-2018-08-08/DeleteAcceleratorOutput
type DeleteAcceleratorOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteAcceleratorOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteAccelerator = "DeleteAccelerator"

// DeleteAcceleratorRequest returns a request value for making API operation for
// AWS Global Accelerator.
//
// Delete an accelerator. Note: before you can delete an accelerator, you must
// disable it and remove all dependent resources (listeners and endpoint groups).
//
//    // Example sending a request using DeleteAcceleratorRequest.
//    req := client.DeleteAcceleratorRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/globalaccelerator-2018-08-08/DeleteAccelerator
func (c *Client) DeleteAcceleratorRequest(input *DeleteAcceleratorInput) DeleteAcceleratorRequest {
	op := &aws.Operation{
		Name:       opDeleteAccelerator,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteAcceleratorInput{}
	}

	req := c.newRequest(op, input, &DeleteAcceleratorOutput{})
	req.Handlers.Unmarshal.Remove(jsonrpc.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return DeleteAcceleratorRequest{Request: req, Input: input, Copy: c.DeleteAcceleratorRequest}
}

// DeleteAcceleratorRequest is the request type for the
// DeleteAccelerator API operation.
type DeleteAcceleratorRequest struct {
	*aws.Request
	Input *DeleteAcceleratorInput
	Copy  func(*DeleteAcceleratorInput) DeleteAcceleratorRequest
}

// Send marshals and sends the DeleteAccelerator API request.
func (r DeleteAcceleratorRequest) Send(ctx context.Context) (*DeleteAcceleratorResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteAcceleratorResponse{
		DeleteAcceleratorOutput: r.Request.Data.(*DeleteAcceleratorOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteAcceleratorResponse is the response type for the
// DeleteAccelerator API operation.
type DeleteAcceleratorResponse struct {
	*DeleteAcceleratorOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteAccelerator request.
func (r *DeleteAcceleratorResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
