// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package opsworks

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/opsworks-2013-02-18/UpdateAppRequest
type UpdateAppInput struct {
	_ struct{} `type:"structure"`

	// The app ID.
	//
	// AppId is a required field
	AppId *string `type:"string" required:"true"`

	// A Source object that specifies the app repository.
	AppSource *Source `type:"structure"`

	// One or more user-defined key/value pairs to be added to the stack attributes.
	Attributes map[string]string `type:"map"`

	// The app's data sources.
	DataSources []DataSource `type:"list"`

	// A description of the app.
	Description *string `type:"string"`

	// The app's virtual host settings, with multiple domains separated by commas.
	// For example: 'www.example.com, example.com'
	Domains []string `type:"list"`

	// Whether SSL is enabled for the app.
	EnableSsl *bool `type:"boolean"`

	// An array of EnvironmentVariable objects that specify environment variables
	// to be associated with the app. After you deploy the app, these variables
	// are defined on the associated app server instances.For more information,
	// see  Environment Variables (http://docs.aws.amazon.com/opsworks/latest/userguide/workingapps-creating.html#workingapps-creating-environment).
	//
	// There is no specific limit on the number of environment variables. However,
	// the size of the associated data structure - which includes the variables'
	// names, values, and protected flag values - cannot exceed 10 KB (10240 Bytes).
	// This limit should accommodate most if not all use cases. Exceeding it will
	// cause an exception with the message, "Environment: is too large (maximum
	// is 10KB)."
	//
	// This parameter is supported only by Chef 11.10 stacks. If you have specified
	// one or more environment variables, you cannot modify the stack's Chef version.
	Environment []EnvironmentVariable `type:"list"`

	// The app name.
	Name *string `type:"string"`

	// An SslConfiguration object with the SSL configuration.
	SslConfiguration *SslConfiguration `type:"structure"`

	// The app type.
	Type AppType `type:"string" enum:"true"`
}

// String returns the string representation
func (s UpdateAppInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateAppInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateAppInput"}

	if s.AppId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AppId"))
	}
	if s.Environment != nil {
		for i, v := range s.Environment {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Environment", i), err.(aws.ErrInvalidParams))
			}
		}
	}
	if s.SslConfiguration != nil {
		if err := s.SslConfiguration.Validate(); err != nil {
			invalidParams.AddNested("SslConfiguration", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/opsworks-2013-02-18/UpdateAppOutput
type UpdateAppOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s UpdateAppOutput) String() string {
	return awsutil.Prettify(s)
}

const opUpdateApp = "UpdateApp"

// UpdateAppRequest returns a request value for making API operation for
// AWS OpsWorks.
//
// Updates a specified app.
//
// Required Permissions: To use this action, an IAM user must have a Deploy
// or Manage permissions level for the stack, or an attached policy that explicitly
// grants permissions. For more information on user permissions, see Managing
// User Permissions (http://docs.aws.amazon.com/opsworks/latest/userguide/opsworks-security-users.html).
//
//    // Example sending a request using UpdateAppRequest.
//    req := client.UpdateAppRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/opsworks-2013-02-18/UpdateApp
func (c *Client) UpdateAppRequest(input *UpdateAppInput) UpdateAppRequest {
	op := &aws.Operation{
		Name:       opUpdateApp,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateAppInput{}
	}

	req := c.newRequest(op, input, &UpdateAppOutput{})
	req.Handlers.Unmarshal.Remove(jsonrpc.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return UpdateAppRequest{Request: req, Input: input, Copy: c.UpdateAppRequest}
}

// UpdateAppRequest is the request type for the
// UpdateApp API operation.
type UpdateAppRequest struct {
	*aws.Request
	Input *UpdateAppInput
	Copy  func(*UpdateAppInput) UpdateAppRequest
}

// Send marshals and sends the UpdateApp API request.
func (r UpdateAppRequest) Send(ctx context.Context) (*UpdateAppResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateAppResponse{
		UpdateAppOutput: r.Request.Data.(*UpdateAppOutput),
		response:        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateAppResponse is the response type for the
// UpdateApp API operation.
type UpdateAppResponse struct {
	*UpdateAppOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateApp request.
func (r *UpdateAppResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
