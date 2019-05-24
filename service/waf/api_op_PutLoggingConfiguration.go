// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package waf

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/waf-2015-08-24/PutLoggingConfigurationRequest
type PutLoggingConfigurationInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Kinesis Data Firehose that contains the inspected traffic information,
	// the redacted fields details, and the Amazon Resource Name (ARN) of the web
	// ACL to monitor.
	//
	// LoggingConfiguration is a required field
	LoggingConfiguration *LoggingConfiguration `type:"structure" required:"true"`
}

// String returns the string representation
func (s PutLoggingConfigurationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PutLoggingConfigurationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "PutLoggingConfigurationInput"}

	if s.LoggingConfiguration == nil {
		invalidParams.Add(aws.NewErrParamRequired("LoggingConfiguration"))
	}
	if s.LoggingConfiguration != nil {
		if err := s.LoggingConfiguration.Validate(); err != nil {
			invalidParams.AddNested("LoggingConfiguration", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/waf-2015-08-24/PutLoggingConfigurationResponse
type PutLoggingConfigurationOutput struct {
	_ struct{} `type:"structure"`

	// The LoggingConfiguration that you submitted in the request.
	LoggingConfiguration *LoggingConfiguration `type:"structure"`
}

// String returns the string representation
func (s PutLoggingConfigurationOutput) String() string {
	return awsutil.Prettify(s)
}

const opPutLoggingConfiguration = "PutLoggingConfiguration"

// PutLoggingConfigurationRequest returns a request value for making API operation for
// AWS WAF.
//
// Associates a LoggingConfiguration with a specified web ACL.
//
// You can access information about all traffic that AWS WAF inspects using
// the following steps:
//
// Create an Amazon Kinesis Data Firehose .
//
// Create the data firehose with a PUT source and in the region that you are
// operating. However, if you are capturing logs for Amazon CloudFront, always
// create the firehose in US East (N. Virginia).
//
// Associate that firehose to your web ACL using a PutLoggingConfiguration request.
//
// When you successfully enable logging using a PutLoggingConfiguration request,
// AWS WAF will create a service linked role with the necessary permissions
// to write logs to the Amazon Kinesis Data Firehose. For more information,
// see Logging Web ACL Traffic Information (https://docs.aws.amazon.com/waf/latest/developerguide/logging.html)
// in the AWS WAF Developer Guide.
//
//    // Example sending a request using PutLoggingConfigurationRequest.
//    req := client.PutLoggingConfigurationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/waf-2015-08-24/PutLoggingConfiguration
func (c *Client) PutLoggingConfigurationRequest(input *PutLoggingConfigurationInput) PutLoggingConfigurationRequest {
	op := &aws.Operation{
		Name:       opPutLoggingConfiguration,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &PutLoggingConfigurationInput{}
	}

	req := c.newRequest(op, input, &PutLoggingConfigurationOutput{})
	return PutLoggingConfigurationRequest{Request: req, Input: input, Copy: c.PutLoggingConfigurationRequest}
}

// PutLoggingConfigurationRequest is the request type for the
// PutLoggingConfiguration API operation.
type PutLoggingConfigurationRequest struct {
	*aws.Request
	Input *PutLoggingConfigurationInput
	Copy  func(*PutLoggingConfigurationInput) PutLoggingConfigurationRequest
}

// Send marshals and sends the PutLoggingConfiguration API request.
func (r PutLoggingConfigurationRequest) Send(ctx context.Context) (*PutLoggingConfigurationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &PutLoggingConfigurationResponse{
		PutLoggingConfigurationOutput: r.Request.Data.(*PutLoggingConfigurationOutput),
		response:                      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// PutLoggingConfigurationResponse is the response type for the
// PutLoggingConfiguration API operation.
type PutLoggingConfigurationResponse struct {
	*PutLoggingConfigurationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// PutLoggingConfiguration request.
func (r *PutLoggingConfigurationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
