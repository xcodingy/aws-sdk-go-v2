// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package wafregional

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/waf-regional-2016-11-28/DeleteSizeConstraintSetRequest
type DeleteSizeConstraintSetInput struct {
	_ struct{} `type:"structure"`

	// The value returned by the most recent call to GetChangeToken.
	//
	// ChangeToken is a required field
	ChangeToken *string `min:"1" type:"string" required:"true"`

	// The SizeConstraintSetId of the SizeConstraintSet that you want to delete.
	// SizeConstraintSetId is returned by CreateSizeConstraintSet and by ListSizeConstraintSets.
	//
	// SizeConstraintSetId is a required field
	SizeConstraintSetId *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteSizeConstraintSetInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteSizeConstraintSetInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteSizeConstraintSetInput"}

	if s.ChangeToken == nil {
		invalidParams.Add(aws.NewErrParamRequired("ChangeToken"))
	}
	if s.ChangeToken != nil && len(*s.ChangeToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ChangeToken", 1))
	}

	if s.SizeConstraintSetId == nil {
		invalidParams.Add(aws.NewErrParamRequired("SizeConstraintSetId"))
	}
	if s.SizeConstraintSetId != nil && len(*s.SizeConstraintSetId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("SizeConstraintSetId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/waf-regional-2016-11-28/DeleteSizeConstraintSetResponse
type DeleteSizeConstraintSetOutput struct {
	_ struct{} `type:"structure"`

	// The ChangeToken that you used to submit the DeleteSizeConstraintSet request.
	// You can also use this value to query the status of the request. For more
	// information, see GetChangeTokenStatus.
	ChangeToken *string `min:"1" type:"string"`
}

// String returns the string representation
func (s DeleteSizeConstraintSetOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteSizeConstraintSet = "DeleteSizeConstraintSet"

// DeleteSizeConstraintSetRequest returns a request value for making API operation for
// AWS WAF Regional.
//
// Permanently deletes a SizeConstraintSet. You can't delete a SizeConstraintSet
// if it's still used in any Rules or if it still includes any SizeConstraint
// objects (any filters).
//
// If you just want to remove a SizeConstraintSet from a Rule, use UpdateRule.
//
// To permanently delete a SizeConstraintSet, perform the following steps:
//
// Update the SizeConstraintSet to remove filters, if any. For more information,
// see UpdateSizeConstraintSet.
//
// Use GetChangeToken to get the change token that you provide in the ChangeToken
// parameter of a DeleteSizeConstraintSet request.
//
// Submit a DeleteSizeConstraintSet request.
//
//    // Example sending a request using DeleteSizeConstraintSetRequest.
//    req := client.DeleteSizeConstraintSetRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/waf-regional-2016-11-28/DeleteSizeConstraintSet
func (c *Client) DeleteSizeConstraintSetRequest(input *DeleteSizeConstraintSetInput) DeleteSizeConstraintSetRequest {
	op := &aws.Operation{
		Name:       opDeleteSizeConstraintSet,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteSizeConstraintSetInput{}
	}

	req := c.newRequest(op, input, &DeleteSizeConstraintSetOutput{})
	return DeleteSizeConstraintSetRequest{Request: req, Input: input, Copy: c.DeleteSizeConstraintSetRequest}
}

// DeleteSizeConstraintSetRequest is the request type for the
// DeleteSizeConstraintSet API operation.
type DeleteSizeConstraintSetRequest struct {
	*aws.Request
	Input *DeleteSizeConstraintSetInput
	Copy  func(*DeleteSizeConstraintSetInput) DeleteSizeConstraintSetRequest
}

// Send marshals and sends the DeleteSizeConstraintSet API request.
func (r DeleteSizeConstraintSetRequest) Send(ctx context.Context) (*DeleteSizeConstraintSetResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteSizeConstraintSetResponse{
		DeleteSizeConstraintSetOutput: r.Request.Data.(*DeleteSizeConstraintSetOutput),
		response:                      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteSizeConstraintSetResponse is the response type for the
// DeleteSizeConstraintSet API operation.
type DeleteSizeConstraintSetResponse struct {
	*DeleteSizeConstraintSetOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteSizeConstraintSet request.
func (r *DeleteSizeConstraintSetResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
