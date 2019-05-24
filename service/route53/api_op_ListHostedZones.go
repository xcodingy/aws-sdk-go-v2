// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package route53

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// A request to retrieve a list of the public and private hosted zones that
// are associated with the current AWS account.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/route53-2013-04-01/ListHostedZonesRequest
type ListHostedZonesInput struct {
	_ struct{} `type:"structure"`

	// If you're using reusable delegation sets and you want to list all of the
	// hosted zones that are associated with a reusable delegation set, specify
	// the ID of that reusable delegation set.
	DelegationSetId *string `location:"querystring" locationName:"delegationsetid" type:"string"`

	// If the value of IsTruncated in the previous response was true, you have more
	// hosted zones. To get more hosted zones, submit another ListHostedZones request.
	//
	// For the value of marker, specify the value of NextMarker from the previous
	// response, which is the ID of the first hosted zone that Amazon Route 53 will
	// return if you submit another request.
	//
	// If the value of IsTruncated in the previous response was false, there are
	// no more hosted zones to get.
	Marker *string `location:"querystring" locationName:"marker" type:"string"`

	// (Optional) The maximum number of hosted zones that you want Amazon Route
	// 53 to return. If you have more than maxitems hosted zones, the value of IsTruncated
	// in the response is true, and the value of NextMarker is the hosted zone ID
	// of the first hosted zone that Route 53 will return if you submit another
	// request.
	MaxItems *string `location:"querystring" locationName:"maxitems" type:"string"`
}

// String returns the string representation
func (s ListHostedZonesInput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListHostedZonesInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.DelegationSetId != nil {
		v := *s.DelegationSetId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "delegationsetid", protocol.StringValue(v), metadata)
	}
	if s.Marker != nil {
		v := *s.Marker

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "marker", protocol.StringValue(v), metadata)
	}
	if s.MaxItems != nil {
		v := *s.MaxItems

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "maxitems", protocol.StringValue(v), metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/route53-2013-04-01/ListHostedZonesResponse
type ListHostedZonesOutput struct {
	_ struct{} `type:"structure"`

	// A complex type that contains general information about the hosted zone.
	//
	// HostedZones is a required field
	HostedZones []HostedZone `locationNameList:"HostedZone" type:"list" required:"true"`

	// A flag indicating whether there are more hosted zones to be listed. If the
	// response was truncated, you can get more hosted zones by submitting another
	// ListHostedZones request and specifying the value of NextMarker in the marker
	// parameter.
	//
	// IsTruncated is a required field
	IsTruncated *bool `type:"boolean" required:"true"`

	// For the second and subsequent calls to ListHostedZones, Marker is the value
	// that you specified for the marker parameter in the request that produced
	// the current response.
	//
	// Marker is a required field
	Marker *string `type:"string" required:"true"`

	// The value that you specified for the maxitems parameter in the call to ListHostedZones
	// that produced the current response.
	//
	// MaxItems is a required field
	MaxItems *string `type:"string" required:"true"`

	// If IsTruncated is true, the value of NextMarker identifies the first hosted
	// zone in the next group of hosted zones. Submit another ListHostedZones request,
	// and specify the value of NextMarker from the response in the marker parameter.
	//
	// This element is present only if IsTruncated is true.
	NextMarker *string `type:"string"`
}

// String returns the string representation
func (s ListHostedZonesOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListHostedZonesOutput) MarshalFields(e protocol.FieldEncoder) error {
	if len(s.HostedZones) > 0 {
		v := s.HostedZones

		metadata := protocol.Metadata{ListLocationName: "HostedZone"}
		ls0 := e.List(protocol.BodyTarget, "HostedZones", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.IsTruncated != nil {
		v := *s.IsTruncated

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "IsTruncated", protocol.BoolValue(v), metadata)
	}
	if s.Marker != nil {
		v := *s.Marker

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Marker", protocol.StringValue(v), metadata)
	}
	if s.MaxItems != nil {
		v := *s.MaxItems

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "MaxItems", protocol.StringValue(v), metadata)
	}
	if s.NextMarker != nil {
		v := *s.NextMarker

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "NextMarker", protocol.StringValue(v), metadata)
	}
	return nil
}

const opListHostedZones = "ListHostedZones"

// ListHostedZonesRequest returns a request value for making API operation for
// Amazon Route 53.
//
// Retrieves a list of the public and private hosted zones that are associated
// with the current AWS account. The response includes a HostedZones child element
// for each hosted zone.
//
// Amazon Route 53 returns a maximum of 100 items in each response. If you have
// a lot of hosted zones, you can use the maxitems parameter to list them in
// groups of up to 100.
//
//    // Example sending a request using ListHostedZonesRequest.
//    req := client.ListHostedZonesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/route53-2013-04-01/ListHostedZones
func (c *Client) ListHostedZonesRequest(input *ListHostedZonesInput) ListHostedZonesRequest {
	op := &aws.Operation{
		Name:       opListHostedZones,
		HTTPMethod: "GET",
		HTTPPath:   "/2013-04-01/hostedzone",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"Marker"},
			OutputTokens:    []string{"NextMarker"},
			LimitToken:      "MaxItems",
			TruncationToken: "IsTruncated",
		},
	}

	if input == nil {
		input = &ListHostedZonesInput{}
	}

	req := c.newRequest(op, input, &ListHostedZonesOutput{})
	return ListHostedZonesRequest{Request: req, Input: input, Copy: c.ListHostedZonesRequest}
}

// ListHostedZonesRequest is the request type for the
// ListHostedZones API operation.
type ListHostedZonesRequest struct {
	*aws.Request
	Input *ListHostedZonesInput
	Copy  func(*ListHostedZonesInput) ListHostedZonesRequest
}

// Send marshals and sends the ListHostedZones API request.
func (r ListHostedZonesRequest) Send(ctx context.Context) (*ListHostedZonesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListHostedZonesResponse{
		ListHostedZonesOutput: r.Request.Data.(*ListHostedZonesOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListHostedZonesRequestPaginator returns a paginator for ListHostedZones.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListHostedZonesRequest(input)
//   p := route53.NewListHostedZonesRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListHostedZonesPaginator(req ListHostedZonesRequest) ListHostedZonesPaginator {
	return ListHostedZonesPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListHostedZonesInput
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

// ListHostedZonesPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListHostedZonesPaginator struct {
	aws.Pager
}

func (p *ListHostedZonesPaginator) CurrentPage() *ListHostedZonesOutput {
	return p.Pager.CurrentPage().(*ListHostedZonesOutput)
}

// ListHostedZonesResponse is the response type for the
// ListHostedZones API operation.
type ListHostedZonesResponse struct {
	*ListHostedZonesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListHostedZones request.
func (r *ListHostedZonesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
