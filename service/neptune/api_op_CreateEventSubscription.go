// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package neptune

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/neptune-2014-10-31/CreateEventSubscriptionMessage
type CreateEventSubscriptionInput struct {
	_ struct{} `type:"structure"`

	// A Boolean value; set to true to activate the subscription, set to false to
	// create the subscription but not active it.
	Enabled *bool `type:"boolean"`

	// A list of event categories for a SourceType that you want to subscribe to.
	// You can see a list of the categories for a given SourceType by using the
	// DescribeEventCategories action.
	EventCategories []string `locationNameList:"EventCategory" type:"list"`

	// The Amazon Resource Name (ARN) of the SNS topic created for event notification.
	// The ARN is created by Amazon SNS when you create a topic and subscribe to
	// it.
	//
	// SnsTopicArn is a required field
	SnsTopicArn *string `type:"string" required:"true"`

	// The list of identifiers of the event sources for which events are returned.
	// If not specified, then all sources are included in the response. An identifier
	// must begin with a letter and must contain only ASCII letters, digits, and
	// hyphens; it can't end with a hyphen or contain two consecutive hyphens.
	//
	// Constraints:
	//
	//    * If SourceIds are supplied, SourceType must also be provided.
	//
	//    * If the source type is a DB instance, then a DBInstanceIdentifier must
	//    be supplied.
	//
	//    * If the source type is a DB security group, a DBSecurityGroupName must
	//    be supplied.
	//
	//    * If the source type is a DB parameter group, a DBParameterGroupName must
	//    be supplied.
	//
	//    * If the source type is a DB snapshot, a DBSnapshotIdentifier must be
	//    supplied.
	SourceIds []string `locationNameList:"SourceId" type:"list"`

	// The type of source that is generating the events. For example, if you want
	// to be notified of events generated by a DB instance, you would set this parameter
	// to db-instance. if this value is not specified, all events are returned.
	//
	// Valid values: db-instance | db-cluster | db-parameter-group | db-security-group
	// | db-snapshot | db-cluster-snapshot
	SourceType *string `type:"string"`

	// The name of the subscription.
	//
	// Constraints: The name must be less than 255 characters.
	//
	// SubscriptionName is a required field
	SubscriptionName *string `type:"string" required:"true"`

	// A list of tags. For more information, see Tagging Amazon Neptune Resources
	// (http://docs.aws.amazon.com/neptune/latest/UserGuide/tagging.ARN.html).
	Tags []Tag `locationNameList:"Tag" type:"list"`
}

// String returns the string representation
func (s CreateEventSubscriptionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateEventSubscriptionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateEventSubscriptionInput"}

	if s.SnsTopicArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("SnsTopicArn"))
	}

	if s.SubscriptionName == nil {
		invalidParams.Add(aws.NewErrParamRequired("SubscriptionName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/neptune-2014-10-31/CreateEventSubscriptionResult
type CreateEventSubscriptionOutput struct {
	_ struct{} `type:"structure"`

	// Contains the results of a successful invocation of the DescribeEventSubscriptions
	// action.
	EventSubscription *EventSubscription `type:"structure"`
}

// String returns the string representation
func (s CreateEventSubscriptionOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateEventSubscription = "CreateEventSubscription"

// CreateEventSubscriptionRequest returns a request value for making API operation for
// Amazon Neptune.
//
// Creates an event notification subscription. This action requires a topic
// ARN (Amazon Resource Name) created by either the Neptune console, the SNS
// console, or the SNS API. To obtain an ARN with SNS, you must create a topic
// in Amazon SNS and subscribe to the topic. The ARN is displayed in the SNS
// console.
//
// You can specify the type of source (SourceType) you want to be notified of,
// provide a list of Neptune sources (SourceIds) that triggers the events, and
// provide a list of event categories (EventCategories) for events you want
// to be notified of. For example, you can specify SourceType = db-instance,
// SourceIds = mydbinstance1, mydbinstance2 and EventCategories = Availability,
// Backup.
//
// If you specify both the SourceType and SourceIds, such as SourceType = db-instance
// and SourceIdentifier = myDBInstance1, you are notified of all the db-instance
// events for the specified source. If you specify a SourceType but do not specify
// a SourceIdentifier, you receive notice of the events for that source type
// for all your Neptune sources. If you do not specify either the SourceType
// nor the SourceIdentifier, you are notified of events generated from all Neptune
// sources belonging to your customer account.
//
//    // Example sending a request using CreateEventSubscriptionRequest.
//    req := client.CreateEventSubscriptionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/neptune-2014-10-31/CreateEventSubscription
func (c *Client) CreateEventSubscriptionRequest(input *CreateEventSubscriptionInput) CreateEventSubscriptionRequest {
	op := &aws.Operation{
		Name:       opCreateEventSubscription,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateEventSubscriptionInput{}
	}

	req := c.newRequest(op, input, &CreateEventSubscriptionOutput{})
	return CreateEventSubscriptionRequest{Request: req, Input: input, Copy: c.CreateEventSubscriptionRequest}
}

// CreateEventSubscriptionRequest is the request type for the
// CreateEventSubscription API operation.
type CreateEventSubscriptionRequest struct {
	*aws.Request
	Input *CreateEventSubscriptionInput
	Copy  func(*CreateEventSubscriptionInput) CreateEventSubscriptionRequest
}

// Send marshals and sends the CreateEventSubscription API request.
func (r CreateEventSubscriptionRequest) Send(ctx context.Context) (*CreateEventSubscriptionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateEventSubscriptionResponse{
		CreateEventSubscriptionOutput: r.Request.Data.(*CreateEventSubscriptionOutput),
		response:                      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateEventSubscriptionResponse is the response type for the
// CreateEventSubscription API operation.
type CreateEventSubscriptionResponse struct {
	*CreateEventSubscriptionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateEventSubscription request.
func (r *CreateEventSubscriptionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
