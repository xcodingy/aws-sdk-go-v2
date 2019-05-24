// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package glue

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/glue-2017-03-31/GetMappingRequest
type GetMappingInput struct {
	_ struct{} `type:"structure"`

	// Parameters for the mapping.
	Location *Location `type:"structure"`

	// A list of target tables.
	Sinks []CatalogEntry `type:"list"`

	// Specifies the source table.
	//
	// Source is a required field
	Source *CatalogEntry `type:"structure" required:"true"`
}

// String returns the string representation
func (s GetMappingInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetMappingInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetMappingInput"}

	if s.Source == nil {
		invalidParams.Add(aws.NewErrParamRequired("Source"))
	}
	if s.Location != nil {
		if err := s.Location.Validate(); err != nil {
			invalidParams.AddNested("Location", err.(aws.ErrInvalidParams))
		}
	}
	if s.Sinks != nil {
		for i, v := range s.Sinks {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Sinks", i), err.(aws.ErrInvalidParams))
			}
		}
	}
	if s.Source != nil {
		if err := s.Source.Validate(); err != nil {
			invalidParams.AddNested("Source", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/glue-2017-03-31/GetMappingResponse
type GetMappingOutput struct {
	_ struct{} `type:"structure"`

	// A list of mappings to the specified targets.
	//
	// Mapping is a required field
	Mapping []MappingEntry `type:"list" required:"true"`
}

// String returns the string representation
func (s GetMappingOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetMapping = "GetMapping"

// GetMappingRequest returns a request value for making API operation for
// AWS Glue.
//
// Creates mappings.
//
//    // Example sending a request using GetMappingRequest.
//    req := client.GetMappingRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/glue-2017-03-31/GetMapping
func (c *Client) GetMappingRequest(input *GetMappingInput) GetMappingRequest {
	op := &aws.Operation{
		Name:       opGetMapping,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetMappingInput{}
	}

	req := c.newRequest(op, input, &GetMappingOutput{})
	return GetMappingRequest{Request: req, Input: input, Copy: c.GetMappingRequest}
}

// GetMappingRequest is the request type for the
// GetMapping API operation.
type GetMappingRequest struct {
	*aws.Request
	Input *GetMappingInput
	Copy  func(*GetMappingInput) GetMappingRequest
}

// Send marshals and sends the GetMapping API request.
func (r GetMappingRequest) Send(ctx context.Context) (*GetMappingResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetMappingResponse{
		GetMappingOutput: r.Request.Data.(*GetMappingOutput),
		response:         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetMappingResponse is the response type for the
// GetMapping API operation.
type GetMappingResponse struct {
	*GetMappingOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetMapping request.
func (r *GetMappingResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
