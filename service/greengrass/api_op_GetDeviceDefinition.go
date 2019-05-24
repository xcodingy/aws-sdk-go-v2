// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package greengrass

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/greengrass-2017-06-07/GetDeviceDefinitionRequest
type GetDeviceDefinitionInput struct {
	_ struct{} `type:"structure"`

	// DeviceDefinitionId is a required field
	DeviceDefinitionId *string `location:"uri" locationName:"DeviceDefinitionId" type:"string" required:"true"`
}

// String returns the string representation
func (s GetDeviceDefinitionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetDeviceDefinitionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetDeviceDefinitionInput"}

	if s.DeviceDefinitionId == nil {
		invalidParams.Add(aws.NewErrParamRequired("DeviceDefinitionId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetDeviceDefinitionInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	if s.DeviceDefinitionId != nil {
		v := *s.DeviceDefinitionId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "DeviceDefinitionId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/greengrass-2017-06-07/GetDeviceDefinitionResponse
type GetDeviceDefinitionOutput struct {
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
func (s GetDeviceDefinitionOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetDeviceDefinitionOutput) MarshalFields(e protocol.FieldEncoder) error {
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

const opGetDeviceDefinition = "GetDeviceDefinition"

// GetDeviceDefinitionRequest returns a request value for making API operation for
// AWS Greengrass.
//
// Retrieves information about a device definition.
//
//    // Example sending a request using GetDeviceDefinitionRequest.
//    req := client.GetDeviceDefinitionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/greengrass-2017-06-07/GetDeviceDefinition
func (c *Client) GetDeviceDefinitionRequest(input *GetDeviceDefinitionInput) GetDeviceDefinitionRequest {
	op := &aws.Operation{
		Name:       opGetDeviceDefinition,
		HTTPMethod: "GET",
		HTTPPath:   "/greengrass/definition/devices/{DeviceDefinitionId}",
	}

	if input == nil {
		input = &GetDeviceDefinitionInput{}
	}

	req := c.newRequest(op, input, &GetDeviceDefinitionOutput{})
	return GetDeviceDefinitionRequest{Request: req, Input: input, Copy: c.GetDeviceDefinitionRequest}
}

// GetDeviceDefinitionRequest is the request type for the
// GetDeviceDefinition API operation.
type GetDeviceDefinitionRequest struct {
	*aws.Request
	Input *GetDeviceDefinitionInput
	Copy  func(*GetDeviceDefinitionInput) GetDeviceDefinitionRequest
}

// Send marshals and sends the GetDeviceDefinition API request.
func (r GetDeviceDefinitionRequest) Send(ctx context.Context) (*GetDeviceDefinitionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetDeviceDefinitionResponse{
		GetDeviceDefinitionOutput: r.Request.Data.(*GetDeviceDefinitionOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetDeviceDefinitionResponse is the response type for the
// GetDeviceDefinition API operation.
type GetDeviceDefinitionResponse struct {
	*GetDeviceDefinitionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetDeviceDefinition request.
func (r *GetDeviceDefinitionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
