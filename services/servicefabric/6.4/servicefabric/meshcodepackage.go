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

// MeshCodePackageClient is the service Fabric REST Client APIs allows management of Service Fabric clusters,
// applications and services.
type MeshCodePackageClient struct {
	BaseClient
}

// NewMeshCodePackageClient creates an instance of the MeshCodePackageClient client.
func NewMeshCodePackageClient() MeshCodePackageClient {
	return NewMeshCodePackageClientWithBaseURI(DefaultBaseURI)
}

// NewMeshCodePackageClientWithBaseURI creates an instance of the MeshCodePackageClient client using a custom endpoint.
// Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewMeshCodePackageClientWithBaseURI(baseURI string) MeshCodePackageClient {
	return MeshCodePackageClient{NewWithBaseURI(baseURI)}
}

// GetContainerLogs gets the logs for the container of the specified code package of the service replica.
// Parameters:
// applicationResourceName - the identity of the application.
// serviceResourceName - the identity of the service.
// replicaName - service Fabric replica name.
// codePackageName - the name of code package of the service.
// tail - number of lines to show from the end of the logs. Default is 100. 'all' to show the complete logs.
func (client MeshCodePackageClient) GetContainerLogs(ctx context.Context, applicationResourceName string, serviceResourceName string, replicaName string, codePackageName string, tail string) (result ContainerLogs, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/MeshCodePackageClient.GetContainerLogs")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetContainerLogsPreparer(ctx, applicationResourceName, serviceResourceName, replicaName, codePackageName, tail)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabric.MeshCodePackageClient", "GetContainerLogs", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetContainerLogsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "servicefabric.MeshCodePackageClient", "GetContainerLogs", resp, "Failure sending request")
		return
	}

	result, err = client.GetContainerLogsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabric.MeshCodePackageClient", "GetContainerLogs", resp, "Failure responding to request")
		return
	}

	return
}

// GetContainerLogsPreparer prepares the GetContainerLogs request.
func (client MeshCodePackageClient) GetContainerLogsPreparer(ctx context.Context, applicationResourceName string, serviceResourceName string, replicaName string, codePackageName string, tail string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"applicationResourceName": applicationResourceName,
		"codePackageName":         autorest.Encode("path", codePackageName),
		"replicaName":             replicaName,
		"serviceResourceName":     serviceResourceName,
	}

	const APIVersion = "6.4-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(tail) > 0 {
		queryParameters["Tail"] = autorest.Encode("query", tail)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/Resources/Applications/{applicationResourceName}/Services/{serviceResourceName}/Replicas/{replicaName}/CodePackages/{codePackageName}/Logs", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetContainerLogsSender sends the GetContainerLogs request. The method will close the
// http.Response Body if it receives an error.
func (client MeshCodePackageClient) GetContainerLogsSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetContainerLogsResponder handles the response to the GetContainerLogs request. The method always
// closes the http.Response Body.
func (client MeshCodePackageClient) GetContainerLogsResponder(resp *http.Response) (result ContainerLogs, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
