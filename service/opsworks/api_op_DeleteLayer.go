// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package opsworks

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/opsworks-2013-02-18/DeleteLayerRequest
type DeleteLayerInput struct {
	_ struct{} `type:"structure"`

	// The layer ID.
	//
	// LayerId is a required field
	LayerId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteLayerInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteLayerInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteLayerInput"}

	if s.LayerId == nil {
		invalidParams.Add(aws.NewErrParamRequired("LayerId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/opsworks-2013-02-18/DeleteLayerOutput
type DeleteLayerOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteLayerOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteLayer = "DeleteLayer"

// DeleteLayerRequest returns a request value for making API operation for
// AWS OpsWorks.
//
// Deletes a specified layer. You must first stop and then delete all associated
// instances or unassign registered instances. For more information, see How
// to Delete a Layer (http://docs.aws.amazon.com/opsworks/latest/userguide/workinglayers-basics-delete.html).
//
// Required Permissions: To use this action, an IAM user must have a Manage
// permissions level for the stack, or an attached policy that explicitly grants
// permissions. For more information on user permissions, see Managing User
// Permissions (http://docs.aws.amazon.com/opsworks/latest/userguide/opsworks-security-users.html).
//
//    // Example sending a request using DeleteLayerRequest.
//    req := client.DeleteLayerRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/opsworks-2013-02-18/DeleteLayer
func (c *Client) DeleteLayerRequest(input *DeleteLayerInput) DeleteLayerRequest {
	op := &aws.Operation{
		Name:       opDeleteLayer,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteLayerInput{}
	}

	req := c.newRequest(op, input, &DeleteLayerOutput{})
	req.Handlers.Unmarshal.Remove(jsonrpc.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return DeleteLayerRequest{Request: req, Input: input, Copy: c.DeleteLayerRequest}
}

// DeleteLayerRequest is the request type for the
// DeleteLayer API operation.
type DeleteLayerRequest struct {
	*aws.Request
	Input *DeleteLayerInput
	Copy  func(*DeleteLayerInput) DeleteLayerRequest
}

// Send marshals and sends the DeleteLayer API request.
func (r DeleteLayerRequest) Send(ctx context.Context) (*DeleteLayerResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteLayerResponse{
		DeleteLayerOutput: r.Request.Data.(*DeleteLayerOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteLayerResponse is the response type for the
// DeleteLayer API operation.
type DeleteLayerResponse struct {
	*DeleteLayerOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteLayer request.
func (r *DeleteLayerResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
