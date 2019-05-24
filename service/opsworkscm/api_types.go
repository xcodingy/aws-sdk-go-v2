// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package opsworkscm

import (
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

var _ aws.Config
var _ = awsutil.Prettify

// Stores account attributes.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/opsworkscm-2016-11-01/AccountAttribute
type AccountAttribute struct {
	_ struct{} `type:"structure"`

	// The maximum allowed value.
	Maximum *int64 `type:"integer"`

	// The attribute name. The following are supported attribute names.
	//
	//    * ServerLimit: The number of current servers/maximum number of servers
	//    allowed. By default, you can have a maximum of 10 servers.
	//
	//    * ManualBackupLimit: The number of current manual backups/maximum number
	//    of backups allowed. By default, you can have a maximum of 50 manual backups
	//    saved.
	Name *string `type:"string"`

	// The current usage, such as the current number of servers that are associated
	// with the account.
	Used *int64 `type:"integer"`
}

// String returns the string representation
func (s AccountAttribute) String() string {
	return awsutil.Prettify(s)
}

// Describes a single backup.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/opsworkscm-2016-11-01/Backup
type Backup struct {
	_ struct{} `type:"structure"`

	// The ARN of the backup.
	BackupArn *string `type:"string"`

	// The generated ID of the backup. Example: myServerName-yyyyMMddHHmmssSSS
	BackupId *string `type:"string"`

	// The backup type. Valid values are automated or manual.
	BackupType BackupType `type:"string" enum:"true"`

	// The time stamp when the backup was created in the database. Example: 2016-07-29T13:38:47.520Z
	CreatedAt *time.Time `type:"timestamp" timestampFormat:"unix"`

	// A user-provided description for a manual backup. This field is empty for
	// automated backups.
	Description *string `type:"string"`

	// The engine type that is obtained from the server when the backup is created.
	Engine *string `type:"string"`

	// The engine model that is obtained from the server when the backup is created.
	EngineModel *string `type:"string"`

	// The engine version that is obtained from the server when the backup is created.
	EngineVersion *string `type:"string"`

	// The EC2 instance profile ARN that is obtained from the server when the backup
	// is created. Because this value is stored, you are not required to provide
	// the InstanceProfileArn again if you restore a backup.
	InstanceProfileArn *string `type:"string"`

	// The instance type that is obtained from the server when the backup is created.
	InstanceType *string `type:"string"`

	// The key pair that is obtained from the server when the backup is created.
	KeyPair *string `type:"string"`

	// The preferred backup period that is obtained from the server when the backup
	// is created.
	PreferredBackupWindow *string `type:"string"`

	// The preferred maintenance period that is obtained from the server when the
	// backup is created.
	PreferredMaintenanceWindow *string `type:"string"`

	// This field is deprecated and is no longer used.
	S3DataSize *int64 `deprecated:"true" type:"integer"`

	// This field is deprecated and is no longer used.
	S3DataUrl *string `deprecated:"true" type:"string"`

	// The Amazon S3 URL of the backup's log file.
	S3LogUrl *string `type:"string"`

	// The security group IDs that are obtained from the server when the backup
	// is created.
	SecurityGroupIds []string `type:"list"`

	// The name of the server from which the backup was made.
	ServerName *string `min:"1" type:"string"`

	// The service role ARN that is obtained from the server when the backup is
	// created.
	ServiceRoleArn *string `type:"string"`

	// The status of a backup while in progress.
	Status BackupStatus `type:"string" enum:"true"`

	// An informational message about backup status.
	StatusDescription *string `type:"string"`

	// The subnet IDs that are obtained from the server when the backup is created.
	SubnetIds []string `type:"list"`

	// The version of AWS OpsWorks CM-specific tools that is obtained from the server
	// when the backup is created.
	ToolsVersion *string `type:"string"`

	// The IAM user ARN of the requester for manual backups. This field is empty
	// for automated backups.
	UserArn *string `type:"string"`
}

// String returns the string representation
func (s Backup) String() string {
	return awsutil.Prettify(s)
}

// A name and value pair that is specific to the engine of the server.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/opsworkscm-2016-11-01/EngineAttribute
type EngineAttribute struct {
	_ struct{} `type:"structure"`

	// The name of the engine attribute.
	Name *string `type:"string"`

	// The value of the engine attribute.
	Value *string `type:"string"`
}

// String returns the string representation
func (s EngineAttribute) String() string {
	return awsutil.Prettify(s)
}

// Describes a configuration management server.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/opsworkscm-2016-11-01/Server
type Server struct {
	_ struct{} `type:"structure"`

	// Associate a public IP address with a server that you are launching.
	AssociatePublicIpAddress *bool `type:"boolean"`

	// The number of automated backups to keep.
	BackupRetentionCount *int64 `type:"integer"`

	// The ARN of the CloudFormation stack that was used to create the server.
	CloudFormationStackArn *string `type:"string"`

	// Time stamp of server creation. Example 2016-07-29T13:38:47.520Z
	CreatedAt *time.Time `type:"timestamp" timestampFormat:"unix"`

	// Disables automated backups. The number of stored backups is dependent on
	// the value of PreferredBackupCount.
	DisableAutomatedBackup *bool `type:"boolean"`

	// A DNS name that can be used to access the engine. Example: myserver-asdfghjkl.us-east-1.opsworks.io
	Endpoint *string `type:"string"`

	// The engine type of the server. Valid values in this release include Chef
	// and Puppet.
	Engine *string `type:"string"`

	// The response of a createServer() request returns the master credential to
	// access the server in EngineAttributes. These credentials are not stored by
	// AWS OpsWorks CM; they are returned only as part of the result of createServer().
	//
	// Attributes returned in a createServer response for Chef
	//
	//    * CHEF_PIVOTAL_KEY: A base64-encoded RSA private key that is generated
	//    by AWS OpsWorks for Chef Automate. This private key is required to access
	//    the Chef API.
	//
	//    * CHEF_STARTER_KIT: A base64-encoded ZIP file. The ZIP file contains a
	//    Chef starter kit, which includes a README, a configuration file, and the
	//    required RSA private key. Save this file, unzip it, and then change to
	//    the directory where you've unzipped the file contents. From this directory,
	//    you can run Knife commands.
	//
	// Attributes returned in a createServer response for Puppet
	//
	//    * PUPPET_STARTER_KIT: A base64-encoded ZIP file. The ZIP file contains
	//    a Puppet starter kit, including a README and a required private key. Save
	//    this file, unzip it, and then change to the directory where you've unzipped
	//    the file contents.
	//
	//    * PUPPET_ADMIN_PASSWORD: An administrator password that you can use to
	//    sign in to the Puppet Enterprise console after the server is online.
	EngineAttributes []EngineAttribute `type:"list"`

	// The engine model of the server. Valid values in this release include Monolithic
	// for Puppet and Single for Chef.
	EngineModel *string `type:"string"`

	// The engine version of the server. For a Chef server, the valid value for
	// EngineVersion is currently 12. For a Puppet server, the valid value is 2017.
	EngineVersion *string `type:"string"`

	// The instance profile ARN of the server.
	InstanceProfileArn *string `type:"string"`

	// The instance type for the server, as specified in the CloudFormation stack.
	// This might not be the same instance type that is shown in the EC2 console.
	InstanceType *string `type:"string"`

	// The key pair associated with the server.
	KeyPair *string `type:"string"`

	// The status of the most recent server maintenance run. Shows SUCCESS or FAILED.
	MaintenanceStatus MaintenanceStatus `type:"string" enum:"true"`

	// The preferred backup period specified for the server.
	PreferredBackupWindow *string `type:"string"`

	// The preferred maintenance period specified for the server.
	PreferredMaintenanceWindow *string `type:"string"`

	// The security group IDs for the server, as specified in the CloudFormation
	// stack. These might not be the same security groups that are shown in the
	// EC2 console.
	SecurityGroupIds []string `type:"list"`

	// The ARN of the server.
	ServerArn *string `type:"string"`

	// The name of the server.
	ServerName *string `type:"string"`

	// The service role ARN used to create the server.
	ServiceRoleArn *string `type:"string"`

	// The server's status. This field displays the states of actions in progress,
	// such as creating, running, or backing up the server, as well as the server's
	// health state.
	Status ServerStatus `type:"string" enum:"true"`

	// Depending on the server status, this field has either a human-readable message
	// (such as a create or backup error), or an escaped block of JSON (used for
	// health check results).
	StatusReason *string `type:"string"`

	// The subnet IDs specified in a CreateServer request.
	SubnetIds []string `type:"list"`
}

// String returns the string representation
func (s Server) String() string {
	return awsutil.Prettify(s)
}

// An event that is related to the server, such as the start of maintenance
// or backup.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/opsworkscm-2016-11-01/ServerEvent
type ServerEvent struct {
	_ struct{} `type:"structure"`

	// The time when the event occurred.
	CreatedAt *time.Time `type:"timestamp" timestampFormat:"unix"`

	// The Amazon S3 URL of the event's log file.
	LogUrl *string `type:"string"`

	// A human-readable informational or status message.
	Message *string `type:"string"`

	// The name of the server on or for which the event occurred.
	ServerName *string `type:"string"`
}

// String returns the string representation
func (s ServerEvent) String() string {
	return awsutil.Prettify(s)
}
