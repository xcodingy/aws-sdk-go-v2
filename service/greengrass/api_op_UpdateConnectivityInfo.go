// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package greengrass

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Information required to update a Greengrass core's connectivity.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/greengrass-2017-06-07/UpdateConnectivityInfoRequest
type UpdateConnectivityInfoInput struct {
	_ struct{} `type:"structure"`

	// A list of connectivity info.
	ConnectivityInfo []ConnectivityInfo `type:"list"`

	// ThingName is a required field
	ThingName *string `location:"uri" locationName:"ThingName" type:"string" required:"true"`
}

// String returns the string representation
func (s UpdateConnectivityInfoInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateConnectivityInfoInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateConnectivityInfoInput"}

	if s.ThingName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ThingName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateConnectivityInfoInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	if len(s.ConnectivityInfo) > 0 {
		v := s.ConnectivityInfo

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "ConnectivityInfo", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.ThingName != nil {
		v := *s.ThingName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "ThingName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/greengrass-2017-06-07/UpdateConnectivityInfoResponse
type UpdateConnectivityInfoOutput struct {
	_ struct{} `type:"structure"`

	// A message about the connectivity info update request.
	Message *string `locationName:"message" type:"string"`

	// The new version of the connectivity info.
	Version *string `type:"string"`
}

// String returns the string representation
func (s UpdateConnectivityInfoOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateConnectivityInfoOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Message != nil {
		v := *s.Message

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "message", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Version != nil {
		v := *s.Version

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Version", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opUpdateConnectivityInfo = "UpdateConnectivityInfo"

// UpdateConnectivityInfoRequest returns a request value for making API operation for
// AWS Greengrass.
//
// Updates the connectivity information for the core. Any devices that belong
// to the group which has this core will receive this information in order to
// find the location of the core and connect to it.
//
//    // Example sending a request using UpdateConnectivityInfoRequest.
//    req := client.UpdateConnectivityInfoRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/greengrass-2017-06-07/UpdateConnectivityInfo
func (c *Client) UpdateConnectivityInfoRequest(input *UpdateConnectivityInfoInput) UpdateConnectivityInfoRequest {
	op := &aws.Operation{
		Name:       opUpdateConnectivityInfo,
		HTTPMethod: "PUT",
		HTTPPath:   "/greengrass/things/{ThingName}/connectivityInfo",
	}

	if input == nil {
		input = &UpdateConnectivityInfoInput{}
	}

	req := c.newRequest(op, input, &UpdateConnectivityInfoOutput{})
	return UpdateConnectivityInfoRequest{Request: req, Input: input, Copy: c.UpdateConnectivityInfoRequest}
}

// UpdateConnectivityInfoRequest is the request type for the
// UpdateConnectivityInfo API operation.
type UpdateConnectivityInfoRequest struct {
	*aws.Request
	Input *UpdateConnectivityInfoInput
	Copy  func(*UpdateConnectivityInfoInput) UpdateConnectivityInfoRequest
}

// Send marshals and sends the UpdateConnectivityInfo API request.
func (r UpdateConnectivityInfoRequest) Send(ctx context.Context) (*UpdateConnectivityInfoResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateConnectivityInfoResponse{
		UpdateConnectivityInfoOutput: r.Request.Data.(*UpdateConnectivityInfoOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateConnectivityInfoResponse is the response type for the
// UpdateConnectivityInfo API operation.
type UpdateConnectivityInfoResponse struct {
	*UpdateConnectivityInfoOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateConnectivityInfo request.
func (r *UpdateConnectivityInfoResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
