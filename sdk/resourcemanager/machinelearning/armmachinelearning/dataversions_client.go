//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armmachinelearning

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
	"strconv"
	"strings"
)

// DataVersionsClient contains the methods for the DataVersions group.
// Don't use this type directly, use NewDataVersionsClient() instead.
type DataVersionsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewDataVersionsClient creates a new instance of DataVersionsClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewDataVersionsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*DataVersionsClient, error) {
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
	client := &DataVersionsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// CreateOrUpdate - Create or update version.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-10-01
// resourceGroupName - The name of the resource group. The name is case insensitive.
// workspaceName - Name of Azure Machine Learning workspace.
// name - Container name.
// version - Version identifier.
// body - Version entity to create or update.
// options - DataVersionsClientCreateOrUpdateOptions contains the optional parameters for the DataVersionsClient.CreateOrUpdate
// method.
func (client *DataVersionsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, name string, version string, body DataVersionBase, options *DataVersionsClientCreateOrUpdateOptions) (DataVersionsClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, workspaceName, name, version, body, options)
	if err != nil {
		return DataVersionsClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DataVersionsClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return DataVersionsClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *DataVersionsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, name string, version string, body DataVersionBase, options *DataVersionsClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}/data/{name}/versions/{version}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	if version == "" {
		return nil, errors.New("parameter version cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{version}", url.PathEscape(version))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, body)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *DataVersionsClient) createOrUpdateHandleResponse(resp *http.Response) (DataVersionsClientCreateOrUpdateResponse, error) {
	result := DataVersionsClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DataVersionBase); err != nil {
		return DataVersionsClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Delete version.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-10-01
// resourceGroupName - The name of the resource group. The name is case insensitive.
// workspaceName - Name of Azure Machine Learning workspace.
// name - Container name.
// version - Version identifier.
// options - DataVersionsClientDeleteOptions contains the optional parameters for the DataVersionsClient.Delete method.
func (client *DataVersionsClient) Delete(ctx context.Context, resourceGroupName string, workspaceName string, name string, version string, options *DataVersionsClientDeleteOptions) (DataVersionsClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, workspaceName, name, version, options)
	if err != nil {
		return DataVersionsClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DataVersionsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return DataVersionsClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return DataVersionsClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *DataVersionsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, name string, version string, options *DataVersionsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}/data/{name}/versions/{version}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	if version == "" {
		return nil, errors.New("parameter version cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{version}", url.PathEscape(version))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get version.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-10-01
// resourceGroupName - The name of the resource group. The name is case insensitive.
// workspaceName - Name of Azure Machine Learning workspace.
// name - Container name.
// version - Version identifier.
// options - DataVersionsClientGetOptions contains the optional parameters for the DataVersionsClient.Get method.
func (client *DataVersionsClient) Get(ctx context.Context, resourceGroupName string, workspaceName string, name string, version string, options *DataVersionsClientGetOptions) (DataVersionsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, workspaceName, name, version, options)
	if err != nil {
		return DataVersionsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DataVersionsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DataVersionsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *DataVersionsClient) getCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, name string, version string, options *DataVersionsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}/data/{name}/versions/{version}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	if version == "" {
		return nil, errors.New("parameter version cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{version}", url.PathEscape(version))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *DataVersionsClient) getHandleResponse(resp *http.Response) (DataVersionsClientGetResponse, error) {
	result := DataVersionsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DataVersionBase); err != nil {
		return DataVersionsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - List data versions in the data container
// Generated from API version 2022-10-01
// resourceGroupName - The name of the resource group. The name is case insensitive.
// workspaceName - Name of Azure Machine Learning workspace.
// name - Data container's name
// options - DataVersionsClientListOptions contains the optional parameters for the DataVersionsClient.List method.
func (client *DataVersionsClient) NewListPager(resourceGroupName string, workspaceName string, name string, options *DataVersionsClientListOptions) *runtime.Pager[DataVersionsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[DataVersionsClientListResponse]{
		More: func(page DataVersionsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *DataVersionsClientListResponse) (DataVersionsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, workspaceName, name, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return DataVersionsClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return DataVersionsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return DataVersionsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *DataVersionsClient) listCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, name string, options *DataVersionsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}/data/{name}/versions"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-10-01")
	if options != nil && options.OrderBy != nil {
		reqQP.Set("$orderBy", *options.OrderBy)
	}
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	if options != nil && options.Skip != nil {
		reqQP.Set("$skip", *options.Skip)
	}
	if options != nil && options.Tags != nil {
		reqQP.Set("$tags", *options.Tags)
	}
	if options != nil && options.ListViewType != nil {
		reqQP.Set("listViewType", string(*options.ListViewType))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *DataVersionsClient) listHandleResponse(resp *http.Response) (DataVersionsClientListResponse, error) {
	result := DataVersionsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DataVersionBaseResourceArmPaginatedResult); err != nil {
		return DataVersionsClientListResponse{}, err
	}
	return result, nil
}
