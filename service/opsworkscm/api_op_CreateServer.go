// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package opsworkscm

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/opsworkscm-2016-11-01/CreateServerRequest
type CreateServerInput struct {
	_ struct{} `type:"structure"`

	// Associate a public IP address with a server that you are launching. Valid
	// values are true or false. The default value is true.
	AssociatePublicIpAddress *bool `type:"boolean"`

	// If you specify this field, AWS OpsWorks CM creates the server by using the
	// backup represented by BackupId.
	BackupId *string `type:"string"`

	// The number of automated backups that you want to keep. Whenever a new backup
	// is created, AWS OpsWorks CM deletes the oldest backups if this number is
	// exceeded. The default value is 1.
	BackupRetentionCount *int64 `min:"1" type:"integer"`

	// Enable or disable scheduled backups. Valid values are true or false. The
	// default value is true.
	DisableAutomatedBackup *bool `type:"boolean"`

	// The configuration management engine to use. Valid values include Chef and
	// Puppet.
	Engine *string `type:"string"`

	// Optional engine attributes on a specified server.
	//
	// Attributes accepted in a Chef createServer request:
	//
	//    * CHEF_PIVOTAL_KEY: A base64-encoded RSA public key. The corresponding
	//    private key is required to access the Chef API. When no CHEF_PIVOTAL_KEY
	//    is set, a private key is generated and returned in the response.
	//
	//    * CHEF_DELIVERY_ADMIN_PASSWORD: The password for the administrative user
	//    in the Chef Automate GUI. The password length is a minimum of eight characters,
	//    and a maximum of 32. The password can contain letters, numbers, and special
	//    characters (!/@#$%^&+=_). The password must contain at least one lower
	//    case letter, one upper case letter, one number, and one special character.
	//    When no CHEF_DELIVERY_ADMIN_PASSWORD is set, one is generated and returned
	//    in the response.
	//
	// Attributes accepted in a Puppet createServer request:
	//
	//    * PUPPET_ADMIN_PASSWORD: To work with the Puppet Enterprise console, a
	//    password must use ASCII characters.
	//
	//    * PUPPET_R10K_REMOTE: The r10k remote is the URL of your control repository
	//    (for example, ssh://git@your.git-repo.com:user/control-repo.git). Specifying
	//    an r10k remote opens TCP port 8170.
	//
	//    * PUPPET_R10K_PRIVATE_KEY: If you are using a private Git repository,
	//    add PUPPET_R10K_PRIVATE_KEY to specify an SSH URL and a PEM-encoded private
	//    SSH key.
	EngineAttributes []EngineAttribute `type:"list"`

	// The engine model of the server. Valid values in this release include Monolithic
	// for Puppet and Single for Chef.
	EngineModel *string `type:"string"`

	// The major release version of the engine that you want to use. For a Chef
	// server, the valid value for EngineVersion is currently 12. For a Puppet server,
	// the valid value is 2017.
	EngineVersion *string `type:"string"`

	// The ARN of the instance profile that your Amazon EC2 instances use. Although
	// the AWS OpsWorks console typically creates the instance profile for you,
	// if you are using API commands instead, run the service-role-creation.yaml
	// AWS CloudFormation template, located at https://s3.amazonaws.com/opsworks-cm-us-east-1-prod-default-assets/misc/opsworks-cm-roles.yaml.
	// This template creates a CloudFormation stack that includes the instance profile
	// you need.
	//
	// InstanceProfileArn is a required field
	InstanceProfileArn *string `type:"string" required:"true"`

	// The Amazon EC2 instance type to use. For example, m4.large. Recommended instance
	// types include t2.medium and greater, m4.*, or c4.xlarge and greater.
	//
	// InstanceType is a required field
	InstanceType *string `type:"string" required:"true"`

	// The Amazon EC2 key pair to set for the instance. This parameter is optional;
	// if desired, you may specify this parameter to connect to your instances by
	// using SSH.
	KeyPair *string `type:"string"`

	// The start time for a one-hour period during which AWS OpsWorks CM backs up
	// application-level data on your server if automated backups are enabled. Valid
	// values must be specified in one of the following formats:
	//
	//    * HH:MM for daily backups
	//
	//    * DDD:HH:MM for weekly backups
	//
	// The specified time is in coordinated universal time (UTC). The default value
	// is a random, daily start time.
	//
	// Example: 08:00, which represents a daily start time of 08:00 UTC.
	//
	// Example: Mon:08:00, which represents a start time of every Monday at 08:00
	// UTC. (8:00 a.m.)
	PreferredBackupWindow *string `type:"string"`

	// The start time for a one-hour period each week during which AWS OpsWorks
	// CM performs maintenance on the instance. Valid values must be specified in
	// the following format: DDD:HH:MM. The specified time is in coordinated universal
	// time (UTC). The default value is a random one-hour period on Tuesday, Wednesday,
	// or Friday. See TimeWindowDefinition for more information.
	//
	// Example: Mon:08:00, which represents a start time of every Monday at 08:00
	// UTC. (8:00 a.m.)
	PreferredMaintenanceWindow *string `type:"string"`

	// A list of security group IDs to attach to the Amazon EC2 instance. If you
	// add this parameter, the specified security groups must be within the VPC
	// that is specified by SubnetIds.
	//
	// If you do not specify this parameter, AWS OpsWorks CM creates one new security
	// group that uses TCP ports 22 and 443, open to 0.0.0.0/0 (everyone).
	SecurityGroupIds []string `type:"list"`

	// The name of the server. The server name must be unique within your AWS account,
	// within each region. Server names must start with a letter; then letters,
	// numbers, or hyphens (-) are allowed, up to a maximum of 40 characters.
	//
	// ServerName is a required field
	ServerName *string `min:"1" type:"string" required:"true"`

	// The service role that the AWS OpsWorks CM service backend uses to work with
	// your account. Although the AWS OpsWorks management console typically creates
	// the service role for you, if you are using the AWS CLI or API commands, run
	// the service-role-creation.yaml AWS CloudFormation template, located at https://s3.amazonaws.com/opsworks-cm-us-east-1-prod-default-assets/misc/opsworks-cm-roles.yaml.
	// This template creates a CloudFormation stack that includes the service role
	// and instance profile that you need.
	//
	// ServiceRoleArn is a required field
	ServiceRoleArn *string `type:"string" required:"true"`

	// The IDs of subnets in which to launch the server EC2 instance.
	//
	// Amazon EC2-Classic customers: This field is required. All servers must run
	// within a VPC. The VPC must have "Auto Assign Public IP" enabled.
	//
	// EC2-VPC customers: This field is optional. If you do not specify subnet IDs,
	// your EC2 instances are created in a default subnet that is selected by Amazon
	// EC2. If you specify subnet IDs, the VPC must have "Auto Assign Public IP"
	// enabled.
	//
	// For more information about supported Amazon EC2 platforms, see Supported
	// Platforms (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-supported-platforms.html).
	SubnetIds []string `type:"list"`
}

// String returns the string representation
func (s CreateServerInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateServerInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateServerInput"}
	if s.BackupRetentionCount != nil && *s.BackupRetentionCount < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("BackupRetentionCount", 1))
	}

	if s.InstanceProfileArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("InstanceProfileArn"))
	}

	if s.InstanceType == nil {
		invalidParams.Add(aws.NewErrParamRequired("InstanceType"))
	}

	if s.ServerName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ServerName"))
	}
	if s.ServerName != nil && len(*s.ServerName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ServerName", 1))
	}

	if s.ServiceRoleArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("ServiceRoleArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/opsworkscm-2016-11-01/CreateServerResponse
type CreateServerOutput struct {
	_ struct{} `type:"structure"`

	// The server that is created by the request.
	Server *Server `type:"structure"`
}

// String returns the string representation
func (s CreateServerOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateServer = "CreateServer"

// CreateServerRequest returns a request value for making API operation for
// AWS OpsWorks for Chef Automate.
//
// Creates and immedately starts a new server. The server is ready to use when
// it is in the HEALTHY state. By default, you can create a maximum of 10 servers.
//
// This operation is asynchronous.
//
// A LimitExceededException is thrown when you have created the maximum number
// of servers (10). A ResourceAlreadyExistsException is thrown when a server
// with the same name already exists in the account. A ResourceNotFoundException
// is thrown when you specify a backup ID that is not valid or is for a backup
// that does not exist. A ValidationException is thrown when parameters of the
// request are not valid.
//
// If you do not specify a security group by adding the SecurityGroupIds parameter,
// AWS OpsWorks creates a new security group.
//
// Chef Automate: The default security group opens the Chef server to the world
// on TCP port 443. If a KeyName is present, AWS OpsWorks enables SSH access.
// SSH is also open to the world on TCP port 22.
//
// Puppet Enterprise: The default security group opens TCP ports 22, 443, 4433,
// 8140, 8142, 8143, and 8170. If a KeyName is present, AWS OpsWorks enables
// SSH access. SSH is also open to the world on TCP port 22.
//
// By default, your server is accessible from any IP address. We recommend that
// you update your security group rules to allow access from known IP addresses
// and address ranges only. To edit security group rules, open Security Groups
// in the navigation pane of the EC2 management console.
//
//    // Example sending a request using CreateServerRequest.
//    req := client.CreateServerRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/opsworkscm-2016-11-01/CreateServer
func (c *Client) CreateServerRequest(input *CreateServerInput) CreateServerRequest {
	op := &aws.Operation{
		Name:       opCreateServer,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateServerInput{}
	}

	req := c.newRequest(op, input, &CreateServerOutput{})
	return CreateServerRequest{Request: req, Input: input, Copy: c.CreateServerRequest}
}

// CreateServerRequest is the request type for the
// CreateServer API operation.
type CreateServerRequest struct {
	*aws.Request
	Input *CreateServerInput
	Copy  func(*CreateServerInput) CreateServerRequest
}

// Send marshals and sends the CreateServer API request.
func (r CreateServerRequest) Send(ctx context.Context) (*CreateServerResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateServerResponse{
		CreateServerOutput: r.Request.Data.(*CreateServerOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateServerResponse is the response type for the
// CreateServer API operation.
type CreateServerResponse struct {
	*CreateServerOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateServer request.
func (r *CreateServerResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}