// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package apigateway

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Requests API Gateway to change information about a Stage resource.
type UpdateStageInput struct {
	_ struct{} `type:"structure"`

	// A list of update operations to be applied to the specified resource and in
	// the order specified in this list.
	PatchOperations []PatchOperation `locationName:"patchOperations" type:"list"`

	// [Required] The string identifier of the associated RestApi.
	//
	// RestApiId is a required field
	RestApiId *string `location:"uri" locationName:"restapi_id" type:"string" required:"true"`

	// [Required] The name of the Stage resource to change information about.
	//
	// StageName is a required field
	StageName *string `location:"uri" locationName:"stage_name" type:"string" required:"true"`
}

// String returns the string representation
func (s UpdateStageInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateStageInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateStageInput"}

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
func (s UpdateStageInput) MarshalFields(e protocol.FieldEncoder) error {

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

// Represents a unique identifier for a version of a deployed RestApi that is
// callable by users.
//
// Deploy an API (https://docs.aws.amazon.com/apigateway/latest/developerguide/how-to-deploy-api.html)
type UpdateStageOutput struct {
	_ struct{} `type:"structure"`

	// Settings for logging access in this stage.
	AccessLogSettings *AccessLogSettings `locationName:"accessLogSettings" type:"structure"`

	// Specifies whether a cache cluster is enabled for the stage.
	CacheClusterEnabled *bool `locationName:"cacheClusterEnabled" type:"boolean"`

	// The size of the cache cluster for the stage, if enabled.
	CacheClusterSize CacheClusterSize `locationName:"cacheClusterSize" type:"string" enum:"true"`

	// The status of the cache cluster for the stage, if enabled.
	CacheClusterStatus CacheClusterStatus `locationName:"cacheClusterStatus" type:"string" enum:"true"`

	// Settings for the canary deployment in this stage.
	CanarySettings *CanarySettings `locationName:"canarySettings" type:"structure"`

	// The identifier of a client certificate for an API stage.
	ClientCertificateId *string `locationName:"clientCertificateId" type:"string"`

	// The timestamp when the stage was created.
	CreatedDate *time.Time `locationName:"createdDate" type:"timestamp" timestampFormat:"unix"`

	// The identifier of the Deployment that the stage points to.
	DeploymentId *string `locationName:"deploymentId" type:"string"`

	// The stage's description.
	Description *string `locationName:"description" type:"string"`

	// The version of the associated API documentation.
	DocumentationVersion *string `locationName:"documentationVersion" type:"string"`

	// The timestamp when the stage last updated.
	LastUpdatedDate *time.Time `locationName:"lastUpdatedDate" type:"timestamp" timestampFormat:"unix"`

	// A map that defines the method settings for a Stage resource. Keys (designated
	// as /{method_setting_key below) are method paths defined as {resource_path}/{http_method}
	// for an individual method override, or /\*/\* for overriding all methods in
	// the stage.
	MethodSettings map[string]MethodSetting `locationName:"methodSettings" type:"map"`

	// The name of the stage is the first path segment in the Uniform Resource Identifier
	// (URI) of a call to API Gateway.
	StageName *string `locationName:"stageName" type:"string"`

	// The collection of tags. Each tag element is associated with a given resource.
	Tags map[string]string `locationName:"tags" type:"map"`

	// Specifies whether active tracing with X-ray is enabled for the Stage.
	TracingEnabled *bool `locationName:"tracingEnabled" type:"boolean"`

	// A map that defines the stage variables for a Stage resource. Variable names
	// can have alphanumeric and underscore characters, and the values must match
	// [A-Za-z0-9-._~:/?#&=,]+.
	Variables map[string]string `locationName:"variables" type:"map"`

	// The ARN of the WebAcl associated with the Stage.
	WebAclArn *string `locationName:"webAclArn" type:"string"`
}

// String returns the string representation
func (s UpdateStageOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateStageOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.AccessLogSettings != nil {
		v := s.AccessLogSettings

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "accessLogSettings", v, metadata)
	}
	if s.CacheClusterEnabled != nil {
		v := *s.CacheClusterEnabled

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "cacheClusterEnabled", protocol.BoolValue(v), metadata)
	}
	if len(s.CacheClusterSize) > 0 {
		v := s.CacheClusterSize

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "cacheClusterSize", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if len(s.CacheClusterStatus) > 0 {
		v := s.CacheClusterStatus

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "cacheClusterStatus", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.CanarySettings != nil {
		v := s.CanarySettings

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "canarySettings", v, metadata)
	}
	if s.ClientCertificateId != nil {
		v := *s.ClientCertificateId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "clientCertificateId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.CreatedDate != nil {
		v := *s.CreatedDate

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "createdDate", protocol.TimeValue{V: v, Format: protocol.UnixTimeFormat}, metadata)
	}
	if s.DeploymentId != nil {
		v := *s.DeploymentId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "deploymentId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Description != nil {
		v := *s.Description

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "description", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.DocumentationVersion != nil {
		v := *s.DocumentationVersion

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "documentationVersion", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.LastUpdatedDate != nil {
		v := *s.LastUpdatedDate

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "lastUpdatedDate", protocol.TimeValue{V: v, Format: protocol.UnixTimeFormat}, metadata)
	}
	if len(s.MethodSettings) > 0 {
		v := s.MethodSettings

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "methodSettings", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetFields(k1, v1)
		}
		ms0.End()

	}
	if s.StageName != nil {
		v := *s.StageName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "stageName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.Tags) > 0 {
		v := s.Tags

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "tags", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetValue(k1, protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ms0.End()

	}
	if s.TracingEnabled != nil {
		v := *s.TracingEnabled

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "tracingEnabled", protocol.BoolValue(v), metadata)
	}
	if len(s.Variables) > 0 {
		v := s.Variables

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "variables", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetValue(k1, protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ms0.End()

	}
	if s.WebAclArn != nil {
		v := *s.WebAclArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "webAclArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opUpdateStage = "UpdateStage"

// UpdateStageRequest returns a request value for making API operation for
// Amazon API Gateway.
//
// Changes information about a Stage resource.
//
//    // Example sending a request using UpdateStageRequest.
//    req := client.UpdateStageRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) UpdateStageRequest(input *UpdateStageInput) UpdateStageRequest {
	op := &aws.Operation{
		Name:       opUpdateStage,
		HTTPMethod: "PATCH",
		HTTPPath:   "/restapis/{restapi_id}/stages/{stage_name}",
	}

	if input == nil {
		input = &UpdateStageInput{}
	}

	req := c.newRequest(op, input, &UpdateStageOutput{})
	return UpdateStageRequest{Request: req, Input: input, Copy: c.UpdateStageRequest}
}

// UpdateStageRequest is the request type for the
// UpdateStage API operation.
type UpdateStageRequest struct {
	*aws.Request
	Input *UpdateStageInput
	Copy  func(*UpdateStageInput) UpdateStageRequest
}

// Send marshals and sends the UpdateStage API request.
func (r UpdateStageRequest) Send(ctx context.Context) (*UpdateStageResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateStageResponse{
		UpdateStageOutput: r.Request.Data.(*UpdateStageOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateStageResponse is the response type for the
// UpdateStage API operation.
type UpdateStageResponse struct {
	*UpdateStageOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateStage request.
func (r *UpdateStageResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
