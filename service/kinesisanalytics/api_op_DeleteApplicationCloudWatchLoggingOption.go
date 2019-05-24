// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package kinesisanalytics

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/kinesisanalytics-2015-08-14/DeleteApplicationCloudWatchLoggingOptionRequest
type DeleteApplicationCloudWatchLoggingOptionInput struct {
	_ struct{} `type:"structure"`

	// The Kinesis Analytics application name.
	//
	// ApplicationName is a required field
	ApplicationName *string `min:"1" type:"string" required:"true"`

	// The CloudWatchLoggingOptionId of the CloudWatch logging option to delete.
	// You can get the CloudWatchLoggingOptionId by using the DescribeApplication
	// operation.
	//
	// CloudWatchLoggingOptionId is a required field
	CloudWatchLoggingOptionId *string `min:"1" type:"string" required:"true"`

	// The version ID of the Kinesis Analytics application.
	//
	// CurrentApplicationVersionId is a required field
	CurrentApplicationVersionId *int64 `min:"1" type:"long" required:"true"`
}

// String returns the string representation
func (s DeleteApplicationCloudWatchLoggingOptionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteApplicationCloudWatchLoggingOptionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteApplicationCloudWatchLoggingOptionInput"}

	if s.ApplicationName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ApplicationName"))
	}
	if s.ApplicationName != nil && len(*s.ApplicationName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ApplicationName", 1))
	}

	if s.CloudWatchLoggingOptionId == nil {
		invalidParams.Add(aws.NewErrParamRequired("CloudWatchLoggingOptionId"))
	}
	if s.CloudWatchLoggingOptionId != nil && len(*s.CloudWatchLoggingOptionId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("CloudWatchLoggingOptionId", 1))
	}

	if s.CurrentApplicationVersionId == nil {
		invalidParams.Add(aws.NewErrParamRequired("CurrentApplicationVersionId"))
	}
	if s.CurrentApplicationVersionId != nil && *s.CurrentApplicationVersionId < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("CurrentApplicationVersionId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/kinesisanalytics-2015-08-14/DeleteApplicationCloudWatchLoggingOptionResponse
type DeleteApplicationCloudWatchLoggingOptionOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteApplicationCloudWatchLoggingOptionOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteApplicationCloudWatchLoggingOption = "DeleteApplicationCloudWatchLoggingOption"

// DeleteApplicationCloudWatchLoggingOptionRequest returns a request value for making API operation for
// Amazon Kinesis Analytics.
//
// Deletes a CloudWatch log stream from an application. For more information
// about using CloudWatch log streams with Amazon Kinesis Analytics applications,
// see Working with Amazon CloudWatch Logs (http://docs.aws.amazon.com/kinesisanalytics/latest/dev/cloudwatch-logs.html).
//
//    // Example sending a request using DeleteApplicationCloudWatchLoggingOptionRequest.
//    req := client.DeleteApplicationCloudWatchLoggingOptionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/kinesisanalytics-2015-08-14/DeleteApplicationCloudWatchLoggingOption
func (c *Client) DeleteApplicationCloudWatchLoggingOptionRequest(input *DeleteApplicationCloudWatchLoggingOptionInput) DeleteApplicationCloudWatchLoggingOptionRequest {
	op := &aws.Operation{
		Name:       opDeleteApplicationCloudWatchLoggingOption,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteApplicationCloudWatchLoggingOptionInput{}
	}

	req := c.newRequest(op, input, &DeleteApplicationCloudWatchLoggingOptionOutput{})
	return DeleteApplicationCloudWatchLoggingOptionRequest{Request: req, Input: input, Copy: c.DeleteApplicationCloudWatchLoggingOptionRequest}
}

// DeleteApplicationCloudWatchLoggingOptionRequest is the request type for the
// DeleteApplicationCloudWatchLoggingOption API operation.
type DeleteApplicationCloudWatchLoggingOptionRequest struct {
	*aws.Request
	Input *DeleteApplicationCloudWatchLoggingOptionInput
	Copy  func(*DeleteApplicationCloudWatchLoggingOptionInput) DeleteApplicationCloudWatchLoggingOptionRequest
}

// Send marshals and sends the DeleteApplicationCloudWatchLoggingOption API request.
func (r DeleteApplicationCloudWatchLoggingOptionRequest) Send(ctx context.Context) (*DeleteApplicationCloudWatchLoggingOptionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteApplicationCloudWatchLoggingOptionResponse{
		DeleteApplicationCloudWatchLoggingOptionOutput: r.Request.Data.(*DeleteApplicationCloudWatchLoggingOptionOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteApplicationCloudWatchLoggingOptionResponse is the response type for the
// DeleteApplicationCloudWatchLoggingOption API operation.
type DeleteApplicationCloudWatchLoggingOptionResponse struct {
	*DeleteApplicationCloudWatchLoggingOptionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteApplicationCloudWatchLoggingOption request.
func (r *DeleteApplicationCloudWatchLoggingOptionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
