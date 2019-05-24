// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package elasticache

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents the input for a ModifyReplicationGroupShardConfiguration operation.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticache-2015-02-02/ModifyReplicationGroupShardConfigurationMessage
type ModifyReplicationGroupShardConfigurationInput struct {
	_ struct{} `type:"structure"`

	// Indicates that the shard reconfiguration process begins immediately. At present,
	// the only permitted value for this parameter is true.
	//
	// Value: true
	//
	// ApplyImmediately is a required field
	ApplyImmediately *bool `type:"boolean" required:"true"`

	// The number of node groups (shards) that results from the modification of
	// the shard configuration.
	//
	// NodeGroupCount is a required field
	NodeGroupCount *int64 `type:"integer" required:"true"`

	// If the value of NodeGroupCount is less than the current number of node groups
	// (shards), the NodeGroupsToRemove or NodeGroupsToRetain is a required list
	// of node group ids to remove from or retain in the cluster.
	//
	// ElastiCache for Redis will attempt to remove all node groups listed by NodeGroupsToRemove
	// from the cluster.
	NodeGroupsToRemove []string `locationNameList:"NodeGroupToRemove" type:"list"`

	// If the value of NodeGroupCount is less than the current number of node groups
	// (shards), the NodeGroupsToRemove or NodeGroupsToRetain is a required list
	// of node group ids to remove from or retain in the cluster.
	//
	// ElastiCache for Redis will attempt to remove all node groups except those
	// listed by NodeGroupsToRetain from the cluster.
	NodeGroupsToRetain []string `locationNameList:"NodeGroupToRetain" type:"list"`

	// The name of the Redis (cluster mode enabled) cluster (replication group)
	// on which the shards are to be configured.
	//
	// ReplicationGroupId is a required field
	ReplicationGroupId *string `type:"string" required:"true"`

	// Specifies the preferred availability zones for each node group in the cluster.
	// If the value of NodeGroupCount is greater than the current number of node
	// groups (shards), you can use this parameter to specify the preferred availability
	// zones of the cluster's shards. If you omit this parameter ElastiCache selects
	// availability zones for you.
	//
	// You can specify this parameter only if the value of NodeGroupCount is greater
	// than the current number of node groups (shards).
	ReshardingConfiguration []ReshardingConfiguration `locationNameList:"ReshardingConfiguration" type:"list"`
}

// String returns the string representation
func (s ModifyReplicationGroupShardConfigurationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ModifyReplicationGroupShardConfigurationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ModifyReplicationGroupShardConfigurationInput"}

	if s.ApplyImmediately == nil {
		invalidParams.Add(aws.NewErrParamRequired("ApplyImmediately"))
	}

	if s.NodeGroupCount == nil {
		invalidParams.Add(aws.NewErrParamRequired("NodeGroupCount"))
	}

	if s.ReplicationGroupId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ReplicationGroupId"))
	}
	if s.ReshardingConfiguration != nil {
		for i, v := range s.ReshardingConfiguration {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "ReshardingConfiguration", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticache-2015-02-02/ModifyReplicationGroupShardConfigurationResult
type ModifyReplicationGroupShardConfigurationOutput struct {
	_ struct{} `type:"structure"`

	// Contains all of the attributes of a specific Redis replication group.
	ReplicationGroup *ReplicationGroup `type:"structure"`
}

// String returns the string representation
func (s ModifyReplicationGroupShardConfigurationOutput) String() string {
	return awsutil.Prettify(s)
}

const opModifyReplicationGroupShardConfiguration = "ModifyReplicationGroupShardConfiguration"

// ModifyReplicationGroupShardConfigurationRequest returns a request value for making API operation for
// Amazon ElastiCache.
//
// Modifies a replication group's shards (node groups) by allowing you to add
// shards, remove shards, or rebalance the keyspaces among exisiting shards.
//
//    // Example sending a request using ModifyReplicationGroupShardConfigurationRequest.
//    req := client.ModifyReplicationGroupShardConfigurationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticache-2015-02-02/ModifyReplicationGroupShardConfiguration
func (c *Client) ModifyReplicationGroupShardConfigurationRequest(input *ModifyReplicationGroupShardConfigurationInput) ModifyReplicationGroupShardConfigurationRequest {
	op := &aws.Operation{
		Name:       opModifyReplicationGroupShardConfiguration,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ModifyReplicationGroupShardConfigurationInput{}
	}

	req := c.newRequest(op, input, &ModifyReplicationGroupShardConfigurationOutput{})
	return ModifyReplicationGroupShardConfigurationRequest{Request: req, Input: input, Copy: c.ModifyReplicationGroupShardConfigurationRequest}
}

// ModifyReplicationGroupShardConfigurationRequest is the request type for the
// ModifyReplicationGroupShardConfiguration API operation.
type ModifyReplicationGroupShardConfigurationRequest struct {
	*aws.Request
	Input *ModifyReplicationGroupShardConfigurationInput
	Copy  func(*ModifyReplicationGroupShardConfigurationInput) ModifyReplicationGroupShardConfigurationRequest
}

// Send marshals and sends the ModifyReplicationGroupShardConfiguration API request.
func (r ModifyReplicationGroupShardConfigurationRequest) Send(ctx context.Context) (*ModifyReplicationGroupShardConfigurationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ModifyReplicationGroupShardConfigurationResponse{
		ModifyReplicationGroupShardConfigurationOutput: r.Request.Data.(*ModifyReplicationGroupShardConfigurationOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ModifyReplicationGroupShardConfigurationResponse is the response type for the
// ModifyReplicationGroupShardConfiguration API operation.
type ModifyReplicationGroupShardConfigurationResponse struct {
	*ModifyReplicationGroupShardConfigurationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ModifyReplicationGroupShardConfiguration request.
func (r *ModifyReplicationGroupShardConfigurationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
