// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package directoryservice

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Updates a conditional forwarder.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ds-2015-04-16/UpdateConditionalForwarderRequest
type UpdateConditionalForwarderInput struct {
	_ struct{} `type:"structure"`

	// The directory ID of the AWS directory for which to update the conditional
	// forwarder.
	//
	// DirectoryId is a required field
	DirectoryId *string `type:"string" required:"true"`

	// The updated IP addresses of the remote DNS server associated with the conditional
	// forwarder.
	//
	// DnsIpAddrs is a required field
	DnsIpAddrs []string `type:"list" required:"true"`

	// The fully qualified domain name (FQDN) of the remote domain with which you
	// will set up a trust relationship.
	//
	// RemoteDomainName is a required field
	RemoteDomainName *string `type:"string" required:"true"`
}

// String returns the string representation
func (s UpdateConditionalForwarderInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateConditionalForwarderInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateConditionalForwarderInput"}

	if s.DirectoryId == nil {
		invalidParams.Add(aws.NewErrParamRequired("DirectoryId"))
	}

	if s.DnsIpAddrs == nil {
		invalidParams.Add(aws.NewErrParamRequired("DnsIpAddrs"))
	}

	if s.RemoteDomainName == nil {
		invalidParams.Add(aws.NewErrParamRequired("RemoteDomainName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// The result of an UpdateConditionalForwarder request.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ds-2015-04-16/UpdateConditionalForwarderResult
type UpdateConditionalForwarderOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s UpdateConditionalForwarderOutput) String() string {
	return awsutil.Prettify(s)
}

const opUpdateConditionalForwarder = "UpdateConditionalForwarder"

// UpdateConditionalForwarderRequest returns a request value for making API operation for
// AWS Directory Service.
//
// Updates a conditional forwarder that has been set up for your AWS directory.
//
//    // Example sending a request using UpdateConditionalForwarderRequest.
//    req := client.UpdateConditionalForwarderRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ds-2015-04-16/UpdateConditionalForwarder
func (c *Client) UpdateConditionalForwarderRequest(input *UpdateConditionalForwarderInput) UpdateConditionalForwarderRequest {
	op := &aws.Operation{
		Name:       opUpdateConditionalForwarder,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateConditionalForwarderInput{}
	}

	req := c.newRequest(op, input, &UpdateConditionalForwarderOutput{})
	return UpdateConditionalForwarderRequest{Request: req, Input: input, Copy: c.UpdateConditionalForwarderRequest}
}

// UpdateConditionalForwarderRequest is the request type for the
// UpdateConditionalForwarder API operation.
type UpdateConditionalForwarderRequest struct {
	*aws.Request
	Input *UpdateConditionalForwarderInput
	Copy  func(*UpdateConditionalForwarderInput) UpdateConditionalForwarderRequest
}

// Send marshals and sends the UpdateConditionalForwarder API request.
func (r UpdateConditionalForwarderRequest) Send(ctx context.Context) (*UpdateConditionalForwarderResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateConditionalForwarderResponse{
		UpdateConditionalForwarderOutput: r.Request.Data.(*UpdateConditionalForwarderOutput),
		response:                         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateConditionalForwarderResponse is the response type for the
// UpdateConditionalForwarder API operation.
type UpdateConditionalForwarderResponse struct {
	*UpdateConditionalForwarderOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateConditionalForwarder request.
func (r *UpdateConditionalForwarderResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
