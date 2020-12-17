package backup

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// ExportJobsOperationResultsClient is the open API 2.0 Specs for Azure RecoveryServices Backup service
type ExportJobsOperationResultsClient struct {
	BaseClient
}

// NewExportJobsOperationResultsClient creates an instance of the ExportJobsOperationResultsClient client.
func NewExportJobsOperationResultsClient(subscriptionID string) ExportJobsOperationResultsClient {
	return NewExportJobsOperationResultsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewExportJobsOperationResultsClientWithBaseURI creates an instance of the ExportJobsOperationResultsClient client
// using a custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign
// clouds, Azure stack).
func NewExportJobsOperationResultsClientWithBaseURI(baseURI string, subscriptionID string) ExportJobsOperationResultsClient {
	return ExportJobsOperationResultsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Get gets the operation result of operation triggered by Export Jobs API. If the operation is successful, then it
// also
// contains URL of a Blob and a SAS key to access the same. The blob contains exported jobs in JSON serialized format.
// Parameters:
// vaultName - the name of the recovery services vault.
// resourceGroupName - the name of the resource group where the recovery services vault is present.
// operationID - operationID which represents the export job.
func (client ExportJobsOperationResultsClient) Get(ctx context.Context, vaultName string, resourceGroupName string, operationID string) (result OperationResultInfoBaseResource, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ExportJobsOperationResultsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, vaultName, resourceGroupName, operationID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "backup.ExportJobsOperationResultsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "backup.ExportJobsOperationResultsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "backup.ExportJobsOperationResultsClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client ExportJobsOperationResultsClient) GetPreparer(ctx context.Context, vaultName string, resourceGroupName string, operationID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"operationId":       autorest.Encode("path", operationID),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"vaultName":         autorest.Encode("path", vaultName),
	}

	const APIVersion = "2019-06-15"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/backupJobs/operationResults/{operationId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ExportJobsOperationResultsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ExportJobsOperationResultsClient) GetResponder(resp *http.Response) (result OperationResultInfoBaseResource, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
