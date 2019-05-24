// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package greengrass

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/greengrass-2017-06-07/GetGroupVersionRequest
type GetGroupVersionInput struct {
	_ struct{} `type:"structure"`

	// GroupId is a required field
	GroupId *string `location:"uri" locationName:"GroupId" type:"string" required:"true"`

	// GroupVersionId is a required field
	GroupVersionId *string `location:"uri" locationName:"GroupVersionId" type:"string" required:"true"`
}

// String returns the string representation
func (s GetGroupVersionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetGroupVersionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetGroupVersionInput"}

	if s.GroupId == nil {
		invalidParams.Add(aws.NewErrParamRequired("GroupId"))
	}

	if s.GroupVersionId == nil {
		invalidParams.Add(aws.NewErrParamRequired("GroupVersionId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetGroupVersionInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	if s.GroupId != nil {
		v := *s.GroupId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "GroupId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.GroupVersionId != nil {
		v := *s.GroupVersionId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "GroupVersionId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Information about a group version.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/greengrass-2017-06-07/GetGroupVersionResponse
type GetGroupVersionOutput struct {
	_ struct{} `type:"structure"`

	// The ARN of the group version.
	Arn *string `type:"string"`

	// The time, in milliseconds since the epoch, when the group version was created.
	CreationTimestamp *string `type:"string"`

	// Information about the group version definition.
	Definition *GroupVersion `type:"structure"`

	// The ID of the group version.
	Id *string `type:"string"`

	// The unique ID for the version of the group.
	Version *string `type:"string"`
}

// String returns the string representation
func (s GetGroupVersionOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetGroupVersionOutput) MarshalFields(e protocol.FieldEncoder) error {
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
	if s.Definition != nil {
		v := s.Definition

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "Definition", v, metadata)
	}
	if s.Id != nil {
		v := *s.Id

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Version != nil {
		v := *s.Version

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Version", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opGetGroupVersion = "GetGroupVersion"

// GetGroupVersionRequest returns a request value for making API operation for
// AWS Greengrass.
//
// Retrieves information about a group version.
//
//    // Example sending a request using GetGroupVersionRequest.
//    req := client.GetGroupVersionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/greengrass-2017-06-07/GetGroupVersion
func (c *Client) GetGroupVersionRequest(input *GetGroupVersionInput) GetGroupVersionRequest {
	op := &aws.Operation{
		Name:       opGetGroupVersion,
		HTTPMethod: "GET",
		HTTPPath:   "/greengrass/groups/{GroupId}/versions/{GroupVersionId}",
	}

	if input == nil {
		input = &GetGroupVersionInput{}
	}

	req := c.newRequest(op, input, &GetGroupVersionOutput{})
	return GetGroupVersionRequest{Request: req, Input: input, Copy: c.GetGroupVersionRequest}
}

// GetGroupVersionRequest is the request type for the
// GetGroupVersion API operation.
type GetGroupVersionRequest struct {
	*aws.Request
	Input *GetGroupVersionInput
	Copy  func(*GetGroupVersionInput) GetGroupVersionRequest
}

// Send marshals and sends the GetGroupVersion API request.
func (r GetGroupVersionRequest) Send(ctx context.Context) (*GetGroupVersionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetGroupVersionResponse{
		GetGroupVersionOutput: r.Request.Data.(*GetGroupVersionOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetGroupVersionResponse is the response type for the
// GetGroupVersion API operation.
type GetGroupVersionResponse struct {
	*GetGroupVersionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetGroupVersion request.
func (r *GetGroupVersionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
