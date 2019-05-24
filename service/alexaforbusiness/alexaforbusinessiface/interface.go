// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package alexaforbusinessiface provides an interface to enable mocking the Alexa For Business service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package alexaforbusinessiface

import (
	"github.com/aws/aws-sdk-go-v2/service/alexaforbusiness"
)

// ClientAPI provides an interface to enable mocking the
// alexaforbusiness.Client methods. This make unit testing your code that
// calls out to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // Alexa For Business.
//    func myFunc(svc alexaforbusinessiface.ClientAPI) bool {
//        // Make svc.ApproveSkill request
//    }
//
//    func main() {
//        cfg, err := external.LoadDefaultAWSConfig()
//        if err != nil {
//            panic("failed to load config, " + err.Error())
//        }
//
//        svc := alexaforbusiness.New(cfg)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockClientClient struct {
//        alexaforbusinessiface.ClientPI
//    }
//    func (m *mockClientClient) ApproveSkill(input *alexaforbusiness.ApproveSkillInput) (*alexaforbusiness.ApproveSkillOutput, error) {
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
	ApproveSkillRequest(*alexaforbusiness.ApproveSkillInput) alexaforbusiness.ApproveSkillRequest

	AssociateContactWithAddressBookRequest(*alexaforbusiness.AssociateContactWithAddressBookInput) alexaforbusiness.AssociateContactWithAddressBookRequest

	AssociateDeviceWithRoomRequest(*alexaforbusiness.AssociateDeviceWithRoomInput) alexaforbusiness.AssociateDeviceWithRoomRequest

	AssociateSkillGroupWithRoomRequest(*alexaforbusiness.AssociateSkillGroupWithRoomInput) alexaforbusiness.AssociateSkillGroupWithRoomRequest

	AssociateSkillWithSkillGroupRequest(*alexaforbusiness.AssociateSkillWithSkillGroupInput) alexaforbusiness.AssociateSkillWithSkillGroupRequest

	AssociateSkillWithUsersRequest(*alexaforbusiness.AssociateSkillWithUsersInput) alexaforbusiness.AssociateSkillWithUsersRequest

	CreateAddressBookRequest(*alexaforbusiness.CreateAddressBookInput) alexaforbusiness.CreateAddressBookRequest

	CreateBusinessReportScheduleRequest(*alexaforbusiness.CreateBusinessReportScheduleInput) alexaforbusiness.CreateBusinessReportScheduleRequest

	CreateConferenceProviderRequest(*alexaforbusiness.CreateConferenceProviderInput) alexaforbusiness.CreateConferenceProviderRequest

	CreateContactRequest(*alexaforbusiness.CreateContactInput) alexaforbusiness.CreateContactRequest

	CreateGatewayGroupRequest(*alexaforbusiness.CreateGatewayGroupInput) alexaforbusiness.CreateGatewayGroupRequest

	CreateProfileRequest(*alexaforbusiness.CreateProfileInput) alexaforbusiness.CreateProfileRequest

	CreateRoomRequest(*alexaforbusiness.CreateRoomInput) alexaforbusiness.CreateRoomRequest

	CreateSkillGroupRequest(*alexaforbusiness.CreateSkillGroupInput) alexaforbusiness.CreateSkillGroupRequest

	CreateUserRequest(*alexaforbusiness.CreateUserInput) alexaforbusiness.CreateUserRequest

	DeleteAddressBookRequest(*alexaforbusiness.DeleteAddressBookInput) alexaforbusiness.DeleteAddressBookRequest

	DeleteBusinessReportScheduleRequest(*alexaforbusiness.DeleteBusinessReportScheduleInput) alexaforbusiness.DeleteBusinessReportScheduleRequest

	DeleteConferenceProviderRequest(*alexaforbusiness.DeleteConferenceProviderInput) alexaforbusiness.DeleteConferenceProviderRequest

	DeleteContactRequest(*alexaforbusiness.DeleteContactInput) alexaforbusiness.DeleteContactRequest

	DeleteDeviceRequest(*alexaforbusiness.DeleteDeviceInput) alexaforbusiness.DeleteDeviceRequest

	DeleteGatewayGroupRequest(*alexaforbusiness.DeleteGatewayGroupInput) alexaforbusiness.DeleteGatewayGroupRequest

	DeleteProfileRequest(*alexaforbusiness.DeleteProfileInput) alexaforbusiness.DeleteProfileRequest

	DeleteRoomRequest(*alexaforbusiness.DeleteRoomInput) alexaforbusiness.DeleteRoomRequest

	DeleteRoomSkillParameterRequest(*alexaforbusiness.DeleteRoomSkillParameterInput) alexaforbusiness.DeleteRoomSkillParameterRequest

	DeleteSkillAuthorizationRequest(*alexaforbusiness.DeleteSkillAuthorizationInput) alexaforbusiness.DeleteSkillAuthorizationRequest

	DeleteSkillGroupRequest(*alexaforbusiness.DeleteSkillGroupInput) alexaforbusiness.DeleteSkillGroupRequest

	DeleteUserRequest(*alexaforbusiness.DeleteUserInput) alexaforbusiness.DeleteUserRequest

	DisassociateContactFromAddressBookRequest(*alexaforbusiness.DisassociateContactFromAddressBookInput) alexaforbusiness.DisassociateContactFromAddressBookRequest

	DisassociateDeviceFromRoomRequest(*alexaforbusiness.DisassociateDeviceFromRoomInput) alexaforbusiness.DisassociateDeviceFromRoomRequest

	DisassociateSkillFromSkillGroupRequest(*alexaforbusiness.DisassociateSkillFromSkillGroupInput) alexaforbusiness.DisassociateSkillFromSkillGroupRequest

	DisassociateSkillFromUsersRequest(*alexaforbusiness.DisassociateSkillFromUsersInput) alexaforbusiness.DisassociateSkillFromUsersRequest

	DisassociateSkillGroupFromRoomRequest(*alexaforbusiness.DisassociateSkillGroupFromRoomInput) alexaforbusiness.DisassociateSkillGroupFromRoomRequest

	ForgetSmartHomeAppliancesRequest(*alexaforbusiness.ForgetSmartHomeAppliancesInput) alexaforbusiness.ForgetSmartHomeAppliancesRequest

	GetAddressBookRequest(*alexaforbusiness.GetAddressBookInput) alexaforbusiness.GetAddressBookRequest

	GetConferencePreferenceRequest(*alexaforbusiness.GetConferencePreferenceInput) alexaforbusiness.GetConferencePreferenceRequest

	GetConferenceProviderRequest(*alexaforbusiness.GetConferenceProviderInput) alexaforbusiness.GetConferenceProviderRequest

	GetContactRequest(*alexaforbusiness.GetContactInput) alexaforbusiness.GetContactRequest

	GetDeviceRequest(*alexaforbusiness.GetDeviceInput) alexaforbusiness.GetDeviceRequest

	GetGatewayRequest(*alexaforbusiness.GetGatewayInput) alexaforbusiness.GetGatewayRequest

	GetGatewayGroupRequest(*alexaforbusiness.GetGatewayGroupInput) alexaforbusiness.GetGatewayGroupRequest

	GetInvitationConfigurationRequest(*alexaforbusiness.GetInvitationConfigurationInput) alexaforbusiness.GetInvitationConfigurationRequest

	GetProfileRequest(*alexaforbusiness.GetProfileInput) alexaforbusiness.GetProfileRequest

	GetRoomRequest(*alexaforbusiness.GetRoomInput) alexaforbusiness.GetRoomRequest

	GetRoomSkillParameterRequest(*alexaforbusiness.GetRoomSkillParameterInput) alexaforbusiness.GetRoomSkillParameterRequest

	GetSkillGroupRequest(*alexaforbusiness.GetSkillGroupInput) alexaforbusiness.GetSkillGroupRequest

	ListBusinessReportSchedulesRequest(*alexaforbusiness.ListBusinessReportSchedulesInput) alexaforbusiness.ListBusinessReportSchedulesRequest

	ListConferenceProvidersRequest(*alexaforbusiness.ListConferenceProvidersInput) alexaforbusiness.ListConferenceProvidersRequest

	ListDeviceEventsRequest(*alexaforbusiness.ListDeviceEventsInput) alexaforbusiness.ListDeviceEventsRequest

	ListGatewayGroupsRequest(*alexaforbusiness.ListGatewayGroupsInput) alexaforbusiness.ListGatewayGroupsRequest

	ListGatewaysRequest(*alexaforbusiness.ListGatewaysInput) alexaforbusiness.ListGatewaysRequest

	ListSkillsRequest(*alexaforbusiness.ListSkillsInput) alexaforbusiness.ListSkillsRequest

	ListSkillsStoreCategoriesRequest(*alexaforbusiness.ListSkillsStoreCategoriesInput) alexaforbusiness.ListSkillsStoreCategoriesRequest

	ListSkillsStoreSkillsByCategoryRequest(*alexaforbusiness.ListSkillsStoreSkillsByCategoryInput) alexaforbusiness.ListSkillsStoreSkillsByCategoryRequest

	ListSmartHomeAppliancesRequest(*alexaforbusiness.ListSmartHomeAppliancesInput) alexaforbusiness.ListSmartHomeAppliancesRequest

	ListTagsRequest(*alexaforbusiness.ListTagsInput) alexaforbusiness.ListTagsRequest

	PutConferencePreferenceRequest(*alexaforbusiness.PutConferencePreferenceInput) alexaforbusiness.PutConferencePreferenceRequest

	PutInvitationConfigurationRequest(*alexaforbusiness.PutInvitationConfigurationInput) alexaforbusiness.PutInvitationConfigurationRequest

	PutRoomSkillParameterRequest(*alexaforbusiness.PutRoomSkillParameterInput) alexaforbusiness.PutRoomSkillParameterRequest

	PutSkillAuthorizationRequest(*alexaforbusiness.PutSkillAuthorizationInput) alexaforbusiness.PutSkillAuthorizationRequest

	RegisterAVSDeviceRequest(*alexaforbusiness.RegisterAVSDeviceInput) alexaforbusiness.RegisterAVSDeviceRequest

	RejectSkillRequest(*alexaforbusiness.RejectSkillInput) alexaforbusiness.RejectSkillRequest

	ResolveRoomRequest(*alexaforbusiness.ResolveRoomInput) alexaforbusiness.ResolveRoomRequest

	RevokeInvitationRequest(*alexaforbusiness.RevokeInvitationInput) alexaforbusiness.RevokeInvitationRequest

	SearchAddressBooksRequest(*alexaforbusiness.SearchAddressBooksInput) alexaforbusiness.SearchAddressBooksRequest

	SearchContactsRequest(*alexaforbusiness.SearchContactsInput) alexaforbusiness.SearchContactsRequest

	SearchDevicesRequest(*alexaforbusiness.SearchDevicesInput) alexaforbusiness.SearchDevicesRequest

	SearchProfilesRequest(*alexaforbusiness.SearchProfilesInput) alexaforbusiness.SearchProfilesRequest

	SearchRoomsRequest(*alexaforbusiness.SearchRoomsInput) alexaforbusiness.SearchRoomsRequest

	SearchSkillGroupsRequest(*alexaforbusiness.SearchSkillGroupsInput) alexaforbusiness.SearchSkillGroupsRequest

	SearchUsersRequest(*alexaforbusiness.SearchUsersInput) alexaforbusiness.SearchUsersRequest

	SendInvitationRequest(*alexaforbusiness.SendInvitationInput) alexaforbusiness.SendInvitationRequest

	StartDeviceSyncRequest(*alexaforbusiness.StartDeviceSyncInput) alexaforbusiness.StartDeviceSyncRequest

	StartSmartHomeApplianceDiscoveryRequest(*alexaforbusiness.StartSmartHomeApplianceDiscoveryInput) alexaforbusiness.StartSmartHomeApplianceDiscoveryRequest

	TagResourceRequest(*alexaforbusiness.TagResourceInput) alexaforbusiness.TagResourceRequest

	UntagResourceRequest(*alexaforbusiness.UntagResourceInput) alexaforbusiness.UntagResourceRequest

	UpdateAddressBookRequest(*alexaforbusiness.UpdateAddressBookInput) alexaforbusiness.UpdateAddressBookRequest

	UpdateBusinessReportScheduleRequest(*alexaforbusiness.UpdateBusinessReportScheduleInput) alexaforbusiness.UpdateBusinessReportScheduleRequest

	UpdateConferenceProviderRequest(*alexaforbusiness.UpdateConferenceProviderInput) alexaforbusiness.UpdateConferenceProviderRequest

	UpdateContactRequest(*alexaforbusiness.UpdateContactInput) alexaforbusiness.UpdateContactRequest

	UpdateDeviceRequest(*alexaforbusiness.UpdateDeviceInput) alexaforbusiness.UpdateDeviceRequest

	UpdateGatewayRequest(*alexaforbusiness.UpdateGatewayInput) alexaforbusiness.UpdateGatewayRequest

	UpdateGatewayGroupRequest(*alexaforbusiness.UpdateGatewayGroupInput) alexaforbusiness.UpdateGatewayGroupRequest

	UpdateProfileRequest(*alexaforbusiness.UpdateProfileInput) alexaforbusiness.UpdateProfileRequest

	UpdateRoomRequest(*alexaforbusiness.UpdateRoomInput) alexaforbusiness.UpdateRoomRequest

	UpdateSkillGroupRequest(*alexaforbusiness.UpdateSkillGroupInput) alexaforbusiness.UpdateSkillGroupRequest
}

var _ ClientAPI = (*alexaforbusiness.Client)(nil)
