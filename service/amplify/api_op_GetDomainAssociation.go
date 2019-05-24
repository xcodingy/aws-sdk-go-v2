// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package amplify

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Request structure for the get Domain Association request.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/amplify-2017-07-25/GetDomainAssociationRequest
type GetDomainAssociationInput struct {
	_ struct{} `type:"structure"`

	// Unique Id for an Amplify App.
	//
	// AppId is a required field
	AppId *string `location:"uri" locationName:"appId" min:"1" type:"string" required:"true"`

	// Name of the domain.
	//
	// DomainName is a required field
	DomainName *string `location:"uri" locationName:"domainName" type:"string" required:"true"`
}

// String returns the string representation
func (s GetDomainAssociationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetDomainAssociationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetDomainAssociationInput"}

	if s.AppId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AppId"))
	}
	if s.AppId != nil && len(*s.AppId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("AppId", 1))
	}

	if s.DomainName == nil {
		invalidParams.Add(aws.NewErrParamRequired("DomainName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetDomainAssociationInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	if s.AppId != nil {
		v := *s.AppId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "appId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.DomainName != nil {
		v := *s.DomainName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "domainName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Result structure for the get Domain Association request.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/amplify-2017-07-25/GetDomainAssociationResult
type GetDomainAssociationOutput struct {
	_ struct{} `type:"structure"`

	// Domain Association structure.
	//
	// DomainAssociation is a required field
	DomainAssociation *DomainAssociation `locationName:"domainAssociation" type:"structure" required:"true"`
}

// String returns the string representation
func (s GetDomainAssociationOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetDomainAssociationOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.DomainAssociation != nil {
		v := s.DomainAssociation

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "domainAssociation", v, metadata)
	}
	return nil
}

const opGetDomainAssociation = "GetDomainAssociation"

// GetDomainAssociationRequest returns a request value for making API operation for
// AWS Amplify.
//
// Retrieves domain info that corresponds to an appId and domainName.
//
//    // Example sending a request using GetDomainAssociationRequest.
//    req := client.GetDomainAssociationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/amplify-2017-07-25/GetDomainAssociation
func (c *Client) GetDomainAssociationRequest(input *GetDomainAssociationInput) GetDomainAssociationRequest {
	op := &aws.Operation{
		Name:       opGetDomainAssociation,
		HTTPMethod: "GET",
		HTTPPath:   "/apps/{appId}/domains/{domainName}",
	}

	if input == nil {
		input = &GetDomainAssociationInput{}
	}

	req := c.newRequest(op, input, &GetDomainAssociationOutput{})
	return GetDomainAssociationRequest{Request: req, Input: input, Copy: c.GetDomainAssociationRequest}
}

// GetDomainAssociationRequest is the request type for the
// GetDomainAssociation API operation.
type GetDomainAssociationRequest struct {
	*aws.Request
	Input *GetDomainAssociationInput
	Copy  func(*GetDomainAssociationInput) GetDomainAssociationRequest
}

// Send marshals and sends the GetDomainAssociation API request.
func (r GetDomainAssociationRequest) Send(ctx context.Context) (*GetDomainAssociationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetDomainAssociationResponse{
		GetDomainAssociationOutput: r.Request.Data.(*GetDomainAssociationOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetDomainAssociationResponse is the response type for the
// GetDomainAssociation API operation.
type GetDomainAssociationResponse struct {
	*GetDomainAssociationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetDomainAssociation request.
func (r *GetDomainAssociationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
