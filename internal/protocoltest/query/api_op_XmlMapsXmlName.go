// Code generated by smithy-go-codegen DO NOT EDIT.

package query

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/internal/protocoltest/query/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

func (c *Client) XmlMapsXmlName(ctx context.Context, params *XmlMapsXmlNameInput, optFns ...func(*Options)) (*XmlMapsXmlNameOutput, error) {
	stack := middleware.NewStack("XmlMapsXmlName", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	AddResolveEndpointMiddleware(stack, options)
	retry.AddRetryMiddlewares(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opXmlMapsXmlName(options.Region), middleware.Before)
	addawsAwsquery_serdeOpXmlMapsXmlNameMiddlewares(stack)

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
			OperationName: "XmlMapsXmlName",
			Err:           err,
		}
	}
	out := result.(*XmlMapsXmlNameOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type XmlMapsXmlNameInput struct {
}

type XmlMapsXmlNameOutput struct {
	MyMap map[string]*types.GreetingStruct

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpXmlMapsXmlNameMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpXmlMapsXmlName{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpXmlMapsXmlName{}, middleware.After)
}

func newServiceMetadataMiddleware_opXmlMapsXmlName(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:         region,
		ServiceName:    "Query Protocol",
		ServiceID:      "queryprotocol",
		EndpointPrefix: "queryprotocol",
		OperationName:  "XmlMapsXmlName",
	}
}
