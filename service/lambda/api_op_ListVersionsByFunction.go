// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package lambda

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/lambda-2015-03-31/ListVersionsByFunctionRequest
type ListVersionsByFunctionInput struct {
	_ struct{} `type:"structure"`

	// The name of the Lambda function.
	//
	// Name formats
	//
	//    * Function name - MyFunction.
	//
	//    * Function ARN - arn:aws:lambda:us-west-2:123456789012:function:MyFunction.
	//
	//    * Partial ARN - 123456789012:function:MyFunction.
	//
	// The length constraint applies only to the full ARN. If you specify only the
	// function name, it is limited to 64 characters in length.
	//
	// FunctionName is a required field
	FunctionName *string `location:"uri" locationName:"FunctionName" min:"1" type:"string" required:"true"`

	// Specify the pagination token that's returned by a previous request to retrieve
	// the next page of results.
	Marker *string `location:"querystring" locationName:"Marker" type:"string"`

	// Limit the number of versions that are returned.
	MaxItems *int64 `location:"querystring" locationName:"MaxItems" min:"1" type:"integer"`
}

// String returns the string representation
func (s ListVersionsByFunctionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListVersionsByFunctionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListVersionsByFunctionInput"}

	if s.FunctionName == nil {
		invalidParams.Add(aws.NewErrParamRequired("FunctionName"))
	}
	if s.FunctionName != nil && len(*s.FunctionName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("FunctionName", 1))
	}
	if s.MaxItems != nil && *s.MaxItems < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxItems", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListVersionsByFunctionInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.FunctionName != nil {
		v := *s.FunctionName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "FunctionName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Marker != nil {
		v := *s.Marker

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "Marker", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.MaxItems != nil {
		v := *s.MaxItems

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "MaxItems", protocol.Int64Value(v), metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/lambda-2015-03-31/ListVersionsByFunctionResponse
type ListVersionsByFunctionOutput struct {
	_ struct{} `type:"structure"`

	// The pagination token that's included if more results are available.
	NextMarker *string `type:"string"`

	// A list of Lambda function versions.
	Versions []FunctionConfiguration `type:"list"`
}

// String returns the string representation
func (s ListVersionsByFunctionOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListVersionsByFunctionOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.NextMarker != nil {
		v := *s.NextMarker

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "NextMarker", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.Versions) > 0 {
		v := s.Versions

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "Versions", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	return nil
}

const opListVersionsByFunction = "ListVersionsByFunction"

// ListVersionsByFunctionRequest returns a request value for making API operation for
// AWS Lambda.
//
// Returns a list of versions (https://docs.aws.amazon.com/lambda/latest/dg/versioning-aliases.html),
// with the version-specific configuration of each.
//
//    // Example sending a request using ListVersionsByFunctionRequest.
//    req := client.ListVersionsByFunctionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/lambda-2015-03-31/ListVersionsByFunction
func (c *Client) ListVersionsByFunctionRequest(input *ListVersionsByFunctionInput) ListVersionsByFunctionRequest {
	op := &aws.Operation{
		Name:       opListVersionsByFunction,
		HTTPMethod: "GET",
		HTTPPath:   "/2015-03-31/functions/{FunctionName}/versions",
	}

	if input == nil {
		input = &ListVersionsByFunctionInput{}
	}

	req := c.newRequest(op, input, &ListVersionsByFunctionOutput{})
	return ListVersionsByFunctionRequest{Request: req, Input: input, Copy: c.ListVersionsByFunctionRequest}
}

// ListVersionsByFunctionRequest is the request type for the
// ListVersionsByFunction API operation.
type ListVersionsByFunctionRequest struct {
	*aws.Request
	Input *ListVersionsByFunctionInput
	Copy  func(*ListVersionsByFunctionInput) ListVersionsByFunctionRequest
}

// Send marshals and sends the ListVersionsByFunction API request.
func (r ListVersionsByFunctionRequest) Send(ctx context.Context) (*ListVersionsByFunctionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListVersionsByFunctionResponse{
		ListVersionsByFunctionOutput: r.Request.Data.(*ListVersionsByFunctionOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListVersionsByFunctionResponse is the response type for the
// ListVersionsByFunction API operation.
type ListVersionsByFunctionResponse struct {
	*ListVersionsByFunctionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListVersionsByFunction request.
func (r *ListVersionsByFunctionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
