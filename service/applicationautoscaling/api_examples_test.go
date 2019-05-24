// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package applicationautoscaling_test

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/awserr"
	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/applicationautoscaling"
)

var _ time.Duration
var _ strings.Reader
var _ aws.Config

func parseTime(layout, value string) *time.Time {
	t, err := time.Parse(layout, value)
	if err != nil {
		panic(err)
	}
	return &t
}

// To delete a scaling policy
//
// This example deletes a scaling policy for the Amazon ECS service called web-app,
// which is running in the default cluster.
func ExampleClient_DeleteScalingPolicyRequest_shared00() {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		panic("failed to load config, " + err.Error())
	}

	svc := applicationautoscaling.New(cfg)
	input := &applicationautoscaling.DeleteScalingPolicyInput{
		PolicyName:        aws.String("web-app-cpu-lt-25"),
		ResourceId:        aws.String("service/default/web-app"),
		ScalableDimension: applicationautoscaling.ScalableDimensionEcsServiceDesiredCount,
		ServiceNamespace:  applicationautoscaling.ServiceNamespaceEcs,
	}

	req := svc.DeleteScalingPolicyRequest(input)
	result, err := req.Send(context.Background())
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case applicationautoscaling.ErrCodeValidationException:
				fmt.Println(applicationautoscaling.ErrCodeValidationException, aerr.Error())
			case applicationautoscaling.ErrCodeObjectNotFoundException:
				fmt.Println(applicationautoscaling.ErrCodeObjectNotFoundException, aerr.Error())
			case applicationautoscaling.ErrCodeConcurrentUpdateException:
				fmt.Println(applicationautoscaling.ErrCodeConcurrentUpdateException, aerr.Error())
			case applicationautoscaling.ErrCodeInternalServiceException:
				fmt.Println(applicationautoscaling.ErrCodeInternalServiceException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To deregister a scalable target
//
// This example deregisters a scalable target for an Amazon ECS service called web-app
// that is running in the default cluster.
func ExampleClient_DeregisterScalableTargetRequest_shared00() {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		panic("failed to load config, " + err.Error())
	}

	svc := applicationautoscaling.New(cfg)
	input := &applicationautoscaling.DeregisterScalableTargetInput{
		ResourceId:        aws.String("service/default/web-app"),
		ScalableDimension: applicationautoscaling.ScalableDimensionEcsServiceDesiredCount,
		ServiceNamespace:  applicationautoscaling.ServiceNamespaceEcs,
	}

	req := svc.DeregisterScalableTargetRequest(input)
	result, err := req.Send(context.Background())
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case applicationautoscaling.ErrCodeValidationException:
				fmt.Println(applicationautoscaling.ErrCodeValidationException, aerr.Error())
			case applicationautoscaling.ErrCodeObjectNotFoundException:
				fmt.Println(applicationautoscaling.ErrCodeObjectNotFoundException, aerr.Error())
			case applicationautoscaling.ErrCodeConcurrentUpdateException:
				fmt.Println(applicationautoscaling.ErrCodeConcurrentUpdateException, aerr.Error())
			case applicationautoscaling.ErrCodeInternalServiceException:
				fmt.Println(applicationautoscaling.ErrCodeInternalServiceException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To describe scalable targets
//
// This example describes the scalable targets for the ecs service namespace.
func ExampleClient_DescribeScalableTargetsRequest_shared00() {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		panic("failed to load config, " + err.Error())
	}

	svc := applicationautoscaling.New(cfg)
	input := &applicationautoscaling.DescribeScalableTargetsInput{
		ServiceNamespace: applicationautoscaling.ServiceNamespaceEcs,
	}

	req := svc.DescribeScalableTargetsRequest(input)
	result, err := req.Send(context.Background())
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case applicationautoscaling.ErrCodeValidationException:
				fmt.Println(applicationautoscaling.ErrCodeValidationException, aerr.Error())
			case applicationautoscaling.ErrCodeInvalidNextTokenException:
				fmt.Println(applicationautoscaling.ErrCodeInvalidNextTokenException, aerr.Error())
			case applicationautoscaling.ErrCodeConcurrentUpdateException:
				fmt.Println(applicationautoscaling.ErrCodeConcurrentUpdateException, aerr.Error())
			case applicationautoscaling.ErrCodeInternalServiceException:
				fmt.Println(applicationautoscaling.ErrCodeInternalServiceException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To describe scaling activities for a scalable target
//
// This example describes the scaling activities for an Amazon ECS service called web-app
// that is running in the default cluster.
func ExampleClient_DescribeScalingActivitiesRequest_shared00() {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		panic("failed to load config, " + err.Error())
	}

	svc := applicationautoscaling.New(cfg)
	input := &applicationautoscaling.DescribeScalingActivitiesInput{
		ResourceId:        aws.String("service/default/web-app"),
		ScalableDimension: applicationautoscaling.ScalableDimensionEcsServiceDesiredCount,
		ServiceNamespace:  applicationautoscaling.ServiceNamespaceEcs,
	}

	req := svc.DescribeScalingActivitiesRequest(input)
	result, err := req.Send(context.Background())
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case applicationautoscaling.ErrCodeValidationException:
				fmt.Println(applicationautoscaling.ErrCodeValidationException, aerr.Error())
			case applicationautoscaling.ErrCodeInvalidNextTokenException:
				fmt.Println(applicationautoscaling.ErrCodeInvalidNextTokenException, aerr.Error())
			case applicationautoscaling.ErrCodeConcurrentUpdateException:
				fmt.Println(applicationautoscaling.ErrCodeConcurrentUpdateException, aerr.Error())
			case applicationautoscaling.ErrCodeInternalServiceException:
				fmt.Println(applicationautoscaling.ErrCodeInternalServiceException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To describe scaling policies
//
// This example describes the scaling policies for the ecs service namespace.
func ExampleClient_DescribeScalingPoliciesRequest_shared00() {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		panic("failed to load config, " + err.Error())
	}

	svc := applicationautoscaling.New(cfg)
	input := &applicationautoscaling.DescribeScalingPoliciesInput{
		ServiceNamespace: applicationautoscaling.ServiceNamespaceEcs,
	}

	req := svc.DescribeScalingPoliciesRequest(input)
	result, err := req.Send(context.Background())
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case applicationautoscaling.ErrCodeValidationException:
				fmt.Println(applicationautoscaling.ErrCodeValidationException, aerr.Error())
			case applicationautoscaling.ErrCodeFailedResourceAccessException:
				fmt.Println(applicationautoscaling.ErrCodeFailedResourceAccessException, aerr.Error())
			case applicationautoscaling.ErrCodeInvalidNextTokenException:
				fmt.Println(applicationautoscaling.ErrCodeInvalidNextTokenException, aerr.Error())
			case applicationautoscaling.ErrCodeConcurrentUpdateException:
				fmt.Println(applicationautoscaling.ErrCodeConcurrentUpdateException, aerr.Error())
			case applicationautoscaling.ErrCodeInternalServiceException:
				fmt.Println(applicationautoscaling.ErrCodeInternalServiceException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To apply a scaling policy to an Amazon ECS service
//
// This example applies a scaling policy to an Amazon ECS service called web-app in
// the default cluster. The policy increases the desired count of the service by 200%,
// with a cool down period of 60 seconds.
func ExampleClient_PutScalingPolicyRequest_shared00() {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		panic("failed to load config, " + err.Error())
	}

	svc := applicationautoscaling.New(cfg)
	input := &applicationautoscaling.PutScalingPolicyInput{
		PolicyName:        aws.String("web-app-cpu-gt-75"),
		PolicyType:        applicationautoscaling.PolicyTypeStepScaling,
		ResourceId:        aws.String("service/default/web-app"),
		ScalableDimension: applicationautoscaling.ScalableDimensionEcsServiceDesiredCount,
		ServiceNamespace:  applicationautoscaling.ServiceNamespaceEcs,
		StepScalingPolicyConfiguration: &applicationautoscaling.StepScalingPolicyConfiguration{
			AdjustmentType: applicationautoscaling.AdjustmentTypePercentChangeInCapacity,
			Cooldown:       aws.Int64(60),
			StepAdjustments: []applicationautoscaling.StepAdjustment{
				{
					MetricIntervalLowerBound: aws.Float64(0.000000),
					ScalingAdjustment:        aws.Int64(200),
				},
			},
		},
	}

	req := svc.PutScalingPolicyRequest(input)
	result, err := req.Send(context.Background())
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case applicationautoscaling.ErrCodeValidationException:
				fmt.Println(applicationautoscaling.ErrCodeValidationException, aerr.Error())
			case applicationautoscaling.ErrCodeLimitExceededException:
				fmt.Println(applicationautoscaling.ErrCodeLimitExceededException, aerr.Error())
			case applicationautoscaling.ErrCodeObjectNotFoundException:
				fmt.Println(applicationautoscaling.ErrCodeObjectNotFoundException, aerr.Error())
			case applicationautoscaling.ErrCodeConcurrentUpdateException:
				fmt.Println(applicationautoscaling.ErrCodeConcurrentUpdateException, aerr.Error())
			case applicationautoscaling.ErrCodeFailedResourceAccessException:
				fmt.Println(applicationautoscaling.ErrCodeFailedResourceAccessException, aerr.Error())
			case applicationautoscaling.ErrCodeInternalServiceException:
				fmt.Println(applicationautoscaling.ErrCodeInternalServiceException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To apply a scaling policy to an Amazon EC2 Spot fleet
//
// This example applies a scaling policy to an Amazon EC2 Spot fleet. The policy increases
// the target capacity of the spot fleet by 200%, with a cool down period of 180 seconds.",
//
//
func ExampleClient_PutScalingPolicyRequest_shared01() {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		panic("failed to load config, " + err.Error())
	}

	svc := applicationautoscaling.New(cfg)
	input := &applicationautoscaling.PutScalingPolicyInput{
		PolicyName:        aws.String("fleet-cpu-gt-75"),
		PolicyType:        applicationautoscaling.PolicyTypeStepScaling,
		ResourceId:        aws.String("spot-fleet-request/sfr-45e69d8a-be48-4539-bbf3-3464e99c50c3"),
		ScalableDimension: applicationautoscaling.ScalableDimensionEc2SpotFleetRequestTargetCapacity,
		ServiceNamespace:  applicationautoscaling.ServiceNamespaceEc2,
		StepScalingPolicyConfiguration: &applicationautoscaling.StepScalingPolicyConfiguration{
			AdjustmentType: applicationautoscaling.AdjustmentTypePercentChangeInCapacity,
			Cooldown:       aws.Int64(180),
			StepAdjustments: []applicationautoscaling.StepAdjustment{
				{
					MetricIntervalLowerBound: aws.Float64(0.000000),
					ScalingAdjustment:        aws.Int64(200),
				},
			},
		},
	}

	req := svc.PutScalingPolicyRequest(input)
	result, err := req.Send(context.Background())
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case applicationautoscaling.ErrCodeValidationException:
				fmt.Println(applicationautoscaling.ErrCodeValidationException, aerr.Error())
			case applicationautoscaling.ErrCodeLimitExceededException:
				fmt.Println(applicationautoscaling.ErrCodeLimitExceededException, aerr.Error())
			case applicationautoscaling.ErrCodeObjectNotFoundException:
				fmt.Println(applicationautoscaling.ErrCodeObjectNotFoundException, aerr.Error())
			case applicationautoscaling.ErrCodeConcurrentUpdateException:
				fmt.Println(applicationautoscaling.ErrCodeConcurrentUpdateException, aerr.Error())
			case applicationautoscaling.ErrCodeFailedResourceAccessException:
				fmt.Println(applicationautoscaling.ErrCodeFailedResourceAccessException, aerr.Error())
			case applicationautoscaling.ErrCodeInternalServiceException:
				fmt.Println(applicationautoscaling.ErrCodeInternalServiceException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To register an ECS service as a scalable target
//
// This example registers a scalable target from an Amazon ECS service called web-app
// that is running on the default cluster, with a minimum desired count of 1 task and
// a maximum desired count of 10 tasks.
func ExampleClient_RegisterScalableTargetRequest_shared00() {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		panic("failed to load config, " + err.Error())
	}

	svc := applicationautoscaling.New(cfg)
	input := &applicationautoscaling.RegisterScalableTargetInput{
		MaxCapacity:       aws.Int64(10),
		MinCapacity:       aws.Int64(1),
		ResourceId:        aws.String("service/default/web-app"),
		RoleARN:           aws.String("arn:aws:iam::012345678910:role/ApplicationAutoscalingECSRole"),
		ScalableDimension: applicationautoscaling.ScalableDimensionEcsServiceDesiredCount,
		ServiceNamespace:  applicationautoscaling.ServiceNamespaceEcs,
	}

	req := svc.RegisterScalableTargetRequest(input)
	result, err := req.Send(context.Background())
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case applicationautoscaling.ErrCodeValidationException:
				fmt.Println(applicationautoscaling.ErrCodeValidationException, aerr.Error())
			case applicationautoscaling.ErrCodeLimitExceededException:
				fmt.Println(applicationautoscaling.ErrCodeLimitExceededException, aerr.Error())
			case applicationautoscaling.ErrCodeConcurrentUpdateException:
				fmt.Println(applicationautoscaling.ErrCodeConcurrentUpdateException, aerr.Error())
			case applicationautoscaling.ErrCodeInternalServiceException:
				fmt.Println(applicationautoscaling.ErrCodeInternalServiceException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To register an EC2 Spot fleet as a scalable target
//
// This example registers a scalable target from an Amazon EC2 Spot fleet with a minimum
// target capacity of 1 and a maximum of 10.
func ExampleClient_RegisterScalableTargetRequest_shared01() {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		panic("failed to load config, " + err.Error())
	}

	svc := applicationautoscaling.New(cfg)
	input := &applicationautoscaling.RegisterScalableTargetInput{
		MaxCapacity:       aws.Int64(10),
		MinCapacity:       aws.Int64(1),
		ResourceId:        aws.String("spot-fleet-request/sfr-45e69d8a-be48-4539-bbf3-3464e99c50c3"),
		RoleARN:           aws.String("arn:aws:iam::012345678910:role/ApplicationAutoscalingSpotRole"),
		ScalableDimension: applicationautoscaling.ScalableDimensionEc2SpotFleetRequestTargetCapacity,
		ServiceNamespace:  applicationautoscaling.ServiceNamespaceEc2,
	}

	req := svc.RegisterScalableTargetRequest(input)
	result, err := req.Send(context.Background())
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case applicationautoscaling.ErrCodeValidationException:
				fmt.Println(applicationautoscaling.ErrCodeValidationException, aerr.Error())
			case applicationautoscaling.ErrCodeLimitExceededException:
				fmt.Println(applicationautoscaling.ErrCodeLimitExceededException, aerr.Error())
			case applicationautoscaling.ErrCodeConcurrentUpdateException:
				fmt.Println(applicationautoscaling.ErrCodeConcurrentUpdateException, aerr.Error())
			case applicationautoscaling.ErrCodeInternalServiceException:
				fmt.Println(applicationautoscaling.ErrCodeInternalServiceException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}
