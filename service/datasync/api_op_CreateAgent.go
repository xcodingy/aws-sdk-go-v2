// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package datasync

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// CreateAgentRequest
// Please also see https://docs.aws.amazon.com/goto/WebAPI/datasync-2018-11-09/CreateAgentRequest
type CreateAgentInput struct {
	_ struct{} `type:"structure"`

	// Your agent activation key. You can get the activation key either by sending
	// an HTTP GET request with redirects that enable you to get the agent IP address
	// (port 80). Alternatively, you can get it from the AWS DataSync console.
	//
	// The redirect URL returned in the response provides you the activation key
	// for your agent in the query string parameter activationKey. It might also
	// include other activation-related parameters; however, these are merely defaults.
	// The arguments you pass to this API call determine the actual configuration
	// of your agent. For more information, see Activating a Sync Agent (https://docs.aws.amazon.com/sync-service/latest/userguide/working-with-sync-agents.html#activating-sync-agent)
	// in the AWS DataSync User Guide.
	//
	// ActivationKey is a required field
	ActivationKey *string `type:"string" required:"true"`

	// The name you configured for your agent. This value is a text reference that
	// is used to identify the agent in the console.
	AgentName *string `min:"1" type:"string"`

	// The key-value pair that represents the tag you want to associate with the
	// agent. The value can be an empty string. This value helps you manage, filter,
	// and search for your agents.
	//
	// Valid characters for key and value are letters, spaces, and numbers representable
	// in UTF-8 format, and the following special characters: + - = . _ : / @.
	Tags []TagListEntry `type:"list"`
}

// String returns the string representation
func (s CreateAgentInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateAgentInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateAgentInput"}

	if s.ActivationKey == nil {
		invalidParams.Add(aws.NewErrParamRequired("ActivationKey"))
	}
	if s.AgentName != nil && len(*s.AgentName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("AgentName", 1))
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

// CreateAgentResponse
// Please also see https://docs.aws.amazon.com/goto/WebAPI/datasync-2018-11-09/CreateAgentResponse
type CreateAgentOutput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the agent. Use the ListAgents operation
	// to return a list of agents for your account and AWS Region.
	AgentArn *string `type:"string"`
}

// String returns the string representation
func (s CreateAgentOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateAgent = "CreateAgent"

// CreateAgentRequest returns a request value for making API operation for
// AWS DataSync.
//
// Activates an AWS DataSync agent that you have deployed on your host. The
// activation process associates your agent with your account. In the activation
// process, you specify information such as the AWS Region that you want to
// activate the agent in. You activate the agent in the AWS Region where your
// target locations (in Amazon S3 or Amazon EFS) reside. Your tasks are created
// in this AWS Region.
//
// You can use an agent for more than one location. If a task uses multiple
// agents, all of them need to have status AVAILABLE for the task to run. If
// you use multiple agents for a source location, the status of all the agents
// must be AVAILABLE for the task to run. For more information, see Activating
// a Sync Agent (https://docs.aws.amazon.com/sync-service/latest/userguide/working-with-sync-agents.html#activating-sync-agent)
// in the AWS DataSync User Guide.
//
// Agents are automatically updated by AWS on a regular basis, using a mechanism
// that ensures minimal interruption to your tasks.
//
//    // Example sending a request using CreateAgentRequest.
//    req := client.CreateAgentRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/datasync-2018-11-09/CreateAgent
func (c *Client) CreateAgentRequest(input *CreateAgentInput) CreateAgentRequest {
	op := &aws.Operation{
		Name:       opCreateAgent,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateAgentInput{}
	}

	req := c.newRequest(op, input, &CreateAgentOutput{})
	return CreateAgentRequest{Request: req, Input: input, Copy: c.CreateAgentRequest}
}

// CreateAgentRequest is the request type for the
// CreateAgent API operation.
type CreateAgentRequest struct {
	*aws.Request
	Input *CreateAgentInput
	Copy  func(*CreateAgentInput) CreateAgentRequest
}

// Send marshals and sends the CreateAgent API request.
func (r CreateAgentRequest) Send(ctx context.Context) (*CreateAgentResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateAgentResponse{
		CreateAgentOutput: r.Request.Data.(*CreateAgentOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateAgentResponse is the response type for the
// CreateAgent API operation.
type CreateAgentResponse struct {
	*CreateAgentOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateAgent request.
func (r *CreateAgentResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
