// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cognitoidentityprovider

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/cognito-idp-2016-04-18/DeleteGroupRequest
type DeleteGroupInput struct {
	_ struct{} `type:"structure"`

	// The name of the group.
	//
	// GroupName is a required field
	GroupName *string `min:"1" type:"string" required:"true"`

	// The user pool ID for the user pool.
	//
	// UserPoolId is a required field
	UserPoolId *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteGroupInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteGroupInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteGroupInput"}

	if s.GroupName == nil {
		invalidParams.Add(aws.NewErrParamRequired("GroupName"))
	}
	if s.GroupName != nil && len(*s.GroupName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("GroupName", 1))
	}

	if s.UserPoolId == nil {
		invalidParams.Add(aws.NewErrParamRequired("UserPoolId"))
	}
	if s.UserPoolId != nil && len(*s.UserPoolId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("UserPoolId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/cognito-idp-2016-04-18/DeleteGroupOutput
type DeleteGroupOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteGroupOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteGroup = "DeleteGroup"

// DeleteGroupRequest returns a request value for making API operation for
// Amazon Cognito Identity Provider.
//
// Deletes a group. Currently only groups with no members can be deleted.
//
// Requires developer credentials.
//
//    // Example sending a request using DeleteGroupRequest.
//    req := client.DeleteGroupRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cognito-idp-2016-04-18/DeleteGroup
func (c *Client) DeleteGroupRequest(input *DeleteGroupInput) DeleteGroupRequest {
	op := &aws.Operation{
		Name:       opDeleteGroup,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteGroupInput{}
	}

	req := c.newRequest(op, input, &DeleteGroupOutput{})
	req.Handlers.Unmarshal.Remove(jsonrpc.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return DeleteGroupRequest{Request: req, Input: input, Copy: c.DeleteGroupRequest}
}

// DeleteGroupRequest is the request type for the
// DeleteGroup API operation.
type DeleteGroupRequest struct {
	*aws.Request
	Input *DeleteGroupInput
	Copy  func(*DeleteGroupInput) DeleteGroupRequest
}

// Send marshals and sends the DeleteGroup API request.
func (r DeleteGroupRequest) Send(ctx context.Context) (*DeleteGroupResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteGroupResponse{
		DeleteGroupOutput: r.Request.Data.(*DeleteGroupOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteGroupResponse is the response type for the
// DeleteGroup API operation.
type DeleteGroupResponse struct {
	*DeleteGroupOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteGroup request.
func (r *DeleteGroupResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
