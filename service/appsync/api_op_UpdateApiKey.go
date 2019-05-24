// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package appsync

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/appsync-2017-07-25/UpdateApiKeyRequest
type UpdateApiKeyInput struct {
	_ struct{} `type:"structure"`

	// The ID for the GraphQL API.
	//
	// ApiId is a required field
	ApiId *string `location:"uri" locationName:"apiId" type:"string" required:"true"`

	// A description of the purpose of the API key.
	Description *string `locationName:"description" type:"string"`

	// The time from update time after which the API key expires. The date is represented
	// as seconds since the epoch. For more information, see .
	Expires *int64 `locationName:"expires" type:"long"`

	// The API key ID.
	//
	// Id is a required field
	Id *string `location:"uri" locationName:"id" type:"string" required:"true"`
}

// String returns the string representation
func (s UpdateApiKeyInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateApiKeyInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateApiKeyInput"}

	if s.ApiId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ApiId"))
	}

	if s.Id == nil {
		invalidParams.Add(aws.NewErrParamRequired("Id"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateApiKeyInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	if s.Description != nil {
		v := *s.Description

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "description", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Expires != nil {
		v := *s.Expires

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "expires", protocol.Int64Value(v), metadata)
	}
	if s.ApiId != nil {
		v := *s.ApiId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "apiId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Id != nil {
		v := *s.Id

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/appsync-2017-07-25/UpdateApiKeyResponse
type UpdateApiKeyOutput struct {
	_ struct{} `type:"structure"`

	// The API key.
	ApiKey *ApiKey `locationName:"apiKey" type:"structure"`
}

// String returns the string representation
func (s UpdateApiKeyOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateApiKeyOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.ApiKey != nil {
		v := s.ApiKey

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "apiKey", v, metadata)
	}
	return nil
}

const opUpdateApiKey = "UpdateApiKey"

// UpdateApiKeyRequest returns a request value for making API operation for
// AWS AppSync.
//
// Updates an API key.
//
//    // Example sending a request using UpdateApiKeyRequest.
//    req := client.UpdateApiKeyRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/appsync-2017-07-25/UpdateApiKey
func (c *Client) UpdateApiKeyRequest(input *UpdateApiKeyInput) UpdateApiKeyRequest {
	op := &aws.Operation{
		Name:       opUpdateApiKey,
		HTTPMethod: "POST",
		HTTPPath:   "/v1/apis/{apiId}/apikeys/{id}",
	}

	if input == nil {
		input = &UpdateApiKeyInput{}
	}

	req := c.newRequest(op, input, &UpdateApiKeyOutput{})
	return UpdateApiKeyRequest{Request: req, Input: input, Copy: c.UpdateApiKeyRequest}
}

// UpdateApiKeyRequest is the request type for the
// UpdateApiKey API operation.
type UpdateApiKeyRequest struct {
	*aws.Request
	Input *UpdateApiKeyInput
	Copy  func(*UpdateApiKeyInput) UpdateApiKeyRequest
}

// Send marshals and sends the UpdateApiKey API request.
func (r UpdateApiKeyRequest) Send(ctx context.Context) (*UpdateApiKeyResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateApiKeyResponse{
		UpdateApiKeyOutput: r.Request.Data.(*UpdateApiKeyOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateApiKeyResponse is the response type for the
// UpdateApiKey API operation.
type UpdateApiKeyResponse struct {
	*UpdateApiKeyOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateApiKey request.
func (r *UpdateApiKeyResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
