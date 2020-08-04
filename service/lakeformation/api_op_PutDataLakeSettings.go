// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package lakeformation

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type PutDataLakeSettingsInput struct {
	_ struct{} `type:"structure"`

	// The identifier for the Data Catalog. By default, the account ID. The Data
	// Catalog is the persistent metadata store. It contains database definitions,
	// table definitions, and other control information to manage your AWS Lake
	// Formation environment.
	CatalogId *string `min:"1" type:"string"`

	// A structure representing a list of AWS Lake Formation principals designated
	// as data lake administrators.
	//
	// DataLakeSettings is a required field
	DataLakeSettings *DataLakeSettings `type:"structure" required:"true"`
}

// String returns the string representation
func (s PutDataLakeSettingsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PutDataLakeSettingsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "PutDataLakeSettingsInput"}
	if s.CatalogId != nil && len(*s.CatalogId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("CatalogId", 1))
	}

	if s.DataLakeSettings == nil {
		invalidParams.Add(aws.NewErrParamRequired("DataLakeSettings"))
	}
	if s.DataLakeSettings != nil {
		if err := s.DataLakeSettings.Validate(); err != nil {
			invalidParams.AddNested("DataLakeSettings", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type PutDataLakeSettingsOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s PutDataLakeSettingsOutput) String() string {
	return awsutil.Prettify(s)
}

const opPutDataLakeSettings = "PutDataLakeSettings"

// PutDataLakeSettingsRequest returns a request value for making API operation for
// AWS Lake Formation.
//
// Sets the list of data lake administrators who have admin privileges on all
// resources managed by Lake Formation. For more information on admin privileges,
// see Granting Lake Formation Permissions (https://docs.aws.amazon.com/lake-formation/latest/dg/lake-formation-permissions.html).
//
// This API replaces the current list of data lake admins with the new list
// being passed. To add an admin, fetch the current list and add the new admin
// to that list and pass that list in this API.
//
//    // Example sending a request using PutDataLakeSettingsRequest.
//    req := client.PutDataLakeSettingsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/lakeformation-2017-03-31/PutDataLakeSettings
func (c *Client) PutDataLakeSettingsRequest(input *PutDataLakeSettingsInput) PutDataLakeSettingsRequest {
	op := &aws.Operation{
		Name:       opPutDataLakeSettings,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &PutDataLakeSettingsInput{}
	}

	req := c.newRequest(op, input, &PutDataLakeSettingsOutput{})

	return PutDataLakeSettingsRequest{Request: req, Input: input, Copy: c.PutDataLakeSettingsRequest}
}

// PutDataLakeSettingsRequest is the request type for the
// PutDataLakeSettings API operation.
type PutDataLakeSettingsRequest struct {
	*aws.Request
	Input *PutDataLakeSettingsInput
	Copy  func(*PutDataLakeSettingsInput) PutDataLakeSettingsRequest
}

// Send marshals and sends the PutDataLakeSettings API request.
func (r PutDataLakeSettingsRequest) Send(ctx context.Context) (*PutDataLakeSettingsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &PutDataLakeSettingsResponse{
		PutDataLakeSettingsOutput: r.Request.Data.(*PutDataLakeSettingsOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// PutDataLakeSettingsResponse is the response type for the
// PutDataLakeSettings API operation.
type PutDataLakeSettingsResponse struct {
	*PutDataLakeSettingsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// PutDataLakeSettings request.
func (r *PutDataLakeSettingsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}