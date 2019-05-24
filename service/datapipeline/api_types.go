// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package datapipeline

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

var _ aws.Config
var _ = awsutil.Prettify

// A key-value pair that describes a property of a pipeline object. The value
// is specified as either a string value (StringValue) or a reference to another
// object (RefValue) but not as both.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/datapipeline-2012-10-29/Field
type Field struct {
	_ struct{} `type:"structure"`

	// The field identifier.
	//
	// Key is a required field
	Key *string `locationName:"key" min:"1" type:"string" required:"true"`

	// The field value, expressed as the identifier of another object.
	RefValue *string `locationName:"refValue" min:"1" type:"string"`

	// The field value, expressed as a String.
	StringValue *string `locationName:"stringValue" type:"string"`
}

// String returns the string representation
func (s Field) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *Field) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "Field"}

	if s.Key == nil {
		invalidParams.Add(aws.NewErrParamRequired("Key"))
	}
	if s.Key != nil && len(*s.Key) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Key", 1))
	}
	if s.RefValue != nil && len(*s.RefValue) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("RefValue", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Identity information for the EC2 instance that is hosting the task runner.
// You can get this value by calling a metadata URI from the EC2 instance. For
// more information, see Instance Metadata (http://docs.aws.amazon.com/AWSEC2/latest/UserGuide/AESDG-chapter-instancedata.html)
// in the Amazon Elastic Compute Cloud User Guide. Passing in this value proves
// that your task runner is running on an EC2 instance, and ensures the proper
// AWS Data Pipeline service charges are applied to your pipeline.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/datapipeline-2012-10-29/InstanceIdentity
type InstanceIdentity struct {
	_ struct{} `type:"structure"`

	// A description of an EC2 instance that is generated when the instance is launched
	// and exposed to the instance via the instance metadata service in the form
	// of a JSON representation of an object.
	Document *string `locationName:"document" type:"string"`

	// A signature which can be used to verify the accuracy and authenticity of
	// the information provided in the instance identity document.
	Signature *string `locationName:"signature" type:"string"`
}

// String returns the string representation
func (s InstanceIdentity) String() string {
	return awsutil.Prettify(s)
}

// Contains a logical operation for comparing the value of a field with a specified
// value.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/datapipeline-2012-10-29/Operator
type Operator struct {
	_ struct{} `type:"structure"`

	// The logical operation to be performed: equal (EQ), equal reference (REF_EQ),
	// less than or equal (LE), greater than or equal (GE), or between (BETWEEN).
	// Equal reference (REF_EQ) can be used only with reference fields. The other
	// comparison types can be used only with String fields. The comparison types
	// you can use apply only to certain object fields, as detailed below.
	//
	// The comparison operators EQ and REF_EQ act on the following fields:
	//
	//    * name
	//    * @sphere
	//    * parent
	//    * @componentParent
	//    * @instanceParent
	//    * @status
	//    * @scheduledStartTime
	//    * @scheduledEndTime
	//    * @actualStartTime
	//    * @actualEndTime
	// The comparison operators GE, LE, and BETWEEN act on the following fields:
	//
	//    * @scheduledStartTime
	//    * @scheduledEndTime
	//    * @actualStartTime
	//    * @actualEndTime
	// Note that fields beginning with the at sign (@) are read-only and set by
	// the web service. When you name fields, you should choose names containing
	// only alpha-numeric values, as symbols may be reserved by AWS Data Pipeline.
	// User-defined fields that you add to a pipeline should prefix their name with
	// the string "my".
	Type OperatorType `locationName:"type" type:"string" enum:"true"`

	// The value that the actual field value will be compared with.
	Values []string `locationName:"values" type:"list"`
}

// String returns the string representation
func (s Operator) String() string {
	return awsutil.Prettify(s)
}

// The attributes allowed or specified with a parameter object.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/datapipeline-2012-10-29/ParameterAttribute
type ParameterAttribute struct {
	_ struct{} `type:"structure"`

	// The field identifier.
	//
	// Key is a required field
	Key *string `locationName:"key" min:"1" type:"string" required:"true"`

	// The field value, expressed as a String.
	//
	// StringValue is a required field
	StringValue *string `locationName:"stringValue" type:"string" required:"true"`
}

// String returns the string representation
func (s ParameterAttribute) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ParameterAttribute) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ParameterAttribute"}

	if s.Key == nil {
		invalidParams.Add(aws.NewErrParamRequired("Key"))
	}
	if s.Key != nil && len(*s.Key) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Key", 1))
	}

	if s.StringValue == nil {
		invalidParams.Add(aws.NewErrParamRequired("StringValue"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Contains information about a parameter object.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/datapipeline-2012-10-29/ParameterObject
type ParameterObject struct {
	_ struct{} `type:"structure"`

	// The attributes of the parameter object.
	//
	// Attributes is a required field
	Attributes []ParameterAttribute `locationName:"attributes" type:"list" required:"true"`

	// The ID of the parameter object.
	//
	// Id is a required field
	Id *string `locationName:"id" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s ParameterObject) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ParameterObject) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ParameterObject"}

	if s.Attributes == nil {
		invalidParams.Add(aws.NewErrParamRequired("Attributes"))
	}

	if s.Id == nil {
		invalidParams.Add(aws.NewErrParamRequired("Id"))
	}
	if s.Id != nil && len(*s.Id) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Id", 1))
	}
	if s.Attributes != nil {
		for i, v := range s.Attributes {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Attributes", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// A value or list of parameter values.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/datapipeline-2012-10-29/ParameterValue
type ParameterValue struct {
	_ struct{} `type:"structure"`

	// The ID of the parameter value.
	//
	// Id is a required field
	Id *string `locationName:"id" min:"1" type:"string" required:"true"`

	// The field value, expressed as a String.
	//
	// StringValue is a required field
	StringValue *string `locationName:"stringValue" type:"string" required:"true"`
}

// String returns the string representation
func (s ParameterValue) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ParameterValue) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ParameterValue"}

	if s.Id == nil {
		invalidParams.Add(aws.NewErrParamRequired("Id"))
	}
	if s.Id != nil && len(*s.Id) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Id", 1))
	}

	if s.StringValue == nil {
		invalidParams.Add(aws.NewErrParamRequired("StringValue"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Contains pipeline metadata.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/datapipeline-2012-10-29/PipelineDescription
type PipelineDescription struct {
	_ struct{} `type:"structure"`

	// Description of the pipeline.
	Description *string `locationName:"description" type:"string"`

	// A list of read-only fields that contain metadata about the pipeline: @userId,
	// @accountId, and @pipelineState.
	//
	// Fields is a required field
	Fields []Field `locationName:"fields" type:"list" required:"true"`

	// The name of the pipeline.
	//
	// Name is a required field
	Name *string `locationName:"name" min:"1" type:"string" required:"true"`

	// The pipeline identifier that was assigned by AWS Data Pipeline. This is a
	// string of the form df-297EG78HU43EEXAMPLE.
	//
	// PipelineId is a required field
	PipelineId *string `locationName:"pipelineId" min:"1" type:"string" required:"true"`

	// A list of tags to associated with a pipeline. Tags let you control access
	// to pipelines. For more information, see Controlling User Access to Pipelines
	// (http://docs.aws.amazon.com/datapipeline/latest/DeveloperGuide/dp-control-access.html)
	// in the AWS Data Pipeline Developer Guide.
	Tags []Tag `locationName:"tags" type:"list"`
}

// String returns the string representation
func (s PipelineDescription) String() string {
	return awsutil.Prettify(s)
}

// Contains the name and identifier of a pipeline.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/datapipeline-2012-10-29/PipelineIdName
type PipelineIdName struct {
	_ struct{} `type:"structure"`

	// The ID of the pipeline that was assigned by AWS Data Pipeline. This is a
	// string of the form df-297EG78HU43EEXAMPLE.
	Id *string `locationName:"id" min:"1" type:"string"`

	// The name of the pipeline.
	Name *string `locationName:"name" min:"1" type:"string"`
}

// String returns the string representation
func (s PipelineIdName) String() string {
	return awsutil.Prettify(s)
}

// Contains information about a pipeline object. This can be a logical, physical,
// or physical attempt pipeline object. The complete set of components of a
// pipeline defines the pipeline.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/datapipeline-2012-10-29/PipelineObject
type PipelineObject struct {
	_ struct{} `type:"structure"`

	// Key-value pairs that define the properties of the object.
	//
	// Fields is a required field
	Fields []Field `locationName:"fields" type:"list" required:"true"`

	// The ID of the object.
	//
	// Id is a required field
	Id *string `locationName:"id" min:"1" type:"string" required:"true"`

	// The name of the object.
	//
	// Name is a required field
	Name *string `locationName:"name" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s PipelineObject) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PipelineObject) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "PipelineObject"}

	if s.Fields == nil {
		invalidParams.Add(aws.NewErrParamRequired("Fields"))
	}

	if s.Id == nil {
		invalidParams.Add(aws.NewErrParamRequired("Id"))
	}
	if s.Id != nil && len(*s.Id) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Id", 1))
	}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}
	if s.Name != nil && len(*s.Name) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Name", 1))
	}
	if s.Fields != nil {
		for i, v := range s.Fields {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Fields", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Defines the query to run against an object.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/datapipeline-2012-10-29/Query
type Query struct {
	_ struct{} `type:"structure"`

	// List of selectors that define the query. An object must satisfy all of the
	// selectors to match the query.
	Selectors []Selector `locationName:"selectors" type:"list"`
}

// String returns the string representation
func (s Query) String() string {
	return awsutil.Prettify(s)
}

// A comparision that is used to determine whether a query should return this
// object.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/datapipeline-2012-10-29/Selector
type Selector struct {
	_ struct{} `type:"structure"`

	// The name of the field that the operator will be applied to. The field name
	// is the "key" portion of the field definition in the pipeline definition syntax
	// that is used by the AWS Data Pipeline API. If the field is not set on the
	// object, the condition fails.
	FieldName *string `locationName:"fieldName" type:"string"`

	// Contains a logical operation for comparing the value of a field with a specified
	// value.
	Operator *Operator `locationName:"operator" type:"structure"`
}

// String returns the string representation
func (s Selector) String() string {
	return awsutil.Prettify(s)
}

// Tags are key/value pairs defined by a user and associated with a pipeline
// to control access. AWS Data Pipeline allows you to associate ten tags per
// pipeline. For more information, see Controlling User Access to Pipelines
// (http://docs.aws.amazon.com/datapipeline/latest/DeveloperGuide/dp-control-access.html)
// in the AWS Data Pipeline Developer Guide.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/datapipeline-2012-10-29/Tag
type Tag struct {
	_ struct{} `type:"structure"`

	// The key name of a tag defined by a user. For more information, see Controlling
	// User Access to Pipelines (http://docs.aws.amazon.com/datapipeline/latest/DeveloperGuide/dp-control-access.html)
	// in the AWS Data Pipeline Developer Guide.
	//
	// Key is a required field
	Key *string `locationName:"key" min:"1" type:"string" required:"true"`

	// The optional value portion of a tag defined by a user. For more information,
	// see Controlling User Access to Pipelines (http://docs.aws.amazon.com/datapipeline/latest/DeveloperGuide/dp-control-access.html)
	// in the AWS Data Pipeline Developer Guide.
	//
	// Value is a required field
	Value *string `locationName:"value" type:"string" required:"true"`
}

// String returns the string representation
func (s Tag) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *Tag) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "Tag"}

	if s.Key == nil {
		invalidParams.Add(aws.NewErrParamRequired("Key"))
	}
	if s.Key != nil && len(*s.Key) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Key", 1))
	}

	if s.Value == nil {
		invalidParams.Add(aws.NewErrParamRequired("Value"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Contains information about a pipeline task that is assigned to a task runner.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/datapipeline-2012-10-29/TaskObject
type TaskObject struct {
	_ struct{} `type:"structure"`

	// The ID of the pipeline task attempt object. AWS Data Pipeline uses this value
	// to track how many times a task is attempted.
	AttemptId *string `locationName:"attemptId" min:"1" type:"string"`

	// Connection information for the location where the task runner will publish
	// the output of the task.
	Objects map[string]PipelineObject `locationName:"objects" type:"map"`

	// The ID of the pipeline that provided the task.
	PipelineId *string `locationName:"pipelineId" min:"1" type:"string"`

	// An internal identifier for the task. This ID is passed to the SetTaskStatus
	// and ReportTaskProgress actions.
	TaskId *string `locationName:"taskId" min:"1" type:"string"`
}

// String returns the string representation
func (s TaskObject) String() string {
	return awsutil.Prettify(s)
}

// Defines a validation error. Validation errors prevent pipeline activation.
// The set of validation errors that can be returned are defined by AWS Data
// Pipeline.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/datapipeline-2012-10-29/ValidationError
type ValidationError struct {
	_ struct{} `type:"structure"`

	// A description of the validation error.
	Errors []string `locationName:"errors" type:"list"`

	// The identifier of the object that contains the validation error.
	Id *string `locationName:"id" min:"1" type:"string"`
}

// String returns the string representation
func (s ValidationError) String() string {
	return awsutil.Prettify(s)
}

// Defines a validation warning. Validation warnings do not prevent pipeline
// activation. The set of validation warnings that can be returned are defined
// by AWS Data Pipeline.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/datapipeline-2012-10-29/ValidationWarning
type ValidationWarning struct {
	_ struct{} `type:"structure"`

	// The identifier of the object that contains the validation warning.
	Id *string `locationName:"id" min:"1" type:"string"`

	// A description of the validation warning.
	Warnings []string `locationName:"warnings" type:"list"`
}

// String returns the string representation
func (s ValidationWarning) String() string {
	return awsutil.Prettify(s)
}
