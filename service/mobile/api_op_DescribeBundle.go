// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package mobile

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Request structure to request the details of a specific bundle.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/mobile-2017-07-01/DescribeBundleRequest
type DescribeBundleInput struct {
	_ struct{} `type:"structure"`

	// Unique bundle identifier.
	//
	// BundleId is a required field
	BundleId *string `location:"uri" locationName:"bundleId" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeBundleInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeBundleInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeBundleInput"}

	if s.BundleId == nil {
		invalidParams.Add(aws.NewErrParamRequired("BundleId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeBundleInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	if s.BundleId != nil {
		v := *s.BundleId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "bundleId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Result structure contains the details of the bundle.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/mobile-2017-07-01/DescribeBundleResult
type DescribeBundleOutput struct {
	_ struct{} `type:"structure"`

	// The details of the bundle.
	Details *BundleDetails `locationName:"details" type:"structure"`
}

// String returns the string representation
func (s DescribeBundleOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeBundleOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Details != nil {
		v := s.Details

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "details", v, metadata)
	}
	return nil
}

const opDescribeBundle = "DescribeBundle"

// DescribeBundleRequest returns a request value for making API operation for
// AWS Mobile.
//
// Get the bundle details for the requested bundle id.
//
//    // Example sending a request using DescribeBundleRequest.
//    req := client.DescribeBundleRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/mobile-2017-07-01/DescribeBundle
func (c *Client) DescribeBundleRequest(input *DescribeBundleInput) DescribeBundleRequest {
	op := &aws.Operation{
		Name:       opDescribeBundle,
		HTTPMethod: "GET",
		HTTPPath:   "/bundles/{bundleId}",
	}

	if input == nil {
		input = &DescribeBundleInput{}
	}

	req := c.newRequest(op, input, &DescribeBundleOutput{})
	return DescribeBundleRequest{Request: req, Input: input, Copy: c.DescribeBundleRequest}
}

// DescribeBundleRequest is the request type for the
// DescribeBundle API operation.
type DescribeBundleRequest struct {
	*aws.Request
	Input *DescribeBundleInput
	Copy  func(*DescribeBundleInput) DescribeBundleRequest
}

// Send marshals and sends the DescribeBundle API request.
func (r DescribeBundleRequest) Send(ctx context.Context) (*DescribeBundleResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeBundleResponse{
		DescribeBundleOutput: r.Request.Data.(*DescribeBundleOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeBundleResponse is the response type for the
// DescribeBundle API operation.
type DescribeBundleResponse struct {
	*DescribeBundleOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeBundle request.
func (r *DescribeBundleResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
