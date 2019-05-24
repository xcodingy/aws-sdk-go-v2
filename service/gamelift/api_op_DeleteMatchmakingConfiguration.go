// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package gamelift

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents the input for a request action.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/gamelift-2015-10-01/DeleteMatchmakingConfigurationInput
type DeleteMatchmakingConfigurationInput struct {
	_ struct{} `type:"structure"`

	// Unique identifier for a matchmaking configuration
	//
	// Name is a required field
	Name *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteMatchmakingConfigurationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteMatchmakingConfigurationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteMatchmakingConfigurationInput"}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}
	if s.Name != nil && len(*s.Name) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Name", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/gamelift-2015-10-01/DeleteMatchmakingConfigurationOutput
type DeleteMatchmakingConfigurationOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteMatchmakingConfigurationOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteMatchmakingConfiguration = "DeleteMatchmakingConfiguration"

// DeleteMatchmakingConfigurationRequest returns a request value for making API operation for
// Amazon GameLift.
//
// Permanently removes a FlexMatch matchmaking configuration. To delete, specify
// the configuration name. A matchmaking configuration cannot be deleted if
// it is being used in any active matchmaking tickets.
//
//    * CreateMatchmakingConfiguration
//
//    * DescribeMatchmakingConfigurations
//
//    * UpdateMatchmakingConfiguration
//
//    * DeleteMatchmakingConfiguration
//
//    * CreateMatchmakingRuleSet
//
//    * DescribeMatchmakingRuleSets
//
//    * ValidateMatchmakingRuleSet
//
//    * DeleteMatchmakingRuleSet
//
//    // Example sending a request using DeleteMatchmakingConfigurationRequest.
//    req := client.DeleteMatchmakingConfigurationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/gamelift-2015-10-01/DeleteMatchmakingConfiguration
func (c *Client) DeleteMatchmakingConfigurationRequest(input *DeleteMatchmakingConfigurationInput) DeleteMatchmakingConfigurationRequest {
	op := &aws.Operation{
		Name:       opDeleteMatchmakingConfiguration,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteMatchmakingConfigurationInput{}
	}

	req := c.newRequest(op, input, &DeleteMatchmakingConfigurationOutput{})
	return DeleteMatchmakingConfigurationRequest{Request: req, Input: input, Copy: c.DeleteMatchmakingConfigurationRequest}
}

// DeleteMatchmakingConfigurationRequest is the request type for the
// DeleteMatchmakingConfiguration API operation.
type DeleteMatchmakingConfigurationRequest struct {
	*aws.Request
	Input *DeleteMatchmakingConfigurationInput
	Copy  func(*DeleteMatchmakingConfigurationInput) DeleteMatchmakingConfigurationRequest
}

// Send marshals and sends the DeleteMatchmakingConfiguration API request.
func (r DeleteMatchmakingConfigurationRequest) Send(ctx context.Context) (*DeleteMatchmakingConfigurationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteMatchmakingConfigurationResponse{
		DeleteMatchmakingConfigurationOutput: r.Request.Data.(*DeleteMatchmakingConfigurationOutput),
		response:                             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteMatchmakingConfigurationResponse is the response type for the
// DeleteMatchmakingConfiguration API operation.
type DeleteMatchmakingConfigurationResponse struct {
	*DeleteMatchmakingConfigurationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteMatchmakingConfiguration request.
func (r *DeleteMatchmakingConfigurationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
