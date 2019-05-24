// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package amplify

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Result structure for get branch request.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/amplify-2017-07-25/GetBranchRequest
type GetBranchInput struct {
	_ struct{} `type:"structure"`

	// Unique Id for an Amplify App.
	//
	// AppId is a required field
	AppId *string `location:"uri" locationName:"appId" min:"1" type:"string" required:"true"`

	// Name for the branch.
	//
	// BranchName is a required field
	BranchName *string `location:"uri" locationName:"branchName" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s GetBranchInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetBranchInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetBranchInput"}

	if s.AppId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AppId"))
	}
	if s.AppId != nil && len(*s.AppId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("AppId", 1))
	}

	if s.BranchName == nil {
		invalidParams.Add(aws.NewErrParamRequired("BranchName"))
	}
	if s.BranchName != nil && len(*s.BranchName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("BranchName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetBranchInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	if s.AppId != nil {
		v := *s.AppId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "appId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.BranchName != nil {
		v := *s.BranchName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "branchName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/amplify-2017-07-25/GetBranchResult
type GetBranchOutput struct {
	_ struct{} `type:"structure"`

	// Branch for an Amplify App, which maps to a 3rd party repository branch.
	//
	// Branch is a required field
	Branch *Branch `locationName:"branch" type:"structure" required:"true"`
}

// String returns the string representation
func (s GetBranchOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetBranchOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Branch != nil {
		v := s.Branch

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "branch", v, metadata)
	}
	return nil
}

const opGetBranch = "GetBranch"

// GetBranchRequest returns a request value for making API operation for
// AWS Amplify.
//
// Retrieves a branch for an Amplify App.
//
//    // Example sending a request using GetBranchRequest.
//    req := client.GetBranchRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/amplify-2017-07-25/GetBranch
func (c *Client) GetBranchRequest(input *GetBranchInput) GetBranchRequest {
	op := &aws.Operation{
		Name:       opGetBranch,
		HTTPMethod: "GET",
		HTTPPath:   "/apps/{appId}/branches/{branchName}",
	}

	if input == nil {
		input = &GetBranchInput{}
	}

	req := c.newRequest(op, input, &GetBranchOutput{})
	return GetBranchRequest{Request: req, Input: input, Copy: c.GetBranchRequest}
}

// GetBranchRequest is the request type for the
// GetBranch API operation.
type GetBranchRequest struct {
	*aws.Request
	Input *GetBranchInput
	Copy  func(*GetBranchInput) GetBranchRequest
}

// Send marshals and sends the GetBranch API request.
func (r GetBranchRequest) Send(ctx context.Context) (*GetBranchResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetBranchResponse{
		GetBranchOutput: r.Request.Data.(*GetBranchOutput),
		response:        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetBranchResponse is the response type for the
// GetBranch API operation.
type GetBranchResponse struct {
	*GetBranchOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetBranch request.
func (r *GetBranchResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
