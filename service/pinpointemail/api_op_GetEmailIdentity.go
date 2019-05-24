// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package pinpointemail

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// A request to return details about an email identity.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-email-2018-07-26/GetEmailIdentityRequest
type GetEmailIdentityInput struct {
	_ struct{} `type:"structure"`

	// The email identity that you want to retrieve details for.
	//
	// EmailIdentity is a required field
	EmailIdentity *string `location:"uri" locationName:"EmailIdentity" type:"string" required:"true"`
}

// String returns the string representation
func (s GetEmailIdentityInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetEmailIdentityInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetEmailIdentityInput"}

	if s.EmailIdentity == nil {
		invalidParams.Add(aws.NewErrParamRequired("EmailIdentity"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetEmailIdentityInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	if s.EmailIdentity != nil {
		v := *s.EmailIdentity

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "EmailIdentity", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Details about an email identity.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-email-2018-07-26/GetEmailIdentityResponse
type GetEmailIdentityOutput struct {
	_ struct{} `type:"structure"`

	// An object that contains information about the DKIM attributes for the identity.
	// This object includes the tokens that you use to create the CNAME records
	// that are required to complete the DKIM verification process.
	DkimAttributes *DkimAttributes `type:"structure"`

	// The feedback forwarding configuration for the identity.
	//
	// If the value is true, Amazon Pinpoint sends you email notifications when
	// bounce or complaint events occur. Amazon Pinpoint sends this notification
	// to the address that you specified in the Return-Path header of the original
	// email.
	//
	// When you set this value to false, Amazon Pinpoint sends notifications through
	// other mechanisms, such as by notifying an Amazon SNS topic or another event
	// destination. You're required to have a method of tracking bounces and complaints.
	// If you haven't set up another mechanism for receiving bounce or complaint
	// notifications, Amazon Pinpoint sends an email notification when these events
	// occur (even if this setting is disabled).
	FeedbackForwardingStatus *bool `type:"boolean"`

	// The email identity type.
	IdentityType IdentityType `type:"string" enum:"true"`

	// An object that contains information about the Mail-From attributes for the
	// email identity.
	MailFromAttributes *MailFromAttributes `type:"structure"`

	// Specifies whether or not the identity is verified. In Amazon Pinpoint, you
	// can only send email from verified email addresses or domains. For more information
	// about verifying identities, see the Amazon Pinpoint User Guide (https://docs.aws.amazon.com/pinpoint/latest/userguide/channels-email-manage-verify.html).
	VerifiedForSendingStatus *bool `type:"boolean"`
}

// String returns the string representation
func (s GetEmailIdentityOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetEmailIdentityOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.DkimAttributes != nil {
		v := s.DkimAttributes

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "DkimAttributes", v, metadata)
	}
	if s.FeedbackForwardingStatus != nil {
		v := *s.FeedbackForwardingStatus

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "FeedbackForwardingStatus", protocol.BoolValue(v), metadata)
	}
	if len(s.IdentityType) > 0 {
		v := s.IdentityType

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "IdentityType", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.MailFromAttributes != nil {
		v := s.MailFromAttributes

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "MailFromAttributes", v, metadata)
	}
	if s.VerifiedForSendingStatus != nil {
		v := *s.VerifiedForSendingStatus

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "VerifiedForSendingStatus", protocol.BoolValue(v), metadata)
	}
	return nil
}

const opGetEmailIdentity = "GetEmailIdentity"

// GetEmailIdentityRequest returns a request value for making API operation for
// Amazon Pinpoint Email Service.
//
// Provides information about a specific identity associated with your Amazon
// Pinpoint account, including the identity's verification status, its DKIM
// authentication status, and its custom Mail-From settings.
//
//    // Example sending a request using GetEmailIdentityRequest.
//    req := client.GetEmailIdentityRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-email-2018-07-26/GetEmailIdentity
func (c *Client) GetEmailIdentityRequest(input *GetEmailIdentityInput) GetEmailIdentityRequest {
	op := &aws.Operation{
		Name:       opGetEmailIdentity,
		HTTPMethod: "GET",
		HTTPPath:   "/v1/email/identities/{EmailIdentity}",
	}

	if input == nil {
		input = &GetEmailIdentityInput{}
	}

	req := c.newRequest(op, input, &GetEmailIdentityOutput{})
	return GetEmailIdentityRequest{Request: req, Input: input, Copy: c.GetEmailIdentityRequest}
}

// GetEmailIdentityRequest is the request type for the
// GetEmailIdentity API operation.
type GetEmailIdentityRequest struct {
	*aws.Request
	Input *GetEmailIdentityInput
	Copy  func(*GetEmailIdentityInput) GetEmailIdentityRequest
}

// Send marshals and sends the GetEmailIdentity API request.
func (r GetEmailIdentityRequest) Send(ctx context.Context) (*GetEmailIdentityResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetEmailIdentityResponse{
		GetEmailIdentityOutput: r.Request.Data.(*GetEmailIdentityOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetEmailIdentityResponse is the response type for the
// GetEmailIdentity API operation.
type GetEmailIdentityResponse struct {
	*GetEmailIdentityOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetEmailIdentity request.
func (r *GetEmailIdentityResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
