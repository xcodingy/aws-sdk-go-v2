// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package apigateway

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Request to list existing Models defined for a RestApi resource.
type GetModelsInput struct {
	_ struct{} `type:"structure"`

	// The maximum number of returned results per page. The default value is 25
	// and the maximum value is 500.
	Limit *int64 `location:"querystring" locationName:"limit" type:"integer"`

	// The current pagination position in the paged result set.
	Position *string `location:"querystring" locationName:"position" type:"string"`

	// [Required] The string identifier of the associated RestApi.
	//
	// RestApiId is a required field
	RestApiId *string `location:"uri" locationName:"restapi_id" type:"string" required:"true"`
}

// String returns the string representation
func (s GetModelsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetModelsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetModelsInput"}

	if s.RestApiId == nil {
		invalidParams.Add(aws.NewErrParamRequired("RestApiId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetModelsInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.RestApiId != nil {
		v := *s.RestApiId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "restapi_id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Limit != nil {
		v := *s.Limit

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "limit", protocol.Int64Value(v), metadata)
	}
	if s.Position != nil {
		v := *s.Position

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "position", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Represents a collection of Model resources.
//
// Method, MethodResponse, Models and Mappings (https://docs.aws.amazon.com/apigateway/latest/developerguide/models-mappings.html)
type GetModelsOutput struct {
	_ struct{} `type:"structure"`

	// The current page of elements from this collection.
	Items []Model `locationName:"item" type:"list"`

	Position *string `locationName:"position" type:"string"`
}

// String returns the string representation
func (s GetModelsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetModelsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if len(s.Items) > 0 {
		v := s.Items

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "item", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.Position != nil {
		v := *s.Position

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "position", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opGetModels = "GetModels"

// GetModelsRequest returns a request value for making API operation for
// Amazon API Gateway.
//
// Describes existing Models defined for a RestApi resource.
//
//    // Example sending a request using GetModelsRequest.
//    req := client.GetModelsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) GetModelsRequest(input *GetModelsInput) GetModelsRequest {
	op := &aws.Operation{
		Name:       opGetModels,
		HTTPMethod: "GET",
		HTTPPath:   "/restapis/{restapi_id}/models",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"position"},
			OutputTokens:    []string{"position"},
			LimitToken:      "limit",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &GetModelsInput{}
	}

	req := c.newRequest(op, input, &GetModelsOutput{})
	return GetModelsRequest{Request: req, Input: input, Copy: c.GetModelsRequest}
}

// GetModelsRequest is the request type for the
// GetModels API operation.
type GetModelsRequest struct {
	*aws.Request
	Input *GetModelsInput
	Copy  func(*GetModelsInput) GetModelsRequest
}

// Send marshals and sends the GetModels API request.
func (r GetModelsRequest) Send(ctx context.Context) (*GetModelsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetModelsResponse{
		GetModelsOutput: r.Request.Data.(*GetModelsOutput),
		response:        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewGetModelsRequestPaginator returns a paginator for GetModels.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.GetModelsRequest(input)
//   p := apigateway.NewGetModelsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewGetModelsPaginator(req GetModelsRequest) GetModelsPaginator {
	return GetModelsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *GetModelsInput
				if req.Input != nil {
					tmp := *req.Input
					inCpy = &tmp
				}

				newReq := req.Copy(inCpy)
				newReq.SetContext(ctx)
				return newReq.Request, nil
			},
		},
	}
}

// GetModelsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type GetModelsPaginator struct {
	aws.Pager
}

func (p *GetModelsPaginator) CurrentPage() *GetModelsOutput {
	return p.Pager.CurrentPage().(*GetModelsOutput)
}

// GetModelsResponse is the response type for the
// GetModels API operation.
type GetModelsResponse struct {
	*GetModelsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetModels request.
func (r *GetModelsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
