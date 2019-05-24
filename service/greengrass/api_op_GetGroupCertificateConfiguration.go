// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package greengrass

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/greengrass-2017-06-07/GetGroupCertificateConfigurationRequest
type GetGroupCertificateConfigurationInput struct {
	_ struct{} `type:"structure"`

	// GroupId is a required field
	GroupId *string `location:"uri" locationName:"GroupId" type:"string" required:"true"`
}

// String returns the string representation
func (s GetGroupCertificateConfigurationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetGroupCertificateConfigurationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetGroupCertificateConfigurationInput"}

	if s.GroupId == nil {
		invalidParams.Add(aws.NewErrParamRequired("GroupId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetGroupCertificateConfigurationInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	if s.GroupId != nil {
		v := *s.GroupId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "GroupId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/greengrass-2017-06-07/GetGroupCertificateConfigurationResponse
type GetGroupCertificateConfigurationOutput struct {
	_ struct{} `type:"structure"`

	CertificateAuthorityExpiryInMilliseconds *string `type:"string"`

	CertificateExpiryInMilliseconds *string `type:"string"`

	GroupId *string `type:"string"`
}

// String returns the string representation
func (s GetGroupCertificateConfigurationOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetGroupCertificateConfigurationOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.CertificateAuthorityExpiryInMilliseconds != nil {
		v := *s.CertificateAuthorityExpiryInMilliseconds

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "CertificateAuthorityExpiryInMilliseconds", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.CertificateExpiryInMilliseconds != nil {
		v := *s.CertificateExpiryInMilliseconds

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "CertificateExpiryInMilliseconds", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.GroupId != nil {
		v := *s.GroupId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "GroupId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opGetGroupCertificateConfiguration = "GetGroupCertificateConfiguration"

// GetGroupCertificateConfigurationRequest returns a request value for making API operation for
// AWS Greengrass.
//
// Retrieves the current configuration for the CA used by the group.
//
//    // Example sending a request using GetGroupCertificateConfigurationRequest.
//    req := client.GetGroupCertificateConfigurationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/greengrass-2017-06-07/GetGroupCertificateConfiguration
func (c *Client) GetGroupCertificateConfigurationRequest(input *GetGroupCertificateConfigurationInput) GetGroupCertificateConfigurationRequest {
	op := &aws.Operation{
		Name:       opGetGroupCertificateConfiguration,
		HTTPMethod: "GET",
		HTTPPath:   "/greengrass/groups/{GroupId}/certificateauthorities/configuration/expiry",
	}

	if input == nil {
		input = &GetGroupCertificateConfigurationInput{}
	}

	req := c.newRequest(op, input, &GetGroupCertificateConfigurationOutput{})
	return GetGroupCertificateConfigurationRequest{Request: req, Input: input, Copy: c.GetGroupCertificateConfigurationRequest}
}

// GetGroupCertificateConfigurationRequest is the request type for the
// GetGroupCertificateConfiguration API operation.
type GetGroupCertificateConfigurationRequest struct {
	*aws.Request
	Input *GetGroupCertificateConfigurationInput
	Copy  func(*GetGroupCertificateConfigurationInput) GetGroupCertificateConfigurationRequest
}

// Send marshals and sends the GetGroupCertificateConfiguration API request.
func (r GetGroupCertificateConfigurationRequest) Send(ctx context.Context) (*GetGroupCertificateConfigurationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetGroupCertificateConfigurationResponse{
		GetGroupCertificateConfigurationOutput: r.Request.Data.(*GetGroupCertificateConfigurationOutput),
		response:                               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetGroupCertificateConfigurationResponse is the response type for the
// GetGroupCertificateConfiguration API operation.
type GetGroupCertificateConfigurationResponse struct {
	*GetGroupCertificateConfigurationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetGroupCertificateConfiguration request.
func (r *GetGroupCertificateConfigurationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
