// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package transfer

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/transfer-2018-11-05/TagResourceRequest
type TagResourceInput struct {
	_ struct{} `type:"structure"`

	// An Amazon Resource Name (ARN) for a specific AWS resource, such as a server,
	// user, or role.
	//
	// Arn is a required field
	Arn *string `min:"20" type:"string" required:"true"`

	// Key-value pairs assigned to ARNs that you can use to group and search for
	// resources by type. You can attach this metadata to user accounts for any
	// purpose.
	//
	// Tags is a required field
	Tags []Tag `min:"1" type:"list" required:"true"`
}

// String returns the string representation
func (s TagResourceInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *TagResourceInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "TagResourceInput"}

	if s.Arn == nil {
		invalidParams.Add(aws.NewErrParamRequired("Arn"))
	}
	if s.Arn != nil && len(*s.Arn) < 20 {
		invalidParams.Add(aws.NewErrParamMinLen("Arn", 20))
	}

	if s.Tags == nil {
		invalidParams.Add(aws.NewErrParamRequired("Tags"))
	}
	if s.Tags != nil && len(s.Tags) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Tags", 1))
	}
	if s.Tags != nil {
		for i, v := range s.Tags {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Tags", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/transfer-2018-11-05/TagResourceOutput
type TagResourceOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s TagResourceOutput) String() string {
	return awsutil.Prettify(s)
}

const opTagResource = "TagResource"

// TagResourceRequest returns a request value for making API operation for
// AWS Transfer for SFTP.
//
// Attaches a key-value pair to a resource, as identified by its Amazon Resource
// Name (ARN). Resources are users, servers, roles, and other entities.
//
// There is no response returned from this call.
//
//    // Example sending a request using TagResourceRequest.
//    req := client.TagResourceRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/transfer-2018-11-05/TagResource
func (c *Client) TagResourceRequest(input *TagResourceInput) TagResourceRequest {
	op := &aws.Operation{
		Name:       opTagResource,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &TagResourceInput{}
	}

	req := c.newRequest(op, input, &TagResourceOutput{})
	req.Handlers.Unmarshal.Remove(jsonrpc.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return TagResourceRequest{Request: req, Input: input, Copy: c.TagResourceRequest}
}

// TagResourceRequest is the request type for the
// TagResource API operation.
type TagResourceRequest struct {
	*aws.Request
	Input *TagResourceInput
	Copy  func(*TagResourceInput) TagResourceRequest
}

// Send marshals and sends the TagResource API request.
func (r TagResourceRequest) Send(ctx context.Context) (*TagResourceResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &TagResourceResponse{
		TagResourceOutput: r.Request.Data.(*TagResourceOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// TagResourceResponse is the response type for the
// TagResource API operation.
type TagResourceResponse struct {
	*TagResourceOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// TagResource request.
func (r *TagResourceResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
