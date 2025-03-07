//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armpostgresql

// CheckNameAvailabilityClientExecuteResponse contains the response from method CheckNameAvailabilityClient.Execute.
type CheckNameAvailabilityClientExecuteResponse struct {
	NameAvailability
}

// ConfigurationsClientCreateOrUpdateResponse contains the response from method ConfigurationsClient.CreateOrUpdate.
type ConfigurationsClientCreateOrUpdateResponse struct {
	Configuration
}

// ConfigurationsClientGetResponse contains the response from method ConfigurationsClient.Get.
type ConfigurationsClientGetResponse struct {
	Configuration
}

// ConfigurationsClientListByServerResponse contains the response from method ConfigurationsClient.ListByServer.
type ConfigurationsClientListByServerResponse struct {
	ConfigurationListResult
}

// DatabasesClientCreateOrUpdateResponse contains the response from method DatabasesClient.CreateOrUpdate.
type DatabasesClientCreateOrUpdateResponse struct {
	Database
}

// DatabasesClientDeleteResponse contains the response from method DatabasesClient.Delete.
type DatabasesClientDeleteResponse struct {
	// placeholder for future response values
}

// DatabasesClientGetResponse contains the response from method DatabasesClient.Get.
type DatabasesClientGetResponse struct {
	Database
}

// DatabasesClientListByServerResponse contains the response from method DatabasesClient.ListByServer.
type DatabasesClientListByServerResponse struct {
	DatabaseListResult
}

// FirewallRulesClientCreateOrUpdateResponse contains the response from method FirewallRulesClient.CreateOrUpdate.
type FirewallRulesClientCreateOrUpdateResponse struct {
	FirewallRule
}

// FirewallRulesClientDeleteResponse contains the response from method FirewallRulesClient.Delete.
type FirewallRulesClientDeleteResponse struct {
	// placeholder for future response values
}

// FirewallRulesClientGetResponse contains the response from method FirewallRulesClient.Get.
type FirewallRulesClientGetResponse struct {
	FirewallRule
}

// FirewallRulesClientListByServerResponse contains the response from method FirewallRulesClient.ListByServer.
type FirewallRulesClientListByServerResponse struct {
	FirewallRuleListResult
}

// LocationBasedPerformanceTierClientListResponse contains the response from method LocationBasedPerformanceTierClient.List.
type LocationBasedPerformanceTierClientListResponse struct {
	PerformanceTierListResult
}

// LogFilesClientListByServerResponse contains the response from method LogFilesClient.ListByServer.
type LogFilesClientListByServerResponse struct {
	LogFileListResult
}

// OperationsClientListResponse contains the response from method OperationsClient.List.
type OperationsClientListResponse struct {
	OperationListResult
}

// PrivateEndpointConnectionsClientCreateOrUpdateResponse contains the response from method PrivateEndpointConnectionsClient.CreateOrUpdate.
type PrivateEndpointConnectionsClientCreateOrUpdateResponse struct {
	PrivateEndpointConnection
}

// PrivateEndpointConnectionsClientDeleteResponse contains the response from method PrivateEndpointConnectionsClient.Delete.
type PrivateEndpointConnectionsClientDeleteResponse struct {
	// placeholder for future response values
}

// PrivateEndpointConnectionsClientGetResponse contains the response from method PrivateEndpointConnectionsClient.Get.
type PrivateEndpointConnectionsClientGetResponse struct {
	PrivateEndpointConnection
}

// PrivateEndpointConnectionsClientListByServerResponse contains the response from method PrivateEndpointConnectionsClient.ListByServer.
type PrivateEndpointConnectionsClientListByServerResponse struct {
	PrivateEndpointConnectionListResult
}

// PrivateEndpointConnectionsClientUpdateTagsResponse contains the response from method PrivateEndpointConnectionsClient.UpdateTags.
type PrivateEndpointConnectionsClientUpdateTagsResponse struct {
	PrivateEndpointConnection
}

// PrivateLinkResourcesClientGetResponse contains the response from method PrivateLinkResourcesClient.Get.
type PrivateLinkResourcesClientGetResponse struct {
	PrivateLinkResource
}

// PrivateLinkResourcesClientListByServerResponse contains the response from method PrivateLinkResourcesClient.ListByServer.
type PrivateLinkResourcesClientListByServerResponse struct {
	PrivateLinkResourceListResult
}

// RecoverableServersClientGetResponse contains the response from method RecoverableServersClient.Get.
type RecoverableServersClientGetResponse struct {
	RecoverableServerResource
}

// ReplicasClientListByServerResponse contains the response from method ReplicasClient.ListByServer.
type ReplicasClientListByServerResponse struct {
	ServerListResult
}

// ServerAdministratorsClientCreateOrUpdateResponse contains the response from method ServerAdministratorsClient.CreateOrUpdate.
type ServerAdministratorsClientCreateOrUpdateResponse struct {
	ServerAdministratorResource
}

// ServerAdministratorsClientDeleteResponse contains the response from method ServerAdministratorsClient.Delete.
type ServerAdministratorsClientDeleteResponse struct {
	// placeholder for future response values
}

// ServerAdministratorsClientGetResponse contains the response from method ServerAdministratorsClient.Get.
type ServerAdministratorsClientGetResponse struct {
	ServerAdministratorResource
}

// ServerAdministratorsClientListResponse contains the response from method ServerAdministratorsClient.List.
type ServerAdministratorsClientListResponse struct {
	ServerAdministratorResourceListResult
}

// ServerBasedPerformanceTierClientListResponse contains the response from method ServerBasedPerformanceTierClient.List.
type ServerBasedPerformanceTierClientListResponse struct {
	PerformanceTierListResult
}

// ServerKeysClientCreateOrUpdateResponse contains the response from method ServerKeysClient.CreateOrUpdate.
type ServerKeysClientCreateOrUpdateResponse struct {
	ServerKey
}

// ServerKeysClientDeleteResponse contains the response from method ServerKeysClient.Delete.
type ServerKeysClientDeleteResponse struct {
	// placeholder for future response values
}

// ServerKeysClientGetResponse contains the response from method ServerKeysClient.Get.
type ServerKeysClientGetResponse struct {
	ServerKey
}

// ServerKeysClientListResponse contains the response from method ServerKeysClient.List.
type ServerKeysClientListResponse struct {
	ServerKeyListResult
}

// ServerParametersClientListUpdateConfigurationsResponse contains the response from method ServerParametersClient.ListUpdateConfigurations.
type ServerParametersClientListUpdateConfigurationsResponse struct {
	ConfigurationListResult
}

// ServerSecurityAlertPoliciesClientCreateOrUpdateResponse contains the response from method ServerSecurityAlertPoliciesClient.CreateOrUpdate.
type ServerSecurityAlertPoliciesClientCreateOrUpdateResponse struct {
	ServerSecurityAlertPolicy
}

// ServerSecurityAlertPoliciesClientGetResponse contains the response from method ServerSecurityAlertPoliciesClient.Get.
type ServerSecurityAlertPoliciesClientGetResponse struct {
	ServerSecurityAlertPolicy
}

// ServerSecurityAlertPoliciesClientListByServerResponse contains the response from method ServerSecurityAlertPoliciesClient.ListByServer.
type ServerSecurityAlertPoliciesClientListByServerResponse struct {
	ServerSecurityAlertPolicyListResult
}

// ServersClientCreateResponse contains the response from method ServersClient.Create.
type ServersClientCreateResponse struct {
	Server
}

// ServersClientDeleteResponse contains the response from method ServersClient.Delete.
type ServersClientDeleteResponse struct {
	// placeholder for future response values
}

// ServersClientGetResponse contains the response from method ServersClient.Get.
type ServersClientGetResponse struct {
	Server
}

// ServersClientListByResourceGroupResponse contains the response from method ServersClient.ListByResourceGroup.
type ServersClientListByResourceGroupResponse struct {
	ServerListResult
}

// ServersClientListResponse contains the response from method ServersClient.List.
type ServersClientListResponse struct {
	ServerListResult
}

// ServersClientRestartResponse contains the response from method ServersClient.Restart.
type ServersClientRestartResponse struct {
	// placeholder for future response values
}

// ServersClientUpdateResponse contains the response from method ServersClient.Update.
type ServersClientUpdateResponse struct {
	Server
}

// VirtualNetworkRulesClientCreateOrUpdateResponse contains the response from method VirtualNetworkRulesClient.CreateOrUpdate.
type VirtualNetworkRulesClientCreateOrUpdateResponse struct {
	VirtualNetworkRule
}

// VirtualNetworkRulesClientDeleteResponse contains the response from method VirtualNetworkRulesClient.Delete.
type VirtualNetworkRulesClientDeleteResponse struct {
	// placeholder for future response values
}

// VirtualNetworkRulesClientGetResponse contains the response from method VirtualNetworkRulesClient.Get.
type VirtualNetworkRulesClientGetResponse struct {
	VirtualNetworkRule
}

// VirtualNetworkRulesClientListByServerResponse contains the response from method VirtualNetworkRulesClient.ListByServer.
type VirtualNetworkRulesClientListByServerResponse struct {
	VirtualNetworkRuleListResult
}
