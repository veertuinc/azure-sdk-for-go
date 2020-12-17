package trafficmanager

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

// HeatMapClient is the client for the HeatMap methods of the Trafficmanager service.
type HeatMapClient struct {
	BaseClient
}

// NewHeatMapClient creates an instance of the HeatMapClient client.
func NewHeatMapClient(subscriptionID string) HeatMapClient {
	return NewHeatMapClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewHeatMapClientWithBaseURI creates an instance of the HeatMapClient client using a custom endpoint.  Use this when
// interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewHeatMapClientWithBaseURI(baseURI string, subscriptionID string) HeatMapClient {
	return HeatMapClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Get gets latest heatmap for Traffic Manager profile.
// Parameters:
// resourceGroupName - the name of the resource group containing the Traffic Manager endpoint.
// profileName - the name of the Traffic Manager profile.
// topLeft - the top left latitude,longitude pair of the rectangular viewport to query for.
// botRight - the bottom right latitude,longitude pair of the rectangular viewport to query for.
func (client HeatMapClient) Get(ctx context.Context, resourceGroupName string, profileName string, topLeft []float64, botRight []float64) (result HeatMapModel, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/HeatMapClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: topLeft,
			Constraints: []validation.Constraint{{Target: "topLeft", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "topLeft", Name: validation.MaxItems, Rule: 2, Chain: nil},
					{Target: "topLeft", Name: validation.MinItems, Rule: 2, Chain: nil},
				}}}},
		{TargetValue: botRight,
			Constraints: []validation.Constraint{{Target: "botRight", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "botRight", Name: validation.MaxItems, Rule: 2, Chain: nil},
					{Target: "botRight", Name: validation.MinItems, Rule: 2, Chain: nil},
				}}}}}); err != nil {
		return result, validation.NewError("trafficmanager.HeatMapClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, resourceGroupName, profileName, topLeft, botRight)
	if err != nil {
		err = autorest.NewErrorWithError(err, "trafficmanager.HeatMapClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "trafficmanager.HeatMapClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "trafficmanager.HeatMapClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client HeatMapClient) GetPreparer(ctx context.Context, resourceGroupName string, profileName string, topLeft []float64, botRight []float64) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"heatMapType":       autorest.Encode("path", "default"),
		"profileName":       autorest.Encode("path", profileName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-09-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if topLeft != nil && len(topLeft) > 0 {
		queryParameters["topLeft"] = autorest.Encode("query", topLeft, ",")
	}
	if botRight != nil && len(botRight) > 0 {
		queryParameters["botRight"] = autorest.Encode("query", botRight, ",")
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/trafficmanagerprofiles/{profileName}/heatMaps/{heatMapType}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client HeatMapClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client HeatMapClient) GetResponder(resp *http.Response) (result HeatMapModel, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
