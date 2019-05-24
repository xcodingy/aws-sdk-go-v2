// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package glue

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/glue-2017-03-31/DeleteClassifierRequest
type DeleteClassifierInput struct {
	_ struct{} `type:"structure"`

	// Name of the classifier to remove.
	//
	// Name is a required field
	Name *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteClassifierInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteClassifierInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteClassifierInput"}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}
	if s.Name != nil && len(*s.Name) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Name", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/glue-2017-03-31/DeleteClassifierResponse
type DeleteClassifierOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteClassifierOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteClassifier = "DeleteClassifier"

// DeleteClassifierRequest returns a request value for making API operation for
// AWS Glue.
//
// Removes a classifier from the Data Catalog.
//
//    // Example sending a request using DeleteClassifierRequest.
//    req := client.DeleteClassifierRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/glue-2017-03-31/DeleteClassifier
func (c *Client) DeleteClassifierRequest(input *DeleteClassifierInput) DeleteClassifierRequest {
	op := &aws.Operation{
		Name:       opDeleteClassifier,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteClassifierInput{}
	}

	req := c.newRequest(op, input, &DeleteClassifierOutput{})
	return DeleteClassifierRequest{Request: req, Input: input, Copy: c.DeleteClassifierRequest}
}

// DeleteClassifierRequest is the request type for the
// DeleteClassifier API operation.
type DeleteClassifierRequest struct {
	*aws.Request
	Input *DeleteClassifierInput
	Copy  func(*DeleteClassifierInput) DeleteClassifierRequest
}

// Send marshals and sends the DeleteClassifier API request.
func (r DeleteClassifierRequest) Send(ctx context.Context) (*DeleteClassifierResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteClassifierResponse{
		DeleteClassifierOutput: r.Request.Data.(*DeleteClassifierOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteClassifierResponse is the response type for the
// DeleteClassifier API operation.
type DeleteClassifierResponse struct {
	*DeleteClassifierOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteClassifier request.
func (r *DeleteClassifierResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
