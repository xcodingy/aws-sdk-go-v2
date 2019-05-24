// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package lambda

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/lambda-2015-03-31/GetLayerVersionPolicyRequest
type GetLayerVersionPolicyInput struct {
	_ struct{} `type:"structure"`

	// The name or Amazon Resource Name (ARN) of the layer.
	//
	// LayerName is a required field
	LayerName *string `location:"uri" locationName:"LayerName" min:"1" type:"string" required:"true"`

	// The version number.
	//
	// VersionNumber is a required field
	VersionNumber *int64 `location:"uri" locationName:"VersionNumber" type:"long" required:"true"`
}

// String returns the string representation
func (s GetLayerVersionPolicyInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetLayerVersionPolicyInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetLayerVersionPolicyInput"}

	if s.LayerName == nil {
		invalidParams.Add(aws.NewErrParamRequired("LayerName"))
	}
	if s.LayerName != nil && len(*s.LayerName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("LayerName", 1))
	}

	if s.VersionNumber == nil {
		invalidParams.Add(aws.NewErrParamRequired("VersionNumber"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetLayerVersionPolicyInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.LayerName != nil {
		v := *s.LayerName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "LayerName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.VersionNumber != nil {
		v := *s.VersionNumber

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "VersionNumber", protocol.Int64Value(v), metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/lambda-2015-03-31/GetLayerVersionPolicyResponse
type GetLayerVersionPolicyOutput struct {
	_ struct{} `type:"structure"`

	// The policy document.
	Policy *string `type:"string"`

	// A unique identifier for the current revision of the policy.
	RevisionId *string `type:"string"`
}

// String returns the string representation
func (s GetLayerVersionPolicyOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetLayerVersionPolicyOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Policy != nil {
		v := *s.Policy

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Policy", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.RevisionId != nil {
		v := *s.RevisionId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "RevisionId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opGetLayerVersionPolicy = "GetLayerVersionPolicy"

// GetLayerVersionPolicyRequest returns a request value for making API operation for
// AWS Lambda.
//
// Returns the permission policy for a version of an AWS Lambda layer (https://docs.aws.amazon.com/lambda/latest/dg/configuration-layers.html).
// For more information, see AddLayerVersionPermission.
//
//    // Example sending a request using GetLayerVersionPolicyRequest.
//    req := client.GetLayerVersionPolicyRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/lambda-2015-03-31/GetLayerVersionPolicy
func (c *Client) GetLayerVersionPolicyRequest(input *GetLayerVersionPolicyInput) GetLayerVersionPolicyRequest {
	op := &aws.Operation{
		Name:       opGetLayerVersionPolicy,
		HTTPMethod: "GET",
		HTTPPath:   "/2018-10-31/layers/{LayerName}/versions/{VersionNumber}/policy",
	}

	if input == nil {
		input = &GetLayerVersionPolicyInput{}
	}

	req := c.newRequest(op, input, &GetLayerVersionPolicyOutput{})
	return GetLayerVersionPolicyRequest{Request: req, Input: input, Copy: c.GetLayerVersionPolicyRequest}
}

// GetLayerVersionPolicyRequest is the request type for the
// GetLayerVersionPolicy API operation.
type GetLayerVersionPolicyRequest struct {
	*aws.Request
	Input *GetLayerVersionPolicyInput
	Copy  func(*GetLayerVersionPolicyInput) GetLayerVersionPolicyRequest
}

// Send marshals and sends the GetLayerVersionPolicy API request.
func (r GetLayerVersionPolicyRequest) Send(ctx context.Context) (*GetLayerVersionPolicyResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetLayerVersionPolicyResponse{
		GetLayerVersionPolicyOutput: r.Request.Data.(*GetLayerVersionPolicyOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetLayerVersionPolicyResponse is the response type for the
// GetLayerVersionPolicy API operation.
type GetLayerVersionPolicyResponse struct {
	*GetLayerVersionPolicyOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetLayerVersionPolicy request.
func (r *GetLayerVersionPolicyResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
