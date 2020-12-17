package postgresql

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

// LocationBasedPerformanceTierClient is the the Microsoft Azure management API provides create, read, update, and
// delete functionality for Azure PostgreSQL resources including servers, databases, firewall rules, VNET rules,
// security alert policies, log files and configurations with new business model.
type LocationBasedPerformanceTierClient struct {
	BaseClient
}

// NewLocationBasedPerformanceTierClient creates an instance of the LocationBasedPerformanceTierClient client.
func NewLocationBasedPerformanceTierClient(subscriptionID string) LocationBasedPerformanceTierClient {
	return NewLocationBasedPerformanceTierClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewLocationBasedPerformanceTierClientWithBaseURI creates an instance of the LocationBasedPerformanceTierClient
// client using a custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI
// (sovereign clouds, Azure stack).
func NewLocationBasedPerformanceTierClientWithBaseURI(baseURI string, subscriptionID string) LocationBasedPerformanceTierClient {
	return LocationBasedPerformanceTierClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// List list all the performance tiers at specified location in a given subscription.
// Parameters:
// locationName - the name of the location.
func (client LocationBasedPerformanceTierClient) List(ctx context.Context, locationName string) (result PerformanceTierListResult, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/LocationBasedPerformanceTierClient.List")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ListPreparer(ctx, locationName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "postgresql.LocationBasedPerformanceTierClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "postgresql.LocationBasedPerformanceTierClient", "List", resp, "Failure sending request")
		return
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "postgresql.LocationBasedPerformanceTierClient", "List", resp, "Failure responding to request")
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client LocationBasedPerformanceTierClient) ListPreparer(ctx context.Context, locationName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"locationName":   autorest.Encode("path", locationName),
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-12-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.DBForPostgreSQL/locations/{locationName}/performanceTiers", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client LocationBasedPerformanceTierClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client LocationBasedPerformanceTierClient) ListResponder(resp *http.Response) (result PerformanceTierListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
