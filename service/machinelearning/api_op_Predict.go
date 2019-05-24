// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package machinelearning

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type PredictInput struct {
	_ struct{} `type:"structure"`

	// A unique identifier of the MLModel.
	//
	// MLModelId is a required field
	MLModelId *string `min:"1" type:"string" required:"true"`

	// PredictEndpoint is a required field
	PredictEndpoint *string `type:"string" required:"true"`

	// A map of variable name-value pairs that represent an observation.
	//
	// Record is a required field
	Record map[string]string `type:"map" required:"true"`
}

// String returns the string representation
func (s PredictInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PredictInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "PredictInput"}

	if s.MLModelId == nil {
		invalidParams.Add(aws.NewErrParamRequired("MLModelId"))
	}
	if s.MLModelId != nil && len(*s.MLModelId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("MLModelId", 1))
	}

	if s.PredictEndpoint == nil {
		invalidParams.Add(aws.NewErrParamRequired("PredictEndpoint"))
	}

	if s.Record == nil {
		invalidParams.Add(aws.NewErrParamRequired("Record"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type PredictOutput struct {
	_ struct{} `type:"structure"`

	// The output from a Predict operation:
	//
	//    * Details - Contains the following attributes: DetailsAttributes.PREDICTIVE_MODEL_TYPE
	//    - REGRESSION | BINARY | MULTICLASSDetailsAttributes.ALGORITHM - SGD
	//
	//    * PredictedLabel - Present for either a BINARY or MULTICLASSMLModel request.
	//
	//
	//    * PredictedScores - Contains the raw classification score corresponding
	//    to each label.
	//
	//    * PredictedValue - Present for a REGRESSIONMLModel request.
	Prediction *Prediction `type:"structure"`
}

// String returns the string representation
func (s PredictOutput) String() string {
	return awsutil.Prettify(s)
}

const opPredict = "Predict"

// PredictRequest returns a request value for making API operation for
// Amazon Machine Learning.
//
// Generates a prediction for the observation using the specified ML Model.
//
// NoteNot all response parameters will be populated. Whether a response parameter
// is populated depends on the type of model requested.
//
//    // Example sending a request using PredictRequest.
//    req := client.PredictRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) PredictRequest(input *PredictInput) PredictRequest {
	op := &aws.Operation{
		Name:       opPredict,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &PredictInput{}
	}

	req := c.newRequest(op, input, &PredictOutput{})
	return PredictRequest{Request: req, Input: input, Copy: c.PredictRequest}
}

// PredictRequest is the request type for the
// Predict API operation.
type PredictRequest struct {
	*aws.Request
	Input *PredictInput
	Copy  func(*PredictInput) PredictRequest
}

// Send marshals and sends the Predict API request.
func (r PredictRequest) Send(ctx context.Context) (*PredictResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &PredictResponse{
		PredictOutput: r.Request.Data.(*PredictOutput),
		response:      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// PredictResponse is the response type for the
// Predict API operation.
type PredictResponse struct {
	*PredictOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// Predict request.
func (r *PredictResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
