// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package lightsail

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/lightsail-2016-11-28/UpdateRelationalDatabaseRequest
type UpdateRelationalDatabaseInput struct {
	_ struct{} `type:"structure"`

	// When true, applies changes immediately. When false, applies changes during
	// the preferred maintenance window. Some changes may cause an outage.
	//
	// Default: false
	ApplyImmediately *bool `locationName:"applyImmediately" type:"boolean"`

	// When true, disables automated backup retention for your database.
	//
	// Disabling backup retention deletes all automated database backups. Before
	// disabling this, you may want to create a snapshot of your database using
	// the create relational database snapshot operation.
	//
	// Updates are applied during the next maintenance window because this can result
	// in an outage.
	DisableBackupRetention *bool `locationName:"disableBackupRetention" type:"boolean"`

	// When true, enables automated backup retention for your database.
	//
	// Updates are applied during the next maintenance window because this can result
	// in an outage.
	EnableBackupRetention *bool `locationName:"enableBackupRetention" type:"boolean"`

	// The password for the master user of your database. The password can include
	// any printable ASCII character except "/", """, or "@".
	//
	// Constraints: Must contain 8 to 41 characters.
	MasterUserPassword *string `locationName:"masterUserPassword" type:"string"`

	// The daily time range during which automated backups are created for your
	// database if automated backups are enabled.
	//
	// Constraints:
	//
	//    * Must be in the hh24:mi-hh24:mi format.
	//
	// Example: 16:00-16:30
	//
	//    * Specified in Universal Coordinated Time (UTC).
	//
	//    * Must not conflict with the preferred maintenance window.
	//
	//    * Must be at least 30 minutes.
	PreferredBackupWindow *string `locationName:"preferredBackupWindow" type:"string"`

	// The weekly time range during which system maintenance can occur on your database.
	//
	// The default is a 30-minute window selected at random from an 8-hour block
	// of time for each AWS Region, occurring on a random day of the week.
	//
	// Constraints:
	//
	//    * Must be in the ddd:hh24:mi-ddd:hh24:mi format.
	//
	//    * Valid days: Mon, Tue, Wed, Thu, Fri, Sat, Sun.
	//
	//    * Must be at least 30 minutes.
	//
	//    * Specified in Universal Coordinated Time (UTC).
	//
	//    * Example: Tue:17:00-Tue:17:30
	PreferredMaintenanceWindow *string `locationName:"preferredMaintenanceWindow" type:"string"`

	// Specifies the accessibility options for your database. A value of true specifies
	// a database that is available to resources outside of your Lightsail account.
	// A value of false specifies a database that is available only to your Lightsail
	// resources in the same region as your database.
	PubliclyAccessible *bool `locationName:"publiclyAccessible" type:"boolean"`

	// The name of your database to update.
	//
	// RelationalDatabaseName is a required field
	RelationalDatabaseName *string `locationName:"relationalDatabaseName" type:"string" required:"true"`

	// When true, the master user password is changed to a new strong password generated
	// by Lightsail.
	//
	// Use the get relational database master user password operation to get the
	// new password.
	RotateMasterUserPassword *bool `locationName:"rotateMasterUserPassword" type:"boolean"`
}

// String returns the string representation
func (s UpdateRelationalDatabaseInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateRelationalDatabaseInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateRelationalDatabaseInput"}

	if s.RelationalDatabaseName == nil {
		invalidParams.Add(aws.NewErrParamRequired("RelationalDatabaseName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/lightsail-2016-11-28/UpdateRelationalDatabaseResult
type UpdateRelationalDatabaseOutput struct {
	_ struct{} `type:"structure"`

	// An object describing the result of your update relational database request.
	Operations []Operation `locationName:"operations" type:"list"`
}

// String returns the string representation
func (s UpdateRelationalDatabaseOutput) String() string {
	return awsutil.Prettify(s)
}

const opUpdateRelationalDatabase = "UpdateRelationalDatabase"

// UpdateRelationalDatabaseRequest returns a request value for making API operation for
// Amazon Lightsail.
//
// Allows the update of one or more attributes of a database in Amazon Lightsail.
//
// Updates are applied immediately, or in cases where the updates could result
// in an outage, are applied during the database's predefined maintenance window.
//
// The update relational database operation supports tag-based access control
// via resource tags applied to the resource identified by relationalDatabaseName.
// For more information, see the Lightsail Dev Guide (https://lightsail.aws.amazon.com/ls/docs/en/articles/amazon-lightsail-controlling-access-using-tags).
//
//    // Example sending a request using UpdateRelationalDatabaseRequest.
//    req := client.UpdateRelationalDatabaseRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/lightsail-2016-11-28/UpdateRelationalDatabase
func (c *Client) UpdateRelationalDatabaseRequest(input *UpdateRelationalDatabaseInput) UpdateRelationalDatabaseRequest {
	op := &aws.Operation{
		Name:       opUpdateRelationalDatabase,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateRelationalDatabaseInput{}
	}

	req := c.newRequest(op, input, &UpdateRelationalDatabaseOutput{})
	return UpdateRelationalDatabaseRequest{Request: req, Input: input, Copy: c.UpdateRelationalDatabaseRequest}
}

// UpdateRelationalDatabaseRequest is the request type for the
// UpdateRelationalDatabase API operation.
type UpdateRelationalDatabaseRequest struct {
	*aws.Request
	Input *UpdateRelationalDatabaseInput
	Copy  func(*UpdateRelationalDatabaseInput) UpdateRelationalDatabaseRequest
}

// Send marshals and sends the UpdateRelationalDatabase API request.
func (r UpdateRelationalDatabaseRequest) Send(ctx context.Context) (*UpdateRelationalDatabaseResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateRelationalDatabaseResponse{
		UpdateRelationalDatabaseOutput: r.Request.Data.(*UpdateRelationalDatabaseOutput),
		response:                       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateRelationalDatabaseResponse is the response type for the
// UpdateRelationalDatabase API operation.
type UpdateRelationalDatabaseResponse struct {
	*UpdateRelationalDatabaseOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateRelationalDatabase request.
func (r *UpdateRelationalDatabaseResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
