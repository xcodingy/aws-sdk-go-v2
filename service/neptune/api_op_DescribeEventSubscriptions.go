// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package neptune

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/neptune-2014-10-31/DescribeEventSubscriptionsMessage
type DescribeEventSubscriptionsInput struct {
	_ struct{} `type:"structure"`

	// This parameter is not currently supported.
	Filters []Filter `locationNameList:"Filter" type:"list"`

	// An optional pagination token provided by a previous DescribeOrderableDBInstanceOptions
	// request. If this parameter is specified, the response includes only records
	// beyond the marker, up to the value specified by MaxRecords .
	Marker *string `type:"string"`

	// The maximum number of records to include in the response. If more records
	// exist than the specified MaxRecords value, a pagination token called a marker
	// is included in the response so that the remaining results can be retrieved.
	//
	// Default: 100
	//
	// Constraints: Minimum 20, maximum 100.
	MaxRecords *int64 `type:"integer"`

	// The name of the event notification subscription you want to describe.
	SubscriptionName *string `type:"string"`
}

// String returns the string representation
func (s DescribeEventSubscriptionsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeEventSubscriptionsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeEventSubscriptionsInput"}
	if s.Filters != nil {
		for i, v := range s.Filters {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Filters", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Data returned by the DescribeEventSubscriptions action.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/neptune-2014-10-31/EventSubscriptionsMessage
type DescribeEventSubscriptionsOutput struct {
	_ struct{} `type:"structure"`

	// A list of EventSubscriptions data types.
	EventSubscriptionsList []EventSubscription `locationNameList:"EventSubscription" type:"list"`

	// An optional pagination token provided by a previous DescribeOrderableDBInstanceOptions
	// request. If this parameter is specified, the response includes only records
	// beyond the marker, up to the value specified by MaxRecords.
	Marker *string `type:"string"`
}

// String returns the string representation
func (s DescribeEventSubscriptionsOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeEventSubscriptions = "DescribeEventSubscriptions"

// DescribeEventSubscriptionsRequest returns a request value for making API operation for
// Amazon Neptune.
//
// Lists all the subscription descriptions for a customer account. The description
// for a subscription includes SubscriptionName, SNSTopicARN, CustomerID, SourceType,
// SourceID, CreationTime, and Status.
//
// If you specify a SubscriptionName, lists the description for that subscription.
//
//    // Example sending a request using DescribeEventSubscriptionsRequest.
//    req := client.DescribeEventSubscriptionsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/neptune-2014-10-31/DescribeEventSubscriptions
func (c *Client) DescribeEventSubscriptionsRequest(input *DescribeEventSubscriptionsInput) DescribeEventSubscriptionsRequest {
	op := &aws.Operation{
		Name:       opDescribeEventSubscriptions,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"Marker"},
			OutputTokens:    []string{"Marker"},
			LimitToken:      "MaxRecords",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &DescribeEventSubscriptionsInput{}
	}

	req := c.newRequest(op, input, &DescribeEventSubscriptionsOutput{})
	return DescribeEventSubscriptionsRequest{Request: req, Input: input, Copy: c.DescribeEventSubscriptionsRequest}
}

// DescribeEventSubscriptionsRequest is the request type for the
// DescribeEventSubscriptions API operation.
type DescribeEventSubscriptionsRequest struct {
	*aws.Request
	Input *DescribeEventSubscriptionsInput
	Copy  func(*DescribeEventSubscriptionsInput) DescribeEventSubscriptionsRequest
}

// Send marshals and sends the DescribeEventSubscriptions API request.
func (r DescribeEventSubscriptionsRequest) Send(ctx context.Context) (*DescribeEventSubscriptionsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeEventSubscriptionsResponse{
		DescribeEventSubscriptionsOutput: r.Request.Data.(*DescribeEventSubscriptionsOutput),
		response:                         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewDescribeEventSubscriptionsRequestPaginator returns a paginator for DescribeEventSubscriptions.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.DescribeEventSubscriptionsRequest(input)
//   p := neptune.NewDescribeEventSubscriptionsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewDescribeEventSubscriptionsPaginator(req DescribeEventSubscriptionsRequest) DescribeEventSubscriptionsPaginator {
	return DescribeEventSubscriptionsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *DescribeEventSubscriptionsInput
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

// DescribeEventSubscriptionsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type DescribeEventSubscriptionsPaginator struct {
	aws.Pager
}

func (p *DescribeEventSubscriptionsPaginator) CurrentPage() *DescribeEventSubscriptionsOutput {
	return p.Pager.CurrentPage().(*DescribeEventSubscriptionsOutput)
}

// DescribeEventSubscriptionsResponse is the response type for the
// DescribeEventSubscriptions API operation.
type DescribeEventSubscriptionsResponse struct {
	*DescribeEventSubscriptionsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeEventSubscriptions request.
func (r *DescribeEventSubscriptionsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
