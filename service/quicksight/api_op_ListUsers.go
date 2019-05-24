// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package quicksight

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/quicksight-2018-04-01/ListUsersRequest
type ListUsersInput struct {
	_ struct{} `type:"structure"`

	// The ID for the AWS account that the user is in. Currently, you use the ID
	// for the AWS account that contains your Amazon QuickSight account.
	//
	// AwsAccountId is a required field
	AwsAccountId *string `location:"uri" locationName:"AwsAccountId" min:"12" type:"string" required:"true"`

	// The maximum number of results to return from this request.
	MaxResults *int64 `location:"querystring" locationName:"max-results" min:"1" type:"integer"`

	// The namespace. Currently, you should set this to default.
	//
	// Namespace is a required field
	Namespace *string `location:"uri" locationName:"Namespace" type:"string" required:"true"`

	// A pagination token that can be used in a subsequent request.
	NextToken *string `location:"querystring" locationName:"next-token" type:"string"`
}

// String returns the string representation
func (s ListUsersInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListUsersInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListUsersInput"}

	if s.AwsAccountId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AwsAccountId"))
	}
	if s.AwsAccountId != nil && len(*s.AwsAccountId) < 12 {
		invalidParams.Add(aws.NewErrParamMinLen("AwsAccountId", 12))
	}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if s.Namespace == nil {
		invalidParams.Add(aws.NewErrParamRequired("Namespace"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListUsersInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.0"), protocol.Metadata{})

	if s.AwsAccountId != nil {
		v := *s.AwsAccountId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "AwsAccountId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Namespace != nil {
		v := *s.Namespace

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "Namespace", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.MaxResults != nil {
		v := *s.MaxResults

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "max-results", protocol.Int64Value(v), metadata)
	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "next-token", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/quicksight-2018-04-01/ListUsersResponse
type ListUsersOutput struct {
	_ struct{} `type:"structure"`

	// A pagination token that can be used in a subsequent request.
	NextToken *string `type:"string"`

	// The AWS request ID for this operation.
	RequestId *string `type:"string"`

	// The http status of the request.
	Status *int64 `location:"statusCode" type:"integer"`

	// The list of users.
	UserList []User `type:"list"`
}

// String returns the string representation
func (s ListUsersOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListUsersOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "NextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.RequestId != nil {
		v := *s.RequestId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "RequestId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.UserList) > 0 {
		v := s.UserList

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "UserList", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	// ignoring invalid encode state, StatusCode. Status
	return nil
}

const opListUsers = "ListUsers"

// ListUsersRequest returns a request value for making API operation for
// Amazon QuickSight.
//
// Returns a list of all of the Amazon QuickSight users belonging to this account.
//
// The permission resource is arn:aws:quicksight:us-east-1:<aws-account-id>:user/default/*.
//
// The response is a list of user objects, containing each user's Amazon Resource
// Name (ARN), AWS Identity and Access Management (IAM) role, and email address.
//
// CLI Sample:
//
// aws quicksight list-users --aws-account-id=111122223333 --namespace=default
//
//    // Example sending a request using ListUsersRequest.
//    req := client.ListUsersRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/quicksight-2018-04-01/ListUsers
func (c *Client) ListUsersRequest(input *ListUsersInput) ListUsersRequest {
	op := &aws.Operation{
		Name:       opListUsers,
		HTTPMethod: "GET",
		HTTPPath:   "/accounts/{AwsAccountId}/namespaces/{Namespace}/users",
	}

	if input == nil {
		input = &ListUsersInput{}
	}

	req := c.newRequest(op, input, &ListUsersOutput{})
	return ListUsersRequest{Request: req, Input: input, Copy: c.ListUsersRequest}
}

// ListUsersRequest is the request type for the
// ListUsers API operation.
type ListUsersRequest struct {
	*aws.Request
	Input *ListUsersInput
	Copy  func(*ListUsersInput) ListUsersRequest
}

// Send marshals and sends the ListUsers API request.
func (r ListUsersRequest) Send(ctx context.Context) (*ListUsersResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListUsersResponse{
		ListUsersOutput: r.Request.Data.(*ListUsersOutput),
		response:        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListUsersResponse is the response type for the
// ListUsers API operation.
type ListUsersResponse struct {
	*ListUsersOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListUsers request.
func (r *ListUsersResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
