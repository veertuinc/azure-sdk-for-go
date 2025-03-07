//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdatalakeanalytics

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

// StorageAccountsClient contains the methods for the StorageAccounts group.
// Don't use this type directly, use NewStorageAccountsClient() instead.
type StorageAccountsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewStorageAccountsClient creates a new instance of StorageAccountsClient with the specified values.
// subscriptionID - Get subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID
// forms part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewStorageAccountsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*StorageAccountsClient, error) {
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
	client := &StorageAccountsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// Add - Updates the specified Data Lake Analytics account to add an Azure Storage account.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2019-11-01-preview
// resourceGroupName - The name of the Azure resource group.
// accountName - The name of the Data Lake Analytics account.
// storageAccountName - The name of the Azure Storage account to add
// parameters - The parameters containing the access key and optional suffix for the Azure Storage Account.
// options - StorageAccountsClientAddOptions contains the optional parameters for the StorageAccountsClient.Add method.
func (client *StorageAccountsClient) Add(ctx context.Context, resourceGroupName string, accountName string, storageAccountName string, parameters AddStorageAccountParameters, options *StorageAccountsClientAddOptions) (StorageAccountsClientAddResponse, error) {
	req, err := client.addCreateRequest(ctx, resourceGroupName, accountName, storageAccountName, parameters, options)
	if err != nil {
		return StorageAccountsClientAddResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return StorageAccountsClientAddResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return StorageAccountsClientAddResponse{}, runtime.NewResponseError(resp)
	}
	return StorageAccountsClientAddResponse{}, nil
}

// addCreateRequest creates the Add request.
func (client *StorageAccountsClient) addCreateRequest(ctx context.Context, resourceGroupName string, accountName string, storageAccountName string, parameters AddStorageAccountParameters, options *StorageAccountsClientAddOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeAnalytics/accounts/{accountName}/storageAccounts/{storageAccountName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if storageAccountName == "" {
		return nil, errors.New("parameter storageAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{storageAccountName}", url.PathEscape(storageAccountName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// Delete - Updates the specified Data Lake Analytics account to remove an Azure Storage account.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2019-11-01-preview
// resourceGroupName - The name of the Azure resource group.
// accountName - The name of the Data Lake Analytics account.
// storageAccountName - The name of the Azure Storage account to remove
// options - StorageAccountsClientDeleteOptions contains the optional parameters for the StorageAccountsClient.Delete method.
func (client *StorageAccountsClient) Delete(ctx context.Context, resourceGroupName string, accountName string, storageAccountName string, options *StorageAccountsClientDeleteOptions) (StorageAccountsClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, accountName, storageAccountName, options)
	if err != nil {
		return StorageAccountsClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return StorageAccountsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return StorageAccountsClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return StorageAccountsClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *StorageAccountsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, accountName string, storageAccountName string, options *StorageAccountsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeAnalytics/accounts/{accountName}/storageAccounts/{storageAccountName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if storageAccountName == "" {
		return nil, errors.New("parameter storageAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{storageAccountName}", url.PathEscape(storageAccountName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets the specified Azure Storage account linked to the given Data Lake Analytics account.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2019-11-01-preview
// resourceGroupName - The name of the Azure resource group.
// accountName - The name of the Data Lake Analytics account.
// storageAccountName - The name of the Azure Storage account for which to retrieve the details.
// options - StorageAccountsClientGetOptions contains the optional parameters for the StorageAccountsClient.Get method.
func (client *StorageAccountsClient) Get(ctx context.Context, resourceGroupName string, accountName string, storageAccountName string, options *StorageAccountsClientGetOptions) (StorageAccountsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, accountName, storageAccountName, options)
	if err != nil {
		return StorageAccountsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return StorageAccountsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return StorageAccountsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *StorageAccountsClient) getCreateRequest(ctx context.Context, resourceGroupName string, accountName string, storageAccountName string, options *StorageAccountsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeAnalytics/accounts/{accountName}/storageAccounts/{storageAccountName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if storageAccountName == "" {
		return nil, errors.New("parameter storageAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{storageAccountName}", url.PathEscape(storageAccountName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *StorageAccountsClient) getHandleResponse(resp *http.Response) (StorageAccountsClientGetResponse, error) {
	result := StorageAccountsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.StorageAccountInformation); err != nil {
		return StorageAccountsClientGetResponse{}, err
	}
	return result, nil
}

// GetStorageContainer - Gets the specified Azure Storage container associated with the given Data Lake Analytics and Azure
// Storage accounts.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2019-11-01-preview
// resourceGroupName - The name of the Azure resource group.
// accountName - The name of the Data Lake Analytics account.
// storageAccountName - The name of the Azure storage account from which to retrieve the blob container.
// containerName - The name of the Azure storage container to retrieve
// options - StorageAccountsClientGetStorageContainerOptions contains the optional parameters for the StorageAccountsClient.GetStorageContainer
// method.
func (client *StorageAccountsClient) GetStorageContainer(ctx context.Context, resourceGroupName string, accountName string, storageAccountName string, containerName string, options *StorageAccountsClientGetStorageContainerOptions) (StorageAccountsClientGetStorageContainerResponse, error) {
	req, err := client.getStorageContainerCreateRequest(ctx, resourceGroupName, accountName, storageAccountName, containerName, options)
	if err != nil {
		return StorageAccountsClientGetStorageContainerResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return StorageAccountsClientGetStorageContainerResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return StorageAccountsClientGetStorageContainerResponse{}, runtime.NewResponseError(resp)
	}
	return client.getStorageContainerHandleResponse(resp)
}

// getStorageContainerCreateRequest creates the GetStorageContainer request.
func (client *StorageAccountsClient) getStorageContainerCreateRequest(ctx context.Context, resourceGroupName string, accountName string, storageAccountName string, containerName string, options *StorageAccountsClientGetStorageContainerOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeAnalytics/accounts/{accountName}/storageAccounts/{storageAccountName}/containers/{containerName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if storageAccountName == "" {
		return nil, errors.New("parameter storageAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{storageAccountName}", url.PathEscape(storageAccountName))
	if containerName == "" {
		return nil, errors.New("parameter containerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{containerName}", url.PathEscape(containerName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getStorageContainerHandleResponse handles the GetStorageContainer response.
func (client *StorageAccountsClient) getStorageContainerHandleResponse(resp *http.Response) (StorageAccountsClientGetStorageContainerResponse, error) {
	result := StorageAccountsClientGetStorageContainerResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.StorageContainer); err != nil {
		return StorageAccountsClientGetStorageContainerResponse{}, err
	}
	return result, nil
}

// NewListByAccountPager - Gets the first page of Azure Storage accounts, if any, linked to the specified Data Lake Analytics
// account. The response includes a link to the next page, if any.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2019-11-01-preview
// resourceGroupName - The name of the Azure resource group.
// accountName - The name of the Data Lake Analytics account.
// options - StorageAccountsClientListByAccountOptions contains the optional parameters for the StorageAccountsClient.ListByAccount
// method.
func (client *StorageAccountsClient) NewListByAccountPager(resourceGroupName string, accountName string, options *StorageAccountsClientListByAccountOptions) *runtime.Pager[StorageAccountsClientListByAccountResponse] {
	return runtime.NewPager(runtime.PagingHandler[StorageAccountsClientListByAccountResponse]{
		More: func(page StorageAccountsClientListByAccountResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *StorageAccountsClientListByAccountResponse) (StorageAccountsClientListByAccountResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByAccountCreateRequest(ctx, resourceGroupName, accountName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return StorageAccountsClientListByAccountResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return StorageAccountsClientListByAccountResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return StorageAccountsClientListByAccountResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByAccountHandleResponse(resp)
		},
	})
}

// listByAccountCreateRequest creates the ListByAccount request.
func (client *StorageAccountsClient) listByAccountCreateRequest(ctx context.Context, resourceGroupName string, accountName string, options *StorageAccountsClientListByAccountOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeAnalytics/accounts/{accountName}/storageAccounts"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	if options != nil && options.Skip != nil {
		reqQP.Set("$skip", strconv.FormatInt(int64(*options.Skip), 10))
	}
	if options != nil && options.Select != nil {
		reqQP.Set("$select", *options.Select)
	}
	if options != nil && options.Orderby != nil {
		reqQP.Set("$orderby", *options.Orderby)
	}
	if options != nil && options.Count != nil {
		reqQP.Set("$count", strconv.FormatBool(*options.Count))
	}
	reqQP.Set("api-version", "2019-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByAccountHandleResponse handles the ListByAccount response.
func (client *StorageAccountsClient) listByAccountHandleResponse(resp *http.Response) (StorageAccountsClientListByAccountResponse, error) {
	result := StorageAccountsClientListByAccountResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.StorageAccountInformationListResult); err != nil {
		return StorageAccountsClientListByAccountResponse{}, err
	}
	return result, nil
}

// NewListSasTokensPager - Gets the SAS token associated with the specified Data Lake Analytics and Azure Storage account
// and container combination.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2019-11-01-preview
// resourceGroupName - The name of the Azure resource group.
// accountName - The name of the Data Lake Analytics account.
// storageAccountName - The name of the Azure storage account for which the SAS token is being requested.
// containerName - The name of the Azure storage container for which the SAS token is being requested.
// options - StorageAccountsClientListSasTokensOptions contains the optional parameters for the StorageAccountsClient.ListSasTokens
// method.
func (client *StorageAccountsClient) NewListSasTokensPager(resourceGroupName string, accountName string, storageAccountName string, containerName string, options *StorageAccountsClientListSasTokensOptions) *runtime.Pager[StorageAccountsClientListSasTokensResponse] {
	return runtime.NewPager(runtime.PagingHandler[StorageAccountsClientListSasTokensResponse]{
		More: func(page StorageAccountsClientListSasTokensResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *StorageAccountsClientListSasTokensResponse) (StorageAccountsClientListSasTokensResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listSasTokensCreateRequest(ctx, resourceGroupName, accountName, storageAccountName, containerName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return StorageAccountsClientListSasTokensResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return StorageAccountsClientListSasTokensResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return StorageAccountsClientListSasTokensResponse{}, runtime.NewResponseError(resp)
			}
			return client.listSasTokensHandleResponse(resp)
		},
	})
}

// listSasTokensCreateRequest creates the ListSasTokens request.
func (client *StorageAccountsClient) listSasTokensCreateRequest(ctx context.Context, resourceGroupName string, accountName string, storageAccountName string, containerName string, options *StorageAccountsClientListSasTokensOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeAnalytics/accounts/{accountName}/storageAccounts/{storageAccountName}/containers/{containerName}/listSasTokens"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if storageAccountName == "" {
		return nil, errors.New("parameter storageAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{storageAccountName}", url.PathEscape(storageAccountName))
	if containerName == "" {
		return nil, errors.New("parameter containerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{containerName}", url.PathEscape(containerName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listSasTokensHandleResponse handles the ListSasTokens response.
func (client *StorageAccountsClient) listSasTokensHandleResponse(resp *http.Response) (StorageAccountsClientListSasTokensResponse, error) {
	result := StorageAccountsClientListSasTokensResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SasTokenInformationListResult); err != nil {
		return StorageAccountsClientListSasTokensResponse{}, err
	}
	return result, nil
}

// NewListStorageContainersPager - Lists the Azure Storage containers, if any, associated with the specified Data Lake Analytics
// and Azure Storage account combination. The response includes a link to the next page of results, if any.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2019-11-01-preview
// resourceGroupName - The name of the Azure resource group.
// accountName - The name of the Data Lake Analytics account.
// storageAccountName - The name of the Azure storage account from which to list blob containers.
// options - StorageAccountsClientListStorageContainersOptions contains the optional parameters for the StorageAccountsClient.ListStorageContainers
// method.
func (client *StorageAccountsClient) NewListStorageContainersPager(resourceGroupName string, accountName string, storageAccountName string, options *StorageAccountsClientListStorageContainersOptions) *runtime.Pager[StorageAccountsClientListStorageContainersResponse] {
	return runtime.NewPager(runtime.PagingHandler[StorageAccountsClientListStorageContainersResponse]{
		More: func(page StorageAccountsClientListStorageContainersResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *StorageAccountsClientListStorageContainersResponse) (StorageAccountsClientListStorageContainersResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listStorageContainersCreateRequest(ctx, resourceGroupName, accountName, storageAccountName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return StorageAccountsClientListStorageContainersResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return StorageAccountsClientListStorageContainersResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return StorageAccountsClientListStorageContainersResponse{}, runtime.NewResponseError(resp)
			}
			return client.listStorageContainersHandleResponse(resp)
		},
	})
}

// listStorageContainersCreateRequest creates the ListStorageContainers request.
func (client *StorageAccountsClient) listStorageContainersCreateRequest(ctx context.Context, resourceGroupName string, accountName string, storageAccountName string, options *StorageAccountsClientListStorageContainersOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeAnalytics/accounts/{accountName}/storageAccounts/{storageAccountName}/containers"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if storageAccountName == "" {
		return nil, errors.New("parameter storageAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{storageAccountName}", url.PathEscape(storageAccountName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listStorageContainersHandleResponse handles the ListStorageContainers response.
func (client *StorageAccountsClient) listStorageContainersHandleResponse(resp *http.Response) (StorageAccountsClientListStorageContainersResponse, error) {
	result := StorageAccountsClientListStorageContainersResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.StorageContainerListResult); err != nil {
		return StorageAccountsClientListStorageContainersResponse{}, err
	}
	return result, nil
}

// Update - Updates the Data Lake Analytics account to replace Azure Storage blob account details, such as the access key
// and/or suffix.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2019-11-01-preview
// resourceGroupName - The name of the Azure resource group.
// accountName - The name of the Data Lake Analytics account.
// storageAccountName - The Azure Storage account to modify
// options - StorageAccountsClientUpdateOptions contains the optional parameters for the StorageAccountsClient.Update method.
func (client *StorageAccountsClient) Update(ctx context.Context, resourceGroupName string, accountName string, storageAccountName string, options *StorageAccountsClientUpdateOptions) (StorageAccountsClientUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, accountName, storageAccountName, options)
	if err != nil {
		return StorageAccountsClientUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return StorageAccountsClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return StorageAccountsClientUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return StorageAccountsClientUpdateResponse{}, nil
}

// updateCreateRequest creates the Update request.
func (client *StorageAccountsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, accountName string, storageAccountName string, options *StorageAccountsClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeAnalytics/accounts/{accountName}/storageAccounts/{storageAccountName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if storageAccountName == "" {
		return nil, errors.New("parameter storageAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{storageAccountName}", url.PathEscape(storageAccountName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if options != nil && options.Parameters != nil {
		return req, runtime.MarshalAsJSON(req, *options.Parameters)
	}
	return req, nil
}
