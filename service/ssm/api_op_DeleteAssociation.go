// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ssm

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ssm-2014-11-06/DeleteAssociationRequest
type DeleteAssociationInput struct {
	_ struct{} `type:"structure"`

	// The association ID that you want to delete.
	AssociationId *string `type:"string"`

	// The ID of the instance.
	InstanceId *string `type:"string"`

	// The name of the Systems Manager document.
	Name *string `type:"string"`
}

// String returns the string representation
func (s DeleteAssociationInput) String() string {
	return awsutil.Prettify(s)
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ssm-2014-11-06/DeleteAssociationResult
type DeleteAssociationOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteAssociationOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteAssociation = "DeleteAssociation"

// DeleteAssociationRequest returns a request value for making API operation for
// Amazon Simple Systems Manager (SSM).
//
// Disassociates the specified Systems Manager document from the specified instance.
//
// When you disassociate a document from an instance, it does not change the
// configuration of the instance. To change the configuration state of an instance
// after you disassociate a document, you must create a new document with the
// desired configuration and associate it with the instance.
//
//    // Example sending a request using DeleteAssociationRequest.
//    req := client.DeleteAssociationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ssm-2014-11-06/DeleteAssociation
func (c *Client) DeleteAssociationRequest(input *DeleteAssociationInput) DeleteAssociationRequest {
	op := &aws.Operation{
		Name:       opDeleteAssociation,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteAssociationInput{}
	}

	req := c.newRequest(op, input, &DeleteAssociationOutput{})
	return DeleteAssociationRequest{Request: req, Input: input, Copy: c.DeleteAssociationRequest}
}

// DeleteAssociationRequest is the request type for the
// DeleteAssociation API operation.
type DeleteAssociationRequest struct {
	*aws.Request
	Input *DeleteAssociationInput
	Copy  func(*DeleteAssociationInput) DeleteAssociationRequest
}

// Send marshals and sends the DeleteAssociation API request.
func (r DeleteAssociationRequest) Send(ctx context.Context) (*DeleteAssociationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteAssociationResponse{
		DeleteAssociationOutput: r.Request.Data.(*DeleteAssociationOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteAssociationResponse is the response type for the
// DeleteAssociation API operation.
type DeleteAssociationResponse struct {
	*DeleteAssociationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteAssociation request.
func (r *DeleteAssociationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
