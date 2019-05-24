// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudsearch

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Container for the parameters to the DeleteExpression operation. Specifies
// the name of the domain you want to update and the name of the expression
// you want to delete.
type DeleteExpressionInput struct {
	_ struct{} `type:"structure"`

	// A string that represents the name of a domain. Domain names are unique across
	// the domains owned by an account within an AWS region. Domain names start
	// with a letter or number and can contain the following characters: a-z (lowercase),
	// 0-9, and - (hyphen).
	//
	// DomainName is a required field
	DomainName *string `min:"3" type:"string" required:"true"`

	// The name of the Expression to delete.
	//
	// ExpressionName is a required field
	ExpressionName *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteExpressionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteExpressionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteExpressionInput"}

	if s.DomainName == nil {
		invalidParams.Add(aws.NewErrParamRequired("DomainName"))
	}
	if s.DomainName != nil && len(*s.DomainName) < 3 {
		invalidParams.Add(aws.NewErrParamMinLen("DomainName", 3))
	}

	if s.ExpressionName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ExpressionName"))
	}
	if s.ExpressionName != nil && len(*s.ExpressionName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ExpressionName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// The result of a DeleteExpression request. Specifies the expression being
// deleted.
type DeleteExpressionOutput struct {
	_ struct{} `type:"structure"`

	// The status of the expression being deleted.
	//
	// Expression is a required field
	Expression *ExpressionStatus `type:"structure" required:"true"`
}

// String returns the string representation
func (s DeleteExpressionOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteExpression = "DeleteExpression"

// DeleteExpressionRequest returns a request value for making API operation for
// Amazon CloudSearch.
//
// Removes an Expression from the search domain. For more information, see Configuring
// Expressions (http://docs.aws.amazon.com/cloudsearch/latest/developerguide/configuring-expressions.html)
// in the Amazon CloudSearch Developer Guide.
//
//    // Example sending a request using DeleteExpressionRequest.
//    req := client.DeleteExpressionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) DeleteExpressionRequest(input *DeleteExpressionInput) DeleteExpressionRequest {
	op := &aws.Operation{
		Name:       opDeleteExpression,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteExpressionInput{}
	}

	req := c.newRequest(op, input, &DeleteExpressionOutput{})
	return DeleteExpressionRequest{Request: req, Input: input, Copy: c.DeleteExpressionRequest}
}

// DeleteExpressionRequest is the request type for the
// DeleteExpression API operation.
type DeleteExpressionRequest struct {
	*aws.Request
	Input *DeleteExpressionInput
	Copy  func(*DeleteExpressionInput) DeleteExpressionRequest
}

// Send marshals and sends the DeleteExpression API request.
func (r DeleteExpressionRequest) Send(ctx context.Context) (*DeleteExpressionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteExpressionResponse{
		DeleteExpressionOutput: r.Request.Data.(*DeleteExpressionOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteExpressionResponse is the response type for the
// DeleteExpression API operation.
type DeleteExpressionResponse struct {
	*DeleteExpressionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteExpression request.
func (r *DeleteExpressionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
