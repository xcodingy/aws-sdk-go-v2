// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package mq

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/mq-2017-11-27/DeleteBrokerRequest
type DeleteBrokerInput struct {
	_ struct{} `type:"structure"`

	// BrokerId is a required field
	BrokerId *string `location:"uri" locationName:"broker-id" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteBrokerInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteBrokerInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteBrokerInput"}

	if s.BrokerId == nil {
		invalidParams.Add(aws.NewErrParamRequired("BrokerId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteBrokerInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	if s.BrokerId != nil {
		v := *s.BrokerId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "broker-id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/mq-2017-11-27/DeleteBrokerResponse
type DeleteBrokerOutput struct {
	_ struct{} `type:"structure"`

	BrokerId *string `locationName:"brokerId" type:"string"`
}

// String returns the string representation
func (s DeleteBrokerOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteBrokerOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.BrokerId != nil {
		v := *s.BrokerId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "brokerId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opDeleteBroker = "DeleteBroker"

// DeleteBrokerRequest returns a request value for making API operation for
// AmazonMQ.
//
// Deletes a broker. Note: This API is asynchronous.
//
//    // Example sending a request using DeleteBrokerRequest.
//    req := client.DeleteBrokerRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/mq-2017-11-27/DeleteBroker
func (c *Client) DeleteBrokerRequest(input *DeleteBrokerInput) DeleteBrokerRequest {
	op := &aws.Operation{
		Name:       opDeleteBroker,
		HTTPMethod: "DELETE",
		HTTPPath:   "/v1/brokers/{broker-id}",
	}

	if input == nil {
		input = &DeleteBrokerInput{}
	}

	req := c.newRequest(op, input, &DeleteBrokerOutput{})
	return DeleteBrokerRequest{Request: req, Input: input, Copy: c.DeleteBrokerRequest}
}

// DeleteBrokerRequest is the request type for the
// DeleteBroker API operation.
type DeleteBrokerRequest struct {
	*aws.Request
	Input *DeleteBrokerInput
	Copy  func(*DeleteBrokerInput) DeleteBrokerRequest
}

// Send marshals and sends the DeleteBroker API request.
func (r DeleteBrokerRequest) Send(ctx context.Context) (*DeleteBrokerResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteBrokerResponse{
		DeleteBrokerOutput: r.Request.Data.(*DeleteBrokerOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteBrokerResponse is the response type for the
// DeleteBroker API operation.
type DeleteBrokerResponse struct {
	*DeleteBrokerOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteBroker request.
func (r *DeleteBrokerResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
