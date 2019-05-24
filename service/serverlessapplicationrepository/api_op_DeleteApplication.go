// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package serverlessapplicationrepository

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/serverlessrepo-2017-09-08/DeleteApplicationRequest
type DeleteApplicationInput struct {
	_ struct{} `type:"structure"`

	// ApplicationId is a required field
	ApplicationId *string `location:"uri" locationName:"applicationId" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteApplicationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteApplicationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteApplicationInput"}

	if s.ApplicationId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ApplicationId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteApplicationInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	if s.ApplicationId != nil {
		v := *s.ApplicationId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "applicationId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/serverlessrepo-2017-09-08/DeleteApplicationOutput
type DeleteApplicationOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteApplicationOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteApplicationOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opDeleteApplication = "DeleteApplication"

// DeleteApplicationRequest returns a request value for making API operation for
// AWSServerlessApplicationRepository.
//
// Deletes the specified application.
//
//    // Example sending a request using DeleteApplicationRequest.
//    req := client.DeleteApplicationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/serverlessrepo-2017-09-08/DeleteApplication
func (c *Client) DeleteApplicationRequest(input *DeleteApplicationInput) DeleteApplicationRequest {
	op := &aws.Operation{
		Name:       opDeleteApplication,
		HTTPMethod: "DELETE",
		HTTPPath:   "/applications/{applicationId}",
	}

	if input == nil {
		input = &DeleteApplicationInput{}
	}

	req := c.newRequest(op, input, &DeleteApplicationOutput{})
	req.Handlers.Unmarshal.Remove(restjson.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return DeleteApplicationRequest{Request: req, Input: input, Copy: c.DeleteApplicationRequest}
}

// DeleteApplicationRequest is the request type for the
// DeleteApplication API operation.
type DeleteApplicationRequest struct {
	*aws.Request
	Input *DeleteApplicationInput
	Copy  func(*DeleteApplicationInput) DeleteApplicationRequest
}

// Send marshals and sends the DeleteApplication API request.
func (r DeleteApplicationRequest) Send(ctx context.Context) (*DeleteApplicationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteApplicationResponse{
		DeleteApplicationOutput: r.Request.Data.(*DeleteApplicationOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteApplicationResponse is the response type for the
// DeleteApplication API operation.
type DeleteApplicationResponse struct {
	*DeleteApplicationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteApplication request.
func (r *DeleteApplicationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
