// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package kms

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/kms-2014-11-01/ScheduleKeyDeletionRequest
type ScheduleKeyDeletionInput struct {
	_ struct{} `type:"structure"`

	// The unique identifier of the customer master key (CMK) to delete.
	//
	// Specify the key ID or the Amazon Resource Name (ARN) of the CMK.
	//
	// For example:
	//
	//    * Key ID: 1234abcd-12ab-34cd-56ef-1234567890ab
	//
	//    * Key ARN: arn:aws:kms:us-east-2:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab
	//
	// To get the key ID and key ARN for a CMK, use ListKeys or DescribeKey.
	//
	// KeyId is a required field
	KeyId *string `min:"1" type:"string" required:"true"`

	// The waiting period, specified in number of days. After the waiting period
	// ends, AWS KMS deletes the customer master key (CMK).
	//
	// This value is optional. If you include a value, it must be between 7 and
	// 30, inclusive. If you do not include a value, it defaults to 30.
	PendingWindowInDays *int64 `min:"1" type:"integer"`
}

// String returns the string representation
func (s ScheduleKeyDeletionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ScheduleKeyDeletionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ScheduleKeyDeletionInput"}

	if s.KeyId == nil {
		invalidParams.Add(aws.NewErrParamRequired("KeyId"))
	}
	if s.KeyId != nil && len(*s.KeyId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("KeyId", 1))
	}
	if s.PendingWindowInDays != nil && *s.PendingWindowInDays < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("PendingWindowInDays", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/kms-2014-11-01/ScheduleKeyDeletionResponse
type ScheduleKeyDeletionOutput struct {
	_ struct{} `type:"structure"`

	// The date and time after which AWS KMS deletes the customer master key (CMK).
	DeletionDate *time.Time `type:"timestamp" timestampFormat:"unix"`

	// The unique identifier of the customer master key (CMK) for which deletion
	// is scheduled.
	KeyId *string `min:"1" type:"string"`
}

// String returns the string representation
func (s ScheduleKeyDeletionOutput) String() string {
	return awsutil.Prettify(s)
}

const opScheduleKeyDeletion = "ScheduleKeyDeletion"

// ScheduleKeyDeletionRequest returns a request value for making API operation for
// AWS Key Management Service.
//
// Schedules the deletion of a customer master key (CMK). You may provide a
// waiting period, specified in days, before deletion occurs. If you do not
// provide a waiting period, the default period of 30 days is used. When this
// operation is successful, the key state of the CMK changes to PendingDeletion.
// Before the waiting period ends, you can use CancelKeyDeletion to cancel the
// deletion of the CMK. After the waiting period ends, AWS KMS deletes the CMK
// and all AWS KMS data associated with it, including all aliases that refer
// to it.
//
// Deleting a CMK is a destructive and potentially dangerous operation. When
// a CMK is deleted, all data that was encrypted under the CMK is unrecoverable.
// To prevent the use of a CMK without deleting it, use DisableKey.
//
// If you schedule deletion of a CMK from a custom key store (http://docs.aws.amazon.com/kms/latest/developerguide/key-store-overview.html),
// when the waiting period expires, ScheduleKeyDeletion deletes the CMK from
// AWS KMS. Then AWS KMS makes a best effort to delete the key material from
// the associated AWS CloudHSM cluster. However, you might need to manually
// delete the orphaned key material (http://docs.aws.amazon.com/kms/latest/developerguide/fix-keystore.html#fix-keystore-orphaned-key)
// from the cluster and its backups.
//
// You cannot perform this operation on a CMK in a different AWS account.
//
// For more information about scheduling a CMK for deletion, see Deleting Customer
// Master Keys (http://docs.aws.amazon.com/kms/latest/developerguide/deleting-keys.html)
// in the AWS Key Management Service Developer Guide.
//
// The result of this operation varies with the key state of the CMK. For details,
// see How Key State Affects Use of a Customer Master Key (http://docs.aws.amazon.com/kms/latest/developerguide/key-state.html)
// in the AWS Key Management Service Developer Guide.
//
//    // Example sending a request using ScheduleKeyDeletionRequest.
//    req := client.ScheduleKeyDeletionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/kms-2014-11-01/ScheduleKeyDeletion
func (c *Client) ScheduleKeyDeletionRequest(input *ScheduleKeyDeletionInput) ScheduleKeyDeletionRequest {
	op := &aws.Operation{
		Name:       opScheduleKeyDeletion,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ScheduleKeyDeletionInput{}
	}

	req := c.newRequest(op, input, &ScheduleKeyDeletionOutput{})
	return ScheduleKeyDeletionRequest{Request: req, Input: input, Copy: c.ScheduleKeyDeletionRequest}
}

// ScheduleKeyDeletionRequest is the request type for the
// ScheduleKeyDeletion API operation.
type ScheduleKeyDeletionRequest struct {
	*aws.Request
	Input *ScheduleKeyDeletionInput
	Copy  func(*ScheduleKeyDeletionInput) ScheduleKeyDeletionRequest
}

// Send marshals and sends the ScheduleKeyDeletion API request.
func (r ScheduleKeyDeletionRequest) Send(ctx context.Context) (*ScheduleKeyDeletionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ScheduleKeyDeletionResponse{
		ScheduleKeyDeletionOutput: r.Request.Data.(*ScheduleKeyDeletionOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ScheduleKeyDeletionResponse is the response type for the
// ScheduleKeyDeletion API operation.
type ScheduleKeyDeletionResponse struct {
	*ScheduleKeyDeletionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ScheduleKeyDeletion request.
func (r *ScheduleKeyDeletionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
