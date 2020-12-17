package consumption

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

// ReservationsSummariesClient is the consumption management client provides access to consumption resources for Azure
// Enterprise Subscriptions.
type ReservationsSummariesClient struct {
	BaseClient
}

// NewReservationsSummariesClient creates an instance of the ReservationsSummariesClient client.
func NewReservationsSummariesClient(subscriptionID string) ReservationsSummariesClient {
	return NewReservationsSummariesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewReservationsSummariesClientWithBaseURI creates an instance of the ReservationsSummariesClient client using a
// custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds,
// Azure stack).
func NewReservationsSummariesClientWithBaseURI(baseURI string, subscriptionID string) ReservationsSummariesClient {
	return ReservationsSummariesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// List lists the reservations summaries for daily or monthly grain.
// Parameters:
// scope - the scope of the reservation summaries. The scope can be
// 'providers/Microsoft.Capacity/reservationorders/{ReservationOrderId}' or
// 'providers/Microsoft.Capacity/reservationorders/{ReservationOrderId}/reservations/{ReservationId}'
// grain - can be daily or monthly
// filter - required only for daily grain. The properties/UsageDate for start date and end date. The filter
// supports 'le' and  'ge'
func (client ReservationsSummariesClient) List(ctx context.Context, scope string, grain Datagrain, filter string) (result ReservationSummariesListResult, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ReservationsSummariesClient.List")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ListPreparer(ctx, scope, grain, filter)
	if err != nil {
		err = autorest.NewErrorWithError(err, "consumption.ReservationsSummariesClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "consumption.ReservationsSummariesClient", "List", resp, "Failure sending request")
		return
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "consumption.ReservationsSummariesClient", "List", resp, "Failure responding to request")
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client ReservationsSummariesClient) ListPreparer(ctx context.Context, scope string, grain Datagrain, filter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"scope": scope,
	}

	const APIVersion = "2017-11-30"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
		"grain":       autorest.Encode("query", grain),
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/{scope}/providers/Microsoft.Consumption/reservationSummaries", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client ReservationsSummariesClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client ReservationsSummariesClient) ListResponder(resp *http.Response) (result ReservationSummariesListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
