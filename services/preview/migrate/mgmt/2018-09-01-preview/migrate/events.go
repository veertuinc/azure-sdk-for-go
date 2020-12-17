package migrate

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

// EventsClient is the migrate your workloads to Azure.
type EventsClient struct {
	BaseClient
}

// NewEventsClient creates an instance of the EventsClient client.
func NewEventsClient(subscriptionID string, acceptLanguage string) EventsClient {
	return NewEventsClientWithBaseURI(DefaultBaseURI, subscriptionID, acceptLanguage)
}

// NewEventsClientWithBaseURI creates an instance of the EventsClient client using a custom endpoint.  Use this when
// interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewEventsClientWithBaseURI(baseURI string, subscriptionID string, acceptLanguage string) EventsClient {
	return EventsClient{NewWithBaseURI(baseURI, subscriptionID, acceptLanguage)}
}

// DeleteEvent delete the migrate event. Deleting non-existent migrate event is a no-operation.
// Parameters:
// resourceGroupName - name of the Azure Resource Group that migrate project is part of.
// migrateProjectName - name of the Azure Migrate project.
// eventName - unique name of an event within a migrate project.
func (client EventsClient) DeleteEvent(ctx context.Context, resourceGroupName string, migrateProjectName string, eventName string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/EventsClient.DeleteEvent")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeleteEventPreparer(ctx, resourceGroupName, migrateProjectName, eventName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "migrate.EventsClient", "DeleteEvent", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteEventSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "migrate.EventsClient", "DeleteEvent", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteEventResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "migrate.EventsClient", "DeleteEvent", resp, "Failure responding to request")
		return
	}

	return
}

// DeleteEventPreparer prepares the DeleteEvent request.
func (client EventsClient) DeleteEventPreparer(ctx context.Context, resourceGroupName string, migrateProjectName string, eventName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"eventName":          autorest.Encode("path", eventName),
		"migrateProjectName": autorest.Encode("path", migrateProjectName),
		"resourceGroupName":  autorest.Encode("path", resourceGroupName),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-09-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Migrate/migrateProjects/{migrateProjectName}/migrateEvents/{eventName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteEventSender sends the DeleteEvent request. The method will close the
// http.Response Body if it receives an error.
func (client EventsClient) DeleteEventSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// DeleteEventResponder handles the response to the DeleteEvent request. The method always
// closes the http.Response Body.
func (client EventsClient) DeleteEventResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// EnumerateEvents sends the enumerate events request.
// Parameters:
// resourceGroupName - name of the Azure Resource Group that migrate project is part of.
// migrateProjectName - name of the Azure Migrate project.
// continuationToken - the continuation token.
// pageSize - the number of items to be returned in a single page. This value is honored only if it is less
// than the 100.
func (client EventsClient) EnumerateEvents(ctx context.Context, resourceGroupName string, migrateProjectName string, continuationToken string, pageSize *int32) (result EventCollection, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/EventsClient.EnumerateEvents")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.EnumerateEventsPreparer(ctx, resourceGroupName, migrateProjectName, continuationToken, pageSize)
	if err != nil {
		err = autorest.NewErrorWithError(err, "migrate.EventsClient", "EnumerateEvents", nil, "Failure preparing request")
		return
	}

	resp, err := client.EnumerateEventsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "migrate.EventsClient", "EnumerateEvents", resp, "Failure sending request")
		return
	}

	result, err = client.EnumerateEventsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "migrate.EventsClient", "EnumerateEvents", resp, "Failure responding to request")
		return
	}

	return
}

// EnumerateEventsPreparer prepares the EnumerateEvents request.
func (client EventsClient) EnumerateEventsPreparer(ctx context.Context, resourceGroupName string, migrateProjectName string, continuationToken string, pageSize *int32) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"migrateProjectName": autorest.Encode("path", migrateProjectName),
		"resourceGroupName":  autorest.Encode("path", resourceGroupName),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-09-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(continuationToken) > 0 {
		queryParameters["continuationToken"] = autorest.Encode("query", continuationToken)
	}
	if pageSize != nil {
		queryParameters["pageSize"] = autorest.Encode("query", *pageSize)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Migrate/migrateProjects/{migrateProjectName}/migrateEvents", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	if len(client.AcceptLanguage) > 0 {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("Accept-Language", autorest.String(client.AcceptLanguage)))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// EnumerateEventsSender sends the EnumerateEvents request. The method will close the
// http.Response Body if it receives an error.
func (client EventsClient) EnumerateEventsSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// EnumerateEventsResponder handles the response to the EnumerateEvents request. The method always
// closes the http.Response Body.
func (client EventsClient) EnumerateEventsResponder(resp *http.Response) (result EventCollection, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetEvent sends the get event request.
// Parameters:
// resourceGroupName - name of the Azure Resource Group that migrate project is part of.
// migrateProjectName - name of the Azure Migrate project.
// eventName - unique name of an event within a migrate project.
func (client EventsClient) GetEvent(ctx context.Context, resourceGroupName string, migrateProjectName string, eventName string) (result Event, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/EventsClient.GetEvent")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetEventPreparer(ctx, resourceGroupName, migrateProjectName, eventName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "migrate.EventsClient", "GetEvent", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetEventSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "migrate.EventsClient", "GetEvent", resp, "Failure sending request")
		return
	}

	result, err = client.GetEventResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "migrate.EventsClient", "GetEvent", resp, "Failure responding to request")
		return
	}

	return
}

// GetEventPreparer prepares the GetEvent request.
func (client EventsClient) GetEventPreparer(ctx context.Context, resourceGroupName string, migrateProjectName string, eventName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"eventName":          autorest.Encode("path", eventName),
		"migrateProjectName": autorest.Encode("path", migrateProjectName),
		"resourceGroupName":  autorest.Encode("path", resourceGroupName),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-09-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Migrate/migrateProjects/{migrateProjectName}/migrateEvents/{eventName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetEventSender sends the GetEvent request. The method will close the
// http.Response Body if it receives an error.
func (client EventsClient) GetEventSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetEventResponder handles the response to the GetEvent request. The method always
// closes the http.Response Body.
func (client EventsClient) GetEventResponder(resp *http.Response) (result Event, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
