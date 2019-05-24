// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package kms

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/kms-2014-11-01/DeleteCustomKeyStoreRequest
type DeleteCustomKeyStoreInput struct {
	_ struct{} `type:"structure"`

	// Enter the ID of the custom key store you want to delete. To find the ID of
	// a custom key store, use the DescribeCustomKeyStores operation.
	//
	// CustomKeyStoreId is a required field
	CustomKeyStoreId *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteCustomKeyStoreInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteCustomKeyStoreInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteCustomKeyStoreInput"}

	if s.CustomKeyStoreId == nil {
		invalidParams.Add(aws.NewErrParamRequired("CustomKeyStoreId"))
	}
	if s.CustomKeyStoreId != nil && len(*s.CustomKeyStoreId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("CustomKeyStoreId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/kms-2014-11-01/DeleteCustomKeyStoreResponse
type DeleteCustomKeyStoreOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteCustomKeyStoreOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteCustomKeyStore = "DeleteCustomKeyStore"

// DeleteCustomKeyStoreRequest returns a request value for making API operation for
// AWS Key Management Service.
//
// Deletes a custom key store (http://docs.aws.amazon.com/kms/latest/developerguide/key-store-overview.html).
// This operation does not delete the AWS CloudHSM cluster that is associated
// with the custom key store, or affect any users or keys in the cluster.
//
// The custom key store that you delete cannot contain any AWS KMS customer
// master keys (CMKs) (http://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#master_keys).
// Before deleting the key store, verify that you will never need to use any
// of the CMKs in the key store for any cryptographic operations. Then, use
// ScheduleKeyDeletion to delete the AWS KMS customer master keys (CMKs) from
// the key store. When the scheduled waiting period expires, the ScheduleKeyDeletion
// operation deletes the CMKs. Then it makes a best effort to delete the key
// material from the associated cluster. However, you might need to manually
// delete the orphaned key material (http://docs.aws.amazon.com/kms/latest/developerguide/fix-keystore.html#fix-keystore-orphaned-key)
// from the cluster and its backups.
//
// After all CMKs are deleted from AWS KMS, use DisconnectCustomKeyStore to
// disconnect the key store from AWS KMS. Then, you can delete the custom key
// store.
//
// Instead of deleting the custom key store, consider using DisconnectCustomKeyStore
// to disconnect it from AWS KMS. While the key store is disconnected, you cannot
// create or use the CMKs in the key store. But, you do not need to delete CMKs
// and you can reconnect a disconnected custom key store at any time.
//
// If the operation succeeds, it returns a JSON object with no properties.
//
// This operation is part of the Custom Key Store feature (http://docs.aws.amazon.com/kms/latest/developerguide/custom-key-store-overview.html)
// feature in AWS KMS, which combines the convenience and extensive integration
// of AWS KMS with the isolation and control of a single-tenant key store.
//
//    // Example sending a request using DeleteCustomKeyStoreRequest.
//    req := client.DeleteCustomKeyStoreRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/kms-2014-11-01/DeleteCustomKeyStore
func (c *Client) DeleteCustomKeyStoreRequest(input *DeleteCustomKeyStoreInput) DeleteCustomKeyStoreRequest {
	op := &aws.Operation{
		Name:       opDeleteCustomKeyStore,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteCustomKeyStoreInput{}
	}

	req := c.newRequest(op, input, &DeleteCustomKeyStoreOutput{})
	return DeleteCustomKeyStoreRequest{Request: req, Input: input, Copy: c.DeleteCustomKeyStoreRequest}
}

// DeleteCustomKeyStoreRequest is the request type for the
// DeleteCustomKeyStore API operation.
type DeleteCustomKeyStoreRequest struct {
	*aws.Request
	Input *DeleteCustomKeyStoreInput
	Copy  func(*DeleteCustomKeyStoreInput) DeleteCustomKeyStoreRequest
}

// Send marshals and sends the DeleteCustomKeyStore API request.
func (r DeleteCustomKeyStoreRequest) Send(ctx context.Context) (*DeleteCustomKeyStoreResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteCustomKeyStoreResponse{
		DeleteCustomKeyStoreOutput: r.Request.Data.(*DeleteCustomKeyStoreOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteCustomKeyStoreResponse is the response type for the
// DeleteCustomKeyStore API operation.
type DeleteCustomKeyStoreResponse struct {
	*DeleteCustomKeyStoreOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteCustomKeyStore request.
func (r *DeleteCustomKeyStoreResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
