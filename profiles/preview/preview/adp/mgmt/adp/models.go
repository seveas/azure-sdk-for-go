// +build go1.9

// Copyright 2020 Microsoft Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/tools/profileBuilder

package adp

import (
	"context"

	original "github.com/Azure/azure-sdk-for-go/services/preview/adp/mgmt/2019-07-01-preview/adp"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type CreatedByType = original.CreatedByType

const (
	Application     CreatedByType = original.Application
	Key             CreatedByType = original.Key
	ManagedIdentity CreatedByType = original.ManagedIdentity
	User            CreatedByType = original.User
)

type ProvisioningState = original.ProvisioningState

const (
	Accepted     ProvisioningState = original.Accepted
	Canceled     ProvisioningState = original.Canceled
	Deleting     ProvisioningState = original.Deleting
	Failed       ProvisioningState = original.Failed
	Provisioning ProvisioningState = original.Provisioning
	Succeeded    ProvisioningState = original.Succeeded
)

type Account = original.Account
type AccountList = original.AccountList
type AccountListIterator = original.AccountListIterator
type AccountListPage = original.AccountListPage
type AccountPatch = original.AccountPatch
type AccountProperties = original.AccountProperties
type AccountsClient = original.AccountsClient
type AccountsCreateOrUpdateFuture = original.AccountsCreateOrUpdateFuture
type AccountsDeleteFuture = original.AccountsDeleteFuture
type AccountsUpdateFuture = original.AccountsUpdateFuture
type AzureEntityResource = original.AzureEntityResource
type BaseClient = original.BaseClient
type DataPool = original.DataPool
type DataPoolBaseProperties = original.DataPoolBaseProperties
type DataPoolList = original.DataPoolList
type DataPoolListIterator = original.DataPoolListIterator
type DataPoolListPage = original.DataPoolListPage
type DataPoolLocation = original.DataPoolLocation
type DataPoolPatch = original.DataPoolPatch
type DataPoolProperties = original.DataPoolProperties
type DataPoolsClient = original.DataPoolsClient
type DataPoolsCreateOrUpdateFuture = original.DataPoolsCreateOrUpdateFuture
type DataPoolsDeleteFuture = original.DataPoolsDeleteFuture
type DataPoolsUpdateFuture = original.DataPoolsUpdateFuture
type ErrorDefinition = original.ErrorDefinition
type ErrorResponse = original.ErrorResponse
type Operation = original.Operation
type OperationDisplay = original.OperationDisplay
type OperationListResult = original.OperationListResult
type OperationListResultIterator = original.OperationListResultIterator
type OperationListResultPage = original.OperationListResultPage
type OperationLogSpecification = original.OperationLogSpecification
type OperationMetricAvailability = original.OperationMetricAvailability
type OperationMetricSpecification = original.OperationMetricSpecification
type OperationProperties = original.OperationProperties
type OperationServiceSpecification = original.OperationServiceSpecification
type OperationsClient = original.OperationsClient
type ProxyResource = original.ProxyResource
type Resource = original.Resource
type SystemData = original.SystemData
type Tags = original.Tags
type TrackedResource = original.TrackedResource

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewAccountListIterator(page AccountListPage) AccountListIterator {
	return original.NewAccountListIterator(page)
}
func NewAccountListPage(cur AccountList, getNextPage func(context.Context, AccountList) (AccountList, error)) AccountListPage {
	return original.NewAccountListPage(cur, getNextPage)
}
func NewAccountsClient(subscriptionID string) AccountsClient {
	return original.NewAccountsClient(subscriptionID)
}
func NewAccountsClientWithBaseURI(baseURI string, subscriptionID string) AccountsClient {
	return original.NewAccountsClientWithBaseURI(baseURI, subscriptionID)
}
func NewDataPoolListIterator(page DataPoolListPage) DataPoolListIterator {
	return original.NewDataPoolListIterator(page)
}
func NewDataPoolListPage(cur DataPoolList, getNextPage func(context.Context, DataPoolList) (DataPoolList, error)) DataPoolListPage {
	return original.NewDataPoolListPage(cur, getNextPage)
}
func NewDataPoolsClient(subscriptionID string) DataPoolsClient {
	return original.NewDataPoolsClient(subscriptionID)
}
func NewDataPoolsClientWithBaseURI(baseURI string, subscriptionID string) DataPoolsClient {
	return original.NewDataPoolsClientWithBaseURI(baseURI, subscriptionID)
}
func NewOperationListResultIterator(page OperationListResultPage) OperationListResultIterator {
	return original.NewOperationListResultIterator(page)
}
func NewOperationListResultPage(cur OperationListResult, getNextPage func(context.Context, OperationListResult) (OperationListResult, error)) OperationListResultPage {
	return original.NewOperationListResultPage(cur, getNextPage)
}
func NewOperationsClient(subscriptionID string) OperationsClient {
	return original.NewOperationsClient(subscriptionID)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func PossibleCreatedByTypeValues() []CreatedByType {
	return original.PossibleCreatedByTypeValues()
}
func PossibleProvisioningStateValues() []ProvisioningState {
	return original.PossibleProvisioningStateValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}