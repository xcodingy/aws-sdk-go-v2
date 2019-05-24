// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package batch

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/batch-2016-08-10/DeleteJobQueueRequest
type DeleteJobQueueInput struct {
	_ struct{} `type:"structure"`

	// The short name or full Amazon Resource Name (ARN) of the queue to delete.
	//
	// JobQueue is a required field
	JobQueue *string `locationName:"jobQueue" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteJobQueueInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteJobQueueInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteJobQueueInput"}

	if s.JobQueue == nil {
		invalidParams.Add(aws.NewErrParamRequired("JobQueue"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteJobQueueInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	if s.JobQueue != nil {
		v := *s.JobQueue

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "jobQueue", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/batch-2016-08-10/DeleteJobQueueResponse
type DeleteJobQueueOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteJobQueueOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteJobQueueOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opDeleteJobQueue = "DeleteJobQueue"

// DeleteJobQueueRequest returns a request value for making API operation for
// AWS Batch.
//
// Deletes the specified job queue. You must first disable submissions for a
// queue with the UpdateJobQueue operation. All jobs in the queue are terminated
// when you delete a job queue.
//
// It is not necessary to disassociate compute environments from a queue before
// submitting a DeleteJobQueue request.
//
//    // Example sending a request using DeleteJobQueueRequest.
//    req := client.DeleteJobQueueRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/batch-2016-08-10/DeleteJobQueue
func (c *Client) DeleteJobQueueRequest(input *DeleteJobQueueInput) DeleteJobQueueRequest {
	op := &aws.Operation{
		Name:       opDeleteJobQueue,
		HTTPMethod: "POST",
		HTTPPath:   "/v1/deletejobqueue",
	}

	if input == nil {
		input = &DeleteJobQueueInput{}
	}

	req := c.newRequest(op, input, &DeleteJobQueueOutput{})
	return DeleteJobQueueRequest{Request: req, Input: input, Copy: c.DeleteJobQueueRequest}
}

// DeleteJobQueueRequest is the request type for the
// DeleteJobQueue API operation.
type DeleteJobQueueRequest struct {
	*aws.Request
	Input *DeleteJobQueueInput
	Copy  func(*DeleteJobQueueInput) DeleteJobQueueRequest
}

// Send marshals and sends the DeleteJobQueue API request.
func (r DeleteJobQueueRequest) Send(ctx context.Context) (*DeleteJobQueueResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteJobQueueResponse{
		DeleteJobQueueOutput: r.Request.Data.(*DeleteJobQueueOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteJobQueueResponse is the response type for the
// DeleteJobQueue API operation.
type DeleteJobQueueResponse struct {
	*DeleteJobQueueOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteJobQueue request.
func (r *DeleteJobQueueResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
