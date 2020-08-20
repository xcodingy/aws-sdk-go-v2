// Code generated by smithy-go-codegen DO NOT EDIT.

package dynamodb

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a new table from an existing backup. Any number of users can execute up
// to 4 concurrent restores (any type of restore) in a given account. You can call
// RestoreTableFromBackup at a maximum rate of 10 times per second. You must
// manually set up the following on the restored table:
//
//     * Auto scaling
// policies
//
//     * IAM policies
//
//     * Amazon CloudWatch metrics and alarms
//
//     *
// Tags
//
//     * Stream settings
//
//     * Time to Live (TTL) settings
func (c *Client) RestoreTableFromBackup(ctx context.Context, params *RestoreTableFromBackupInput, optFns ...func(*Options)) (*RestoreTableFromBackupOutput, error) {
	stack := middleware.NewStack("RestoreTableFromBackup", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson10_serdeOpRestoreTableFromBackupMiddlewares(stack)
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	AddResolveEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	retry.AddRetryMiddlewares(stack, options)
	v4.AddHTTPSignerMiddleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	addOpRestoreTableFromBackupValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opRestoreTableFromBackup(options.Region), middleware.Before)

	for _, fn := range options.APIOptions {
		if err := fn(stack); err != nil {
			return nil, err
		}
	}
	handler := middleware.DecorateHandler(smithyhttp.NewClientHandler(options.HTTPClient), stack)
	result, metadata, err := handler.Handle(ctx, params)
	if err != nil {
		return nil, &smithy.OperationError{
			ServiceID:     c.ServiceID(),
			OperationName: "RestoreTableFromBackup",
			Err:           err,
		}
	}
	out := result.(*RestoreTableFromBackupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RestoreTableFromBackupInput struct {
	// The Amazon Resource Name (ARN) associated with the backup.
	BackupArn *string
	// The billing mode of the restored table.
	BillingModeOverride types.BillingMode
	// List of global secondary indexes for the restored table. The indexes provided
	// should match existing secondary indexes. You can choose to exclude some or all
	// of the indexes at the time of restore.
	GlobalSecondaryIndexOverride []*types.GlobalSecondaryIndex
	// List of local secondary indexes for the restored table. The indexes provided
	// should match existing secondary indexes. You can choose to exclude some or all
	// of the indexes at the time of restore.
	LocalSecondaryIndexOverride []*types.LocalSecondaryIndex
	// Provisioned throughput settings for the restored table.
	ProvisionedThroughputOverride *types.ProvisionedThroughput
	// The name of the new table to which the backup must be restored.
	TargetTableName *string
}

type RestoreTableFromBackupOutput struct {
	// The description of the table created from an existing backup.
	TableDescription *types.TableDescription

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson10_serdeOpRestoreTableFromBackupMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson10_serializeOpRestoreTableFromBackup{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson10_deserializeOpRestoreTableFromBackup{}, middleware.After)
}

func newServiceMetadataMiddleware_opRestoreTableFromBackup(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:         region,
		ServiceName:    "DynamoDB GoV2",
		ServiceID:      "dynamodbgov2",
		EndpointPrefix: "dynamodbgov2",
		SigningName:    "dynamodb",
		OperationName:  "RestoreTableFromBackup",
	}
}
