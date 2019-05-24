// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package lambda

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/lambda-2015-03-31/UpdateEventSourceMappingRequest
type UpdateEventSourceMappingInput struct {
	_ struct{} `type:"structure"`

	// The maximum number of items to retrieve in a single batch.
	//
	//    * Amazon Kinesis - Default 100. Max 10,000.
	//
	//    * Amazon DynamoDB Streams - Default 100. Max 1,000.
	//
	//    * Amazon Simple Queue Service - Default 10. Max 10.
	BatchSize *int64 `min:"1" type:"integer"`

	// Disables the event source mapping to pause polling and invocation.
	Enabled *bool `type:"boolean"`

	// The name of the Lambda function.
	//
	// Name formats
	//
	//    * Function name - MyFunction.
	//
	//    * Function ARN - arn:aws:lambda:us-west-2:123456789012:function:MyFunction.
	//
	//    * Version or Alias ARN - arn:aws:lambda:us-west-2:123456789012:function:MyFunction:PROD.
	//
	//    * Partial ARN - 123456789012:function:MyFunction.
	//
	// The length constraint applies only to the full ARN. If you specify only the
	// function name, it's limited to 64 characters in length.
	FunctionName *string `min:"1" type:"string"`

	// The identifier of the event source mapping.
	//
	// UUID is a required field
	UUID *string `location:"uri" locationName:"UUID" type:"string" required:"true"`
}

// String returns the string representation
func (s UpdateEventSourceMappingInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateEventSourceMappingInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateEventSourceMappingInput"}
	if s.BatchSize != nil && *s.BatchSize < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("BatchSize", 1))
	}
	if s.FunctionName != nil && len(*s.FunctionName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("FunctionName", 1))
	}

	if s.UUID == nil {
		invalidParams.Add(aws.NewErrParamRequired("UUID"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateEventSourceMappingInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.BatchSize != nil {
		v := *s.BatchSize

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "BatchSize", protocol.Int64Value(v), metadata)
	}
	if s.Enabled != nil {
		v := *s.Enabled

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Enabled", protocol.BoolValue(v), metadata)
	}
	if s.FunctionName != nil {
		v := *s.FunctionName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "FunctionName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.UUID != nil {
		v := *s.UUID

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "UUID", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// A mapping between an AWS resource and an AWS Lambda function. See CreateEventSourceMapping
// for details.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/lambda-2015-03-31/EventSourceMappingConfiguration
type UpdateEventSourceMappingOutput struct {
	_ struct{} `type:"structure"`

	// The maximum number of items to retrieve in a single batch.
	BatchSize *int64 `min:"1" type:"integer"`

	// The Amazon Resource Name (ARN) of the event source.
	EventSourceArn *string `type:"string"`

	// The ARN of the Lambda function.
	FunctionArn *string `type:"string"`

	// The date that the event source mapping was last updated.
	LastModified *time.Time `type:"timestamp" timestampFormat:"unix"`

	// The result of the last AWS Lambda invocation of your Lambda function.
	LastProcessingResult *string `type:"string"`

	// The state of the event source mapping. It can be one of the following: Creating,
	// Enabling, Enabled, Disabling, Disabled, Updating, or Deleting.
	State *string `type:"string"`

	// The cause of the last state change, either User initiated or Lambda initiated.
	StateTransitionReason *string `type:"string"`

	// The identifier of the event source mapping.
	UUID *string `type:"string"`
}

// String returns the string representation
func (s UpdateEventSourceMappingOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateEventSourceMappingOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.BatchSize != nil {
		v := *s.BatchSize

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "BatchSize", protocol.Int64Value(v), metadata)
	}
	if s.EventSourceArn != nil {
		v := *s.EventSourceArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "EventSourceArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.FunctionArn != nil {
		v := *s.FunctionArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "FunctionArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.LastModified != nil {
		v := *s.LastModified

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "LastModified", protocol.TimeValue{V: v, Format: protocol.UnixTimeFormat}, metadata)
	}
	if s.LastProcessingResult != nil {
		v := *s.LastProcessingResult

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "LastProcessingResult", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.State != nil {
		v := *s.State

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "State", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.StateTransitionReason != nil {
		v := *s.StateTransitionReason

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "StateTransitionReason", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.UUID != nil {
		v := *s.UUID

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "UUID", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opUpdateEventSourceMapping = "UpdateEventSourceMapping"

// UpdateEventSourceMappingRequest returns a request value for making API operation for
// AWS Lambda.
//
// Updates an event source mapping. You can change the function that AWS Lambda
// invokes, or pause invocation and resume later from the same location.
//
//    // Example sending a request using UpdateEventSourceMappingRequest.
//    req := client.UpdateEventSourceMappingRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/lambda-2015-03-31/UpdateEventSourceMapping
func (c *Client) UpdateEventSourceMappingRequest(input *UpdateEventSourceMappingInput) UpdateEventSourceMappingRequest {
	op := &aws.Operation{
		Name:       opUpdateEventSourceMapping,
		HTTPMethod: "PUT",
		HTTPPath:   "/2015-03-31/event-source-mappings/{UUID}",
	}

	if input == nil {
		input = &UpdateEventSourceMappingInput{}
	}

	req := c.newRequest(op, input, &UpdateEventSourceMappingOutput{})
	return UpdateEventSourceMappingRequest{Request: req, Input: input, Copy: c.UpdateEventSourceMappingRequest}
}

// UpdateEventSourceMappingRequest is the request type for the
// UpdateEventSourceMapping API operation.
type UpdateEventSourceMappingRequest struct {
	*aws.Request
	Input *UpdateEventSourceMappingInput
	Copy  func(*UpdateEventSourceMappingInput) UpdateEventSourceMappingRequest
}

// Send marshals and sends the UpdateEventSourceMapping API request.
func (r UpdateEventSourceMappingRequest) Send(ctx context.Context) (*UpdateEventSourceMappingResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateEventSourceMappingResponse{
		UpdateEventSourceMappingOutput: r.Request.Data.(*UpdateEventSourceMappingOutput),
		response:                       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateEventSourceMappingResponse is the response type for the
// UpdateEventSourceMapping API operation.
type UpdateEventSourceMappingResponse struct {
	*UpdateEventSourceMappingOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateEventSourceMapping request.
func (r *UpdateEventSourceMappingResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
