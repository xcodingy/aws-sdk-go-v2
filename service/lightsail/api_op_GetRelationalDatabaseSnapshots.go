// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package lightsail

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/lightsail-2016-11-28/GetRelationalDatabaseSnapshotsRequest
type GetRelationalDatabaseSnapshotsInput struct {
	_ struct{} `type:"structure"`

	// A token used for advancing to a specific page of results for your get relational
	// database snapshots request.
	PageToken *string `locationName:"pageToken" type:"string"`
}

// String returns the string representation
func (s GetRelationalDatabaseSnapshotsInput) String() string {
	return awsutil.Prettify(s)
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/lightsail-2016-11-28/GetRelationalDatabaseSnapshotsResult
type GetRelationalDatabaseSnapshotsOutput struct {
	_ struct{} `type:"structure"`

	// A token used for advancing to the next page of results from your get relational
	// database snapshots request.
	NextPageToken *string `locationName:"nextPageToken" type:"string"`

	// An object describing the result of your get relational database snapshots
	// request.
	RelationalDatabaseSnapshots []RelationalDatabaseSnapshot `locationName:"relationalDatabaseSnapshots" type:"list"`
}

// String returns the string representation
func (s GetRelationalDatabaseSnapshotsOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetRelationalDatabaseSnapshots = "GetRelationalDatabaseSnapshots"

// GetRelationalDatabaseSnapshotsRequest returns a request value for making API operation for
// Amazon Lightsail.
//
// Returns information about all of your database snapshots in Amazon Lightsail.
//
//    // Example sending a request using GetRelationalDatabaseSnapshotsRequest.
//    req := client.GetRelationalDatabaseSnapshotsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/lightsail-2016-11-28/GetRelationalDatabaseSnapshots
func (c *Client) GetRelationalDatabaseSnapshotsRequest(input *GetRelationalDatabaseSnapshotsInput) GetRelationalDatabaseSnapshotsRequest {
	op := &aws.Operation{
		Name:       opGetRelationalDatabaseSnapshots,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetRelationalDatabaseSnapshotsInput{}
	}

	req := c.newRequest(op, input, &GetRelationalDatabaseSnapshotsOutput{})
	return GetRelationalDatabaseSnapshotsRequest{Request: req, Input: input, Copy: c.GetRelationalDatabaseSnapshotsRequest}
}

// GetRelationalDatabaseSnapshotsRequest is the request type for the
// GetRelationalDatabaseSnapshots API operation.
type GetRelationalDatabaseSnapshotsRequest struct {
	*aws.Request
	Input *GetRelationalDatabaseSnapshotsInput
	Copy  func(*GetRelationalDatabaseSnapshotsInput) GetRelationalDatabaseSnapshotsRequest
}

// Send marshals and sends the GetRelationalDatabaseSnapshots API request.
func (r GetRelationalDatabaseSnapshotsRequest) Send(ctx context.Context) (*GetRelationalDatabaseSnapshotsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetRelationalDatabaseSnapshotsResponse{
		GetRelationalDatabaseSnapshotsOutput: r.Request.Data.(*GetRelationalDatabaseSnapshotsOutput),
		response:                             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetRelationalDatabaseSnapshotsResponse is the response type for the
// GetRelationalDatabaseSnapshots API operation.
type GetRelationalDatabaseSnapshotsResponse struct {
	*GetRelationalDatabaseSnapshotsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetRelationalDatabaseSnapshots request.
func (r *GetRelationalDatabaseSnapshotsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
