package dataprotectionapi

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/dataprotection/mgmt/2021-07-01/dataprotection"
	"github.com/Azure/go-autorest/autorest"
)

// BackupVaultsClientAPI contains the set of methods on the BackupVaultsClient type.
type BackupVaultsClientAPI interface {
	CheckNameAvailability(ctx context.Context, resourceGroupName string, location string, parameters dataprotection.CheckNameAvailabilityRequest) (result dataprotection.CheckNameAvailabilityResult, err error)
	CreateOrUpdate(ctx context.Context, vaultName string, resourceGroupName string, parameters dataprotection.BackupVaultResource) (result dataprotection.BackupVaultsCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, vaultName string, resourceGroupName string) (result autorest.Response, err error)
	Get(ctx context.Context, vaultName string, resourceGroupName string) (result dataprotection.BackupVaultResource, err error)
	GetInResourceGroup(ctx context.Context, resourceGroupName string) (result dataprotection.BackupVaultResourceListPage, err error)
	GetInResourceGroupComplete(ctx context.Context, resourceGroupName string) (result dataprotection.BackupVaultResourceListIterator, err error)
	GetInSubscription(ctx context.Context) (result dataprotection.BackupVaultResourceListPage, err error)
	GetInSubscriptionComplete(ctx context.Context) (result dataprotection.BackupVaultResourceListIterator, err error)
	Update(ctx context.Context, vaultName string, resourceGroupName string, parameters dataprotection.PatchResourceRequestInput) (result dataprotection.BackupVaultsUpdateFuture, err error)
}

var _ BackupVaultsClientAPI = (*dataprotection.BackupVaultsClient)(nil)

// OperationResultClientAPI contains the set of methods on the OperationResultClient type.
type OperationResultClientAPI interface {
	Get(ctx context.Context, operationID string, location string) (result dataprotection.OperationJobExtendedInfo, err error)
}

var _ OperationResultClientAPI = (*dataprotection.OperationResultClient)(nil)

// OperationStatusClientAPI contains the set of methods on the OperationStatusClient type.
type OperationStatusClientAPI interface {
	Get(ctx context.Context, location string, operationID string) (result dataprotection.OperationResource, err error)
}

var _ OperationStatusClientAPI = (*dataprotection.OperationStatusClient)(nil)

// BackupVaultOperationResultsClientAPI contains the set of methods on the BackupVaultOperationResultsClient type.
type BackupVaultOperationResultsClientAPI interface {
	Get(ctx context.Context, vaultName string, resourceGroupName string, operationID string) (result dataprotection.BackupVaultResource, err error)
}

var _ BackupVaultOperationResultsClientAPI = (*dataprotection.BackupVaultOperationResultsClient)(nil)

// ClientAPI contains the set of methods on the Client type.
type ClientAPI interface {
	CheckFeatureSupport(ctx context.Context, location string, parameters dataprotection.BasicFeatureValidationRequestBase) (result dataprotection.FeatureValidationResponseBaseModel, err error)
}

var _ ClientAPI = (*dataprotection.Client)(nil)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	List(ctx context.Context) (result dataprotection.ClientDiscoveryResponsePage, err error)
	ListComplete(ctx context.Context) (result dataprotection.ClientDiscoveryResponseIterator, err error)
}

var _ OperationsClientAPI = (*dataprotection.OperationsClient)(nil)

// BackupPoliciesClientAPI contains the set of methods on the BackupPoliciesClient type.
type BackupPoliciesClientAPI interface {
	CreateOrUpdate(ctx context.Context, vaultName string, resourceGroupName string, backupPolicyName string, parameters dataprotection.BaseBackupPolicyResource) (result dataprotection.BaseBackupPolicyResource, err error)
	Delete(ctx context.Context, vaultName string, resourceGroupName string, backupPolicyName string) (result autorest.Response, err error)
	Get(ctx context.Context, vaultName string, resourceGroupName string, backupPolicyName string) (result dataprotection.BaseBackupPolicyResource, err error)
	List(ctx context.Context, vaultName string, resourceGroupName string) (result dataprotection.BaseBackupPolicyResourceListPage, err error)
	ListComplete(ctx context.Context, vaultName string, resourceGroupName string) (result dataprotection.BaseBackupPolicyResourceListIterator, err error)
}

var _ BackupPoliciesClientAPI = (*dataprotection.BackupPoliciesClient)(nil)

// BackupInstancesClientAPI contains the set of methods on the BackupInstancesClient type.
type BackupInstancesClientAPI interface {
	AdhocBackup(ctx context.Context, vaultName string, resourceGroupName string, backupInstanceName string, parameters dataprotection.TriggerBackupRequest) (result dataprotection.BackupInstancesAdhocBackupFuture, err error)
	CreateOrUpdate(ctx context.Context, vaultName string, resourceGroupName string, backupInstanceName string, parameters dataprotection.BackupInstanceResource) (result dataprotection.BackupInstancesCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, vaultName string, resourceGroupName string, backupInstanceName string) (result dataprotection.BackupInstancesDeleteFuture, err error)
	Get(ctx context.Context, vaultName string, resourceGroupName string, backupInstanceName string) (result dataprotection.BackupInstanceResource, err error)
	List(ctx context.Context, vaultName string, resourceGroupName string) (result dataprotection.BackupInstanceResourceListPage, err error)
	ListComplete(ctx context.Context, vaultName string, resourceGroupName string) (result dataprotection.BackupInstanceResourceListIterator, err error)
	TriggerRehydrate(ctx context.Context, resourceGroupName string, vaultName string, parameters dataprotection.AzureBackupRehydrationRequest, backupInstanceName string) (result dataprotection.BackupInstancesTriggerRehydrateFuture, err error)
	TriggerRestore(ctx context.Context, vaultName string, resourceGroupName string, backupInstanceName string, parameters dataprotection.BasicAzureBackupRestoreRequest) (result dataprotection.BackupInstancesTriggerRestoreFuture, err error)
	ValidateForBackup(ctx context.Context, vaultName string, resourceGroupName string, parameters dataprotection.ValidateForBackupRequest) (result dataprotection.BackupInstancesValidateForBackupFuture, err error)
	ValidateForRestore(ctx context.Context, vaultName string, resourceGroupName string, backupInstanceName string, parameters dataprotection.ValidateRestoreRequestObject) (result dataprotection.BackupInstancesValidateForRestoreFuture, err error)
}

var _ BackupInstancesClientAPI = (*dataprotection.BackupInstancesClient)(nil)

// RecoveryPointsClientAPI contains the set of methods on the RecoveryPointsClient type.
type RecoveryPointsClientAPI interface {
	Get(ctx context.Context, vaultName string, resourceGroupName string, backupInstanceName string, recoveryPointID string) (result dataprotection.AzureBackupRecoveryPointResource, err error)
	List(ctx context.Context, vaultName string, resourceGroupName string, backupInstanceName string, filter string, skipToken string) (result dataprotection.AzureBackupRecoveryPointResourceListPage, err error)
	ListComplete(ctx context.Context, vaultName string, resourceGroupName string, backupInstanceName string, filter string, skipToken string) (result dataprotection.AzureBackupRecoveryPointResourceListIterator, err error)
}

var _ RecoveryPointsClientAPI = (*dataprotection.RecoveryPointsClient)(nil)

// JobsClientAPI contains the set of methods on the JobsClient type.
type JobsClientAPI interface {
	Get(ctx context.Context, resourceGroupName string, vaultName string, jobID string) (result dataprotection.AzureBackupJobResource, err error)
	List(ctx context.Context, resourceGroupName string, vaultName string) (result dataprotection.AzureBackupJobResourceListPage, err error)
	ListComplete(ctx context.Context, resourceGroupName string, vaultName string) (result dataprotection.AzureBackupJobResourceListIterator, err error)
}

var _ JobsClientAPI = (*dataprotection.JobsClient)(nil)

// RestorableTimeRangesClientAPI contains the set of methods on the RestorableTimeRangesClient type.
type RestorableTimeRangesClientAPI interface {
	Find(ctx context.Context, vaultName string, resourceGroupName string, backupInstanceName string, parameters dataprotection.AzureBackupFindRestorableTimeRangesRequest) (result dataprotection.AzureBackupFindRestorableTimeRangesResponseResource, err error)
}

var _ RestorableTimeRangesClientAPI = (*dataprotection.RestorableTimeRangesClient)(nil)

// ExportJobsClientAPI contains the set of methods on the ExportJobsClient type.
type ExportJobsClientAPI interface {
	Trigger(ctx context.Context, resourceGroupName string, vaultName string) (result dataprotection.ExportJobsTriggerFuture, err error)
}

var _ ExportJobsClientAPI = (*dataprotection.ExportJobsClient)(nil)

// ExportJobsOperationResultClientAPI contains the set of methods on the ExportJobsOperationResultClient type.
type ExportJobsOperationResultClientAPI interface {
	Get(ctx context.Context, resourceGroupName string, vaultName string, operationID string) (result dataprotection.ExportJobsResult, err error)
}

var _ ExportJobsOperationResultClientAPI = (*dataprotection.ExportJobsOperationResultClient)(nil)

// ResourceGuardsClientAPI contains the set of methods on the ResourceGuardsClient type.
type ResourceGuardsClientAPI interface {
	Delete(ctx context.Context, resourceGroupName string, resourceGuardsName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, resourceGuardsName string) (result dataprotection.ResourceGuardResource, err error)
	GetBackupSecurityPINRequestsObjects(ctx context.Context, resourceGroupName string, resourceGuardsName string) (result dataprotection.DppBaseResourceListPage, err error)
	GetBackupSecurityPINRequestsObjectsComplete(ctx context.Context, resourceGroupName string, resourceGuardsName string) (result dataprotection.DppBaseResourceListIterator, err error)
	GetDefaultBackupSecurityPINRequestsObject(ctx context.Context, resourceGroupName string, resourceGuardsName string, requestName string) (result dataprotection.DppBaseResource, err error)
	GetDefaultDeleteProtectedItemRequestsObject(ctx context.Context, resourceGroupName string, resourceGuardsName string, requestName string) (result dataprotection.DppBaseResource, err error)
	GetDefaultDeleteResourceGuardProxyRequestsObject(ctx context.Context, resourceGroupName string, resourceGuardsName string, requestName string) (result dataprotection.DppBaseResource, err error)
	GetDefaultDisableSoftDeleteRequestsObject(ctx context.Context, resourceGroupName string, resourceGuardsName string, requestName string) (result dataprotection.DppBaseResource, err error)
	GetDefaultUpdateProtectedItemRequestsObject(ctx context.Context, resourceGroupName string, resourceGuardsName string, requestName string) (result dataprotection.DppBaseResource, err error)
	GetDefaultUpdateProtectionPolicyRequestsObject(ctx context.Context, resourceGroupName string, resourceGuardsName string, requestName string) (result dataprotection.DppBaseResource, err error)
	GetDeleteProtectedItemRequestsObjects(ctx context.Context, resourceGroupName string, resourceGuardsName string) (result dataprotection.DppBaseResourceListPage, err error)
	GetDeleteProtectedItemRequestsObjectsComplete(ctx context.Context, resourceGroupName string, resourceGuardsName string) (result dataprotection.DppBaseResourceListIterator, err error)
	GetDeleteResourceGuardProxyRequestsObjects(ctx context.Context, resourceGroupName string, resourceGuardsName string) (result dataprotection.DppBaseResourceListPage, err error)
	GetDeleteResourceGuardProxyRequestsObjectsComplete(ctx context.Context, resourceGroupName string, resourceGuardsName string) (result dataprotection.DppBaseResourceListIterator, err error)
	GetDisableSoftDeleteRequestsObjects(ctx context.Context, resourceGroupName string, resourceGuardsName string) (result dataprotection.DppBaseResourceListPage, err error)
	GetDisableSoftDeleteRequestsObjectsComplete(ctx context.Context, resourceGroupName string, resourceGuardsName string) (result dataprotection.DppBaseResourceListIterator, err error)
	GetResourcesInResourceGroup(ctx context.Context, resourceGroupName string) (result dataprotection.ResourceGuardResourceListPage, err error)
	GetResourcesInResourceGroupComplete(ctx context.Context, resourceGroupName string) (result dataprotection.ResourceGuardResourceListIterator, err error)
	GetResourcesInSubscription(ctx context.Context) (result dataprotection.ResourceGuardResourceListPage, err error)
	GetResourcesInSubscriptionComplete(ctx context.Context) (result dataprotection.ResourceGuardResourceListIterator, err error)
	GetUpdateProtectedItemRequestsObjects(ctx context.Context, resourceGroupName string, resourceGuardsName string) (result dataprotection.DppBaseResourceListPage, err error)
	GetUpdateProtectedItemRequestsObjectsComplete(ctx context.Context, resourceGroupName string, resourceGuardsName string) (result dataprotection.DppBaseResourceListIterator, err error)
	GetUpdateProtectionPolicyRequestsObjects(ctx context.Context, resourceGroupName string, resourceGuardsName string) (result dataprotection.DppBaseResourceListPage, err error)
	GetUpdateProtectionPolicyRequestsObjectsComplete(ctx context.Context, resourceGroupName string, resourceGuardsName string) (result dataprotection.DppBaseResourceListIterator, err error)
	Patch(ctx context.Context, resourceGroupName string, resourceGuardsName string, parameters dataprotection.PatchResourceRequestInput) (result dataprotection.ResourceGuardResource, err error)
	Put(ctx context.Context, resourceGroupName string, resourceGuardsName string, parameters dataprotection.ResourceGuardResource) (result dataprotection.ResourceGuardResource, err error)
}

var _ ResourceGuardsClientAPI = (*dataprotection.ResourceGuardsClient)(nil)
