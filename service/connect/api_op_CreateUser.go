// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package connect

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/connect-2017-08-08/CreateUserRequest
type CreateUserInput struct {
	_ struct{} `type:"structure"`

	// The unique identifier for the user account in the directory service directory
	// used for identity management. If Amazon Connect is unable to access the existing
	// directory, you can use the DirectoryUserId to authenticate users. If you
	// include the parameter, it is assumed that Amazon Connect cannot access the
	// directory. If the parameter is not included, the UserIdentityInfo is used
	// to authenticate users from your existing directory.
	//
	// This parameter is required if you are using an existing directory for identity
	// management in Amazon Connect when Amazon Connect cannot access your directory
	// to authenticate users. If you are using SAML for identity management and
	// include this parameter, an InvalidRequestException is returned.
	DirectoryUserId *string `type:"string"`

	// The unique identifier for the hierarchy group to assign to the user created.
	HierarchyGroupId *string `type:"string"`

	// Information about the user, including email address, first name, and last
	// name.
	IdentityInfo *UserIdentityInfo `type:"structure"`

	// The identifier for your Amazon Connect instance. To find the ID of your instance,
	// open the AWS console and select Amazon Connect. Select the alias of the instance
	// in the Instance alias column. The instance ID is displayed in the Overview
	// section of your instance settings. For example, the instance ID is the set
	// of characters at the end of the instance ARN, after instance/, such as 10a4c4eb-f57e-4d4c-b602-bf39176ced07.
	//
	// InstanceId is a required field
	InstanceId *string `location:"uri" locationName:"InstanceId" min:"1" type:"string" required:"true"`

	// The password for the user account to create. This is required if you are
	// using Amazon Connect for identity management. If you are using SAML for identity
	// management and include this parameter, an InvalidRequestException is returned.
	Password *string `type:"string"`

	// Specifies the phone settings for the user, including AfterContactWorkTimeLimit,
	// AutoAccept, DeskPhoneNumber, and PhoneType.
	//
	// PhoneConfig is a required field
	PhoneConfig *UserPhoneConfig `type:"structure" required:"true"`

	// The unique identifier for the routing profile to assign to the user created.
	//
	// RoutingProfileId is a required field
	RoutingProfileId *string `type:"string" required:"true"`

	// The unique identifier of the security profile to assign to the user created.
	//
	// SecurityProfileIds is a required field
	SecurityProfileIds []string `min:"1" type:"list" required:"true"`

	// The user name in Amazon Connect for the account to create. If you are using
	// SAML for identity management in your Amazon Connect, the value for Username
	// can include up to 64 characters from [a-zA-Z0-9_-.\@]+.
	//
	// Username is a required field
	Username *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s CreateUserInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateUserInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateUserInput"}

	if s.InstanceId == nil {
		invalidParams.Add(aws.NewErrParamRequired("InstanceId"))
	}
	if s.InstanceId != nil && len(*s.InstanceId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("InstanceId", 1))
	}

	if s.PhoneConfig == nil {
		invalidParams.Add(aws.NewErrParamRequired("PhoneConfig"))
	}

	if s.RoutingProfileId == nil {
		invalidParams.Add(aws.NewErrParamRequired("RoutingProfileId"))
	}

	if s.SecurityProfileIds == nil {
		invalidParams.Add(aws.NewErrParamRequired("SecurityProfileIds"))
	}
	if s.SecurityProfileIds != nil && len(s.SecurityProfileIds) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("SecurityProfileIds", 1))
	}

	if s.Username == nil {
		invalidParams.Add(aws.NewErrParamRequired("Username"))
	}
	if s.Username != nil && len(*s.Username) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Username", 1))
	}
	if s.IdentityInfo != nil {
		if err := s.IdentityInfo.Validate(); err != nil {
			invalidParams.AddNested("IdentityInfo", err.(aws.ErrInvalidParams))
		}
	}
	if s.PhoneConfig != nil {
		if err := s.PhoneConfig.Validate(); err != nil {
			invalidParams.AddNested("PhoneConfig", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateUserInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	if s.DirectoryUserId != nil {
		v := *s.DirectoryUserId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "DirectoryUserId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.HierarchyGroupId != nil {
		v := *s.HierarchyGroupId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "HierarchyGroupId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.IdentityInfo != nil {
		v := s.IdentityInfo

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "IdentityInfo", v, metadata)
	}
	if s.Password != nil {
		v := *s.Password

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Password", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.PhoneConfig != nil {
		v := s.PhoneConfig

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "PhoneConfig", v, metadata)
	}
	if s.RoutingProfileId != nil {
		v := *s.RoutingProfileId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "RoutingProfileId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.SecurityProfileIds) > 0 {
		v := s.SecurityProfileIds

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "SecurityProfileIds", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	if s.Username != nil {
		v := *s.Username

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Username", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.InstanceId != nil {
		v := *s.InstanceId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "InstanceId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/connect-2017-08-08/CreateUserResponse
type CreateUserOutput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the user account created.
	UserArn *string `type:"string"`

	// The unique identifier for the user account in Amazon Connect
	UserId *string `type:"string"`
}

// String returns the string representation
func (s CreateUserOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateUserOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.UserArn != nil {
		v := *s.UserArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "UserArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.UserId != nil {
		v := *s.UserId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "UserId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opCreateUser = "CreateUser"

// CreateUserRequest returns a request value for making API operation for
// Amazon Connect Service.
//
// Creates a new user account in your Amazon Connect instance.
//
//    // Example sending a request using CreateUserRequest.
//    req := client.CreateUserRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/connect-2017-08-08/CreateUser
func (c *Client) CreateUserRequest(input *CreateUserInput) CreateUserRequest {
	op := &aws.Operation{
		Name:       opCreateUser,
		HTTPMethod: "PUT",
		HTTPPath:   "/users/{InstanceId}",
	}

	if input == nil {
		input = &CreateUserInput{}
	}

	req := c.newRequest(op, input, &CreateUserOutput{})
	return CreateUserRequest{Request: req, Input: input, Copy: c.CreateUserRequest}
}

// CreateUserRequest is the request type for the
// CreateUser API operation.
type CreateUserRequest struct {
	*aws.Request
	Input *CreateUserInput
	Copy  func(*CreateUserInput) CreateUserRequest
}

// Send marshals and sends the CreateUser API request.
func (r CreateUserRequest) Send(ctx context.Context) (*CreateUserResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateUserResponse{
		CreateUserOutput: r.Request.Data.(*CreateUserOutput),
		response:         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateUserResponse is the response type for the
// CreateUser API operation.
type CreateUserResponse struct {
	*CreateUserOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateUser request.
func (r *CreateUserResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
