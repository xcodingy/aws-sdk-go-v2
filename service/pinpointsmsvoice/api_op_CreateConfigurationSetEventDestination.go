// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package pinpointsmsvoice

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Create a new event destination in a configuration set.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-sms-voice-2018-09-05/CreateConfigurationSetEventDestinationRequest
type CreateConfigurationSetEventDestinationInput struct {
	_ struct{} `type:"structure"`

	// ConfigurationSetName is a required field
	ConfigurationSetName *string `location:"uri" locationName:"ConfigurationSetName" type:"string" required:"true"`

	// An object that defines a single event destination.
	EventDestination *EventDestinationDefinition `type:"structure"`

	// A name that identifies the event destination.
	EventDestinationName *string `type:"string"`
}

// String returns the string representation
func (s CreateConfigurationSetEventDestinationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateConfigurationSetEventDestinationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateConfigurationSetEventDestinationInput"}

	if s.ConfigurationSetName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ConfigurationSetName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateConfigurationSetEventDestinationInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	if s.EventDestination != nil {
		v := s.EventDestination

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "EventDestination", v, metadata)
	}
	if s.EventDestinationName != nil {
		v := *s.EventDestinationName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "EventDestinationName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ConfigurationSetName != nil {
		v := *s.ConfigurationSetName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "ConfigurationSetName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// An empty object that indicates that the event destination was created successfully.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-sms-voice-2018-09-05/CreateConfigurationSetEventDestinationResponse
type CreateConfigurationSetEventDestinationOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s CreateConfigurationSetEventDestinationOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateConfigurationSetEventDestinationOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opCreateConfigurationSetEventDestination = "CreateConfigurationSetEventDestination"

// CreateConfigurationSetEventDestinationRequest returns a request value for making API operation for
// Amazon Pinpoint SMS and Voice Service.
//
// Create a new event destination in a configuration set.
//
//    // Example sending a request using CreateConfigurationSetEventDestinationRequest.
//    req := client.CreateConfigurationSetEventDestinationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-sms-voice-2018-09-05/CreateConfigurationSetEventDestination
func (c *Client) CreateConfigurationSetEventDestinationRequest(input *CreateConfigurationSetEventDestinationInput) CreateConfigurationSetEventDestinationRequest {
	op := &aws.Operation{
		Name:       opCreateConfigurationSetEventDestination,
		HTTPMethod: "POST",
		HTTPPath:   "/v1/sms-voice/configuration-sets/{ConfigurationSetName}/event-destinations",
	}

	if input == nil {
		input = &CreateConfigurationSetEventDestinationInput{}
	}

	req := c.newRequest(op, input, &CreateConfigurationSetEventDestinationOutput{})
	return CreateConfigurationSetEventDestinationRequest{Request: req, Input: input, Copy: c.CreateConfigurationSetEventDestinationRequest}
}

// CreateConfigurationSetEventDestinationRequest is the request type for the
// CreateConfigurationSetEventDestination API operation.
type CreateConfigurationSetEventDestinationRequest struct {
	*aws.Request
	Input *CreateConfigurationSetEventDestinationInput
	Copy  func(*CreateConfigurationSetEventDestinationInput) CreateConfigurationSetEventDestinationRequest
}

// Send marshals and sends the CreateConfigurationSetEventDestination API request.
func (r CreateConfigurationSetEventDestinationRequest) Send(ctx context.Context) (*CreateConfigurationSetEventDestinationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateConfigurationSetEventDestinationResponse{
		CreateConfigurationSetEventDestinationOutput: r.Request.Data.(*CreateConfigurationSetEventDestinationOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateConfigurationSetEventDestinationResponse is the response type for the
// CreateConfigurationSetEventDestination API operation.
type CreateConfigurationSetEventDestinationResponse struct {
	*CreateConfigurationSetEventDestinationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateConfigurationSetEventDestination request.
func (r *CreateConfigurationSetEventDestinationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
