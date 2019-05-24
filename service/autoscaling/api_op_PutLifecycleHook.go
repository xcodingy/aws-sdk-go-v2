// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package autoscaling

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/autoscaling-2011-01-01/PutLifecycleHookType
type PutLifecycleHookInput struct {
	_ struct{} `type:"structure"`

	// The name of the Auto Scaling group.
	//
	// AutoScalingGroupName is a required field
	AutoScalingGroupName *string `min:"1" type:"string" required:"true"`

	// Defines the action the Auto Scaling group should take when the lifecycle
	// hook timeout elapses or if an unexpected failure occurs. This parameter can
	// be either CONTINUE or ABANDON. The default value is ABANDON.
	DefaultResult *string `type:"string"`

	// The maximum time, in seconds, that can elapse before the lifecycle hook times
	// out. The range is from 30 to 7200 seconds. The default value is 3600 seconds
	// (1 hour).
	//
	// If the lifecycle hook times out, Amazon EC2 Auto Scaling performs the action
	// that you specified in the DefaultResult parameter. You can prevent the lifecycle
	// hook from timing out by calling RecordLifecycleActionHeartbeat.
	HeartbeatTimeout *int64 `type:"integer"`

	// The name of the lifecycle hook.
	//
	// LifecycleHookName is a required field
	LifecycleHookName *string `min:"1" type:"string" required:"true"`

	// The instance state to which you want to attach the lifecycle hook. The valid
	// values are:
	//
	//    * autoscaling:EC2_INSTANCE_LAUNCHING
	//
	//    * autoscaling:EC2_INSTANCE_TERMINATING
	//
	// Conditional: This parameter is required for new lifecycle hooks, but optional
	// when updating existing hooks.
	LifecycleTransition *string `type:"string"`

	// Additional information that you want to include any time Amazon EC2 Auto
	// Scaling sends a message to the notification target.
	NotificationMetadata *string `min:"1" type:"string"`

	// The ARN of the notification target that Amazon EC2 Auto Scaling uses to notify
	// you when an instance is in the transition state for the lifecycle hook. This
	// target can be either an SQS queue or an SNS topic.
	//
	// If you specify an empty string, this overrides the current ARN.
	//
	// This operation uses the JSON format when sending notifications to an Amazon
	// SQS queue, and an email key-value pair format when sending notifications
	// to an Amazon SNS topic.
	//
	// When you specify a notification target, Amazon EC2 Auto Scaling sends it
	// a test message. Test messages contain the following additional key-value
	// pair: "Event": "autoscaling:TEST_NOTIFICATION".
	NotificationTargetARN *string `type:"string"`

	// The ARN of the IAM role that allows the Auto Scaling group to publish to
	// the specified notification target, for example, an Amazon SNS topic or an
	// Amazon SQS queue.
	//
	// Conditional: This parameter is required for new lifecycle hooks, but optional
	// when updating existing hooks.
	RoleARN *string `min:"1" type:"string"`
}

// String returns the string representation
func (s PutLifecycleHookInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PutLifecycleHookInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "PutLifecycleHookInput"}

	if s.AutoScalingGroupName == nil {
		invalidParams.Add(aws.NewErrParamRequired("AutoScalingGroupName"))
	}
	if s.AutoScalingGroupName != nil && len(*s.AutoScalingGroupName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("AutoScalingGroupName", 1))
	}

	if s.LifecycleHookName == nil {
		invalidParams.Add(aws.NewErrParamRequired("LifecycleHookName"))
	}
	if s.LifecycleHookName != nil && len(*s.LifecycleHookName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("LifecycleHookName", 1))
	}
	if s.NotificationMetadata != nil && len(*s.NotificationMetadata) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NotificationMetadata", 1))
	}
	if s.RoleARN != nil && len(*s.RoleARN) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("RoleARN", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/autoscaling-2011-01-01/PutLifecycleHookAnswer
type PutLifecycleHookOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s PutLifecycleHookOutput) String() string {
	return awsutil.Prettify(s)
}

const opPutLifecycleHook = "PutLifecycleHook"

// PutLifecycleHookRequest returns a request value for making API operation for
// Auto Scaling.
//
// Creates or updates a lifecycle hook for the specified Auto Scaling group.
//
// A lifecycle hook tells Amazon EC2 Auto Scaling to perform an action on an
// instance when the instance launches (before it is put into service) or as
// the instance terminates (before it is fully terminated).
//
// This step is a part of the procedure for adding a lifecycle hook to an Auto
// Scaling group:
//
// (Optional) Create a Lambda function and a rule that allows CloudWatch Events
// to invoke your Lambda function when Amazon EC2 Auto Scaling launches or terminates
// instances.
//
// (Optional) Create a notification target and an IAM role. The target can be
// either an Amazon SQS queue or an Amazon SNS topic. The role allows Amazon
// EC2 Auto Scaling to publish lifecycle notifications to the target.
//
// Create the lifecycle hook. Specify whether the hook is used when the instances
// launch or terminate.
//
// If you need more time, record the lifecycle action heartbeat to keep the
// instance in a pending state using using RecordLifecycleActionHeartbeat.
//
// If you finish before the timeout period ends, complete the lifecycle action
// using CompleteLifecycleAction.
//
// For more information, see Amazon EC2 Auto Scaling Lifecycle Hooks (https://docs.aws.amazon.com/autoscaling/ec2/userguide/lifecycle-hooks.html)
// in the Amazon EC2 Auto Scaling User Guide.
//
// If you exceed your maximum limit of lifecycle hooks, which by default is
// 50 per Auto Scaling group, the call fails.
//
// You can view the lifecycle hooks for an Auto Scaling group using DescribeLifecycleHooks.
// If you are no longer using a lifecycle hook, you can delete it using DeleteLifecycleHook.
//
//    // Example sending a request using PutLifecycleHookRequest.
//    req := client.PutLifecycleHookRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/autoscaling-2011-01-01/PutLifecycleHook
func (c *Client) PutLifecycleHookRequest(input *PutLifecycleHookInput) PutLifecycleHookRequest {
	op := &aws.Operation{
		Name:       opPutLifecycleHook,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &PutLifecycleHookInput{}
	}

	req := c.newRequest(op, input, &PutLifecycleHookOutput{})
	return PutLifecycleHookRequest{Request: req, Input: input, Copy: c.PutLifecycleHookRequest}
}

// PutLifecycleHookRequest is the request type for the
// PutLifecycleHook API operation.
type PutLifecycleHookRequest struct {
	*aws.Request
	Input *PutLifecycleHookInput
	Copy  func(*PutLifecycleHookInput) PutLifecycleHookRequest
}

// Send marshals and sends the PutLifecycleHook API request.
func (r PutLifecycleHookRequest) Send(ctx context.Context) (*PutLifecycleHookResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &PutLifecycleHookResponse{
		PutLifecycleHookOutput: r.Request.Data.(*PutLifecycleHookOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// PutLifecycleHookResponse is the response type for the
// PutLifecycleHook API operation.
type PutLifecycleHookResponse struct {
	*PutLifecycleHookOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// PutLifecycleHook request.
func (r *PutLifecycleHookResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
