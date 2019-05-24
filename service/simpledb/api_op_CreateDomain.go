// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package simpledb

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/query"
)

type CreateDomainInput struct {
	_ struct{} `type:"structure"`

	// The name of the domain to create. The name can range between 3 and 255 characters
	// and can contain the following characters: a-z, A-Z, 0-9, '_', '-', and '.'.
	//
	// DomainName is a required field
	DomainName *string `type:"string" required:"true"`
}

// String returns the string representation
func (s CreateDomainInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateDomainInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateDomainInput"}

	if s.DomainName == nil {
		invalidParams.Add(aws.NewErrParamRequired("DomainName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreateDomainOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s CreateDomainOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateDomain = "CreateDomain"

// CreateDomainRequest returns a request value for making API operation for
// Amazon SimpleDB.
//
// The CreateDomain operation creates a new domain. The domain name should be
// unique among the domains associated with the Access Key ID provided in the
// request. The CreateDomain operation may take 10 or more seconds to complete.
//
// CreateDomain is an idempotent operation; running it multiple times using
// the same domain name will not result in an error response. The client can create up to 100 domains per account.
//
// If the client requires additional domains, go to  http://aws.amazon.com/contact-us/simpledb-limit-request/
// (http://aws.amazon.com/contact-us/simpledb-limit-request/).
//
//    // Example sending a request using CreateDomainRequest.
//    req := client.CreateDomainRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) CreateDomainRequest(input *CreateDomainInput) CreateDomainRequest {
	op := &aws.Operation{
		Name:       opCreateDomain,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateDomainInput{}
	}

	req := c.newRequest(op, input, &CreateDomainOutput{})
	req.Handlers.Unmarshal.Remove(query.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return CreateDomainRequest{Request: req, Input: input, Copy: c.CreateDomainRequest}
}

// CreateDomainRequest is the request type for the
// CreateDomain API operation.
type CreateDomainRequest struct {
	*aws.Request
	Input *CreateDomainInput
	Copy  func(*CreateDomainInput) CreateDomainRequest
}

// Send marshals and sends the CreateDomain API request.
func (r CreateDomainRequest) Send(ctx context.Context) (*CreateDomainResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateDomainResponse{
		CreateDomainOutput: r.Request.Data.(*CreateDomainOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateDomainResponse is the response type for the
// CreateDomain API operation.
type CreateDomainResponse struct {
	*CreateDomainOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateDomain request.
func (r *CreateDomainResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
