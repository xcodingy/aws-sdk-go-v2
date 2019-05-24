// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package storagegateway

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/storagegateway-2013-06-30/UpdateVTLDeviceTypeInput
type UpdateVTLDeviceTypeInput struct {
	_ struct{} `type:"structure"`

	// The type of medium changer you want to select.
	//
	// Valid Values: "STK-L700", "AWS-Gateway-VTL"
	//
	// DeviceType is a required field
	DeviceType *string `min:"2" type:"string" required:"true"`

	// The Amazon Resource Name (ARN) of the medium changer you want to select.
	//
	// VTLDeviceARN is a required field
	VTLDeviceARN *string `min:"50" type:"string" required:"true"`
}

// String returns the string representation
func (s UpdateVTLDeviceTypeInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateVTLDeviceTypeInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateVTLDeviceTypeInput"}

	if s.DeviceType == nil {
		invalidParams.Add(aws.NewErrParamRequired("DeviceType"))
	}
	if s.DeviceType != nil && len(*s.DeviceType) < 2 {
		invalidParams.Add(aws.NewErrParamMinLen("DeviceType", 2))
	}

	if s.VTLDeviceARN == nil {
		invalidParams.Add(aws.NewErrParamRequired("VTLDeviceARN"))
	}
	if s.VTLDeviceARN != nil && len(*s.VTLDeviceARN) < 50 {
		invalidParams.Add(aws.NewErrParamMinLen("VTLDeviceARN", 50))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// UpdateVTLDeviceTypeOutput
// Please also see https://docs.aws.amazon.com/goto/WebAPI/storagegateway-2013-06-30/UpdateVTLDeviceTypeOutput
type UpdateVTLDeviceTypeOutput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the medium changer you have selected.
	VTLDeviceARN *string `min:"50" type:"string"`
}

// String returns the string representation
func (s UpdateVTLDeviceTypeOutput) String() string {
	return awsutil.Prettify(s)
}

const opUpdateVTLDeviceType = "UpdateVTLDeviceType"

// UpdateVTLDeviceTypeRequest returns a request value for making API operation for
// AWS Storage Gateway.
//
// Updates the type of medium changer in a tape gateway. When you activate a
// tape gateway, you select a medium changer type for the tape gateway. This
// operation enables you to select a different type of medium changer after
// a tape gateway is activated. This operation is only supported in the tape
// gateway type.
//
//    // Example sending a request using UpdateVTLDeviceTypeRequest.
//    req := client.UpdateVTLDeviceTypeRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/storagegateway-2013-06-30/UpdateVTLDeviceType
func (c *Client) UpdateVTLDeviceTypeRequest(input *UpdateVTLDeviceTypeInput) UpdateVTLDeviceTypeRequest {
	op := &aws.Operation{
		Name:       opUpdateVTLDeviceType,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateVTLDeviceTypeInput{}
	}

	req := c.newRequest(op, input, &UpdateVTLDeviceTypeOutput{})
	return UpdateVTLDeviceTypeRequest{Request: req, Input: input, Copy: c.UpdateVTLDeviceTypeRequest}
}

// UpdateVTLDeviceTypeRequest is the request type for the
// UpdateVTLDeviceType API operation.
type UpdateVTLDeviceTypeRequest struct {
	*aws.Request
	Input *UpdateVTLDeviceTypeInput
	Copy  func(*UpdateVTLDeviceTypeInput) UpdateVTLDeviceTypeRequest
}

// Send marshals and sends the UpdateVTLDeviceType API request.
func (r UpdateVTLDeviceTypeRequest) Send(ctx context.Context) (*UpdateVTLDeviceTypeResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateVTLDeviceTypeResponse{
		UpdateVTLDeviceTypeOutput: r.Request.Data.(*UpdateVTLDeviceTypeOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateVTLDeviceTypeResponse is the response type for the
// UpdateVTLDeviceType API operation.
type UpdateVTLDeviceTypeResponse struct {
	*UpdateVTLDeviceTypeOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateVTLDeviceType request.
func (r *UpdateVTLDeviceTypeResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
