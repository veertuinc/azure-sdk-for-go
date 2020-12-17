package adhybridhealthservice

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

// UpdateClient is the REST APIs for Azure Active Directory Connect Health
type UpdateClient struct {
	BaseClient
}

// NewUpdateClient creates an instance of the UpdateClient client.
func NewUpdateClient() UpdateClient {
	return NewUpdateClientWithBaseURI(DefaultBaseURI)
}

// NewUpdateClientWithBaseURI creates an instance of the UpdateClient client using a custom endpoint.  Use this when
// interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewUpdateClientWithBaseURI(baseURI string) UpdateClient {
	return UpdateClient{NewWithBaseURI(baseURI)}
}

// IPAddressAggregateSettings updates the IP address aggregate settings alert thresholds.
// Parameters:
// serviceName - the name of the service.
// IPAddressAggregateSetting - the IP address aggregate setting object.
func (client UpdateClient) IPAddressAggregateSettings(ctx context.Context, serviceName string, IPAddressAggregateSetting IPAddressAggregateSetting) (result IPAddressAggregateSetting, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/UpdateClient.IPAddressAggregateSettings")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.IPAddressAggregateSettingsPreparer(ctx, serviceName, IPAddressAggregateSetting)
	if err != nil {
		err = autorest.NewErrorWithError(err, "adhybridhealthservice.UpdateClient", "IPAddressAggregateSettings", nil, "Failure preparing request")
		return
	}

	resp, err := client.IPAddressAggregateSettingsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "adhybridhealthservice.UpdateClient", "IPAddressAggregateSettings", resp, "Failure sending request")
		return
	}

	result, err = client.IPAddressAggregateSettingsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "adhybridhealthservice.UpdateClient", "IPAddressAggregateSettings", resp, "Failure responding to request")
		return
	}

	return
}

// IPAddressAggregateSettingsPreparer prepares the IPAddressAggregateSettings request.
func (client UpdateClient) IPAddressAggregateSettingsPreparer(ctx context.Context, serviceName string, IPAddressAggregateSetting IPAddressAggregateSetting) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"serviceName": autorest.Encode("path", serviceName),
	}

	const APIVersion = "2014-01-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.ADHybridHealthService/services/{serviceName}/ipAddressAggregateSettings", pathParameters),
		autorest.WithJSON(IPAddressAggregateSetting),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// IPAddressAggregateSettingsSender sends the IPAddressAggregateSettings request. The method will close the
// http.Response Body if it receives an error.
func (client UpdateClient) IPAddressAggregateSettingsSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// IPAddressAggregateSettingsResponder handles the response to the IPAddressAggregateSettings request. The method always
// closes the http.Response Body.
func (client UpdateClient) IPAddressAggregateSettingsResponder(resp *http.Response) (result IPAddressAggregateSetting, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
