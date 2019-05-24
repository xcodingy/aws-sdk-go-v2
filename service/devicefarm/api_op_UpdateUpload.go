// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package devicefarm

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/devicefarm-2015-06-23/UpdateUploadRequest
type UpdateUploadInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the uploaded test spec.
	//
	// Arn is a required field
	Arn *string `locationName:"arn" min:"32" type:"string" required:"true"`

	// The upload's content type (for example, "application/x-yaml").
	ContentType *string `locationName:"contentType" type:"string"`

	// Set to true if the YAML file has changed and needs to be updated; otherwise,
	// set to false.
	EditContent *bool `locationName:"editContent" type:"boolean"`

	// The upload's test spec file name. The name should not contain the '/' character.
	// The test spec file name must end with the .yaml or .yml file extension.
	Name *string `locationName:"name" type:"string"`
}

// String returns the string representation
func (s UpdateUploadInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateUploadInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateUploadInput"}

	if s.Arn == nil {
		invalidParams.Add(aws.NewErrParamRequired("Arn"))
	}
	if s.Arn != nil && len(*s.Arn) < 32 {
		invalidParams.Add(aws.NewErrParamMinLen("Arn", 32))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/devicefarm-2015-06-23/UpdateUploadResult
type UpdateUploadOutput struct {
	_ struct{} `type:"structure"`

	// A test spec uploaded to Device Farm.
	Upload *Upload `locationName:"upload" type:"structure"`
}

// String returns the string representation
func (s UpdateUploadOutput) String() string {
	return awsutil.Prettify(s)
}

const opUpdateUpload = "UpdateUpload"

// UpdateUploadRequest returns a request value for making API operation for
// AWS Device Farm.
//
// Update an uploaded test specification (test spec).
//
//    // Example sending a request using UpdateUploadRequest.
//    req := client.UpdateUploadRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/devicefarm-2015-06-23/UpdateUpload
func (c *Client) UpdateUploadRequest(input *UpdateUploadInput) UpdateUploadRequest {
	op := &aws.Operation{
		Name:       opUpdateUpload,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateUploadInput{}
	}

	req := c.newRequest(op, input, &UpdateUploadOutput{})
	return UpdateUploadRequest{Request: req, Input: input, Copy: c.UpdateUploadRequest}
}

// UpdateUploadRequest is the request type for the
// UpdateUpload API operation.
type UpdateUploadRequest struct {
	*aws.Request
	Input *UpdateUploadInput
	Copy  func(*UpdateUploadInput) UpdateUploadRequest
}

// Send marshals and sends the UpdateUpload API request.
func (r UpdateUploadRequest) Send(ctx context.Context) (*UpdateUploadResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateUploadResponse{
		UpdateUploadOutput: r.Request.Data.(*UpdateUploadOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateUploadResponse is the response type for the
// UpdateUpload API operation.
type UpdateUploadResponse struct {
	*UpdateUploadOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateUpload request.
func (r *UpdateUploadResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
