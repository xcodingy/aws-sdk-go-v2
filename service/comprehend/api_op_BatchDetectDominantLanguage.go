// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package comprehend

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/comprehend-2017-11-27/BatchDetectDominantLanguageRequest
type BatchDetectDominantLanguageInput struct {
	_ struct{} `type:"structure"`

	// A list containing the text of the input documents. The list can contain a
	// maximum of 25 documents. Each document should contain at least 20 characters
	// and must contain fewer than 5,000 bytes of UTF-8 encoded characters.
	//
	// TextList is a required field
	TextList []string `type:"list" required:"true"`
}

// String returns the string representation
func (s BatchDetectDominantLanguageInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *BatchDetectDominantLanguageInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "BatchDetectDominantLanguageInput"}

	if s.TextList == nil {
		invalidParams.Add(aws.NewErrParamRequired("TextList"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/comprehend-2017-11-27/BatchDetectDominantLanguageResponse
type BatchDetectDominantLanguageOutput struct {
	_ struct{} `type:"structure"`

	// A list containing one object for each document that contained an error. The
	// results are sorted in ascending order by the Index field and match the order
	// of the documents in the input list. If there are no errors in the batch,
	// the ErrorList is empty.
	//
	// ErrorList is a required field
	ErrorList []BatchItemError `type:"list" required:"true"`

	// A list of objects containing the results of the operation. The results are
	// sorted in ascending order by the Index field and match the order of the documents
	// in the input list. If all of the documents contain an error, the ResultList
	// is empty.
	//
	// ResultList is a required field
	ResultList []BatchDetectDominantLanguageItemResult `type:"list" required:"true"`
}

// String returns the string representation
func (s BatchDetectDominantLanguageOutput) String() string {
	return awsutil.Prettify(s)
}

const opBatchDetectDominantLanguage = "BatchDetectDominantLanguage"

// BatchDetectDominantLanguageRequest returns a request value for making API operation for
// Amazon Comprehend.
//
// Determines the dominant language of the input text for a batch of documents.
// For a list of languages that Amazon Comprehend can detect, see Amazon Comprehend
// Supported Languages (https://docs.aws.amazon.com/comprehend/latest/dg/how-languages.html).
//
//    // Example sending a request using BatchDetectDominantLanguageRequest.
//    req := client.BatchDetectDominantLanguageRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/comprehend-2017-11-27/BatchDetectDominantLanguage
func (c *Client) BatchDetectDominantLanguageRequest(input *BatchDetectDominantLanguageInput) BatchDetectDominantLanguageRequest {
	op := &aws.Operation{
		Name:       opBatchDetectDominantLanguage,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &BatchDetectDominantLanguageInput{}
	}

	req := c.newRequest(op, input, &BatchDetectDominantLanguageOutput{})
	return BatchDetectDominantLanguageRequest{Request: req, Input: input, Copy: c.BatchDetectDominantLanguageRequest}
}

// BatchDetectDominantLanguageRequest is the request type for the
// BatchDetectDominantLanguage API operation.
type BatchDetectDominantLanguageRequest struct {
	*aws.Request
	Input *BatchDetectDominantLanguageInput
	Copy  func(*BatchDetectDominantLanguageInput) BatchDetectDominantLanguageRequest
}

// Send marshals and sends the BatchDetectDominantLanguage API request.
func (r BatchDetectDominantLanguageRequest) Send(ctx context.Context) (*BatchDetectDominantLanguageResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &BatchDetectDominantLanguageResponse{
		BatchDetectDominantLanguageOutput: r.Request.Data.(*BatchDetectDominantLanguageOutput),
		response:                          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// BatchDetectDominantLanguageResponse is the response type for the
// BatchDetectDominantLanguage API operation.
type BatchDetectDominantLanguageResponse struct {
	*BatchDetectDominantLanguageOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// BatchDetectDominantLanguage request.
func (r *BatchDetectDominantLanguageResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
