package programmatic

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
	"github.com/satori/go.uuid"
	"net/http"
)

// FeaturesClient is the client for the Features methods of the Programmatic service.
type FeaturesClient struct {
	BaseClient
}

// NewFeaturesClient creates an instance of the FeaturesClient client.
func NewFeaturesClient(azureRegion AzureRegions) FeaturesClient {
	return FeaturesClient{New(azureRegion)}
}

// AddPhraseList creates a new phraselist feature.
// Parameters:
// appID - the application ID.
// versionID - the version ID.
// phraselistCreateObject - a Phraselist object containing Name, comma-separated Phrases and the isExchangeable
// boolean. Default value for isExchangeable is true.
func (client FeaturesClient) AddPhraseList(ctx context.Context, appID uuid.UUID, versionID string, phraselistCreateObject PhraselistCreateObject) (result Int32, err error) {
	req, err := client.AddPhraseListPreparer(ctx, appID, versionID, phraselistCreateObject)
	if err != nil {
		err = autorest.NewErrorWithError(err, "programmatic.FeaturesClient", "AddPhraseList", nil, "Failure preparing request")
		return
	}

	resp, err := client.AddPhraseListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "programmatic.FeaturesClient", "AddPhraseList", resp, "Failure sending request")
		return
	}

	result, err = client.AddPhraseListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "programmatic.FeaturesClient", "AddPhraseList", resp, "Failure responding to request")
		return
	}

	return
}

// AddPhraseListPreparer prepares the AddPhraseList request.
func (client FeaturesClient) AddPhraseListPreparer(ctx context.Context, appID uuid.UUID, versionID string, phraselistCreateObject PhraselistCreateObject) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"AzureRegion": client.AzureRegion,
	}

	pathParameters := map[string]interface{}{
		"appId":     autorest.Encode("path", appID),
		"versionId": autorest.Encode("path", versionID),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithCustomBaseURL("https://{AzureRegion}.api.cognitive.microsoft.com/luis/api/v2.0", urlParameters),
		autorest.WithPathParameters("/apps/{appId}/versions/{versionId}/phraselists", pathParameters),
		autorest.WithJSON(phraselistCreateObject))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// AddPhraseListSender sends the AddPhraseList request. The method will close the
// http.Response Body if it receives an error.
func (client FeaturesClient) AddPhraseListSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// AddPhraseListResponder handles the response to the AddPhraseList request. The method always
// closes the http.Response Body.
func (client FeaturesClient) AddPhraseListResponder(resp *http.Response) (result Int32, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// DeletePhraseList deletes a phraselist feature.
// Parameters:
// appID - the application ID.
// versionID - the version ID.
// phraselistID - the ID of the feature to be deleted.
func (client FeaturesClient) DeletePhraseList(ctx context.Context, appID uuid.UUID, versionID string, phraselistID int32) (result OperationStatus, err error) {
	req, err := client.DeletePhraseListPreparer(ctx, appID, versionID, phraselistID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "programmatic.FeaturesClient", "DeletePhraseList", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeletePhraseListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "programmatic.FeaturesClient", "DeletePhraseList", resp, "Failure sending request")
		return
	}

	result, err = client.DeletePhraseListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "programmatic.FeaturesClient", "DeletePhraseList", resp, "Failure responding to request")
		return
	}

	return
}

// DeletePhraseListPreparer prepares the DeletePhraseList request.
func (client FeaturesClient) DeletePhraseListPreparer(ctx context.Context, appID uuid.UUID, versionID string, phraselistID int32) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"AzureRegion": client.AzureRegion,
	}

	pathParameters := map[string]interface{}{
		"appId":        autorest.Encode("path", appID),
		"phraselistId": autorest.Encode("path", phraselistID),
		"versionId":    autorest.Encode("path", versionID),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithCustomBaseURL("https://{AzureRegion}.api.cognitive.microsoft.com/luis/api/v2.0", urlParameters),
		autorest.WithPathParameters("/apps/{appId}/versions/{versionId}/phraselists/{phraselistId}", pathParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeletePhraseListSender sends the DeletePhraseList request. The method will close the
// http.Response Body if it receives an error.
func (client FeaturesClient) DeletePhraseListSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// DeletePhraseListResponder handles the response to the DeletePhraseList request. The method always
// closes the http.Response Body.
func (client FeaturesClient) DeletePhraseListResponder(resp *http.Response) (result OperationStatus, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetPhraseList gets phraselist feature info.
// Parameters:
// appID - the application ID.
// versionID - the version ID.
// phraselistID - the ID of the feature to be retrieved.
func (client FeaturesClient) GetPhraseList(ctx context.Context, appID uuid.UUID, versionID string, phraselistID int32) (result PhraseListFeatureInfo, err error) {
	req, err := client.GetPhraseListPreparer(ctx, appID, versionID, phraselistID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "programmatic.FeaturesClient", "GetPhraseList", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetPhraseListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "programmatic.FeaturesClient", "GetPhraseList", resp, "Failure sending request")
		return
	}

	result, err = client.GetPhraseListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "programmatic.FeaturesClient", "GetPhraseList", resp, "Failure responding to request")
		return
	}

	return
}

// GetPhraseListPreparer prepares the GetPhraseList request.
func (client FeaturesClient) GetPhraseListPreparer(ctx context.Context, appID uuid.UUID, versionID string, phraselistID int32) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"AzureRegion": client.AzureRegion,
	}

	pathParameters := map[string]interface{}{
		"appId":        autorest.Encode("path", appID),
		"phraselistId": autorest.Encode("path", phraselistID),
		"versionId":    autorest.Encode("path", versionID),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithCustomBaseURL("https://{AzureRegion}.api.cognitive.microsoft.com/luis/api/v2.0", urlParameters),
		autorest.WithPathParameters("/apps/{appId}/versions/{versionId}/phraselists/{phraselistId}", pathParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetPhraseListSender sends the GetPhraseList request. The method will close the
// http.Response Body if it receives an error.
func (client FeaturesClient) GetPhraseListSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetPhraseListResponder handles the response to the GetPhraseList request. The method always
// closes the http.Response Body.
func (client FeaturesClient) GetPhraseListResponder(resp *http.Response) (result PhraseListFeatureInfo, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List gets all the extraction features for the specified application version.
// Parameters:
// appID - the application ID.
// versionID - the version ID.
// skip - the number of entries to skip. Default value is 0.
// take - the number of entries to return. Maximum page size is 500. Default is 100.
func (client FeaturesClient) List(ctx context.Context, appID uuid.UUID, versionID string, skip *int32, take *int32) (result FeaturesResponseObject, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: skip,
			Constraints: []validation.Constraint{{Target: "skip", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "skip", Name: validation.InclusiveMinimum, Rule: 0, Chain: nil}}}}},
		{TargetValue: take,
			Constraints: []validation.Constraint{{Target: "take", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "take", Name: validation.InclusiveMaximum, Rule: 500, Chain: nil},
					{Target: "take", Name: validation.InclusiveMinimum, Rule: 0, Chain: nil},
				}}}}}); err != nil {
		return result, validation.NewError("programmatic.FeaturesClient", "List", err.Error())
	}

	req, err := client.ListPreparer(ctx, appID, versionID, skip, take)
	if err != nil {
		err = autorest.NewErrorWithError(err, "programmatic.FeaturesClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "programmatic.FeaturesClient", "List", resp, "Failure sending request")
		return
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "programmatic.FeaturesClient", "List", resp, "Failure responding to request")
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client FeaturesClient) ListPreparer(ctx context.Context, appID uuid.UUID, versionID string, skip *int32, take *int32) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"AzureRegion": client.AzureRegion,
	}

	pathParameters := map[string]interface{}{
		"appId":     autorest.Encode("path", appID),
		"versionId": autorest.Encode("path", versionID),
	}

	queryParameters := map[string]interface{}{}
	if skip != nil {
		queryParameters["skip"] = autorest.Encode("query", *skip)
	} else {
		queryParameters["skip"] = autorest.Encode("query", 0)
	}
	if take != nil {
		queryParameters["take"] = autorest.Encode("query", *take)
	} else {
		queryParameters["take"] = autorest.Encode("query", 100)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithCustomBaseURL("https://{AzureRegion}.api.cognitive.microsoft.com/luis/api/v2.0", urlParameters),
		autorest.WithPathParameters("/apps/{appId}/versions/{versionId}/features", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client FeaturesClient) ListSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client FeaturesClient) ListResponder(resp *http.Response) (result FeaturesResponseObject, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListPhraseLists gets all the phraselist features.
// Parameters:
// appID - the application ID.
// versionID - the version ID.
// skip - the number of entries to skip. Default value is 0.
// take - the number of entries to return. Maximum page size is 500. Default is 100.
func (client FeaturesClient) ListPhraseLists(ctx context.Context, appID uuid.UUID, versionID string, skip *int32, take *int32) (result ListPhraseListFeatureInfo, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: skip,
			Constraints: []validation.Constraint{{Target: "skip", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "skip", Name: validation.InclusiveMinimum, Rule: 0, Chain: nil}}}}},
		{TargetValue: take,
			Constraints: []validation.Constraint{{Target: "take", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "take", Name: validation.InclusiveMaximum, Rule: 500, Chain: nil},
					{Target: "take", Name: validation.InclusiveMinimum, Rule: 0, Chain: nil},
				}}}}}); err != nil {
		return result, validation.NewError("programmatic.FeaturesClient", "ListPhraseLists", err.Error())
	}

	req, err := client.ListPhraseListsPreparer(ctx, appID, versionID, skip, take)
	if err != nil {
		err = autorest.NewErrorWithError(err, "programmatic.FeaturesClient", "ListPhraseLists", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListPhraseListsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "programmatic.FeaturesClient", "ListPhraseLists", resp, "Failure sending request")
		return
	}

	result, err = client.ListPhraseListsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "programmatic.FeaturesClient", "ListPhraseLists", resp, "Failure responding to request")
		return
	}

	return
}

// ListPhraseListsPreparer prepares the ListPhraseLists request.
func (client FeaturesClient) ListPhraseListsPreparer(ctx context.Context, appID uuid.UUID, versionID string, skip *int32, take *int32) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"AzureRegion": client.AzureRegion,
	}

	pathParameters := map[string]interface{}{
		"appId":     autorest.Encode("path", appID),
		"versionId": autorest.Encode("path", versionID),
	}

	queryParameters := map[string]interface{}{}
	if skip != nil {
		queryParameters["skip"] = autorest.Encode("query", *skip)
	} else {
		queryParameters["skip"] = autorest.Encode("query", 0)
	}
	if take != nil {
		queryParameters["take"] = autorest.Encode("query", *take)
	} else {
		queryParameters["take"] = autorest.Encode("query", 100)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithCustomBaseURL("https://{AzureRegion}.api.cognitive.microsoft.com/luis/api/v2.0", urlParameters),
		autorest.WithPathParameters("/apps/{appId}/versions/{versionId}/phraselists", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListPhraseListsSender sends the ListPhraseLists request. The method will close the
// http.Response Body if it receives an error.
func (client FeaturesClient) ListPhraseListsSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListPhraseListsResponder handles the response to the ListPhraseLists request. The method always
// closes the http.Response Body.
func (client FeaturesClient) ListPhraseListsResponder(resp *http.Response) (result ListPhraseListFeatureInfo, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// UpdatePhraseList updates the phrases, the state and the name of the phraselist feature.
// Parameters:
// appID - the application ID.
// versionID - the version ID.
// phraselistID - the ID of the feature to be updated.
// phraselistUpdateObject - the new values for: - Just a boolean called IsActive, in which case the status of
// the feature will be changed. - Name, Pattern, Mode, and a boolean called IsActive to update the feature.
func (client FeaturesClient) UpdatePhraseList(ctx context.Context, appID uuid.UUID, versionID string, phraselistID int32, phraselistUpdateObject *PhraselistUpdateObject) (result OperationStatus, err error) {
	req, err := client.UpdatePhraseListPreparer(ctx, appID, versionID, phraselistID, phraselistUpdateObject)
	if err != nil {
		err = autorest.NewErrorWithError(err, "programmatic.FeaturesClient", "UpdatePhraseList", nil, "Failure preparing request")
		return
	}

	resp, err := client.UpdatePhraseListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "programmatic.FeaturesClient", "UpdatePhraseList", resp, "Failure sending request")
		return
	}

	result, err = client.UpdatePhraseListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "programmatic.FeaturesClient", "UpdatePhraseList", resp, "Failure responding to request")
		return
	}

	return
}

// UpdatePhraseListPreparer prepares the UpdatePhraseList request.
func (client FeaturesClient) UpdatePhraseListPreparer(ctx context.Context, appID uuid.UUID, versionID string, phraselistID int32, phraselistUpdateObject *PhraselistUpdateObject) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"AzureRegion": client.AzureRegion,
	}

	pathParameters := map[string]interface{}{
		"appId":        autorest.Encode("path", appID),
		"phraselistId": autorest.Encode("path", phraselistID),
		"versionId":    autorest.Encode("path", versionID),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithCustomBaseURL("https://{AzureRegion}.api.cognitive.microsoft.com/luis/api/v2.0", urlParameters),
		autorest.WithPathParameters("/apps/{appId}/versions/{versionId}/phraselists/{phraselistId}", pathParameters))
	if phraselistUpdateObject != nil {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithJSON(phraselistUpdateObject))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdatePhraseListSender sends the UpdatePhraseList request. The method will close the
// http.Response Body if it receives an error.
func (client FeaturesClient) UpdatePhraseListSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// UpdatePhraseListResponder handles the response to the UpdatePhraseList request. The method always
// closes the http.Response Body.
func (client FeaturesClient) UpdatePhraseListResponder(resp *http.Response) (result OperationStatus, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
