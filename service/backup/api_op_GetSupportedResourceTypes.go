// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package backup

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/backup-2018-11-15/GetSupportedResourceTypesInput
type GetSupportedResourceTypesInput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s GetSupportedResourceTypesInput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetSupportedResourceTypesInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/backup-2018-11-15/GetSupportedResourceTypesOutput
type GetSupportedResourceTypesOutput struct {
	_ struct{} `type:"structure"`

	// Contains a string with the supported AWS resource types:
	//
	//    * EBS for Amazon Elastic Block Store
	//
	//    * SGW for AWS Storage Gateway
	//
	//    * RDS for Amazon Relational Database Service
	//
	//    * DDB for Amazon DynamoDB
	//
	//    * EFS for Amazon Elastic File System
	ResourceTypes []string `type:"list"`
}

// String returns the string representation
func (s GetSupportedResourceTypesOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetSupportedResourceTypesOutput) MarshalFields(e protocol.FieldEncoder) error {
	if len(s.ResourceTypes) > 0 {
		v := s.ResourceTypes

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "ResourceTypes", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	return nil
}

const opGetSupportedResourceTypes = "GetSupportedResourceTypes"

// GetSupportedResourceTypesRequest returns a request value for making API operation for
// AWS Backup.
//
// Returns the AWS resource types supported by AWS Backup.
//
//    // Example sending a request using GetSupportedResourceTypesRequest.
//    req := client.GetSupportedResourceTypesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/backup-2018-11-15/GetSupportedResourceTypes
func (c *Client) GetSupportedResourceTypesRequest(input *GetSupportedResourceTypesInput) GetSupportedResourceTypesRequest {
	op := &aws.Operation{
		Name:       opGetSupportedResourceTypes,
		HTTPMethod: "GET",
		HTTPPath:   "/supported-resource-types",
	}

	if input == nil {
		input = &GetSupportedResourceTypesInput{}
	}

	req := c.newRequest(op, input, &GetSupportedResourceTypesOutput{})
	return GetSupportedResourceTypesRequest{Request: req, Input: input, Copy: c.GetSupportedResourceTypesRequest}
}

// GetSupportedResourceTypesRequest is the request type for the
// GetSupportedResourceTypes API operation.
type GetSupportedResourceTypesRequest struct {
	*aws.Request
	Input *GetSupportedResourceTypesInput
	Copy  func(*GetSupportedResourceTypesInput) GetSupportedResourceTypesRequest
}

// Send marshals and sends the GetSupportedResourceTypes API request.
func (r GetSupportedResourceTypesRequest) Send(ctx context.Context) (*GetSupportedResourceTypesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetSupportedResourceTypesResponse{
		GetSupportedResourceTypesOutput: r.Request.Data.(*GetSupportedResourceTypesOutput),
		response:                        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetSupportedResourceTypesResponse is the response type for the
// GetSupportedResourceTypes API operation.
type GetSupportedResourceTypesResponse struct {
	*GetSupportedResourceTypesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetSupportedResourceTypes request.
func (r *GetSupportedResourceTypesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
