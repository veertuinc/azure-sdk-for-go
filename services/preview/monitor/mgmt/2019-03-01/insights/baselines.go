package insights

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

// BaselinesClient is the monitor Management Client
type BaselinesClient struct {
	BaseClient
}

// NewBaselinesClient creates an instance of the BaselinesClient client.
func NewBaselinesClient(subscriptionID string) BaselinesClient {
	return NewBaselinesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewBaselinesClientWithBaseURI creates an instance of the BaselinesClient client using a custom endpoint.  Use this
// when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewBaselinesClientWithBaseURI(baseURI string, subscriptionID string) BaselinesClient {
	return BaselinesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// List **Lists the metric baseline values for a resource**.
// Parameters:
// resourceURI - the identifier of the resource.
// metricnames - the names of the metrics (comma separated) to retrieve.
// metricnamespace - metric namespace to query metric definitions for.
// timespan - the timespan of the query. It is a string with the following format
// 'startDateTime_ISO/endDateTime_ISO'.
// interval - the interval (i.e. timegrain) of the query.
// aggregation - the list of aggregation types (comma separated) to retrieve.
// sensitivities - the list of sensitivities (comma separated) to retrieve.
// filter - the **$filter** is used to reduce the set of metric data returned.<br>Example:<br>Metric contains
// metadata A, B and C.<br>- Return all time series of C where A = a1 and B = b1 or b2<br>**$filter=A eq ‘a1’
// and B eq ‘b1’ or B eq ‘b2’ and C eq ‘*’**<br>- Invalid variant:<br>**$filter=A eq ‘a1’ and B eq ‘b1’ and C
// eq ‘*’ or B = ‘b2’**<br>This is invalid because the logical or operator cannot separate two different
// metadata names.<br>- Return all time series where A = a1, B = b1 and C = c1:<br>**$filter=A eq ‘a1’ and B eq
// ‘b1’ and C eq ‘c1’**<br>- Return all time series where A = a1<br>**$filter=A eq ‘a1’ and B eq ‘*’ and C eq
// ‘*’**.
// resultType - allows retrieving only metadata of the baseline. On data request all information is retrieved.
func (client BaselinesClient) List(ctx context.Context, resourceURI string, metricnames string, metricnamespace string, timespan string, interval *string, aggregation string, sensitivities string, filter string, resultType ResultType) (result MetricBaselinesResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BaselinesClient.List")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ListPreparer(ctx, resourceURI, metricnames, metricnamespace, timespan, interval, aggregation, sensitivities, filter, resultType)
	if err != nil {
		err = autorest.NewErrorWithError(err, "insights.BaselinesClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "insights.BaselinesClient", "List", resp, "Failure sending request")
		return
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "insights.BaselinesClient", "List", resp, "Failure responding to request")
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client BaselinesClient) ListPreparer(ctx context.Context, resourceURI string, metricnames string, metricnamespace string, timespan string, interval *string, aggregation string, sensitivities string, filter string, resultType ResultType) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceUri": resourceURI,
	}

	const APIVersion = "2019-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(metricnames) > 0 {
		queryParameters["metricnames"] = autorest.Encode("query", metricnames)
	}
	if len(metricnamespace) > 0 {
		queryParameters["metricnamespace"] = autorest.Encode("query", metricnamespace)
	}
	if len(timespan) > 0 {
		queryParameters["timespan"] = autorest.Encode("query", timespan)
	}
	if interval != nil {
		queryParameters["interval"] = autorest.Encode("query", *interval)
	}
	if len(aggregation) > 0 {
		queryParameters["aggregation"] = autorest.Encode("query", aggregation)
	}
	if len(sensitivities) > 0 {
		queryParameters["sensitivities"] = autorest.Encode("query", sensitivities)
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}
	if len(string(resultType)) > 0 {
		queryParameters["resultType"] = autorest.Encode("query", resultType)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/{resourceUri}/providers/microsoft.insights/metricBaselines", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client BaselinesClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client BaselinesClient) ListResponder(resp *http.Response) (result MetricBaselinesResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
