//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// Code generated by Microsoft (R) AutoRest Code Generator.Changes may cause incorrect behavior and will be lost if the code
// is regenerated.
// DO NOT EDIT.

package armelastic

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// AllTrafficFiltersClient contains the methods for the AllTrafficFilters group.
// Don't use this type directly, use NewAllTrafficFiltersClient() instead.
type AllTrafficFiltersClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewAllTrafficFiltersClient creates a new instance of AllTrafficFiltersClient with the specified values.
// subscriptionID - The Azure subscription ID. This is a GUID-formatted string (e.g. 00000000-0000-0000-0000-000000000000)
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewAllTrafficFiltersClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*AllTrafficFiltersClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublic.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &AllTrafficFiltersClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// List - Get the list of all traffic filters for the account.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-07-01-preview
// resourceGroupName - The name of the resource group to which the Elastic resource belongs.
// monitorName - Monitor resource name
// options - AllTrafficFiltersClientListOptions contains the optional parameters for the AllTrafficFiltersClient.List method.
func (client *AllTrafficFiltersClient) List(ctx context.Context, resourceGroupName string, monitorName string, options *AllTrafficFiltersClientListOptions) (AllTrafficFiltersClientListResponse, error) {
	req, err := client.listCreateRequest(ctx, resourceGroupName, monitorName, options)
	if err != nil {
		return AllTrafficFiltersClientListResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AllTrafficFiltersClientListResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return AllTrafficFiltersClientListResponse{}, runtime.NewResponseError(resp)
	}
	return client.listHandleResponse(resp)
}

// listCreateRequest creates the List request.
func (client *AllTrafficFiltersClient) listCreateRequest(ctx context.Context, resourceGroupName string, monitorName string, options *AllTrafficFiltersClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Elastic/monitors/{monitorName}/listAllTrafficFilters"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if monitorName == "" {
		return nil, errors.New("parameter monitorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{monitorName}", url.PathEscape(monitorName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-07-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *AllTrafficFiltersClient) listHandleResponse(resp *http.Response) (AllTrafficFiltersClientListResponse, error) {
	result := AllTrafficFiltersClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.TrafficFilterResponse); err != nil {
		return AllTrafficFiltersClientListResponse{}, err
	}
	return result, nil
}
