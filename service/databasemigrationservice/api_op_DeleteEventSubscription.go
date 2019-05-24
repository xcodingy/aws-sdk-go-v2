// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package databasemigrationservice

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/dms-2016-01-01/DeleteEventSubscriptionMessage
type DeleteEventSubscriptionInput struct {
	_ struct{} `type:"structure"`

	// The name of the DMS event notification subscription to be deleted.
	//
	// SubscriptionName is a required field
	SubscriptionName *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteEventSubscriptionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteEventSubscriptionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteEventSubscriptionInput"}

	if s.SubscriptionName == nil {
		invalidParams.Add(aws.NewErrParamRequired("SubscriptionName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/dms-2016-01-01/DeleteEventSubscriptionResponse
type DeleteEventSubscriptionOutput struct {
	_ struct{} `type:"structure"`

	// The event subscription that was deleted.
	EventSubscription *EventSubscription `type:"structure"`
}

// String returns the string representation
func (s DeleteEventSubscriptionOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteEventSubscription = "DeleteEventSubscription"

// DeleteEventSubscriptionRequest returns a request value for making API operation for
// AWS Database Migration Service.
//
// Deletes an AWS DMS event subscription.
//
//    // Example sending a request using DeleteEventSubscriptionRequest.
//    req := client.DeleteEventSubscriptionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/dms-2016-01-01/DeleteEventSubscription
func (c *Client) DeleteEventSubscriptionRequest(input *DeleteEventSubscriptionInput) DeleteEventSubscriptionRequest {
	op := &aws.Operation{
		Name:       opDeleteEventSubscription,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteEventSubscriptionInput{}
	}

	req := c.newRequest(op, input, &DeleteEventSubscriptionOutput{})
	return DeleteEventSubscriptionRequest{Request: req, Input: input, Copy: c.DeleteEventSubscriptionRequest}
}

// DeleteEventSubscriptionRequest is the request type for the
// DeleteEventSubscription API operation.
type DeleteEventSubscriptionRequest struct {
	*aws.Request
	Input *DeleteEventSubscriptionInput
	Copy  func(*DeleteEventSubscriptionInput) DeleteEventSubscriptionRequest
}

// Send marshals and sends the DeleteEventSubscription API request.
func (r DeleteEventSubscriptionRequest) Send(ctx context.Context) (*DeleteEventSubscriptionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteEventSubscriptionResponse{
		DeleteEventSubscriptionOutput: r.Request.Data.(*DeleteEventSubscriptionOutput),
		response:                      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteEventSubscriptionResponse is the response type for the
// DeleteEventSubscription API operation.
type DeleteEventSubscriptionResponse struct {
	*DeleteEventSubscriptionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteEventSubscription request.
func (r *DeleteEventSubscriptionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
