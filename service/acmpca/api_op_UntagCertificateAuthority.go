// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package acmpca

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/acm-pca-2017-08-22/UntagCertificateAuthorityRequest
type UntagCertificateAuthorityInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) that was returned when you called CreateCertificateAuthority.
	// This must be of the form:
	//
	// arn:aws:acm-pca:region:account:certificate-authority/12345678-1234-1234-1234-123456789012
	//
	// CertificateAuthorityArn is a required field
	CertificateAuthorityArn *string `min:"5" type:"string" required:"true"`

	// List of tags to be removed from the CA.
	//
	// Tags is a required field
	Tags []Tag `min:"1" type:"list" required:"true"`
}

// String returns the string representation
func (s UntagCertificateAuthorityInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UntagCertificateAuthorityInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UntagCertificateAuthorityInput"}

	if s.CertificateAuthorityArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("CertificateAuthorityArn"))
	}
	if s.CertificateAuthorityArn != nil && len(*s.CertificateAuthorityArn) < 5 {
		invalidParams.Add(aws.NewErrParamMinLen("CertificateAuthorityArn", 5))
	}

	if s.Tags == nil {
		invalidParams.Add(aws.NewErrParamRequired("Tags"))
	}
	if s.Tags != nil && len(s.Tags) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Tags", 1))
	}
	if s.Tags != nil {
		for i, v := range s.Tags {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Tags", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/acm-pca-2017-08-22/UntagCertificateAuthorityOutput
type UntagCertificateAuthorityOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s UntagCertificateAuthorityOutput) String() string {
	return awsutil.Prettify(s)
}

const opUntagCertificateAuthority = "UntagCertificateAuthority"

// UntagCertificateAuthorityRequest returns a request value for making API operation for
// AWS Certificate Manager Private Certificate Authority.
//
// Remove one or more tags from your private CA. A tag consists of a key-value
// pair. If you do not specify the value portion of the tag when calling this
// operation, the tag will be removed regardless of value. If you specify a
// value, the tag is removed only if it is associated with the specified value.
// To add tags to a private CA, use the TagCertificateAuthority. Call the ListTags
// operation to see what tags are associated with your CA.
//
//    // Example sending a request using UntagCertificateAuthorityRequest.
//    req := client.UntagCertificateAuthorityRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/acm-pca-2017-08-22/UntagCertificateAuthority
func (c *Client) UntagCertificateAuthorityRequest(input *UntagCertificateAuthorityInput) UntagCertificateAuthorityRequest {
	op := &aws.Operation{
		Name:       opUntagCertificateAuthority,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UntagCertificateAuthorityInput{}
	}

	req := c.newRequest(op, input, &UntagCertificateAuthorityOutput{})
	req.Handlers.Unmarshal.Remove(jsonrpc.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return UntagCertificateAuthorityRequest{Request: req, Input: input, Copy: c.UntagCertificateAuthorityRequest}
}

// UntagCertificateAuthorityRequest is the request type for the
// UntagCertificateAuthority API operation.
type UntagCertificateAuthorityRequest struct {
	*aws.Request
	Input *UntagCertificateAuthorityInput
	Copy  func(*UntagCertificateAuthorityInput) UntagCertificateAuthorityRequest
}

// Send marshals and sends the UntagCertificateAuthority API request.
func (r UntagCertificateAuthorityRequest) Send(ctx context.Context) (*UntagCertificateAuthorityResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UntagCertificateAuthorityResponse{
		UntagCertificateAuthorityOutput: r.Request.Data.(*UntagCertificateAuthorityOutput),
		response:                        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UntagCertificateAuthorityResponse is the response type for the
// UntagCertificateAuthority API operation.
type UntagCertificateAuthorityResponse struct {
	*UntagCertificateAuthorityOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UntagCertificateAuthority request.
func (r *UntagCertificateAuthorityResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
