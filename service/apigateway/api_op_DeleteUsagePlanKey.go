// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package apigateway

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
)

// The DELETE request to delete a usage plan key and remove the underlying API
// key from the associated usage plan.
type DeleteUsagePlanKeyInput struct {
	_ struct{} `type:"structure"`

	// [Required] The Id of the UsagePlanKey resource to be deleted.
	//
	// KeyId is a required field
	KeyId *string `location:"uri" locationName:"keyId" type:"string" required:"true"`

	// [Required] The Id of the UsagePlan resource representing the usage plan containing
	// the to-be-deleted UsagePlanKey resource representing a plan customer.
	//
	// UsagePlanId is a required field
	UsagePlanId *string `location:"uri" locationName:"usageplanId" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteUsagePlanKeyInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteUsagePlanKeyInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteUsagePlanKeyInput"}

	if s.KeyId == nil {
		invalidParams.Add(aws.NewErrParamRequired("KeyId"))
	}

	if s.UsagePlanId == nil {
		invalidParams.Add(aws.NewErrParamRequired("UsagePlanId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteUsagePlanKeyInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.KeyId != nil {
		v := *s.KeyId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "keyId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.UsagePlanId != nil {
		v := *s.UsagePlanId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "usageplanId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type DeleteUsagePlanKeyOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteUsagePlanKeyOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteUsagePlanKeyOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opDeleteUsagePlanKey = "DeleteUsagePlanKey"

// DeleteUsagePlanKeyRequest returns a request value for making API operation for
// Amazon API Gateway.
//
// Deletes a usage plan key and remove the underlying API key from the associated
// usage plan.
//
//    // Example sending a request using DeleteUsagePlanKeyRequest.
//    req := client.DeleteUsagePlanKeyRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) DeleteUsagePlanKeyRequest(input *DeleteUsagePlanKeyInput) DeleteUsagePlanKeyRequest {
	op := &aws.Operation{
		Name:       opDeleteUsagePlanKey,
		HTTPMethod: "DELETE",
		HTTPPath:   "/usageplans/{usageplanId}/keys/{keyId}",
	}

	if input == nil {
		input = &DeleteUsagePlanKeyInput{}
	}

	req := c.newRequest(op, input, &DeleteUsagePlanKeyOutput{})
	req.Handlers.Unmarshal.Remove(restjson.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return DeleteUsagePlanKeyRequest{Request: req, Input: input, Copy: c.DeleteUsagePlanKeyRequest}
}

// DeleteUsagePlanKeyRequest is the request type for the
// DeleteUsagePlanKey API operation.
type DeleteUsagePlanKeyRequest struct {
	*aws.Request
	Input *DeleteUsagePlanKeyInput
	Copy  func(*DeleteUsagePlanKeyInput) DeleteUsagePlanKeyRequest
}

// Send marshals and sends the DeleteUsagePlanKey API request.
func (r DeleteUsagePlanKeyRequest) Send(ctx context.Context) (*DeleteUsagePlanKeyResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteUsagePlanKeyResponse{
		DeleteUsagePlanKeyOutput: r.Request.Data.(*DeleteUsagePlanKeyOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteUsagePlanKeyResponse is the response type for the
// DeleteUsagePlanKey API operation.
type DeleteUsagePlanKeyResponse struct {
	*DeleteUsagePlanKeyOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteUsagePlanKey request.
func (r *DeleteUsagePlanKeyResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
