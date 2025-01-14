//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdomainservices

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// OuContainerClient contains the methods for the OuContainer group.
// Don't use this type directly, use NewOuContainerClient() instead.
type OuContainerClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewOuContainerClient creates a new instance of OuContainerClient with the specified values.
// subscriptionID - Gets subscription credentials which uniquely identify the Microsoft Azure subscription. The subscription
// ID forms part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewOuContainerClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *OuContainerClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Endpoint) == 0 {
		cp.Endpoint = arm.AzurePublicCloud
	}
	client := &OuContainerClient{
		subscriptionID: subscriptionID,
		host:           string(cp.Endpoint),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, &cp),
	}
	return client
}

// BeginCreate - The Create OuContainer operation creates a new OuContainer under the specified Domain Service instance.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group within the user's subscription. The name is case insensitive.
// domainServiceName - The name of the domain service.
// ouContainerName - The name of the OuContainer.
// containerAccount - Container Account Description.
// options - OuContainerClientBeginCreateOptions contains the optional parameters for the OuContainerClient.BeginCreate method.
func (client *OuContainerClient) BeginCreate(ctx context.Context, resourceGroupName string, domainServiceName string, ouContainerName string, containerAccount ContainerAccount, options *OuContainerClientBeginCreateOptions) (OuContainerClientCreatePollerResponse, error) {
	resp, err := client.create(ctx, resourceGroupName, domainServiceName, ouContainerName, containerAccount, options)
	if err != nil {
		return OuContainerClientCreatePollerResponse{}, err
	}
	result := OuContainerClientCreatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("OuContainerClient.Create", "", resp, client.pl)
	if err != nil {
		return OuContainerClientCreatePollerResponse{}, err
	}
	result.Poller = &OuContainerClientCreatePoller{
		pt: pt,
	}
	return result, nil
}

// Create - The Create OuContainer operation creates a new OuContainer under the specified Domain Service instance.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *OuContainerClient) create(ctx context.Context, resourceGroupName string, domainServiceName string, ouContainerName string, containerAccount ContainerAccount, options *OuContainerClientBeginCreateOptions) (*http.Response, error) {
	req, err := client.createCreateRequest(ctx, resourceGroupName, domainServiceName, ouContainerName, containerAccount, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createCreateRequest creates the Create request.
func (client *OuContainerClient) createCreateRequest(ctx context.Context, resourceGroupName string, domainServiceName string, ouContainerName string, containerAccount ContainerAccount, options *OuContainerClientBeginCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Aad/domainServices/{domainServiceName}/ouContainer/{ouContainerName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if domainServiceName == "" {
		return nil, errors.New("parameter domainServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{domainServiceName}", url.PathEscape(domainServiceName))
	if ouContainerName == "" {
		return nil, errors.New("parameter ouContainerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ouContainerName}", url.PathEscape(ouContainerName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, containerAccount)
}

// BeginDelete - The Delete OuContainer operation deletes specified OuContainer.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group within the user's subscription. The name is case insensitive.
// domainServiceName - The name of the domain service.
// ouContainerName - The name of the OuContainer.
// options - OuContainerClientBeginDeleteOptions contains the optional parameters for the OuContainerClient.BeginDelete method.
func (client *OuContainerClient) BeginDelete(ctx context.Context, resourceGroupName string, domainServiceName string, ouContainerName string, options *OuContainerClientBeginDeleteOptions) (OuContainerClientDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, domainServiceName, ouContainerName, options)
	if err != nil {
		return OuContainerClientDeletePollerResponse{}, err
	}
	result := OuContainerClientDeletePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("OuContainerClient.Delete", "", resp, client.pl)
	if err != nil {
		return OuContainerClientDeletePollerResponse{}, err
	}
	result.Poller = &OuContainerClientDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - The Delete OuContainer operation deletes specified OuContainer.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *OuContainerClient) deleteOperation(ctx context.Context, resourceGroupName string, domainServiceName string, ouContainerName string, options *OuContainerClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, domainServiceName, ouContainerName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *OuContainerClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, domainServiceName string, ouContainerName string, options *OuContainerClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Aad/domainServices/{domainServiceName}/ouContainer/{ouContainerName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if domainServiceName == "" {
		return nil, errors.New("parameter domainServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{domainServiceName}", url.PathEscape(domainServiceName))
	if ouContainerName == "" {
		return nil, errors.New("parameter ouContainerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ouContainerName}", url.PathEscape(ouContainerName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Get - Get OuContainer in DomainService instance.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group within the user's subscription. The name is case insensitive.
// domainServiceName - The name of the domain service.
// ouContainerName - The name of the OuContainer.
// options - OuContainerClientGetOptions contains the optional parameters for the OuContainerClient.Get method.
func (client *OuContainerClient) Get(ctx context.Context, resourceGroupName string, domainServiceName string, ouContainerName string, options *OuContainerClientGetOptions) (OuContainerClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, domainServiceName, ouContainerName, options)
	if err != nil {
		return OuContainerClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return OuContainerClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return OuContainerClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *OuContainerClient) getCreateRequest(ctx context.Context, resourceGroupName string, domainServiceName string, ouContainerName string, options *OuContainerClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Aad/domainServices/{domainServiceName}/ouContainer/{ouContainerName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if domainServiceName == "" {
		return nil, errors.New("parameter domainServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{domainServiceName}", url.PathEscape(domainServiceName))
	if ouContainerName == "" {
		return nil, errors.New("parameter ouContainerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ouContainerName}", url.PathEscape(ouContainerName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *OuContainerClient) getHandleResponse(resp *http.Response) (OuContainerClientGetResponse, error) {
	result := OuContainerClientGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.OuContainer); err != nil {
		return OuContainerClientGetResponse{}, err
	}
	return result, nil
}

// List - The List of OuContainers in DomainService instance.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group within the user's subscription. The name is case insensitive.
// domainServiceName - The name of the domain service.
// options - OuContainerClientListOptions contains the optional parameters for the OuContainerClient.List method.
func (client *OuContainerClient) List(resourceGroupName string, domainServiceName string, options *OuContainerClientListOptions) *OuContainerClientListPager {
	return &OuContainerClientListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, resourceGroupName, domainServiceName, options)
		},
		advancer: func(ctx context.Context, resp OuContainerClientListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.OuContainerListResult.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *OuContainerClient) listCreateRequest(ctx context.Context, resourceGroupName string, domainServiceName string, options *OuContainerClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Aad/domainServices/{domainServiceName}/ouContainer"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if domainServiceName == "" {
		return nil, errors.New("parameter domainServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{domainServiceName}", url.PathEscape(domainServiceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *OuContainerClient) listHandleResponse(resp *http.Response) (OuContainerClientListResponse, error) {
	result := OuContainerClientListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.OuContainerListResult); err != nil {
		return OuContainerClientListResponse{}, err
	}
	return result, nil
}

// BeginUpdate - The Update OuContainer operation can be used to update the existing OuContainers.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group within the user's subscription. The name is case insensitive.
// domainServiceName - The name of the domain service.
// ouContainerName - The name of the OuContainer.
// containerAccount - Container Account Description.
// options - OuContainerClientBeginUpdateOptions contains the optional parameters for the OuContainerClient.BeginUpdate method.
func (client *OuContainerClient) BeginUpdate(ctx context.Context, resourceGroupName string, domainServiceName string, ouContainerName string, containerAccount ContainerAccount, options *OuContainerClientBeginUpdateOptions) (OuContainerClientUpdatePollerResponse, error) {
	resp, err := client.update(ctx, resourceGroupName, domainServiceName, ouContainerName, containerAccount, options)
	if err != nil {
		return OuContainerClientUpdatePollerResponse{}, err
	}
	result := OuContainerClientUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("OuContainerClient.Update", "", resp, client.pl)
	if err != nil {
		return OuContainerClientUpdatePollerResponse{}, err
	}
	result.Poller = &OuContainerClientUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// Update - The Update OuContainer operation can be used to update the existing OuContainers.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *OuContainerClient) update(ctx context.Context, resourceGroupName string, domainServiceName string, ouContainerName string, containerAccount ContainerAccount, options *OuContainerClientBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, domainServiceName, ouContainerName, containerAccount, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// updateCreateRequest creates the Update request.
func (client *OuContainerClient) updateCreateRequest(ctx context.Context, resourceGroupName string, domainServiceName string, ouContainerName string, containerAccount ContainerAccount, options *OuContainerClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Aad/domainServices/{domainServiceName}/ouContainer/{ouContainerName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if domainServiceName == "" {
		return nil, errors.New("parameter domainServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{domainServiceName}", url.PathEscape(domainServiceName))
	if ouContainerName == "" {
		return nil, errors.New("parameter ouContainerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ouContainerName}", url.PathEscape(ouContainerName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, containerAccount)
}
