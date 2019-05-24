// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package glue

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/glue-2017-03-31/GetUserDefinedFunctionRequest
type GetUserDefinedFunctionInput struct {
	_ struct{} `type:"structure"`

	// The ID of the Data Catalog where the function to be retrieved is located.
	// If none is supplied, the AWS account ID is used by default.
	CatalogId *string `min:"1" type:"string"`

	// The name of the catalog database where the function is located.
	//
	// DatabaseName is a required field
	DatabaseName *string `min:"1" type:"string" required:"true"`

	// The name of the function.
	//
	// FunctionName is a required field
	FunctionName *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s GetUserDefinedFunctionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetUserDefinedFunctionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetUserDefinedFunctionInput"}
	if s.CatalogId != nil && len(*s.CatalogId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("CatalogId", 1))
	}

	if s.DatabaseName == nil {
		invalidParams.Add(aws.NewErrParamRequired("DatabaseName"))
	}
	if s.DatabaseName != nil && len(*s.DatabaseName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("DatabaseName", 1))
	}

	if s.FunctionName == nil {
		invalidParams.Add(aws.NewErrParamRequired("FunctionName"))
	}
	if s.FunctionName != nil && len(*s.FunctionName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("FunctionName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/glue-2017-03-31/GetUserDefinedFunctionResponse
type GetUserDefinedFunctionOutput struct {
	_ struct{} `type:"structure"`

	// The requested function definition.
	UserDefinedFunction *UserDefinedFunction `type:"structure"`
}

// String returns the string representation
func (s GetUserDefinedFunctionOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetUserDefinedFunction = "GetUserDefinedFunction"

// GetUserDefinedFunctionRequest returns a request value for making API operation for
// AWS Glue.
//
// Retrieves a specified function definition from the Data Catalog.
//
//    // Example sending a request using GetUserDefinedFunctionRequest.
//    req := client.GetUserDefinedFunctionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/glue-2017-03-31/GetUserDefinedFunction
func (c *Client) GetUserDefinedFunctionRequest(input *GetUserDefinedFunctionInput) GetUserDefinedFunctionRequest {
	op := &aws.Operation{
		Name:       opGetUserDefinedFunction,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetUserDefinedFunctionInput{}
	}

	req := c.newRequest(op, input, &GetUserDefinedFunctionOutput{})
	return GetUserDefinedFunctionRequest{Request: req, Input: input, Copy: c.GetUserDefinedFunctionRequest}
}

// GetUserDefinedFunctionRequest is the request type for the
// GetUserDefinedFunction API operation.
type GetUserDefinedFunctionRequest struct {
	*aws.Request
	Input *GetUserDefinedFunctionInput
	Copy  func(*GetUserDefinedFunctionInput) GetUserDefinedFunctionRequest
}

// Send marshals and sends the GetUserDefinedFunction API request.
func (r GetUserDefinedFunctionRequest) Send(ctx context.Context) (*GetUserDefinedFunctionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetUserDefinedFunctionResponse{
		GetUserDefinedFunctionOutput: r.Request.Data.(*GetUserDefinedFunctionOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetUserDefinedFunctionResponse is the response type for the
// GetUserDefinedFunction API operation.
type GetUserDefinedFunctionResponse struct {
	*GetUserDefinedFunctionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetUserDefinedFunction request.
func (r *GetUserDefinedFunctionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
