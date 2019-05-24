// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package docdb

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents the input to DescribeEvents.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/docdb-2014-10-31/DescribeEventsMessage
type DescribeEventsInput struct {
	_ struct{} `type:"structure"`

	// The number of minutes to retrieve events for.
	//
	// Default: 60
	Duration *int64 `type:"integer"`

	// The end of the time interval for which to retrieve events, specified in ISO
	// 8601 format.
	//
	// Example: 2009-07-08T18:00Z
	EndTime *time.Time `type:"timestamp" timestampFormat:"iso8601"`

	// A list of event categories that trigger notifications for an event notification
	// subscription.
	EventCategories []string `locationNameList:"EventCategory" type:"list"`

	// This parameter is not currently supported.
	Filters []Filter `locationNameList:"Filter" type:"list"`

	// An optional pagination token provided by a previous request. If this parameter
	// is specified, the response includes only records beyond the marker, up to
	// the value specified by MaxRecords.
	Marker *string `type:"string"`

	// The maximum number of records to include in the response. If more records
	// exist than the specified MaxRecords value, a pagination token (marker) is
	// included in the response so that the remaining results can be retrieved.
	//
	// Default: 100
	//
	// Constraints: Minimum 20, maximum 100.
	MaxRecords *int64 `type:"integer"`

	// The identifier of the event source for which events are returned. If not
	// specified, then all sources are included in the response.
	//
	// Constraints:
	//
	//    * If SourceIdentifier is provided, SourceType must also be provided.
	//
	//    * If the source type is DBInstance, a DBInstanceIdentifier must be provided.
	//
	//    * If the source type is DBSecurityGroup, a DBSecurityGroupName must be
	//    provided.
	//
	//    * If the source type is DBParameterGroup, a DBParameterGroupName must
	//    be provided.
	//
	//    * If the source type is DBSnapshot, a DBSnapshotIdentifier must be provided.
	//
	//    * Cannot end with a hyphen or contain two consecutive hyphens.
	SourceIdentifier *string `type:"string"`

	// The event source to retrieve events for. If no value is specified, all events
	// are returned.
	SourceType SourceType `type:"string" enum:"true"`

	// The beginning of the time interval to retrieve events for, specified in ISO
	// 8601 format.
	//
	// Example: 2009-07-08T18:00Z
	StartTime *time.Time `type:"timestamp" timestampFormat:"iso8601"`
}

// String returns the string representation
func (s DescribeEventsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeEventsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeEventsInput"}
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

// Represents the output of DescribeEvents.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/docdb-2014-10-31/EventsMessage
type DescribeEventsOutput struct {
	_ struct{} `type:"structure"`

	// Detailed information about one or more events.
	Events []Event `locationNameList:"Event" type:"list"`

	// An optional pagination token provided by a previous request. If this parameter
	// is specified, the response includes only records beyond the marker, up to
	// the value specified by MaxRecords.
	Marker *string `type:"string"`
}

// String returns the string representation
func (s DescribeEventsOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeEvents = "DescribeEvents"

// DescribeEventsRequest returns a request value for making API operation for
// Amazon DocumentDB with MongoDB compatibility.
//
// Returns events related to DB instances, DB security groups, DB snapshots,
// and DB parameter groups for the past 14 days. You can obtain events specific
// to a particular DB instance, DB security group, DB snapshot, or DB parameter
// group by providing the name as a parameter. By default, the events of the
// past hour are returned.
//
//    // Example sending a request using DescribeEventsRequest.
//    req := client.DescribeEventsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/docdb-2014-10-31/DescribeEvents
func (c *Client) DescribeEventsRequest(input *DescribeEventsInput) DescribeEventsRequest {
	op := &aws.Operation{
		Name:       opDescribeEvents,
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
		input = &DescribeEventsInput{}
	}

	req := c.newRequest(op, input, &DescribeEventsOutput{})
	return DescribeEventsRequest{Request: req, Input: input, Copy: c.DescribeEventsRequest}
}

// DescribeEventsRequest is the request type for the
// DescribeEvents API operation.
type DescribeEventsRequest struct {
	*aws.Request
	Input *DescribeEventsInput
	Copy  func(*DescribeEventsInput) DescribeEventsRequest
}

// Send marshals and sends the DescribeEvents API request.
func (r DescribeEventsRequest) Send(ctx context.Context) (*DescribeEventsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeEventsResponse{
		DescribeEventsOutput: r.Request.Data.(*DescribeEventsOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewDescribeEventsRequestPaginator returns a paginator for DescribeEvents.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.DescribeEventsRequest(input)
//   p := docdb.NewDescribeEventsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewDescribeEventsPaginator(req DescribeEventsRequest) DescribeEventsPaginator {
	return DescribeEventsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *DescribeEventsInput
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

// DescribeEventsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type DescribeEventsPaginator struct {
	aws.Pager
}

func (p *DescribeEventsPaginator) CurrentPage() *DescribeEventsOutput {
	return p.Pager.CurrentPage().(*DescribeEventsOutput)
}

// DescribeEventsResponse is the response type for the
// DescribeEvents API operation.
type DescribeEventsResponse struct {
	*DescribeEventsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeEvents request.
func (r *DescribeEventsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
