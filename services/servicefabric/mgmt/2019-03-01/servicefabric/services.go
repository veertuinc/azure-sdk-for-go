package servicefabric

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

// ServicesClient is the service Fabric Management Client
type ServicesClient struct {
	BaseClient
}

// NewServicesClient creates an instance of the ServicesClient client.
func NewServicesClient(subscriptionID string) ServicesClient {
	return NewServicesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewServicesClientWithBaseURI creates an instance of the ServicesClient client using a custom endpoint.  Use this
// when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewServicesClientWithBaseURI(baseURI string, subscriptionID string) ServicesClient {
	return ServicesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate create or update a Service Fabric service resource with the specified name.
// Parameters:
// resourceGroupName - the name of the resource group.
// clusterName - the name of the cluster resource.
// applicationName - the name of the application resource.
// serviceName - the name of the service resource in the format of {applicationName}~{serviceName}.
// parameters - the service resource.
func (client ServicesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, clusterName string, applicationName string, serviceName string, parameters ServiceResource) (result ServicesCreateOrUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ServicesClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, clusterName, applicationName, serviceName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabric.ServicesClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabric.ServicesClient", "CreateOrUpdate", result.Response(), "Failure sending request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client ServicesClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, clusterName string, applicationName string, serviceName string, parameters ServiceResource) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"applicationName":   autorest.Encode("path", applicationName),
		"clusterName":       autorest.Encode("path", clusterName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serviceName":       autorest.Encode("path", serviceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceFabric/clusters/{clusterName}/applications/{applicationName}/services/{serviceName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client ServicesClient) CreateOrUpdateSender(req *http.Request) (future ServicesCreateOrUpdateFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client ServicesClient) CreateOrUpdateResponder(resp *http.Response) (result ServiceResource, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete delete a Service Fabric service resource with the specified name.
// Parameters:
// resourceGroupName - the name of the resource group.
// clusterName - the name of the cluster resource.
// applicationName - the name of the application resource.
// serviceName - the name of the service resource in the format of {applicationName}~{serviceName}.
func (client ServicesClient) Delete(ctx context.Context, resourceGroupName string, clusterName string, applicationName string, serviceName string) (result ServicesDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ServicesClient.Delete")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, resourceGroupName, clusterName, applicationName, serviceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabric.ServicesClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabric.ServicesClient", "Delete", result.Response(), "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client ServicesClient) DeletePreparer(ctx context.Context, resourceGroupName string, clusterName string, applicationName string, serviceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"applicationName":   autorest.Encode("path", applicationName),
		"clusterName":       autorest.Encode("path", clusterName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serviceName":       autorest.Encode("path", serviceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceFabric/clusters/{clusterName}/applications/{applicationName}/services/{serviceName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client ServicesClient) DeleteSender(req *http.Request) (future ServicesDeleteFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client ServicesClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get get a Service Fabric service resource created or in the process of being created in the Service Fabric
// application resource.
// Parameters:
// resourceGroupName - the name of the resource group.
// clusterName - the name of the cluster resource.
// applicationName - the name of the application resource.
// serviceName - the name of the service resource in the format of {applicationName}~{serviceName}.
func (client ServicesClient) Get(ctx context.Context, resourceGroupName string, clusterName string, applicationName string, serviceName string) (result ServiceResource, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ServicesClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, clusterName, applicationName, serviceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabric.ServicesClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "servicefabric.ServicesClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabric.ServicesClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client ServicesClient) GetPreparer(ctx context.Context, resourceGroupName string, clusterName string, applicationName string, serviceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"applicationName":   autorest.Encode("path", applicationName),
		"clusterName":       autorest.Encode("path", clusterName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serviceName":       autorest.Encode("path", serviceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceFabric/clusters/{clusterName}/applications/{applicationName}/services/{serviceName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ServicesClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ServicesClient) GetResponder(resp *http.Response) (result ServiceResource, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List gets all service resources created or in the process of being created in the Service Fabric application
// resource.
// Parameters:
// resourceGroupName - the name of the resource group.
// clusterName - the name of the cluster resource.
// applicationName - the name of the application resource.
func (client ServicesClient) List(ctx context.Context, resourceGroupName string, clusterName string, applicationName string) (result ServiceResourceList, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ServicesClient.List")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ListPreparer(ctx, resourceGroupName, clusterName, applicationName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabric.ServicesClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "servicefabric.ServicesClient", "List", resp, "Failure sending request")
		return
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabric.ServicesClient", "List", resp, "Failure responding to request")
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client ServicesClient) ListPreparer(ctx context.Context, resourceGroupName string, clusterName string, applicationName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"applicationName":   autorest.Encode("path", applicationName),
		"clusterName":       autorest.Encode("path", clusterName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceFabric/clusters/{clusterName}/applications/{applicationName}/services", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client ServicesClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client ServicesClient) ListResponder(resp *http.Response) (result ServiceResourceList, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Update update a Service Fabric service resource with the specified name.
// Parameters:
// resourceGroupName - the name of the resource group.
// clusterName - the name of the cluster resource.
// applicationName - the name of the application resource.
// serviceName - the name of the service resource in the format of {applicationName}~{serviceName}.
// parameters - the service resource for patch operations.
func (client ServicesClient) Update(ctx context.Context, resourceGroupName string, clusterName string, applicationName string, serviceName string, parameters ServiceResourceUpdate) (result ServicesUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ServicesClient.Update")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.UpdatePreparer(ctx, resourceGroupName, clusterName, applicationName, serviceName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabric.ServicesClient", "Update", nil, "Failure preparing request")
		return
	}

	result, err = client.UpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabric.ServicesClient", "Update", result.Response(), "Failure sending request")
		return
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client ServicesClient) UpdatePreparer(ctx context.Context, resourceGroupName string, clusterName string, applicationName string, serviceName string, parameters ServiceResourceUpdate) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"applicationName":   autorest.Encode("path", applicationName),
		"clusterName":       autorest.Encode("path", clusterName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serviceName":       autorest.Encode("path", serviceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceFabric/clusters/{clusterName}/applications/{applicationName}/services/{serviceName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client ServicesClient) UpdateSender(req *http.Request) (future ServicesUpdateFuture, err error) {
	var resp *http.Response
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client ServicesClient) UpdateResponder(resp *http.Response) (result ServiceResource, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
