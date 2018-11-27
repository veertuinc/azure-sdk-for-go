// +build go1.9

// Copyright 2018 Microsoft Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/tools/profileBuilder

package servicefabricmesh

import (
	"context"

	original "github.com/Azure/azure-sdk-for-go/services/preview/servicefabricmesh/mgmt/2018-09-01-preview/servicefabricmesh"
)

type ApplicationClient = original.ApplicationClient

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type BaseClient = original.BaseClient
type CodePackageClient = original.CodePackageClient
type GatewayClient = original.GatewayClient
type ApplicationScopedVolumeKind = original.ApplicationScopedVolumeKind

const (
	ServiceFabricVolumeDisk ApplicationScopedVolumeKind = original.ServiceFabricVolumeDisk
)

type AutoScalingMechanismKind = original.AutoScalingMechanismKind

const (
	AddRemoveReplica AutoScalingMechanismKind = original.AddRemoveReplica
)

type AutoScalingMetricKind = original.AutoScalingMetricKind

const (
	AutoScalingMetricKindResource AutoScalingMetricKind = original.AutoScalingMetricKindResource
)

type AutoScalingResourceMetricName = original.AutoScalingResourceMetricName

const (
	CPU        AutoScalingResourceMetricName = original.CPU
	MemoryInGB AutoScalingResourceMetricName = original.MemoryInGB
)

type AutoScalingTriggerKind = original.AutoScalingTriggerKind

const (
	AverageLoad AutoScalingTriggerKind = original.AverageLoad
)

type DiagnosticsSinkKind = original.DiagnosticsSinkKind

const (
	AzureInternalMonitoringPipeline DiagnosticsSinkKind = original.AzureInternalMonitoringPipeline
	Invalid                         DiagnosticsSinkKind = original.Invalid
)

type HeaderMatchType = original.HeaderMatchType

const (
	Exact HeaderMatchType = original.Exact
)

type HealthState = original.HealthState

const (
	HealthStateError   HealthState = original.HealthStateError
	HealthStateInvalid HealthState = original.HealthStateInvalid
	HealthStateOk      HealthState = original.HealthStateOk
	HealthStateUnknown HealthState = original.HealthStateUnknown
	HealthStateWarning HealthState = original.HealthStateWarning
)

type Kind = original.Kind

const (
	KindInlinedValue                 Kind = original.KindInlinedValue
	KindSecretResourceProperties     Kind = original.KindSecretResourceProperties
	KindSecretResourcePropertiesBase Kind = original.KindSecretResourcePropertiesBase
)

type KindBasicApplicationScopedVolumeCreationParameters = original.KindBasicApplicationScopedVolumeCreationParameters

const (
	KindApplicationScopedVolumeCreationParameters KindBasicApplicationScopedVolumeCreationParameters = original.KindApplicationScopedVolumeCreationParameters
	KindServiceFabricVolumeDisk                   KindBasicApplicationScopedVolumeCreationParameters = original.KindServiceFabricVolumeDisk
)

type KindBasicAutoScalingMechanism = original.KindBasicAutoScalingMechanism

const (
	KindAddRemoveReplica     KindBasicAutoScalingMechanism = original.KindAddRemoveReplica
	KindAutoScalingMechanism KindBasicAutoScalingMechanism = original.KindAutoScalingMechanism
)

type KindBasicAutoScalingMetric = original.KindBasicAutoScalingMetric

const (
	KindAutoScalingMetric KindBasicAutoScalingMetric = original.KindAutoScalingMetric
	KindResource          KindBasicAutoScalingMetric = original.KindResource
)

type KindBasicAutoScalingTrigger = original.KindBasicAutoScalingTrigger

const (
	KindAutoScalingTrigger KindBasicAutoScalingTrigger = original.KindAutoScalingTrigger
	KindAverageLoad        KindBasicAutoScalingTrigger = original.KindAverageLoad
)

type KindBasicDiagnosticsSinkProperties = original.KindBasicDiagnosticsSinkProperties

const (
	KindAzureInternalMonitoringPipeline KindBasicDiagnosticsSinkProperties = original.KindAzureInternalMonitoringPipeline
	KindDiagnosticsSinkProperties       KindBasicDiagnosticsSinkProperties = original.KindDiagnosticsSinkProperties
)

type KindBasicNetworkResourcePropertiesBase = original.KindBasicNetworkResourcePropertiesBase

const (
	KindLocal                         KindBasicNetworkResourcePropertiesBase = original.KindLocal
	KindNetworkResourceProperties     KindBasicNetworkResourcePropertiesBase = original.KindNetworkResourceProperties
	KindNetworkResourcePropertiesBase KindBasicNetworkResourcePropertiesBase = original.KindNetworkResourcePropertiesBase
)

type NetworkKind = original.NetworkKind

const (
	Local NetworkKind = original.Local
)

type OperatingSystemType = original.OperatingSystemType

const (
	Linux   OperatingSystemType = original.Linux
	Windows OperatingSystemType = original.Windows
)

type ResourceStatus = original.ResourceStatus

const (
	Creating  ResourceStatus = original.Creating
	Deleting  ResourceStatus = original.Deleting
	Failed    ResourceStatus = original.Failed
	Ready     ResourceStatus = original.Ready
	Unknown   ResourceStatus = original.Unknown
	Upgrading ResourceStatus = original.Upgrading
)

type SecretKind = original.SecretKind

const (
	InlinedValue SecretKind = original.InlinedValue
)

type SizeTypes = original.SizeTypes

const (
	Large  SizeTypes = original.Large
	Medium SizeTypes = original.Medium
	Small  SizeTypes = original.Small
)

type VolumeProvider = original.VolumeProvider

const (
	SFAzureFile VolumeProvider = original.SFAzureFile
)

type AddRemoveReplicaScalingMechanism = original.AddRemoveReplicaScalingMechanism
type ApplicationProperties = original.ApplicationProperties
type ApplicationResourceDescription = original.ApplicationResourceDescription
type ApplicationResourceDescriptionList = original.ApplicationResourceDescriptionList
type ApplicationResourceDescriptionListIterator = original.ApplicationResourceDescriptionListIterator
type ApplicationResourceDescriptionListPage = original.ApplicationResourceDescriptionListPage
type ApplicationResourceProperties = original.ApplicationResourceProperties
type ApplicationScopedVolume = original.ApplicationScopedVolume
type BasicApplicationScopedVolumeCreationParameters = original.BasicApplicationScopedVolumeCreationParameters
type ApplicationScopedVolumeCreationParameters = original.ApplicationScopedVolumeCreationParameters
type ApplicationScopedVolumeCreationParametersServiceFabricVolumeDisk = original.ApplicationScopedVolumeCreationParametersServiceFabricVolumeDisk
type BasicAutoScalingMechanism = original.BasicAutoScalingMechanism
type AutoScalingMechanism = original.AutoScalingMechanism
type BasicAutoScalingMetric = original.BasicAutoScalingMetric
type AutoScalingMetric = original.AutoScalingMetric
type AutoScalingPolicy = original.AutoScalingPolicy
type AutoScalingResourceMetric = original.AutoScalingResourceMetric
type BasicAutoScalingTrigger = original.BasicAutoScalingTrigger
type AutoScalingTrigger = original.AutoScalingTrigger
type AvailableOperationDisplay = original.AvailableOperationDisplay
type AverageLoadScalingTrigger = original.AverageLoadScalingTrigger
type AzureInternalMonitoringPipelineSinkDescription = original.AzureInternalMonitoringPipelineSinkDescription
type ContainerCodePackageProperties = original.ContainerCodePackageProperties
type ContainerEvent = original.ContainerEvent
type ContainerInstanceView = original.ContainerInstanceView
type ContainerLabel = original.ContainerLabel
type ContainerLogs = original.ContainerLogs
type ContainerState = original.ContainerState
type DiagnosticsDescription = original.DiagnosticsDescription
type DiagnosticsRef = original.DiagnosticsRef
type BasicDiagnosticsSinkProperties = original.BasicDiagnosticsSinkProperties
type DiagnosticsSinkProperties = original.DiagnosticsSinkProperties
type EndpointProperties = original.EndpointProperties
type EndpointRef = original.EndpointRef
type EnvironmentVariable = original.EnvironmentVariable
type ErrorDetailsModel = original.ErrorDetailsModel
type ErrorErrorModel = original.ErrorErrorModel
type ErrorModel = original.ErrorModel
type GatewayDestination = original.GatewayDestination
type GatewayProperties = original.GatewayProperties
type GatewayResourceDescription = original.GatewayResourceDescription
type GatewayResourceDescriptionList = original.GatewayResourceDescriptionList
type GatewayResourceDescriptionListIterator = original.GatewayResourceDescriptionListIterator
type GatewayResourceDescriptionListPage = original.GatewayResourceDescriptionListPage
type GatewayResourceProperties = original.GatewayResourceProperties
type HTTPConfig = original.HTTPConfig
type HTTPHostConfig = original.HTTPHostConfig
type HTTPRouteConfig = original.HTTPRouteConfig
type HTTPRouteMatchHeader = original.HTTPRouteMatchHeader
type HTTPRouteMatchPath = original.HTTPRouteMatchPath
type HTTPRouteMatchRule = original.HTTPRouteMatchRule
type ImageRegistryCredential = original.ImageRegistryCredential
type InlinedValueSecretResourceProperties = original.InlinedValueSecretResourceProperties
type LocalNetworkResourceProperties = original.LocalNetworkResourceProperties
type ManagedProxyResource = original.ManagedProxyResource
type NetworkRef = original.NetworkRef
type NetworkResourceDescription = original.NetworkResourceDescription
type NetworkResourceDescriptionList = original.NetworkResourceDescriptionList
type NetworkResourceDescriptionListIterator = original.NetworkResourceDescriptionListIterator
type NetworkResourceDescriptionListPage = original.NetworkResourceDescriptionListPage
type BasicNetworkResourceProperties = original.BasicNetworkResourceProperties
type NetworkResourceProperties = original.NetworkResourceProperties
type BasicNetworkResourcePropertiesBase = original.BasicNetworkResourcePropertiesBase
type NetworkResourcePropertiesBase = original.NetworkResourcePropertiesBase
type OperationListResult = original.OperationListResult
type OperationListResultIterator = original.OperationListResultIterator
type OperationListResultPage = original.OperationListResultPage
type OperationResult = original.OperationResult
type ProvisionedResourceProperties = original.ProvisionedResourceProperties
type ProxyResource = original.ProxyResource
type ReliableCollectionsRef = original.ReliableCollectionsRef
type Resource = original.Resource
type ResourceLimits = original.ResourceLimits
type ResourceRequests = original.ResourceRequests
type ResourceRequirements = original.ResourceRequirements
type SecretResourceDescription = original.SecretResourceDescription
type SecretResourceDescriptionList = original.SecretResourceDescriptionList
type SecretResourceDescriptionListIterator = original.SecretResourceDescriptionListIterator
type SecretResourceDescriptionListPage = original.SecretResourceDescriptionListPage
type BasicSecretResourceProperties = original.BasicSecretResourceProperties
type SecretResourceProperties = original.SecretResourceProperties
type BasicSecretResourcePropertiesBase = original.BasicSecretResourcePropertiesBase
type SecretResourcePropertiesBase = original.SecretResourcePropertiesBase
type SecretValue = original.SecretValue
type SecretValueProperties = original.SecretValueProperties
type SecretValueResourceDescription = original.SecretValueResourceDescription
type SecretValueResourceDescriptionList = original.SecretValueResourceDescriptionList
type SecretValueResourceDescriptionListIterator = original.SecretValueResourceDescriptionListIterator
type SecretValueResourceDescriptionListPage = original.SecretValueResourceDescriptionListPage
type SecretValueResourceProperties = original.SecretValueResourceProperties
type ServiceProperties = original.ServiceProperties
type ServiceReplicaDescription = original.ServiceReplicaDescription
type ServiceReplicaDescriptionList = original.ServiceReplicaDescriptionList
type ServiceReplicaDescriptionListIterator = original.ServiceReplicaDescriptionListIterator
type ServiceReplicaDescriptionListPage = original.ServiceReplicaDescriptionListPage
type ServiceReplicaProperties = original.ServiceReplicaProperties
type ServiceResourceDescription = original.ServiceResourceDescription
type ServiceResourceDescriptionList = original.ServiceResourceDescriptionList
type ServiceResourceDescriptionListIterator = original.ServiceResourceDescriptionListIterator
type ServiceResourceDescriptionListPage = original.ServiceResourceDescriptionListPage
type ServiceResourceProperties = original.ServiceResourceProperties
type Setting = original.Setting
type TCPConfig = original.TCPConfig
type TrackedResource = original.TrackedResource
type VolumeProperties = original.VolumeProperties
type VolumeProviderParametersAzureFile = original.VolumeProviderParametersAzureFile
type VolumeReference = original.VolumeReference
type VolumeResourceDescription = original.VolumeResourceDescription
type VolumeResourceDescriptionList = original.VolumeResourceDescriptionList
type VolumeResourceDescriptionListIterator = original.VolumeResourceDescriptionListIterator
type VolumeResourceDescriptionListPage = original.VolumeResourceDescriptionListPage
type VolumeResourceProperties = original.VolumeResourceProperties
type NetworkClient = original.NetworkClient
type OperationsClient = original.OperationsClient
type SecretClient = original.SecretClient
type SecretValueClient = original.SecretValueClient
type ServiceClient = original.ServiceClient
type ServiceReplicaClient = original.ServiceReplicaClient
type VolumeClient = original.VolumeClient

func NewApplicationClient(subscriptionID string) ApplicationClient {
	return original.NewApplicationClient(subscriptionID)
}
func NewApplicationClientWithBaseURI(baseURI string, subscriptionID string) ApplicationClient {
	return original.NewApplicationClientWithBaseURI(baseURI, subscriptionID)
}
func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func NewCodePackageClient(subscriptionID string) CodePackageClient {
	return original.NewCodePackageClient(subscriptionID)
}
func NewCodePackageClientWithBaseURI(baseURI string, subscriptionID string) CodePackageClient {
	return original.NewCodePackageClientWithBaseURI(baseURI, subscriptionID)
}
func NewGatewayClient(subscriptionID string) GatewayClient {
	return original.NewGatewayClient(subscriptionID)
}
func NewGatewayClientWithBaseURI(baseURI string, subscriptionID string) GatewayClient {
	return original.NewGatewayClientWithBaseURI(baseURI, subscriptionID)
}
func PossibleApplicationScopedVolumeKindValues() []ApplicationScopedVolumeKind {
	return original.PossibleApplicationScopedVolumeKindValues()
}
func PossibleAutoScalingMechanismKindValues() []AutoScalingMechanismKind {
	return original.PossibleAutoScalingMechanismKindValues()
}
func PossibleAutoScalingMetricKindValues() []AutoScalingMetricKind {
	return original.PossibleAutoScalingMetricKindValues()
}
func PossibleAutoScalingResourceMetricNameValues() []AutoScalingResourceMetricName {
	return original.PossibleAutoScalingResourceMetricNameValues()
}
func PossibleAutoScalingTriggerKindValues() []AutoScalingTriggerKind {
	return original.PossibleAutoScalingTriggerKindValues()
}
func PossibleDiagnosticsSinkKindValues() []DiagnosticsSinkKind {
	return original.PossibleDiagnosticsSinkKindValues()
}
func PossibleHeaderMatchTypeValues() []HeaderMatchType {
	return original.PossibleHeaderMatchTypeValues()
}
func PossibleHealthStateValues() []HealthState {
	return original.PossibleHealthStateValues()
}
func PossibleKindValues() []Kind {
	return original.PossibleKindValues()
}
func PossibleKindBasicApplicationScopedVolumeCreationParametersValues() []KindBasicApplicationScopedVolumeCreationParameters {
	return original.PossibleKindBasicApplicationScopedVolumeCreationParametersValues()
}
func PossibleKindBasicAutoScalingMechanismValues() []KindBasicAutoScalingMechanism {
	return original.PossibleKindBasicAutoScalingMechanismValues()
}
func PossibleKindBasicAutoScalingMetricValues() []KindBasicAutoScalingMetric {
	return original.PossibleKindBasicAutoScalingMetricValues()
}
func PossibleKindBasicAutoScalingTriggerValues() []KindBasicAutoScalingTrigger {
	return original.PossibleKindBasicAutoScalingTriggerValues()
}
func PossibleKindBasicDiagnosticsSinkPropertiesValues() []KindBasicDiagnosticsSinkProperties {
	return original.PossibleKindBasicDiagnosticsSinkPropertiesValues()
}
func PossibleKindBasicNetworkResourcePropertiesBaseValues() []KindBasicNetworkResourcePropertiesBase {
	return original.PossibleKindBasicNetworkResourcePropertiesBaseValues()
}
func PossibleNetworkKindValues() []NetworkKind {
	return original.PossibleNetworkKindValues()
}
func PossibleOperatingSystemTypeValues() []OperatingSystemType {
	return original.PossibleOperatingSystemTypeValues()
}
func PossibleResourceStatusValues() []ResourceStatus {
	return original.PossibleResourceStatusValues()
}
func PossibleSecretKindValues() []SecretKind {
	return original.PossibleSecretKindValues()
}
func PossibleSizeTypesValues() []SizeTypes {
	return original.PossibleSizeTypesValues()
}
func PossibleVolumeProviderValues() []VolumeProvider {
	return original.PossibleVolumeProviderValues()
}
func NewApplicationResourceDescriptionListIterator(page ApplicationResourceDescriptionListPage) ApplicationResourceDescriptionListIterator {
	return original.NewApplicationResourceDescriptionListIterator(page)
}
func NewApplicationResourceDescriptionListPage(getNextPage func(context.Context, ApplicationResourceDescriptionList) (ApplicationResourceDescriptionList, error)) ApplicationResourceDescriptionListPage {
	return original.NewApplicationResourceDescriptionListPage(getNextPage)
}
func NewGatewayResourceDescriptionListIterator(page GatewayResourceDescriptionListPage) GatewayResourceDescriptionListIterator {
	return original.NewGatewayResourceDescriptionListIterator(page)
}
func NewGatewayResourceDescriptionListPage(getNextPage func(context.Context, GatewayResourceDescriptionList) (GatewayResourceDescriptionList, error)) GatewayResourceDescriptionListPage {
	return original.NewGatewayResourceDescriptionListPage(getNextPage)
}
func NewNetworkResourceDescriptionListIterator(page NetworkResourceDescriptionListPage) NetworkResourceDescriptionListIterator {
	return original.NewNetworkResourceDescriptionListIterator(page)
}
func NewNetworkResourceDescriptionListPage(getNextPage func(context.Context, NetworkResourceDescriptionList) (NetworkResourceDescriptionList, error)) NetworkResourceDescriptionListPage {
	return original.NewNetworkResourceDescriptionListPage(getNextPage)
}
func NewOperationListResultIterator(page OperationListResultPage) OperationListResultIterator {
	return original.NewOperationListResultIterator(page)
}
func NewOperationListResultPage(getNextPage func(context.Context, OperationListResult) (OperationListResult, error)) OperationListResultPage {
	return original.NewOperationListResultPage(getNextPage)
}
func NewSecretResourceDescriptionListIterator(page SecretResourceDescriptionListPage) SecretResourceDescriptionListIterator {
	return original.NewSecretResourceDescriptionListIterator(page)
}
func NewSecretResourceDescriptionListPage(getNextPage func(context.Context, SecretResourceDescriptionList) (SecretResourceDescriptionList, error)) SecretResourceDescriptionListPage {
	return original.NewSecretResourceDescriptionListPage(getNextPage)
}
func NewSecretValueResourceDescriptionListIterator(page SecretValueResourceDescriptionListPage) SecretValueResourceDescriptionListIterator {
	return original.NewSecretValueResourceDescriptionListIterator(page)
}
func NewSecretValueResourceDescriptionListPage(getNextPage func(context.Context, SecretValueResourceDescriptionList) (SecretValueResourceDescriptionList, error)) SecretValueResourceDescriptionListPage {
	return original.NewSecretValueResourceDescriptionListPage(getNextPage)
}
func NewServiceReplicaDescriptionListIterator(page ServiceReplicaDescriptionListPage) ServiceReplicaDescriptionListIterator {
	return original.NewServiceReplicaDescriptionListIterator(page)
}
func NewServiceReplicaDescriptionListPage(getNextPage func(context.Context, ServiceReplicaDescriptionList) (ServiceReplicaDescriptionList, error)) ServiceReplicaDescriptionListPage {
	return original.NewServiceReplicaDescriptionListPage(getNextPage)
}
func NewServiceResourceDescriptionListIterator(page ServiceResourceDescriptionListPage) ServiceResourceDescriptionListIterator {
	return original.NewServiceResourceDescriptionListIterator(page)
}
func NewServiceResourceDescriptionListPage(getNextPage func(context.Context, ServiceResourceDescriptionList) (ServiceResourceDescriptionList, error)) ServiceResourceDescriptionListPage {
	return original.NewServiceResourceDescriptionListPage(getNextPage)
}
func NewVolumeResourceDescriptionListIterator(page VolumeResourceDescriptionListPage) VolumeResourceDescriptionListIterator {
	return original.NewVolumeResourceDescriptionListIterator(page)
}
func NewVolumeResourceDescriptionListPage(getNextPage func(context.Context, VolumeResourceDescriptionList) (VolumeResourceDescriptionList, error)) VolumeResourceDescriptionListPage {
	return original.NewVolumeResourceDescriptionListPage(getNextPage)
}
func NewNetworkClient(subscriptionID string) NetworkClient {
	return original.NewNetworkClient(subscriptionID)
}
func NewNetworkClientWithBaseURI(baseURI string, subscriptionID string) NetworkClient {
	return original.NewNetworkClientWithBaseURI(baseURI, subscriptionID)
}
func NewOperationsClient(subscriptionID string) OperationsClient {
	return original.NewOperationsClient(subscriptionID)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewSecretClient(subscriptionID string) SecretClient {
	return original.NewSecretClient(subscriptionID)
}
func NewSecretClientWithBaseURI(baseURI string, subscriptionID string) SecretClient {
	return original.NewSecretClientWithBaseURI(baseURI, subscriptionID)
}
func NewSecretValueClient(subscriptionID string) SecretValueClient {
	return original.NewSecretValueClient(subscriptionID)
}
func NewSecretValueClientWithBaseURI(baseURI string, subscriptionID string) SecretValueClient {
	return original.NewSecretValueClientWithBaseURI(baseURI, subscriptionID)
}
func NewServiceClient(subscriptionID string) ServiceClient {
	return original.NewServiceClient(subscriptionID)
}
func NewServiceClientWithBaseURI(baseURI string, subscriptionID string) ServiceClient {
	return original.NewServiceClientWithBaseURI(baseURI, subscriptionID)
}
func NewServiceReplicaClient(subscriptionID string) ServiceReplicaClient {
	return original.NewServiceReplicaClient(subscriptionID)
}
func NewServiceReplicaClientWithBaseURI(baseURI string, subscriptionID string) ServiceReplicaClient {
	return original.NewServiceReplicaClientWithBaseURI(baseURI, subscriptionID)
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
func NewVolumeClient(subscriptionID string) VolumeClient {
	return original.NewVolumeClient(subscriptionID)
}
func NewVolumeClientWithBaseURI(baseURI string, subscriptionID string) VolumeClient {
	return original.NewVolumeClientWithBaseURI(baseURI, subscriptionID)
}
