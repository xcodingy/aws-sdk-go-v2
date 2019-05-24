// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sts

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/sts-2011-06-15/GetSessionTokenRequest
type GetSessionTokenInput struct {
	_ struct{} `type:"structure"`

	// The duration, in seconds, that the credentials should remain valid. Acceptable
	// durations for IAM user sessions range from 900 seconds (15 minutes) to 129600
	// seconds (36 hours), with 43200 seconds (12 hours) as the default. Sessions
	// for AWS account owners are restricted to a maximum of 3600 seconds (one hour).
	// If the duration is longer than one hour, the session for AWS account owners
	// defaults to one hour.
	DurationSeconds *int64 `min:"900" type:"integer"`

	// The identification number of the MFA device that is associated with the IAM
	// user who is making the GetSessionToken call. Specify this value if the IAM
	// user has a policy that requires MFA authentication. The value is either the
	// serial number for a hardware device (such as GAHT12345678) or an Amazon Resource
	// Name (ARN) for a virtual device (such as arn:aws:iam::123456789012:mfa/user).
	// You can find the device for an IAM user by going to the AWS Management Console
	// and viewing the user's security credentials.
	//
	// The regex used to validated this parameter is a string of characters consisting
	// of upper- and lower-case alphanumeric characters with no spaces. You can
	// also include underscores or any of the following characters: =,.@:/-
	SerialNumber *string `min:"9" type:"string"`

	// The value provided by the MFA device, if MFA is required. If any policy requires
	// the IAM user to submit an MFA code, specify this value. If MFA authentication
	// is required, and the user does not provide a code when requesting a set of
	// temporary security credentials, the user will receive an "access denied"
	// response when requesting resources that require MFA authentication.
	//
	// The format for this parameter, as described by its regex pattern, is a sequence
	// of six numeric digits.
	TokenCode *string `min:"6" type:"string"`
}

// String returns the string representation
func (s GetSessionTokenInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetSessionTokenInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetSessionTokenInput"}
	if s.DurationSeconds != nil && *s.DurationSeconds < 900 {
		invalidParams.Add(aws.NewErrParamMinValue("DurationSeconds", 900))
	}
	if s.SerialNumber != nil && len(*s.SerialNumber) < 9 {
		invalidParams.Add(aws.NewErrParamMinLen("SerialNumber", 9))
	}
	if s.TokenCode != nil && len(*s.TokenCode) < 6 {
		invalidParams.Add(aws.NewErrParamMinLen("TokenCode", 6))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Contains the response to a successful GetSessionToken request, including
// temporary AWS credentials that can be used to make AWS requests.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/sts-2011-06-15/GetSessionTokenResponse
type GetSessionTokenOutput struct {
	_ struct{} `type:"structure"`

	// The temporary security credentials, which include an access key ID, a secret
	// access key, and a security (or session) token.
	//
	// Note: The size of the security token that STS APIs return is not fixed. We
	// strongly recommend that you make no assumptions about the maximum size. As
	// of this writing, the typical size is less than 4096 bytes, but that can vary.
	// Also, future updates to AWS might require larger sizes.
	Credentials *Credentials `type:"structure"`
}

// String returns the string representation
func (s GetSessionTokenOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetSessionToken = "GetSessionToken"

// GetSessionTokenRequest returns a request value for making API operation for
// AWS Security Token Service.
//
// Returns a set of temporary credentials for an AWS account or IAM user. The
// credentials consist of an access key ID, a secret access key, and a security
// token. Typically, you use GetSessionToken if you want to use MFA to protect
// programmatic calls to specific AWS APIs like Amazon EC2 StopInstances. MFA-enabled
// IAM users would need to call GetSessionToken and submit an MFA code that
// is associated with their MFA device. Using the temporary security credentials
// that are returned from the call, IAM users can then make programmatic calls
// to APIs that require MFA authentication. If you do not supply a correct MFA
// code, then the API returns an access denied error. For a comparison of GetSessionToken
// with the other APIs that produce temporary credentials, see Requesting Temporary
// Security Credentials (http://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_temp_request.html)
// and Comparing the AWS STS APIs (http://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_temp_request.html#stsapi_comparison)
// in the IAM User Guide.
//
// The GetSessionToken action must be called by using the long-term AWS security
// credentials of the AWS account or an IAM user. Credentials that are created
// by IAM users are valid for the duration that you specify, from 900 seconds
// (15 minutes) up to a maximum of 129600 seconds (36 hours), with a default
// of 43200 seconds (12 hours); credentials that are created by using account
// credentials can range from 900 seconds (15 minutes) up to a maximum of 3600
// seconds (1 hour), with a default of 1 hour.
//
// The temporary security credentials created by GetSessionToken can be used
// to make API calls to any AWS service with the following exceptions:
//
//    * You cannot call any IAM APIs unless MFA authentication information is
//    included in the request.
//
//    * You cannot call any STS API exceptAssumeRole or GetCallerIdentity.
//
// We recommend that you do not call GetSessionToken with root account credentials.
// Instead, follow our best practices (http://docs.aws.amazon.com/IAM/latest/UserGuide/best-practices.html#create-iam-users)
// by creating one or more IAM users, giving them the necessary permissions,
// and using IAM users for everyday interaction with AWS.
//
// The permissions associated with the temporary security credentials returned
// by GetSessionToken are based on the permissions associated with account or
// IAM user whose credentials are used to call the action. If GetSessionToken
// is called using root account credentials, the temporary credentials have
// root account permissions. Similarly, if GetSessionToken is called using the
// credentials of an IAM user, the temporary credentials have the same permissions
// as the IAM user.
//
// For more information about using GetSessionToken to create temporary credentials,
// go to Temporary Credentials for Users in Untrusted Environments (http://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_temp_request.html#api_getsessiontoken)
// in the IAM User Guide.
//
//    // Example sending a request using GetSessionTokenRequest.
//    req := client.GetSessionTokenRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/sts-2011-06-15/GetSessionToken
func (c *Client) GetSessionTokenRequest(input *GetSessionTokenInput) GetSessionTokenRequest {
	op := &aws.Operation{
		Name:       opGetSessionToken,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetSessionTokenInput{}
	}

	req := c.newRequest(op, input, &GetSessionTokenOutput{})
	return GetSessionTokenRequest{Request: req, Input: input, Copy: c.GetSessionTokenRequest}
}

// GetSessionTokenRequest is the request type for the
// GetSessionToken API operation.
type GetSessionTokenRequest struct {
	*aws.Request
	Input *GetSessionTokenInput
	Copy  func(*GetSessionTokenInput) GetSessionTokenRequest
}

// Send marshals and sends the GetSessionToken API request.
func (r GetSessionTokenRequest) Send(ctx context.Context) (*GetSessionTokenResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetSessionTokenResponse{
		GetSessionTokenOutput: r.Request.Data.(*GetSessionTokenOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetSessionTokenResponse is the response type for the
// GetSessionToken API operation.
type GetSessionTokenResponse struct {
	*GetSessionTokenOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetSessionToken request.
func (r *GetSessionTokenResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
