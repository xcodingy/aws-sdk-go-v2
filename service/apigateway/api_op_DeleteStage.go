// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package apigateway

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
)

// Requests API Gateway to delete a Stage resource.
type DeleteStageInput struct {
	_ struct{} `type:"structure"`

	// [Required] The string identifier of the associated RestApi.
	//
	// RestApiId is a required field
	RestApiId *string `location:"uri" locationName:"restapi_id" type:"string" required:"true"`

	// [Required] The name of the Stage resource to delete.
	//
	// StageName is a required field
	StageName *string `location:"uri" locationName:"stage_name" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteStageInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteStageInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteStageInput"}

	if s.RestApiId == nil {
		invalidParams.Add(aws.NewErrParamRequired("RestApiId"))
	}

	if s.StageName == nil {
		invalidParams.Add(aws.NewErrParamRequired("StageName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteStageInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.RestApiId != nil {
		v := *s.RestApiId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "restapi_id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.StageName != nil {
		v := *s.StageName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "stage_name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type DeleteStageOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteStageOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteStageOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opDeleteStage = "DeleteStage"

// DeleteStageRequest returns a request value for making API operation for
// Amazon API Gateway.
//
// Deletes a Stage resource.
//
//    // Example sending a request using DeleteStageRequest.
//    req := client.DeleteStageRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) DeleteStageRequest(input *DeleteStageInput) DeleteStageRequest {
	op := &aws.Operation{
		Name:       opDeleteStage,
		HTTPMethod: "DELETE",
		HTTPPath:   "/restapis/{restapi_id}/stages/{stage_name}",
	}

	if input == nil {
		input = &DeleteStageInput{}
	}

	req := c.newRequest(op, input, &DeleteStageOutput{})
	req.Handlers.Unmarshal.Remove(restjson.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return DeleteStageRequest{Request: req, Input: input, Copy: c.DeleteStageRequest}
}

// DeleteStageRequest is the request type for the
// DeleteStage API operation.
type DeleteStageRequest struct {
	*aws.Request
	Input *DeleteStageInput
	Copy  func(*DeleteStageInput) DeleteStageRequest
}

// Send marshals and sends the DeleteStage API request.
func (r DeleteStageRequest) Send(ctx context.Context) (*DeleteStageResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteStageResponse{
		DeleteStageOutput: r.Request.Data.(*DeleteStageOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteStageResponse is the response type for the
// DeleteStage API operation.
type DeleteStageResponse struct {
	*DeleteStageOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteStage request.
func (r *DeleteStageResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
