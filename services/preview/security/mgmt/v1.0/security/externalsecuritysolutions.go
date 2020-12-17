package security

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

// ExternalSecuritySolutionsClient is the API spec for Microsoft.Security (Azure Security Center) resource provider
type ExternalSecuritySolutionsClient struct {
	BaseClient
}

// NewExternalSecuritySolutionsClient creates an instance of the ExternalSecuritySolutionsClient client.
func NewExternalSecuritySolutionsClient(subscriptionID string, ascLocation string) ExternalSecuritySolutionsClient {
	return NewExternalSecuritySolutionsClientWithBaseURI(DefaultBaseURI, subscriptionID, ascLocation)
}

// NewExternalSecuritySolutionsClientWithBaseURI creates an instance of the ExternalSecuritySolutionsClient client
// using a custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign
// clouds, Azure stack).
func NewExternalSecuritySolutionsClientWithBaseURI(baseURI string, subscriptionID string, ascLocation string) ExternalSecuritySolutionsClient {
	return ExternalSecuritySolutionsClient{NewWithBaseURI(baseURI, subscriptionID, ascLocation)}
}

// Get gets a specific external Security Solution.
// Parameters:
// resourceGroupName - the name of the resource group within the user's subscription. The name is case
// insensitive.
// externalSecuritySolutionsName - name of an external security solution.
func (client ExternalSecuritySolutionsClient) Get(ctx context.Context, resourceGroupName string, externalSecuritySolutionsName string) (result ExternalSecuritySolutionModel, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ExternalSecuritySolutionsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.Pattern, Rule: `^[0-9A-Fa-f]{8}-([0-9A-Fa-f]{4}-){3}[0-9A-Fa-f]{12}$`, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("security.ExternalSecuritySolutionsClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, resourceGroupName, externalSecuritySolutionsName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "security.ExternalSecuritySolutionsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "security.ExternalSecuritySolutionsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "security.ExternalSecuritySolutionsClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client ExternalSecuritySolutionsClient) GetPreparer(ctx context.Context, resourceGroupName string, externalSecuritySolutionsName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"ascLocation":                   autorest.Encode("path", client.AscLocation),
		"externalSecuritySolutionsName": autorest.Encode("path", externalSecuritySolutionsName),
		"resourceGroupName":             autorest.Encode("path", resourceGroupName),
		"subscriptionId":                autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Security/locations/{ascLocation}/ExternalSecuritySolutions/{externalSecuritySolutionsName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ExternalSecuritySolutionsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ExternalSecuritySolutionsClient) GetResponder(resp *http.Response) (result ExternalSecuritySolutionModel, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List gets a list of external security solutions for the subscription.
func (client ExternalSecuritySolutionsClient) List(ctx context.Context) (result ExternalSecuritySolutionListPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ExternalSecuritySolutionsClient.List")
		defer func() {
			sc := -1
			if result.essl.Response.Response != nil {
				sc = result.essl.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.Pattern, Rule: `^[0-9A-Fa-f]{8}-([0-9A-Fa-f]{4}-){3}[0-9A-Fa-f]{12}$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("security.ExternalSecuritySolutionsClient", "List", err.Error())
	}

	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "security.ExternalSecuritySolutionsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.essl.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "security.ExternalSecuritySolutionsClient", "List", resp, "Failure sending request")
		return
	}

	result.essl, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "security.ExternalSecuritySolutionsClient", "List", resp, "Failure responding to request")
		return
	}
	if result.essl.hasNextLink() && result.essl.IsEmpty() {
		err = result.NextWithContext(ctx)
	}

	return
}

// ListPreparer prepares the List request.
func (client ExternalSecuritySolutionsClient) ListPreparer(ctx context.Context) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Security/externalSecuritySolutions", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client ExternalSecuritySolutionsClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client ExternalSecuritySolutionsClient) ListResponder(resp *http.Response) (result ExternalSecuritySolutionList, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client ExternalSecuritySolutionsClient) listNextResults(ctx context.Context, lastResults ExternalSecuritySolutionList) (result ExternalSecuritySolutionList, err error) {
	req, err := lastResults.externalSecuritySolutionListPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "security.ExternalSecuritySolutionsClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "security.ExternalSecuritySolutionsClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "security.ExternalSecuritySolutionsClient", "listNextResults", resp, "Failure responding to next results request")
		return
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client ExternalSecuritySolutionsClient) ListComplete(ctx context.Context) (result ExternalSecuritySolutionListIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ExternalSecuritySolutionsClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx)
	return
}

// ListByHomeRegion gets a list of external Security Solutions for the subscription and location.
func (client ExternalSecuritySolutionsClient) ListByHomeRegion(ctx context.Context) (result ExternalSecuritySolutionListPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ExternalSecuritySolutionsClient.ListByHomeRegion")
		defer func() {
			sc := -1
			if result.essl.Response.Response != nil {
				sc = result.essl.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.Pattern, Rule: `^[0-9A-Fa-f]{8}-([0-9A-Fa-f]{4}-){3}[0-9A-Fa-f]{12}$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("security.ExternalSecuritySolutionsClient", "ListByHomeRegion", err.Error())
	}

	result.fn = client.listByHomeRegionNextResults
	req, err := client.ListByHomeRegionPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "security.ExternalSecuritySolutionsClient", "ListByHomeRegion", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByHomeRegionSender(req)
	if err != nil {
		result.essl.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "security.ExternalSecuritySolutionsClient", "ListByHomeRegion", resp, "Failure sending request")
		return
	}

	result.essl, err = client.ListByHomeRegionResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "security.ExternalSecuritySolutionsClient", "ListByHomeRegion", resp, "Failure responding to request")
		return
	}
	if result.essl.hasNextLink() && result.essl.IsEmpty() {
		err = result.NextWithContext(ctx)
	}

	return
}

// ListByHomeRegionPreparer prepares the ListByHomeRegion request.
func (client ExternalSecuritySolutionsClient) ListByHomeRegionPreparer(ctx context.Context) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"ascLocation":    autorest.Encode("path", client.AscLocation),
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Security/locations/{ascLocation}/ExternalSecuritySolutions", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByHomeRegionSender sends the ListByHomeRegion request. The method will close the
// http.Response Body if it receives an error.
func (client ExternalSecuritySolutionsClient) ListByHomeRegionSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByHomeRegionResponder handles the response to the ListByHomeRegion request. The method always
// closes the http.Response Body.
func (client ExternalSecuritySolutionsClient) ListByHomeRegionResponder(resp *http.Response) (result ExternalSecuritySolutionList, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByHomeRegionNextResults retrieves the next set of results, if any.
func (client ExternalSecuritySolutionsClient) listByHomeRegionNextResults(ctx context.Context, lastResults ExternalSecuritySolutionList) (result ExternalSecuritySolutionList, err error) {
	req, err := lastResults.externalSecuritySolutionListPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "security.ExternalSecuritySolutionsClient", "listByHomeRegionNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByHomeRegionSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "security.ExternalSecuritySolutionsClient", "listByHomeRegionNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByHomeRegionResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "security.ExternalSecuritySolutionsClient", "listByHomeRegionNextResults", resp, "Failure responding to next results request")
		return
	}
	return
}

// ListByHomeRegionComplete enumerates all values, automatically crossing page boundaries as required.
func (client ExternalSecuritySolutionsClient) ListByHomeRegionComplete(ctx context.Context) (result ExternalSecuritySolutionListIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ExternalSecuritySolutionsClient.ListByHomeRegion")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByHomeRegion(ctx)
	return
}
