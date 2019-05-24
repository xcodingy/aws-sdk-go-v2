// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ram

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ram-2018-01-04/EnableSharingWithAwsOrganizationRequest
type EnableSharingWithAwsOrganizationInput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s EnableSharingWithAwsOrganizationInput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s EnableSharingWithAwsOrganizationInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ram-2018-01-04/EnableSharingWithAwsOrganizationResponse
type EnableSharingWithAwsOrganizationOutput struct {
	_ struct{} `type:"structure"`

	// Indicates whether the request succeeded.
	ReturnValue *bool `locationName:"returnValue" type:"boolean"`
}

// String returns the string representation
func (s EnableSharingWithAwsOrganizationOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s EnableSharingWithAwsOrganizationOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.ReturnValue != nil {
		v := *s.ReturnValue

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "returnValue", protocol.BoolValue(v), metadata)
	}
	return nil
}

const opEnableSharingWithAwsOrganization = "EnableSharingWithAwsOrganization"

// EnableSharingWithAwsOrganizationRequest returns a request value for making API operation for
// AWS Resource Access Manager.
//
// Enables resource sharing within your organization.
//
//    // Example sending a request using EnableSharingWithAwsOrganizationRequest.
//    req := client.EnableSharingWithAwsOrganizationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ram-2018-01-04/EnableSharingWithAwsOrganization
func (c *Client) EnableSharingWithAwsOrganizationRequest(input *EnableSharingWithAwsOrganizationInput) EnableSharingWithAwsOrganizationRequest {
	op := &aws.Operation{
		Name:       opEnableSharingWithAwsOrganization,
		HTTPMethod: "POST",
		HTTPPath:   "/enablesharingwithawsorganization",
	}

	if input == nil {
		input = &EnableSharingWithAwsOrganizationInput{}
	}

	req := c.newRequest(op, input, &EnableSharingWithAwsOrganizationOutput{})
	return EnableSharingWithAwsOrganizationRequest{Request: req, Input: input, Copy: c.EnableSharingWithAwsOrganizationRequest}
}

// EnableSharingWithAwsOrganizationRequest is the request type for the
// EnableSharingWithAwsOrganization API operation.
type EnableSharingWithAwsOrganizationRequest struct {
	*aws.Request
	Input *EnableSharingWithAwsOrganizationInput
	Copy  func(*EnableSharingWithAwsOrganizationInput) EnableSharingWithAwsOrganizationRequest
}

// Send marshals and sends the EnableSharingWithAwsOrganization API request.
func (r EnableSharingWithAwsOrganizationRequest) Send(ctx context.Context) (*EnableSharingWithAwsOrganizationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &EnableSharingWithAwsOrganizationResponse{
		EnableSharingWithAwsOrganizationOutput: r.Request.Data.(*EnableSharingWithAwsOrganizationOutput),
		response:                               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// EnableSharingWithAwsOrganizationResponse is the response type for the
// EnableSharingWithAwsOrganization API operation.
type EnableSharingWithAwsOrganizationResponse struct {
	*EnableSharingWithAwsOrganizationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// EnableSharingWithAwsOrganization request.
func (r *EnableSharingWithAwsOrganizationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
