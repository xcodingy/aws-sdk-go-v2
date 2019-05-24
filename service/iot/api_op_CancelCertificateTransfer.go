// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iot

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
)

// The input for the CancelCertificateTransfer operation.
type CancelCertificateTransferInput struct {
	_ struct{} `type:"structure"`

	// The ID of the certificate. (The last part of the certificate ARN contains
	// the certificate ID.)
	//
	// CertificateId is a required field
	CertificateId *string `location:"uri" locationName:"certificateId" min:"64" type:"string" required:"true"`
}

// String returns the string representation
func (s CancelCertificateTransferInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CancelCertificateTransferInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CancelCertificateTransferInput"}

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
func (s CancelCertificateTransferInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.CertificateId != nil {
		v := *s.CertificateId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "certificateId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type CancelCertificateTransferOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s CancelCertificateTransferOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CancelCertificateTransferOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opCancelCertificateTransfer = "CancelCertificateTransfer"

// CancelCertificateTransferRequest returns a request value for making API operation for
// AWS IoT.
//
// Cancels a pending transfer for the specified certificate.
//
// Note Only the transfer source account can use this operation to cancel a
// transfer. (Transfer destinations can use RejectCertificateTransfer instead.)
// After transfer, AWS IoT returns the certificate to the source account in
// the INACTIVE state. After the destination account has accepted the transfer,
// the transfer cannot be cancelled.
//
// After a certificate transfer is cancelled, the status of the certificate
// changes from PENDING_TRANSFER to INACTIVE.
//
//    // Example sending a request using CancelCertificateTransferRequest.
//    req := client.CancelCertificateTransferRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) CancelCertificateTransferRequest(input *CancelCertificateTransferInput) CancelCertificateTransferRequest {
	op := &aws.Operation{
		Name:       opCancelCertificateTransfer,
		HTTPMethod: "PATCH",
		HTTPPath:   "/cancel-certificate-transfer/{certificateId}",
	}

	if input == nil {
		input = &CancelCertificateTransferInput{}
	}

	req := c.newRequest(op, input, &CancelCertificateTransferOutput{})
	req.Handlers.Unmarshal.Remove(restjson.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return CancelCertificateTransferRequest{Request: req, Input: input, Copy: c.CancelCertificateTransferRequest}
}

// CancelCertificateTransferRequest is the request type for the
// CancelCertificateTransfer API operation.
type CancelCertificateTransferRequest struct {
	*aws.Request
	Input *CancelCertificateTransferInput
	Copy  func(*CancelCertificateTransferInput) CancelCertificateTransferRequest
}

// Send marshals and sends the CancelCertificateTransfer API request.
func (r CancelCertificateTransferRequest) Send(ctx context.Context) (*CancelCertificateTransferResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CancelCertificateTransferResponse{
		CancelCertificateTransferOutput: r.Request.Data.(*CancelCertificateTransferOutput),
		response:                        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CancelCertificateTransferResponse is the response type for the
// CancelCertificateTransfer API operation.
type CancelCertificateTransferResponse struct {
	*CancelCertificateTransferOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CancelCertificateTransfer request.
func (r *CancelCertificateTransferResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
