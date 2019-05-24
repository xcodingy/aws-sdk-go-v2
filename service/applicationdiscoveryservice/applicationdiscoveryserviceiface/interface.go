// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package applicationdiscoveryserviceiface provides an interface to enable mocking the AWS Application Discovery Service service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package applicationdiscoveryserviceiface

import (
	"github.com/aws/aws-sdk-go-v2/service/applicationdiscoveryservice"
)

// ClientAPI provides an interface to enable mocking the
// applicationdiscoveryservice.Client methods. This make unit testing your code that
// calls out to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // AWS Application Discovery Service.
//    func myFunc(svc applicationdiscoveryserviceiface.ClientAPI) bool {
//        // Make svc.AssociateConfigurationItemsToApplication request
//    }
//
//    func main() {
//        cfg, err := external.LoadDefaultAWSConfig()
//        if err != nil {
//            panic("failed to load config, " + err.Error())
//        }
//
//        svc := applicationdiscoveryservice.New(cfg)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockClientClient struct {
//        applicationdiscoveryserviceiface.ClientPI
//    }
//    func (m *mockClientClient) AssociateConfigurationItemsToApplication(input *applicationdiscoveryservice.AssociateConfigurationItemsToApplicationInput) (*applicationdiscoveryservice.AssociateConfigurationItemsToApplicationOutput, error) {
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
	AssociateConfigurationItemsToApplicationRequest(*applicationdiscoveryservice.AssociateConfigurationItemsToApplicationInput) applicationdiscoveryservice.AssociateConfigurationItemsToApplicationRequest

	BatchDeleteImportDataRequest(*applicationdiscoveryservice.BatchDeleteImportDataInput) applicationdiscoveryservice.BatchDeleteImportDataRequest

	CreateApplicationRequest(*applicationdiscoveryservice.CreateApplicationInput) applicationdiscoveryservice.CreateApplicationRequest

	CreateTagsRequest(*applicationdiscoveryservice.CreateTagsInput) applicationdiscoveryservice.CreateTagsRequest

	DeleteApplicationsRequest(*applicationdiscoveryservice.DeleteApplicationsInput) applicationdiscoveryservice.DeleteApplicationsRequest

	DeleteTagsRequest(*applicationdiscoveryservice.DeleteTagsInput) applicationdiscoveryservice.DeleteTagsRequest

	DescribeAgentsRequest(*applicationdiscoveryservice.DescribeAgentsInput) applicationdiscoveryservice.DescribeAgentsRequest

	DescribeConfigurationsRequest(*applicationdiscoveryservice.DescribeConfigurationsInput) applicationdiscoveryservice.DescribeConfigurationsRequest

	DescribeContinuousExportsRequest(*applicationdiscoveryservice.DescribeContinuousExportsInput) applicationdiscoveryservice.DescribeContinuousExportsRequest

	DescribeExportConfigurationsRequest(*applicationdiscoveryservice.DescribeExportConfigurationsInput) applicationdiscoveryservice.DescribeExportConfigurationsRequest

	DescribeExportTasksRequest(*applicationdiscoveryservice.DescribeExportTasksInput) applicationdiscoveryservice.DescribeExportTasksRequest

	DescribeImportTasksRequest(*applicationdiscoveryservice.DescribeImportTasksInput) applicationdiscoveryservice.DescribeImportTasksRequest

	DescribeTagsRequest(*applicationdiscoveryservice.DescribeTagsInput) applicationdiscoveryservice.DescribeTagsRequest

	DisassociateConfigurationItemsFromApplicationRequest(*applicationdiscoveryservice.DisassociateConfigurationItemsFromApplicationInput) applicationdiscoveryservice.DisassociateConfigurationItemsFromApplicationRequest

	ExportConfigurationsRequest(*applicationdiscoveryservice.ExportConfigurationsInput) applicationdiscoveryservice.ExportConfigurationsRequest

	GetDiscoverySummaryRequest(*applicationdiscoveryservice.GetDiscoverySummaryInput) applicationdiscoveryservice.GetDiscoverySummaryRequest

	ListConfigurationsRequest(*applicationdiscoveryservice.ListConfigurationsInput) applicationdiscoveryservice.ListConfigurationsRequest

	ListServerNeighborsRequest(*applicationdiscoveryservice.ListServerNeighborsInput) applicationdiscoveryservice.ListServerNeighborsRequest

	StartContinuousExportRequest(*applicationdiscoveryservice.StartContinuousExportInput) applicationdiscoveryservice.StartContinuousExportRequest

	StartDataCollectionByAgentIdsRequest(*applicationdiscoveryservice.StartDataCollectionByAgentIdsInput) applicationdiscoveryservice.StartDataCollectionByAgentIdsRequest

	StartExportTaskRequest(*applicationdiscoveryservice.StartExportTaskInput) applicationdiscoveryservice.StartExportTaskRequest

	StartImportTaskRequest(*applicationdiscoveryservice.StartImportTaskInput) applicationdiscoveryservice.StartImportTaskRequest

	StopContinuousExportRequest(*applicationdiscoveryservice.StopContinuousExportInput) applicationdiscoveryservice.StopContinuousExportRequest

	StopDataCollectionByAgentIdsRequest(*applicationdiscoveryservice.StopDataCollectionByAgentIdsInput) applicationdiscoveryservice.StopDataCollectionByAgentIdsRequest

	UpdateApplicationRequest(*applicationdiscoveryservice.UpdateApplicationInput) applicationdiscoveryservice.UpdateApplicationRequest
}

var _ ClientAPI = (*applicationdiscoveryservice.Client)(nil)
