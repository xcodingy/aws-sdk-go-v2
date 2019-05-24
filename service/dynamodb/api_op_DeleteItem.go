// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package dynamodb

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents the input of a DeleteItem operation.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/dynamodb-2012-08-10/DeleteItemInput
type DeleteItemInput struct {
	_ struct{} `type:"structure"`

	// A condition that must be satisfied in order for a conditional DeleteItem
	// to succeed.
	//
	// An expression can contain any of the following:
	//
	//    * Functions: attribute_exists | attribute_not_exists | attribute_type
	//    | contains | begins_with | size
	//
	// These function names are case-sensitive.
	//
	//    * Comparison operators: = | <> | < | > | <= | >= | BETWEEN | IN
	//
	//    *  Logical operators: AND | OR | NOT
	//
	// For more information on condition expressions, see Specifying Conditions
	// (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Expressions.SpecifyingConditions.html)
	// in the Amazon DynamoDB Developer Guide.
	ConditionExpression *string `type:"string"`

	// This is a legacy parameter. Use ConditionExpression instead. For more information,
	// see ConditionalOperator (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/LegacyConditionalParameters.ConditionalOperator.html)
	// in the Amazon DynamoDB Developer Guide.
	ConditionalOperator ConditionalOperator `type:"string" enum:"true"`

	// This is a legacy parameter. Use ConditionExpression instead. For more information,
	// see Expected (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/LegacyConditionalParameters.Expected.html)
	// in the Amazon DynamoDB Developer Guide.
	Expected map[string]ExpectedAttributeValue `type:"map"`

	// One or more substitution tokens for attribute names in an expression. The
	// following are some use cases for using ExpressionAttributeNames:
	//
	//    * To access an attribute whose name conflicts with a DynamoDB reserved
	//    word.
	//
	//    * To create a placeholder for repeating occurrences of an attribute name
	//    in an expression.
	//
	//    * To prevent special characters in an attribute name from being misinterpreted
	//    in an expression.
	//
	// Use the # character in an expression to dereference an attribute name. For
	// example, consider the following attribute name:
	//
	//    * Percentile
	//
	// The name of this attribute conflicts with a reserved word, so it cannot be
	// used directly in an expression. (For the complete list of reserved words,
	// see Reserved Words (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/ReservedWords.html)
	// in the Amazon DynamoDB Developer Guide). To work around this, you could specify
	// the following for ExpressionAttributeNames:
	//
	//    * {"#P":"Percentile"}
	//
	// You could then use this substitution in an expression, as in this example:
	//
	//    * #P = :val
	//
	// Tokens that begin with the : character are expression attribute values, which
	// are placeholders for the actual value at runtime.
	//
	// For more information on expression attribute names, see Accessing Item Attributes
	// (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Expressions.AccessingItemAttributes.html)
	// in the Amazon DynamoDB Developer Guide.
	ExpressionAttributeNames map[string]string `type:"map"`

	// One or more values that can be substituted in an expression.
	//
	// Use the : (colon) character in an expression to dereference an attribute
	// value. For example, suppose that you wanted to check whether the value of
	// the ProductStatus attribute was one of the following:
	//
	// Available | Backordered | Discontinued
	//
	// You would first need to specify ExpressionAttributeValues as follows:
	//
	// { ":avail":{"S":"Available"}, ":back":{"S":"Backordered"}, ":disc":{"S":"Discontinued"}
	// }
	//
	// You could then use these values in an expression, such as this:
	//
	// ProductStatus IN (:avail, :back, :disc)
	//
	// For more information on expression attribute values, see Specifying Conditions
	// (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Expressions.SpecifyingConditions.html)
	// in the Amazon DynamoDB Developer Guide.
	ExpressionAttributeValues map[string]AttributeValue `type:"map"`

	// A map of attribute names to AttributeValue objects, representing the primary
	// key of the item to delete.
	//
	// For the primary key, you must provide all of the attributes. For example,
	// with a simple primary key, you only need to provide a value for the partition
	// key. For a composite primary key, you must provide values for both the partition
	// key and the sort key.
	//
	// Key is a required field
	Key map[string]AttributeValue `type:"map" required:"true"`

	// Determines the level of detail about provisioned throughput consumption that
	// is returned in the response:
	//
	//    * INDEXES - The response includes the aggregate ConsumedCapacity for the
	//    operation, together with ConsumedCapacity for each table and secondary
	//    index that was accessed.
	//
	// Note that some operations, such as GetItem and BatchGetItem, do not access
	//    any indexes at all. In these cases, specifying INDEXES will only return
	//    ConsumedCapacity information for table(s).
	//
	//    * TOTAL - The response includes only the aggregate ConsumedCapacity for
	//    the operation.
	//
	//    * NONE - No ConsumedCapacity details are included in the response.
	ReturnConsumedCapacity ReturnConsumedCapacity `type:"string" enum:"true"`

	// Determines whether item collection metrics are returned. If set to SIZE,
	// the response includes statistics about item collections, if any, that were
	// modified during the operation are returned in the response. If set to NONE
	// (the default), no statistics are returned.
	ReturnItemCollectionMetrics ReturnItemCollectionMetrics `type:"string" enum:"true"`

	// Use ReturnValues if you want to get the item attributes as they appeared
	// before they were deleted. For DeleteItem, the valid values are:
	//
	//    * NONE - If ReturnValues is not specified, or if its value is NONE, then
	//    nothing is returned. (This setting is the default for ReturnValues.)
	//
	//    * ALL_OLD - The content of the old item is returned.
	//
	// The ReturnValues parameter is used by several DynamoDB operations; however,
	// DeleteItem does not recognize any values other than NONE or ALL_OLD.
	ReturnValues ReturnValue `type:"string" enum:"true"`

	// The name of the table from which to delete the item.
	//
	// TableName is a required field
	TableName *string `min:"3" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteItemInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteItemInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteItemInput"}

	if s.Key == nil {
		invalidParams.Add(aws.NewErrParamRequired("Key"))
	}

	if s.TableName == nil {
		invalidParams.Add(aws.NewErrParamRequired("TableName"))
	}
	if s.TableName != nil && len(*s.TableName) < 3 {
		invalidParams.Add(aws.NewErrParamMinLen("TableName", 3))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents the output of a DeleteItem operation.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/dynamodb-2012-08-10/DeleteItemOutput
type DeleteItemOutput struct {
	_ struct{} `type:"structure"`

	// A map of attribute names to AttributeValue objects, representing the item
	// as it appeared before the DeleteItem operation. This map appears in the response
	// only if ReturnValues was specified as ALL_OLD in the request.
	Attributes map[string]AttributeValue `type:"map"`

	// The capacity units consumed by the DeleteItem operation. The data returned
	// includes the total provisioned throughput consumed, along with statistics
	// for the table and any indexes involved in the operation. ConsumedCapacity
	// is only returned if the ReturnConsumedCapacity parameter was specified. For
	// more information, see Provisioned Throughput (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/ProvisionedThroughputIntro.html)
	// in the Amazon DynamoDB Developer Guide.
	ConsumedCapacity *ConsumedCapacity `type:"structure"`

	// Information about item collections, if any, that were affected by the DeleteItem
	// operation. ItemCollectionMetrics is only returned if the ReturnItemCollectionMetrics
	// parameter was specified. If the table does not have any local secondary indexes,
	// this information is not returned in the response.
	//
	// Each ItemCollectionMetrics element consists of:
	//
	//    * ItemCollectionKey - The partition key value of the item collection.
	//    This is the same as the partition key value of the item itself.
	//
	//    * SizeEstimateRangeGB - An estimate of item collection size, in gigabytes.
	//    This value is a two-element array containing a lower bound and an upper
	//    bound for the estimate. The estimate includes the size of all the items
	//    in the table, plus the size of all attributes projected into all of the
	//    local secondary indexes on that table. Use this estimate to measure whether
	//    a local secondary index is approaching its size limit.
	//
	// The estimate is subject to change over time; therefore, do not rely on the
	//    precision or accuracy of the estimate.
	ItemCollectionMetrics *ItemCollectionMetrics `type:"structure"`
}

// String returns the string representation
func (s DeleteItemOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteItem = "DeleteItem"

// DeleteItemRequest returns a request value for making API operation for
// Amazon DynamoDB.
//
// Deletes a single item in a table by primary key. You can perform a conditional
// delete operation that deletes the item if it exists, or if it has an expected
// attribute value.
//
// In addition to deleting an item, you can also return the item's attribute
// values in the same operation, using the ReturnValues parameter.
//
// Unless you specify conditions, the DeleteItem is an idempotent operation;
// running it multiple times on the same item or attribute does not result in
// an error response.
//
// Conditional deletes are useful for deleting items only if specific conditions
// are met. If those conditions are met, DynamoDB performs the delete. Otherwise,
// the item is not deleted.
//
//    // Example sending a request using DeleteItemRequest.
//    req := client.DeleteItemRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/dynamodb-2012-08-10/DeleteItem
func (c *Client) DeleteItemRequest(input *DeleteItemInput) DeleteItemRequest {
	op := &aws.Operation{
		Name:       opDeleteItem,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteItemInput{}
	}

	req := c.newRequest(op, input, &DeleteItemOutput{})
	return DeleteItemRequest{Request: req, Input: input, Copy: c.DeleteItemRequest}
}

// DeleteItemRequest is the request type for the
// DeleteItem API operation.
type DeleteItemRequest struct {
	*aws.Request
	Input *DeleteItemInput
	Copy  func(*DeleteItemInput) DeleteItemRequest
}

// Send marshals and sends the DeleteItem API request.
func (r DeleteItemRequest) Send(ctx context.Context) (*DeleteItemResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteItemResponse{
		DeleteItemOutput: r.Request.Data.(*DeleteItemOutput),
		response:         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteItemResponse is the response type for the
// DeleteItem API operation.
type DeleteItemResponse struct {
	*DeleteItemOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteItem request.
func (r *DeleteItemResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
