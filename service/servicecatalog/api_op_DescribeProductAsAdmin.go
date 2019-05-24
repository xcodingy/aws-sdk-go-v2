// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package servicecatalog

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/servicecatalog-2015-12-10/DescribeProductAsAdminInput
type DescribeProductAsAdminInput struct {
	_ struct{} `type:"structure"`

	// The language code.
	//
	//    * en - English (default)
	//
	//    * jp - Japanese
	//
	//    * zh - Chinese
	AcceptLanguage *string `type:"string"`

	// The product identifier.
	//
	// Id is a required field
	Id *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeProductAsAdminInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeProductAsAdminInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeProductAsAdminInput"}

	if s.Id == nil {
		invalidParams.Add(aws.NewErrParamRequired("Id"))
	}
	if s.Id != nil && len(*s.Id) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Id", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/servicecatalog-2015-12-10/DescribeProductAsAdminOutput
type DescribeProductAsAdminOutput struct {
	_ struct{} `type:"structure"`

	// Information about the product view.
	ProductViewDetail *ProductViewDetail `type:"structure"`

	// Information about the provisioning artifacts (also known as versions) for
	// the specified product.
	ProvisioningArtifactSummaries []ProvisioningArtifactSummary `type:"list"`

	// Information about the TagOptions associated with the product.
	TagOptions []TagOptionDetail `type:"list"`

	// Information about the tags associated with the product.
	Tags []Tag `type:"list"`
}

// String returns the string representation
func (s DescribeProductAsAdminOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeProductAsAdmin = "DescribeProductAsAdmin"

// DescribeProductAsAdminRequest returns a request value for making API operation for
// AWS Service Catalog.
//
// Gets information about the specified product. This operation is run with
// administrator access.
//
//    // Example sending a request using DescribeProductAsAdminRequest.
//    req := client.DescribeProductAsAdminRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/servicecatalog-2015-12-10/DescribeProductAsAdmin
func (c *Client) DescribeProductAsAdminRequest(input *DescribeProductAsAdminInput) DescribeProductAsAdminRequest {
	op := &aws.Operation{
		Name:       opDescribeProductAsAdmin,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeProductAsAdminInput{}
	}

	req := c.newRequest(op, input, &DescribeProductAsAdminOutput{})
	return DescribeProductAsAdminRequest{Request: req, Input: input, Copy: c.DescribeProductAsAdminRequest}
}

// DescribeProductAsAdminRequest is the request type for the
// DescribeProductAsAdmin API operation.
type DescribeProductAsAdminRequest struct {
	*aws.Request
	Input *DescribeProductAsAdminInput
	Copy  func(*DescribeProductAsAdminInput) DescribeProductAsAdminRequest
}

// Send marshals and sends the DescribeProductAsAdmin API request.
func (r DescribeProductAsAdminRequest) Send(ctx context.Context) (*DescribeProductAsAdminResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeProductAsAdminResponse{
		DescribeProductAsAdminOutput: r.Request.Data.(*DescribeProductAsAdminOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeProductAsAdminResponse is the response type for the
// DescribeProductAsAdmin API operation.
type DescribeProductAsAdminResponse struct {
	*DescribeProductAsAdminOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeProductAsAdmin request.
func (r *DescribeProductAsAdminResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
