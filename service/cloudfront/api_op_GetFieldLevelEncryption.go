// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudfront

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/cloudfront-2018-11-05/GetFieldLevelEncryptionRequest
type GetFieldLevelEncryptionInput struct {
	_ struct{} `type:"structure"`

	// Request the ID for the field-level encryption configuration information.
	//
	// Id is a required field
	Id *string `location:"uri" locationName:"Id" type:"string" required:"true"`
}

// String returns the string representation
func (s GetFieldLevelEncryptionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetFieldLevelEncryptionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetFieldLevelEncryptionInput"}

	if s.Id == nil {
		invalidParams.Add(aws.NewErrParamRequired("Id"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetFieldLevelEncryptionInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.Id != nil {
		v := *s.Id

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "Id", protocol.StringValue(v), metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/cloudfront-2018-11-05/GetFieldLevelEncryptionResult
type GetFieldLevelEncryptionOutput struct {
	_ struct{} `type:"structure" payload:"FieldLevelEncryption"`

	// The current version of the field level encryption configuration. For example:
	// E2QWRUHAPOMQZL.
	ETag *string `location:"header" locationName:"ETag" type:"string"`

	// Return the field-level encryption configuration information.
	FieldLevelEncryption *FieldLevelEncryption `type:"structure"`
}

// String returns the string representation
func (s GetFieldLevelEncryptionOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetFieldLevelEncryptionOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.ETag != nil {
		v := *s.ETag

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "ETag", protocol.StringValue(v), metadata)
	}
	if s.FieldLevelEncryption != nil {
		v := s.FieldLevelEncryption

		metadata := protocol.Metadata{}
		e.SetFields(protocol.PayloadTarget, "FieldLevelEncryption", v, metadata)
	}
	return nil
}

const opGetFieldLevelEncryption = "GetFieldLevelEncryption2018_11_05"

// GetFieldLevelEncryptionRequest returns a request value for making API operation for
// Amazon CloudFront.
//
// Get the field-level encryption configuration information.
//
//    // Example sending a request using GetFieldLevelEncryptionRequest.
//    req := client.GetFieldLevelEncryptionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cloudfront-2018-11-05/GetFieldLevelEncryption
func (c *Client) GetFieldLevelEncryptionRequest(input *GetFieldLevelEncryptionInput) GetFieldLevelEncryptionRequest {
	op := &aws.Operation{
		Name:       opGetFieldLevelEncryption,
		HTTPMethod: "GET",
		HTTPPath:   "/2018-11-05/field-level-encryption/{Id}",
	}

	if input == nil {
		input = &GetFieldLevelEncryptionInput{}
	}

	req := c.newRequest(op, input, &GetFieldLevelEncryptionOutput{})
	return GetFieldLevelEncryptionRequest{Request: req, Input: input, Copy: c.GetFieldLevelEncryptionRequest}
}

// GetFieldLevelEncryptionRequest is the request type for the
// GetFieldLevelEncryption API operation.
type GetFieldLevelEncryptionRequest struct {
	*aws.Request
	Input *GetFieldLevelEncryptionInput
	Copy  func(*GetFieldLevelEncryptionInput) GetFieldLevelEncryptionRequest
}

// Send marshals and sends the GetFieldLevelEncryption API request.
func (r GetFieldLevelEncryptionRequest) Send(ctx context.Context) (*GetFieldLevelEncryptionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetFieldLevelEncryptionResponse{
		GetFieldLevelEncryptionOutput: r.Request.Data.(*GetFieldLevelEncryptionOutput),
		response:                      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetFieldLevelEncryptionResponse is the response type for the
// GetFieldLevelEncryption API operation.
type GetFieldLevelEncryptionResponse struct {
	*GetFieldLevelEncryptionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetFieldLevelEncryption request.
func (r *GetFieldLevelEncryptionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
