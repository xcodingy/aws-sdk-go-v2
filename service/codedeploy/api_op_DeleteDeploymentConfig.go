// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codedeploy

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
)

// Represents the input of a DeleteDeploymentConfig operation.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codedeploy-2014-10-06/DeleteDeploymentConfigInput
type DeleteDeploymentConfigInput struct {
	_ struct{} `type:"structure"`

	// The name of a deployment configuration associated with the IAM user or AWS
	// account.
	//
	// DeploymentConfigName is a required field
	DeploymentConfigName *string `locationName:"deploymentConfigName" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteDeploymentConfigInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteDeploymentConfigInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteDeploymentConfigInput"}

	if s.DeploymentConfigName == nil {
		invalidParams.Add(aws.NewErrParamRequired("DeploymentConfigName"))
	}
	if s.DeploymentConfigName != nil && len(*s.DeploymentConfigName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("DeploymentConfigName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/codedeploy-2014-10-06/DeleteDeploymentConfigOutput
type DeleteDeploymentConfigOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteDeploymentConfigOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteDeploymentConfig = "DeleteDeploymentConfig"

// DeleteDeploymentConfigRequest returns a request value for making API operation for
// AWS CodeDeploy.
//
// Deletes a deployment configuration.
//
// A deployment configuration cannot be deleted if it is currently in use. Predefined
// configurations cannot be deleted.
//
//    // Example sending a request using DeleteDeploymentConfigRequest.
//    req := client.DeleteDeploymentConfigRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codedeploy-2014-10-06/DeleteDeploymentConfig
func (c *Client) DeleteDeploymentConfigRequest(input *DeleteDeploymentConfigInput) DeleteDeploymentConfigRequest {
	op := &aws.Operation{
		Name:       opDeleteDeploymentConfig,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteDeploymentConfigInput{}
	}

	req := c.newRequest(op, input, &DeleteDeploymentConfigOutput{})
	req.Handlers.Unmarshal.Remove(jsonrpc.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return DeleteDeploymentConfigRequest{Request: req, Input: input, Copy: c.DeleteDeploymentConfigRequest}
}

// DeleteDeploymentConfigRequest is the request type for the
// DeleteDeploymentConfig API operation.
type DeleteDeploymentConfigRequest struct {
	*aws.Request
	Input *DeleteDeploymentConfigInput
	Copy  func(*DeleteDeploymentConfigInput) DeleteDeploymentConfigRequest
}

// Send marshals and sends the DeleteDeploymentConfig API request.
func (r DeleteDeploymentConfigRequest) Send(ctx context.Context) (*DeleteDeploymentConfigResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteDeploymentConfigResponse{
		DeleteDeploymentConfigOutput: r.Request.Data.(*DeleteDeploymentConfigOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteDeploymentConfigResponse is the response type for the
// DeleteDeploymentConfig API operation.
type DeleteDeploymentConfigResponse struct {
	*DeleteDeploymentConfigOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteDeploymentConfig request.
func (r *DeleteDeploymentConfigResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
