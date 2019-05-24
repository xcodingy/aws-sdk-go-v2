// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package pinpointemail

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// A request to retrieve a list of the blacklists that your dedicated IP addresses
// appear on.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-email-2018-07-26/GetBlacklistReportsRequest
type GetBlacklistReportsInput struct {
	_ struct{} `type:"structure"`

	// A list of IP addresses that you want to retrieve blacklist information about.
	// You can only specify the dedicated IP addresses that you use to send email
	// using Amazon Pinpoint or Amazon SES.
	//
	// BlacklistItemNames is a required field
	BlacklistItemNames []string `type:"list" required:"true"`
}

// String returns the string representation
func (s GetBlacklistReportsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetBlacklistReportsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetBlacklistReportsInput"}

	if s.BlacklistItemNames == nil {
		invalidParams.Add(aws.NewErrParamRequired("BlacklistItemNames"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetBlacklistReportsInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	if len(s.BlacklistItemNames) > 0 {
		v := s.BlacklistItemNames

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "BlacklistItemNames", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	return nil
}

// An object that contains information about blacklist events.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-email-2018-07-26/GetBlacklistReportsResponse
type GetBlacklistReportsOutput struct {
	_ struct{} `type:"structure"`

	// An object that contains information about a blacklist that one of your dedicated
	// IP addresses appears on.
	//
	// BlacklistReport is a required field
	BlacklistReport map[string][]BlacklistEntry `type:"map" required:"true"`
}

// String returns the string representation
func (s GetBlacklistReportsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetBlacklistReportsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if len(s.BlacklistReport) > 0 {
		v := s.BlacklistReport

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "BlacklistReport", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ls1 := ms0.List(k1)
			ls1.Start()
			for _, v2 := range v1 {
				ls1.ListAddFields(v2)
			}
			ls1.End()
		}
		ms0.End()

	}
	return nil
}

const opGetBlacklistReports = "GetBlacklistReports"

// GetBlacklistReportsRequest returns a request value for making API operation for
// Amazon Pinpoint Email Service.
//
// Retrieve a list of the blacklists that your dedicated IP addresses appear
// on.
//
//    // Example sending a request using GetBlacklistReportsRequest.
//    req := client.GetBlacklistReportsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-email-2018-07-26/GetBlacklistReports
func (c *Client) GetBlacklistReportsRequest(input *GetBlacklistReportsInput) GetBlacklistReportsRequest {
	op := &aws.Operation{
		Name:       opGetBlacklistReports,
		HTTPMethod: "GET",
		HTTPPath:   "/v1/email/deliverability-dashboard/blacklist-report",
	}

	if input == nil {
		input = &GetBlacklistReportsInput{}
	}

	req := c.newRequest(op, input, &GetBlacklistReportsOutput{})
	return GetBlacklistReportsRequest{Request: req, Input: input, Copy: c.GetBlacklistReportsRequest}
}

// GetBlacklistReportsRequest is the request type for the
// GetBlacklistReports API operation.
type GetBlacklistReportsRequest struct {
	*aws.Request
	Input *GetBlacklistReportsInput
	Copy  func(*GetBlacklistReportsInput) GetBlacklistReportsRequest
}

// Send marshals and sends the GetBlacklistReports API request.
func (r GetBlacklistReportsRequest) Send(ctx context.Context) (*GetBlacklistReportsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetBlacklistReportsResponse{
		GetBlacklistReportsOutput: r.Request.Data.(*GetBlacklistReportsOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetBlacklistReportsResponse is the response type for the
// GetBlacklistReports API operation.
type GetBlacklistReportsResponse struct {
	*GetBlacklistReportsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetBlacklistReports request.
func (r *GetBlacklistReportsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
