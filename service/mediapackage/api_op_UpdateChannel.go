// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package mediapackage

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/mediapackage-2017-10-12/UpdateChannelRequest
type UpdateChannelInput struct {
	_ struct{} `type:"structure"`

	Description *string `locationName:"description" type:"string"`

	// Id is a required field
	Id *string `location:"uri" locationName:"id" type:"string" required:"true"`
}

// String returns the string representation
func (s UpdateChannelInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateChannelInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateChannelInput"}

	if s.Id == nil {
		invalidParams.Add(aws.NewErrParamRequired("Id"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateChannelInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	if s.Description != nil {
		v := *s.Description

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "description", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Id != nil {
		v := *s.Id

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/mediapackage-2017-10-12/UpdateChannelResponse
type UpdateChannelOutput struct {
	_ struct{} `type:"structure"`

	Arn *string `locationName:"arn" type:"string"`

	Description *string `locationName:"description" type:"string"`

	// An HTTP Live Streaming (HLS) ingest resource configuration.
	HlsIngest *HlsIngest `locationName:"hlsIngest" type:"structure"`

	Id *string `locationName:"id" type:"string"`

	// A collection of tags associated with a resource
	Tags map[string]string `locationName:"tags" type:"map"`
}

// String returns the string representation
func (s UpdateChannelOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateChannelOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Arn != nil {
		v := *s.Arn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "arn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Description != nil {
		v := *s.Description

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "description", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.HlsIngest != nil {
		v := s.HlsIngest

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "hlsIngest", v, metadata)
	}
	if s.Id != nil {
		v := *s.Id

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
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

const opUpdateChannel = "UpdateChannel"

// UpdateChannelRequest returns a request value for making API operation for
// AWS Elemental MediaPackage.
//
// Updates an existing Channel.
//
//    // Example sending a request using UpdateChannelRequest.
//    req := client.UpdateChannelRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/mediapackage-2017-10-12/UpdateChannel
func (c *Client) UpdateChannelRequest(input *UpdateChannelInput) UpdateChannelRequest {
	op := &aws.Operation{
		Name:       opUpdateChannel,
		HTTPMethod: "PUT",
		HTTPPath:   "/channels/{id}",
	}

	if input == nil {
		input = &UpdateChannelInput{}
	}

	req := c.newRequest(op, input, &UpdateChannelOutput{})
	return UpdateChannelRequest{Request: req, Input: input, Copy: c.UpdateChannelRequest}
}

// UpdateChannelRequest is the request type for the
// UpdateChannel API operation.
type UpdateChannelRequest struct {
	*aws.Request
	Input *UpdateChannelInput
	Copy  func(*UpdateChannelInput) UpdateChannelRequest
}

// Send marshals and sends the UpdateChannel API request.
func (r UpdateChannelRequest) Send(ctx context.Context) (*UpdateChannelResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateChannelResponse{
		UpdateChannelOutput: r.Request.Data.(*UpdateChannelOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateChannelResponse is the response type for the
// UpdateChannel API operation.
type UpdateChannelResponse struct {
	*UpdateChannelOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateChannel request.
func (r *UpdateChannelResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
