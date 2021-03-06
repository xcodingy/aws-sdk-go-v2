// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package kendra

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeIndexInput struct {
	_ struct{} `type:"structure"`

	// The name of the index to describe.
	//
	// Id is a required field
	Id *string `min:"36" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeIndexInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeIndexInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeIndexInput"}

	if s.Id == nil {
		invalidParams.Add(aws.NewErrParamRequired("Id"))
	}
	if s.Id != nil && len(*s.Id) < 36 {
		invalidParams.Add(aws.NewErrParamMinLen("Id", 36))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeIndexOutput struct {
	_ struct{} `type:"structure"`

	// For enterprise edtion indexes, you can choose to use additional capacity
	// to meet the needs of your application. This contains the capacity units used
	// for the index. A 0 for the query capacity or the storage capacity indicates
	// that the index is using the default capacity for the index.
	CapacityUnits *CapacityUnitsConfiguration `type:"structure"`

	// The Unix datetime that the index was created.
	CreatedAt *time.Time `type:"timestamp"`

	// The description of the index.
	Description *string `min:"1" type:"string"`

	// Configuration settings for any metadata applied to the documents in the index.
	DocumentMetadataConfigurations []DocumentMetadataConfiguration `type:"list"`

	// The Amazon Kendra edition used for the index. You decide the edition when
	// you create the index.
	Edition IndexEdition `type:"string" enum:"true"`

	// When th eStatus field value is FAILED, the ErrorMessage field contains a
	// message that explains why.
	ErrorMessage *string `min:"1" type:"string"`

	// the name of the index.
	Id *string `min:"36" type:"string"`

	// Provides information about the number of FAQ questions and answers and the
	// number of text documents indexed.
	IndexStatistics *IndexStatistics `type:"structure"`

	// The name of the index.
	Name *string `min:"1" type:"string"`

	// The Amazon Resource Name (ARN) of the IAM role that gives Amazon Kendra permission
	// to write to your Amazon Cloudwatch logs.
	RoleArn *string `min:"1" type:"string"`

	// The identifier of the AWS KMS customer master key (CMK) used to encrypt your
	// data. Amazon Kendra doesn't support asymmetric CMKs.
	ServerSideEncryptionConfiguration *ServerSideEncryptionConfiguration `type:"structure"`

	// The current status of the index. When the value is ACTIVE, the index is ready
	// for use. If the Status field value is FAILED, the ErrorMessage field contains
	// a message that explains why.
	Status IndexStatus `type:"string" enum:"true"`

	// The Unix datetime that the index was last updated.
	UpdatedAt *time.Time `type:"timestamp"`
}

// String returns the string representation
func (s DescribeIndexOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeIndex = "DescribeIndex"

// DescribeIndexRequest returns a request value for making API operation for
// AWSKendraFrontendService.
//
// Describes an existing Amazon Kendra index
//
//    // Example sending a request using DescribeIndexRequest.
//    req := client.DescribeIndexRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/kendra-2019-02-03/DescribeIndex
func (c *Client) DescribeIndexRequest(input *DescribeIndexInput) DescribeIndexRequest {
	op := &aws.Operation{
		Name:       opDescribeIndex,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeIndexInput{}
	}

	req := c.newRequest(op, input, &DescribeIndexOutput{})

	return DescribeIndexRequest{Request: req, Input: input, Copy: c.DescribeIndexRequest}
}

// DescribeIndexRequest is the request type for the
// DescribeIndex API operation.
type DescribeIndexRequest struct {
	*aws.Request
	Input *DescribeIndexInput
	Copy  func(*DescribeIndexInput) DescribeIndexRequest
}

// Send marshals and sends the DescribeIndex API request.
func (r DescribeIndexRequest) Send(ctx context.Context) (*DescribeIndexResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeIndexResponse{
		DescribeIndexOutput: r.Request.Data.(*DescribeIndexOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeIndexResponse is the response type for the
// DescribeIndex API operation.
type DescribeIndexResponse struct {
	*DescribeIndexOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeIndex request.
func (r *DescribeIndexResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
