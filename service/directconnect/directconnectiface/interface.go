// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package directconnectiface provides an interface to enable mocking the AWS Direct Connect service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package directconnectiface

import (
	"github.com/aws/aws-sdk-go-v2/service/directconnect"
)

// ClientAPI provides an interface to enable mocking the
// directconnect.Client methods. This make unit testing your code that
// calls out to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // AWS Direct Connect.
//    func myFunc(svc directconnectiface.ClientAPI) bool {
//        // Make svc.AcceptDirectConnectGatewayAssociationProposal request
//    }
//
//    func main() {
//        cfg, err := external.LoadDefaultAWSConfig()
//        if err != nil {
//            panic("failed to load config, " + err.Error())
//        }
//
//        svc := directconnect.New(cfg)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockClientClient struct {
//        directconnectiface.ClientPI
//    }
//    func (m *mockClientClient) AcceptDirectConnectGatewayAssociationProposal(input *directconnect.AcceptDirectConnectGatewayAssociationProposalInput) (*directconnect.AcceptDirectConnectGatewayAssociationProposalOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockClientClient{}
//
//        myfunc(mockSvc)
//
//        // Verify myFunc's functionality
//    }
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type ClientAPI interface {
	AcceptDirectConnectGatewayAssociationProposalRequest(*directconnect.AcceptDirectConnectGatewayAssociationProposalInput) directconnect.AcceptDirectConnectGatewayAssociationProposalRequest

	AllocateConnectionOnInterconnectRequest(*directconnect.AllocateConnectionOnInterconnectInput) directconnect.AllocateConnectionOnInterconnectRequest

	AllocateHostedConnectionRequest(*directconnect.AllocateHostedConnectionInput) directconnect.AllocateHostedConnectionRequest

	AllocatePrivateVirtualInterfaceRequest(*directconnect.AllocatePrivateVirtualInterfaceInput) directconnect.AllocatePrivateVirtualInterfaceRequest

	AllocatePublicVirtualInterfaceRequest(*directconnect.AllocatePublicVirtualInterfaceInput) directconnect.AllocatePublicVirtualInterfaceRequest

	AssociateConnectionWithLagRequest(*directconnect.AssociateConnectionWithLagInput) directconnect.AssociateConnectionWithLagRequest

	AssociateHostedConnectionRequest(*directconnect.AssociateHostedConnectionInput) directconnect.AssociateHostedConnectionRequest

	AssociateVirtualInterfaceRequest(*directconnect.AssociateVirtualInterfaceInput) directconnect.AssociateVirtualInterfaceRequest

	ConfirmConnectionRequest(*directconnect.ConfirmConnectionInput) directconnect.ConfirmConnectionRequest

	ConfirmPrivateVirtualInterfaceRequest(*directconnect.ConfirmPrivateVirtualInterfaceInput) directconnect.ConfirmPrivateVirtualInterfaceRequest

	ConfirmPublicVirtualInterfaceRequest(*directconnect.ConfirmPublicVirtualInterfaceInput) directconnect.ConfirmPublicVirtualInterfaceRequest

	CreateBGPPeerRequest(*directconnect.CreateBGPPeerInput) directconnect.CreateBGPPeerRequest

	CreateConnectionRequest(*directconnect.CreateConnectionInput) directconnect.CreateConnectionRequest

	CreateDirectConnectGatewayRequest(*directconnect.CreateDirectConnectGatewayInput) directconnect.CreateDirectConnectGatewayRequest

	CreateDirectConnectGatewayAssociationRequest(*directconnect.CreateDirectConnectGatewayAssociationInput) directconnect.CreateDirectConnectGatewayAssociationRequest

	CreateDirectConnectGatewayAssociationProposalRequest(*directconnect.CreateDirectConnectGatewayAssociationProposalInput) directconnect.CreateDirectConnectGatewayAssociationProposalRequest

	CreateInterconnectRequest(*directconnect.CreateInterconnectInput) directconnect.CreateInterconnectRequest

	CreateLagRequest(*directconnect.CreateLagInput) directconnect.CreateLagRequest

	CreatePrivateVirtualInterfaceRequest(*directconnect.CreatePrivateVirtualInterfaceInput) directconnect.CreatePrivateVirtualInterfaceRequest

	CreatePublicVirtualInterfaceRequest(*directconnect.CreatePublicVirtualInterfaceInput) directconnect.CreatePublicVirtualInterfaceRequest

	DeleteBGPPeerRequest(*directconnect.DeleteBGPPeerInput) directconnect.DeleteBGPPeerRequest

	DeleteConnectionRequest(*directconnect.DeleteConnectionInput) directconnect.DeleteConnectionRequest

	DeleteDirectConnectGatewayRequest(*directconnect.DeleteDirectConnectGatewayInput) directconnect.DeleteDirectConnectGatewayRequest

	DeleteDirectConnectGatewayAssociationRequest(*directconnect.DeleteDirectConnectGatewayAssociationInput) directconnect.DeleteDirectConnectGatewayAssociationRequest

	DeleteDirectConnectGatewayAssociationProposalRequest(*directconnect.DeleteDirectConnectGatewayAssociationProposalInput) directconnect.DeleteDirectConnectGatewayAssociationProposalRequest

	DeleteInterconnectRequest(*directconnect.DeleteInterconnectInput) directconnect.DeleteInterconnectRequest

	DeleteLagRequest(*directconnect.DeleteLagInput) directconnect.DeleteLagRequest

	DeleteVirtualInterfaceRequest(*directconnect.DeleteVirtualInterfaceInput) directconnect.DeleteVirtualInterfaceRequest

	DescribeConnectionLoaRequest(*directconnect.DescribeConnectionLoaInput) directconnect.DescribeConnectionLoaRequest

	DescribeConnectionsRequest(*directconnect.DescribeConnectionsInput) directconnect.DescribeConnectionsRequest

	DescribeConnectionsOnInterconnectRequest(*directconnect.DescribeConnectionsOnInterconnectInput) directconnect.DescribeConnectionsOnInterconnectRequest

	DescribeDirectConnectGatewayAssociationProposalsRequest(*directconnect.DescribeDirectConnectGatewayAssociationProposalsInput) directconnect.DescribeDirectConnectGatewayAssociationProposalsRequest

	DescribeDirectConnectGatewayAssociationsRequest(*directconnect.DescribeDirectConnectGatewayAssociationsInput) directconnect.DescribeDirectConnectGatewayAssociationsRequest

	DescribeDirectConnectGatewayAttachmentsRequest(*directconnect.DescribeDirectConnectGatewayAttachmentsInput) directconnect.DescribeDirectConnectGatewayAttachmentsRequest

	DescribeDirectConnectGatewaysRequest(*directconnect.DescribeDirectConnectGatewaysInput) directconnect.DescribeDirectConnectGatewaysRequest

	DescribeHostedConnectionsRequest(*directconnect.DescribeHostedConnectionsInput) directconnect.DescribeHostedConnectionsRequest

	DescribeInterconnectLoaRequest(*directconnect.DescribeInterconnectLoaInput) directconnect.DescribeInterconnectLoaRequest

	DescribeInterconnectsRequest(*directconnect.DescribeInterconnectsInput) directconnect.DescribeInterconnectsRequest

	DescribeLagsRequest(*directconnect.DescribeLagsInput) directconnect.DescribeLagsRequest

	DescribeLoaRequest(*directconnect.DescribeLoaInput) directconnect.DescribeLoaRequest

	DescribeLocationsRequest(*directconnect.DescribeLocationsInput) directconnect.DescribeLocationsRequest

	DescribeTagsRequest(*directconnect.DescribeTagsInput) directconnect.DescribeTagsRequest

	DescribeVirtualGatewaysRequest(*directconnect.DescribeVirtualGatewaysInput) directconnect.DescribeVirtualGatewaysRequest

	DescribeVirtualInterfacesRequest(*directconnect.DescribeVirtualInterfacesInput) directconnect.DescribeVirtualInterfacesRequest

	DisassociateConnectionFromLagRequest(*directconnect.DisassociateConnectionFromLagInput) directconnect.DisassociateConnectionFromLagRequest

	TagResourceRequest(*directconnect.TagResourceInput) directconnect.TagResourceRequest

	UntagResourceRequest(*directconnect.UntagResourceInput) directconnect.UntagResourceRequest

	UpdateDirectConnectGatewayAssociationRequest(*directconnect.UpdateDirectConnectGatewayAssociationInput) directconnect.UpdateDirectConnectGatewayAssociationRequest

	UpdateLagRequest(*directconnect.UpdateLagInput) directconnect.UpdateLagRequest

	UpdateVirtualInterfaceAttributesRequest(*directconnect.UpdateVirtualInterfaceAttributesInput) directconnect.UpdateVirtualInterfaceAttributesRequest
}

var _ ClientAPI = (*directconnect.Client)(nil)
