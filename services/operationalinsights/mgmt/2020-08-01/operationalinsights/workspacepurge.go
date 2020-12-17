package operationalinsights

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
	"github.com/Azure/go-autorest/autorest/validation"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// WorkspacePurgeClient is the operational Insights Client
type WorkspacePurgeClient struct {
	BaseClient
}

// NewWorkspacePurgeClient creates an instance of the WorkspacePurgeClient client.
func NewWorkspacePurgeClient(subscriptionID string) WorkspacePurgeClient {
	return NewWorkspacePurgeClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewWorkspacePurgeClientWithBaseURI creates an instance of the WorkspacePurgeClient client using a custom endpoint.
// Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewWorkspacePurgeClientWithBaseURI(baseURI string, subscriptionID string) WorkspacePurgeClient {
	return WorkspacePurgeClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// GetPurgeStatus gets status of an ongoing purge operation.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// workspaceName - the name of the workspace.
// purgeID - in a purge status request, this is the Id of the operation the status of which is returned.
func (client WorkspacePurgeClient) GetPurgeStatus(ctx context.Context, resourceGroupName string, workspaceName string, purgeID string) (result WorkspacePurgeStatusResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/WorkspacePurgeClient.GetPurgeStatus")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}},
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: workspaceName,
			Constraints: []validation.Constraint{{Target: "workspaceName", Name: validation.MaxLength, Rule: 63, Chain: nil},
				{Target: "workspaceName", Name: validation.MinLength, Rule: 4, Chain: nil},
				{Target: "workspaceName", Name: validation.Pattern, Rule: `^[A-Za-z0-9][A-Za-z0-9-]+[A-Za-z0-9]$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("operationalinsights.WorkspacePurgeClient", "GetPurgeStatus", err.Error())
	}

	req, err := client.GetPurgeStatusPreparer(ctx, resourceGroupName, workspaceName, purgeID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "operationalinsights.WorkspacePurgeClient", "GetPurgeStatus", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetPurgeStatusSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "operationalinsights.WorkspacePurgeClient", "GetPurgeStatus", resp, "Failure sending request")
		return
	}

	result, err = client.GetPurgeStatusResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "operationalinsights.WorkspacePurgeClient", "GetPurgeStatus", resp, "Failure responding to request")
		return
	}

	return
}

// GetPurgeStatusPreparer prepares the GetPurgeStatus request.
func (client WorkspacePurgeClient) GetPurgeStatusPreparer(ctx context.Context, resourceGroupName string, workspaceName string, purgeID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"purgeId":           autorest.Encode("path", purgeID),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"workspaceName":     autorest.Encode("path", workspaceName),
	}

	const APIVersion = "2020-08-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/operations/{purgeId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetPurgeStatusSender sends the GetPurgeStatus request. The method will close the
// http.Response Body if it receives an error.
func (client WorkspacePurgeClient) GetPurgeStatusSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetPurgeStatusResponder handles the response to the GetPurgeStatus request. The method always
// closes the http.Response Body.
func (client WorkspacePurgeClient) GetPurgeStatusResponder(resp *http.Response) (result WorkspacePurgeStatusResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Purge purges data in an Log Analytics workspace by a set of user-defined filters.
//
// In order to manage system resources, purge requests are throttled at 50 requests per hour. You should batch the
// execution of purge requests by sending a single command whose predicate includes all user identities that require
// purging. Use the in operator to specify multiple identities. You should run the query prior to using for a purge
// request to verify that the results are expected.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// workspaceName - the name of the workspace.
// body - describes the body of a request to purge data in a single table of an Log Analytics Workspace
func (client WorkspacePurgeClient) Purge(ctx context.Context, resourceGroupName string, workspaceName string, body WorkspacePurgeBody) (result WorkspacePurgeResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/WorkspacePurgeClient.Purge")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}},
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: workspaceName,
			Constraints: []validation.Constraint{{Target: "workspaceName", Name: validation.MaxLength, Rule: 63, Chain: nil},
				{Target: "workspaceName", Name: validation.MinLength, Rule: 4, Chain: nil},
				{Target: "workspaceName", Name: validation.Pattern, Rule: `^[A-Za-z0-9][A-Za-z0-9-]+[A-Za-z0-9]$`, Chain: nil}}},
		{TargetValue: body,
			Constraints: []validation.Constraint{{Target: "body.Table", Name: validation.Null, Rule: true, Chain: nil},
				{Target: "body.Filters", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("operationalinsights.WorkspacePurgeClient", "Purge", err.Error())
	}

	req, err := client.PurgePreparer(ctx, resourceGroupName, workspaceName, body)
	if err != nil {
		err = autorest.NewErrorWithError(err, "operationalinsights.WorkspacePurgeClient", "Purge", nil, "Failure preparing request")
		return
	}

	resp, err := client.PurgeSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "operationalinsights.WorkspacePurgeClient", "Purge", resp, "Failure sending request")
		return
	}

	result, err = client.PurgeResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "operationalinsights.WorkspacePurgeClient", "Purge", resp, "Failure responding to request")
		return
	}

	return
}

// PurgePreparer prepares the Purge request.
func (client WorkspacePurgeClient) PurgePreparer(ctx context.Context, resourceGroupName string, workspaceName string, body WorkspacePurgeBody) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"workspaceName":     autorest.Encode("path", workspaceName),
	}

	const APIVersion = "2020-08-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/purge", pathParameters),
		autorest.WithJSON(body),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// PurgeSender sends the Purge request. The method will close the
// http.Response Body if it receives an error.
func (client WorkspacePurgeClient) PurgeSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// PurgeResponder handles the response to the Purge request. The method always
// closes the http.Response Body.
func (client WorkspacePurgeClient) PurgeResponder(resp *http.Response) (result WorkspacePurgeResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
