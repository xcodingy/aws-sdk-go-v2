// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package apigateway

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Represents an update integration response request.
type UpdateIntegrationResponseInput struct {
	_ struct{} `type:"structure"`

	// [Required] Specifies an update integration response request's HTTP method.
	//
	// HttpMethod is a required field
	HttpMethod *string `location:"uri" locationName:"http_method" type:"string" required:"true"`

	// A list of update operations to be applied to the specified resource and in
	// the order specified in this list.
	PatchOperations []PatchOperation `locationName:"patchOperations" type:"list"`

	// [Required] Specifies an update integration response request's resource identifier.
	//
	// ResourceId is a required field
	ResourceId *string `location:"uri" locationName:"resource_id" type:"string" required:"true"`

	// [Required] The string identifier of the associated RestApi.
	//
	// RestApiId is a required field
	RestApiId *string `location:"uri" locationName:"restapi_id" type:"string" required:"true"`

	// [Required] Specifies an update integration response request's status code.
	//
	// StatusCode is a required field
	StatusCode *string `location:"uri" locationName:"status_code" type:"string" required:"true"`
}

// String returns the string representation
func (s UpdateIntegrationResponseInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateIntegrationResponseInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateIntegrationResponseInput"}

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
func (s UpdateIntegrationResponseInput) MarshalFields(e protocol.FieldEncoder) error {

	if len(s.PatchOperations) > 0 {
		v := s.PatchOperations

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "patchOperations", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
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

// Represents an integration response. The status code must map to an existing
// MethodResponse, and parameters and templates can be used to transform the
// back-end response.
//
// Creating an API (https://docs.aws.amazon.com/apigateway/latest/developerguide/how-to-create-api.html)
type UpdateIntegrationResponseOutput struct {
	_ struct{} `type:"structure"`

	// Specifies how to handle response payload content type conversions. Supported
	// values are CONVERT_TO_BINARY and CONVERT_TO_TEXT, with the following behaviors:
	//
	//    * CONVERT_TO_BINARY: Converts a response payload from a Base64-encoded
	//    string to the corresponding binary blob.
	//
	//    * CONVERT_TO_TEXT: Converts a response payload from a binary blob to a
	//    Base64-encoded string.
	//
	// If this property is not defined, the response payload will be passed through
	// from the integration response to the method response without modification.
	ContentHandling ContentHandlingStrategy `locationName:"contentHandling" type:"string" enum:"true"`

	// A key-value map specifying response parameters that are passed to the method
	// response from the back end. The key is a method response header parameter
	// name and the mapped value is an integration response header value, a static
	// value enclosed within a pair of single quotes, or a JSON expression from
	// the integration response body. The mapping key must match the pattern of
	// method.response.header.{name}, where name is a valid and unique header name.
	// The mapped non-static value must match the pattern of integration.response.header.{name}
	// or integration.response.body.{JSON-expression}, where name is a valid and
	// unique response header name and JSON-expression is a valid JSON expression
	// without the $ prefix.
	ResponseParameters map[string]string `locationName:"responseParameters" type:"map"`

	// Specifies the templates used to transform the integration response body.
	// Response templates are represented as a key/value map, with a content-type
	// as the key and a template as the value.
	ResponseTemplates map[string]string `locationName:"responseTemplates" type:"map"`

	// Specifies the regular expression (regex) pattern used to choose an integration
	// response based on the response from the back end. For example, if the success
	// response returns nothing and the error response returns some string, you
	// could use the .+ regex to match error response. However, make sure that the
	// error response does not contain any newline (\n) character in such cases.
	// If the back end is an AWS Lambda function, the AWS Lambda function error
	// header is matched. For all other HTTP and AWS back ends, the HTTP status
	// code is matched.
	SelectionPattern *string `locationName:"selectionPattern" type:"string"`

	// Specifies the status code that is used to map the integration response to
	// an existing MethodResponse.
	StatusCode *string `locationName:"statusCode" type:"string"`
}

// String returns the string representation
func (s UpdateIntegrationResponseOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateIntegrationResponseOutput) MarshalFields(e protocol.FieldEncoder) error {
	if len(s.ContentHandling) > 0 {
		v := s.ContentHandling

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "contentHandling", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if len(s.ResponseParameters) > 0 {
		v := s.ResponseParameters

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "responseParameters", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetValue(k1, protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ms0.End()

	}
	if len(s.ResponseTemplates) > 0 {
		v := s.ResponseTemplates

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "responseTemplates", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetValue(k1, protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ms0.End()

	}
	if s.SelectionPattern != nil {
		v := *s.SelectionPattern

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "selectionPattern", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.StatusCode != nil {
		v := *s.StatusCode

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "statusCode", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opUpdateIntegrationResponse = "UpdateIntegrationResponse"

// UpdateIntegrationResponseRequest returns a request value for making API operation for
// Amazon API Gateway.
//
// Represents an update integration response.
//
//    // Example sending a request using UpdateIntegrationResponseRequest.
//    req := client.UpdateIntegrationResponseRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) UpdateIntegrationResponseRequest(input *UpdateIntegrationResponseInput) UpdateIntegrationResponseRequest {
	op := &aws.Operation{
		Name:       opUpdateIntegrationResponse,
		HTTPMethod: "PATCH",
		HTTPPath:   "/restapis/{restapi_id}/resources/{resource_id}/methods/{http_method}/integration/responses/{status_code}",
	}

	if input == nil {
		input = &UpdateIntegrationResponseInput{}
	}

	req := c.newRequest(op, input, &UpdateIntegrationResponseOutput{})
	return UpdateIntegrationResponseRequest{Request: req, Input: input, Copy: c.UpdateIntegrationResponseRequest}
}

// UpdateIntegrationResponseRequest is the request type for the
// UpdateIntegrationResponse API operation.
type UpdateIntegrationResponseRequest struct {
	*aws.Request
	Input *UpdateIntegrationResponseInput
	Copy  func(*UpdateIntegrationResponseInput) UpdateIntegrationResponseRequest
}

// Send marshals and sends the UpdateIntegrationResponse API request.
func (r UpdateIntegrationResponseRequest) Send(ctx context.Context) (*UpdateIntegrationResponseResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateIntegrationResponseResponse{
		UpdateIntegrationResponseOutput: r.Request.Data.(*UpdateIntegrationResponseOutput),
		response:                        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateIntegrationResponseResponse is the response type for the
// UpdateIntegrationResponse API operation.
type UpdateIntegrationResponseResponse struct {
	*UpdateIntegrationResponseOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateIntegrationResponse request.
func (r *UpdateIntegrationResponseResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
