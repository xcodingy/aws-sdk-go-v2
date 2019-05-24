// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package connect

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/connect-2017-08-08/DescribeUserRequest
type DescribeUserInput struct {
	_ struct{} `type:"structure"`

	// The identifier for your Amazon Connect instance. To find the ID of your instance,
	// open the AWS console and select Amazon Connect. Select the alias of the instance
	// in the Instance alias column. The instance ID is displayed in the Overview
	// section of your instance settings. For example, the instance ID is the set
	// of characters at the end of the instance ARN, after instance/, such as 10a4c4eb-f57e-4d4c-b602-bf39176ced07.
	//
	// InstanceId is a required field
	InstanceId *string `location:"uri" locationName:"InstanceId" min:"1" type:"string" required:"true"`

	// Unique identifier for the user account to return.
	//
	// UserId is a required field
	UserId *string `location:"uri" locationName:"UserId" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeUserInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeUserInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeUserInput"}

	if s.InstanceId == nil {
		invalidParams.Add(aws.NewErrParamRequired("InstanceId"))
	}
	if s.InstanceId != nil && len(*s.InstanceId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("InstanceId", 1))
	}

	if s.UserId == nil {
		invalidParams.Add(aws.NewErrParamRequired("UserId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeUserInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	if s.InstanceId != nil {
		v := *s.InstanceId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "InstanceId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.UserId != nil {
		v := *s.UserId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "UserId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/connect-2017-08-08/DescribeUserResponse
type DescribeUserOutput struct {
	_ struct{} `type:"structure"`

	// A User object that contains information about the user account and configuration
	// settings.
	User *User `type:"structure"`
}

// String returns the string representation
func (s DescribeUserOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeUserOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.User != nil {
		v := s.User

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "User", v, metadata)
	}
	return nil
}

const opDescribeUser = "DescribeUser"

// DescribeUserRequest returns a request value for making API operation for
// Amazon Connect Service.
//
// Returns a User object that contains information about the user account specified
// by the UserId.
//
//    // Example sending a request using DescribeUserRequest.
//    req := client.DescribeUserRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/connect-2017-08-08/DescribeUser
func (c *Client) DescribeUserRequest(input *DescribeUserInput) DescribeUserRequest {
	op := &aws.Operation{
		Name:       opDescribeUser,
		HTTPMethod: "GET",
		HTTPPath:   "/users/{InstanceId}/{UserId}",
	}

	if input == nil {
		input = &DescribeUserInput{}
	}

	req := c.newRequest(op, input, &DescribeUserOutput{})
	return DescribeUserRequest{Request: req, Input: input, Copy: c.DescribeUserRequest}
}

// DescribeUserRequest is the request type for the
// DescribeUser API operation.
type DescribeUserRequest struct {
	*aws.Request
	Input *DescribeUserInput
	Copy  func(*DescribeUserInput) DescribeUserRequest
}

// Send marshals and sends the DescribeUser API request.
func (r DescribeUserRequest) Send(ctx context.Context) (*DescribeUserResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeUserResponse{
		DescribeUserOutput: r.Request.Data.(*DescribeUserOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeUserResponse is the response type for the
// DescribeUser API operation.
type DescribeUserResponse struct {
	*DescribeUserOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeUser request.
func (r *DescribeUserResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
