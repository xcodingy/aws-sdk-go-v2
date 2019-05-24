// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package athena

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/athena-2017-05-18/CreateWorkGroupInput
type CreateWorkGroupInput struct {
	_ struct{} `type:"structure"`

	// The configuration for the workgroup, which includes the location in Amazon
	// S3 where query results are stored, the encryption configuration, if any,
	// used for encrypting query results, whether the Amazon CloudWatch Metrics
	// are enabled for the workgroup, the limit for the amount of bytes scanned
	// (cutoff) per query, if it is specified, and whether workgroup's settings
	// (specified with EnforceWorkGroupConfiguration) in the WorkGroupConfiguration
	// override client-side settings. See WorkGroupConfiguration$EnforceWorkGroupConfiguration.
	Configuration *WorkGroupConfiguration `type:"structure"`

	// The workgroup description.
	Description *string `type:"string"`

	// The workgroup name.
	//
	// Name is a required field
	Name *string `type:"string" required:"true"`

	// One or more tags, separated by commas, that you want to attach to the workgroup
	// as you create it.
	Tags []Tag `type:"list"`
}

// String returns the string representation
func (s CreateWorkGroupInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateWorkGroupInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateWorkGroupInput"}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}
	if s.Configuration != nil {
		if err := s.Configuration.Validate(); err != nil {
			invalidParams.AddNested("Configuration", err.(aws.ErrInvalidParams))
		}
	}
	if s.Tags != nil {
		for i, v := range s.Tags {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Tags", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/athena-2017-05-18/CreateWorkGroupOutput
type CreateWorkGroupOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s CreateWorkGroupOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateWorkGroup = "CreateWorkGroup"

// CreateWorkGroupRequest returns a request value for making API operation for
// Amazon Athena.
//
// Creates a workgroup with the specified name.
//
//    // Example sending a request using CreateWorkGroupRequest.
//    req := client.CreateWorkGroupRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/athena-2017-05-18/CreateWorkGroup
func (c *Client) CreateWorkGroupRequest(input *CreateWorkGroupInput) CreateWorkGroupRequest {
	op := &aws.Operation{
		Name:       opCreateWorkGroup,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateWorkGroupInput{}
	}

	req := c.newRequest(op, input, &CreateWorkGroupOutput{})
	return CreateWorkGroupRequest{Request: req, Input: input, Copy: c.CreateWorkGroupRequest}
}

// CreateWorkGroupRequest is the request type for the
// CreateWorkGroup API operation.
type CreateWorkGroupRequest struct {
	*aws.Request
	Input *CreateWorkGroupInput
	Copy  func(*CreateWorkGroupInput) CreateWorkGroupRequest
}

// Send marshals and sends the CreateWorkGroup API request.
func (r CreateWorkGroupRequest) Send(ctx context.Context) (*CreateWorkGroupResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateWorkGroupResponse{
		CreateWorkGroupOutput: r.Request.Data.(*CreateWorkGroupOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateWorkGroupResponse is the response type for the
// CreateWorkGroup API operation.
type CreateWorkGroupResponse struct {
	*CreateWorkGroupOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateWorkGroup request.
func (r *CreateWorkGroupResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
