// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DescribeAvailabilityZonesRequest
type DescribeAvailabilityZonesInput struct {
	_ struct{} `type:"structure"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `locationName:"dryRun" type:"boolean"`

	// The filters.
	//
	//    * message - Information about the Availability Zone.
	//
	//    * region-name - The name of the region for the Availability Zone (for
	//    example, us-east-1).
	//
	//    * state - The state of the Availability Zone (available | information
	//    | impaired | unavailable).
	//
	//    * zone-id - The ID of the Availability Zone (for example, use1-az1).
	//
	//    * zone-name - The name of the Availability Zone (for example, us-east-1a).
	Filters []Filter `locationName:"Filter" locationNameList:"Filter" type:"list"`

	// The IDs of the Availability Zones.
	ZoneIds []string `locationName:"ZoneId" locationNameList:"ZoneId" type:"list"`

	// The names of the Availability Zones.
	ZoneNames []string `locationName:"ZoneName" locationNameList:"ZoneName" type:"list"`
}

// String returns the string representation
func (s DescribeAvailabilityZonesInput) String() string {
	return awsutil.Prettify(s)
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DescribeAvailabilityZonesResult
type DescribeAvailabilityZonesOutput struct {
	_ struct{} `type:"structure"`

	// Information about the Availability Zones.
	AvailabilityZones []AvailabilityZone `locationName:"availabilityZoneInfo" locationNameList:"item" type:"list"`
}

// String returns the string representation
func (s DescribeAvailabilityZonesOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeAvailabilityZones = "DescribeAvailabilityZones"

// DescribeAvailabilityZonesRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Describes the Availability Zones that are available to you. The results include
// zones only for the region you're currently using. If there is an event impacting
// an Availability Zone, you can use this request to view the state and any
// provided message for that Availability Zone.
//
// For more information, see Regions and Availability Zones (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/using-regions-availability-zones.html)
// in the Amazon Elastic Compute Cloud User Guide.
//
//    // Example sending a request using DescribeAvailabilityZonesRequest.
//    req := client.DescribeAvailabilityZonesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DescribeAvailabilityZones
func (c *Client) DescribeAvailabilityZonesRequest(input *DescribeAvailabilityZonesInput) DescribeAvailabilityZonesRequest {
	op := &aws.Operation{
		Name:       opDescribeAvailabilityZones,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeAvailabilityZonesInput{}
	}

	req := c.newRequest(op, input, &DescribeAvailabilityZonesOutput{})
	return DescribeAvailabilityZonesRequest{Request: req, Input: input, Copy: c.DescribeAvailabilityZonesRequest}
}

// DescribeAvailabilityZonesRequest is the request type for the
// DescribeAvailabilityZones API operation.
type DescribeAvailabilityZonesRequest struct {
	*aws.Request
	Input *DescribeAvailabilityZonesInput
	Copy  func(*DescribeAvailabilityZonesInput) DescribeAvailabilityZonesRequest
}

// Send marshals and sends the DescribeAvailabilityZones API request.
func (r DescribeAvailabilityZonesRequest) Send(ctx context.Context) (*DescribeAvailabilityZonesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeAvailabilityZonesResponse{
		DescribeAvailabilityZonesOutput: r.Request.Data.(*DescribeAvailabilityZonesOutput),
		response:                        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeAvailabilityZonesResponse is the response type for the
// DescribeAvailabilityZones API operation.
type DescribeAvailabilityZonesResponse struct {
	*DescribeAvailabilityZonesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeAvailabilityZones request.
func (r *DescribeAvailabilityZonesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
