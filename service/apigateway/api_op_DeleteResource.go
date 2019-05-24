// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package apigateway

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
)

// Request to delete a Resource.
type DeleteResourceInput struct {
	_ struct{} `type:"structure"`

	// [Required] The identifier of the Resource resource.
	//
	// ResourceId is a required field
	ResourceId *string `location:"uri" locationName:"resource_id" type:"string" required:"true"`

	// [Required] The string identifier of the associated RestApi.
	//
	// RestApiId is a required field
	RestApiId *string `location:"uri" locationName:"restapi_id" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteResourceInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteResourceInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteResourceInput"}

	if s.ResourceId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ResourceId"))
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
func (s DeleteResourceInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.ResourceId != nil {
		v := *s.ResourceId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "resource_id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.RestApiId != nil {
		v := *s.RestApiId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "restapi_id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type DeleteResourceOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteResourceOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteResourceOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opDeleteResource = "DeleteResource"

// DeleteResourceRequest returns a request value for making API operation for
// Amazon API Gateway.
//
// Deletes a Resource resource.
//
//    // Example sending a request using DeleteResourceRequest.
//    req := client.DeleteResourceRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) DeleteResourceRequest(input *DeleteResourceInput) DeleteResourceRequest {
	op := &aws.Operation{
		Name:       opDeleteResource,
		HTTPMethod: "DELETE",
		HTTPPath:   "/restapis/{restapi_id}/resources/{resource_id}",
	}

	if input == nil {
		input = &DeleteResourceInput{}
	}

	req := c.newRequest(op, input, &DeleteResourceOutput{})
	req.Handlers.Unmarshal.Remove(restjson.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return DeleteResourceRequest{Request: req, Input: input, Copy: c.DeleteResourceRequest}
}

// DeleteResourceRequest is the request type for the
// DeleteResource API operation.
type DeleteResourceRequest struct {
	*aws.Request
	Input *DeleteResourceInput
	Copy  func(*DeleteResourceInput) DeleteResourceRequest
}

// Send marshals and sends the DeleteResource API request.
func (r DeleteResourceRequest) Send(ctx context.Context) (*DeleteResourceResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteResourceResponse{
		DeleteResourceOutput: r.Request.Data.(*DeleteResourceOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteResourceResponse is the response type for the
// DeleteResource API operation.
type DeleteResourceResponse struct {
	*DeleteResourceOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteResource request.
func (r *DeleteResourceResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
