// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package pinpointemail

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// A request to associate a configuration set with a dedicated IP pool.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-email-2018-07-26/PutConfigurationSetDeliveryOptionsRequest
type PutConfigurationSetDeliveryOptionsInput struct {
	_ struct{} `type:"structure"`

	// The name of the configuration set that you want to associate with a dedicated
	// IP pool.
	//
	// ConfigurationSetName is a required field
	ConfigurationSetName *string `location:"uri" locationName:"ConfigurationSetName" type:"string" required:"true"`

	// The name of the dedicated IP pool that you want to associate with the configuration
	// set.
	SendingPoolName *string `type:"string"`
}

// String returns the string representation
func (s PutConfigurationSetDeliveryOptionsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PutConfigurationSetDeliveryOptionsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "PutConfigurationSetDeliveryOptionsInput"}

	if s.ConfigurationSetName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ConfigurationSetName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s PutConfigurationSetDeliveryOptionsInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	if s.SendingPoolName != nil {
		v := *s.SendingPoolName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "SendingPoolName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ConfigurationSetName != nil {
		v := *s.ConfigurationSetName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "ConfigurationSetName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// An HTTP 200 response if the request succeeds, or an error message if the
// request fails.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-email-2018-07-26/PutConfigurationSetDeliveryOptionsResponse
type PutConfigurationSetDeliveryOptionsOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s PutConfigurationSetDeliveryOptionsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s PutConfigurationSetDeliveryOptionsOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opPutConfigurationSetDeliveryOptions = "PutConfigurationSetDeliveryOptions"

// PutConfigurationSetDeliveryOptionsRequest returns a request value for making API operation for
// Amazon Pinpoint Email Service.
//
// Associate a configuration set with a dedicated IP pool. You can use dedicated
// IP pools to create groups of dedicated IP addresses for sending specific
// types of email.
//
//    // Example sending a request using PutConfigurationSetDeliveryOptionsRequest.
//    req := client.PutConfigurationSetDeliveryOptionsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-email-2018-07-26/PutConfigurationSetDeliveryOptions
func (c *Client) PutConfigurationSetDeliveryOptionsRequest(input *PutConfigurationSetDeliveryOptionsInput) PutConfigurationSetDeliveryOptionsRequest {
	op := &aws.Operation{
		Name:       opPutConfigurationSetDeliveryOptions,
		HTTPMethod: "PUT",
		HTTPPath:   "/v1/email/configuration-sets/{ConfigurationSetName}/delivery-options",
	}

	if input == nil {
		input = &PutConfigurationSetDeliveryOptionsInput{}
	}

	req := c.newRequest(op, input, &PutConfigurationSetDeliveryOptionsOutput{})
	return PutConfigurationSetDeliveryOptionsRequest{Request: req, Input: input, Copy: c.PutConfigurationSetDeliveryOptionsRequest}
}

// PutConfigurationSetDeliveryOptionsRequest is the request type for the
// PutConfigurationSetDeliveryOptions API operation.
type PutConfigurationSetDeliveryOptionsRequest struct {
	*aws.Request
	Input *PutConfigurationSetDeliveryOptionsInput
	Copy  func(*PutConfigurationSetDeliveryOptionsInput) PutConfigurationSetDeliveryOptionsRequest
}

// Send marshals and sends the PutConfigurationSetDeliveryOptions API request.
func (r PutConfigurationSetDeliveryOptionsRequest) Send(ctx context.Context) (*PutConfigurationSetDeliveryOptionsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &PutConfigurationSetDeliveryOptionsResponse{
		PutConfigurationSetDeliveryOptionsOutput: r.Request.Data.(*PutConfigurationSetDeliveryOptionsOutput),
		response:                                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// PutConfigurationSetDeliveryOptionsResponse is the response type for the
// PutConfigurationSetDeliveryOptions API operation.
type PutConfigurationSetDeliveryOptionsResponse struct {
	*PutConfigurationSetDeliveryOptionsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// PutConfigurationSetDeliveryOptions request.
func (r *PutConfigurationSetDeliveryOptionsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
