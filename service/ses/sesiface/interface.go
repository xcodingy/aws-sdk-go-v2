// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package sesiface provides an interface to enable mocking the Amazon Simple Email Service service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package sesiface

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ses"
)

// ClientAPI provides an interface to enable mocking the
// ses.Client methods. This make unit testing your code that
// calls out to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // Amazon SES.
//    func myFunc(svc sesiface.ClientAPI) bool {
//        // Make svc.CloneReceiptRuleSet request
//    }
//
//    func main() {
//        cfg, err := external.LoadDefaultAWSConfig()
//        if err != nil {
//            panic("failed to load config, " + err.Error())
//        }
//
//        svc := ses.New(cfg)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockClientClient struct {
//        sesiface.ClientPI
//    }
//    func (m *mockClientClient) CloneReceiptRuleSet(input *ses.CloneReceiptRuleSetInput) (*ses.CloneReceiptRuleSetOutput, error) {
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
	CloneReceiptRuleSetRequest(*ses.CloneReceiptRuleSetInput) ses.CloneReceiptRuleSetRequest

	CreateConfigurationSetRequest(*ses.CreateConfigurationSetInput) ses.CreateConfigurationSetRequest

	CreateConfigurationSetEventDestinationRequest(*ses.CreateConfigurationSetEventDestinationInput) ses.CreateConfigurationSetEventDestinationRequest

	CreateConfigurationSetTrackingOptionsRequest(*ses.CreateConfigurationSetTrackingOptionsInput) ses.CreateConfigurationSetTrackingOptionsRequest

	CreateCustomVerificationEmailTemplateRequest(*ses.CreateCustomVerificationEmailTemplateInput) ses.CreateCustomVerificationEmailTemplateRequest

	CreateReceiptFilterRequest(*ses.CreateReceiptFilterInput) ses.CreateReceiptFilterRequest

	CreateReceiptRuleRequest(*ses.CreateReceiptRuleInput) ses.CreateReceiptRuleRequest

	CreateReceiptRuleSetRequest(*ses.CreateReceiptRuleSetInput) ses.CreateReceiptRuleSetRequest

	CreateTemplateRequest(*ses.CreateTemplateInput) ses.CreateTemplateRequest

	DeleteConfigurationSetRequest(*ses.DeleteConfigurationSetInput) ses.DeleteConfigurationSetRequest

	DeleteConfigurationSetEventDestinationRequest(*ses.DeleteConfigurationSetEventDestinationInput) ses.DeleteConfigurationSetEventDestinationRequest

	DeleteConfigurationSetTrackingOptionsRequest(*ses.DeleteConfigurationSetTrackingOptionsInput) ses.DeleteConfigurationSetTrackingOptionsRequest

	DeleteCustomVerificationEmailTemplateRequest(*ses.DeleteCustomVerificationEmailTemplateInput) ses.DeleteCustomVerificationEmailTemplateRequest

	DeleteIdentityRequest(*ses.DeleteIdentityInput) ses.DeleteIdentityRequest

	DeleteIdentityPolicyRequest(*ses.DeleteIdentityPolicyInput) ses.DeleteIdentityPolicyRequest

	DeleteReceiptFilterRequest(*ses.DeleteReceiptFilterInput) ses.DeleteReceiptFilterRequest

	DeleteReceiptRuleRequest(*ses.DeleteReceiptRuleInput) ses.DeleteReceiptRuleRequest

	DeleteReceiptRuleSetRequest(*ses.DeleteReceiptRuleSetInput) ses.DeleteReceiptRuleSetRequest

	DeleteTemplateRequest(*ses.DeleteTemplateInput) ses.DeleteTemplateRequest

	DeleteVerifiedEmailAddressRequest(*ses.DeleteVerifiedEmailAddressInput) ses.DeleteVerifiedEmailAddressRequest

	DescribeActiveReceiptRuleSetRequest(*ses.DescribeActiveReceiptRuleSetInput) ses.DescribeActiveReceiptRuleSetRequest

	DescribeConfigurationSetRequest(*ses.DescribeConfigurationSetInput) ses.DescribeConfigurationSetRequest

	DescribeReceiptRuleRequest(*ses.DescribeReceiptRuleInput) ses.DescribeReceiptRuleRequest

	DescribeReceiptRuleSetRequest(*ses.DescribeReceiptRuleSetInput) ses.DescribeReceiptRuleSetRequest

	GetAccountSendingEnabledRequest(*ses.GetAccountSendingEnabledInput) ses.GetAccountSendingEnabledRequest

	GetCustomVerificationEmailTemplateRequest(*ses.GetCustomVerificationEmailTemplateInput) ses.GetCustomVerificationEmailTemplateRequest

	GetIdentityDkimAttributesRequest(*ses.GetIdentityDkimAttributesInput) ses.GetIdentityDkimAttributesRequest

	GetIdentityMailFromDomainAttributesRequest(*ses.GetIdentityMailFromDomainAttributesInput) ses.GetIdentityMailFromDomainAttributesRequest

	GetIdentityNotificationAttributesRequest(*ses.GetIdentityNotificationAttributesInput) ses.GetIdentityNotificationAttributesRequest

	GetIdentityPoliciesRequest(*ses.GetIdentityPoliciesInput) ses.GetIdentityPoliciesRequest

	GetIdentityVerificationAttributesRequest(*ses.GetIdentityVerificationAttributesInput) ses.GetIdentityVerificationAttributesRequest

	GetSendQuotaRequest(*ses.GetSendQuotaInput) ses.GetSendQuotaRequest

	GetSendStatisticsRequest(*ses.GetSendStatisticsInput) ses.GetSendStatisticsRequest

	GetTemplateRequest(*ses.GetTemplateInput) ses.GetTemplateRequest

	ListConfigurationSetsRequest(*ses.ListConfigurationSetsInput) ses.ListConfigurationSetsRequest

	ListCustomVerificationEmailTemplatesRequest(*ses.ListCustomVerificationEmailTemplatesInput) ses.ListCustomVerificationEmailTemplatesRequest

	ListIdentitiesRequest(*ses.ListIdentitiesInput) ses.ListIdentitiesRequest

	ListIdentityPoliciesRequest(*ses.ListIdentityPoliciesInput) ses.ListIdentityPoliciesRequest

	ListReceiptFiltersRequest(*ses.ListReceiptFiltersInput) ses.ListReceiptFiltersRequest

	ListReceiptRuleSetsRequest(*ses.ListReceiptRuleSetsInput) ses.ListReceiptRuleSetsRequest

	ListTemplatesRequest(*ses.ListTemplatesInput) ses.ListTemplatesRequest

	ListVerifiedEmailAddressesRequest(*ses.ListVerifiedEmailAddressesInput) ses.ListVerifiedEmailAddressesRequest

	PutIdentityPolicyRequest(*ses.PutIdentityPolicyInput) ses.PutIdentityPolicyRequest

	ReorderReceiptRuleSetRequest(*ses.ReorderReceiptRuleSetInput) ses.ReorderReceiptRuleSetRequest

	SendBounceRequest(*ses.SendBounceInput) ses.SendBounceRequest

	SendBulkTemplatedEmailRequest(*ses.SendBulkTemplatedEmailInput) ses.SendBulkTemplatedEmailRequest

	SendCustomVerificationEmailRequest(*ses.SendCustomVerificationEmailInput) ses.SendCustomVerificationEmailRequest

	SendEmailRequest(*ses.SendEmailInput) ses.SendEmailRequest

	SendRawEmailRequest(*ses.SendRawEmailInput) ses.SendRawEmailRequest

	SendTemplatedEmailRequest(*ses.SendTemplatedEmailInput) ses.SendTemplatedEmailRequest

	SetActiveReceiptRuleSetRequest(*ses.SetActiveReceiptRuleSetInput) ses.SetActiveReceiptRuleSetRequest

	SetIdentityDkimEnabledRequest(*ses.SetIdentityDkimEnabledInput) ses.SetIdentityDkimEnabledRequest

	SetIdentityFeedbackForwardingEnabledRequest(*ses.SetIdentityFeedbackForwardingEnabledInput) ses.SetIdentityFeedbackForwardingEnabledRequest

	SetIdentityHeadersInNotificationsEnabledRequest(*ses.SetIdentityHeadersInNotificationsEnabledInput) ses.SetIdentityHeadersInNotificationsEnabledRequest

	SetIdentityMailFromDomainRequest(*ses.SetIdentityMailFromDomainInput) ses.SetIdentityMailFromDomainRequest

	SetIdentityNotificationTopicRequest(*ses.SetIdentityNotificationTopicInput) ses.SetIdentityNotificationTopicRequest

	SetReceiptRulePositionRequest(*ses.SetReceiptRulePositionInput) ses.SetReceiptRulePositionRequest

	TestRenderTemplateRequest(*ses.TestRenderTemplateInput) ses.TestRenderTemplateRequest

	UpdateAccountSendingEnabledRequest(*ses.UpdateAccountSendingEnabledInput) ses.UpdateAccountSendingEnabledRequest

	UpdateConfigurationSetEventDestinationRequest(*ses.UpdateConfigurationSetEventDestinationInput) ses.UpdateConfigurationSetEventDestinationRequest

	UpdateConfigurationSetReputationMetricsEnabledRequest(*ses.UpdateConfigurationSetReputationMetricsEnabledInput) ses.UpdateConfigurationSetReputationMetricsEnabledRequest

	UpdateConfigurationSetSendingEnabledRequest(*ses.UpdateConfigurationSetSendingEnabledInput) ses.UpdateConfigurationSetSendingEnabledRequest

	UpdateConfigurationSetTrackingOptionsRequest(*ses.UpdateConfigurationSetTrackingOptionsInput) ses.UpdateConfigurationSetTrackingOptionsRequest

	UpdateCustomVerificationEmailTemplateRequest(*ses.UpdateCustomVerificationEmailTemplateInput) ses.UpdateCustomVerificationEmailTemplateRequest

	UpdateReceiptRuleRequest(*ses.UpdateReceiptRuleInput) ses.UpdateReceiptRuleRequest

	UpdateTemplateRequest(*ses.UpdateTemplateInput) ses.UpdateTemplateRequest

	VerifyDomainDkimRequest(*ses.VerifyDomainDkimInput) ses.VerifyDomainDkimRequest

	VerifyDomainIdentityRequest(*ses.VerifyDomainIdentityInput) ses.VerifyDomainIdentityRequest

	VerifyEmailAddressRequest(*ses.VerifyEmailAddressInput) ses.VerifyEmailAddressRequest

	VerifyEmailIdentityRequest(*ses.VerifyEmailIdentityInput) ses.VerifyEmailIdentityRequest

	WaitUntilIdentityExists(context.Context, *ses.GetIdentityVerificationAttributesInput, ...aws.WaiterOption) error
}

var _ ClientAPI = (*ses.Client)(nil)
