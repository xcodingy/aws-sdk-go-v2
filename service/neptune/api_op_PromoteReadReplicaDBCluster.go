// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package neptune

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/neptune-2014-10-31/PromoteReadReplicaDBClusterMessage
type PromoteReadReplicaDBClusterInput struct {
	_ struct{} `type:"structure"`

	// The identifier of the DB cluster Read Replica to promote. This parameter
	// is not case-sensitive.
	//
	// Constraints:
	//
	//    * Must match the identifier of an existing DBCluster Read Replica.
	//
	// Example: my-cluster-replica1
	//
	// DBClusterIdentifier is a required field
	DBClusterIdentifier *string `type:"string" required:"true"`
}

// String returns the string representation
func (s PromoteReadReplicaDBClusterInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PromoteReadReplicaDBClusterInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "PromoteReadReplicaDBClusterInput"}

	if s.DBClusterIdentifier == nil {
		invalidParams.Add(aws.NewErrParamRequired("DBClusterIdentifier"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/neptune-2014-10-31/PromoteReadReplicaDBClusterResult
type PromoteReadReplicaDBClusterOutput struct {
	_ struct{} `type:"structure"`

	// Contains the details of an Amazon Neptune DB cluster.
	//
	// This data type is used as a response element in the DescribeDBClusters action.
	DBCluster *DBCluster `type:"structure"`
}

// String returns the string representation
func (s PromoteReadReplicaDBClusterOutput) String() string {
	return awsutil.Prettify(s)
}

const opPromoteReadReplicaDBCluster = "PromoteReadReplicaDBCluster"

// PromoteReadReplicaDBClusterRequest returns a request value for making API operation for
// Amazon Neptune.
//
// Promotes a Read Replica DB cluster to a standalone DB cluster.
//
//    // Example sending a request using PromoteReadReplicaDBClusterRequest.
//    req := client.PromoteReadReplicaDBClusterRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/neptune-2014-10-31/PromoteReadReplicaDBCluster
func (c *Client) PromoteReadReplicaDBClusterRequest(input *PromoteReadReplicaDBClusterInput) PromoteReadReplicaDBClusterRequest {
	op := &aws.Operation{
		Name:       opPromoteReadReplicaDBCluster,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &PromoteReadReplicaDBClusterInput{}
	}

	req := c.newRequest(op, input, &PromoteReadReplicaDBClusterOutput{})
	return PromoteReadReplicaDBClusterRequest{Request: req, Input: input, Copy: c.PromoteReadReplicaDBClusterRequest}
}

// PromoteReadReplicaDBClusterRequest is the request type for the
// PromoteReadReplicaDBCluster API operation.
type PromoteReadReplicaDBClusterRequest struct {
	*aws.Request
	Input *PromoteReadReplicaDBClusterInput
	Copy  func(*PromoteReadReplicaDBClusterInput) PromoteReadReplicaDBClusterRequest
}

// Send marshals and sends the PromoteReadReplicaDBCluster API request.
func (r PromoteReadReplicaDBClusterRequest) Send(ctx context.Context) (*PromoteReadReplicaDBClusterResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &PromoteReadReplicaDBClusterResponse{
		PromoteReadReplicaDBClusterOutput: r.Request.Data.(*PromoteReadReplicaDBClusterOutput),
		response:                          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// PromoteReadReplicaDBClusterResponse is the response type for the
// PromoteReadReplicaDBCluster API operation.
type PromoteReadReplicaDBClusterResponse struct {
	*PromoteReadReplicaDBClusterOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// PromoteReadReplicaDBCluster request.
func (r *PromoteReadReplicaDBClusterResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
