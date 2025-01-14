//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcosmos_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos"
)

// x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2021-10-15/examples/CosmosDBCollectionGetMetrics.json
func ExampleCollectionClient_ListMetrics() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armcosmos.NewCollectionClient("<subscription-id>", cred, nil)
	res, err := client.ListMetrics(ctx,
		"<resource-group-name>",
		"<account-name>",
		"<database-rid>",
		"<collection-rid>",
		"<filter>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.CollectionClientListMetricsResult)
}

// x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2021-10-15/examples/CosmosDBCollectionGetUsages.json
func ExampleCollectionClient_ListUsages() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armcosmos.NewCollectionClient("<subscription-id>", cred, nil)
	res, err := client.ListUsages(ctx,
		"<resource-group-name>",
		"<account-name>",
		"<database-rid>",
		"<collection-rid>",
		&armcosmos.CollectionClientListUsagesOptions{Filter: to.StringPtr("<filter>")})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.CollectionClientListUsagesResult)
}

// x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2021-10-15/examples/CosmosDBCollectionGetMetricDefinitions.json
func ExampleCollectionClient_ListMetricDefinitions() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armcosmos.NewCollectionClient("<subscription-id>", cred, nil)
	res, err := client.ListMetricDefinitions(ctx,
		"<resource-group-name>",
		"<account-name>",
		"<database-rid>",
		"<collection-rid>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.CollectionClientListMetricDefinitionsResult)
}
