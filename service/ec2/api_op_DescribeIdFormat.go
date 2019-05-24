// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DescribeIdFormatRequest
type DescribeIdFormatInput struct {
	_ struct{} `type:"structure"`

	// The type of resource: bundle | conversion-task | customer-gateway | dhcp-options
	// | elastic-ip-allocation | elastic-ip-association | export-task | flow-log
	// | image | import-task | instance | internet-gateway | network-acl | network-acl-association
	// | network-interface | network-interface-attachment | prefix-list | reservation
	// | route-table | route-table-association | security-group | snapshot | subnet
	// | subnet-cidr-block-association | volume | vpc | vpc-cidr-block-association
	// | vpc-endpoint | vpc-peering-connection | vpn-connection | vpn-gateway
	Resource *string `type:"string"`
}

// String returns the string representation
func (s DescribeIdFormatInput) String() string {
	return awsutil.Prettify(s)
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DescribeIdFormatResult
type DescribeIdFormatOutput struct {
	_ struct{} `type:"structure"`

	// Information about the ID format for the resource.
	Statuses []IdFormat `locationName:"statusSet" locationNameList:"item" type:"list"`
}

// String returns the string representation
func (s DescribeIdFormatOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeIdFormat = "DescribeIdFormat"

// DescribeIdFormatRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Describes the ID format settings for your resources on a per-region basis,
// for example, to view which resource types are enabled for longer IDs. This
// request only returns information about resource types whose ID formats can
// be modified; it does not return information about other resource types.
//
// The following resource types support longer IDs: bundle | conversion-task
// | customer-gateway | dhcp-options | elastic-ip-allocation | elastic-ip-association
// | export-task | flow-log | image | import-task | instance | internet-gateway
// | network-acl | network-acl-association | network-interface | network-interface-attachment
// | prefix-list | reservation | route-table | route-table-association | security-group
// | snapshot | subnet | subnet-cidr-block-association | volume | vpc | vpc-cidr-block-association
// | vpc-endpoint | vpc-peering-connection | vpn-connection | vpn-gateway.
//
// These settings apply to the IAM user who makes the request; they do not apply
// to the entire AWS account. By default, an IAM user defaults to the same settings
// as the root user, unless they explicitly override the settings by running
// the ModifyIdFormat command. Resources created with longer IDs are visible
// to all IAM users, regardless of these settings and provided that they have
// permission to use the relevant Describe command for the resource type.
//
//    // Example sending a request using DescribeIdFormatRequest.
//    req := client.DescribeIdFormatRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DescribeIdFormat
func (c *Client) DescribeIdFormatRequest(input *DescribeIdFormatInput) DescribeIdFormatRequest {
	op := &aws.Operation{
		Name:       opDescribeIdFormat,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeIdFormatInput{}
	}

	req := c.newRequest(op, input, &DescribeIdFormatOutput{})
	return DescribeIdFormatRequest{Request: req, Input: input, Copy: c.DescribeIdFormatRequest}
}

// DescribeIdFormatRequest is the request type for the
// DescribeIdFormat API operation.
type DescribeIdFormatRequest struct {
	*aws.Request
	Input *DescribeIdFormatInput
	Copy  func(*DescribeIdFormatInput) DescribeIdFormatRequest
}

// Send marshals and sends the DescribeIdFormat API request.
func (r DescribeIdFormatRequest) Send(ctx context.Context) (*DescribeIdFormatResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeIdFormatResponse{
		DescribeIdFormatOutput: r.Request.Data.(*DescribeIdFormatOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeIdFormatResponse is the response type for the
// DescribeIdFormat API operation.
type DescribeIdFormatResponse struct {
	*DescribeIdFormatOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeIdFormat request.
func (r *DescribeIdFormatResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
