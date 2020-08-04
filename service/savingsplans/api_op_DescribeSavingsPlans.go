// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package savingsplans

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type DescribeSavingsPlansInput struct {
	_ struct{} `type:"structure"`

	// The filters.
	Filters []SavingsPlanFilter `locationName:"filters" type:"list"`

	// The maximum number of results to return with a single call. To retrieve additional
	// results, make another call with the returned token value.
	MaxResults *int64 `locationName:"maxResults" min:"1" type:"integer"`

	// The token for the next page of results.
	NextToken *string `locationName:"nextToken" type:"string"`

	// The Amazon Resource Names (ARN) of the Savings Plans.
	SavingsPlanArns []string `locationName:"savingsPlanArns" type:"list"`

	// The IDs of the Savings Plans.
	SavingsPlanIds []string `locationName:"savingsPlanIds" type:"list"`

	// The states.
	States []SavingsPlanState `locationName:"states" type:"list"`
}

// String returns the string representation
func (s DescribeSavingsPlansInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeSavingsPlansInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeSavingsPlansInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeSavingsPlansInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.Filters != nil {
		v := s.Filters

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "filters", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.MaxResults != nil {
		v := *s.MaxResults

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "maxResults", protocol.Int64Value(v), metadata)
	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.SavingsPlanArns != nil {
		v := s.SavingsPlanArns

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "savingsPlanArns", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	if s.SavingsPlanIds != nil {
		v := s.SavingsPlanIds

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "savingsPlanIds", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	if s.States != nil {
		v := s.States

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "states", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	return nil
}

type DescribeSavingsPlansOutput struct {
	_ struct{} `type:"structure"`

	// The token to use to retrieve the next page of results. This value is null
	// when there are no more results to return.
	NextToken *string `locationName:"nextToken" type:"string"`

	// Information about the Savings Plans.
	SavingsPlans []SavingsPlan `locationName:"savingsPlans" type:"list"`
}

// String returns the string representation
func (s DescribeSavingsPlansOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeSavingsPlansOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.SavingsPlans != nil {
		v := s.SavingsPlans

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "savingsPlans", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	return nil
}

const opDescribeSavingsPlans = "DescribeSavingsPlans"

// DescribeSavingsPlansRequest returns a request value for making API operation for
// AWS Savings Plans.
//
// Describes the specified Savings Plans.
//
//    // Example sending a request using DescribeSavingsPlansRequest.
//    req := client.DescribeSavingsPlansRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/savingsplans-2019-06-28/DescribeSavingsPlans
func (c *Client) DescribeSavingsPlansRequest(input *DescribeSavingsPlansInput) DescribeSavingsPlansRequest {
	op := &aws.Operation{
		Name:       opDescribeSavingsPlans,
		HTTPMethod: "POST",
		HTTPPath:   "/DescribeSavingsPlans",
	}

	if input == nil {
		input = &DescribeSavingsPlansInput{}
	}

	req := c.newRequest(op, input, &DescribeSavingsPlansOutput{})

	return DescribeSavingsPlansRequest{Request: req, Input: input, Copy: c.DescribeSavingsPlansRequest}
}

// DescribeSavingsPlansRequest is the request type for the
// DescribeSavingsPlans API operation.
type DescribeSavingsPlansRequest struct {
	*aws.Request
	Input *DescribeSavingsPlansInput
	Copy  func(*DescribeSavingsPlansInput) DescribeSavingsPlansRequest
}

// Send marshals and sends the DescribeSavingsPlans API request.
func (r DescribeSavingsPlansRequest) Send(ctx context.Context) (*DescribeSavingsPlansResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeSavingsPlansResponse{
		DescribeSavingsPlansOutput: r.Request.Data.(*DescribeSavingsPlansOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeSavingsPlansResponse is the response type for the
// DescribeSavingsPlans API operation.
type DescribeSavingsPlansResponse struct {
	*DescribeSavingsPlansOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeSavingsPlans request.
func (r *DescribeSavingsPlansResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}