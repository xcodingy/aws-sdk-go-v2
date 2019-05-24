// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package workdocs

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/workdocs-2016-05-01/DeleteDocumentRequest
type DeleteDocumentInput struct {
	_ struct{} `type:"structure"`

	// Amazon WorkDocs authentication token. Do not set this field when using administrative
	// API actions, as in accessing the API using AWS credentials.
	AuthenticationToken *string `location:"header" locationName:"Authentication" min:"1" type:"string"`

	// The ID of the document.
	//
	// DocumentId is a required field
	DocumentId *string `location:"uri" locationName:"DocumentId" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteDocumentInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteDocumentInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteDocumentInput"}
	if s.AuthenticationToken != nil && len(*s.AuthenticationToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("AuthenticationToken", 1))
	}

	if s.DocumentId == nil {
		invalidParams.Add(aws.NewErrParamRequired("DocumentId"))
	}
	if s.DocumentId != nil && len(*s.DocumentId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("DocumentId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteDocumentInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	if s.AuthenticationToken != nil {
		v := *s.AuthenticationToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "Authentication", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.DocumentId != nil {
		v := *s.DocumentId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "DocumentId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/workdocs-2016-05-01/DeleteDocumentOutput
type DeleteDocumentOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteDocumentOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteDocumentOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opDeleteDocument = "DeleteDocument"

// DeleteDocumentRequest returns a request value for making API operation for
// Amazon WorkDocs.
//
// Permanently deletes the specified document and its associated metadata.
//
//    // Example sending a request using DeleteDocumentRequest.
//    req := client.DeleteDocumentRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/workdocs-2016-05-01/DeleteDocument
func (c *Client) DeleteDocumentRequest(input *DeleteDocumentInput) DeleteDocumentRequest {
	op := &aws.Operation{
		Name:       opDeleteDocument,
		HTTPMethod: "DELETE",
		HTTPPath:   "/api/v1/documents/{DocumentId}",
	}

	if input == nil {
		input = &DeleteDocumentInput{}
	}

	req := c.newRequest(op, input, &DeleteDocumentOutput{})
	req.Handlers.Unmarshal.Remove(restjson.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return DeleteDocumentRequest{Request: req, Input: input, Copy: c.DeleteDocumentRequest}
}

// DeleteDocumentRequest is the request type for the
// DeleteDocument API operation.
type DeleteDocumentRequest struct {
	*aws.Request
	Input *DeleteDocumentInput
	Copy  func(*DeleteDocumentInput) DeleteDocumentRequest
}

// Send marshals and sends the DeleteDocument API request.
func (r DeleteDocumentRequest) Send(ctx context.Context) (*DeleteDocumentResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteDocumentResponse{
		DeleteDocumentOutput: r.Request.Data.(*DeleteDocumentOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteDocumentResponse is the response type for the
// DeleteDocument API operation.
type DeleteDocumentResponse struct {
	*DeleteDocumentOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteDocument request.
func (r *DeleteDocumentResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
