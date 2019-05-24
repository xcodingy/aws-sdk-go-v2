// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package apigateway

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
)

// A request to delete an existing MethodResponse resource.
type DeleteMethodResponseInput struct {
	_ struct{} `type:"structure"`

	// [Required] The HTTP verb of the Method resource.
	//
	// HttpMethod is a required field
	HttpMethod *string `location:"uri" locationName:"http_method" type:"string" required:"true"`

	// [Required] The Resource identifier for the MethodResponse resource.
	//
	// ResourceId is a required field
	ResourceId *string `location:"uri" locationName:"resource_id" type:"string" required:"true"`

	// [Required] The string identifier of the associated RestApi.
	//
	// RestApiId is a required field
	RestApiId *string `location:"uri" locationName:"restapi_id" type:"string" required:"true"`

	// [Required] The status code identifier for the MethodResponse resource.
	//
	// StatusCode is a required field
	StatusCode *string `location:"uri" locationName:"status_code" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteMethodResponseInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteMethodResponseInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteMethodResponseInput"}

	if s.HttpMethod == nil {
		invalidParams.Add(aws.NewErrParamRequired("HttpMethod"))
	}

	if s.ResourceId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ResourceId"))
	}

	if s.RestApiId == nil {
		invalidParams.Add(aws.NewErrParamRequired("RestApiId"))
	}

	if s.StatusCode == nil {
		invalidParams.Add(aws.NewErrParamRequired("StatusCode"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteMethodResponseInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.HttpMethod != nil {
		v := *s.HttpMethod

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "http_method", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
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
	if s.StatusCode != nil {
		v := *s.StatusCode

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "status_code", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type DeleteMethodResponseOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteMethodResponseOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteMethodResponseOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opDeleteMethodResponse = "DeleteMethodResponse"

// DeleteMethodResponseRequest returns a request value for making API operation for
// Amazon API Gateway.
//
// Deletes an existing MethodResponse resource.
//
//    // Example sending a request using DeleteMethodResponseRequest.
//    req := client.DeleteMethodResponseRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) DeleteMethodResponseRequest(input *DeleteMethodResponseInput) DeleteMethodResponseRequest {
	op := &aws.Operation{
		Name:       opDeleteMethodResponse,
		HTTPMethod: "DELETE",
		HTTPPath:   "/restapis/{restapi_id}/resources/{resource_id}/methods/{http_method}/responses/{status_code}",
	}

	if input == nil {
		input = &DeleteMethodResponseInput{}
	}

	req := c.newRequest(op, input, &DeleteMethodResponseOutput{})
	req.Handlers.Unmarshal.Remove(restjson.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return DeleteMethodResponseRequest{Request: req, Input: input, Copy: c.DeleteMethodResponseRequest}
}

// DeleteMethodResponseRequest is the request type for the
// DeleteMethodResponse API operation.
type DeleteMethodResponseRequest struct {
	*aws.Request
	Input *DeleteMethodResponseInput
	Copy  func(*DeleteMethodResponseInput) DeleteMethodResponseRequest
}

// Send marshals and sends the DeleteMethodResponse API request.
func (r DeleteMethodResponseRequest) Send(ctx context.Context) (*DeleteMethodResponseResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteMethodResponseResponse{
		DeleteMethodResponseOutput: r.Request.Data.(*DeleteMethodResponseOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteMethodResponseResponse is the response type for the
// DeleteMethodResponse API operation.
type DeleteMethodResponseResponse struct {
	*DeleteMethodResponseOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteMethodResponse request.
func (r *DeleteMethodResponseResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
