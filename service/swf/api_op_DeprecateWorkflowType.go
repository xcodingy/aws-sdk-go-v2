// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package swf

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
)

type DeprecateWorkflowTypeInput struct {
	_ struct{} `type:"structure"`

	// The name of the domain in which the workflow type is registered.
	//
	// Domain is a required field
	Domain *string `locationName:"domain" min:"1" type:"string" required:"true"`

	// The workflow type to deprecate.
	//
	// WorkflowType is a required field
	WorkflowType *WorkflowType `locationName:"workflowType" type:"structure" required:"true"`
}

// String returns the string representation
func (s DeprecateWorkflowTypeInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeprecateWorkflowTypeInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeprecateWorkflowTypeInput"}

	if s.Domain == nil {
		invalidParams.Add(aws.NewErrParamRequired("Domain"))
	}
	if s.Domain != nil && len(*s.Domain) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Domain", 1))
	}

	if s.WorkflowType == nil {
		invalidParams.Add(aws.NewErrParamRequired("WorkflowType"))
	}
	if s.WorkflowType != nil {
		if err := s.WorkflowType.Validate(); err != nil {
			invalidParams.AddNested("WorkflowType", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeprecateWorkflowTypeOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeprecateWorkflowTypeOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeprecateWorkflowType = "DeprecateWorkflowType"

// DeprecateWorkflowTypeRequest returns a request value for making API operation for
// Amazon Simple Workflow Service.
//
// Deprecates the specified workflow type. After a workflow type has been deprecated,
// you cannot create new executions of that type. Executions that were started
// before the type was deprecated continues to run. A deprecated workflow type
// may still be used when calling visibility actions.
//
// This operation is eventually consistent. The results are best effort and
// may not exactly reflect recent updates and changes.
//
// Access Control
//
// You can use IAM policies to control this action's access to Amazon SWF resources
// as follows:
//
//    * Use a Resource element with the domain name to limit the action to only
//    specified domains.
//
//    * Use an Action element to allow or deny permission to call this action.
//
//    * Constrain the following parameters by using a Condition element with
//    the appropriate keys.
//
// workflowType.name: String constraint. The key is swf:workflowType.name.
//
// workflowType.version: String constraint. The key is swf:workflowType.version.
//
// If the caller doesn't have sufficient permissions to invoke the action, or
// the parameter values fall outside the specified constraints, the action fails.
// The associated event attribute's cause parameter is set to OPERATION_NOT_PERMITTED.
// For details and example IAM policies, see Using IAM to Manage Access to Amazon
// SWF Workflows (http://docs.aws.amazon.com/amazonswf/latest/developerguide/swf-dev-iam.html)
// in the Amazon SWF Developer Guide.
//
//    // Example sending a request using DeprecateWorkflowTypeRequest.
//    req := client.DeprecateWorkflowTypeRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) DeprecateWorkflowTypeRequest(input *DeprecateWorkflowTypeInput) DeprecateWorkflowTypeRequest {
	op := &aws.Operation{
		Name:       opDeprecateWorkflowType,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeprecateWorkflowTypeInput{}
	}

	req := c.newRequest(op, input, &DeprecateWorkflowTypeOutput{})
	req.Handlers.Unmarshal.Remove(jsonrpc.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return DeprecateWorkflowTypeRequest{Request: req, Input: input, Copy: c.DeprecateWorkflowTypeRequest}
}

// DeprecateWorkflowTypeRequest is the request type for the
// DeprecateWorkflowType API operation.
type DeprecateWorkflowTypeRequest struct {
	*aws.Request
	Input *DeprecateWorkflowTypeInput
	Copy  func(*DeprecateWorkflowTypeInput) DeprecateWorkflowTypeRequest
}

// Send marshals and sends the DeprecateWorkflowType API request.
func (r DeprecateWorkflowTypeRequest) Send(ctx context.Context) (*DeprecateWorkflowTypeResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeprecateWorkflowTypeResponse{
		DeprecateWorkflowTypeOutput: r.Request.Data.(*DeprecateWorkflowTypeOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeprecateWorkflowTypeResponse is the response type for the
// DeprecateWorkflowType API operation.
type DeprecateWorkflowTypeResponse struct {
	*DeprecateWorkflowTypeOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeprecateWorkflowType request.
func (r *DeprecateWorkflowTypeResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
