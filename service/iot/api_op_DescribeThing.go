// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iot

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// The input for the DescribeThing operation.
type DescribeThingInput struct {
	_ struct{} `type:"structure"`

	// The name of the thing.
	//
	// ThingName is a required field
	ThingName *string `location:"uri" locationName:"thingName" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeThingInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeThingInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeThingInput"}

	if s.ThingName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ThingName"))
	}
	if s.ThingName != nil && len(*s.ThingName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ThingName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeThingInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.ThingName != nil {
		v := *s.ThingName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "thingName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// The output from the DescribeThing operation.
type DescribeThingOutput struct {
	_ struct{} `type:"structure"`

	// The thing attributes.
	Attributes map[string]string `locationName:"attributes" type:"map"`

	// The name of the billing group the thing belongs to.
	BillingGroupName *string `locationName:"billingGroupName" min:"1" type:"string"`

	// The default client ID.
	DefaultClientId *string `locationName:"defaultClientId" type:"string"`

	// The ARN of the thing to describe.
	ThingArn *string `locationName:"thingArn" type:"string"`

	// The ID of the thing to describe.
	ThingId *string `locationName:"thingId" type:"string"`

	// The name of the thing.
	ThingName *string `locationName:"thingName" min:"1" type:"string"`

	// The thing type name.
	ThingTypeName *string `locationName:"thingTypeName" min:"1" type:"string"`

	// The current version of the thing record in the registry.
	//
	// To avoid unintentional changes to the information in the registry, you can
	// pass the version information in the expectedVersion parameter of the UpdateThing
	// and DeleteThing calls.
	Version *int64 `locationName:"version" type:"long"`
}

// String returns the string representation
func (s DescribeThingOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeThingOutput) MarshalFields(e protocol.FieldEncoder) error {
	if len(s.Attributes) > 0 {
		v := s.Attributes

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "attributes", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetValue(k1, protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ms0.End()

	}
	if s.BillingGroupName != nil {
		v := *s.BillingGroupName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "billingGroupName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.DefaultClientId != nil {
		v := *s.DefaultClientId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "defaultClientId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ThingArn != nil {
		v := *s.ThingArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "thingArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ThingId != nil {
		v := *s.ThingId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "thingId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ThingName != nil {
		v := *s.ThingName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "thingName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ThingTypeName != nil {
		v := *s.ThingTypeName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "thingTypeName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Version != nil {
		v := *s.Version

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "version", protocol.Int64Value(v), metadata)
	}
	return nil
}

const opDescribeThing = "DescribeThing"

// DescribeThingRequest returns a request value for making API operation for
// AWS IoT.
//
// Gets information about the specified thing.
//
//    // Example sending a request using DescribeThingRequest.
//    req := client.DescribeThingRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) DescribeThingRequest(input *DescribeThingInput) DescribeThingRequest {
	op := &aws.Operation{
		Name:       opDescribeThing,
		HTTPMethod: "GET",
		HTTPPath:   "/things/{thingName}",
	}

	if input == nil {
		input = &DescribeThingInput{}
	}

	req := c.newRequest(op, input, &DescribeThingOutput{})
	return DescribeThingRequest{Request: req, Input: input, Copy: c.DescribeThingRequest}
}

// DescribeThingRequest is the request type for the
// DescribeThing API operation.
type DescribeThingRequest struct {
	*aws.Request
	Input *DescribeThingInput
	Copy  func(*DescribeThingInput) DescribeThingRequest
}

// Send marshals and sends the DescribeThing API request.
func (r DescribeThingRequest) Send(ctx context.Context) (*DescribeThingResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeThingResponse{
		DescribeThingOutput: r.Request.Data.(*DescribeThingOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeThingResponse is the response type for the
// DescribeThing API operation.
type DescribeThingResponse struct {
	*DescribeThingOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeThing request.
func (r *DescribeThingResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
