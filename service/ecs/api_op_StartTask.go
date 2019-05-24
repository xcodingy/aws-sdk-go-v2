// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ecs

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ecs-2014-11-13/StartTaskRequest
type StartTaskInput struct {
	_ struct{} `type:"structure"`

	// The short name or full Amazon Resource Name (ARN) of the cluster on which
	// to start your task. If you do not specify a cluster, the default cluster
	// is assumed.
	Cluster *string `locationName:"cluster" type:"string"`

	// The container instance IDs or full ARN entries for the container instances
	// on which you would like to place your task. You can specify up to 10 container
	// instances.
	//
	// ContainerInstances is a required field
	ContainerInstances []string `locationName:"containerInstances" type:"list" required:"true"`

	// Specifies whether to enable Amazon ECS managed tags for the task. For more
	// information, see Tagging Your Amazon ECS Resources (http://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-using-tags.html)
	// in the Amazon Elastic Container Service Developer Guide.
	EnableECSManagedTags *bool `locationName:"enableECSManagedTags" type:"boolean"`

	// The name of the task group to associate with the task. The default value
	// is the family name of the task definition (for example, family:my-family-name).
	Group *string `locationName:"group" type:"string"`

	// The VPC subnet and security group configuration for tasks that receive their
	// own elastic network interface by using the awsvpc networking mode.
	NetworkConfiguration *NetworkConfiguration `locationName:"networkConfiguration" type:"structure"`

	// A list of container overrides in JSON format that specify the name of a container
	// in the specified task definition and the overrides it should receive. You
	// can override the default command for a container (that is specified in the
	// task definition or Docker image) with a command override. You can also override
	// existing environment variables (that are specified in the task definition
	// or Docker image) on a container or add new environment variables to it with
	// an environment override.
	//
	// A total of 8192 characters are allowed for overrides. This limit includes
	// the JSON formatting characters of the override structure.
	Overrides *TaskOverride `locationName:"overrides" type:"structure"`

	// Specifies whether to propagate the tags from the task definition or the service
	// to the task. If no value is specified, the tags are not propagated.
	PropagateTags PropagateTags `locationName:"propagateTags" type:"string" enum:"true"`

	// An optional tag specified when a task is started. For example, if you automatically
	// trigger a task to run a batch process job, you could apply a unique identifier
	// for that job to your task with the startedBy parameter. You can then identify
	// which tasks belong to that job by filtering the results of a ListTasks call
	// with the startedBy value. Up to 36 letters (uppercase and lowercase), numbers,
	// hyphens, and underscores are allowed.
	//
	// If a task is started by an Amazon ECS service, then the startedBy parameter
	// contains the deployment ID of the service that starts it.
	StartedBy *string `locationName:"startedBy" type:"string"`

	// The metadata that you apply to the task to help you categorize and organize
	// them. Each tag consists of a key and an optional value, both of which you
	// define. Tag keys can have a maximum character length of 128 characters, and
	// tag values can have a maximum length of 256 characters.
	Tags []Tag `locationName:"tags" type:"list"`

	// The family and revision (family:revision) or full ARN of the task definition
	// to start. If a revision is not specified, the latest ACTIVE revision is used.
	//
	// TaskDefinition is a required field
	TaskDefinition *string `locationName:"taskDefinition" type:"string" required:"true"`
}

// String returns the string representation
func (s StartTaskInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *StartTaskInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "StartTaskInput"}

	if s.ContainerInstances == nil {
		invalidParams.Add(aws.NewErrParamRequired("ContainerInstances"))
	}

	if s.TaskDefinition == nil {
		invalidParams.Add(aws.NewErrParamRequired("TaskDefinition"))
	}
	if s.NetworkConfiguration != nil {
		if err := s.NetworkConfiguration.Validate(); err != nil {
			invalidParams.AddNested("NetworkConfiguration", err.(aws.ErrInvalidParams))
		}
	}
	if s.Overrides != nil {
		if err := s.Overrides.Validate(); err != nil {
			invalidParams.AddNested("Overrides", err.(aws.ErrInvalidParams))
		}
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

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ecs-2014-11-13/StartTaskResponse
type StartTaskOutput struct {
	_ struct{} `type:"structure"`

	// Any failures associated with the call.
	Failures []Failure `locationName:"failures" type:"list"`

	// A full description of the tasks that were started. Each task that was successfully
	// placed on your container instances is described.
	Tasks []Task `locationName:"tasks" type:"list"`
}

// String returns the string representation
func (s StartTaskOutput) String() string {
	return awsutil.Prettify(s)
}

const opStartTask = "StartTask"

// StartTaskRequest returns a request value for making API operation for
// Amazon EC2 Container Service.
//
// Starts a new task from the specified task definition on the specified container
// instance or instances.
//
// Alternatively, you can use RunTask to place tasks for you. For more information,
// see Scheduling Tasks (https://docs.aws.amazon.com/AmazonECS/latest/developerguide/scheduling_tasks.html)
// in the Amazon Elastic Container Service Developer Guide.
//
//    // Example sending a request using StartTaskRequest.
//    req := client.StartTaskRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ecs-2014-11-13/StartTask
func (c *Client) StartTaskRequest(input *StartTaskInput) StartTaskRequest {
	op := &aws.Operation{
		Name:       opStartTask,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &StartTaskInput{}
	}

	req := c.newRequest(op, input, &StartTaskOutput{})
	return StartTaskRequest{Request: req, Input: input, Copy: c.StartTaskRequest}
}

// StartTaskRequest is the request type for the
// StartTask API operation.
type StartTaskRequest struct {
	*aws.Request
	Input *StartTaskInput
	Copy  func(*StartTaskInput) StartTaskRequest
}

// Send marshals and sends the StartTask API request.
func (r StartTaskRequest) Send(ctx context.Context) (*StartTaskResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &StartTaskResponse{
		StartTaskOutput: r.Request.Data.(*StartTaskOutput),
		response:        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// StartTaskResponse is the response type for the
// StartTask API operation.
type StartTaskResponse struct {
	*StartTaskOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// StartTask request.
func (r *StartTaskResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
