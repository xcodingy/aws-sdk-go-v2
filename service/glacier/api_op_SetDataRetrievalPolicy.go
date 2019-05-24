// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package glacier

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
)

// SetDataRetrievalPolicy input.
type SetDataRetrievalPolicyInput struct {
	_ struct{} `type:"structure"`

	// The AccountId value is the AWS account ID. This value must match the AWS
	// account ID associated with the credentials used to sign the request. You
	// can either specify an AWS account ID or optionally a single '-' (hyphen),
	// in which case Amazon Glacier uses the AWS account ID associated with the
	// credentials used to sign the request. If you specify your account ID, do
	// not include any hyphens ('-') in the ID.
	//
	// AccountId is a required field
	AccountId *string `location:"uri" locationName:"accountId" type:"string" required:"true"`

	// The data retrieval policy in JSON format.
	Policy *DataRetrievalPolicy `type:"structure"`
}

// String returns the string representation
func (s SetDataRetrievalPolicyInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *SetDataRetrievalPolicyInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "SetDataRetrievalPolicyInput"}

	if s.AccountId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AccountId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s SetDataRetrievalPolicyInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.Policy != nil {
		v := s.Policy

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "Policy", v, metadata)
	}
	if s.AccountId != nil {
		v := *s.AccountId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "accountId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type SetDataRetrievalPolicyOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s SetDataRetrievalPolicyOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s SetDataRetrievalPolicyOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opSetDataRetrievalPolicy = "SetDataRetrievalPolicy"

// SetDataRetrievalPolicyRequest returns a request value for making API operation for
// Amazon Glacier.
//
// This operation sets and then enacts a data retrieval policy in the region
// specified in the PUT request. You can set one policy per region for an AWS
// account. The policy is enacted within a few minutes of a successful PUT operation.
//
// The set policy operation does not affect retrieval jobs that were in progress
// before the policy was enacted. For more information about data retrieval
// policies, see Amazon Glacier Data Retrieval Policies (http://docs.aws.amazon.com/amazonglacier/latest/dev/data-retrieval-policy.html).
//
//    // Example sending a request using SetDataRetrievalPolicyRequest.
//    req := client.SetDataRetrievalPolicyRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) SetDataRetrievalPolicyRequest(input *SetDataRetrievalPolicyInput) SetDataRetrievalPolicyRequest {
	op := &aws.Operation{
		Name:       opSetDataRetrievalPolicy,
		HTTPMethod: "PUT",
		HTTPPath:   "/{accountId}/policies/data-retrieval",
	}

	if input == nil {
		input = &SetDataRetrievalPolicyInput{}
	}

	req := c.newRequest(op, input, &SetDataRetrievalPolicyOutput{})
	req.Handlers.Unmarshal.Remove(restjson.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return SetDataRetrievalPolicyRequest{Request: req, Input: input, Copy: c.SetDataRetrievalPolicyRequest}
}

// SetDataRetrievalPolicyRequest is the request type for the
// SetDataRetrievalPolicy API operation.
type SetDataRetrievalPolicyRequest struct {
	*aws.Request
	Input *SetDataRetrievalPolicyInput
	Copy  func(*SetDataRetrievalPolicyInput) SetDataRetrievalPolicyRequest
}

// Send marshals and sends the SetDataRetrievalPolicy API request.
func (r SetDataRetrievalPolicyRequest) Send(ctx context.Context) (*SetDataRetrievalPolicyResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &SetDataRetrievalPolicyResponse{
		SetDataRetrievalPolicyOutput: r.Request.Data.(*SetDataRetrievalPolicyOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// SetDataRetrievalPolicyResponse is the response type for the
// SetDataRetrievalPolicy API operation.
type SetDataRetrievalPolicyResponse struct {
	*SetDataRetrievalPolicyOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// SetDataRetrievalPolicy request.
func (r *SetDataRetrievalPolicyResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
