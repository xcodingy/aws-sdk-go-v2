// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package elasticsearchservice

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
)

type DeleteElasticsearchServiceRoleInput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteElasticsearchServiceRoleInput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteElasticsearchServiceRoleInput) MarshalFields(e protocol.FieldEncoder) error {

	return nil
}

type DeleteElasticsearchServiceRoleOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteElasticsearchServiceRoleOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteElasticsearchServiceRoleOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opDeleteElasticsearchServiceRole = "DeleteElasticsearchServiceRole"

// DeleteElasticsearchServiceRoleRequest returns a request value for making API operation for
// Amazon Elasticsearch Service.
//
// Deletes the service-linked role that Elasticsearch Service uses to manage
// and maintain VPC domains. Role deletion will fail if any existing VPC domains
// use the role. You must delete any such Elasticsearch domains before deleting
// the role. See Deleting Elasticsearch Service Role (http://docs.aws.amazon.com/elasticsearch-service/latest/developerguide/es-vpc.html#es-enabling-slr)
// in VPC Endpoints for Amazon Elasticsearch Service Domains.
//
//    // Example sending a request using DeleteElasticsearchServiceRoleRequest.
//    req := client.DeleteElasticsearchServiceRoleRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) DeleteElasticsearchServiceRoleRequest(input *DeleteElasticsearchServiceRoleInput) DeleteElasticsearchServiceRoleRequest {
	op := &aws.Operation{
		Name:       opDeleteElasticsearchServiceRole,
		HTTPMethod: "DELETE",
		HTTPPath:   "/2015-01-01/es/role",
	}

	if input == nil {
		input = &DeleteElasticsearchServiceRoleInput{}
	}

	req := c.newRequest(op, input, &DeleteElasticsearchServiceRoleOutput{})
	req.Handlers.Unmarshal.Remove(restjson.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return DeleteElasticsearchServiceRoleRequest{Request: req, Input: input, Copy: c.DeleteElasticsearchServiceRoleRequest}
}

// DeleteElasticsearchServiceRoleRequest is the request type for the
// DeleteElasticsearchServiceRole API operation.
type DeleteElasticsearchServiceRoleRequest struct {
	*aws.Request
	Input *DeleteElasticsearchServiceRoleInput
	Copy  func(*DeleteElasticsearchServiceRoleInput) DeleteElasticsearchServiceRoleRequest
}

// Send marshals and sends the DeleteElasticsearchServiceRole API request.
func (r DeleteElasticsearchServiceRoleRequest) Send(ctx context.Context) (*DeleteElasticsearchServiceRoleResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteElasticsearchServiceRoleResponse{
		DeleteElasticsearchServiceRoleOutput: r.Request.Data.(*DeleteElasticsearchServiceRoleOutput),
		response:                             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteElasticsearchServiceRoleResponse is the response type for the
// DeleteElasticsearchServiceRole API operation.
type DeleteElasticsearchServiceRoleResponse struct {
	*DeleteElasticsearchServiceRoleOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteElasticsearchServiceRole request.
func (r *DeleteElasticsearchServiceRoleResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
