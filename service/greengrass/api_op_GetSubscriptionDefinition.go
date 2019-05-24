// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package greengrass

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/greengrass-2017-06-07/GetSubscriptionDefinitionRequest
type GetSubscriptionDefinitionInput struct {
	_ struct{} `type:"structure"`

	// SubscriptionDefinitionId is a required field
	SubscriptionDefinitionId *string `location:"uri" locationName:"SubscriptionDefinitionId" type:"string" required:"true"`
}

// String returns the string representation
func (s GetSubscriptionDefinitionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetSubscriptionDefinitionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetSubscriptionDefinitionInput"}

	if s.SubscriptionDefinitionId == nil {
		invalidParams.Add(aws.NewErrParamRequired("SubscriptionDefinitionId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetSubscriptionDefinitionInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	if s.SubscriptionDefinitionId != nil {
		v := *s.SubscriptionDefinitionId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "SubscriptionDefinitionId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/greengrass-2017-06-07/GetSubscriptionDefinitionResponse
type GetSubscriptionDefinitionOutput struct {
	_ struct{} `type:"structure"`

	Arn *string `type:"string"`

	CreationTimestamp *string `type:"string"`

	Id *string `type:"string"`

	LastUpdatedTimestamp *string `type:"string"`

	LatestVersion *string `type:"string"`

	LatestVersionArn *string `type:"string"`

	Name *string `type:"string"`

	Tags map[string]string `locationName:"tags" type:"map"`
}

// String returns the string representation
func (s GetSubscriptionDefinitionOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetSubscriptionDefinitionOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Arn != nil {
		v := *s.Arn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Arn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.CreationTimestamp != nil {
		v := *s.CreationTimestamp

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "CreationTimestamp", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Id != nil {
		v := *s.Id

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.LastUpdatedTimestamp != nil {
		v := *s.LastUpdatedTimestamp

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "LastUpdatedTimestamp", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.LatestVersion != nil {
		v := *s.LatestVersion

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "LatestVersion", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.LatestVersionArn != nil {
		v := *s.LatestVersionArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "LatestVersionArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.Tags) > 0 {
		v := s.Tags

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "tags", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetValue(k1, protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ms0.End()

	}
	return nil
}

const opGetSubscriptionDefinition = "GetSubscriptionDefinition"

// GetSubscriptionDefinitionRequest returns a request value for making API operation for
// AWS Greengrass.
//
// Retrieves information about a subscription definition.
//
//    // Example sending a request using GetSubscriptionDefinitionRequest.
//    req := client.GetSubscriptionDefinitionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/greengrass-2017-06-07/GetSubscriptionDefinition
func (c *Client) GetSubscriptionDefinitionRequest(input *GetSubscriptionDefinitionInput) GetSubscriptionDefinitionRequest {
	op := &aws.Operation{
		Name:       opGetSubscriptionDefinition,
		HTTPMethod: "GET",
		HTTPPath:   "/greengrass/definition/subscriptions/{SubscriptionDefinitionId}",
	}

	if input == nil {
		input = &GetSubscriptionDefinitionInput{}
	}

	req := c.newRequest(op, input, &GetSubscriptionDefinitionOutput{})
	return GetSubscriptionDefinitionRequest{Request: req, Input: input, Copy: c.GetSubscriptionDefinitionRequest}
}

// GetSubscriptionDefinitionRequest is the request type for the
// GetSubscriptionDefinition API operation.
type GetSubscriptionDefinitionRequest struct {
	*aws.Request
	Input *GetSubscriptionDefinitionInput
	Copy  func(*GetSubscriptionDefinitionInput) GetSubscriptionDefinitionRequest
}

// Send marshals and sends the GetSubscriptionDefinition API request.
func (r GetSubscriptionDefinitionRequest) Send(ctx context.Context) (*GetSubscriptionDefinitionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetSubscriptionDefinitionResponse{
		GetSubscriptionDefinitionOutput: r.Request.Data.(*GetSubscriptionDefinitionOutput),
		response:                        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetSubscriptionDefinitionResponse is the response type for the
// GetSubscriptionDefinition API operation.
type GetSubscriptionDefinitionResponse struct {
	*GetSubscriptionDefinitionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetSubscriptionDefinition request.
func (r *GetSubscriptionDefinitionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
