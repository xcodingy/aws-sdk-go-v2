// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iot

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type GetEffectivePoliciesInput struct {
	_ struct{} `type:"structure"`

	// The Cognito identity pool ID.
	CognitoIdentityPoolId *string `locationName:"cognitoIdentityPoolId" type:"string"`

	// The principal.
	Principal *string `locationName:"principal" type:"string"`

	// The thing name.
	ThingName *string `location:"querystring" locationName:"thingName" min:"1" type:"string"`
}

// String returns the string representation
func (s GetEffectivePoliciesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetEffectivePoliciesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetEffectivePoliciesInput"}
	if s.ThingName != nil && len(*s.ThingName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ThingName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetEffectivePoliciesInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.CognitoIdentityPoolId != nil {
		v := *s.CognitoIdentityPoolId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "cognitoIdentityPoolId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Principal != nil {
		v := *s.Principal

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "principal", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ThingName != nil {
		v := *s.ThingName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "thingName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type GetEffectivePoliciesOutput struct {
	_ struct{} `type:"structure"`

	// The effective policies.
	EffectivePolicies []EffectivePolicy `locationName:"effectivePolicies" type:"list"`
}

// String returns the string representation
func (s GetEffectivePoliciesOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetEffectivePoliciesOutput) MarshalFields(e protocol.FieldEncoder) error {
	if len(s.EffectivePolicies) > 0 {
		v := s.EffectivePolicies

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "effectivePolicies", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	return nil
}

const opGetEffectivePolicies = "GetEffectivePolicies"

// GetEffectivePoliciesRequest returns a request value for making API operation for
// AWS IoT.
//
// Gets a list of the policies that have an effect on the authorization behavior
// of the specified device when it connects to the AWS IoT device gateway.
//
//    // Example sending a request using GetEffectivePoliciesRequest.
//    req := client.GetEffectivePoliciesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) GetEffectivePoliciesRequest(input *GetEffectivePoliciesInput) GetEffectivePoliciesRequest {
	op := &aws.Operation{
		Name:       opGetEffectivePolicies,
		HTTPMethod: "POST",
		HTTPPath:   "/effective-policies",
	}

	if input == nil {
		input = &GetEffectivePoliciesInput{}
	}

	req := c.newRequest(op, input, &GetEffectivePoliciesOutput{})
	return GetEffectivePoliciesRequest{Request: req, Input: input, Copy: c.GetEffectivePoliciesRequest}
}

// GetEffectivePoliciesRequest is the request type for the
// GetEffectivePolicies API operation.
type GetEffectivePoliciesRequest struct {
	*aws.Request
	Input *GetEffectivePoliciesInput
	Copy  func(*GetEffectivePoliciesInput) GetEffectivePoliciesRequest
}

// Send marshals and sends the GetEffectivePolicies API request.
func (r GetEffectivePoliciesRequest) Send(ctx context.Context) (*GetEffectivePoliciesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetEffectivePoliciesResponse{
		GetEffectivePoliciesOutput: r.Request.Data.(*GetEffectivePoliciesOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetEffectivePoliciesResponse is the response type for the
// GetEffectivePolicies API operation.
type GetEffectivePoliciesResponse struct {
	*GetEffectivePoliciesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetEffectivePolicies request.
func (r *GetEffectivePoliciesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
