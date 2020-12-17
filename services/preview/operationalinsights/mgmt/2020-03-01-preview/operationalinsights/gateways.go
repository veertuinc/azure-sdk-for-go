package operationalinsights

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

// GatewaysClient is the operational Insights Client
type GatewaysClient struct {
	BaseClient
}

// NewGatewaysClient creates an instance of the GatewaysClient client.
func NewGatewaysClient(subscriptionID string) GatewaysClient {
	return NewGatewaysClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewGatewaysClientWithBaseURI creates an instance of the GatewaysClient client using a custom endpoint.  Use this
// when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewGatewaysClientWithBaseURI(baseURI string, subscriptionID string) GatewaysClient {
	return GatewaysClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Delete delete a Log Analytics gateway.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// workspaceName - the name of the workspace.
// gatewayID - the Log Analytics gateway Id.
func (client GatewaysClient) Delete(ctx context.Context, resourceGroupName string, workspaceName string, gatewayID string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/GatewaysClient.Delete")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}},
		{TargetValue: workspaceName,
			Constraints: []validation.Constraint{{Target: "workspaceName", Name: validation.MaxLength, Rule: 63, Chain: nil},
				{Target: "workspaceName", Name: validation.MinLength, Rule: 4, Chain: nil},
				{Target: "workspaceName", Name: validation.Pattern, Rule: `^[A-Za-z0-9][A-Za-z0-9-]+[A-Za-z0-9]$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("operationalinsights.GatewaysClient", "Delete", err.Error())
	}

	req, err := client.DeletePreparer(ctx, resourceGroupName, workspaceName, gatewayID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "operationalinsights.GatewaysClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "operationalinsights.GatewaysClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "operationalinsights.GatewaysClient", "Delete", resp, "Failure responding to request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client GatewaysClient) DeletePreparer(ctx context.Context, resourceGroupName string, workspaceName string, gatewayID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"gatewayId":         autorest.Encode("path", gatewayID),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"workspaceName":     autorest.Encode("path", workspaceName),
	}

	const APIVersion = "2020-03-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/gateways/{gatewayId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client GatewaysClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client GatewaysClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}
