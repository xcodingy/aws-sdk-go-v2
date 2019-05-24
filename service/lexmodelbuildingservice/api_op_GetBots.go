// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package lexmodelbuildingservice

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/lex-models-2017-04-19/GetBotsRequest
type GetBotsInput struct {
	_ struct{} `type:"structure"`

	// The maximum number of bots to return in the response that the request will
	// return. The default is 10.
	MaxResults *int64 `location:"querystring" locationName:"maxResults" min:"1" type:"integer"`

	// Substring to match in bot names. A bot will be returned if any part of its
	// name matches the substring. For example, "xyz" matches both "xyzabc" and
	// "abcxyz."
	NameContains *string `location:"querystring" locationName:"nameContains" min:"2" type:"string"`

	// A pagination token that fetches the next page of bots. If the response to
	// this call is truncated, Amazon Lex returns a pagination token in the response.
	// To fetch the next page of bots, specify the pagination token in the next
	// request.
	NextToken *string `location:"querystring" locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s GetBotsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetBotsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetBotsInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}
	if s.NameContains != nil && len(*s.NameContains) < 2 {
		invalidParams.Add(aws.NewErrParamMinLen("NameContains", 2))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetBotsInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	if s.MaxResults != nil {
		v := *s.MaxResults

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "maxResults", protocol.Int64Value(v), metadata)
	}
	if s.NameContains != nil {
		v := *s.NameContains

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "nameContains", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/lex-models-2017-04-19/GetBotsResponse
type GetBotsOutput struct {
	_ struct{} `type:"structure"`

	// An array of botMetadata objects, with one entry for each bot.
	Bots []BotMetadata `locationName:"bots" type:"list"`

	// If the response is truncated, it includes a pagination token that you can
	// specify in your next request to fetch the next page of bots.
	NextToken *string `locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s GetBotsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetBotsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if len(s.Bots) > 0 {
		v := s.Bots

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "bots", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opGetBots = "GetBots"

// GetBotsRequest returns a request value for making API operation for
// Amazon Lex Model Building Service.
//
// Returns bot information as follows:
//
//    * If you provide the nameContains field, the response includes information
//    for the $LATEST version of all bots whose name contains the specified
//    string.
//
//    * If you don't specify the nameContains field, the operation returns information
//    about the $LATEST version of all of your bots.
//
// This operation requires permission for the lex:GetBots action.
//
//    // Example sending a request using GetBotsRequest.
//    req := client.GetBotsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/lex-models-2017-04-19/GetBots
func (c *Client) GetBotsRequest(input *GetBotsInput) GetBotsRequest {
	op := &aws.Operation{
		Name:       opGetBots,
		HTTPMethod: "GET",
		HTTPPath:   "/bots/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"nextToken"},
			OutputTokens:    []string{"nextToken"},
			LimitToken:      "maxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &GetBotsInput{}
	}

	req := c.newRequest(op, input, &GetBotsOutput{})
	return GetBotsRequest{Request: req, Input: input, Copy: c.GetBotsRequest}
}

// GetBotsRequest is the request type for the
// GetBots API operation.
type GetBotsRequest struct {
	*aws.Request
	Input *GetBotsInput
	Copy  func(*GetBotsInput) GetBotsRequest
}

// Send marshals and sends the GetBots API request.
func (r GetBotsRequest) Send(ctx context.Context) (*GetBotsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetBotsResponse{
		GetBotsOutput: r.Request.Data.(*GetBotsOutput),
		response:      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewGetBotsRequestPaginator returns a paginator for GetBots.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.GetBotsRequest(input)
//   p := lexmodelbuildingservice.NewGetBotsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewGetBotsPaginator(req GetBotsRequest) GetBotsPaginator {
	return GetBotsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *GetBotsInput
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

// GetBotsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type GetBotsPaginator struct {
	aws.Pager
}

func (p *GetBotsPaginator) CurrentPage() *GetBotsOutput {
	return p.Pager.CurrentPage().(*GetBotsOutput)
}

// GetBotsResponse is the response type for the
// GetBots API operation.
type GetBotsResponse struct {
	*GetBotsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetBots request.
func (r *GetBotsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
