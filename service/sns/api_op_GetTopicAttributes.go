// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sns

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Input for GetTopicAttributes action.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/sns-2010-03-31/GetTopicAttributesInput
type GetTopicAttributesInput struct {
	_ struct{} `type:"structure"`

	// The ARN of the topic whose properties you want to get.
	//
	// TopicArn is a required field
	TopicArn *string `type:"string" required:"true"`
}

// String returns the string representation
func (s GetTopicAttributesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetTopicAttributesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetTopicAttributesInput"}

	if s.TopicArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("TopicArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Response for GetTopicAttributes action.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/sns-2010-03-31/GetTopicAttributesResponse
type GetTopicAttributesOutput struct {
	_ struct{} `type:"structure"`

	// A map of the topic's attributes. Attributes in this map include the following:
	//
	//    * TopicArn – the topic's ARN
	//
	//    * Owner – the AWS account ID of the topic's owner
	//
	//    * Policy – the JSON serialization of the topic's access control policy
	//
	//    * DisplayName – the human-readable name used in the "From" field for notifications
	//    to email and email-json endpoints
	//
	//    * SubscriptionsPending – the number of subscriptions pending confirmation
	//    on this topic
	//
	//    * SubscriptionsConfirmed – the number of confirmed subscriptions on this
	//    topic
	//
	//    * SubscriptionsDeleted – the number of deleted subscriptions on this topic
	//
	//    * DeliveryPolicy – the JSON serialization of the topic's delivery policy
	//
	//    * EffectiveDeliveryPolicy – the JSON serialization of the effective delivery
	//    policy that takes into account system defaults
	Attributes map[string]string `type:"map"`
}

// String returns the string representation
func (s GetTopicAttributesOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetTopicAttributes = "GetTopicAttributes"

// GetTopicAttributesRequest returns a request value for making API operation for
// Amazon Simple Notification Service.
//
// Returns all of the properties of a topic. Topic properties returned might
// differ based on the authorization of the user.
//
//    // Example sending a request using GetTopicAttributesRequest.
//    req := client.GetTopicAttributesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/sns-2010-03-31/GetTopicAttributes
func (c *Client) GetTopicAttributesRequest(input *GetTopicAttributesInput) GetTopicAttributesRequest {
	op := &aws.Operation{
		Name:       opGetTopicAttributes,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetTopicAttributesInput{}
	}

	req := c.newRequest(op, input, &GetTopicAttributesOutput{})
	return GetTopicAttributesRequest{Request: req, Input: input, Copy: c.GetTopicAttributesRequest}
}

// GetTopicAttributesRequest is the request type for the
// GetTopicAttributes API operation.
type GetTopicAttributesRequest struct {
	*aws.Request
	Input *GetTopicAttributesInput
	Copy  func(*GetTopicAttributesInput) GetTopicAttributesRequest
}

// Send marshals and sends the GetTopicAttributes API request.
func (r GetTopicAttributesRequest) Send(ctx context.Context) (*GetTopicAttributesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetTopicAttributesResponse{
		GetTopicAttributesOutput: r.Request.Data.(*GetTopicAttributesOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetTopicAttributesResponse is the response type for the
// GetTopicAttributes API operation.
type GetTopicAttributesResponse struct {
	*GetTopicAttributesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetTopicAttributes request.
func (r *GetTopicAttributesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
