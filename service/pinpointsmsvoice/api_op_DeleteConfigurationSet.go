// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package pinpointsmsvoice

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-sms-voice-2018-09-05/DeleteConfigurationSetRequest
type DeleteConfigurationSetInput struct {
	_ struct{} `type:"structure"`

	// ConfigurationSetName is a required field
	ConfigurationSetName *string `location:"uri" locationName:"ConfigurationSetName" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteConfigurationSetInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteConfigurationSetInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteConfigurationSetInput"}

	if s.ConfigurationSetName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ConfigurationSetName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteConfigurationSetInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	if s.ConfigurationSetName != nil {
		v := *s.ConfigurationSetName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "ConfigurationSetName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// An empty object that indicates that the configuration set was deleted successfully.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-sms-voice-2018-09-05/DeleteConfigurationSetResponse
type DeleteConfigurationSetOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteConfigurationSetOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteConfigurationSetOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opDeleteConfigurationSet = "DeleteConfigurationSet"

// DeleteConfigurationSetRequest returns a request value for making API operation for
// Amazon Pinpoint SMS and Voice Service.
//
// Deletes an existing configuration set.
//
//    // Example sending a request using DeleteConfigurationSetRequest.
//    req := client.DeleteConfigurationSetRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-sms-voice-2018-09-05/DeleteConfigurationSet
func (c *Client) DeleteConfigurationSetRequest(input *DeleteConfigurationSetInput) DeleteConfigurationSetRequest {
	op := &aws.Operation{
		Name:       opDeleteConfigurationSet,
		HTTPMethod: "DELETE",
		HTTPPath:   "/v1/sms-voice/configuration-sets/{ConfigurationSetName}",
	}

	if input == nil {
		input = &DeleteConfigurationSetInput{}
	}

	req := c.newRequest(op, input, &DeleteConfigurationSetOutput{})
	return DeleteConfigurationSetRequest{Request: req, Input: input, Copy: c.DeleteConfigurationSetRequest}
}

// DeleteConfigurationSetRequest is the request type for the
// DeleteConfigurationSet API operation.
type DeleteConfigurationSetRequest struct {
	*aws.Request
	Input *DeleteConfigurationSetInput
	Copy  func(*DeleteConfigurationSetInput) DeleteConfigurationSetRequest
}

// Send marshals and sends the DeleteConfigurationSet API request.
func (r DeleteConfigurationSetRequest) Send(ctx context.Context) (*DeleteConfigurationSetResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteConfigurationSetResponse{
		DeleteConfigurationSetOutput: r.Request.Data.(*DeleteConfigurationSetOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteConfigurationSetResponse is the response type for the
// DeleteConfigurationSet API operation.
type DeleteConfigurationSetResponse struct {
	*DeleteConfigurationSetOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteConfigurationSet request.
func (r *DeleteConfigurationSetResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
