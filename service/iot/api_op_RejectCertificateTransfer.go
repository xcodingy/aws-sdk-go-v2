// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iot

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
)

// The input for the RejectCertificateTransfer operation.
type RejectCertificateTransferInput struct {
	_ struct{} `type:"structure"`

	// The ID of the certificate. (The last part of the certificate ARN contains
	// the certificate ID.)
	//
	// CertificateId is a required field
	CertificateId *string `location:"uri" locationName:"certificateId" min:"64" type:"string" required:"true"`

	// The reason the certificate transfer was rejected.
	RejectReason *string `locationName:"rejectReason" type:"string"`
}

// String returns the string representation
func (s RejectCertificateTransferInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *RejectCertificateTransferInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "RejectCertificateTransferInput"}

	if s.CertificateId == nil {
		invalidParams.Add(aws.NewErrParamRequired("CertificateId"))
	}
	if s.CertificateId != nil && len(*s.CertificateId) < 64 {
		invalidParams.Add(aws.NewErrParamMinLen("CertificateId", 64))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s RejectCertificateTransferInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.RejectReason != nil {
		v := *s.RejectReason

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "rejectReason", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.CertificateId != nil {
		v := *s.CertificateId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "certificateId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type RejectCertificateTransferOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s RejectCertificateTransferOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s RejectCertificateTransferOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opRejectCertificateTransfer = "RejectCertificateTransfer"

// RejectCertificateTransferRequest returns a request value for making API operation for
// AWS IoT.
//
// Rejects a pending certificate transfer. After AWS IoT rejects a certificate
// transfer, the certificate status changes from PENDING_TRANSFER to INACTIVE.
//
// To check for pending certificate transfers, call ListCertificates to enumerate
// your certificates.
//
// This operation can only be called by the transfer destination. After it is
// called, the certificate will be returned to the source's account in the INACTIVE
// state.
//
//    // Example sending a request using RejectCertificateTransferRequest.
//    req := client.RejectCertificateTransferRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) RejectCertificateTransferRequest(input *RejectCertificateTransferInput) RejectCertificateTransferRequest {
	op := &aws.Operation{
		Name:       opRejectCertificateTransfer,
		HTTPMethod: "PATCH",
		HTTPPath:   "/reject-certificate-transfer/{certificateId}",
	}

	if input == nil {
		input = &RejectCertificateTransferInput{}
	}

	req := c.newRequest(op, input, &RejectCertificateTransferOutput{})
	req.Handlers.Unmarshal.Remove(restjson.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return RejectCertificateTransferRequest{Request: req, Input: input, Copy: c.RejectCertificateTransferRequest}
}

// RejectCertificateTransferRequest is the request type for the
// RejectCertificateTransfer API operation.
type RejectCertificateTransferRequest struct {
	*aws.Request
	Input *RejectCertificateTransferInput
	Copy  func(*RejectCertificateTransferInput) RejectCertificateTransferRequest
}

// Send marshals and sends the RejectCertificateTransfer API request.
func (r RejectCertificateTransferRequest) Send(ctx context.Context) (*RejectCertificateTransferResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &RejectCertificateTransferResponse{
		RejectCertificateTransferOutput: r.Request.Data.(*RejectCertificateTransferOutput),
		response:                        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// RejectCertificateTransferResponse is the response type for the
// RejectCertificateTransfer API operation.
type RejectCertificateTransferResponse struct {
	*RejectCertificateTransferOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// RejectCertificateTransfer request.
func (r *RejectCertificateTransferResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
