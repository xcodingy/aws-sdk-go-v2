// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ssm

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ssm-2014-11-06/DeleteInventoryRequest
type DeleteInventoryInput struct {
	_ struct{} `type:"structure"`

	// User-provided idempotency token.
	ClientToken *string `min:"1" type:"string" idempotencyToken:"true"`

	// Use this option to view a summary of the deletion request without deleting
	// any data or the data type. This option is useful when you only want to understand
	// what will be deleted. Once you validate that the data to be deleted is what
	// you intend to delete, you can run the same command without specifying the
	// DryRun option.
	DryRun *bool `type:"boolean"`

	// Use the SchemaDeleteOption to delete a custom inventory type (schema). If
	// you don't choose this option, the system only deletes existing inventory
	// data associated with the custom inventory type. Choose one of the following
	// options:
	//
	// DisableSchema: If you choose this option, the system ignores all inventory
	// data for the specified version, and any earlier versions. To enable this
	// schema again, you must call the PutInventory action for a version greater
	// than the disbled version.
	//
	// DeleteSchema: This option deletes the specified custom type from the Inventory
	// service. You can recreate the schema later, if you want.
	SchemaDeleteOption InventorySchemaDeleteOption `type:"string" enum:"true"`

	// The name of the custom inventory type for which you want to delete either
	// all previously collected data, or the inventory type itself.
	//
	// TypeName is a required field
	TypeName *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteInventoryInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteInventoryInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteInventoryInput"}
	if s.ClientToken != nil && len(*s.ClientToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ClientToken", 1))
	}

	if s.TypeName == nil {
		invalidParams.Add(aws.NewErrParamRequired("TypeName"))
	}
	if s.TypeName != nil && len(*s.TypeName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("TypeName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ssm-2014-11-06/DeleteInventoryResult
type DeleteInventoryOutput struct {
	_ struct{} `type:"structure"`

	// Every DeleteInventory action is assigned a unique ID. This option returns
	// a unique ID. You can use this ID to query the status of a delete operation.
	// This option is useful for ensuring that a delete operation has completed
	// before you begin other actions.
	DeletionId *string `type:"string"`

	// A summary of the delete operation. For more information about this summary,
	// see Understanding the Delete Inventory Summary (http://docs.aws.amazon.com/systems-manager/latest/userguide/sysman-inventory-delete.html#sysman-inventory-delete-summary)
	// in the AWS Systems Manager User Guide.
	DeletionSummary *InventoryDeletionSummary `type:"structure"`

	// The name of the inventory data type specified in the request.
	TypeName *string `min:"1" type:"string"`
}

// String returns the string representation
func (s DeleteInventoryOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteInventory = "DeleteInventory"

// DeleteInventoryRequest returns a request value for making API operation for
// Amazon Simple Systems Manager (SSM).
//
// Delete a custom inventory type, or the data associated with a custom Inventory
// type. Deleting a custom inventory type is also referred to as deleting a
// custom inventory schema.
//
//    // Example sending a request using DeleteInventoryRequest.
//    req := client.DeleteInventoryRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ssm-2014-11-06/DeleteInventory
func (c *Client) DeleteInventoryRequest(input *DeleteInventoryInput) DeleteInventoryRequest {
	op := &aws.Operation{
		Name:       opDeleteInventory,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteInventoryInput{}
	}

	req := c.newRequest(op, input, &DeleteInventoryOutput{})
	return DeleteInventoryRequest{Request: req, Input: input, Copy: c.DeleteInventoryRequest}
}

// DeleteInventoryRequest is the request type for the
// DeleteInventory API operation.
type DeleteInventoryRequest struct {
	*aws.Request
	Input *DeleteInventoryInput
	Copy  func(*DeleteInventoryInput) DeleteInventoryRequest
}

// Send marshals and sends the DeleteInventory API request.
func (r DeleteInventoryRequest) Send(ctx context.Context) (*DeleteInventoryResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteInventoryResponse{
		DeleteInventoryOutput: r.Request.Data.(*DeleteInventoryOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteInventoryResponse is the response type for the
// DeleteInventory API operation.
type DeleteInventoryResponse struct {
	*DeleteInventoryOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteInventory request.
func (r *DeleteInventoryResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
