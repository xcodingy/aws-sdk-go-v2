// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudformation

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/cloudformation-2010-05-15/CreateStackSetInput
type CreateStackSetInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Number (ARN) of the IAM role to use to create this stack
	// set.
	//
	// Specify an IAM role only if you are using customized administrator roles
	// to control which users or groups can manage specific stack sets within the
	// same administrator account. For more information, see Prerequisites: Granting
	// Permissions for Stack Set Operations (http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/stacksets-prereqs.html)
	// in the AWS CloudFormation User Guide.
	AdministrationRoleARN *string `min:"20" type:"string"`

	// In some cases, you must explicity acknowledge that your stack set template
	// contains certain capabilities in order for AWS CloudFormation to create the
	// stack set and related stack instances.
	//
	//    * CAPABILITY_IAM and CAPABILITY_NAMED_IAM
	//
	// Some stack templates might include resources that can affect permissions
	//    in your AWS account; for example, by creating new AWS Identity and Access
	//    Management (IAM) users. For those stack sets, you must explicitly acknowledge
	//    this by specifying one of these capabilities.
	//
	// The following IAM resources require you to specify either the CAPABILITY_IAM
	//    or CAPABILITY_NAMED_IAM capability.
	//
	// If you have IAM resources, you can specify either capability.
	//
	// If you have IAM resources with custom names, you must specify CAPABILITY_NAMED_IAM.
	//
	//
	// If you don't specify either of these capabilities, AWS CloudFormation returns
	//    an InsufficientCapabilities error.
	//
	// If your stack template contains these resources, we recommend that you review
	//    all permissions associated with them and edit their permissions if necessary.
	//
	//  AWS::IAM::AccessKey (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-accesskey.html)
	//
	//  AWS::IAM::Group (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-group.html)
	//
	//  AWS::IAM::InstanceProfile (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-instanceprofile.html)
	//
	//  AWS::IAM::Policy (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-policy.html)
	//
	//  AWS::IAM::Role (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-role.html)
	//
	//  AWS::IAM::User (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-user.html)
	//
	//  AWS::IAM::UserToGroupAddition (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-addusertogroup.html)
	//
	// For more information, see Acknowledging IAM Resources in AWS CloudFormation
	//    Templates (http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-iam-template.html#capabilities).
	//
	//    * CAPABILITY_AUTO_EXPAND
	//
	// Some templates contain macros. If your stack template contains one or more
	//    macros, and you choose to create a stack directly from the processed template,
	//    without first reviewing the resulting changes in a change set, you must
	//    acknowledge this capability. For more information, see Using AWS CloudFormation
	//    Macros to Perform Custom Processing on Templates (http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/template-macros.html).
	//
	// Stack sets do not currently support macros in stack templates. (This includes
	//    the AWS::Include (http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/create-reusable-transform-function-snippets-and-add-to-your-template-with-aws-include-transform.html)
	//    and AWS::Serverless (http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/transform-aws-serverless.html)
	//    transforms, which are macros hosted by AWS CloudFormation.) Even if you
	//    specify this capability, if you include a macro in your template the stack
	//    set operation will fail.
	Capabilities []Capability `type:"list"`

	// A unique identifier for this CreateStackSet request. Specify this token if
	// you plan to retry requests so that AWS CloudFormation knows that you're not
	// attempting to create another stack set with the same name. You might retry
	// CreateStackSet requests to ensure that AWS CloudFormation successfully received
	// them.
	//
	// If you don't specify an operation ID, the SDK generates one automatically.
	ClientRequestToken *string `min:"1" type:"string" idempotencyToken:"true"`

	// A description of the stack set. You can use the description to identify the
	// stack set's purpose or other important information.
	Description *string `min:"1" type:"string"`

	// The name of the IAM execution role to use to create the stack set. If you
	// do not specify an execution role, AWS CloudFormation uses the AWSCloudFormationStackSetExecutionRole
	// role for the stack set operation.
	//
	// Specify an IAM role only if you are using customized execution roles to control
	// which stack resources users and groups can include in their stack sets.
	ExecutionRoleName *string `min:"1" type:"string"`

	// The input parameters for the stack set template.
	Parameters []Parameter `type:"list"`

	// The name to associate with the stack set. The name must be unique in the
	// region where you create your stack set.
	//
	// A stack name can contain only alphanumeric characters (case-sensitive) and
	// hyphens. It must start with an alphabetic character and can't be longer than
	// 128 characters.
	//
	// StackSetName is a required field
	StackSetName *string `type:"string" required:"true"`

	// The key-value pairs to associate with this stack set and the stacks created
	// from it. AWS CloudFormation also propagates these tags to supported resources
	// that are created in the stacks. A maximum number of 50 tags can be specified.
	//
	// If you specify tags as part of a CreateStackSet action, AWS CloudFormation
	// checks to see if you have the required IAM permission to tag resources. If
	// you don't, the entire CreateStackSet action fails with an access denied error,
	// and the stack set is not created.
	Tags []Tag `type:"list"`

	// The structure that contains the template body, with a minimum length of 1
	// byte and a maximum length of 51,200 bytes. For more information, see Template
	// Anatomy (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/template-anatomy.html)
	// in the AWS CloudFormation User Guide.
	//
	// Conditional: You must specify either the TemplateBody or the TemplateURL
	// parameter, but not both.
	TemplateBody *string `min:"1" type:"string"`

	// The location of the file that contains the template body. The URL must point
	// to a template (maximum size: 460,800 bytes) that's located in an Amazon S3
	// bucket. For more information, see Template Anatomy (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/template-anatomy.html)
	// in the AWS CloudFormation User Guide.
	//
	// Conditional: You must specify either the TemplateBody or the TemplateURL
	// parameter, but not both.
	TemplateURL *string `min:"1" type:"string"`
}

// String returns the string representation
func (s CreateStackSetInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateStackSetInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateStackSetInput"}
	if s.AdministrationRoleARN != nil && len(*s.AdministrationRoleARN) < 20 {
		invalidParams.Add(aws.NewErrParamMinLen("AdministrationRoleARN", 20))
	}
	if s.ClientRequestToken != nil && len(*s.ClientRequestToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ClientRequestToken", 1))
	}
	if s.Description != nil && len(*s.Description) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Description", 1))
	}
	if s.ExecutionRoleName != nil && len(*s.ExecutionRoleName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ExecutionRoleName", 1))
	}

	if s.StackSetName == nil {
		invalidParams.Add(aws.NewErrParamRequired("StackSetName"))
	}
	if s.TemplateBody != nil && len(*s.TemplateBody) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("TemplateBody", 1))
	}
	if s.TemplateURL != nil && len(*s.TemplateURL) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("TemplateURL", 1))
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

// Please also see https://docs.aws.amazon.com/goto/WebAPI/cloudformation-2010-05-15/CreateStackSetOutput
type CreateStackSetOutput struct {
	_ struct{} `type:"structure"`

	// The ID of the stack set that you're creating.
	StackSetId *string `type:"string"`
}

// String returns the string representation
func (s CreateStackSetOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateStackSet = "CreateStackSet"

// CreateStackSetRequest returns a request value for making API operation for
// AWS CloudFormation.
//
// Creates a stack set.
//
//    // Example sending a request using CreateStackSetRequest.
//    req := client.CreateStackSetRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cloudformation-2010-05-15/CreateStackSet
func (c *Client) CreateStackSetRequest(input *CreateStackSetInput) CreateStackSetRequest {
	op := &aws.Operation{
		Name:       opCreateStackSet,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateStackSetInput{}
	}

	req := c.newRequest(op, input, &CreateStackSetOutput{})
	return CreateStackSetRequest{Request: req, Input: input, Copy: c.CreateStackSetRequest}
}

// CreateStackSetRequest is the request type for the
// CreateStackSet API operation.
type CreateStackSetRequest struct {
	*aws.Request
	Input *CreateStackSetInput
	Copy  func(*CreateStackSetInput) CreateStackSetRequest
}

// Send marshals and sends the CreateStackSet API request.
func (r CreateStackSetRequest) Send(ctx context.Context) (*CreateStackSetResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateStackSetResponse{
		CreateStackSetOutput: r.Request.Data.(*CreateStackSetOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateStackSetResponse is the response type for the
// CreateStackSet API operation.
type CreateStackSetResponse struct {
	*CreateStackSetOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateStackSet request.
func (r *CreateStackSetResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
