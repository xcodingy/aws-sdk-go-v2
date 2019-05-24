// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iot

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// The input for the CreateKeysAndCertificate operation.
type CreateKeysAndCertificateInput struct {
	_ struct{} `type:"structure"`

	// Specifies whether the certificate is active.
	SetAsActive *bool `location:"querystring" locationName:"setAsActive" type:"boolean"`
}

// String returns the string representation
func (s CreateKeysAndCertificateInput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateKeysAndCertificateInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.SetAsActive != nil {
		v := *s.SetAsActive

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "setAsActive", protocol.BoolValue(v), metadata)
	}
	return nil
}

// The output of the CreateKeysAndCertificate operation.
type CreateKeysAndCertificateOutput struct {
	_ struct{} `type:"structure"`

	// The ARN of the certificate.
	CertificateArn *string `locationName:"certificateArn" type:"string"`

	// The ID of the certificate. AWS IoT issues a default subject name for the
	// certificate (for example, AWS IoT Certificate).
	CertificateId *string `locationName:"certificateId" min:"64" type:"string"`

	// The certificate data, in PEM format.
	CertificatePem *string `locationName:"certificatePem" min:"1" type:"string"`

	// The generated key pair.
	KeyPair *KeyPair `locationName:"keyPair" type:"structure"`
}

// String returns the string representation
func (s CreateKeysAndCertificateOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateKeysAndCertificateOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.CertificateArn != nil {
		v := *s.CertificateArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "certificateArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.CertificateId != nil {
		v := *s.CertificateId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "certificateId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.CertificatePem != nil {
		v := *s.CertificatePem

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "certificatePem", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.KeyPair != nil {
		v := s.KeyPair

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "keyPair", v, metadata)
	}
	return nil
}

const opCreateKeysAndCertificate = "CreateKeysAndCertificate"

// CreateKeysAndCertificateRequest returns a request value for making API operation for
// AWS IoT.
//
// Creates a 2048-bit RSA key pair and issues an X.509 certificate using the
// issued public key.
//
// Note This is the only time AWS IoT issues the private key for this certificate,
// so it is important to keep it in a secure location.
//
//    // Example sending a request using CreateKeysAndCertificateRequest.
//    req := client.CreateKeysAndCertificateRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) CreateKeysAndCertificateRequest(input *CreateKeysAndCertificateInput) CreateKeysAndCertificateRequest {
	op := &aws.Operation{
		Name:       opCreateKeysAndCertificate,
		HTTPMethod: "POST",
		HTTPPath:   "/keys-and-certificate",
	}

	if input == nil {
		input = &CreateKeysAndCertificateInput{}
	}

	req := c.newRequest(op, input, &CreateKeysAndCertificateOutput{})
	return CreateKeysAndCertificateRequest{Request: req, Input: input, Copy: c.CreateKeysAndCertificateRequest}
}

// CreateKeysAndCertificateRequest is the request type for the
// CreateKeysAndCertificate API operation.
type CreateKeysAndCertificateRequest struct {
	*aws.Request
	Input *CreateKeysAndCertificateInput
	Copy  func(*CreateKeysAndCertificateInput) CreateKeysAndCertificateRequest
}

// Send marshals and sends the CreateKeysAndCertificate API request.
func (r CreateKeysAndCertificateRequest) Send(ctx context.Context) (*CreateKeysAndCertificateResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateKeysAndCertificateResponse{
		CreateKeysAndCertificateOutput: r.Request.Data.(*CreateKeysAndCertificateOutput),
		response:                       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateKeysAndCertificateResponse is the response type for the
// CreateKeysAndCertificate API operation.
type CreateKeysAndCertificateResponse struct {
	*CreateKeysAndCertificateOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateKeysAndCertificate request.
func (r *CreateKeysAndCertificateResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
