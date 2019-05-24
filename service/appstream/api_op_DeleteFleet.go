// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package appstream

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/appstream-2016-12-01/DeleteFleetRequest
type DeleteFleetInput struct {
	_ struct{} `type:"structure"`

	// The name of the fleet.
	//
	// Name is a required field
	Name *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteFleetInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteFleetInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteFleetInput"}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}
	if s.Name != nil && len(*s.Name) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Name", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/appstream-2016-12-01/DeleteFleetResult
type DeleteFleetOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteFleetOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteFleet = "DeleteFleet"

// DeleteFleetRequest returns a request value for making API operation for
// Amazon AppStream.
//
// Deletes the specified fleet.
//
//    // Example sending a request using DeleteFleetRequest.
//    req := client.DeleteFleetRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/appstream-2016-12-01/DeleteFleet
func (c *Client) DeleteFleetRequest(input *DeleteFleetInput) DeleteFleetRequest {
	op := &aws.Operation{
		Name:       opDeleteFleet,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteFleetInput{}
	}

	req := c.newRequest(op, input, &DeleteFleetOutput{})
	return DeleteFleetRequest{Request: req, Input: input, Copy: c.DeleteFleetRequest}
}

// DeleteFleetRequest is the request type for the
// DeleteFleet API operation.
type DeleteFleetRequest struct {
	*aws.Request
	Input *DeleteFleetInput
	Copy  func(*DeleteFleetInput) DeleteFleetRequest
}

// Send marshals and sends the DeleteFleet API request.
func (r DeleteFleetRequest) Send(ctx context.Context) (*DeleteFleetResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteFleetResponse{
		DeleteFleetOutput: r.Request.Data.(*DeleteFleetOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteFleetResponse is the response type for the
// DeleteFleet API operation.
type DeleteFleetResponse struct {
	*DeleteFleetOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteFleet request.
func (r *DeleteFleetResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
