// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package backup

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/backup-2018-11-15/GetBackupVaultAccessPolicyInput
type GetBackupVaultAccessPolicyInput struct {
	_ struct{} `type:"structure"`

	// The name of a logical container where backups are stored. Backup vaults are
	// identified by names that are unique to the account used to create them and
	// the AWS Region where they are created. They consist of lowercase letters,
	// numbers, and hyphens.
	//
	// BackupVaultName is a required field
	BackupVaultName *string `location:"uri" locationName:"backupVaultName" type:"string" required:"true"`
}

// String returns the string representation
func (s GetBackupVaultAccessPolicyInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetBackupVaultAccessPolicyInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetBackupVaultAccessPolicyInput"}

	if s.BackupVaultName == nil {
		invalidParams.Add(aws.NewErrParamRequired("BackupVaultName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetBackupVaultAccessPolicyInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	if s.BackupVaultName != nil {
		v := *s.BackupVaultName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "backupVaultName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/backup-2018-11-15/GetBackupVaultAccessPolicyOutput
type GetBackupVaultAccessPolicyOutput struct {
	_ struct{} `type:"structure"`

	// An Amazon Resource Name (ARN) that uniquely identifies a backup vault; for
	// example, arn:aws:backup:us-east-1:123456789012:vault:aBackupVault.
	BackupVaultArn *string `type:"string"`

	// The name of a logical container where backups are stored. Backup vaults are
	// identified by names that are unique to the account used to create them and
	// the Region where they are created. They consist of lowercase letters, numbers,
	// and hyphens.
	BackupVaultName *string `type:"string"`

	// The backup vault access policy document in JSON format.
	Policy *string `type:"string"`
}

// String returns the string representation
func (s GetBackupVaultAccessPolicyOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetBackupVaultAccessPolicyOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.BackupVaultArn != nil {
		v := *s.BackupVaultArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "BackupVaultArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.BackupVaultName != nil {
		v := *s.BackupVaultName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "BackupVaultName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Policy != nil {
		v := *s.Policy

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Policy", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opGetBackupVaultAccessPolicy = "GetBackupVaultAccessPolicy"

// GetBackupVaultAccessPolicyRequest returns a request value for making API operation for
// AWS Backup.
//
// Returns the access policy document that is associated with the named backup
// vault.
//
//    // Example sending a request using GetBackupVaultAccessPolicyRequest.
//    req := client.GetBackupVaultAccessPolicyRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/backup-2018-11-15/GetBackupVaultAccessPolicy
func (c *Client) GetBackupVaultAccessPolicyRequest(input *GetBackupVaultAccessPolicyInput) GetBackupVaultAccessPolicyRequest {
	op := &aws.Operation{
		Name:       opGetBackupVaultAccessPolicy,
		HTTPMethod: "GET",
		HTTPPath:   "/backup-vaults/{backupVaultName}/access-policy",
	}

	if input == nil {
		input = &GetBackupVaultAccessPolicyInput{}
	}

	req := c.newRequest(op, input, &GetBackupVaultAccessPolicyOutput{})
	return GetBackupVaultAccessPolicyRequest{Request: req, Input: input, Copy: c.GetBackupVaultAccessPolicyRequest}
}

// GetBackupVaultAccessPolicyRequest is the request type for the
// GetBackupVaultAccessPolicy API operation.
type GetBackupVaultAccessPolicyRequest struct {
	*aws.Request
	Input *GetBackupVaultAccessPolicyInput
	Copy  func(*GetBackupVaultAccessPolicyInput) GetBackupVaultAccessPolicyRequest
}

// Send marshals and sends the GetBackupVaultAccessPolicy API request.
func (r GetBackupVaultAccessPolicyRequest) Send(ctx context.Context) (*GetBackupVaultAccessPolicyResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetBackupVaultAccessPolicyResponse{
		GetBackupVaultAccessPolicyOutput: r.Request.Data.(*GetBackupVaultAccessPolicyOutput),
		response:                         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetBackupVaultAccessPolicyResponse is the response type for the
// GetBackupVaultAccessPolicy API operation.
type GetBackupVaultAccessPolicyResponse struct {
	*GetBackupVaultAccessPolicyOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetBackupVaultAccessPolicy request.
func (r *GetBackupVaultAccessPolicyResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
