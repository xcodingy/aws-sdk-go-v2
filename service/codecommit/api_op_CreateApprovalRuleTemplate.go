// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codecommit

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type CreateApprovalRuleTemplateInput struct {
	_ struct{} `type:"structure"`

	// The content of the approval rule that is created on pull requests in associated
	// repositories. If you specify one or more destination references (branches),
	// approval rules are created in an associated repository only if their destination
	// references (branches) match those specified in the template.
	//
	// When you create the content of the approval rule template, you can specify
	// approvers in an approval pool in one of two ways:
	//
	//    * CodeCommitApprovers: This option only requires an AWS account and a
	//    resource. It can be used for both IAM users and federated access users
	//    whose name matches the provided resource name. This is a very powerful
	//    option that offers a great deal of flexibility. For example, if you specify
	//    the AWS account 123456789012 and Mary_Major, all of the following are
	//    counted as approvals coming from that user: An IAM user in the account
	//    (arn:aws:iam::123456789012:user/Mary_Major) A federated user identified
	//    in IAM as Mary_Major (arn:aws:sts::123456789012:federated-user/Mary_Major)
	//    This option does not recognize an active session of someone assuming the
	//    role of CodeCommitReview with a role session name of Mary_Major (arn:aws:sts::123456789012:assumed-role/CodeCommitReview/Mary_Major)
	//    unless you include a wildcard (*Mary_Major).
	//
	//    * Fully qualified ARN: This option allows you to specify the fully qualified
	//    Amazon Resource Name (ARN) of the IAM user or role.
	//
	// For more information about IAM ARNs, wildcards, and formats, see IAM Identifiers
	// (https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_identifiers.html)
	// in the IAM User Guide.
	//
	// ApprovalRuleTemplateContent is a required field
	ApprovalRuleTemplateContent *string `locationName:"approvalRuleTemplateContent" min:"1" type:"string" required:"true"`

	// The description of the approval rule template. Consider providing a description
	// that explains what this template does and when it might be appropriate to
	// associate it with repositories.
	ApprovalRuleTemplateDescription *string `locationName:"approvalRuleTemplateDescription" type:"string"`

	// The name of the approval rule template. Provide descriptive names, because
	// this name is applied to the approval rules created automatically in associated
	// repositories.
	//
	// ApprovalRuleTemplateName is a required field
	ApprovalRuleTemplateName *string `locationName:"approvalRuleTemplateName" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s CreateApprovalRuleTemplateInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateApprovalRuleTemplateInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateApprovalRuleTemplateInput"}

	if s.ApprovalRuleTemplateContent == nil {
		invalidParams.Add(aws.NewErrParamRequired("ApprovalRuleTemplateContent"))
	}
	if s.ApprovalRuleTemplateContent != nil && len(*s.ApprovalRuleTemplateContent) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ApprovalRuleTemplateContent", 1))
	}

	if s.ApprovalRuleTemplateName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ApprovalRuleTemplateName"))
	}
	if s.ApprovalRuleTemplateName != nil && len(*s.ApprovalRuleTemplateName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ApprovalRuleTemplateName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreateApprovalRuleTemplateOutput struct {
	_ struct{} `type:"structure"`

	// The content and structure of the created approval rule template.
	//
	// ApprovalRuleTemplate is a required field
	ApprovalRuleTemplate *ApprovalRuleTemplate `locationName:"approvalRuleTemplate" type:"structure" required:"true"`
}

// String returns the string representation
func (s CreateApprovalRuleTemplateOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateApprovalRuleTemplate = "CreateApprovalRuleTemplate"

// CreateApprovalRuleTemplateRequest returns a request value for making API operation for
// AWS CodeCommit.
//
// Creates a template for approval rules that can then be associated with one
// or more repositories in your AWS account. When you associate a template with
// a repository, AWS CodeCommit creates an approval rule that matches the conditions
// of the template for all pull requests that meet the conditions of the template.
// For more information, see AssociateApprovalRuleTemplateWithRepository.
//
//    // Example sending a request using CreateApprovalRuleTemplateRequest.
//    req := client.CreateApprovalRuleTemplateRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codecommit-2015-04-13/CreateApprovalRuleTemplate
func (c *Client) CreateApprovalRuleTemplateRequest(input *CreateApprovalRuleTemplateInput) CreateApprovalRuleTemplateRequest {
	op := &aws.Operation{
		Name:       opCreateApprovalRuleTemplate,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateApprovalRuleTemplateInput{}
	}

	req := c.newRequest(op, input, &CreateApprovalRuleTemplateOutput{})

	return CreateApprovalRuleTemplateRequest{Request: req, Input: input, Copy: c.CreateApprovalRuleTemplateRequest}
}

// CreateApprovalRuleTemplateRequest is the request type for the
// CreateApprovalRuleTemplate API operation.
type CreateApprovalRuleTemplateRequest struct {
	*aws.Request
	Input *CreateApprovalRuleTemplateInput
	Copy  func(*CreateApprovalRuleTemplateInput) CreateApprovalRuleTemplateRequest
}

// Send marshals and sends the CreateApprovalRuleTemplate API request.
func (r CreateApprovalRuleTemplateRequest) Send(ctx context.Context) (*CreateApprovalRuleTemplateResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateApprovalRuleTemplateResponse{
		CreateApprovalRuleTemplateOutput: r.Request.Data.(*CreateApprovalRuleTemplateOutput),
		response:                         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateApprovalRuleTemplateResponse is the response type for the
// CreateApprovalRuleTemplate API operation.
type CreateApprovalRuleTemplateResponse struct {
	*CreateApprovalRuleTemplateOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateApprovalRuleTemplate request.
func (r *CreateApprovalRuleTemplateResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}