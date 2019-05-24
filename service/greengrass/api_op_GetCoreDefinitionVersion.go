// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package greengrass

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/greengrass-2017-06-07/GetCoreDefinitionVersionRequest
type GetCoreDefinitionVersionInput struct {
	_ struct{} `type:"structure"`

	// CoreDefinitionId is a required field
	CoreDefinitionId *string `location:"uri" locationName:"CoreDefinitionId" type:"string" required:"true"`

	// CoreDefinitionVersionId is a required field
	CoreDefinitionVersionId *string `location:"uri" locationName:"CoreDefinitionVersionId" type:"string" required:"true"`
}

// String returns the string representation
func (s GetCoreDefinitionVersionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetCoreDefinitionVersionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetCoreDefinitionVersionInput"}

	if s.CoreDefinitionId == nil {
		invalidParams.Add(aws.NewErrParamRequired("CoreDefinitionId"))
	}

	if s.CoreDefinitionVersionId == nil {
		invalidParams.Add(aws.NewErrParamRequired("CoreDefinitionVersionId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetCoreDefinitionVersionInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	if s.CoreDefinitionId != nil {
		v := *s.CoreDefinitionId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "CoreDefinitionId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.CoreDefinitionVersionId != nil {
		v := *s.CoreDefinitionVersionId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "CoreDefinitionVersionId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/greengrass-2017-06-07/GetCoreDefinitionVersionResponse
type GetCoreDefinitionVersionOutput struct {
	_ struct{} `type:"structure"`

	// The ARN of the core definition version.
	Arn *string `type:"string"`

	// The time, in milliseconds since the epoch, when the core definition version
	// was created.
	CreationTimestamp *string `type:"string"`

	// Information about the core definition version.
	Definition *CoreDefinitionVersion `type:"structure"`

	// The ID of the core definition version.
	Id *string `type:"string"`

	// The token for the next set of results, or ''null'' if there are no additional
	// results.
	NextToken *string `type:"string"`

	// The version of the core definition version.
	Version *string `type:"string"`
}

// String returns the string representation
func (s GetCoreDefinitionVersionOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetCoreDefinitionVersionOutput) MarshalFields(e protocol.FieldEncoder) error {
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
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "NextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Version != nil {
		v := *s.Version

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Version", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opGetCoreDefinitionVersion = "GetCoreDefinitionVersion"

// GetCoreDefinitionVersionRequest returns a request value for making API operation for
// AWS Greengrass.
//
// Retrieves information about a core definition version.
//
//    // Example sending a request using GetCoreDefinitionVersionRequest.
//    req := client.GetCoreDefinitionVersionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/greengrass-2017-06-07/GetCoreDefinitionVersion
func (c *Client) GetCoreDefinitionVersionRequest(input *GetCoreDefinitionVersionInput) GetCoreDefinitionVersionRequest {
	op := &aws.Operation{
		Name:       opGetCoreDefinitionVersion,
		HTTPMethod: "GET",
		HTTPPath:   "/greengrass/definition/cores/{CoreDefinitionId}/versions/{CoreDefinitionVersionId}",
	}

	if input == nil {
		input = &GetCoreDefinitionVersionInput{}
	}

	req := c.newRequest(op, input, &GetCoreDefinitionVersionOutput{})
	return GetCoreDefinitionVersionRequest{Request: req, Input: input, Copy: c.GetCoreDefinitionVersionRequest}
}

// GetCoreDefinitionVersionRequest is the request type for the
// GetCoreDefinitionVersion API operation.
type GetCoreDefinitionVersionRequest struct {
	*aws.Request
	Input *GetCoreDefinitionVersionInput
	Copy  func(*GetCoreDefinitionVersionInput) GetCoreDefinitionVersionRequest
}

// Send marshals and sends the GetCoreDefinitionVersion API request.
func (r GetCoreDefinitionVersionRequest) Send(ctx context.Context) (*GetCoreDefinitionVersionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetCoreDefinitionVersionResponse{
		GetCoreDefinitionVersionOutput: r.Request.Data.(*GetCoreDefinitionVersionOutput),
		response:                       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetCoreDefinitionVersionResponse is the response type for the
// GetCoreDefinitionVersion API operation.
type GetCoreDefinitionVersionResponse struct {
	*GetCoreDefinitionVersionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetCoreDefinitionVersion request.
func (r *GetCoreDefinitionVersionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
