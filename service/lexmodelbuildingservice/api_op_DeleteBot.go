// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package lexmodelbuildingservice

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/lex-models-2017-04-19/DeleteBotRequest
type DeleteBotInput struct {
	_ struct{} `type:"structure"`

	// The name of the bot. The name is case sensitive.
	//
	// Name is a required field
	Name *string `location:"uri" locationName:"name" min:"2" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteBotInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteBotInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteBotInput"}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}
	if s.Name != nil && len(*s.Name) < 2 {
		invalidParams.Add(aws.NewErrParamMinLen("Name", 2))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteBotInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/lex-models-2017-04-19/DeleteBotOutput
type DeleteBotOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteBotOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteBotOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opDeleteBot = "DeleteBot"

// DeleteBotRequest returns a request value for making API operation for
// Amazon Lex Model Building Service.
//
// Deletes all versions of the bot, including the $LATEST version. To delete
// a specific version of the bot, use the DeleteBotVersion operation.
//
// If a bot has an alias, you can't delete it. Instead, the DeleteBot operation
// returns a ResourceInUseException exception that includes a reference to the
// alias that refers to the bot. To remove the reference to the bot, delete
// the alias. If you get the same exception again, delete the referring alias
// until the DeleteBot operation is successful.
//
// This operation requires permissions for the lex:DeleteBot action.
//
//    // Example sending a request using DeleteBotRequest.
//    req := client.DeleteBotRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/lex-models-2017-04-19/DeleteBot
func (c *Client) DeleteBotRequest(input *DeleteBotInput) DeleteBotRequest {
	op := &aws.Operation{
		Name:       opDeleteBot,
		HTTPMethod: "DELETE",
		HTTPPath:   "/bots/{name}",
	}

	if input == nil {
		input = &DeleteBotInput{}
	}

	req := c.newRequest(op, input, &DeleteBotOutput{})
	req.Handlers.Unmarshal.Remove(restjson.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return DeleteBotRequest{Request: req, Input: input, Copy: c.DeleteBotRequest}
}

// DeleteBotRequest is the request type for the
// DeleteBot API operation.
type DeleteBotRequest struct {
	*aws.Request
	Input *DeleteBotInput
	Copy  func(*DeleteBotInput) DeleteBotRequest
}

// Send marshals and sends the DeleteBot API request.
func (r DeleteBotRequest) Send(ctx context.Context) (*DeleteBotResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteBotResponse{
		DeleteBotOutput: r.Request.Data.(*DeleteBotOutput),
		response:        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteBotResponse is the response type for the
// DeleteBot API operation.
type DeleteBotResponse struct {
	*DeleteBotOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteBot request.
func (r *DeleteBotResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
