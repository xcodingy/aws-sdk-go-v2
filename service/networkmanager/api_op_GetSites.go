// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package networkmanager

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type GetSitesInput struct {
	_ struct{} `type:"structure"`

	// The ID of the global network.
	//
	// GlobalNetworkId is a required field
	GlobalNetworkId *string `location:"uri" locationName:"globalNetworkId" type:"string" required:"true"`

	// The maximum number of results to return.
	MaxResults *int64 `location:"querystring" locationName:"maxResults" min:"1" type:"integer"`

	// The token for the next page of results.
	NextToken *string `location:"querystring" locationName:"nextToken" type:"string"`

	// One or more site IDs. The maximum is 10.
	SiteIds []string `location:"querystring" locationName:"siteIds" type:"list"`
}

// String returns the string representation
func (s GetSitesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetSitesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetSitesInput"}

	if s.GlobalNetworkId == nil {
		invalidParams.Add(aws.NewErrParamRequired("GlobalNetworkId"))
	}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetSitesInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.GlobalNetworkId != nil {
		v := *s.GlobalNetworkId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "globalNetworkId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.MaxResults != nil {
		v := *s.MaxResults

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "maxResults", protocol.Int64Value(v), metadata)
	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.SiteIds != nil {
		v := s.SiteIds

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.QueryTarget, "siteIds", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	return nil
}

type GetSitesOutput struct {
	_ struct{} `type:"structure"`

	// The token for the next page of results.
	NextToken *string `type:"string"`

	// The sites.
	Sites []Site `type:"list"`
}

// String returns the string representation
func (s GetSitesOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetSitesOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "NextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Sites != nil {
		v := s.Sites

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "Sites", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	return nil
}

const opGetSites = "GetSites"

// GetSitesRequest returns a request value for making API operation for
// AWS Network Manager.
//
// Gets information about one or more of your sites in a global network.
//
//    // Example sending a request using GetSitesRequest.
//    req := client.GetSitesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/networkmanager-2019-07-05/GetSites
func (c *Client) GetSitesRequest(input *GetSitesInput) GetSitesRequest {
	op := &aws.Operation{
		Name:       opGetSites,
		HTTPMethod: "GET",
		HTTPPath:   "/global-networks/{globalNetworkId}/sites",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &GetSitesInput{}
	}

	req := c.newRequest(op, input, &GetSitesOutput{})

	return GetSitesRequest{Request: req, Input: input, Copy: c.GetSitesRequest}
}

// GetSitesRequest is the request type for the
// GetSites API operation.
type GetSitesRequest struct {
	*aws.Request
	Input *GetSitesInput
	Copy  func(*GetSitesInput) GetSitesRequest
}

// Send marshals and sends the GetSites API request.
func (r GetSitesRequest) Send(ctx context.Context) (*GetSitesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetSitesResponse{
		GetSitesOutput: r.Request.Data.(*GetSitesOutput),
		response:       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewGetSitesRequestPaginator returns a paginator for GetSites.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.GetSitesRequest(input)
//   p := networkmanager.NewGetSitesRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewGetSitesPaginator(req GetSitesRequest) GetSitesPaginator {
	return GetSitesPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *GetSitesInput
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

// GetSitesPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type GetSitesPaginator struct {
	aws.Pager
}

func (p *GetSitesPaginator) CurrentPage() *GetSitesOutput {
	return p.Pager.CurrentPage().(*GetSitesOutput)
}

// GetSitesResponse is the response type for the
// GetSites API operation.
type GetSitesResponse struct {
	*GetSitesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetSites request.
func (r *GetSitesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
