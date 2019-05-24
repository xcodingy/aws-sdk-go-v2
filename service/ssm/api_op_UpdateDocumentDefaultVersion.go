// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ssm

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ssm-2014-11-06/UpdateDocumentDefaultVersionRequest
type UpdateDocumentDefaultVersionInput struct {
	_ struct{} `type:"structure"`

	// The version of a custom document that you want to set as the default version.
	//
	// DocumentVersion is a required field
	DocumentVersion *string `type:"string" required:"true"`

	// The name of a custom document that you want to set as the default version.
	//
	// Name is a required field
	Name *string `type:"string" required:"true"`
}

// String returns the string representation
func (s UpdateDocumentDefaultVersionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateDocumentDefaultVersionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateDocumentDefaultVersionInput"}

	if s.DocumentVersion == nil {
		invalidParams.Add(aws.NewErrParamRequired("DocumentVersion"))
	}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ssm-2014-11-06/UpdateDocumentDefaultVersionResult
type UpdateDocumentDefaultVersionOutput struct {
	_ struct{} `type:"structure"`

	// The description of a custom document that you want to set as the default
	// version.
	Description *DocumentDefaultVersionDescription `type:"structure"`
}

// String returns the string representation
func (s UpdateDocumentDefaultVersionOutput) String() string {
	return awsutil.Prettify(s)
}

const opUpdateDocumentDefaultVersion = "UpdateDocumentDefaultVersion"

// UpdateDocumentDefaultVersionRequest returns a request value for making API operation for
// Amazon Simple Systems Manager (SSM).
//
// Set the default version of a document.
//
//    // Example sending a request using UpdateDocumentDefaultVersionRequest.
//    req := client.UpdateDocumentDefaultVersionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ssm-2014-11-06/UpdateDocumentDefaultVersion
func (c *Client) UpdateDocumentDefaultVersionRequest(input *UpdateDocumentDefaultVersionInput) UpdateDocumentDefaultVersionRequest {
	op := &aws.Operation{
		Name:       opUpdateDocumentDefaultVersion,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateDocumentDefaultVersionInput{}
	}

	req := c.newRequest(op, input, &UpdateDocumentDefaultVersionOutput{})
	return UpdateDocumentDefaultVersionRequest{Request: req, Input: input, Copy: c.UpdateDocumentDefaultVersionRequest}
}

// UpdateDocumentDefaultVersionRequest is the request type for the
// UpdateDocumentDefaultVersion API operation.
type UpdateDocumentDefaultVersionRequest struct {
	*aws.Request
	Input *UpdateDocumentDefaultVersionInput
	Copy  func(*UpdateDocumentDefaultVersionInput) UpdateDocumentDefaultVersionRequest
}

// Send marshals and sends the UpdateDocumentDefaultVersion API request.
func (r UpdateDocumentDefaultVersionRequest) Send(ctx context.Context) (*UpdateDocumentDefaultVersionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateDocumentDefaultVersionResponse{
		UpdateDocumentDefaultVersionOutput: r.Request.Data.(*UpdateDocumentDefaultVersionOutput),
		response:                           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateDocumentDefaultVersionResponse is the response type for the
// UpdateDocumentDefaultVersion API operation.
type UpdateDocumentDefaultVersionResponse struct {
	*UpdateDocumentDefaultVersionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateDocumentDefaultVersion request.
func (r *UpdateDocumentDefaultVersionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
