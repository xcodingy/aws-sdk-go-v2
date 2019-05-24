// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package lightsail

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/lightsail-2016-11-28/DeleteRelationalDatabaseSnapshotRequest
type DeleteRelationalDatabaseSnapshotInput struct {
	_ struct{} `type:"structure"`

	// The name of the database snapshot that you are deleting.
	//
	// RelationalDatabaseSnapshotName is a required field
	RelationalDatabaseSnapshotName *string `locationName:"relationalDatabaseSnapshotName" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteRelationalDatabaseSnapshotInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteRelationalDatabaseSnapshotInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteRelationalDatabaseSnapshotInput"}

	if s.RelationalDatabaseSnapshotName == nil {
		invalidParams.Add(aws.NewErrParamRequired("RelationalDatabaseSnapshotName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/lightsail-2016-11-28/DeleteRelationalDatabaseSnapshotResult
type DeleteRelationalDatabaseSnapshotOutput struct {
	_ struct{} `type:"structure"`

	// An object describing the result of your delete relational database snapshot
	// request.
	Operations []Operation `locationName:"operations" type:"list"`
}

// String returns the string representation
func (s DeleteRelationalDatabaseSnapshotOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteRelationalDatabaseSnapshot = "DeleteRelationalDatabaseSnapshot"

// DeleteRelationalDatabaseSnapshotRequest returns a request value for making API operation for
// Amazon Lightsail.
//
// Deletes a database snapshot in Amazon Lightsail.
//
// The delete relational database snapshot operation supports tag-based access
// control via resource tags applied to the resource identified by relationalDatabaseName.
// For more information, see the Lightsail Dev Guide (https://lightsail.aws.amazon.com/ls/docs/en/articles/amazon-lightsail-controlling-access-using-tags).
//
//    // Example sending a request using DeleteRelationalDatabaseSnapshotRequest.
//    req := client.DeleteRelationalDatabaseSnapshotRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/lightsail-2016-11-28/DeleteRelationalDatabaseSnapshot
func (c *Client) DeleteRelationalDatabaseSnapshotRequest(input *DeleteRelationalDatabaseSnapshotInput) DeleteRelationalDatabaseSnapshotRequest {
	op := &aws.Operation{
		Name:       opDeleteRelationalDatabaseSnapshot,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteRelationalDatabaseSnapshotInput{}
	}

	req := c.newRequest(op, input, &DeleteRelationalDatabaseSnapshotOutput{})
	return DeleteRelationalDatabaseSnapshotRequest{Request: req, Input: input, Copy: c.DeleteRelationalDatabaseSnapshotRequest}
}

// DeleteRelationalDatabaseSnapshotRequest is the request type for the
// DeleteRelationalDatabaseSnapshot API operation.
type DeleteRelationalDatabaseSnapshotRequest struct {
	*aws.Request
	Input *DeleteRelationalDatabaseSnapshotInput
	Copy  func(*DeleteRelationalDatabaseSnapshotInput) DeleteRelationalDatabaseSnapshotRequest
}

// Send marshals and sends the DeleteRelationalDatabaseSnapshot API request.
func (r DeleteRelationalDatabaseSnapshotRequest) Send(ctx context.Context) (*DeleteRelationalDatabaseSnapshotResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteRelationalDatabaseSnapshotResponse{
		DeleteRelationalDatabaseSnapshotOutput: r.Request.Data.(*DeleteRelationalDatabaseSnapshotOutput),
		response:                               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteRelationalDatabaseSnapshotResponse is the response type for the
// DeleteRelationalDatabaseSnapshot API operation.
type DeleteRelationalDatabaseSnapshotResponse struct {
	*DeleteRelationalDatabaseSnapshotOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteRelationalDatabaseSnapshot request.
func (r *DeleteRelationalDatabaseSnapshotResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
