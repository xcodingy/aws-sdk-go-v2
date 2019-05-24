// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package lambda

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/lambda-2015-03-31/ListAliasesRequest
type ListAliasesInput struct {
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

	// Specify a function version to only list aliases that invoke that version.
	FunctionVersion *string `location:"querystring" locationName:"FunctionVersion" min:"1" type:"string"`

	// Specify the pagination token that's returned by a previous request to retrieve
	// the next page of results.
	Marker *string `location:"querystring" locationName:"Marker" type:"string"`

	// Limit the number of aliases returned.
	MaxItems *int64 `location:"querystring" locationName:"MaxItems" min:"1" type:"integer"`
}

// String returns the string representation
func (s ListAliasesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListAliasesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListAliasesInput"}

	if s.FunctionName == nil {
		invalidParams.Add(aws.NewErrParamRequired("FunctionName"))
	}
	if s.FunctionName != nil && len(*s.FunctionName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("FunctionName", 1))
	}
	if s.FunctionVersion != nil && len(*s.FunctionVersion) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("FunctionVersion", 1))
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
func (s ListAliasesInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.FunctionName != nil {
		v := *s.FunctionName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "FunctionName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.FunctionVersion != nil {
		v := *s.FunctionVersion

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "FunctionVersion", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
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

// Please also see https://docs.aws.amazon.com/goto/WebAPI/lambda-2015-03-31/ListAliasesResponse
type ListAliasesOutput struct {
	_ struct{} `type:"structure"`

	// A list of aliases.
	Aliases []AliasConfiguration `type:"list"`

	// The pagination token that's included if more results are available.
	NextMarker *string `type:"string"`
}

// String returns the string representation
func (s ListAliasesOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListAliasesOutput) MarshalFields(e protocol.FieldEncoder) error {
	if len(s.Aliases) > 0 {
		v := s.Aliases

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "Aliases", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.NextMarker != nil {
		v := *s.NextMarker

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "NextMarker", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opListAliases = "ListAliases"

// ListAliasesRequest returns a request value for making API operation for
// AWS Lambda.
//
// Returns a list of aliases (https://docs.aws.amazon.com/lambda/latest/dg/versioning-aliases.html)
// for a Lambda function.
//
//    // Example sending a request using ListAliasesRequest.
//    req := client.ListAliasesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/lambda-2015-03-31/ListAliases
func (c *Client) ListAliasesRequest(input *ListAliasesInput) ListAliasesRequest {
	op := &aws.Operation{
		Name:       opListAliases,
		HTTPMethod: "GET",
		HTTPPath:   "/2015-03-31/functions/{FunctionName}/aliases",
	}

	if input == nil {
		input = &ListAliasesInput{}
	}

	req := c.newRequest(op, input, &ListAliasesOutput{})
	return ListAliasesRequest{Request: req, Input: input, Copy: c.ListAliasesRequest}
}

// ListAliasesRequest is the request type for the
// ListAliases API operation.
type ListAliasesRequest struct {
	*aws.Request
	Input *ListAliasesInput
	Copy  func(*ListAliasesInput) ListAliasesRequest
}

// Send marshals and sends the ListAliases API request.
func (r ListAliasesRequest) Send(ctx context.Context) (*ListAliasesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListAliasesResponse{
		ListAliasesOutput: r.Request.Data.(*ListAliasesOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListAliasesResponse is the response type for the
// ListAliases API operation.
type ListAliasesResponse struct {
	*ListAliasesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListAliases request.
func (r *ListAliasesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
