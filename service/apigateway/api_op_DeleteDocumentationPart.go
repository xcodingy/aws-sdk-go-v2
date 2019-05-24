// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package apigateway

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
)

// Deletes an existing documentation part of an API.
type DeleteDocumentationPartInput struct {
	_ struct{} `type:"structure"`

	// [Required] The identifier of the to-be-deleted documentation part.
	//
	// DocumentationPartId is a required field
	DocumentationPartId *string `location:"uri" locationName:"part_id" type:"string" required:"true"`

	// [Required] The string identifier of the associated RestApi.
	//
	// RestApiId is a required field
	RestApiId *string `location:"uri" locationName:"restapi_id" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteDocumentationPartInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteDocumentationPartInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteDocumentationPartInput"}

	if s.DocumentationPartId == nil {
		invalidParams.Add(aws.NewErrParamRequired("DocumentationPartId"))
	}

	if s.RestApiId == nil {
		invalidParams.Add(aws.NewErrParamRequired("RestApiId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteDocumentationPartInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.DocumentationPartId != nil {
		v := *s.DocumentationPartId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "part_id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.RestApiId != nil {
		v := *s.RestApiId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "restapi_id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type DeleteDocumentationPartOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteDocumentationPartOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteDocumentationPartOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opDeleteDocumentationPart = "DeleteDocumentationPart"

// DeleteDocumentationPartRequest returns a request value for making API operation for
// Amazon API Gateway.
//
//    // Example sending a request using DeleteDocumentationPartRequest.
//    req := client.DeleteDocumentationPartRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) DeleteDocumentationPartRequest(input *DeleteDocumentationPartInput) DeleteDocumentationPartRequest {
	op := &aws.Operation{
		Name:       opDeleteDocumentationPart,
		HTTPMethod: "DELETE",
		HTTPPath:   "/restapis/{restapi_id}/documentation/parts/{part_id}",
	}

	if input == nil {
		input = &DeleteDocumentationPartInput{}
	}

	req := c.newRequest(op, input, &DeleteDocumentationPartOutput{})
	req.Handlers.Unmarshal.Remove(restjson.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return DeleteDocumentationPartRequest{Request: req, Input: input, Copy: c.DeleteDocumentationPartRequest}
}

// DeleteDocumentationPartRequest is the request type for the
// DeleteDocumentationPart API operation.
type DeleteDocumentationPartRequest struct {
	*aws.Request
	Input *DeleteDocumentationPartInput
	Copy  func(*DeleteDocumentationPartInput) DeleteDocumentationPartRequest
}

// Send marshals and sends the DeleteDocumentationPart API request.
func (r DeleteDocumentationPartRequest) Send(ctx context.Context) (*DeleteDocumentationPartResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteDocumentationPartResponse{
		DeleteDocumentationPartOutput: r.Request.Data.(*DeleteDocumentationPartOutput),
		response:                      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteDocumentationPartResponse is the response type for the
// DeleteDocumentationPart API operation.
type DeleteDocumentationPartResponse struct {
	*DeleteDocumentationPartOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteDocumentationPart request.
func (r *DeleteDocumentationPartResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
