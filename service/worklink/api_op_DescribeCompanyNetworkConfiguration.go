// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package worklink

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/worklink-2018-09-25/DescribeCompanyNetworkConfigurationRequest
type DescribeCompanyNetworkConfigurationInput struct {
	_ struct{} `type:"structure"`

	// The ARN of the fleet.
	//
	// FleetArn is a required field
	FleetArn *string `min:"20" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeCompanyNetworkConfigurationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeCompanyNetworkConfigurationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeCompanyNetworkConfigurationInput"}

	if s.FleetArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("FleetArn"))
	}
	if s.FleetArn != nil && len(*s.FleetArn) < 20 {
		invalidParams.Add(aws.NewErrParamMinLen("FleetArn", 20))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeCompanyNetworkConfigurationInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	if s.FleetArn != nil {
		v := *s.FleetArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "FleetArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/worklink-2018-09-25/DescribeCompanyNetworkConfigurationResponse
type DescribeCompanyNetworkConfigurationOutput struct {
	_ struct{} `type:"structure"`

	// The security groups associated with access to the provided subnets.
	SecurityGroupIds []string `type:"list"`

	// The subnets used for X-ENI connections from Amazon WorkLink rendering containers.
	SubnetIds []string `type:"list"`

	// The VPC with connectivity to associated websites.
	VpcId *string `type:"string"`
}

// String returns the string representation
func (s DescribeCompanyNetworkConfigurationOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeCompanyNetworkConfigurationOutput) MarshalFields(e protocol.FieldEncoder) error {
	if len(s.SecurityGroupIds) > 0 {
		v := s.SecurityGroupIds

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "SecurityGroupIds", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	if len(s.SubnetIds) > 0 {
		v := s.SubnetIds

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "SubnetIds", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	if s.VpcId != nil {
		v := *s.VpcId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "VpcId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opDescribeCompanyNetworkConfiguration = "DescribeCompanyNetworkConfiguration"

// DescribeCompanyNetworkConfigurationRequest returns a request value for making API operation for
// Amazon WorkLink.
//
// Describes the networking configuration to access the internal websites associated
// with the specified fleet.
//
//    // Example sending a request using DescribeCompanyNetworkConfigurationRequest.
//    req := client.DescribeCompanyNetworkConfigurationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/worklink-2018-09-25/DescribeCompanyNetworkConfiguration
func (c *Client) DescribeCompanyNetworkConfigurationRequest(input *DescribeCompanyNetworkConfigurationInput) DescribeCompanyNetworkConfigurationRequest {
	op := &aws.Operation{
		Name:       opDescribeCompanyNetworkConfiguration,
		HTTPMethod: "POST",
		HTTPPath:   "/describeCompanyNetworkConfiguration",
	}

	if input == nil {
		input = &DescribeCompanyNetworkConfigurationInput{}
	}

	req := c.newRequest(op, input, &DescribeCompanyNetworkConfigurationOutput{})
	return DescribeCompanyNetworkConfigurationRequest{Request: req, Input: input, Copy: c.DescribeCompanyNetworkConfigurationRequest}
}

// DescribeCompanyNetworkConfigurationRequest is the request type for the
// DescribeCompanyNetworkConfiguration API operation.
type DescribeCompanyNetworkConfigurationRequest struct {
	*aws.Request
	Input *DescribeCompanyNetworkConfigurationInput
	Copy  func(*DescribeCompanyNetworkConfigurationInput) DescribeCompanyNetworkConfigurationRequest
}

// Send marshals and sends the DescribeCompanyNetworkConfiguration API request.
func (r DescribeCompanyNetworkConfigurationRequest) Send(ctx context.Context) (*DescribeCompanyNetworkConfigurationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeCompanyNetworkConfigurationResponse{
		DescribeCompanyNetworkConfigurationOutput: r.Request.Data.(*DescribeCompanyNetworkConfigurationOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeCompanyNetworkConfigurationResponse is the response type for the
// DescribeCompanyNetworkConfiguration API operation.
type DescribeCompanyNetworkConfigurationResponse struct {
	*DescribeCompanyNetworkConfigurationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeCompanyNetworkConfiguration request.
func (r *DescribeCompanyNetworkConfigurationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
