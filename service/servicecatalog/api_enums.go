// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package servicecatalog

type AccessLevelFilterKey string

// Enum values for AccessLevelFilterKey
const (
	AccessLevelFilterKeyAccount AccessLevelFilterKey = "Account"
	AccessLevelFilterKeyRole    AccessLevelFilterKey = "Role"
	AccessLevelFilterKeyUser    AccessLevelFilterKey = "User"
)

func (enum AccessLevelFilterKey) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum AccessLevelFilterKey) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type AccessStatus string

// Enum values for AccessStatus
const (
	AccessStatusEnabled     AccessStatus = "ENABLED"
	AccessStatusUnderChange AccessStatus = "UNDER_CHANGE"
	AccessStatusDisabled    AccessStatus = "DISABLED"
)

func (enum AccessStatus) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum AccessStatus) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ChangeAction string

// Enum values for ChangeAction
const (
	ChangeActionAdd    ChangeAction = "ADD"
	ChangeActionModify ChangeAction = "MODIFY"
	ChangeActionRemove ChangeAction = "REMOVE"
)

func (enum ChangeAction) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ChangeAction) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type CopyOption string

// Enum values for CopyOption
const (
	CopyOptionCopyTags CopyOption = "CopyTags"
)

func (enum CopyOption) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum CopyOption) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type CopyProductStatus string

// Enum values for CopyProductStatus
const (
	CopyProductStatusSucceeded  CopyProductStatus = "SUCCEEDED"
	CopyProductStatusInProgress CopyProductStatus = "IN_PROGRESS"
	CopyProductStatusFailed     CopyProductStatus = "FAILED"
)

func (enum CopyProductStatus) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum CopyProductStatus) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type EvaluationType string

// Enum values for EvaluationType
const (
	EvaluationTypeStatic  EvaluationType = "STATIC"
	EvaluationTypeDynamic EvaluationType = "DYNAMIC"
)

func (enum EvaluationType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum EvaluationType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type OrganizationNodeType string

// Enum values for OrganizationNodeType
const (
	OrganizationNodeTypeOrganization       OrganizationNodeType = "ORGANIZATION"
	OrganizationNodeTypeOrganizationalUnit OrganizationNodeType = "ORGANIZATIONAL_UNIT"
	OrganizationNodeTypeAccount            OrganizationNodeType = "ACCOUNT"
)

func (enum OrganizationNodeType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum OrganizationNodeType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type PortfolioShareType string

// Enum values for PortfolioShareType
const (
	PortfolioShareTypeImported          PortfolioShareType = "IMPORTED"
	PortfolioShareTypeAwsServicecatalog PortfolioShareType = "AWS_SERVICECATALOG"
	PortfolioShareTypeAwsOrganizations  PortfolioShareType = "AWS_ORGANIZATIONS"
)

func (enum PortfolioShareType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum PortfolioShareType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type PrincipalType string

// Enum values for PrincipalType
const (
	PrincipalTypeIam PrincipalType = "IAM"
)

func (enum PrincipalType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum PrincipalType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ProductSource string

// Enum values for ProductSource
const (
	ProductSourceAccount ProductSource = "ACCOUNT"
)

func (enum ProductSource) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ProductSource) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ProductType string

// Enum values for ProductType
const (
	ProductTypeCloudFormationTemplate ProductType = "CLOUD_FORMATION_TEMPLATE"
	ProductTypeMarketplace            ProductType = "MARKETPLACE"
)

func (enum ProductType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ProductType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ProductViewFilterBy string

// Enum values for ProductViewFilterBy
const (
	ProductViewFilterByFullTextSearch  ProductViewFilterBy = "FullTextSearch"
	ProductViewFilterByOwner           ProductViewFilterBy = "Owner"
	ProductViewFilterByProductType     ProductViewFilterBy = "ProductType"
	ProductViewFilterBySourceProductId ProductViewFilterBy = "SourceProductId"
)

func (enum ProductViewFilterBy) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ProductViewFilterBy) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ProductViewSortBy string

// Enum values for ProductViewSortBy
const (
	ProductViewSortByTitle        ProductViewSortBy = "Title"
	ProductViewSortByVersionCount ProductViewSortBy = "VersionCount"
	ProductViewSortByCreationDate ProductViewSortBy = "CreationDate"
)

func (enum ProductViewSortBy) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ProductViewSortBy) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ProvisionedProductPlanStatus string

// Enum values for ProvisionedProductPlanStatus
const (
	ProvisionedProductPlanStatusCreateInProgress  ProvisionedProductPlanStatus = "CREATE_IN_PROGRESS"
	ProvisionedProductPlanStatusCreateSuccess     ProvisionedProductPlanStatus = "CREATE_SUCCESS"
	ProvisionedProductPlanStatusCreateFailed      ProvisionedProductPlanStatus = "CREATE_FAILED"
	ProvisionedProductPlanStatusExecuteInProgress ProvisionedProductPlanStatus = "EXECUTE_IN_PROGRESS"
	ProvisionedProductPlanStatusExecuteSuccess    ProvisionedProductPlanStatus = "EXECUTE_SUCCESS"
	ProvisionedProductPlanStatusExecuteFailed     ProvisionedProductPlanStatus = "EXECUTE_FAILED"
)

func (enum ProvisionedProductPlanStatus) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ProvisionedProductPlanStatus) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ProvisionedProductPlanType string

// Enum values for ProvisionedProductPlanType
const (
	ProvisionedProductPlanTypeCloudformation ProvisionedProductPlanType = "CLOUDFORMATION"
)

func (enum ProvisionedProductPlanType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ProvisionedProductPlanType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ProvisionedProductStatus string

// Enum values for ProvisionedProductStatus
const (
	ProvisionedProductStatusAvailable      ProvisionedProductStatus = "AVAILABLE"
	ProvisionedProductStatusUnderChange    ProvisionedProductStatus = "UNDER_CHANGE"
	ProvisionedProductStatusTainted        ProvisionedProductStatus = "TAINTED"
	ProvisionedProductStatusError          ProvisionedProductStatus = "ERROR"
	ProvisionedProductStatusPlanInProgress ProvisionedProductStatus = "PLAN_IN_PROGRESS"
)

func (enum ProvisionedProductStatus) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ProvisionedProductStatus) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ProvisionedProductViewFilterBy string

// Enum values for ProvisionedProductViewFilterBy
const (
	ProvisionedProductViewFilterBySearchQuery ProvisionedProductViewFilterBy = "SearchQuery"
)

func (enum ProvisionedProductViewFilterBy) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ProvisionedProductViewFilterBy) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ProvisioningArtifactPropertyName string

// Enum values for ProvisioningArtifactPropertyName
const (
	ProvisioningArtifactPropertyNameId ProvisioningArtifactPropertyName = "Id"
)

func (enum ProvisioningArtifactPropertyName) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ProvisioningArtifactPropertyName) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ProvisioningArtifactType string

// Enum values for ProvisioningArtifactType
const (
	ProvisioningArtifactTypeCloudFormationTemplate ProvisioningArtifactType = "CLOUD_FORMATION_TEMPLATE"
	ProvisioningArtifactTypeMarketplaceAmi         ProvisioningArtifactType = "MARKETPLACE_AMI"
	ProvisioningArtifactTypeMarketplaceCar         ProvisioningArtifactType = "MARKETPLACE_CAR"
)

func (enum ProvisioningArtifactType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ProvisioningArtifactType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type RecordStatus string

// Enum values for RecordStatus
const (
	RecordStatusCreated           RecordStatus = "CREATED"
	RecordStatusInProgress        RecordStatus = "IN_PROGRESS"
	RecordStatusInProgressInError RecordStatus = "IN_PROGRESS_IN_ERROR"
	RecordStatusSucceeded         RecordStatus = "SUCCEEDED"
	RecordStatusFailed            RecordStatus = "FAILED"
)

func (enum RecordStatus) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum RecordStatus) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type Replacement string

// Enum values for Replacement
const (
	ReplacementTrue        Replacement = "TRUE"
	ReplacementFalse       Replacement = "FALSE"
	ReplacementConditional Replacement = "CONDITIONAL"
)

func (enum Replacement) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum Replacement) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type RequiresRecreation string

// Enum values for RequiresRecreation
const (
	RequiresRecreationNever         RequiresRecreation = "NEVER"
	RequiresRecreationConditionally RequiresRecreation = "CONDITIONALLY"
	RequiresRecreationAlways        RequiresRecreation = "ALWAYS"
)

func (enum RequiresRecreation) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum RequiresRecreation) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ResourceAttribute string

// Enum values for ResourceAttribute
const (
	ResourceAttributeProperties     ResourceAttribute = "PROPERTIES"
	ResourceAttributeMetadata       ResourceAttribute = "METADATA"
	ResourceAttributeCreationpolicy ResourceAttribute = "CREATIONPOLICY"
	ResourceAttributeUpdatepolicy   ResourceAttribute = "UPDATEPOLICY"
	ResourceAttributeDeletionpolicy ResourceAttribute = "DELETIONPOLICY"
	ResourceAttributeTags           ResourceAttribute = "TAGS"
)

func (enum ResourceAttribute) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ResourceAttribute) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ServiceActionAssociationErrorCode string

// Enum values for ServiceActionAssociationErrorCode
const (
	ServiceActionAssociationErrorCodeDuplicateResource ServiceActionAssociationErrorCode = "DUPLICATE_RESOURCE"
	ServiceActionAssociationErrorCodeInternalFailure   ServiceActionAssociationErrorCode = "INTERNAL_FAILURE"
	ServiceActionAssociationErrorCodeLimitExceeded     ServiceActionAssociationErrorCode = "LIMIT_EXCEEDED"
	ServiceActionAssociationErrorCodeResourceNotFound  ServiceActionAssociationErrorCode = "RESOURCE_NOT_FOUND"
	ServiceActionAssociationErrorCodeThrottling        ServiceActionAssociationErrorCode = "THROTTLING"
)

func (enum ServiceActionAssociationErrorCode) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ServiceActionAssociationErrorCode) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ServiceActionDefinitionKey string

// Enum values for ServiceActionDefinitionKey
const (
	ServiceActionDefinitionKeyName       ServiceActionDefinitionKey = "Name"
	ServiceActionDefinitionKeyVersion    ServiceActionDefinitionKey = "Version"
	ServiceActionDefinitionKeyAssumeRole ServiceActionDefinitionKey = "AssumeRole"
	ServiceActionDefinitionKeyParameters ServiceActionDefinitionKey = "Parameters"
)

func (enum ServiceActionDefinitionKey) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ServiceActionDefinitionKey) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ServiceActionDefinitionType string

// Enum values for ServiceActionDefinitionType
const (
	ServiceActionDefinitionTypeSsmAutomation ServiceActionDefinitionType = "SSM_AUTOMATION"
)

func (enum ServiceActionDefinitionType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ServiceActionDefinitionType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ShareStatus string

// Enum values for ShareStatus
const (
	ShareStatusNotStarted          ShareStatus = "NOT_STARTED"
	ShareStatusInProgress          ShareStatus = "IN_PROGRESS"
	ShareStatusCompleted           ShareStatus = "COMPLETED"
	ShareStatusCompletedWithErrors ShareStatus = "COMPLETED_WITH_ERRORS"
	ShareStatusError               ShareStatus = "ERROR"
)

func (enum ShareStatus) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ShareStatus) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type SortOrder string

// Enum values for SortOrder
const (
	SortOrderAscending  SortOrder = "ASCENDING"
	SortOrderDescending SortOrder = "DESCENDING"
)

func (enum SortOrder) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum SortOrder) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type StackSetOperationType string

// Enum values for StackSetOperationType
const (
	StackSetOperationTypeCreate StackSetOperationType = "CREATE"
	StackSetOperationTypeUpdate StackSetOperationType = "UPDATE"
	StackSetOperationTypeDelete StackSetOperationType = "DELETE"
)

func (enum StackSetOperationType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum StackSetOperationType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type Status string

// Enum values for Status
const (
	StatusAvailable Status = "AVAILABLE"
	StatusCreating  Status = "CREATING"
	StatusFailed    Status = "FAILED"
)

func (enum Status) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum Status) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}
