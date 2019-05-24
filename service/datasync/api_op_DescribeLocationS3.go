// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package datasync

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// DescribeLocationS3Request
// Please also see https://docs.aws.amazon.com/goto/WebAPI/datasync-2018-11-09/DescribeLocationS3Request
type DescribeLocationS3Input struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the Amazon S3 bucket location to describe.
	//
	// LocationArn is a required field
	LocationArn *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeLocationS3Input) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeLocationS3Input) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeLocationS3Input"}

	if s.LocationArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("LocationArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// DescribeLocationS3Response
// Please also see https://docs.aws.amazon.com/goto/WebAPI/datasync-2018-11-09/DescribeLocationS3Response
type DescribeLocationS3Output struct {
	_ struct{} `type:"structure"`

	// The time that the Amazon S3 bucket location was created.
	CreationTime *time.Time `type:"timestamp" timestampFormat:"unix"`

	// The Amazon Resource Name (ARN) of the Amazon S3 bucket location.
	LocationArn *string `type:"string"`

	// The URL of the Amazon S3 location that was described.
	LocationUri *string `type:"string"`

	// The Amazon Resource Name (ARN) of the AWS Identity and Access Management
	// (IAM) role that is used to access an Amazon S3 bucket. For detailed information
	// about using such a role, see Components and Terminology (https://alpha-aws-docs.aws.amazon.com/sync-service/latest/userguide/create-locations-cli.html#create-location-s3-cli)
	// in the AWS DataSync User Guide.
	S3Config *S3Config `type:"structure"`
}

// String returns the string representation
func (s DescribeLocationS3Output) String() string {
	return awsutil.Prettify(s)
}

const opDescribeLocationS3 = "DescribeLocationS3"

// DescribeLocationS3Request returns a request value for making API operation for
// AWS DataSync.
//
// Returns metadata, such as bucket name, about an Amazon S3 bucket location.
//
//    // Example sending a request using DescribeLocationS3Request.
//    req := client.DescribeLocationS3Request(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/datasync-2018-11-09/DescribeLocationS3
func (c *Client) DescribeLocationS3Request(input *DescribeLocationS3Input) DescribeLocationS3Request {
	op := &aws.Operation{
		Name:       opDescribeLocationS3,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeLocationS3Input{}
	}

	req := c.newRequest(op, input, &DescribeLocationS3Output{})
	return DescribeLocationS3Request{Request: req, Input: input, Copy: c.DescribeLocationS3Request}
}

// DescribeLocationS3Request is the request type for the
// DescribeLocationS3 API operation.
type DescribeLocationS3Request struct {
	*aws.Request
	Input *DescribeLocationS3Input
	Copy  func(*DescribeLocationS3Input) DescribeLocationS3Request
}

// Send marshals and sends the DescribeLocationS3 API request.
func (r DescribeLocationS3Request) Send(ctx context.Context) (*DescribeLocationS3Response, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeLocationS3Response{
		DescribeLocationS3Output: r.Request.Data.(*DescribeLocationS3Output),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeLocationS3Response is the response type for the
// DescribeLocationS3 API operation.
type DescribeLocationS3Response struct {
	*DescribeLocationS3Output

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeLocationS3 request.
func (r *DescribeLocationS3Response) SDKResponseMetdata() *aws.Response {
	return r.response
}
