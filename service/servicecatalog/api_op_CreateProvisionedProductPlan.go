// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package servicecatalog

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/servicecatalog-2015-12-10/CreateProvisionedProductPlanInput
type CreateProvisionedProductPlanInput struct {
	_ struct{} `type:"structure"`

	// The language code.
	//
	//    * en - English (default)
	//
	//    * jp - Japanese
	//
	//    * zh - Chinese
	AcceptLanguage *string `type:"string"`

	// A unique identifier that you provide to ensure idempotency. If multiple requests
	// differ only by the idempotency token, the same response is returned for each
	// repeated request.
	//
	// IdempotencyToken is a required field
	IdempotencyToken *string `min:"1" type:"string" required:"true" idempotencyToken:"true"`

	// Passed to CloudFormation. The SNS topic ARNs to which to publish stack-related
	// events.
	NotificationArns []string `type:"list"`

	// The path identifier of the product. This value is optional if the product
	// has a default path, and required if the product has more than one path. To
	// list the paths for a product, use ListLaunchPaths.
	PathId *string `min:"1" type:"string"`

	// The name of the plan.
	//
	// PlanName is a required field
	PlanName *string `type:"string" required:"true"`

	// The plan type.
	//
	// PlanType is a required field
	PlanType ProvisionedProductPlanType `type:"string" required:"true" enum:"true"`

	// The product identifier.
	//
	// ProductId is a required field
	ProductId *string `min:"1" type:"string" required:"true"`

	// A user-friendly name for the provisioned product. This value must be unique
	// for the AWS account and cannot be updated after the product is provisioned.
	//
	// ProvisionedProductName is a required field
	ProvisionedProductName *string `min:"1" type:"string" required:"true"`

	// The identifier of the provisioning artifact.
	//
	// ProvisioningArtifactId is a required field
	ProvisioningArtifactId *string `min:"1" type:"string" required:"true"`

	// Parameters specified by the administrator that are required for provisioning
	// the product.
	ProvisioningParameters []UpdateProvisioningParameter `type:"list"`

	// One or more tags.
	Tags []Tag `type:"list"`
}

// String returns the string representation
func (s CreateProvisionedProductPlanInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateProvisionedProductPlanInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateProvisionedProductPlanInput"}

	if s.IdempotencyToken == nil {
		invalidParams.Add(aws.NewErrParamRequired("IdempotencyToken"))
	}
	if s.IdempotencyToken != nil && len(*s.IdempotencyToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("IdempotencyToken", 1))
	}
	if s.PathId != nil && len(*s.PathId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("PathId", 1))
	}

	if s.PlanName == nil {
		invalidParams.Add(aws.NewErrParamRequired("PlanName"))
	}
	if len(s.PlanType) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("PlanType"))
	}

	if s.ProductId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ProductId"))
	}
	if s.ProductId != nil && len(*s.ProductId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ProductId", 1))
	}

	if s.ProvisionedProductName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ProvisionedProductName"))
	}
	if s.ProvisionedProductName != nil && len(*s.ProvisionedProductName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ProvisionedProductName", 1))
	}

	if s.ProvisioningArtifactId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ProvisioningArtifactId"))
	}
	if s.ProvisioningArtifactId != nil && len(*s.ProvisioningArtifactId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ProvisioningArtifactId", 1))
	}
	if s.ProvisioningParameters != nil {
		for i, v := range s.ProvisioningParameters {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "ProvisioningParameters", i), err.(aws.ErrInvalidParams))
			}
		}
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

// Please also see https://docs.aws.amazon.com/goto/WebAPI/servicecatalog-2015-12-10/CreateProvisionedProductPlanOutput
type CreateProvisionedProductPlanOutput struct {
	_ struct{} `type:"structure"`

	// The plan identifier.
	PlanId *string `min:"1" type:"string"`

	// The name of the plan.
	PlanName *string `type:"string"`

	// The product identifier.
	ProvisionProductId *string `min:"1" type:"string"`

	// The user-friendly name of the provisioned product.
	ProvisionedProductName *string `min:"1" type:"string"`

	// The identifier of the provisioning artifact.
	ProvisioningArtifactId *string `min:"1" type:"string"`
}

// String returns the string representation
func (s CreateProvisionedProductPlanOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateProvisionedProductPlan = "CreateProvisionedProductPlan"

// CreateProvisionedProductPlanRequest returns a request value for making API operation for
// AWS Service Catalog.
//
// Creates a plan. A plan includes the list of resources to be created (when
// provisioning a new product) or modified (when updating a provisioned product)
// when the plan is executed.
//
// You can create one plan per provisioned product. To create a plan for an
// existing provisioned product, the product status must be AVAILBLE or TAINTED.
//
// To view the resource changes in the change set, use DescribeProvisionedProductPlan.
// To create or modify the provisioned product, use ExecuteProvisionedProductPlan.
//
//    // Example sending a request using CreateProvisionedProductPlanRequest.
//    req := client.CreateProvisionedProductPlanRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/servicecatalog-2015-12-10/CreateProvisionedProductPlan
func (c *Client) CreateProvisionedProductPlanRequest(input *CreateProvisionedProductPlanInput) CreateProvisionedProductPlanRequest {
	op := &aws.Operation{
		Name:       opCreateProvisionedProductPlan,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateProvisionedProductPlanInput{}
	}

	req := c.newRequest(op, input, &CreateProvisionedProductPlanOutput{})
	return CreateProvisionedProductPlanRequest{Request: req, Input: input, Copy: c.CreateProvisionedProductPlanRequest}
}

// CreateProvisionedProductPlanRequest is the request type for the
// CreateProvisionedProductPlan API operation.
type CreateProvisionedProductPlanRequest struct {
	*aws.Request
	Input *CreateProvisionedProductPlanInput
	Copy  func(*CreateProvisionedProductPlanInput) CreateProvisionedProductPlanRequest
}

// Send marshals and sends the CreateProvisionedProductPlan API request.
func (r CreateProvisionedProductPlanRequest) Send(ctx context.Context) (*CreateProvisionedProductPlanResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateProvisionedProductPlanResponse{
		CreateProvisionedProductPlanOutput: r.Request.Data.(*CreateProvisionedProductPlanOutput),
		response:                           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateProvisionedProductPlanResponse is the response type for the
// CreateProvisionedProductPlan API operation.
type CreateProvisionedProductPlanResponse struct {
	*CreateProvisionedProductPlanOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateProvisionedProductPlan request.
func (r *CreateProvisionedProductPlanResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
