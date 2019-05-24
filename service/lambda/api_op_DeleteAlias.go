// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package lambda

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/lambda-2015-03-31/DeleteAliasRequest
type DeleteAliasInput struct {
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

	// The name of the alias.
	//
	// Name is a required field
	Name *string `location:"uri" locationName:"Name" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteAliasInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteAliasInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteAliasInput"}

	if s.FunctionName == nil {
		invalidParams.Add(aws.NewErrParamRequired("FunctionName"))
	}
	if s.FunctionName != nil && len(*s.FunctionName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("FunctionName", 1))
	}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}
	if s.Name != nil && len(*s.Name) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Name", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteAliasInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.FunctionName != nil {
		v := *s.FunctionName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "FunctionName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "Name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/lambda-2015-03-31/DeleteAliasOutput
type DeleteAliasOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteAliasOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteAliasOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opDeleteAlias = "DeleteAlias"

// DeleteAliasRequest returns a request value for making API operation for
// AWS Lambda.
//
// Deletes a Lambda function alias (https://docs.aws.amazon.com/lambda/latest/dg/versioning-aliases.html).
//
//    // Example sending a request using DeleteAliasRequest.
//    req := client.DeleteAliasRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/lambda-2015-03-31/DeleteAlias
func (c *Client) DeleteAliasRequest(input *DeleteAliasInput) DeleteAliasRequest {
	op := &aws.Operation{
		Name:       opDeleteAlias,
		HTTPMethod: "DELETE",
		HTTPPath:   "/2015-03-31/functions/{FunctionName}/aliases/{Name}",
	}

	if input == nil {
		input = &DeleteAliasInput{}
	}

	req := c.newRequest(op, input, &DeleteAliasOutput{})
	req.Handlers.Unmarshal.Remove(restjson.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return DeleteAliasRequest{Request: req, Input: input, Copy: c.DeleteAliasRequest}
}

// DeleteAliasRequest is the request type for the
// DeleteAlias API operation.
type DeleteAliasRequest struct {
	*aws.Request
	Input *DeleteAliasInput
	Copy  func(*DeleteAliasInput) DeleteAliasRequest
}

// Send marshals and sends the DeleteAlias API request.
func (r DeleteAliasRequest) Send(ctx context.Context) (*DeleteAliasResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteAliasResponse{
		DeleteAliasOutput: r.Request.Data.(*DeleteAliasOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteAliasResponse is the response type for the
// DeleteAlias API operation.
type DeleteAliasResponse struct {
	*DeleteAliasOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteAlias request.
func (r *DeleteAliasResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
