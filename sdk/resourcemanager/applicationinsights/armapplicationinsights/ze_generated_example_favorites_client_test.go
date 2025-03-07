//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armapplicationinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/applicationinsights/armapplicationinsights/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2015-05-01/examples/FavoritesList.json
func ExampleFavoritesClient_List() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armapplicationinsights.NewFavoritesClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.List(ctx,
		"my-resource-group",
		"my-ai-component",
		&armapplicationinsights.FavoritesClientListOptions{FavoriteType: nil,
			SourceType:      nil,
			CanFetchContent: nil,
			Tags:            []string{},
		})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2015-05-01/examples/FavoriteGet.json
func ExampleFavoritesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armapplicationinsights.NewFavoritesClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Get(ctx,
		"my-resource-group",
		"my-ai-component",
		"deadb33f-5e0d-4064-8ebb-1a4ed0313eb2",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2015-05-01/examples/FavoriteAdd.json
func ExampleFavoritesClient_Add() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armapplicationinsights.NewFavoritesClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Add(ctx,
		"my-resource-group",
		"my-ai-component",
		"deadb33f-8bee-4d3b-a059-9be8dac93960",
		armapplicationinsights.ComponentFavorite{
			Config:                  to.Ptr("{\"MEDataModelRawJSON\":\"{\\n  \\\"version\\\": \\\"1.4.1\\\",\\n  \\\"isCustomDataModel\\\": true,\\n  \\\"items\\\": [\\n    {\\n      \\\"id\\\": \\\"90a7134d-9a38-4c25-88d3-a495209873eb\\\",\\n      \\\"chartType\\\": \\\"Area\\\",\\n      \\\"chartHeight\\\": 4,\\n      \\\"metrics\\\": [\\n        {\\n          \\\"id\\\": \\\"preview/requests/count\\\",\\n          \\\"metricAggregation\\\": \\\"Sum\\\",\\n          \\\"color\\\": \\\"msportalfx-bgcolor-d0\\\"\\n        }\\n      ],\\n      \\\"priorPeriod\\\": false,\\n      \\\"clickAction\\\": {\\n        \\\"defaultBlade\\\": \\\"SearchBlade\\\"\\n      },\\n      \\\"horizontalBars\\\": true,\\n      \\\"showOther\\\": true,\\n      \\\"aggregation\\\": \\\"Sum\\\",\\n      \\\"percentage\\\": false,\\n      \\\"palette\\\": \\\"fail\\\",\\n      \\\"yAxisOption\\\": 0,\\n      \\\"title\\\": \\\"\\\"\\n    },\\n    {\\n      \\\"id\\\": \\\"0c289098-88e8-4010-b212-546815cddf70\\\",\\n      \\\"chartType\\\": \\\"Area\\\",\\n      \\\"chartHeight\\\": 2,\\n      \\\"metrics\\\": [\\n        {\\n          \\\"id\\\": \\\"preview/requests/duration\\\",\\n          \\\"metricAggregation\\\": \\\"Avg\\\",\\n          \\\"color\\\": \\\"msportalfx-bgcolor-j1\\\"\\n        }\\n      ],\\n      \\\"priorPeriod\\\": false,\\n      \\\"clickAction\\\": {\\n        \\\"defaultBlade\\\": \\\"SearchBlade\\\"\\n      },\\n      \\\"horizontalBars\\\": true,\\n      \\\"showOther\\\": true,\\n      \\\"aggregation\\\": \\\"Avg\\\",\\n      \\\"percentage\\\": false,\\n      \\\"palette\\\": \\\"greenHues\\\",\\n      \\\"yAxisOption\\\": 0,\\n      \\\"title\\\": \\\"\\\"\\n    },\\n    {\\n      \\\"id\\\": \\\"cbdaab6f-a808-4f71-aca5-b3976cbb7345\\\",\\n      \\\"chartType\\\": \\\"Bar\\\",\\n      \\\"chartHeight\\\": 4,\\n      \\\"metrics\\\": [\\n        {\\n          \\\"id\\\": \\\"preview/requests/duration\\\",\\n          \\\"metricAggregation\\\": \\\"Avg\\\",\\n          \\\"color\\\": \\\"msportalfx-bgcolor-d0\\\"\\n        }\\n      ],\\n      \\\"priorPeriod\\\": false,\\n      \\\"clickAction\\\": {\\n        \\\"defaultBlade\\\": \\\"SearchBlade\\\"\\n      },\\n      \\\"horizontalBars\\\": true,\\n      \\\"showOther\\\": true,\\n      \\\"aggregation\\\": \\\"Avg\\\",\\n      \\\"percentage\\\": false,\\n      \\\"palette\\\": \\\"magentaHues\\\",\\n      \\\"yAxisOption\\\": 0,\\n      \\\"title\\\": \\\"\\\"\\n    },\\n    {\\n      \\\"id\\\": \\\"1d5a6a3a-9fa1-4099-9cf9-05eff72d1b02\\\",\\n      \\\"grouping\\\": {\\n        \\\"kind\\\": \\\"ByDimension\\\",\\n        \\\"dimension\\\": \\\"context.application.version\\\"\\n      },\\n      \\\"chartType\\\": \\\"Grid\\\",\\n      \\\"chartHeight\\\": 1,\\n      \\\"metrics\\\": [\\n        {\\n          \\\"id\\\": \\\"basicException.count\\\",\\n          \\\"metricAggregation\\\": \\\"Sum\\\",\\n          \\\"color\\\": \\\"msportalfx-bgcolor-g0\\\"\\n        },\\n        {\\n          \\\"id\\\": \\\"requestFailed.count\\\",\\n          \\\"metricAggregation\\\": \\\"Sum\\\",\\n          \\\"color\\\": \\\"msportalfx-bgcolor-f0s2\\\"\\n        }\\n      ],\\n      \\\"priorPeriod\\\": true,\\n      \\\"clickAction\\\": {\\n        \\\"defaultBlade\\\": \\\"SearchBlade\\\"\\n      },\\n      \\\"horizontalBars\\\": true,\\n      \\\"showOther\\\": true,\\n      \\\"percentage\\\": false,\\n      \\\"palette\\\": \\\"blueHues\\\",\\n      \\\"yAxisOption\\\": 0,\\n      \\\"title\\\": \\\"\\\"\\n    }\\n  ],\\n  \\\"currentFilter\\\": {\\n    \\\"eventTypes\\\": [\\n      1,\\n      2\\n    ],\\n    \\\"typeFacets\\\": {},\\n    \\\"isPermissive\\\": false\\n  },\\n  \\\"timeContext\\\": {\\n    \\\"durationMs\\\": 75600000,\\n    \\\"endTime\\\": \\\"2018-01-31T20:30:00.000Z\\\",\\n    \\\"createdTime\\\": \\\"2018-01-31T23:54:26.280Z\\\",\\n    \\\"isInitialTime\\\": false,\\n    \\\"grain\\\": 1,\\n    \\\"useDashboardTimeRange\\\": false\\n  },\\n  \\\"jsonUri\\\": \\\"Favorite_BlankChart\\\",\\n  \\\"timeSource\\\": 0\\n}\"}"),
			FavoriteID:              to.Ptr("deadb33f-8bee-4d3b-a059-9be8dac93960"),
			FavoriteType:            to.Ptr(armapplicationinsights.FavoriteTypeShared),
			IsGeneratedFromTemplate: to.Ptr(false),
			Name:                    to.Ptr("Blah Blah Blah"),
			Tags: []*string{
				to.Ptr("TagSample01"),
				to.Ptr("TagSample02")},
			Version: to.Ptr("ME"),
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2015-05-01/examples/FavoriteUpdate.json
func ExampleFavoritesClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armapplicationinsights.NewFavoritesClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Update(ctx,
		"my-resource-group",
		"my-ai-component",
		"deadb33f-5e0d-4064-8ebb-1a4ed0313eb2",
		armapplicationinsights.ComponentFavorite{
			Config:                  to.Ptr("{\"MEDataModelRawJSON\":\"{\\\"version\\\": \\\"1.4.1\\\",\\\"isCustomDataModel\\\": true,\\\"items\\\": [{\\\"id\\\": \\\"90a7134d-9a38-4c25-88d3-a495209873eb\\\",\\\"chartType\\\": \\\"Area\\\",\\\"chartHeight\\\": 4,\\\"metrics\\\": [{\\\"id\\\": \\\"preview/requests/count\\\",\\\"metricAggregation\\\": \\\"Sum\\\",\\\"color\\\": \\\"msportalfx-bgcolor-d0\\\"}],\\\"priorPeriod\\\": false,\\\"clickAction\\\": {\\\"defaultBlade\\\": \\\"SearchBlade\\\"},\\\"horizontalBars\\\": true,\\\"showOther\\\": true,\\\"aggregation\\\": \\\"Sum\\\",\\\"percentage\\\": false,\\\"palette\\\": \\\"fail\\\",\\\"yAxisOption\\\": 0,\\\"title\\\": \\\"\\\"},{\\\"id\\\": \\\"0c289098-88e8-4010-b212-546815cddf70\\\",\\\"chartType\\\": \\\"Area\\\",\\\"chartHeight\\\": 2,\\\"metrics\\\": [{\\\"id\\\": \\\"preview/requests/duration\\\",\\\"metricAggregation\\\": \\\"Avg\\\",\\\"color\\\": \\\"msportalfx-bgcolor-j1\\\"}],\\\"priorPeriod\\\": false,\\\"clickAction\\\": {\\\"defaultBlade\\\": \\\"SearchBlade\\\"},\\\"horizontalBars\\\": true,\\\"showOther\\\": true,\\\"aggregation\\\": \\\"Avg\\\",\\\"percentage\\\": false,\\\"palette\\\": \\\"greenHues\\\",\\\"yAxisOption\\\": 0,\\\"title\\\": \\\"\\\"},{\\\"id\\\": \\\"cbdaab6f-a808-4f71-aca5-b3976cbb7345\\\",\\\"chartType\\\": \\\"Bar\\\",\\\"chartHeight\\\": 4,\\\"metrics\\\": [{\\\"id\\\": \\\"preview/requests/duration\\\",\\\"metricAggregation\\\": \\\"Avg\\\",\\\"color\\\": \\\"msportalfx-bgcolor-d0\\\"}],\\\"priorPeriod\\\": false,\\\"clickAction\\\": {\\\"defaultBlade\\\": \\\"SearchBlade\\\"},\\\"horizontalBars\\\": true,\\\"showOther\\\": true,\\\"aggregation\\\": \\\"Avg\\\",\\\"percentage\\\": false,\\\"palette\\\": \\\"magentaHues\\\",\\\"yAxisOption\\\": 0,\\\"title\\\": \\\"\\\"},{\\\"id\\\": \\\"1d5a6a3a-9fa1-4099-9cf9-05eff72d1b02\\\",\\\"grouping\\\": {\\\"kind\\\": \\\"ByDimension\\\",\\\"dimension\\\": \\\"context.application.version\\\"},\\\"chartType\\\": \\\"Grid\\\",\\\"chartHeight\\\": 1,\\\"metrics\\\": [{\\\"id\\\": \\\"basicException.count\\\",\\\"metricAggregation\\\": \\\"Sum\\\",\\\"color\\\": \\\"msportalfx-bgcolor-g0\\\"},{\\\"id\\\": \\\"requestFailed.count\\\",\\\"metricAggregation\\\": \\\"Sum\\\",\\\"color\\\": \\\"msportalfx-bgcolor-f0s2\\\"}],\\\"priorPeriod\\\": true,\\\"clickAction\\\": {\\\"defaultBlade\\\": \\\"SearchBlade\\\"},\\\"horizontalBars\\\": true,\\\"showOther\\\": true,\\\"percentage\\\": false,\\\"palette\\\": \\\"blueHues\\\",\\\"yAxisOption\\\": 0,\\\"title\\\": \\\"\\\"}],\\\"currentFilter\\\": {\\\"eventTypes\\\": [1,2],\\\"typeFacets\\\": {},\\\"isPermissive\\\": false},\\\"timeContext\\\": {\\\"durationMs\\\": 75600000,\\\"endTime\\\": \\\"2018-01-31T20:30:00.000Z\\\",\\\"createdTime\\\": \\\"2018-01-31T23:54:26.280Z\\\",\\\"isInitialTime\\\": false,\\\"grain\\\": 1,\\\"useDashboardTimeRange\\\": false},\\\"jsonUri\\\": \\\"Favorite_BlankChart\\\",\\\"timeSource\\\": 0}\"}"),
			FavoriteID:              to.Ptr("deadb33f-5e0d-4064-8ebb-1a4ed0313eb2"),
			FavoriteType:            to.Ptr(armapplicationinsights.FavoriteTypeShared),
			IsGeneratedFromTemplate: to.Ptr(false),
			Name:                    to.Ptr("Derek Changed This"),
			Tags: []*string{
				to.Ptr("TagSample01"),
				to.Ptr("TagSample02"),
				to.Ptr("TagSample03")},
			TimeModified: to.Ptr("2018-02-02T18:39:11.6569686Z"),
			Version:      to.Ptr("ME"),
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2015-05-01/examples/FavoriteDelete.json
func ExampleFavoritesClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armapplicationinsights.NewFavoritesClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = client.Delete(ctx,
		"my-resource-group",
		"my-ai-component",
		"deadb33f-5e0d-4064-8ebb-1a4ed0313eb2",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
