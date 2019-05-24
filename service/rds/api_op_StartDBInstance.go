// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package rds

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/rds-2014-10-31/StartDBInstanceMessage
type StartDBInstanceInput struct {
	_ struct{} `type:"structure"`

	// The user-supplied instance identifier.
	//
	// DBInstanceIdentifier is a required field
	DBInstanceIdentifier *string `type:"string" required:"true"`
}

// String returns the string representation
func (s StartDBInstanceInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *StartDBInstanceInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "StartDBInstanceInput"}

	if s.DBInstanceIdentifier == nil {
		invalidParams.Add(aws.NewErrParamRequired("DBInstanceIdentifier"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/rds-2014-10-31/StartDBInstanceResult
type StartDBInstanceOutput struct {
	_ struct{} `type:"structure"`

	// Contains the details of an Amazon RDS DB instance.
	//
	// This data type is used as a response element in the DescribeDBInstances action.
	DBInstance *DBInstance `type:"structure"`
}

// String returns the string representation
func (s StartDBInstanceOutput) String() string {
	return awsutil.Prettify(s)
}

const opStartDBInstance = "StartDBInstance"

// StartDBInstanceRequest returns a request value for making API operation for
// Amazon Relational Database Service.
//
// Starts an Amazon RDS DB instance that was stopped using the AWS console,
// the stop-db-instance AWS CLI command, or the StopDBInstance action.
//
// For more information, see  Starting an Amazon RDS DB instance That Was Previously
// Stopped (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_StartInstance.html)
// in the Amazon RDS User Guide.
//
// This command doesn't apply to Aurora MySQL and Aurora PostgreSQL. For Aurora
// DB clusters, use StartDBCluster instead.
//
//    // Example sending a request using StartDBInstanceRequest.
//    req := client.StartDBInstanceRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/rds-2014-10-31/StartDBInstance
func (c *Client) StartDBInstanceRequest(input *StartDBInstanceInput) StartDBInstanceRequest {
	op := &aws.Operation{
		Name:       opStartDBInstance,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &StartDBInstanceInput{}
	}

	req := c.newRequest(op, input, &StartDBInstanceOutput{})
	return StartDBInstanceRequest{Request: req, Input: input, Copy: c.StartDBInstanceRequest}
}

// StartDBInstanceRequest is the request type for the
// StartDBInstance API operation.
type StartDBInstanceRequest struct {
	*aws.Request
	Input *StartDBInstanceInput
	Copy  func(*StartDBInstanceInput) StartDBInstanceRequest
}

// Send marshals and sends the StartDBInstance API request.
func (r StartDBInstanceRequest) Send(ctx context.Context) (*StartDBInstanceResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &StartDBInstanceResponse{
		StartDBInstanceOutput: r.Request.Data.(*StartDBInstanceOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// StartDBInstanceResponse is the response type for the
// StartDBInstance API operation.
type StartDBInstanceResponse struct {
	*StartDBInstanceOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// StartDBInstance request.
func (r *StartDBInstanceResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
