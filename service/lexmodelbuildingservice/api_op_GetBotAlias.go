// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package lexmodelbuildingservice

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/lex-models-2017-04-19/GetBotAliasRequest
type GetBotAliasInput struct {
	_ struct{} `type:"structure"`

	// The name of the bot.
	//
	// BotName is a required field
	BotName *string `location:"uri" locationName:"botName" min:"2" type:"string" required:"true"`

	// The name of the bot alias. The name is case sensitive.
	//
	// Name is a required field
	Name *string `location:"uri" locationName:"name" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s GetBotAliasInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetBotAliasInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetBotAliasInput"}

	if s.BotName == nil {
		invalidParams.Add(aws.NewErrParamRequired("BotName"))
	}
	if s.BotName != nil && len(*s.BotName) < 2 {
		invalidParams.Add(aws.NewErrParamMinLen("BotName", 2))
	}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}
	if s.Name != nil && len(*s.Name) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Name", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetBotAliasInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	if s.BotName != nil {
		v := *s.BotName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "botName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/lex-models-2017-04-19/GetBotAliasResponse
type GetBotAliasOutput struct {
	_ struct{} `type:"structure"`

	// The name of the bot that the alias points to.
	BotName *string `locationName:"botName" min:"2" type:"string"`

	// The version of the bot that the alias points to.
	BotVersion *string `locationName:"botVersion" min:"1" type:"string"`

	// Checksum of the bot alias.
	Checksum *string `locationName:"checksum" type:"string"`

	// The date that the bot alias was created.
	CreatedDate *time.Time `locationName:"createdDate" type:"timestamp" timestampFormat:"unix"`

	// A description of the bot alias.
	Description *string `locationName:"description" type:"string"`

	// The date that the bot alias was updated. When you create a resource, the
	// creation date and the last updated date are the same.
	LastUpdatedDate *time.Time `locationName:"lastUpdatedDate" type:"timestamp" timestampFormat:"unix"`

	// The name of the bot alias.
	Name *string `locationName:"name" min:"1" type:"string"`
}

// String returns the string representation
func (s GetBotAliasOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetBotAliasOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.BotName != nil {
		v := *s.BotName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "botName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.BotVersion != nil {
		v := *s.BotVersion

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "botVersion", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Checksum != nil {
		v := *s.Checksum

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "checksum", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.CreatedDate != nil {
		v := *s.CreatedDate

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "createdDate", protocol.TimeValue{V: v, Format: protocol.UnixTimeFormat}, metadata)
	}
	if s.Description != nil {
		v := *s.Description

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "description", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.LastUpdatedDate != nil {
		v := *s.LastUpdatedDate

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "lastUpdatedDate", protocol.TimeValue{V: v, Format: protocol.UnixTimeFormat}, metadata)
	}
	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opGetBotAlias = "GetBotAlias"

// GetBotAliasRequest returns a request value for making API operation for
// Amazon Lex Model Building Service.
//
// Returns information about an Amazon Lex bot alias. For more information about
// aliases, see versioning-aliases.
//
// This operation requires permissions for the lex:GetBotAlias action.
//
//    // Example sending a request using GetBotAliasRequest.
//    req := client.GetBotAliasRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/lex-models-2017-04-19/GetBotAlias
func (c *Client) GetBotAliasRequest(input *GetBotAliasInput) GetBotAliasRequest {
	op := &aws.Operation{
		Name:       opGetBotAlias,
		HTTPMethod: "GET",
		HTTPPath:   "/bots/{botName}/aliases/{name}",
	}

	if input == nil {
		input = &GetBotAliasInput{}
	}

	req := c.newRequest(op, input, &GetBotAliasOutput{})
	return GetBotAliasRequest{Request: req, Input: input, Copy: c.GetBotAliasRequest}
}

// GetBotAliasRequest is the request type for the
// GetBotAlias API operation.
type GetBotAliasRequest struct {
	*aws.Request
	Input *GetBotAliasInput
	Copy  func(*GetBotAliasInput) GetBotAliasRequest
}

// Send marshals and sends the GetBotAlias API request.
func (r GetBotAliasRequest) Send(ctx context.Context) (*GetBotAliasResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetBotAliasResponse{
		GetBotAliasOutput: r.Request.Data.(*GetBotAliasOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetBotAliasResponse is the response type for the
// GetBotAlias API operation.
type GetBotAliasResponse struct {
	*GetBotAliasOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetBotAlias request.
func (r *GetBotAliasResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
