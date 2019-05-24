// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package gamelift

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents the input for a request action.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/gamelift-2015-10-01/CreateGameSessionInput
type CreateGameSessionInput struct {
	_ struct{} `type:"structure"`

	// Unique identifier for an alias associated with the fleet to create a game
	// session in. Each request must reference either a fleet ID or alias ID, but
	// not both.
	AliasId *string `type:"string"`

	// Unique identifier for a player or entity creating the game session. This
	// ID is used to enforce a resource protection policy (if one exists) that limits
	// the number of concurrent active game sessions one player can have.
	CreatorId *string `min:"1" type:"string"`

	// Unique identifier for a fleet to create a game session in. Each request must
	// reference either a fleet ID or alias ID, but not both.
	FleetId *string `type:"string"`

	// Set of custom properties for a game session, formatted as key:value pairs.
	// These properties are passed to a game server process in the GameSession object
	// with a request to start a new game session (see Start a Game Session (https://docs.aws.amazon.com/gamelift/latest/developerguide/gamelift-sdk-server-api.html#gamelift-sdk-server-startsession)).
	GameProperties []GameProperty `type:"list"`

	// Set of custom game session properties, formatted as a single string value.
	// This data is passed to a game server process in the GameSession object with
	// a request to start a new game session (see Start a Game Session (https://docs.aws.amazon.com/gamelift/latest/developerguide/gamelift-sdk-server-api.html#gamelift-sdk-server-startsession)).
	GameSessionData *string `min:"1" type:"string"`

	// This parameter is no longer preferred. Please use IdempotencyToken instead.
	// Custom string that uniquely identifies a request for a new game session.
	// Maximum token length is 48 characters. If provided, this string is included
	// in the new game session's ID. (A game session ARN has the following format:
	// arn:aws:gamelift:<region>::gamesession/<fleet ID>/<custom ID string or idempotency
	// token>.)
	GameSessionId *string `min:"1" type:"string"`

	// Custom string that uniquely identifies a request for a new game session.
	// Maximum token length is 48 characters. If provided, this string is included
	// in the new game session's ID. (A game session ARN has the following format:
	// arn:aws:gamelift:<region>::gamesession/<fleet ID>/<custom ID string or idempotency
	// token>.) Idempotency tokens remain in use for 30 days after a game session
	// has ended; game session objects are retained for this time period and then
	// deleted.
	IdempotencyToken *string `min:"1" type:"string"`

	// Maximum number of players that can be connected simultaneously to the game
	// session.
	//
	// MaximumPlayerSessionCount is a required field
	MaximumPlayerSessionCount *int64 `type:"integer" required:"true"`

	// Descriptive label that is associated with a game session. Session names do
	// not need to be unique.
	Name *string `min:"1" type:"string"`
}

// String returns the string representation
func (s CreateGameSessionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateGameSessionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateGameSessionInput"}
	if s.CreatorId != nil && len(*s.CreatorId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("CreatorId", 1))
	}
	if s.GameSessionData != nil && len(*s.GameSessionData) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("GameSessionData", 1))
	}
	if s.GameSessionId != nil && len(*s.GameSessionId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("GameSessionId", 1))
	}
	if s.IdempotencyToken != nil && len(*s.IdempotencyToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("IdempotencyToken", 1))
	}

	if s.MaximumPlayerSessionCount == nil {
		invalidParams.Add(aws.NewErrParamRequired("MaximumPlayerSessionCount"))
	}
	if s.Name != nil && len(*s.Name) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Name", 1))
	}
	if s.GameProperties != nil {
		for i, v := range s.GameProperties {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "GameProperties", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents the returned data in response to a request action.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/gamelift-2015-10-01/CreateGameSessionOutput
type CreateGameSessionOutput struct {
	_ struct{} `type:"structure"`

	// Object that describes the newly created game session record.
	GameSession *GameSession `type:"structure"`
}

// String returns the string representation
func (s CreateGameSessionOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateGameSession = "CreateGameSession"

// CreateGameSessionRequest returns a request value for making API operation for
// Amazon GameLift.
//
// Creates a multiplayer game session for players. This action creates a game
// session record and assigns an available server process in the specified fleet
// to host the game session. A fleet must have an ACTIVE status before a game
// session can be created in it.
//
// To create a game session, specify either fleet ID or alias ID and indicate
// a maximum number of players to allow in the game session. You can also provide
// a name and game-specific properties for this game session. If successful,
// a GameSession object is returned containing the game session properties and
// other settings you specified.
//
// Idempotency tokens. You can add a token that uniquely identifies game session
// requests. This is useful for ensuring that game session requests are idempotent.
// Multiple requests with the same idempotency token are processed only once;
// subsequent requests return the original result. All response values are the
// same with the exception of game session status, which may change.
//
// Resource creation limits. If you are creating a game session on a fleet with
// a resource creation limit policy in force, then you must specify a creator
// ID. Without this ID, Amazon GameLift has no way to evaluate the policy for
// this new game session request.
//
// Player acceptance policy. By default, newly created game sessions are open
// to new players. You can restrict new player access by using UpdateGameSession
// to change the game session's player session creation policy.
//
// Game session logs. Logs are retained for all active game sessions for 14
// days. To access the logs, call GetGameSessionLogUrl to download the log files.
//
// Available in Amazon GameLift Local.
//
//    * CreateGameSession
//
//    * DescribeGameSessions
//
//    * DescribeGameSessionDetails
//
//    * SearchGameSessions
//
//    * UpdateGameSession
//
//    * GetGameSessionLogUrl
//
//    * Game session placements
//
// StartGameSessionPlacement
//
// DescribeGameSessionPlacement
//
// StopGameSessionPlacement
//
//    // Example sending a request using CreateGameSessionRequest.
//    req := client.CreateGameSessionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/gamelift-2015-10-01/CreateGameSession
func (c *Client) CreateGameSessionRequest(input *CreateGameSessionInput) CreateGameSessionRequest {
	op := &aws.Operation{
		Name:       opCreateGameSession,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateGameSessionInput{}
	}

	req := c.newRequest(op, input, &CreateGameSessionOutput{})
	return CreateGameSessionRequest{Request: req, Input: input, Copy: c.CreateGameSessionRequest}
}

// CreateGameSessionRequest is the request type for the
// CreateGameSession API operation.
type CreateGameSessionRequest struct {
	*aws.Request
	Input *CreateGameSessionInput
	Copy  func(*CreateGameSessionInput) CreateGameSessionRequest
}

// Send marshals and sends the CreateGameSession API request.
func (r CreateGameSessionRequest) Send(ctx context.Context) (*CreateGameSessionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateGameSessionResponse{
		CreateGameSessionOutput: r.Request.Data.(*CreateGameSessionOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateGameSessionResponse is the response type for the
// CreateGameSession API operation.
type CreateGameSessionResponse struct {
	*CreateGameSessionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateGameSession request.
func (r *CreateGameSessionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
