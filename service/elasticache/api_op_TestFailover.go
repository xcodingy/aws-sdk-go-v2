// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package elasticache

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticache-2015-02-02/TestFailoverMessage
type TestFailoverInput struct {
	_ struct{} `type:"structure"`

	// The name of the node group (called shard in the console) in this replication
	// group on which automatic failover is to be tested. You may test automatic
	// failover on up to 5 node groups in any rolling 24-hour period.
	//
	// NodeGroupId is a required field
	NodeGroupId *string `min:"1" type:"string" required:"true"`

	// The name of the replication group (console: cluster) whose automatic failover
	// is being tested by this operation.
	//
	// ReplicationGroupId is a required field
	ReplicationGroupId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s TestFailoverInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *TestFailoverInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "TestFailoverInput"}

	if s.NodeGroupId == nil {
		invalidParams.Add(aws.NewErrParamRequired("NodeGroupId"))
	}
	if s.NodeGroupId != nil && len(*s.NodeGroupId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NodeGroupId", 1))
	}

	if s.ReplicationGroupId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ReplicationGroupId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticache-2015-02-02/TestFailoverResult
type TestFailoverOutput struct {
	_ struct{} `type:"structure"`

	// Contains all of the attributes of a specific Redis replication group.
	ReplicationGroup *ReplicationGroup `type:"structure"`
}

// String returns the string representation
func (s TestFailoverOutput) String() string {
	return awsutil.Prettify(s)
}

const opTestFailover = "TestFailover"

// TestFailoverRequest returns a request value for making API operation for
// Amazon ElastiCache.
//
// Represents the input of a TestFailover operation which test automatic failover
// on a specified node group (called shard in the console) in a replication
// group (called cluster in the console).
//
// Note the following
//
//    * A customer can use this operation to test automatic failover on up to
//    5 shards (called node groups in the ElastiCache API and AWS CLI) in any
//    rolling 24-hour period.
//
//    * If calling this operation on shards in different clusters (called replication
//    groups in the API and CLI), the calls can be made concurrently.
//
//    * If calling this operation multiple times on different shards in the
//    same Redis (cluster mode enabled) replication group, the first node replacement
//    must complete before a subsequent call can be made.
//
//    * To determine whether the node replacement is complete you can check
//    Events using the Amazon ElastiCache console, the AWS CLI, or the ElastiCache
//    API. Look for the following automatic failover related events, listed
//    here in order of occurrance:
//
// Replication group message: Test Failover API called for node group <node-group-id>
//
// Cache cluster message: Failover from master node <primary-node-id> to replica
//    node <node-id> completed
//
// Replication group message: Failover from master node <primary-node-id> to
//    replica node <node-id> completed
//
// Cache cluster message: Recovering cache nodes <node-id>
//
// Cache cluster message: Finished recovery for cache nodes <node-id>
//
// For more information see:
//
// Viewing ElastiCache Events (http://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/ECEvents.Viewing.html)
//    in the ElastiCache User Guide
//
// DescribeEvents (http://docs.aws.amazon.com/AmazonElastiCache/latest/APIReference/API_DescribeEvents.html)
//    in the ElastiCache API Reference
//
// Also see, Testing Multi-AZ with Automatic Failover (http://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/AutoFailover.html#auto-failover-test)
// in the ElastiCache User Guide.
//
//    // Example sending a request using TestFailoverRequest.
//    req := client.TestFailoverRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticache-2015-02-02/TestFailover
func (c *Client) TestFailoverRequest(input *TestFailoverInput) TestFailoverRequest {
	op := &aws.Operation{
		Name:       opTestFailover,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &TestFailoverInput{}
	}

	req := c.newRequest(op, input, &TestFailoverOutput{})
	return TestFailoverRequest{Request: req, Input: input, Copy: c.TestFailoverRequest}
}

// TestFailoverRequest is the request type for the
// TestFailover API operation.
type TestFailoverRequest struct {
	*aws.Request
	Input *TestFailoverInput
	Copy  func(*TestFailoverInput) TestFailoverRequest
}

// Send marshals and sends the TestFailover API request.
func (r TestFailoverRequest) Send(ctx context.Context) (*TestFailoverResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &TestFailoverResponse{
		TestFailoverOutput: r.Request.Data.(*TestFailoverOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// TestFailoverResponse is the response type for the
// TestFailover API operation.
type TestFailoverResponse struct {
	*TestFailoverOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// TestFailover request.
func (r *TestFailoverResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
