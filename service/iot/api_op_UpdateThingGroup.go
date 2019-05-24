// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iot

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type UpdateThingGroupInput struct {
	_ struct{} `type:"structure"`

	// The expected version of the thing group. If this does not match the version
	// of the thing group being updated, the update will fail.
	ExpectedVersion *int64 `locationName:"expectedVersion" type:"long"`

	// The thing group to update.
	//
	// ThingGroupName is a required field
	ThingGroupName *string `location:"uri" locationName:"thingGroupName" min:"1" type:"string" required:"true"`

	// The thing group properties.
	//
	// ThingGroupProperties is a required field
	ThingGroupProperties *ThingGroupProperties `locationName:"thingGroupProperties" type:"structure" required:"true"`
}

// String returns the string representation
func (s UpdateThingGroupInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateThingGroupInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateThingGroupInput"}

	if s.ThingGroupName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ThingGroupName"))
	}
	if s.ThingGroupName != nil && len(*s.ThingGroupName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ThingGroupName", 1))
	}

	if s.ThingGroupProperties == nil {
		invalidParams.Add(aws.NewErrParamRequired("ThingGroupProperties"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateThingGroupInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.ExpectedVersion != nil {
		v := *s.ExpectedVersion

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "expectedVersion", protocol.Int64Value(v), metadata)
	}
	if s.ThingGroupProperties != nil {
		v := s.ThingGroupProperties

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "thingGroupProperties", v, metadata)
	}
	if s.ThingGroupName != nil {
		v := *s.ThingGroupName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "thingGroupName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type UpdateThingGroupOutput struct {
	_ struct{} `type:"structure"`

	// The version of the updated thing group.
	Version *int64 `locationName:"version" type:"long"`
}

// String returns the string representation
func (s UpdateThingGroupOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateThingGroupOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Version != nil {
		v := *s.Version

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "version", protocol.Int64Value(v), metadata)
	}
	return nil
}

const opUpdateThingGroup = "UpdateThingGroup"

// UpdateThingGroupRequest returns a request value for making API operation for
// AWS IoT.
//
// Update a thing group.
//
//    // Example sending a request using UpdateThingGroupRequest.
//    req := client.UpdateThingGroupRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) UpdateThingGroupRequest(input *UpdateThingGroupInput) UpdateThingGroupRequest {
	op := &aws.Operation{
		Name:       opUpdateThingGroup,
		HTTPMethod: "PATCH",
		HTTPPath:   "/thing-groups/{thingGroupName}",
	}

	if input == nil {
		input = &UpdateThingGroupInput{}
	}

	req := c.newRequest(op, input, &UpdateThingGroupOutput{})
	return UpdateThingGroupRequest{Request: req, Input: input, Copy: c.UpdateThingGroupRequest}
}

// UpdateThingGroupRequest is the request type for the
// UpdateThingGroup API operation.
type UpdateThingGroupRequest struct {
	*aws.Request
	Input *UpdateThingGroupInput
	Copy  func(*UpdateThingGroupInput) UpdateThingGroupRequest
}

// Send marshals and sends the UpdateThingGroup API request.
func (r UpdateThingGroupRequest) Send(ctx context.Context) (*UpdateThingGroupResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateThingGroupResponse{
		UpdateThingGroupOutput: r.Request.Data.(*UpdateThingGroupOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateThingGroupResponse is the response type for the
// UpdateThingGroup API operation.
type UpdateThingGroupResponse struct {
	*UpdateThingGroupOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateThingGroup request.
func (r *UpdateThingGroupResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
