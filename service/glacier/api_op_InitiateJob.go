// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package glacier

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Provides options for initiating an Amazon Glacier job.
type InitiateJobInput struct {
	_ struct{} `type:"structure" payload:"JobParameters"`

	// The AccountId value is the AWS account ID of the account that owns the vault.
	// You can either specify an AWS account ID or optionally a single '-' (hyphen),
	// in which case Amazon Glacier uses the AWS account ID associated with the
	// credentials used to sign the request. If you use an account ID, do not include
	// any hyphens ('-') in the ID.
	//
	// AccountId is a required field
	AccountId *string `location:"uri" locationName:"accountId" type:"string" required:"true"`

	// Provides options for specifying job information.
	JobParameters *JobParameters `locationName:"jobParameters" type:"structure"`

	// The name of the vault.
	//
	// VaultName is a required field
	VaultName *string `location:"uri" locationName:"vaultName" type:"string" required:"true"`
}

// String returns the string representation
func (s InitiateJobInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *InitiateJobInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "InitiateJobInput"}

	if s.AccountId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AccountId"))
	}

	if s.VaultName == nil {
		invalidParams.Add(aws.NewErrParamRequired("VaultName"))
	}
	if s.JobParameters != nil {
		if err := s.JobParameters.Validate(); err != nil {
			invalidParams.AddNested("JobParameters", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s InitiateJobInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.AccountId != nil {
		v := *s.AccountId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "accountId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.VaultName != nil {
		v := *s.VaultName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "vaultName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.JobParameters != nil {
		v := s.JobParameters

		metadata := protocol.Metadata{}
		e.SetFields(protocol.PayloadTarget, "jobParameters", v, metadata)
	}
	return nil
}

// Contains the Amazon Glacier response to your request.
type InitiateJobOutput struct {
	_ struct{} `type:"structure"`

	// The ID of the job.
	JobId *string `location:"header" locationName:"x-amz-job-id" type:"string"`

	// The path to the location of where the select results are stored.
	JobOutputPath *string `location:"header" locationName:"x-amz-job-output-path" type:"string"`

	// The relative URI path of the job.
	Location *string `location:"header" locationName:"Location" type:"string"`
}

// String returns the string representation
func (s InitiateJobOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s InitiateJobOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.JobId != nil {
		v := *s.JobId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-job-id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.JobOutputPath != nil {
		v := *s.JobOutputPath

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-job-output-path", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Location != nil {
		v := *s.Location

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "Location", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opInitiateJob = "InitiateJob"

// InitiateJobRequest returns a request value for making API operation for
// Amazon Glacier.
//
// This operation initiates a job of the specified type, which can be a select,
// an archival retrieval, or a vault retrieval. For more information about using
// this operation, see the documentation for the underlying REST API Initiate
// a Job (http://docs.aws.amazon.com/amazonglacier/latest/dev/api-initiate-job-post.html).
//
//    // Example sending a request using InitiateJobRequest.
//    req := client.InitiateJobRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) InitiateJobRequest(input *InitiateJobInput) InitiateJobRequest {
	op := &aws.Operation{
		Name:       opInitiateJob,
		HTTPMethod: "POST",
		HTTPPath:   "/{accountId}/vaults/{vaultName}/jobs",
	}

	if input == nil {
		input = &InitiateJobInput{}
	}

	req := c.newRequest(op, input, &InitiateJobOutput{})
	return InitiateJobRequest{Request: req, Input: input, Copy: c.InitiateJobRequest}
}

// InitiateJobRequest is the request type for the
// InitiateJob API operation.
type InitiateJobRequest struct {
	*aws.Request
	Input *InitiateJobInput
	Copy  func(*InitiateJobInput) InitiateJobRequest
}

// Send marshals and sends the InitiateJob API request.
func (r InitiateJobRequest) Send(ctx context.Context) (*InitiateJobResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &InitiateJobResponse{
		InitiateJobOutput: r.Request.Data.(*InitiateJobOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// InitiateJobResponse is the response type for the
// InitiateJob API operation.
type InitiateJobResponse struct {
	*InitiateJobOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// InitiateJob request.
func (r *InitiateJobResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
