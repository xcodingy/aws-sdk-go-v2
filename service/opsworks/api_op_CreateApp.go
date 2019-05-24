// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package opsworks

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/opsworks-2013-02-18/CreateAppRequest
type CreateAppInput struct {
	_ struct{} `type:"structure"`

	// A Source object that specifies the app repository.
	AppSource *Source `type:"structure"`

	// One or more user-defined key/value pairs to be added to the stack attributes.
	Attributes map[string]string `type:"map"`

	// The app's data source.
	DataSources []DataSource `type:"list"`

	// A description of the app.
	Description *string `type:"string"`

	// The app virtual host settings, with multiple domains separated by commas.
	// For example: 'www.example.com, example.com'
	Domains []string `type:"list"`

	// Whether to enable SSL for the app.
	EnableSsl *bool `type:"boolean"`

	// An array of EnvironmentVariable objects that specify environment variables
	// to be associated with the app. After you deploy the app, these variables
	// are defined on the associated app server instance. For more information,
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
	//
	// Name is a required field
	Name *string `type:"string" required:"true"`

	// The app's short name.
	Shortname *string `type:"string"`

	// An SslConfiguration object with the SSL configuration.
	SslConfiguration *SslConfiguration `type:"structure"`

	// The stack ID.
	//
	// StackId is a required field
	StackId *string `type:"string" required:"true"`

	// The app type. Each supported type is associated with a particular layer.
	// For example, PHP applications are associated with a PHP layer. AWS OpsWorks
	// Stacks deploys an application to those instances that are members of the
	// corresponding layer. If your app isn't one of the standard types, or you
	// prefer to implement your own Deploy recipes, specify other.
	//
	// Type is a required field
	Type AppType `type:"string" required:"true" enum:"true"`
}

// String returns the string representation
func (s CreateAppInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateAppInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateAppInput"}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}

	if s.StackId == nil {
		invalidParams.Add(aws.NewErrParamRequired("StackId"))
	}
	if len(s.Type) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("Type"))
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

// Contains the response to a CreateApp request.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/opsworks-2013-02-18/CreateAppResult
type CreateAppOutput struct {
	_ struct{} `type:"structure"`

	// The app ID.
	AppId *string `type:"string"`
}

// String returns the string representation
func (s CreateAppOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateApp = "CreateApp"

// CreateAppRequest returns a request value for making API operation for
// AWS OpsWorks.
//
// Creates an app for a specified stack. For more information, see Creating
// Apps (http://docs.aws.amazon.com/opsworks/latest/userguide/workingapps-creating.html).
//
// Required Permissions: To use this action, an IAM user must have a Manage
// permissions level for the stack, or an attached policy that explicitly grants
// permissions. For more information on user permissions, see Managing User
// Permissions (http://docs.aws.amazon.com/opsworks/latest/userguide/opsworks-security-users.html).
//
//    // Example sending a request using CreateAppRequest.
//    req := client.CreateAppRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/opsworks-2013-02-18/CreateApp
func (c *Client) CreateAppRequest(input *CreateAppInput) CreateAppRequest {
	op := &aws.Operation{
		Name:       opCreateApp,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateAppInput{}
	}

	req := c.newRequest(op, input, &CreateAppOutput{})
	return CreateAppRequest{Request: req, Input: input, Copy: c.CreateAppRequest}
}

// CreateAppRequest is the request type for the
// CreateApp API operation.
type CreateAppRequest struct {
	*aws.Request
	Input *CreateAppInput
	Copy  func(*CreateAppInput) CreateAppRequest
}

// Send marshals and sends the CreateApp API request.
func (r CreateAppRequest) Send(ctx context.Context) (*CreateAppResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateAppResponse{
		CreateAppOutput: r.Request.Data.(*CreateAppOutput),
		response:        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateAppResponse is the response type for the
// CreateApp API operation.
type CreateAppResponse struct {
	*CreateAppOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateApp request.
func (r *CreateAppResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
