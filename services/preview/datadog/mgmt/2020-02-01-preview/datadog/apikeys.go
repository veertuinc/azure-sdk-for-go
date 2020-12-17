package datadog

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

// APIKeysClient is the client for the APIKeys methods of the Datadog service.
type APIKeysClient struct {
	BaseClient
}

// NewAPIKeysClient creates an instance of the APIKeysClient client.
func NewAPIKeysClient(subscriptionID string) APIKeysClient {
	return NewAPIKeysClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewAPIKeysClientWithBaseURI creates an instance of the APIKeysClient client using a custom endpoint.  Use this when
// interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewAPIKeysClientWithBaseURI(baseURI string, subscriptionID string) APIKeysClient {
	return APIKeysClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// GetDefaultKey sends the get default key request.
// Parameters:
// resourceGroupName - the name of the resource group to which the Datadog resource belongs.
// monitorName - monitor resource name
func (client APIKeysClient) GetDefaultKey(ctx context.Context, resourceGroupName string, monitorName string) (result APIKey, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/APIKeysClient.GetDefaultKey")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetDefaultKeyPreparer(ctx, resourceGroupName, monitorName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datadog.APIKeysClient", "GetDefaultKey", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetDefaultKeySender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "datadog.APIKeysClient", "GetDefaultKey", resp, "Failure sending request")
		return
	}

	result, err = client.GetDefaultKeyResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datadog.APIKeysClient", "GetDefaultKey", resp, "Failure responding to request")
		return
	}

	return
}

// GetDefaultKeyPreparer prepares the GetDefaultKey request.
func (client APIKeysClient) GetDefaultKeyPreparer(ctx context.Context, resourceGroupName string, monitorName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"monitorName":       autorest.Encode("path", monitorName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-02-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Datadog/monitors/{monitorName}/getDefaultKey", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetDefaultKeySender sends the GetDefaultKey request. The method will close the
// http.Response Body if it receives an error.
func (client APIKeysClient) GetDefaultKeySender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetDefaultKeyResponder handles the response to the GetDefaultKey request. The method always
// closes the http.Response Body.
func (client APIKeysClient) GetDefaultKeyResponder(resp *http.Response) (result APIKey, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List sends the list request.
// Parameters:
// resourceGroupName - the name of the resource group to which the Datadog resource belongs.
// monitorName - monitor resource name
func (client APIKeysClient) List(ctx context.Context, resourceGroupName string, monitorName string) (result APIKeyListResponsePage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/APIKeysClient.List")
		defer func() {
			sc := -1
			if result.aklr.Response.Response != nil {
				sc = result.aklr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, resourceGroupName, monitorName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datadog.APIKeysClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.aklr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "datadog.APIKeysClient", "List", resp, "Failure sending request")
		return
	}

	result.aklr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datadog.APIKeysClient", "List", resp, "Failure responding to request")
		return
	}
	if result.aklr.hasNextLink() && result.aklr.IsEmpty() {
		err = result.NextWithContext(ctx)
	}

	return
}

// ListPreparer prepares the List request.
func (client APIKeysClient) ListPreparer(ctx context.Context, resourceGroupName string, monitorName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"monitorName":       autorest.Encode("path", monitorName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-02-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Datadog/monitors/{monitorName}/listApiKeys", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client APIKeysClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client APIKeysClient) ListResponder(resp *http.Response) (result APIKeyListResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client APIKeysClient) listNextResults(ctx context.Context, lastResults APIKeyListResponse) (result APIKeyListResponse, err error) {
	req, err := lastResults.aPIKeyListResponsePreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "datadog.APIKeysClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "datadog.APIKeysClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datadog.APIKeysClient", "listNextResults", resp, "Failure responding to next results request")
		return
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client APIKeysClient) ListComplete(ctx context.Context, resourceGroupName string, monitorName string) (result APIKeyListResponseIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/APIKeysClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, resourceGroupName, monitorName)
	return
}

// SetDefaultKey sends the set default key request.
// Parameters:
// resourceGroupName - the name of the resource group to which the Datadog resource belongs.
// monitorName - monitor resource name
func (client APIKeysClient) SetDefaultKey(ctx context.Context, resourceGroupName string, monitorName string, body *APIKey) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/APIKeysClient.SetDefaultKey")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: body,
			Constraints: []validation.Constraint{{Target: "body", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "body.Key", Name: validation.Null, Rule: true, Chain: nil}}}}}}); err != nil {
		return result, validation.NewError("datadog.APIKeysClient", "SetDefaultKey", err.Error())
	}

	req, err := client.SetDefaultKeyPreparer(ctx, resourceGroupName, monitorName, body)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datadog.APIKeysClient", "SetDefaultKey", nil, "Failure preparing request")
		return
	}

	resp, err := client.SetDefaultKeySender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "datadog.APIKeysClient", "SetDefaultKey", resp, "Failure sending request")
		return
	}

	result, err = client.SetDefaultKeyResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datadog.APIKeysClient", "SetDefaultKey", resp, "Failure responding to request")
		return
	}

	return
}

// SetDefaultKeyPreparer prepares the SetDefaultKey request.
func (client APIKeysClient) SetDefaultKeyPreparer(ctx context.Context, resourceGroupName string, monitorName string, body *APIKey) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"monitorName":       autorest.Encode("path", monitorName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-02-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Datadog/monitors/{monitorName}/setDefaultKey", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	if body != nil {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithJSON(body))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// SetDefaultKeySender sends the SetDefaultKey request. The method will close the
// http.Response Body if it receives an error.
func (client APIKeysClient) SetDefaultKeySender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// SetDefaultKeyResponder handles the response to the SetDefaultKey request. The method always
// closes the http.Response Body.
func (client APIKeysClient) SetDefaultKeyResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}
