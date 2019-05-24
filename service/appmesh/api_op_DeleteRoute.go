// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package appmesh

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/appmesh-2019-01-25/DeleteRouteInput
type DeleteRouteInput struct {
	_ struct{} `type:"structure"`

	// The name of the service mesh to delete the route in.
	//
	// MeshName is a required field
	MeshName *string `location:"uri" locationName:"meshName" min:"1" type:"string" required:"true"`

	// The name of the route to delete.
	//
	// RouteName is a required field
	RouteName *string `location:"uri" locationName:"routeName" min:"1" type:"string" required:"true"`

	// The name of the virtual router to delete the route in.
	//
	// VirtualRouterName is a required field
	VirtualRouterName *string `location:"uri" locationName:"virtualRouterName" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteRouteInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteRouteInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteRouteInput"}

	if s.MeshName == nil {
		invalidParams.Add(aws.NewErrParamRequired("MeshName"))
	}
	if s.MeshName != nil && len(*s.MeshName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("MeshName", 1))
	}

	if s.RouteName == nil {
		invalidParams.Add(aws.NewErrParamRequired("RouteName"))
	}
	if s.RouteName != nil && len(*s.RouteName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("RouteName", 1))
	}

	if s.VirtualRouterName == nil {
		invalidParams.Add(aws.NewErrParamRequired("VirtualRouterName"))
	}
	if s.VirtualRouterName != nil && len(*s.VirtualRouterName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("VirtualRouterName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteRouteInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	if s.MeshName != nil {
		v := *s.MeshName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "meshName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.RouteName != nil {
		v := *s.RouteName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "routeName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.VirtualRouterName != nil {
		v := *s.VirtualRouterName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "virtualRouterName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/appmesh-2019-01-25/DeleteRouteOutput
type DeleteRouteOutput struct {
	_ struct{} `type:"structure" payload:"Route"`

	// The route that was deleted.
	//
	// Route is a required field
	Route *RouteData `locationName:"route" type:"structure" required:"true"`
}

// String returns the string representation
func (s DeleteRouteOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteRouteOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Route != nil {
		v := s.Route

		metadata := protocol.Metadata{}
		e.SetFields(protocol.PayloadTarget, "route", v, metadata)
	}
	return nil
}

const opDeleteRoute = "DeleteRoute"

// DeleteRouteRequest returns a request value for making API operation for
// AWS App Mesh.
//
// Deletes an existing route.
//
//    // Example sending a request using DeleteRouteRequest.
//    req := client.DeleteRouteRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/appmesh-2019-01-25/DeleteRoute
func (c *Client) DeleteRouteRequest(input *DeleteRouteInput) DeleteRouteRequest {
	op := &aws.Operation{
		Name:       opDeleteRoute,
		HTTPMethod: "DELETE",
		HTTPPath:   "/v20190125/meshes/{meshName}/virtualRouter/{virtualRouterName}/routes/{routeName}",
	}

	if input == nil {
		input = &DeleteRouteInput{}
	}

	req := c.newRequest(op, input, &DeleteRouteOutput{})
	return DeleteRouteRequest{Request: req, Input: input, Copy: c.DeleteRouteRequest}
}

// DeleteRouteRequest is the request type for the
// DeleteRoute API operation.
type DeleteRouteRequest struct {
	*aws.Request
	Input *DeleteRouteInput
	Copy  func(*DeleteRouteInput) DeleteRouteRequest
}

// Send marshals and sends the DeleteRoute API request.
func (r DeleteRouteRequest) Send(ctx context.Context) (*DeleteRouteResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteRouteResponse{
		DeleteRouteOutput: r.Request.Data.(*DeleteRouteOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteRouteResponse is the response type for the
// DeleteRoute API operation.
type DeleteRouteResponse struct {
	*DeleteRouteOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteRoute request.
func (r *DeleteRouteResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
