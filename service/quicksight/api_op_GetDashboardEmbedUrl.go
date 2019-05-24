// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package quicksight

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/quicksight-2018-04-01/GetDashboardEmbedUrlRequest
type GetDashboardEmbedUrlInput struct {
	_ struct{} `type:"structure"`

	// AWS account ID that contains the dashboard you are embedding.
	//
	// AwsAccountId is a required field
	AwsAccountId *string `location:"uri" locationName:"AwsAccountId" min:"12" type:"string" required:"true"`

	// The ID for the dashboard, also added to IAM policy
	//
	// DashboardId is a required field
	DashboardId *string `location:"uri" locationName:"DashboardId" type:"string" required:"true"`

	// The authentication method the user uses to sign in (IAM only).
	//
	// IdentityType is a required field
	IdentityType IdentityType `location:"querystring" locationName:"creds-type" type:"string" required:"true" enum:"true"`

	// Remove the reset button on embedded dashboard. The default is FALSE, which
	// allows the reset button.
	ResetDisabled *bool `location:"querystring" locationName:"reset-disabled" type:"boolean"`

	// How many minutes the session is valid. The session lifetime must be between
	// 15 and 600 minutes.
	SessionLifetimeInMinutes *int64 `location:"querystring" locationName:"session-lifetime" min:"15" type:"long"`

	// Remove the undo/redo button on embedded dashboard. The default is FALSE,
	// which enables the undo/redo button.
	UndoRedoDisabled *bool `location:"querystring" locationName:"undo-redo-disabled" type:"boolean"`
}

// String returns the string representation
func (s GetDashboardEmbedUrlInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetDashboardEmbedUrlInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetDashboardEmbedUrlInput"}

	if s.AwsAccountId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AwsAccountId"))
	}
	if s.AwsAccountId != nil && len(*s.AwsAccountId) < 12 {
		invalidParams.Add(aws.NewErrParamMinLen("AwsAccountId", 12))
	}

	if s.DashboardId == nil {
		invalidParams.Add(aws.NewErrParamRequired("DashboardId"))
	}
	if len(s.IdentityType) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("IdentityType"))
	}
	if s.SessionLifetimeInMinutes != nil && *s.SessionLifetimeInMinutes < 15 {
		invalidParams.Add(aws.NewErrParamMinValue("SessionLifetimeInMinutes", 15))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetDashboardEmbedUrlInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.0"), protocol.Metadata{})

	if s.AwsAccountId != nil {
		v := *s.AwsAccountId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "AwsAccountId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.DashboardId != nil {
		v := *s.DashboardId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "DashboardId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.IdentityType) > 0 {
		v := s.IdentityType

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "creds-type", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.ResetDisabled != nil {
		v := *s.ResetDisabled

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "reset-disabled", protocol.BoolValue(v), metadata)
	}
	if s.SessionLifetimeInMinutes != nil {
		v := *s.SessionLifetimeInMinutes

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "session-lifetime", protocol.Int64Value(v), metadata)
	}
	if s.UndoRedoDisabled != nil {
		v := *s.UndoRedoDisabled

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "undo-redo-disabled", protocol.BoolValue(v), metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/quicksight-2018-04-01/GetDashboardEmbedUrlResponse
type GetDashboardEmbedUrlOutput struct {
	_ struct{} `type:"structure"`

	// URL that you can put into your server-side webpage to embed your dashboard.
	// This URL is valid for 5 minutes, and the resulting session is valid for 10
	// hours. The API provides the URL with an auth_code that enables a single-signon
	// session.
	EmbedUrl *string `type:"string"`

	// The AWS request ID for this operation.
	RequestId *string `type:"string"`

	// The http status of the request.
	Status *int64 `location:"statusCode" type:"integer"`
}

// String returns the string representation
func (s GetDashboardEmbedUrlOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetDashboardEmbedUrlOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.EmbedUrl != nil {
		v := *s.EmbedUrl

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "EmbedUrl", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.RequestId != nil {
		v := *s.RequestId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "RequestId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	// ignoring invalid encode state, StatusCode. Status
	return nil
}

const opGetDashboardEmbedUrl = "GetDashboardEmbedUrl"

// GetDashboardEmbedUrlRequest returns a request value for making API operation for
// Amazon QuickSight.
//
// Generates a server-side embeddable URL and authorization code. Before this
// can work properly, first you need to configure the dashboards and user permissions.
// For more information, see  Embedding Amazon QuickSight Dashboards (https://docs.aws.amazon.com/en_us/quicksight/latest/user/embedding.html).
//
// Currently, you can use GetDashboardEmbedURL only from the server, not from
// the user’s browser.
//
// CLI Sample:
//
// Assume the role with permissions enabled for actions: quickSight:RegisterUser
// and quicksight:GetDashboardEmbedURL. You can use assume-role, assume-role-with-web-identity,
// or assume-role-with-saml.
//
// aws sts assume-role --role-arn "arn:aws:iam::111122223333:role/embedding_quicksight_dashboard_role"
// --role-session-name embeddingsession
//
// If the user does not exist in QuickSight, register the user:
//
// aws quicksight register-user --aws-account-id 111122223333 --namespace default
// --identity-type IAM --iam-arn "arn:aws:iam::111122223333:role/embedding_quicksight_dashboard_role"
// --user-role READER --session-name "embeddingsession" --email user123@example.com
// --region us-east-1
//
// Get the URL for the embedded dashboard
//
// aws quicksight get-dashboard-embed-url --aws-account-id 111122223333 --dashboard-id
// 1a1ac2b2-3fc3-4b44-5e5d-c6db6778df89 --identity-type IAM
//
//    // Example sending a request using GetDashboardEmbedUrlRequest.
//    req := client.GetDashboardEmbedUrlRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/quicksight-2018-04-01/GetDashboardEmbedUrl
func (c *Client) GetDashboardEmbedUrlRequest(input *GetDashboardEmbedUrlInput) GetDashboardEmbedUrlRequest {
	op := &aws.Operation{
		Name:       opGetDashboardEmbedUrl,
		HTTPMethod: "GET",
		HTTPPath:   "/accounts/{AwsAccountId}/dashboards/{DashboardId}/embed-url",
	}

	if input == nil {
		input = &GetDashboardEmbedUrlInput{}
	}

	req := c.newRequest(op, input, &GetDashboardEmbedUrlOutput{})
	return GetDashboardEmbedUrlRequest{Request: req, Input: input, Copy: c.GetDashboardEmbedUrlRequest}
}

// GetDashboardEmbedUrlRequest is the request type for the
// GetDashboardEmbedUrl API operation.
type GetDashboardEmbedUrlRequest struct {
	*aws.Request
	Input *GetDashboardEmbedUrlInput
	Copy  func(*GetDashboardEmbedUrlInput) GetDashboardEmbedUrlRequest
}

// Send marshals and sends the GetDashboardEmbedUrl API request.
func (r GetDashboardEmbedUrlRequest) Send(ctx context.Context) (*GetDashboardEmbedUrlResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetDashboardEmbedUrlResponse{
		GetDashboardEmbedUrlOutput: r.Request.Data.(*GetDashboardEmbedUrlOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetDashboardEmbedUrlResponse is the response type for the
// GetDashboardEmbedUrl API operation.
type GetDashboardEmbedUrlResponse struct {
	*GetDashboardEmbedUrlOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetDashboardEmbedUrl request.
func (r *GetDashboardEmbedUrlResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
