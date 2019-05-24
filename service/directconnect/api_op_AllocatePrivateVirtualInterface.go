// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package directconnect

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/directconnect-2012-10-25/AllocatePrivateVirtualInterfaceRequest
type AllocatePrivateVirtualInterfaceInput struct {
	_ struct{} `type:"structure"`

	// The ID of the connection on which the private virtual interface is provisioned.
	//
	// ConnectionId is a required field
	ConnectionId *string `locationName:"connectionId" type:"string" required:"true"`

	// Information about the private virtual interface.
	//
	// NewPrivateVirtualInterfaceAllocation is a required field
	NewPrivateVirtualInterfaceAllocation *NewPrivateVirtualInterfaceAllocation `locationName:"newPrivateVirtualInterfaceAllocation" type:"structure" required:"true"`

	// The ID of the AWS account that owns the virtual private interface.
	//
	// OwnerAccount is a required field
	OwnerAccount *string `locationName:"ownerAccount" type:"string" required:"true"`
}

// String returns the string representation
func (s AllocatePrivateVirtualInterfaceInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AllocatePrivateVirtualInterfaceInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "AllocatePrivateVirtualInterfaceInput"}

	if s.ConnectionId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ConnectionId"))
	}

	if s.NewPrivateVirtualInterfaceAllocation == nil {
		invalidParams.Add(aws.NewErrParamRequired("NewPrivateVirtualInterfaceAllocation"))
	}

	if s.OwnerAccount == nil {
		invalidParams.Add(aws.NewErrParamRequired("OwnerAccount"))
	}
	if s.NewPrivateVirtualInterfaceAllocation != nil {
		if err := s.NewPrivateVirtualInterfaceAllocation.Validate(); err != nil {
			invalidParams.AddNested("NewPrivateVirtualInterfaceAllocation", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Information about a virtual interface.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/directconnect-2012-10-25/VirtualInterface
type AllocatePrivateVirtualInterfaceOutput struct {
	_ struct{} `type:"structure"`

	// The address family for the BGP peer.
	AddressFamily AddressFamily `locationName:"addressFamily" type:"string" enum:"true"`

	// The IP address assigned to the Amazon interface.
	AmazonAddress *string `locationName:"amazonAddress" type:"string"`

	// The autonomous system number (ASN) for the Amazon side of the connection.
	AmazonSideAsn *int64 `locationName:"amazonSideAsn" type:"long"`

	// The autonomous system (AS) number for Border Gateway Protocol (BGP) configuration.
	Asn *int64 `locationName:"asn" type:"integer"`

	// The authentication key for BGP configuration.
	AuthKey *string `locationName:"authKey" type:"string"`

	// The Direct Connect endpoint on which the virtual interface terminates.
	AwsDeviceV2 *string `locationName:"awsDeviceV2" type:"string"`

	// The BGP peers configured on this virtual interface.
	BgpPeers []BGPPeer `locationName:"bgpPeers" type:"list"`

	// The ID of the connection.
	ConnectionId *string `locationName:"connectionId" type:"string"`

	// The IP address assigned to the customer interface.
	CustomerAddress *string `locationName:"customerAddress" type:"string"`

	// The customer router configuration.
	CustomerRouterConfig *string `locationName:"customerRouterConfig" type:"string"`

	// The ID of the Direct Connect gateway.
	DirectConnectGatewayId *string `locationName:"directConnectGatewayId" type:"string"`

	// Indicates whether jumbo frames (9001 MTU) are supported.
	JumboFrameCapable *bool `locationName:"jumboFrameCapable" type:"boolean"`

	// The location of the connection.
	Location *string `locationName:"location" type:"string"`

	// The maximum transmission unit (MTU), in bytes. The supported values are 1500
	// and 9001. The default value is 1500.
	Mtu *int64 `locationName:"mtu" type:"integer"`

	// The ID of the AWS account that owns the virtual interface.
	OwnerAccount *string `locationName:"ownerAccount" type:"string"`

	// The AWS Region where the virtual interface is located.
	Region *string `locationName:"region" type:"string"`

	// The routes to be advertised to the AWS network in this Region. Applies to
	// public virtual interfaces.
	RouteFilterPrefixes []RouteFilterPrefix `locationName:"routeFilterPrefixes" type:"list"`

	// The ID of the virtual private gateway. Applies only to private virtual interfaces.
	VirtualGatewayId *string `locationName:"virtualGatewayId" deprecated:"true" type:"string"`

	// The ID of the virtual interface.
	VirtualInterfaceId *string `locationName:"virtualInterfaceId" type:"string"`

	// The name of the virtual interface assigned by the customer network.
	VirtualInterfaceName *string `locationName:"virtualInterfaceName" type:"string"`

	// The state of the virtual interface. The following are the possible values:
	//
	//    * confirming: The creation of the virtual interface is pending confirmation
	//    from the virtual interface owner. If the owner of the virtual interface
	//    is different from the owner of the connection on which it is provisioned,
	//    then the virtual interface will remain in this state until it is confirmed
	//    by the virtual interface owner.
	//
	//    * verifying: This state only applies to public virtual interfaces. Each
	//    public virtual interface needs validation before the virtual interface
	//    can be created.
	//
	//    * pending: A virtual interface is in this state from the time that it
	//    is created until the virtual interface is ready to forward traffic.
	//
	//    * available: A virtual interface that is able to forward traffic.
	//
	//    * down: A virtual interface that is BGP down.
	//
	//    * deleting: A virtual interface is in this state immediately after calling
	//    DeleteVirtualInterface until it can no longer forward traffic.
	//
	//    * deleted: A virtual interface that cannot forward traffic.
	//
	//    * rejected: The virtual interface owner has declined creation of the virtual
	//    interface. If a virtual interface in the Confirming state is deleted by
	//    the virtual interface owner, the virtual interface enters the Rejected
	//    state.
	//
	//    * unknown: The state of the virtual interface is not available.
	VirtualInterfaceState VirtualInterfaceState `locationName:"virtualInterfaceState" type:"string" enum:"true"`

	// The type of virtual interface. The possible values are private and public.
	VirtualInterfaceType *string `locationName:"virtualInterfaceType" type:"string"`

	// The ID of the VLAN.
	Vlan *int64 `locationName:"vlan" type:"integer"`
}

// String returns the string representation
func (s AllocatePrivateVirtualInterfaceOutput) String() string {
	return awsutil.Prettify(s)
}

const opAllocatePrivateVirtualInterface = "AllocatePrivateVirtualInterface"

// AllocatePrivateVirtualInterfaceRequest returns a request value for making API operation for
// AWS Direct Connect.
//
// Provisions a private virtual interface to be owned by the specified AWS account.
//
// Virtual interfaces created using this action must be confirmed by the owner
// using ConfirmPrivateVirtualInterface. Until then, the virtual interface is
// in the Confirming state and is not available to handle traffic.
//
//    // Example sending a request using AllocatePrivateVirtualInterfaceRequest.
//    req := client.AllocatePrivateVirtualInterfaceRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/directconnect-2012-10-25/AllocatePrivateVirtualInterface
func (c *Client) AllocatePrivateVirtualInterfaceRequest(input *AllocatePrivateVirtualInterfaceInput) AllocatePrivateVirtualInterfaceRequest {
	op := &aws.Operation{
		Name:       opAllocatePrivateVirtualInterface,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &AllocatePrivateVirtualInterfaceInput{}
	}

	req := c.newRequest(op, input, &AllocatePrivateVirtualInterfaceOutput{})
	return AllocatePrivateVirtualInterfaceRequest{Request: req, Input: input, Copy: c.AllocatePrivateVirtualInterfaceRequest}
}

// AllocatePrivateVirtualInterfaceRequest is the request type for the
// AllocatePrivateVirtualInterface API operation.
type AllocatePrivateVirtualInterfaceRequest struct {
	*aws.Request
	Input *AllocatePrivateVirtualInterfaceInput
	Copy  func(*AllocatePrivateVirtualInterfaceInput) AllocatePrivateVirtualInterfaceRequest
}

// Send marshals and sends the AllocatePrivateVirtualInterface API request.
func (r AllocatePrivateVirtualInterfaceRequest) Send(ctx context.Context) (*AllocatePrivateVirtualInterfaceResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &AllocatePrivateVirtualInterfaceResponse{
		AllocatePrivateVirtualInterfaceOutput: r.Request.Data.(*AllocatePrivateVirtualInterfaceOutput),
		response:                              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// AllocatePrivateVirtualInterfaceResponse is the response type for the
// AllocatePrivateVirtualInterface API operation.
type AllocatePrivateVirtualInterfaceResponse struct {
	*AllocatePrivateVirtualInterfaceOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// AllocatePrivateVirtualInterface request.
func (r *AllocatePrivateVirtualInterfaceResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
