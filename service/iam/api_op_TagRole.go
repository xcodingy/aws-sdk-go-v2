// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iam

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/query"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/iam-2010-05-08/TagRoleRequest
type TagRoleInput struct {
	_ struct{} `type:"structure"`

	// The name of the role that you want to add tags to.
	//
	// This parameter accepts (through its regex pattern (http://wikipedia.org/wiki/regex))
	// a string of characters that consist of upper and lowercase alphanumeric characters
	// with no spaces. You can also include any of the following characters: _+=,.@-
	//
	// RoleName is a required field
	RoleName *string `min:"1" type:"string" required:"true"`

	// The list of tags that you want to attach to the role. Each tag consists of
	// a key name and an associated value. You can specify this with a JSON string.
	//
	// Tags is a required field
	Tags []Tag `type:"list" required:"true"`
}

// String returns the string representation
func (s TagRoleInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *TagRoleInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "TagRoleInput"}

	if s.RoleName == nil {
		invalidParams.Add(aws.NewErrParamRequired("RoleName"))
	}
	if s.RoleName != nil && len(*s.RoleName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("RoleName", 1))
	}

	if s.Tags == nil {
		invalidParams.Add(aws.NewErrParamRequired("Tags"))
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

// Please also see https://docs.aws.amazon.com/goto/WebAPI/iam-2010-05-08/TagRoleOutput
type TagRoleOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s TagRoleOutput) String() string {
	return awsutil.Prettify(s)
}

const opTagRole = "TagRole"

// TagRoleRequest returns a request value for making API operation for
// AWS Identity and Access Management.
//
// Adds one or more tags to an IAM role. The role can be a regular role or a
// service-linked role. If a tag with the same key name already exists, then
// that tag is overwritten with the new value.
//
// A tag consists of a key name and an associated value. By assigning tags to
// your resources, you can do the following:
//
//    * Administrative grouping and discovery - Attach tags to resources to
//    aid in organization and search. For example, you could search for all
//    resources with the key name Project and the value MyImportantProject.
//    Or search for all resources with the key name Cost Center and the value
//    41200.
//
//    * Access control - Reference tags in IAM user-based and resource-based
//    policies. You can use tags to restrict access to only an IAM user or role
//    that has a specified tag attached. You can also restrict access to only
//    those resources that have a certain tag attached. For examples of policies
//    that show how to use tags to control access, see Control Access Using
//    IAM Tags (https://docs.aws.amazon.com/IAM/latest/UserGuide/access_tags.html)
//    in the IAM User Guide.
//
//    * Cost allocation - Use tags to help track which individuals and teams
//    are using which AWS resources.
//
// Make sure that you have no invalid tags and that you do not exceed the allowed
// number of tags per role. In either case, the entire request fails and no
// tags are added to the role.
//
// AWS always interprets the tag Value as a single string. If you need to store
// an array, you can store comma-separated values in the string. However, you
// must interpret the value in your code.
//
// For more information about tagging, see Tagging IAM Identities (https://docs.aws.amazon.com/IAM/latest/UserGuide/id_tags.html)
// in the IAM User Guide.
//
//    // Example sending a request using TagRoleRequest.
//    req := client.TagRoleRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iam-2010-05-08/TagRole
func (c *Client) TagRoleRequest(input *TagRoleInput) TagRoleRequest {
	op := &aws.Operation{
		Name:       opTagRole,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &TagRoleInput{}
	}

	req := c.newRequest(op, input, &TagRoleOutput{})
	req.Handlers.Unmarshal.Remove(query.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return TagRoleRequest{Request: req, Input: input, Copy: c.TagRoleRequest}
}

// TagRoleRequest is the request type for the
// TagRole API operation.
type TagRoleRequest struct {
	*aws.Request
	Input *TagRoleInput
	Copy  func(*TagRoleInput) TagRoleRequest
}

// Send marshals and sends the TagRole API request.
func (r TagRoleRequest) Send(ctx context.Context) (*TagRoleResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &TagRoleResponse{
		TagRoleOutput: r.Request.Data.(*TagRoleOutput),
		response:      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// TagRoleResponse is the response type for the
// TagRole API operation.
type TagRoleResponse struct {
	*TagRoleOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// TagRole request.
func (r *TagRoleResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
