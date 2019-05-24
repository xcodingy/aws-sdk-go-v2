// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sagemaker

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/sagemaker-2017-07-24/UpdateWorkteamRequest
type UpdateWorkteamInput struct {
	_ struct{} `type:"structure"`

	// An updated description for the work team.
	Description *string `min:"1" type:"string"`

	// A list of MemberDefinition objects that contain the updated work team members.
	MemberDefinitions []MemberDefinition `min:"1" type:"list"`

	// The name of the work team to update.
	//
	// WorkteamName is a required field
	WorkteamName *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s UpdateWorkteamInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateWorkteamInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateWorkteamInput"}
	if s.Description != nil && len(*s.Description) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Description", 1))
	}
	if s.MemberDefinitions != nil && len(s.MemberDefinitions) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("MemberDefinitions", 1))
	}

	if s.WorkteamName == nil {
		invalidParams.Add(aws.NewErrParamRequired("WorkteamName"))
	}
	if s.WorkteamName != nil && len(*s.WorkteamName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("WorkteamName", 1))
	}
	if s.MemberDefinitions != nil {
		for i, v := range s.MemberDefinitions {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "MemberDefinitions", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/sagemaker-2017-07-24/UpdateWorkteamResponse
type UpdateWorkteamOutput struct {
	_ struct{} `type:"structure"`

	// A Workteam object that describes the updated work team.
	//
	// Workteam is a required field
	Workteam *Workteam `type:"structure" required:"true"`
}

// String returns the string representation
func (s UpdateWorkteamOutput) String() string {
	return awsutil.Prettify(s)
}

const opUpdateWorkteam = "UpdateWorkteam"

// UpdateWorkteamRequest returns a request value for making API operation for
// Amazon SageMaker Service.
//
// Updates an existing work team with new member definitions or description.
//
//    // Example sending a request using UpdateWorkteamRequest.
//    req := client.UpdateWorkteamRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/sagemaker-2017-07-24/UpdateWorkteam
func (c *Client) UpdateWorkteamRequest(input *UpdateWorkteamInput) UpdateWorkteamRequest {
	op := &aws.Operation{
		Name:       opUpdateWorkteam,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateWorkteamInput{}
	}

	req := c.newRequest(op, input, &UpdateWorkteamOutput{})
	return UpdateWorkteamRequest{Request: req, Input: input, Copy: c.UpdateWorkteamRequest}
}

// UpdateWorkteamRequest is the request type for the
// UpdateWorkteam API operation.
type UpdateWorkteamRequest struct {
	*aws.Request
	Input *UpdateWorkteamInput
	Copy  func(*UpdateWorkteamInput) UpdateWorkteamRequest
}

// Send marshals and sends the UpdateWorkteam API request.
func (r UpdateWorkteamRequest) Send(ctx context.Context) (*UpdateWorkteamResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateWorkteamResponse{
		UpdateWorkteamOutput: r.Request.Data.(*UpdateWorkteamOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateWorkteamResponse is the response type for the
// UpdateWorkteam API operation.
type UpdateWorkteamResponse struct {
	*UpdateWorkteamOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateWorkteam request.
func (r *UpdateWorkteamResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
