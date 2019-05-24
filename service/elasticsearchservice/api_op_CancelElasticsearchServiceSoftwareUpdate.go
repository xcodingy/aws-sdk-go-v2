// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package elasticsearchservice

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Container for the parameters to the CancelElasticsearchServiceSoftwareUpdate
// operation. Specifies the name of the Elasticsearch domain that you wish to
// cancel a service software update on.
type CancelElasticsearchServiceSoftwareUpdateInput struct {
	_ struct{} `type:"structure"`

	// The name of the domain that you want to stop the latest service software
	// update on.
	//
	// DomainName is a required field
	DomainName *string `min:"3" type:"string" required:"true"`
}

// String returns the string representation
func (s CancelElasticsearchServiceSoftwareUpdateInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CancelElasticsearchServiceSoftwareUpdateInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CancelElasticsearchServiceSoftwareUpdateInput"}

	if s.DomainName == nil {
		invalidParams.Add(aws.NewErrParamRequired("DomainName"))
	}
	if s.DomainName != nil && len(*s.DomainName) < 3 {
		invalidParams.Add(aws.NewErrParamMinLen("DomainName", 3))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CancelElasticsearchServiceSoftwareUpdateInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.DomainName != nil {
		v := *s.DomainName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "DomainName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// The result of a CancelElasticsearchServiceSoftwareUpdate operation. Contains
// the status of the update.
type CancelElasticsearchServiceSoftwareUpdateOutput struct {
	_ struct{} `type:"structure"`

	// The current status of the Elasticsearch service software update.
	ServiceSoftwareOptions *ServiceSoftwareOptions `type:"structure"`
}

// String returns the string representation
func (s CancelElasticsearchServiceSoftwareUpdateOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CancelElasticsearchServiceSoftwareUpdateOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.ServiceSoftwareOptions != nil {
		v := s.ServiceSoftwareOptions

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "ServiceSoftwareOptions", v, metadata)
	}
	return nil
}

const opCancelElasticsearchServiceSoftwareUpdate = "CancelElasticsearchServiceSoftwareUpdate"

// CancelElasticsearchServiceSoftwareUpdateRequest returns a request value for making API operation for
// Amazon Elasticsearch Service.
//
// Cancels a scheduled service software update for an Amazon ES domain. You
// can only perform this operation before the AutomatedUpdateDate and when the
// UpdateStatus is in the PENDING_UPDATE state.
//
//    // Example sending a request using CancelElasticsearchServiceSoftwareUpdateRequest.
//    req := client.CancelElasticsearchServiceSoftwareUpdateRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) CancelElasticsearchServiceSoftwareUpdateRequest(input *CancelElasticsearchServiceSoftwareUpdateInput) CancelElasticsearchServiceSoftwareUpdateRequest {
	op := &aws.Operation{
		Name:       opCancelElasticsearchServiceSoftwareUpdate,
		HTTPMethod: "POST",
		HTTPPath:   "/2015-01-01/es/serviceSoftwareUpdate/cancel",
	}

	if input == nil {
		input = &CancelElasticsearchServiceSoftwareUpdateInput{}
	}

	req := c.newRequest(op, input, &CancelElasticsearchServiceSoftwareUpdateOutput{})
	return CancelElasticsearchServiceSoftwareUpdateRequest{Request: req, Input: input, Copy: c.CancelElasticsearchServiceSoftwareUpdateRequest}
}

// CancelElasticsearchServiceSoftwareUpdateRequest is the request type for the
// CancelElasticsearchServiceSoftwareUpdate API operation.
type CancelElasticsearchServiceSoftwareUpdateRequest struct {
	*aws.Request
	Input *CancelElasticsearchServiceSoftwareUpdateInput
	Copy  func(*CancelElasticsearchServiceSoftwareUpdateInput) CancelElasticsearchServiceSoftwareUpdateRequest
}

// Send marshals and sends the CancelElasticsearchServiceSoftwareUpdate API request.
func (r CancelElasticsearchServiceSoftwareUpdateRequest) Send(ctx context.Context) (*CancelElasticsearchServiceSoftwareUpdateResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CancelElasticsearchServiceSoftwareUpdateResponse{
		CancelElasticsearchServiceSoftwareUpdateOutput: r.Request.Data.(*CancelElasticsearchServiceSoftwareUpdateOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CancelElasticsearchServiceSoftwareUpdateResponse is the response type for the
// CancelElasticsearchServiceSoftwareUpdate API operation.
type CancelElasticsearchServiceSoftwareUpdateResponse struct {
	*CancelElasticsearchServiceSoftwareUpdateOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CancelElasticsearchServiceSoftwareUpdate request.
func (r *CancelElasticsearchServiceSoftwareUpdateResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
