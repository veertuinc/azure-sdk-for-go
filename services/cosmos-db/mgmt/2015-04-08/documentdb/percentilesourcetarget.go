package documentdb

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

// PercentileSourceTargetClient is the azure Cosmos DB Database Service Resource Provider REST API
type PercentileSourceTargetClient struct {
	BaseClient
}

// NewPercentileSourceTargetClient creates an instance of the PercentileSourceTargetClient client.
func NewPercentileSourceTargetClient(subscriptionID string) PercentileSourceTargetClient {
	return NewPercentileSourceTargetClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewPercentileSourceTargetClientWithBaseURI creates an instance of the PercentileSourceTargetClient client using a
// custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds,
// Azure stack).
func NewPercentileSourceTargetClientWithBaseURI(baseURI string, subscriptionID string) PercentileSourceTargetClient {
	return PercentileSourceTargetClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// ListMetrics retrieves the metrics determined by the given filter for the given account, source and target region.
// This url is only for PBS and Replication Latency data
// Parameters:
// resourceGroupName - name of an Azure resource group.
// accountName - cosmos DB database account name.
// sourceRegion - source region from which data is written. Cosmos DB region, with spaces between words and
// each word capitalized.
// targetRegion - target region to which data is written. Cosmos DB region, with spaces between words and each
// word capitalized.
// filter - an OData filter expression that describes a subset of metrics to return. The parameters that can be
// filtered are name.value (name of the metric, can have an or of multiple names), startTime, endTime, and
// timeGrain. The supported operator is eq.
func (client PercentileSourceTargetClient) ListMetrics(ctx context.Context, resourceGroupName string, accountName string, sourceRegion string, targetRegion string, filter string) (result PercentileMetricListResult, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PercentileSourceTargetClient.ListMetrics")
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
		{TargetValue: accountName,
			Constraints: []validation.Constraint{{Target: "accountName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "accountName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "accountName", Name: validation.Pattern, Rule: `^[a-z0-9]+(-[a-z0-9]+)*`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("documentdb.PercentileSourceTargetClient", "ListMetrics", err.Error())
	}

	req, err := client.ListMetricsPreparer(ctx, resourceGroupName, accountName, sourceRegion, targetRegion, filter)
	if err != nil {
		err = autorest.NewErrorWithError(err, "documentdb.PercentileSourceTargetClient", "ListMetrics", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListMetricsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "documentdb.PercentileSourceTargetClient", "ListMetrics", resp, "Failure sending request")
		return
	}

	result, err = client.ListMetricsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "documentdb.PercentileSourceTargetClient", "ListMetrics", resp, "Failure responding to request")
		return
	}

	return
}

// ListMetricsPreparer prepares the ListMetrics request.
func (client PercentileSourceTargetClient) ListMetricsPreparer(ctx context.Context, resourceGroupName string, accountName string, sourceRegion string, targetRegion string, filter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":       autorest.Encode("path", accountName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"sourceRegion":      autorest.Encode("path", sourceRegion),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"targetRegion":      autorest.Encode("path", targetRegion),
	}

	const APIVersion = "2015-04-08"
	queryParameters := map[string]interface{}{
		"$filter":     autorest.Encode("query", filter),
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/sourceRegion/{sourceRegion}/targetRegion/{targetRegion}/percentile/metrics", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListMetricsSender sends the ListMetrics request. The method will close the
// http.Response Body if it receives an error.
func (client PercentileSourceTargetClient) ListMetricsSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListMetricsResponder handles the response to the ListMetrics request. The method always
// closes the http.Response Body.
func (client PercentileSourceTargetClient) ListMetricsResponder(resp *http.Response) (result PercentileMetricListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
