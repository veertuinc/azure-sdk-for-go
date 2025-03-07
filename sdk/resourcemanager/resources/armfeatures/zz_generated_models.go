//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armfeatures

import "time"

// AuthorizationProfile - Authorization Profile
type AuthorizationProfile struct {
	// READ-ONLY; The approved time
	ApprovedTime *time.Time `json:"approvedTime,omitempty" azure:"ro"`

	// READ-ONLY; The approver
	Approver *string `json:"approver,omitempty" azure:"ro"`

	// READ-ONLY; The requested time
	RequestedTime *time.Time `json:"requestedTime,omitempty" azure:"ro"`

	// READ-ONLY; The requester
	Requester *string `json:"requester,omitempty" azure:"ro"`

	// READ-ONLY; The requester object id
	RequesterObjectID *string `json:"requesterObjectId,omitempty" azure:"ro"`
}

// ClientGetOptions contains the optional parameters for the Client.Get method.
type ClientGetOptions struct {
	// placeholder for future optional parameters
}

// ClientListAllOptions contains the optional parameters for the Client.ListAll method.
type ClientListAllOptions struct {
	// placeholder for future optional parameters
}

// ClientListOptions contains the optional parameters for the Client.List method.
type ClientListOptions struct {
	// placeholder for future optional parameters
}

// ClientRegisterOptions contains the optional parameters for the Client.Register method.
type ClientRegisterOptions struct {
	// placeholder for future optional parameters
}

// ClientUnregisterOptions contains the optional parameters for the Client.Unregister method.
type ClientUnregisterOptions struct {
	// placeholder for future optional parameters
}

// ErrorDefinition - Error definition.
type ErrorDefinition struct {
	// Internal error details.
	Details []*ErrorDefinition `json:"details,omitempty"`

	// READ-ONLY; Service specific error code which serves as the substatus for the HTTP error code.
	Code *string `json:"code,omitempty" azure:"ro"`

	// READ-ONLY; Description of the error.
	Message *string `json:"message,omitempty" azure:"ro"`
}

// ErrorResponse - Error response indicates that the service is not able to process the incoming request.
type ErrorResponse struct {
	// The error details.
	Error *ErrorDefinition `json:"error,omitempty"`
}

// FeatureClientListOperationsOptions contains the optional parameters for the FeatureClient.ListOperations method.
type FeatureClientListOperationsOptions struct {
	// placeholder for future optional parameters
}

// FeatureOperationsListResult - List of previewed features.
type FeatureOperationsListResult struct {
	// The URL to use for getting the next set of results.
	NextLink *string `json:"nextLink,omitempty"`

	// The array of features.
	Value []*FeatureResult `json:"value,omitempty"`
}

// FeatureProperties - Information about feature.
type FeatureProperties struct {
	// The registration state of the feature for the subscription.
	State *string `json:"state,omitempty"`
}

// FeatureResult - Previewed feature information.
type FeatureResult struct {
	// The resource ID of the feature.
	ID *string `json:"id,omitempty"`

	// The name of the feature.
	Name *string `json:"name,omitempty"`

	// Properties of the previewed feature.
	Properties *FeatureProperties `json:"properties,omitempty"`

	// The resource type of the feature.
	Type *string `json:"type,omitempty"`
}

// Operation - Microsoft.Features operation
type Operation struct {
	// The object that represents the operation.
	Display *OperationDisplay `json:"display,omitempty"`

	// Operation name: {provider}/{resource}/{operation}
	Name *string `json:"name,omitempty"`
}

// OperationDisplay - The object that represents the operation.
type OperationDisplay struct {
	// Operation type: Read, write, delete, etc.
	Operation *string `json:"operation,omitempty"`

	// Service provider: Microsoft.Features
	Provider *string `json:"provider,omitempty"`

	// Resource on which the operation is performed: Profile, endpoint, etc.
	Resource *string `json:"resource,omitempty"`
}

// OperationListResult - Result of the request to list Microsoft.Features operations. It contains a list of operations and
// a URL link to get the next set of results.
type OperationListResult struct {
	// URL to get the next set of operation list results if there are any.
	NextLink *string `json:"nextLink,omitempty"`

	// List of Microsoft.Features operations.
	Value []*Operation `json:"value,omitempty"`
}

// ProxyResource - An Azure proxy resource.
type ProxyResource struct {
	// READ-ONLY; Azure resource Id.
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; Azure resource name.
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; Azure resource type.
	Type *string `json:"type,omitempty" azure:"ro"`
}

// SubscriptionFeatureRegistration - Subscription feature registration details
type SubscriptionFeatureRegistration struct {
	Properties *SubscriptionFeatureRegistrationProperties `json:"properties,omitempty"`

	// READ-ONLY; Azure resource Id.
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; Azure resource name.
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; Azure resource type.
	Type *string `json:"type,omitempty" azure:"ro"`
}

// SubscriptionFeatureRegistrationList - The list of subscription feature registrations.
type SubscriptionFeatureRegistrationList struct {
	// The link used to get the next page of subscription feature registrations list.
	NextLink *string `json:"nextLink,omitempty"`

	// The list of subscription feature registrations.
	Value []*SubscriptionFeatureRegistration `json:"value,omitempty"`
}

type SubscriptionFeatureRegistrationProperties struct {
	// Authorization Profile
	AuthorizationProfile *AuthorizationProfile `json:"authorizationProfile,omitempty"`

	// The feature description.
	Description *string `json:"description,omitempty"`

	// Key-value pairs for meta data.
	Metadata map[string]*string `json:"metadata,omitempty"`

	// Indicates whether feature should be displayed in Portal.
	ShouldFeatureDisplayInPortal *bool `json:"shouldFeatureDisplayInPortal,omitempty"`

	// The state.
	State *SubscriptionFeatureRegistrationState `json:"state,omitempty"`

	// READ-ONLY; The feature approval type.
	ApprovalType *SubscriptionFeatureRegistrationApprovalType `json:"approvalType,omitempty" azure:"ro"`

	// READ-ONLY; The featureDisplayName.
	DisplayName *string `json:"displayName,omitempty" azure:"ro"`

	// READ-ONLY; The feature documentation link.
	DocumentationLink *string `json:"documentationLink,omitempty" azure:"ro"`

	// READ-ONLY; The featureName.
	FeatureName *string `json:"featureName,omitempty" azure:"ro"`

	// READ-ONLY; The providerNamespace.
	ProviderNamespace *string `json:"providerNamespace,omitempty" azure:"ro"`

	// READ-ONLY; The feature registration date.
	RegistrationDate *time.Time `json:"registrationDate,omitempty" azure:"ro"`

	// READ-ONLY; The feature release date.
	ReleaseDate *time.Time `json:"releaseDate,omitempty" azure:"ro"`

	// READ-ONLY; The subscriptionId.
	SubscriptionID *string `json:"subscriptionId,omitempty" azure:"ro"`

	// READ-ONLY; The tenantId.
	TenantID *string `json:"tenantId,omitempty" azure:"ro"`
}

// SubscriptionFeatureRegistrationsClientCreateOrUpdateOptions contains the optional parameters for the SubscriptionFeatureRegistrationsClient.CreateOrUpdate
// method.
type SubscriptionFeatureRegistrationsClientCreateOrUpdateOptions struct {
	// Subscription Feature Registration Type details.
	SubscriptionFeatureRegistrationType *SubscriptionFeatureRegistration
}

// SubscriptionFeatureRegistrationsClientDeleteOptions contains the optional parameters for the SubscriptionFeatureRegistrationsClient.Delete
// method.
type SubscriptionFeatureRegistrationsClientDeleteOptions struct {
	// placeholder for future optional parameters
}

// SubscriptionFeatureRegistrationsClientGetOptions contains the optional parameters for the SubscriptionFeatureRegistrationsClient.Get
// method.
type SubscriptionFeatureRegistrationsClientGetOptions struct {
	// placeholder for future optional parameters
}

// SubscriptionFeatureRegistrationsClientListAllBySubscriptionOptions contains the optional parameters for the SubscriptionFeatureRegistrationsClient.ListAllBySubscription
// method.
type SubscriptionFeatureRegistrationsClientListAllBySubscriptionOptions struct {
	// placeholder for future optional parameters
}

// SubscriptionFeatureRegistrationsClientListBySubscriptionOptions contains the optional parameters for the SubscriptionFeatureRegistrationsClient.ListBySubscription
// method.
type SubscriptionFeatureRegistrationsClientListBySubscriptionOptions struct {
	// placeholder for future optional parameters
}
