// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iotanalytics

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/iotanalytics-2017-11-27/UpdateDatasetRequest
type UpdateDatasetInput struct {
	_ struct{} `type:"structure"`

	// A list of "DatasetAction" objects.
	//
	// Actions is a required field
	Actions []DatasetAction `locationName:"actions" min:"1" type:"list" required:"true"`

	// When data set contents are created they are delivered to destinations specified
	// here.
	ContentDeliveryRules []DatasetContentDeliveryRule `locationName:"contentDeliveryRules" type:"list"`

	// The name of the data set to update.
	//
	// DatasetName is a required field
	DatasetName *string `location:"uri" locationName:"datasetName" min:"1" type:"string" required:"true"`

	// How long, in days, data set contents are kept for the data set.
	RetentionPeriod *RetentionPeriod `locationName:"retentionPeriod" type:"structure"`

	// A list of "DatasetTrigger" objects. The list can be empty or can contain
	// up to five DataSetTrigger objects.
	Triggers []DatasetTrigger `locationName:"triggers" type:"list"`

	// [Optional] How many versions of data set contents are kept. If not specified
	// or set to null, only the latest version plus the latest succeeded version
	// (if they are different) are kept for the time period specified by the "retentionPeriod"
	// parameter. (For more information, see https://docs.aws.amazon.com/iotanalytics/latest/userguide/getting-started.html#aws-iot-analytics-dataset-versions)
	VersioningConfiguration *VersioningConfiguration `locationName:"versioningConfiguration" type:"structure"`
}

// String returns the string representation
func (s UpdateDatasetInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateDatasetInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateDatasetInput"}

	if s.Actions == nil {
		invalidParams.Add(aws.NewErrParamRequired("Actions"))
	}
	if s.Actions != nil && len(s.Actions) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Actions", 1))
	}

	if s.DatasetName == nil {
		invalidParams.Add(aws.NewErrParamRequired("DatasetName"))
	}
	if s.DatasetName != nil && len(*s.DatasetName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("DatasetName", 1))
	}
	if s.Actions != nil {
		for i, v := range s.Actions {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Actions", i), err.(aws.ErrInvalidParams))
			}
		}
	}
	if s.ContentDeliveryRules != nil {
		for i, v := range s.ContentDeliveryRules {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "ContentDeliveryRules", i), err.(aws.ErrInvalidParams))
			}
		}
	}
	if s.RetentionPeriod != nil {
		if err := s.RetentionPeriod.Validate(); err != nil {
			invalidParams.AddNested("RetentionPeriod", err.(aws.ErrInvalidParams))
		}
	}
	if s.Triggers != nil {
		for i, v := range s.Triggers {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Triggers", i), err.(aws.ErrInvalidParams))
			}
		}
	}
	if s.VersioningConfiguration != nil {
		if err := s.VersioningConfiguration.Validate(); err != nil {
			invalidParams.AddNested("VersioningConfiguration", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateDatasetInput) MarshalFields(e protocol.FieldEncoder) error {

	if len(s.Actions) > 0 {
		v := s.Actions

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "actions", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if len(s.ContentDeliveryRules) > 0 {
		v := s.ContentDeliveryRules

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "contentDeliveryRules", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.RetentionPeriod != nil {
		v := s.RetentionPeriod

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "retentionPeriod", v, metadata)
	}
	if len(s.Triggers) > 0 {
		v := s.Triggers

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "triggers", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.VersioningConfiguration != nil {
		v := s.VersioningConfiguration

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "versioningConfiguration", v, metadata)
	}
	if s.DatasetName != nil {
		v := *s.DatasetName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "datasetName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/iotanalytics-2017-11-27/UpdateDatasetOutput
type UpdateDatasetOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s UpdateDatasetOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateDatasetOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opUpdateDataset = "UpdateDataset"

// UpdateDatasetRequest returns a request value for making API operation for
// AWS IoT Analytics.
//
// Updates the settings of a data set.
//
//    // Example sending a request using UpdateDatasetRequest.
//    req := client.UpdateDatasetRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iotanalytics-2017-11-27/UpdateDataset
func (c *Client) UpdateDatasetRequest(input *UpdateDatasetInput) UpdateDatasetRequest {
	op := &aws.Operation{
		Name:       opUpdateDataset,
		HTTPMethod: "PUT",
		HTTPPath:   "/datasets/{datasetName}",
	}

	if input == nil {
		input = &UpdateDatasetInput{}
	}

	req := c.newRequest(op, input, &UpdateDatasetOutput{})
	req.Handlers.Unmarshal.Remove(restjson.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return UpdateDatasetRequest{Request: req, Input: input, Copy: c.UpdateDatasetRequest}
}

// UpdateDatasetRequest is the request type for the
// UpdateDataset API operation.
type UpdateDatasetRequest struct {
	*aws.Request
	Input *UpdateDatasetInput
	Copy  func(*UpdateDatasetInput) UpdateDatasetRequest
}

// Send marshals and sends the UpdateDataset API request.
func (r UpdateDatasetRequest) Send(ctx context.Context) (*UpdateDatasetResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateDatasetResponse{
		UpdateDatasetOutput: r.Request.Data.(*UpdateDatasetOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateDatasetResponse is the response type for the
// UpdateDataset API operation.
type UpdateDatasetResponse struct {
	*UpdateDatasetOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateDataset request.
func (r *UpdateDatasetResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
