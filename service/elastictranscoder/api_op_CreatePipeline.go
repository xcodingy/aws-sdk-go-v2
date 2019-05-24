// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package elastictranscoder

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// The CreatePipelineRequest structure.
type CreatePipelineInput struct {
	_ struct{} `type:"structure"`

	// The AWS Key Management Service (AWS KMS) key that you want to use with this
	// pipeline.
	//
	// If you use either s3 or s3-aws-kms as your Encryption:Mode, you don't need
	// to provide a key with your job because a default key, known as an AWS-KMS
	// key, is created for you automatically. You need to provide an AWS-KMS key
	// only if you want to use a non-default AWS-KMS key, or if you are using an
	// Encryption:Mode of aes-cbc-pkcs7, aes-ctr, or aes-gcm.
	AwsKmsKeyArn *string `type:"string"`

	// The optional ContentConfig object specifies information about the Amazon
	// S3 bucket in which you want Elastic Transcoder to save transcoded files and
	// playlists: which bucket to use, which users you want to have access to the
	// files, the type of access you want users to have, and the storage class that
	// you want to assign to the files.
	//
	// If you specify values for ContentConfig, you must also specify values for
	// ThumbnailConfig.
	//
	// If you specify values for ContentConfig and ThumbnailConfig, omit the OutputBucket
	// object.
	//
	//    * Bucket: The Amazon S3 bucket in which you want Elastic Transcoder to
	//    save transcoded files and playlists.
	//
	//    * Permissions (Optional): The Permissions object specifies which users
	//    you want to have access to transcoded files and the type of access you
	//    want them to have. You can grant permissions to a maximum of 30 users
	//    and/or predefined Amazon S3 groups.
	//
	//    * Grantee Type: Specify the type of value that appears in the Grantee
	//    object:
	//
	// Canonical: The value in the Grantee object is either the canonical user ID
	//    for an AWS account or an origin access identity for an Amazon CloudFront
	//    distribution. For more information about canonical user IDs, see Access
	//    Control List (ACL) Overview in the Amazon Simple Storage Service Developer
	//    Guide. For more information about using CloudFront origin access identities
	//    to require that users use CloudFront URLs instead of Amazon S3 URLs, see
	//    Using an Origin Access Identity to Restrict Access to Your Amazon S3 Content.
	//
	// A canonical user ID is not the same as an AWS account number.
	//
	// Email: The value in the Grantee object is the registered email address of
	//    an AWS account.
	//
	// Group: The value in the Grantee object is one of the following predefined
	//    Amazon S3 groups: AllUsers, AuthenticatedUsers, or LogDelivery.
	//
	//    * Grantee: The AWS user or group that you want to have access to transcoded
	//    files and playlists. To identify the user or group, you can specify the
	//    canonical user ID for an AWS account, an origin access identity for a
	//    CloudFront distribution, the registered email address of an AWS account,
	//    or a predefined Amazon S3 group
	//
	//    * Access: The permission that you want to give to the AWS user that you
	//    specified in Grantee. Permissions are granted on the files that Elastic
	//    Transcoder adds to the bucket, including playlists and video files. Valid
	//    values include:
	//
	// READ: The grantee can read the objects and metadata for objects that Elastic
	//    Transcoder adds to the Amazon S3 bucket.
	//
	// READ_ACP: The grantee can read the object ACL for objects that Elastic Transcoder
	//    adds to the Amazon S3 bucket.
	//
	// WRITE_ACP: The grantee can write the ACL for the objects that Elastic Transcoder
	//    adds to the Amazon S3 bucket.
	//
	// FULL_CONTROL: The grantee has READ, READ_ACP, and WRITE_ACP permissions for
	//    the objects that Elastic Transcoder adds to the Amazon S3 bucket.
	//
	//    * StorageClass: The Amazon S3 storage class, Standard or ReducedRedundancy,
	//    that you want Elastic Transcoder to assign to the video files and playlists
	//    that it stores in your Amazon S3 bucket.
	ContentConfig *PipelineOutputConfig `type:"structure"`

	// The Amazon S3 bucket in which you saved the media files that you want to
	// transcode.
	//
	// InputBucket is a required field
	InputBucket *string `type:"string" required:"true"`

	// The name of the pipeline. We recommend that the name be unique within the
	// AWS account, but uniqueness is not enforced.
	//
	// Constraints: Maximum 40 characters.
	//
	// Name is a required field
	Name *string `min:"1" type:"string" required:"true"`

	// The Amazon Simple Notification Service (Amazon SNS) topic that you want to
	// notify to report job status.
	//
	// To receive notifications, you must also subscribe to the new topic in the
	// Amazon SNS console.
	//
	//    * Progressing: The topic ARN for the Amazon Simple Notification Service
	//    (Amazon SNS) topic that you want to notify when Elastic Transcoder has
	//    started to process a job in this pipeline. This is the ARN that Amazon
	//    SNS returned when you created the topic. For more information, see Create
	//    a Topic in the Amazon Simple Notification Service Developer Guide.
	//
	//    * Complete: The topic ARN for the Amazon SNS topic that you want to notify
	//    when Elastic Transcoder has finished processing a job in this pipeline.
	//    This is the ARN that Amazon SNS returned when you created the topic.
	//
	//    * Warning: The topic ARN for the Amazon SNS topic that you want to notify
	//    when Elastic Transcoder encounters a warning condition while processing
	//    a job in this pipeline. This is the ARN that Amazon SNS returned when
	//    you created the topic.
	//
	//    * Error: The topic ARN for the Amazon SNS topic that you want to notify
	//    when Elastic Transcoder encounters an error condition while processing
	//    a job in this pipeline. This is the ARN that Amazon SNS returned when
	//    you created the topic.
	Notifications *Notifications `type:"structure"`

	// The Amazon S3 bucket in which you want Elastic Transcoder to save the transcoded
	// files. (Use this, or use ContentConfig:Bucket plus ThumbnailConfig:Bucket.)
	//
	// Specify this value when all of the following are true:
	//
	//    * You want to save transcoded files, thumbnails (if any), and playlists
	//    (if any) together in one bucket.
	//
	//    * You do not want to specify the users or groups who have access to the
	//    transcoded files, thumbnails, and playlists.
	//
	//    * You do not want to specify the permissions that Elastic Transcoder grants
	//    to the files.
	//
	// When Elastic Transcoder saves files in OutputBucket, it grants full control
	//    over the files only to the AWS account that owns the role that is specified
	//    by Role.
	//
	//    * You want to associate the transcoded files and thumbnails with the Amazon
	//    S3 Standard storage class.
	//
	// If you want to save transcoded files and playlists in one bucket and thumbnails
	// in another bucket, specify which users can access the transcoded files or
	// the permissions the users have, or change the Amazon S3 storage class, omit
	// OutputBucket and specify values for ContentConfig and ThumbnailConfig instead.
	OutputBucket *string `type:"string"`

	// The IAM Amazon Resource Name (ARN) for the role that you want Elastic Transcoder
	// to use to create the pipeline.
	//
	// Role is a required field
	Role *string `type:"string" required:"true"`

	// The ThumbnailConfig object specifies several values, including the Amazon
	// S3 bucket in which you want Elastic Transcoder to save thumbnail files, which
	// users you want to have access to the files, the type of access you want users
	// to have, and the storage class that you want to assign to the files.
	//
	// If you specify values for ContentConfig, you must also specify values for
	// ThumbnailConfig even if you don't want to create thumbnails.
	//
	// If you specify values for ContentConfig and ThumbnailConfig, omit the OutputBucket
	// object.
	//
	//    * Bucket: The Amazon S3 bucket in which you want Elastic Transcoder to
	//    save thumbnail files.
	//
	//    * Permissions (Optional): The Permissions object specifies which users
	//    and/or predefined Amazon S3 groups you want to have access to thumbnail
	//    files, and the type of access you want them to have. You can grant permissions
	//    to a maximum of 30 users and/or predefined Amazon S3 groups.
	//
	//    * GranteeType: Specify the type of value that appears in the Grantee object:
	//
	//
	// Canonical: The value in the Grantee object is either the canonical user ID
	//    for an AWS account or an origin access identity for an Amazon CloudFront
	//    distribution.
	//
	// A canonical user ID is not the same as an AWS account number.
	//
	// Email: The value in the Grantee object is the registered email address of
	//    an AWS account.
	//
	// Group: The value in the Grantee object is one of the following predefined
	//    Amazon S3 groups: AllUsers, AuthenticatedUsers, or LogDelivery.
	//
	//    * Grantee: The AWS user or group that you want to have access to thumbnail
	//    files. To identify the user or group, you can specify the canonical user
	//    ID for an AWS account, an origin access identity for a CloudFront distribution,
	//    the registered email address of an AWS account, or a predefined Amazon
	//    S3 group.
	//
	//    * Access: The permission that you want to give to the AWS user that you
	//    specified in Grantee. Permissions are granted on the thumbnail files that
	//    Elastic Transcoder adds to the bucket. Valid values include:
	//
	// READ: The grantee can read the thumbnails and metadata for objects that Elastic
	//    Transcoder adds to the Amazon S3 bucket.
	//
	// READ_ACP: The grantee can read the object ACL for thumbnails that Elastic
	//    Transcoder adds to the Amazon S3 bucket.
	//
	// WRITE_ACP: The grantee can write the ACL for the thumbnails that Elastic
	//    Transcoder adds to the Amazon S3 bucket.
	//
	// FULL_CONTROL: The grantee has READ, READ_ACP, and WRITE_ACP permissions for
	//    the thumbnails that Elastic Transcoder adds to the Amazon S3 bucket.
	//
	//    * StorageClass: The Amazon S3 storage class, Standard or ReducedRedundancy,
	//    that you want Elastic Transcoder to assign to the thumbnails that it stores
	//    in your Amazon S3 bucket.
	ThumbnailConfig *PipelineOutputConfig `type:"structure"`
}

// String returns the string representation
func (s CreatePipelineInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreatePipelineInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreatePipelineInput"}

	if s.InputBucket == nil {
		invalidParams.Add(aws.NewErrParamRequired("InputBucket"))
	}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}
	if s.Name != nil && len(*s.Name) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Name", 1))
	}

	if s.Role == nil {
		invalidParams.Add(aws.NewErrParamRequired("Role"))
	}
	if s.ContentConfig != nil {
		if err := s.ContentConfig.Validate(); err != nil {
			invalidParams.AddNested("ContentConfig", err.(aws.ErrInvalidParams))
		}
	}
	if s.ThumbnailConfig != nil {
		if err := s.ThumbnailConfig.Validate(); err != nil {
			invalidParams.AddNested("ThumbnailConfig", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreatePipelineInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.AwsKmsKeyArn != nil {
		v := *s.AwsKmsKeyArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "AwsKmsKeyArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ContentConfig != nil {
		v := s.ContentConfig

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "ContentConfig", v, metadata)
	}
	if s.InputBucket != nil {
		v := *s.InputBucket

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "InputBucket", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Notifications != nil {
		v := s.Notifications

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "Notifications", v, metadata)
	}
	if s.OutputBucket != nil {
		v := *s.OutputBucket

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "OutputBucket", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Role != nil {
		v := *s.Role

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Role", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ThumbnailConfig != nil {
		v := s.ThumbnailConfig

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "ThumbnailConfig", v, metadata)
	}
	return nil
}

// When you create a pipeline, Elastic Transcoder returns the values that you
// specified in the request.
type CreatePipelineOutput struct {
	_ struct{} `type:"structure"`

	// A section of the response body that provides information about the pipeline
	// that is created.
	Pipeline *Pipeline `type:"structure"`

	// Elastic Transcoder returns a warning if the resources used by your pipeline
	// are not in the same region as the pipeline.
	//
	// Using resources in the same region, such as your Amazon S3 buckets, Amazon
	// SNS notification topics, and AWS KMS key, reduces processing time and prevents
	// cross-regional charges.
	Warnings []Warning `type:"list"`
}

// String returns the string representation
func (s CreatePipelineOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreatePipelineOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Pipeline != nil {
		v := s.Pipeline

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "Pipeline", v, metadata)
	}
	if len(s.Warnings) > 0 {
		v := s.Warnings

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "Warnings", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	return nil
}

const opCreatePipeline = "CreatePipeline"

// CreatePipelineRequest returns a request value for making API operation for
// Amazon Elastic Transcoder.
//
// The CreatePipeline operation creates a pipeline with settings that you specify.
//
//    // Example sending a request using CreatePipelineRequest.
//    req := client.CreatePipelineRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) CreatePipelineRequest(input *CreatePipelineInput) CreatePipelineRequest {
	op := &aws.Operation{
		Name:       opCreatePipeline,
		HTTPMethod: "POST",
		HTTPPath:   "/2012-09-25/pipelines",
	}

	if input == nil {
		input = &CreatePipelineInput{}
	}

	req := c.newRequest(op, input, &CreatePipelineOutput{})
	return CreatePipelineRequest{Request: req, Input: input, Copy: c.CreatePipelineRequest}
}

// CreatePipelineRequest is the request type for the
// CreatePipeline API operation.
type CreatePipelineRequest struct {
	*aws.Request
	Input *CreatePipelineInput
	Copy  func(*CreatePipelineInput) CreatePipelineRequest
}

// Send marshals and sends the CreatePipeline API request.
func (r CreatePipelineRequest) Send(ctx context.Context) (*CreatePipelineResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreatePipelineResponse{
		CreatePipelineOutput: r.Request.Data.(*CreatePipelineOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreatePipelineResponse is the response type for the
// CreatePipeline API operation.
type CreatePipelineResponse struct {
	*CreatePipelineOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreatePipeline request.
func (r *CreatePipelineResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
