// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package transcribe

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/transcribe-2017-10-26/GetVocabularyRequest
type GetVocabularyInput struct {
	_ struct{} `type:"structure"`

	// The name of the vocabulary to return information about. The name is case-sensitive.
	//
	// VocabularyName is a required field
	VocabularyName *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s GetVocabularyInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetVocabularyInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetVocabularyInput"}

	if s.VocabularyName == nil {
		invalidParams.Add(aws.NewErrParamRequired("VocabularyName"))
	}
	if s.VocabularyName != nil && len(*s.VocabularyName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("VocabularyName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/transcribe-2017-10-26/GetVocabularyResponse
type GetVocabularyOutput struct {
	_ struct{} `type:"structure"`

	// The S3 location where the vocabulary is stored. Use this URI to get the contents
	// of the vocabulary. The URI is available for a limited time.
	DownloadUri *string `min:"1" type:"string"`

	// If the VocabularyState field is FAILED, this field contains information about
	// why the job failed.
	FailureReason *string `type:"string"`

	// The language code of the vocabulary entries.
	LanguageCode LanguageCode `type:"string" enum:"true"`

	// The date and time that the vocabulary was last modified.
	LastModifiedTime *time.Time `type:"timestamp" timestampFormat:"unix"`

	// The name of the vocabulary to return.
	VocabularyName *string `min:"1" type:"string"`

	// The processing state of the vocabulary.
	VocabularyState VocabularyState `type:"string" enum:"true"`
}

// String returns the string representation
func (s GetVocabularyOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetVocabulary = "GetVocabulary"

// GetVocabularyRequest returns a request value for making API operation for
// Amazon Transcribe Service.
//
// Gets information about a vocabulary.
//
//    // Example sending a request using GetVocabularyRequest.
//    req := client.GetVocabularyRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/transcribe-2017-10-26/GetVocabulary
func (c *Client) GetVocabularyRequest(input *GetVocabularyInput) GetVocabularyRequest {
	op := &aws.Operation{
		Name:       opGetVocabulary,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetVocabularyInput{}
	}

	req := c.newRequest(op, input, &GetVocabularyOutput{})
	return GetVocabularyRequest{Request: req, Input: input, Copy: c.GetVocabularyRequest}
}

// GetVocabularyRequest is the request type for the
// GetVocabulary API operation.
type GetVocabularyRequest struct {
	*aws.Request
	Input *GetVocabularyInput
	Copy  func(*GetVocabularyInput) GetVocabularyRequest
}

// Send marshals and sends the GetVocabulary API request.
func (r GetVocabularyRequest) Send(ctx context.Context) (*GetVocabularyResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetVocabularyResponse{
		GetVocabularyOutput: r.Request.Data.(*GetVocabularyOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetVocabularyResponse is the response type for the
// GetVocabulary API operation.
type GetVocabularyResponse struct {
	*GetVocabularyOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetVocabulary request.
func (r *GetVocabularyResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
