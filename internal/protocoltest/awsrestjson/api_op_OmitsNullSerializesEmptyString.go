// Code generated by smithy-go-codegen DO NOT EDIT.
package awsrestjson

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Omits null, but serializes empty string value.
func (c *Client) OmitsNullSerializesEmptyString(ctx context.Context, params *OmitsNullSerializesEmptyStringInput, optFns ...func(*Options)) (*OmitsNullSerializesEmptyStringOutput, error) {
	stack := middleware.NewStack("OmitsNullSerializesEmptyString", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	stack.Build.Add(awsmiddleware.RequestInvocationIDMiddleware{}, middleware.After)
	stack.Serialize.Add(&awsRestjson1_serializeOpOmitsNullSerializesEmptyString{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpOmitsNullSerializesEmptyString{}, middleware.After)
	stack.Deserialize.Add(awsmiddleware.AttemptClockSkewMiddleware{}, middleware.After)
	stack.Finalize.Add(retry.NewAttemptMiddleware(options.Retryer, smithyhttp.RequestCloner), middleware.After)
	stack.Finalize.Add(retry.MetricsHeaderMiddleware{}, middleware.After)

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
			OperationName: "OmitsNullSerializesEmptyString",
			Err:           err,
		}
	}
	out := result.(*OmitsNullSerializesEmptyStringOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type OmitsNullSerializesEmptyStringInput struct {
	EmptyString *string
	NullValue   *string
}

type OmitsNullSerializesEmptyStringOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}