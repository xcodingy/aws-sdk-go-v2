// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package glue

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/glue-2017-03-31/BatchDeletePartitionRequest
type BatchDeletePartitionInput struct {
	_ struct{} `type:"structure"`

	// The ID of the Data Catalog where the partition to be deleted resides. If
	// none is supplied, the AWS account ID is used by default.
	CatalogId *string `min:"1" type:"string"`

	// The name of the catalog database in which the table in question resides.
	//
	// DatabaseName is a required field
	DatabaseName *string `min:"1" type:"string" required:"true"`

	// A list of PartitionInput structures that define the partitions to be deleted.
	//
	// PartitionsToDelete is a required field
	PartitionsToDelete []PartitionValueList `type:"list" required:"true"`

	// The name of the table where the partitions to be deleted is located.
	//
	// TableName is a required field
	TableName *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s BatchDeletePartitionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *BatchDeletePartitionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "BatchDeletePartitionInput"}
	if s.CatalogId != nil && len(*s.CatalogId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("CatalogId", 1))
	}

	if s.DatabaseName == nil {
		invalidParams.Add(aws.NewErrParamRequired("DatabaseName"))
	}
	if s.DatabaseName != nil && len(*s.DatabaseName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("DatabaseName", 1))
	}

	if s.PartitionsToDelete == nil {
		invalidParams.Add(aws.NewErrParamRequired("PartitionsToDelete"))
	}

	if s.TableName == nil {
		invalidParams.Add(aws.NewErrParamRequired("TableName"))
	}
	if s.TableName != nil && len(*s.TableName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("TableName", 1))
	}
	if s.PartitionsToDelete != nil {
		for i, v := range s.PartitionsToDelete {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "PartitionsToDelete", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/glue-2017-03-31/BatchDeletePartitionResponse
type BatchDeletePartitionOutput struct {
	_ struct{} `type:"structure"`

	// Errors encountered when trying to delete the requested partitions.
	Errors []PartitionError `type:"list"`
}

// String returns the string representation
func (s BatchDeletePartitionOutput) String() string {
	return awsutil.Prettify(s)
}

const opBatchDeletePartition = "BatchDeletePartition"

// BatchDeletePartitionRequest returns a request value for making API operation for
// AWS Glue.
//
// Deletes one or more partitions in a batch operation.
//
//    // Example sending a request using BatchDeletePartitionRequest.
//    req := client.BatchDeletePartitionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/glue-2017-03-31/BatchDeletePartition
func (c *Client) BatchDeletePartitionRequest(input *BatchDeletePartitionInput) BatchDeletePartitionRequest {
	op := &aws.Operation{
		Name:       opBatchDeletePartition,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &BatchDeletePartitionInput{}
	}

	req := c.newRequest(op, input, &BatchDeletePartitionOutput{})
	return BatchDeletePartitionRequest{Request: req, Input: input, Copy: c.BatchDeletePartitionRequest}
}

// BatchDeletePartitionRequest is the request type for the
// BatchDeletePartition API operation.
type BatchDeletePartitionRequest struct {
	*aws.Request
	Input *BatchDeletePartitionInput
	Copy  func(*BatchDeletePartitionInput) BatchDeletePartitionRequest
}

// Send marshals and sends the BatchDeletePartition API request.
func (r BatchDeletePartitionRequest) Send(ctx context.Context) (*BatchDeletePartitionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &BatchDeletePartitionResponse{
		BatchDeletePartitionOutput: r.Request.Data.(*BatchDeletePartitionOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// BatchDeletePartitionResponse is the response type for the
// BatchDeletePartition API operation.
type BatchDeletePartitionResponse struct {
	*BatchDeletePartitionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// BatchDeletePartition request.
func (r *BatchDeletePartitionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
