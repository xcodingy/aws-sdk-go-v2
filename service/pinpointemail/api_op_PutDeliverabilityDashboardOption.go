// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package pinpointemail

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// A request to enable or disable the Deliverability dashboard. When you enable
// the Deliverability dashboard, you gain access to reputation metrics for the
// domains that you use to send email using Amazon Pinpoint. You also gain the
// ability to perform predictive inbox placement tests.
//
// When you use the Deliverability dashboard, you pay a monthly charge of USD$1,250.00,
// in addition to any other fees that you accrue by using Amazon Pinpoint. If
// you enable the Deliverability dashboard after the first day of a calendar
// month, we prorate the monthly charge based on how many days have elapsed
// in the current calendar month.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-email-2018-07-26/PutDeliverabilityDashboardOptionRequest
type PutDeliverabilityDashboardOptionInput struct {
	_ struct{} `type:"structure"`

	// Indicates whether the Deliverability dashboard is enabled. If the value is
	// true, then the dashboard is enabled.
	//
	// DashboardEnabled is a required field
	DashboardEnabled *bool `type:"boolean" required:"true"`
}

// String returns the string representation
func (s PutDeliverabilityDashboardOptionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PutDeliverabilityDashboardOptionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "PutDeliverabilityDashboardOptionInput"}

	if s.DashboardEnabled == nil {
		invalidParams.Add(aws.NewErrParamRequired("DashboardEnabled"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s PutDeliverabilityDashboardOptionInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	if s.DashboardEnabled != nil {
		v := *s.DashboardEnabled

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "DashboardEnabled", protocol.BoolValue(v), metadata)
	}
	return nil
}

// A response that indicates whether the Deliverability dashboard is enabled
// for your Amazon Pinpoint account.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-email-2018-07-26/PutDeliverabilityDashboardOptionResponse
type PutDeliverabilityDashboardOptionOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s PutDeliverabilityDashboardOptionOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s PutDeliverabilityDashboardOptionOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opPutDeliverabilityDashboardOption = "PutDeliverabilityDashboardOption"

// PutDeliverabilityDashboardOptionRequest returns a request value for making API operation for
// Amazon Pinpoint Email Service.
//
// Enable or disable the Deliverability dashboard. When you enable the Deliverability
// dashboard, you gain access to reputation metrics for the domains that you
// use to send email using Amazon Pinpoint. You also gain the ability to perform
// predictive inbox placement tests.
//
// When you use the Deliverability dashboard, you pay a monthly charge of USD$1,250.00,
// in addition to any other fees that you accrue by using Amazon Pinpoint. If
// you enable the Deliverability dashboard after the first day of a calendar
// month, we prorate the monthly charge based on how many days have elapsed
// in the current calendar month.
//
//    // Example sending a request using PutDeliverabilityDashboardOptionRequest.
//    req := client.PutDeliverabilityDashboardOptionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-email-2018-07-26/PutDeliverabilityDashboardOption
func (c *Client) PutDeliverabilityDashboardOptionRequest(input *PutDeliverabilityDashboardOptionInput) PutDeliverabilityDashboardOptionRequest {
	op := &aws.Operation{
		Name:       opPutDeliverabilityDashboardOption,
		HTTPMethod: "PUT",
		HTTPPath:   "/v1/email/deliverability-dashboard",
	}

	if input == nil {
		input = &PutDeliverabilityDashboardOptionInput{}
	}

	req := c.newRequest(op, input, &PutDeliverabilityDashboardOptionOutput{})
	return PutDeliverabilityDashboardOptionRequest{Request: req, Input: input, Copy: c.PutDeliverabilityDashboardOptionRequest}
}

// PutDeliverabilityDashboardOptionRequest is the request type for the
// PutDeliverabilityDashboardOption API operation.
type PutDeliverabilityDashboardOptionRequest struct {
	*aws.Request
	Input *PutDeliverabilityDashboardOptionInput
	Copy  func(*PutDeliverabilityDashboardOptionInput) PutDeliverabilityDashboardOptionRequest
}

// Send marshals and sends the PutDeliverabilityDashboardOption API request.
func (r PutDeliverabilityDashboardOptionRequest) Send(ctx context.Context) (*PutDeliverabilityDashboardOptionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &PutDeliverabilityDashboardOptionResponse{
		PutDeliverabilityDashboardOptionOutput: r.Request.Data.(*PutDeliverabilityDashboardOptionOutput),
		response:                               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// PutDeliverabilityDashboardOptionResponse is the response type for the
// PutDeliverabilityDashboardOption API operation.
type PutDeliverabilityDashboardOptionResponse struct {
	*PutDeliverabilityDashboardOptionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// PutDeliverabilityDashboardOption request.
func (r *PutDeliverabilityDashboardOptionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
