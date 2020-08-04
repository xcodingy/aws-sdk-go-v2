// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package translate

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type StartTextTranslationJobInput struct {
	_ struct{} `type:"structure"`

	// The client token of the EC2 instance calling the request. This token is auto-generated
	// when using the Amazon Translate SDK. Otherwise, use the DescribeInstances
	// (docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeInstances.html)
	// EC2 operation to retreive an instance's client token. For more information,
	// see Client Tokens (docs.aws.amazon.com/AWSEC2/latest/APIReference/Run_Instance_Idempotency.html#client-tokens)
	// in the EC2 User Guide.
	//
	// ClientToken is a required field
	ClientToken *string `min:"1" type:"string" required:"true" idempotencyToken:"true"`

	// The Amazon Resource Name (ARN) of an AWS Identity Access and Management (IAM)
	// role that grants Amazon Translate read access to your input data. For more
	// nformation, see identity-and-access-management.
	//
	// DataAccessRoleArn is a required field
	DataAccessRoleArn *string `min:"20" type:"string" required:"true"`

	// Specifies the format and S3 location of the input documents for the translation
	// job.
	//
	// InputDataConfig is a required field
	InputDataConfig *InputDataConfig `type:"structure" required:"true"`

	// The name of the batch translation job to be performed.
	JobName *string `min:"1" type:"string"`

	// Specifies the S3 folder to which your job output will be saved.
	//
	// OutputDataConfig is a required field
	OutputDataConfig *OutputDataConfig `type:"structure" required:"true"`

	// The language code of the input language. For a list of language codes, see
	// what-is-languages.
	//
	// Amazon Translate does not automatically detect a source language during batch
	// translation jobs.
	//
	// SourceLanguageCode is a required field
	SourceLanguageCode *string `min:"2" type:"string" required:"true"`

	// The language code of the output language.
	//
	// TargetLanguageCodes is a required field
	TargetLanguageCodes []string `min:"1" type:"list" required:"true"`

	// The name of the terminology to use in the batch translation job. For a list
	// of available terminologies, use the ListTerminologies operation.
	TerminologyNames []string `type:"list"`
}

// String returns the string representation
func (s StartTextTranslationJobInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *StartTextTranslationJobInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "StartTextTranslationJobInput"}

	if s.ClientToken == nil {
		invalidParams.Add(aws.NewErrParamRequired("ClientToken"))
	}
	if s.ClientToken != nil && len(*s.ClientToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ClientToken", 1))
	}

	if s.DataAccessRoleArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("DataAccessRoleArn"))
	}
	if s.DataAccessRoleArn != nil && len(*s.DataAccessRoleArn) < 20 {
		invalidParams.Add(aws.NewErrParamMinLen("DataAccessRoleArn", 20))
	}

	if s.InputDataConfig == nil {
		invalidParams.Add(aws.NewErrParamRequired("InputDataConfig"))
	}
	if s.JobName != nil && len(*s.JobName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("JobName", 1))
	}

	if s.OutputDataConfig == nil {
		invalidParams.Add(aws.NewErrParamRequired("OutputDataConfig"))
	}

	if s.SourceLanguageCode == nil {
		invalidParams.Add(aws.NewErrParamRequired("SourceLanguageCode"))
	}
	if s.SourceLanguageCode != nil && len(*s.SourceLanguageCode) < 2 {
		invalidParams.Add(aws.NewErrParamMinLen("SourceLanguageCode", 2))
	}

	if s.TargetLanguageCodes == nil {
		invalidParams.Add(aws.NewErrParamRequired("TargetLanguageCodes"))
	}
	if s.TargetLanguageCodes != nil && len(s.TargetLanguageCodes) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("TargetLanguageCodes", 1))
	}
	if s.InputDataConfig != nil {
		if err := s.InputDataConfig.Validate(); err != nil {
			invalidParams.AddNested("InputDataConfig", err.(aws.ErrInvalidParams))
		}
	}
	if s.OutputDataConfig != nil {
		if err := s.OutputDataConfig.Validate(); err != nil {
			invalidParams.AddNested("OutputDataConfig", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type StartTextTranslationJobOutput struct {
	_ struct{} `type:"structure"`

	// The identifier generated for the job. To get the status of a job, use this
	// ID with the DescribeTextTranslationJob operation.
	JobId *string `min:"1" type:"string"`

	// The status of the job. Possible values include:
	//
	//    * SUBMITTED - The job has been received and is queued for processing.
	//
	//    * IN_PROGRESS - Amazon Translate is processing the job.
	//
	//    * COMPLETED - The job was successfully completed and the output is available.
	//
	//    * COMPLETED_WITH_ERRORS - The job was completed with errors. The errors
	//    can be analyzed in the job's output.
	//
	//    * FAILED - The job did not complete. To get details, use the DescribeTextTranslationJob
	//    operation.
	//
	//    * STOP_REQUESTED - The user who started the job has requested that it
	//    be stopped.
	//
	//    * STOPPED - The job has been stopped.
	JobStatus JobStatus `type:"string" enum:"true"`
}

// String returns the string representation
func (s StartTextTranslationJobOutput) String() string {
	return awsutil.Prettify(s)
}

const opStartTextTranslationJob = "StartTextTranslationJob"

// StartTextTranslationJobRequest returns a request value for making API operation for
// Amazon Translate.
//
// Starts an asynchronous batch translation job. Batch translation jobs can
// be used to translate large volumes of text across multiple documents at once.
// For more information, see async.
//
// Batch translation jobs can be described with the DescribeTextTranslationJob
// operation, listed with the ListTextTranslationJobs operation, and stopped
// with the StopTextTranslationJob operation.
//
// Amazon Translate does not support batch translation of multiple source languages
// at once.
//
//    // Example sending a request using StartTextTranslationJobRequest.
//    req := client.StartTextTranslationJobRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/translate-2017-07-01/StartTextTranslationJob
func (c *Client) StartTextTranslationJobRequest(input *StartTextTranslationJobInput) StartTextTranslationJobRequest {
	op := &aws.Operation{
		Name:       opStartTextTranslationJob,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &StartTextTranslationJobInput{}
	}

	req := c.newRequest(op, input, &StartTextTranslationJobOutput{})

	return StartTextTranslationJobRequest{Request: req, Input: input, Copy: c.StartTextTranslationJobRequest}
}

// StartTextTranslationJobRequest is the request type for the
// StartTextTranslationJob API operation.
type StartTextTranslationJobRequest struct {
	*aws.Request
	Input *StartTextTranslationJobInput
	Copy  func(*StartTextTranslationJobInput) StartTextTranslationJobRequest
}

// Send marshals and sends the StartTextTranslationJob API request.
func (r StartTextTranslationJobRequest) Send(ctx context.Context) (*StartTextTranslationJobResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &StartTextTranslationJobResponse{
		StartTextTranslationJobOutput: r.Request.Data.(*StartTextTranslationJobOutput),
		response:                      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// StartTextTranslationJobResponse is the response type for the
// StartTextTranslationJob API operation.
type StartTextTranslationJobResponse struct {
	*StartTextTranslationJobOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// StartTextTranslationJob request.
func (r *StartTextTranslationJobResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}