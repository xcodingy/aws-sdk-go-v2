// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package backup

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/backup-2018-11-15/PutBackupVaultNotificationsInput
type PutBackupVaultNotificationsInput struct {
	_ struct{} `type:"structure"`

	// An array of events that indicate the status of jobs to back up resources
	// to the backup vault.
	//
	// BackupVaultEvents is a required field
	BackupVaultEvents []BackupVaultEvent `type:"list" required:"true"`

	// The name of a logical container where backups are stored. Backup vaults are
	// identified by names that are unique to the account used to create them and
	// the AWS Region where they are created. They consist of lowercase letters,
	// numbers, and hyphens.
	//
	// BackupVaultName is a required field
	BackupVaultName *string `location:"uri" locationName:"backupVaultName" type:"string" required:"true"`

	// The Amazon Resource Name (ARN) that specifies the topic for a backup vault’s
	// events; for example, arn:aws:sns:us-west-2:111122223333:MyVaultTopic.
	//
	// SNSTopicArn is a required field
	SNSTopicArn *string `type:"string" required:"true"`
}

// String returns the string representation
func (s PutBackupVaultNotificationsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PutBackupVaultNotificationsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "PutBackupVaultNotificationsInput"}

	if s.BackupVaultEvents == nil {
		invalidParams.Add(aws.NewErrParamRequired("BackupVaultEvents"))
	}

	if s.BackupVaultName == nil {
		invalidParams.Add(aws.NewErrParamRequired("BackupVaultName"))
	}

	if s.SNSTopicArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("SNSTopicArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s PutBackupVaultNotificationsInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	if len(s.BackupVaultEvents) > 0 {
		v := s.BackupVaultEvents

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "BackupVaultEvents", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	if s.SNSTopicArn != nil {
		v := *s.SNSTopicArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "SNSTopicArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.BackupVaultName != nil {
		v := *s.BackupVaultName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "backupVaultName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/backup-2018-11-15/PutBackupVaultNotificationsOutput
type PutBackupVaultNotificationsOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s PutBackupVaultNotificationsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s PutBackupVaultNotificationsOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opPutBackupVaultNotifications = "PutBackupVaultNotifications"

// PutBackupVaultNotificationsRequest returns a request value for making API operation for
// AWS Backup.
//
// Turns on notifications on a backup vault for the specified topic and events.
//
//    // Example sending a request using PutBackupVaultNotificationsRequest.
//    req := client.PutBackupVaultNotificationsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/backup-2018-11-15/PutBackupVaultNotifications
func (c *Client) PutBackupVaultNotificationsRequest(input *PutBackupVaultNotificationsInput) PutBackupVaultNotificationsRequest {
	op := &aws.Operation{
		Name:       opPutBackupVaultNotifications,
		HTTPMethod: "PUT",
		HTTPPath:   "/backup-vaults/{backupVaultName}/notification-configuration",
	}

	if input == nil {
		input = &PutBackupVaultNotificationsInput{}
	}

	req := c.newRequest(op, input, &PutBackupVaultNotificationsOutput{})
	req.Handlers.Unmarshal.Remove(restjson.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return PutBackupVaultNotificationsRequest{Request: req, Input: input, Copy: c.PutBackupVaultNotificationsRequest}
}

// PutBackupVaultNotificationsRequest is the request type for the
// PutBackupVaultNotifications API operation.
type PutBackupVaultNotificationsRequest struct {
	*aws.Request
	Input *PutBackupVaultNotificationsInput
	Copy  func(*PutBackupVaultNotificationsInput) PutBackupVaultNotificationsRequest
}

// Send marshals and sends the PutBackupVaultNotifications API request.
func (r PutBackupVaultNotificationsRequest) Send(ctx context.Context) (*PutBackupVaultNotificationsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &PutBackupVaultNotificationsResponse{
		PutBackupVaultNotificationsOutput: r.Request.Data.(*PutBackupVaultNotificationsOutput),
		response:                          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// PutBackupVaultNotificationsResponse is the response type for the
// PutBackupVaultNotifications API operation.
type PutBackupVaultNotificationsResponse struct {
	*PutBackupVaultNotificationsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// PutBackupVaultNotifications request.
func (r *PutBackupVaultNotificationsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
