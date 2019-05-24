// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package guardduty

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/guardduty-2017-11-28/GetDetectorRequest
type GetDetectorInput struct {
	_ struct{} `type:"structure"`

	// DetectorId is a required field
	DetectorId *string `location:"uri" locationName:"detectorId" type:"string" required:"true"`
}

// String returns the string representation
func (s GetDetectorInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetDetectorInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetDetectorInput"}

	if s.DetectorId == nil {
		invalidParams.Add(aws.NewErrParamRequired("DetectorId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetDetectorInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	if s.DetectorId != nil {
		v := *s.DetectorId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "detectorId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// GetDetector response object.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/guardduty-2017-11-28/GetDetectorResponse
type GetDetectorOutput struct {
	_ struct{} `type:"structure"`

	// The first time a resource was created. The format will be ISO-8601.
	CreatedAt *string `locationName:"createdAt" type:"string"`

	// A enum value that specifies how frequently customer got Finding updates published.
	FindingPublishingFrequency FindingPublishingFrequency `locationName:"findingPublishingFrequency" type:"string" enum:"true"`

	// Customer serviceRole name or ARN for accessing customer resources
	ServiceRole *string `locationName:"serviceRole" type:"string"`

	// The status of detector.
	Status DetectorStatus `locationName:"status" type:"string" enum:"true"`

	// The first time a resource was created. The format will be ISO-8601.
	UpdatedAt *string `locationName:"updatedAt" type:"string"`
}

// String returns the string representation
func (s GetDetectorOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetDetectorOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.CreatedAt != nil {
		v := *s.CreatedAt

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "createdAt", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.FindingPublishingFrequency) > 0 {
		v := s.FindingPublishingFrequency

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "findingPublishingFrequency", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.ServiceRole != nil {
		v := *s.ServiceRole

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "serviceRole", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.Status) > 0 {
		v := s.Status

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "status", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.UpdatedAt != nil {
		v := *s.UpdatedAt

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "updatedAt", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opGetDetector = "GetDetector"

// GetDetectorRequest returns a request value for making API operation for
// Amazon GuardDuty.
//
// Retrieves an Amazon GuardDuty detector specified by the detectorId.
//
//    // Example sending a request using GetDetectorRequest.
//    req := client.GetDetectorRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/guardduty-2017-11-28/GetDetector
func (c *Client) GetDetectorRequest(input *GetDetectorInput) GetDetectorRequest {
	op := &aws.Operation{
		Name:       opGetDetector,
		HTTPMethod: "GET",
		HTTPPath:   "/detector/{detectorId}",
	}

	if input == nil {
		input = &GetDetectorInput{}
	}

	req := c.newRequest(op, input, &GetDetectorOutput{})
	return GetDetectorRequest{Request: req, Input: input, Copy: c.GetDetectorRequest}
}

// GetDetectorRequest is the request type for the
// GetDetector API operation.
type GetDetectorRequest struct {
	*aws.Request
	Input *GetDetectorInput
	Copy  func(*GetDetectorInput) GetDetectorRequest
}

// Send marshals and sends the GetDetector API request.
func (r GetDetectorRequest) Send(ctx context.Context) (*GetDetectorResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetDetectorResponse{
		GetDetectorOutput: r.Request.Data.(*GetDetectorOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetDetectorResponse is the response type for the
// GetDetector API operation.
type GetDetectorResponse struct {
	*GetDetectorOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetDetector request.
func (r *GetDetectorResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
