// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package pinpointemail

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// A request to obtain more information about dedicated IP pools.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-email-2018-07-26/GetDedicatedIpsRequest
type GetDedicatedIpsInput struct {
	_ struct{} `type:"structure"`

	// A token returned from a previous call to GetDedicatedIps to indicate the
	// position of the dedicated IP pool in the list of IP pools.
	NextToken *string `type:"string"`

	// The number of results to show in a single call to GetDedicatedIpsRequest.
	// If the number of results is larger than the number you specified in this
	// parameter, then the response includes a NextToken element, which you can
	// use to obtain additional results.
	PageSize *int64 `type:"integer"`

	// The name of the IP pool that the dedicated IP address is associated with.
	PoolName *string `type:"string"`
}

// String returns the string representation
func (s GetDedicatedIpsInput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetDedicatedIpsInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "NextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.PageSize != nil {
		v := *s.PageSize

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "PageSize", protocol.Int64Value(v), metadata)
	}
	if s.PoolName != nil {
		v := *s.PoolName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "PoolName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Information about the dedicated IP addresses that are associated with your
// Amazon Pinpoint account.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-email-2018-07-26/GetDedicatedIpsResponse
type GetDedicatedIpsOutput struct {
	_ struct{} `type:"structure"`

	// A list of dedicated IP addresses that are reserved for use by your Amazon
	// Pinpoint account.
	DedicatedIps []DedicatedIp `type:"list"`

	// A token that indicates that there are additional dedicated IP addresses to
	// list. To view additional addresses, issue another request to GetDedicatedIps,
	// passing this token in the NextToken parameter.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s GetDedicatedIpsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetDedicatedIpsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if len(s.DedicatedIps) > 0 {
		v := s.DedicatedIps

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "DedicatedIps", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "NextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opGetDedicatedIps = "GetDedicatedIps"

// GetDedicatedIpsRequest returns a request value for making API operation for
// Amazon Pinpoint Email Service.
//
// List the dedicated IP addresses that are associated with your Amazon Pinpoint
// account.
//
//    // Example sending a request using GetDedicatedIpsRequest.
//    req := client.GetDedicatedIpsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-email-2018-07-26/GetDedicatedIps
func (c *Client) GetDedicatedIpsRequest(input *GetDedicatedIpsInput) GetDedicatedIpsRequest {
	op := &aws.Operation{
		Name:       opGetDedicatedIps,
		HTTPMethod: "GET",
		HTTPPath:   "/v1/email/dedicated-ips",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "PageSize",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &GetDedicatedIpsInput{}
	}

	req := c.newRequest(op, input, &GetDedicatedIpsOutput{})
	return GetDedicatedIpsRequest{Request: req, Input: input, Copy: c.GetDedicatedIpsRequest}
}

// GetDedicatedIpsRequest is the request type for the
// GetDedicatedIps API operation.
type GetDedicatedIpsRequest struct {
	*aws.Request
	Input *GetDedicatedIpsInput
	Copy  func(*GetDedicatedIpsInput) GetDedicatedIpsRequest
}

// Send marshals and sends the GetDedicatedIps API request.
func (r GetDedicatedIpsRequest) Send(ctx context.Context) (*GetDedicatedIpsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetDedicatedIpsResponse{
		GetDedicatedIpsOutput: r.Request.Data.(*GetDedicatedIpsOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewGetDedicatedIpsRequestPaginator returns a paginator for GetDedicatedIps.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.GetDedicatedIpsRequest(input)
//   p := pinpointemail.NewGetDedicatedIpsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewGetDedicatedIpsPaginator(req GetDedicatedIpsRequest) GetDedicatedIpsPaginator {
	return GetDedicatedIpsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *GetDedicatedIpsInput
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

// GetDedicatedIpsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type GetDedicatedIpsPaginator struct {
	aws.Pager
}

func (p *GetDedicatedIpsPaginator) CurrentPage() *GetDedicatedIpsOutput {
	return p.Pager.CurrentPage().(*GetDedicatedIpsOutput)
}

// GetDedicatedIpsResponse is the response type for the
// GetDedicatedIps API operation.
type GetDedicatedIpsResponse struct {
	*GetDedicatedIpsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetDedicatedIps request.
func (r *GetDedicatedIpsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
