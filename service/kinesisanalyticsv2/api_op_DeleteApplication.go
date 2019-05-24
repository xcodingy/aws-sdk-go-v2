// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package kinesisanalyticsv2

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/kinesisanalyticsv2-2018-05-23/DeleteApplicationRequest
type DeleteApplicationInput struct {
	_ struct{} `type:"structure"`

	// The name of the application to delete.
	//
	// ApplicationName is a required field
	ApplicationName *string `min:"1" type:"string" required:"true"`

	// Use the DescribeApplication operation to get this value.
	//
	// CreateTimestamp is a required field
	CreateTimestamp *time.Time `type:"timestamp" timestampFormat:"unix" required:"true"`
}

// String returns the string representation
func (s DeleteApplicationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteApplicationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteApplicationInput"}

	if s.ApplicationName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ApplicationName"))
	}
	if s.ApplicationName != nil && len(*s.ApplicationName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ApplicationName", 1))
	}

	if s.CreateTimestamp == nil {
		invalidParams.Add(aws.NewErrParamRequired("CreateTimestamp"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/kinesisanalyticsv2-2018-05-23/DeleteApplicationResponse
type DeleteApplicationOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteApplicationOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteApplication = "DeleteApplication"

// DeleteApplicationRequest returns a request value for making API operation for
// Amazon Kinesis Analytics.
//
// Deletes the specified application. Kinesis Data Analytics halts application
// execution and deletes the application.
//
//    // Example sending a request using DeleteApplicationRequest.
//    req := client.DeleteApplicationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/kinesisanalyticsv2-2018-05-23/DeleteApplication
func (c *Client) DeleteApplicationRequest(input *DeleteApplicationInput) DeleteApplicationRequest {
	op := &aws.Operation{
		Name:       opDeleteApplication,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteApplicationInput{}
	}

	req := c.newRequest(op, input, &DeleteApplicationOutput{})
	return DeleteApplicationRequest{Request: req, Input: input, Copy: c.DeleteApplicationRequest}
}

// DeleteApplicationRequest is the request type for the
// DeleteApplication API operation.
type DeleteApplicationRequest struct {
	*aws.Request
	Input *DeleteApplicationInput
	Copy  func(*DeleteApplicationInput) DeleteApplicationRequest
}

// Send marshals and sends the DeleteApplication API request.
func (r DeleteApplicationRequest) Send(ctx context.Context) (*DeleteApplicationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteApplicationResponse{
		DeleteApplicationOutput: r.Request.Data.(*DeleteApplicationOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteApplicationResponse is the response type for the
// DeleteApplication API operation.
type DeleteApplicationResponse struct {
	*DeleteApplicationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteApplication request.
func (r *DeleteApplicationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
