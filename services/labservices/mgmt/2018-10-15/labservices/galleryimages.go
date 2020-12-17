package labservices

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

// GalleryImagesClient is the the Managed Labs Client.
type GalleryImagesClient struct {
	BaseClient
}

// NewGalleryImagesClient creates an instance of the GalleryImagesClient client.
func NewGalleryImagesClient(subscriptionID string) GalleryImagesClient {
	return NewGalleryImagesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewGalleryImagesClientWithBaseURI creates an instance of the GalleryImagesClient client using a custom endpoint.
// Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewGalleryImagesClientWithBaseURI(baseURI string, subscriptionID string) GalleryImagesClient {
	return GalleryImagesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate create or replace an existing Gallery Image.
// Parameters:
// resourceGroupName - the name of the resource group.
// labAccountName - the name of the lab Account.
// galleryImageName - the name of the gallery Image.
// galleryImage - represents an image from the Azure Marketplace
func (client GalleryImagesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, labAccountName string, galleryImageName string, galleryImage GalleryImage) (result GalleryImage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/GalleryImagesClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, labAccountName, galleryImageName, galleryImage)
	if err != nil {
		err = autorest.NewErrorWithError(err, "labservices.GalleryImagesClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "labservices.GalleryImagesClient", "CreateOrUpdate", resp, "Failure sending request")
		return
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "labservices.GalleryImagesClient", "CreateOrUpdate", resp, "Failure responding to request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client GalleryImagesClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, labAccountName string, galleryImageName string, galleryImage GalleryImage) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"galleryImageName":  autorest.Encode("path", galleryImageName),
		"labAccountName":    autorest.Encode("path", labAccountName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-10-15"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LabServices/labaccounts/{labAccountName}/galleryimages/{galleryImageName}", pathParameters),
		autorest.WithJSON(galleryImage),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client GalleryImagesClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client GalleryImagesClient) CreateOrUpdateResponder(resp *http.Response) (result GalleryImage, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete delete gallery image.
// Parameters:
// resourceGroupName - the name of the resource group.
// labAccountName - the name of the lab Account.
// galleryImageName - the name of the gallery Image.
func (client GalleryImagesClient) Delete(ctx context.Context, resourceGroupName string, labAccountName string, galleryImageName string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/GalleryImagesClient.Delete")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, resourceGroupName, labAccountName, galleryImageName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "labservices.GalleryImagesClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "labservices.GalleryImagesClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "labservices.GalleryImagesClient", "Delete", resp, "Failure responding to request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client GalleryImagesClient) DeletePreparer(ctx context.Context, resourceGroupName string, labAccountName string, galleryImageName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"galleryImageName":  autorest.Encode("path", galleryImageName),
		"labAccountName":    autorest.Encode("path", labAccountName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-10-15"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LabServices/labaccounts/{labAccountName}/galleryimages/{galleryImageName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client GalleryImagesClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client GalleryImagesClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get get gallery image
// Parameters:
// resourceGroupName - the name of the resource group.
// labAccountName - the name of the lab Account.
// galleryImageName - the name of the gallery Image.
// expand - specify the $expand query. Example: 'properties($select=author)'
func (client GalleryImagesClient) Get(ctx context.Context, resourceGroupName string, labAccountName string, galleryImageName string, expand string) (result GalleryImage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/GalleryImagesClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, labAccountName, galleryImageName, expand)
	if err != nil {
		err = autorest.NewErrorWithError(err, "labservices.GalleryImagesClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "labservices.GalleryImagesClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "labservices.GalleryImagesClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client GalleryImagesClient) GetPreparer(ctx context.Context, resourceGroupName string, labAccountName string, galleryImageName string, expand string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"galleryImageName":  autorest.Encode("path", galleryImageName),
		"labAccountName":    autorest.Encode("path", labAccountName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-10-15"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(expand) > 0 {
		queryParameters["$expand"] = autorest.Encode("query", expand)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LabServices/labaccounts/{labAccountName}/galleryimages/{galleryImageName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client GalleryImagesClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client GalleryImagesClient) GetResponder(resp *http.Response) (result GalleryImage, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List list gallery images in a given lab account.
// Parameters:
// resourceGroupName - the name of the resource group.
// labAccountName - the name of the lab Account.
// expand - specify the $expand query. Example: 'properties($select=author)'
// filter - the filter to apply to the operation.
// top - the maximum number of resources to return from the operation.
// orderby - the ordering expression for the results, using OData notation.
func (client GalleryImagesClient) List(ctx context.Context, resourceGroupName string, labAccountName string, expand string, filter string, top *int32, orderby string) (result ResponseWithContinuationGalleryImagePage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/GalleryImagesClient.List")
		defer func() {
			sc := -1
			if result.rwcgi.Response.Response != nil {
				sc = result.rwcgi.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, resourceGroupName, labAccountName, expand, filter, top, orderby)
	if err != nil {
		err = autorest.NewErrorWithError(err, "labservices.GalleryImagesClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.rwcgi.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "labservices.GalleryImagesClient", "List", resp, "Failure sending request")
		return
	}

	result.rwcgi, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "labservices.GalleryImagesClient", "List", resp, "Failure responding to request")
		return
	}
	if result.rwcgi.hasNextLink() && result.rwcgi.IsEmpty() {
		err = result.NextWithContext(ctx)
	}

	return
}

// ListPreparer prepares the List request.
func (client GalleryImagesClient) ListPreparer(ctx context.Context, resourceGroupName string, labAccountName string, expand string, filter string, top *int32, orderby string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"labAccountName":    autorest.Encode("path", labAccountName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-10-15"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(expand) > 0 {
		queryParameters["$expand"] = autorest.Encode("query", expand)
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}
	if top != nil {
		queryParameters["$top"] = autorest.Encode("query", *top)
	}
	if len(orderby) > 0 {
		queryParameters["$orderby"] = autorest.Encode("query", orderby)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LabServices/labaccounts/{labAccountName}/galleryimages", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client GalleryImagesClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client GalleryImagesClient) ListResponder(resp *http.Response) (result ResponseWithContinuationGalleryImage, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client GalleryImagesClient) listNextResults(ctx context.Context, lastResults ResponseWithContinuationGalleryImage) (result ResponseWithContinuationGalleryImage, err error) {
	req, err := lastResults.responseWithContinuationGalleryImagePreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "labservices.GalleryImagesClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "labservices.GalleryImagesClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "labservices.GalleryImagesClient", "listNextResults", resp, "Failure responding to next results request")
		return
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client GalleryImagesClient) ListComplete(ctx context.Context, resourceGroupName string, labAccountName string, expand string, filter string, top *int32, orderby string) (result ResponseWithContinuationGalleryImageIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/GalleryImagesClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, resourceGroupName, labAccountName, expand, filter, top, orderby)
	return
}

// Update modify properties of gallery images.
// Parameters:
// resourceGroupName - the name of the resource group.
// labAccountName - the name of the lab Account.
// galleryImageName - the name of the gallery Image.
// galleryImage - represents an image from the Azure Marketplace
func (client GalleryImagesClient) Update(ctx context.Context, resourceGroupName string, labAccountName string, galleryImageName string, galleryImage GalleryImageFragment) (result GalleryImage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/GalleryImagesClient.Update")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.UpdatePreparer(ctx, resourceGroupName, labAccountName, galleryImageName, galleryImage)
	if err != nil {
		err = autorest.NewErrorWithError(err, "labservices.GalleryImagesClient", "Update", nil, "Failure preparing request")
		return
	}

	resp, err := client.UpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "labservices.GalleryImagesClient", "Update", resp, "Failure sending request")
		return
	}

	result, err = client.UpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "labservices.GalleryImagesClient", "Update", resp, "Failure responding to request")
		return
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client GalleryImagesClient) UpdatePreparer(ctx context.Context, resourceGroupName string, labAccountName string, galleryImageName string, galleryImage GalleryImageFragment) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"galleryImageName":  autorest.Encode("path", galleryImageName),
		"labAccountName":    autorest.Encode("path", labAccountName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-10-15"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LabServices/labaccounts/{labAccountName}/galleryimages/{galleryImageName}", pathParameters),
		autorest.WithJSON(galleryImage),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client GalleryImagesClient) UpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client GalleryImagesClient) UpdateResponder(resp *http.Response) (result GalleryImage, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
