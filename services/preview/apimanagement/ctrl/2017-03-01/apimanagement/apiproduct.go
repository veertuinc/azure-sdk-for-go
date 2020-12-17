package apimanagement

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

// APIProductClient is the client for the APIProduct methods of the Apimanagement service.
type APIProductClient struct {
	BaseClient
}

// NewAPIProductClient creates an instance of the APIProductClient client.
func NewAPIProductClient() APIProductClient {
	return APIProductClient{New()}
}

// ListByApis lists all Products, which the API is part of.
// Parameters:
// apimBaseURL - the management endpoint of the API Management service, for example
// https://myapimservice.management.azure-api.net.
// apiid - API identifier. Must be unique in the current API Management service instance.
// filter - | Field | Supported operators    | Supported functions                         |
// |-------|------------------------|---------------------------------------------|
// | name  | ge, le, eq, ne, gt, lt | substringof, contains, startswith, endswith |
// top - number of records to return.
// skip - number of records to skip.
func (client APIProductClient) ListByApis(ctx context.Context, apimBaseURL string, apiid string, filter string, top *int32, skip *int32) (result ProductCollectionPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/APIProductClient.ListByApis")
		defer func() {
			sc := -1
			if result.pc.Response.Response != nil {
				sc = result.pc.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: apiid,
			Constraints: []validation.Constraint{{Target: "apiid", Name: validation.MaxLength, Rule: 256, Chain: nil},
				{Target: "apiid", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "apiid", Name: validation.Pattern, Rule: `^[^*#&+:<>?]+$`, Chain: nil}}},
		{TargetValue: top,
			Constraints: []validation.Constraint{{Target: "top", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "top", Name: validation.InclusiveMinimum, Rule: int64(1), Chain: nil}}}}},
		{TargetValue: skip,
			Constraints: []validation.Constraint{{Target: "skip", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "skip", Name: validation.InclusiveMinimum, Rule: int64(0), Chain: nil}}}}}}); err != nil {
		return result, validation.NewError("apimanagement.APIProductClient", "ListByApis", err.Error())
	}

	result.fn = client.listByApisNextResults
	req, err := client.ListByApisPreparer(ctx, apimBaseURL, apiid, filter, top, skip)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.APIProductClient", "ListByApis", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByApisSender(req)
	if err != nil {
		result.pc.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "apimanagement.APIProductClient", "ListByApis", resp, "Failure sending request")
		return
	}

	result.pc, err = client.ListByApisResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.APIProductClient", "ListByApis", resp, "Failure responding to request")
		return
	}
	if result.pc.hasNextLink() && result.pc.IsEmpty() {
		err = result.NextWithContext(ctx)
	}

	return
}

// ListByApisPreparer prepares the ListByApis request.
func (client APIProductClient) ListByApisPreparer(ctx context.Context, apimBaseURL string, apiid string, filter string, top *int32, skip *int32) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"apimBaseUrl": apimBaseURL,
	}

	pathParameters := map[string]interface{}{
		"apiId": autorest.Encode("path", apiid),
	}

	const APIVersion = "2017-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}
	if top != nil {
		queryParameters["$top"] = autorest.Encode("query", *top)
	}
	if skip != nil {
		queryParameters["$skip"] = autorest.Encode("query", *skip)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithCustomBaseURL("{apimBaseUrl}", urlParameters),
		autorest.WithPathParameters("/apis/{apiId}/products", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByApisSender sends the ListByApis request. The method will close the
// http.Response Body if it receives an error.
func (client APIProductClient) ListByApisSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListByApisResponder handles the response to the ListByApis request. The method always
// closes the http.Response Body.
func (client APIProductClient) ListByApisResponder(resp *http.Response) (result ProductCollection, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByApisNextResults retrieves the next set of results, if any.
func (client APIProductClient) listByApisNextResults(ctx context.Context, lastResults ProductCollection) (result ProductCollection, err error) {
	req, err := lastResults.productCollectionPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "apimanagement.APIProductClient", "listByApisNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByApisSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "apimanagement.APIProductClient", "listByApisNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByApisResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.APIProductClient", "listByApisNextResults", resp, "Failure responding to next results request")
		return
	}
	return
}

// ListByApisComplete enumerates all values, automatically crossing page boundaries as required.
func (client APIProductClient) ListByApisComplete(ctx context.Context, apimBaseURL string, apiid string, filter string, top *int32, skip *int32) (result ProductCollectionIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/APIProductClient.ListByApis")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByApis(ctx, apimBaseURL, apiid, filter, top, skip)
	return
}
