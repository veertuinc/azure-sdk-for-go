//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/sql/resource-manager/Microsoft.Sql/preview/2021-05-01-preview/examples/LongTermRetentionBackupCopy.json
func ExampleLongTermRetentionBackupsClient_BeginCopy() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsql.NewLongTermRetentionBackupsClient("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCopy(ctx,
		"japaneast",
		"testserver",
		"testDatabase",
		"55555555-6666-7777-8888-999999999999;131637960820000000",
		armsql.CopyLongTermRetentionBackupParameters{
			Properties: &armsql.CopyLongTermRetentionBackupParametersProperties{
				TargetBackupStorageRedundancy: to.Ptr(armsql.BackupStorageRedundancyGeo),
				TargetDatabaseName:            to.Ptr("testDatabase2"),
				TargetServerResourceID:        to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/providers/Microsoft.Sql/resourceGroups/resourceGroup/servers/testserver2"),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/sql/resource-manager/Microsoft.Sql/preview/2021-05-01-preview/examples/LongTermRetentionBackupUpdate.json
func ExampleLongTermRetentionBackupsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsql.NewLongTermRetentionBackupsClient("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginUpdate(ctx,
		"japaneast",
		"testserver",
		"testDatabase",
		"55555555-6666-7777-8888-999999999999;131637960820000000",
		armsql.UpdateLongTermRetentionBackupParameters{
			Properties: &armsql.UpdateLongTermRetentionBackupParametersProperties{
				RequestedBackupStorageRedundancy: to.Ptr(armsql.BackupStorageRedundancyGeo),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/sql/resource-manager/Microsoft.Sql/preview/2021-05-01-preview/examples/LongTermRetentionBackupGet.json
func ExampleLongTermRetentionBackupsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsql.NewLongTermRetentionBackupsClient("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Get(ctx,
		"japaneast",
		"testserver",
		"testDatabase",
		"55555555-6666-7777-8888-999999999999;131637960820000000",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/sql/resource-manager/Microsoft.Sql/preview/2021-05-01-preview/examples/LongTermRetentionBackupDelete.json
func ExampleLongTermRetentionBackupsClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsql.NewLongTermRetentionBackupsClient("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginDelete(ctx,
		"japaneast",
		"testserver",
		"testDatabase",
		"55555555-6666-7777-8888-999999999999;131637960820000000",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/sql/resource-manager/Microsoft.Sql/preview/2021-05-01-preview/examples/LongTermRetentionBackupListByDatabase.json
func ExampleLongTermRetentionBackupsClient_NewListByDatabasePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsql.NewLongTermRetentionBackupsClient("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListByDatabasePager("japaneast",
		"testserver",
		"testDatabase",
		&armsql.LongTermRetentionBackupsClientListByDatabaseOptions{OnlyLatestPerDatabase: nil,
			DatabaseState: nil,
		})
	for pager.More() {
		nextResult, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range nextResult.Value {
			// TODO: use page item
			_ = v
		}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/sql/resource-manager/Microsoft.Sql/preview/2021-05-01-preview/examples/LongTermRetentionBackupListByLocation.json
func ExampleLongTermRetentionBackupsClient_NewListByLocationPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsql.NewLongTermRetentionBackupsClient("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListByLocationPager("japaneast",
		&armsql.LongTermRetentionBackupsClientListByLocationOptions{OnlyLatestPerDatabase: nil,
			DatabaseState: nil,
		})
	for pager.More() {
		nextResult, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range nextResult.Value {
			// TODO: use page item
			_ = v
		}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/sql/resource-manager/Microsoft.Sql/preview/2021-05-01-preview/examples/LongTermRetentionBackupListByServer.json
func ExampleLongTermRetentionBackupsClient_NewListByServerPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsql.NewLongTermRetentionBackupsClient("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListByServerPager("japaneast",
		"testserver",
		&armsql.LongTermRetentionBackupsClientListByServerOptions{OnlyLatestPerDatabase: nil,
			DatabaseState: nil,
		})
	for pager.More() {
		nextResult, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range nextResult.Value {
			// TODO: use page item
			_ = v
		}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/sql/resource-manager/Microsoft.Sql/preview/2021-05-01-preview/examples/ResourceGroupBasedLongTermRetentionBackupCopy.json
func ExampleLongTermRetentionBackupsClient_BeginCopyByResourceGroup() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsql.NewLongTermRetentionBackupsClient("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCopyByResourceGroup(ctx,
		"testResourceGroup",
		"japaneast",
		"testserver",
		"testDatabase",
		"55555555-6666-7777-8888-999999999999;131637960820000000",
		armsql.CopyLongTermRetentionBackupParameters{
			Properties: &armsql.CopyLongTermRetentionBackupParametersProperties{
				TargetBackupStorageRedundancy: to.Ptr(armsql.BackupStorageRedundancyGeo),
				TargetDatabaseName:            to.Ptr("testDatabase2"),
				TargetServerResourceID:        to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/providers/Microsoft.Sql/resourceGroups/resourceGroup/servers/testserver2"),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/sql/resource-manager/Microsoft.Sql/preview/2021-05-01-preview/examples/ResourceGroupBasedLongTermRetentionBackupUpdate.json
func ExampleLongTermRetentionBackupsClient_BeginUpdateByResourceGroup() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsql.NewLongTermRetentionBackupsClient("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginUpdateByResourceGroup(ctx,
		"testResourceGroup",
		"japaneast",
		"testserver",
		"testDatabase",
		"55555555-6666-7777-8888-999999999999;131637960820000000",
		armsql.UpdateLongTermRetentionBackupParameters{
			Properties: &armsql.UpdateLongTermRetentionBackupParametersProperties{
				RequestedBackupStorageRedundancy: to.Ptr(armsql.BackupStorageRedundancyGeo),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/sql/resource-manager/Microsoft.Sql/preview/2021-05-01-preview/examples/ResourceGroupBasedLongTermRetentionBackupGet.json
func ExampleLongTermRetentionBackupsClient_GetByResourceGroup() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsql.NewLongTermRetentionBackupsClient("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.GetByResourceGroup(ctx,
		"testResourceGroup",
		"japaneast",
		"testserver",
		"testDatabase",
		"55555555-6666-7777-8888-999999999999;131637960820000000",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/sql/resource-manager/Microsoft.Sql/preview/2021-05-01-preview/examples/ResourceGroupBasedLongTermRetentionBackupDelete.json
func ExampleLongTermRetentionBackupsClient_BeginDeleteByResourceGroup() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsql.NewLongTermRetentionBackupsClient("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginDeleteByResourceGroup(ctx,
		"testResourceGroup",
		"japaneast",
		"testserver",
		"testDatabase",
		"55555555-6666-7777-8888-999999999999;131637960820000000",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/sql/resource-manager/Microsoft.Sql/preview/2021-05-01-preview/examples/ResourceGroupBasedLongTermRetentionBackupListByDatabase.json
func ExampleLongTermRetentionBackupsClient_NewListByResourceGroupDatabasePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsql.NewLongTermRetentionBackupsClient("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListByResourceGroupDatabasePager("testResourceGroup",
		"japaneast",
		"testserver",
		"testDatabase",
		&armsql.LongTermRetentionBackupsClientListByResourceGroupDatabaseOptions{OnlyLatestPerDatabase: nil,
			DatabaseState: nil,
		})
	for pager.More() {
		nextResult, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range nextResult.Value {
			// TODO: use page item
			_ = v
		}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/sql/resource-manager/Microsoft.Sql/preview/2021-05-01-preview/examples/ResourceGroupBasedLongTermRetentionBackupListByLocation.json
func ExampleLongTermRetentionBackupsClient_NewListByResourceGroupLocationPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsql.NewLongTermRetentionBackupsClient("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListByResourceGroupLocationPager("testResourceGroup",
		"japaneast",
		&armsql.LongTermRetentionBackupsClientListByResourceGroupLocationOptions{OnlyLatestPerDatabase: nil,
			DatabaseState: nil,
		})
	for pager.More() {
		nextResult, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range nextResult.Value {
			// TODO: use page item
			_ = v
		}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/sql/resource-manager/Microsoft.Sql/preview/2021-05-01-preview/examples/ResourceGroupBasedLongTermRetentionBackupListByServer.json
func ExampleLongTermRetentionBackupsClient_NewListByResourceGroupServerPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsql.NewLongTermRetentionBackupsClient("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListByResourceGroupServerPager("testResourceGroup",
		"japaneast",
		"testserver",
		&armsql.LongTermRetentionBackupsClientListByResourceGroupServerOptions{OnlyLatestPerDatabase: nil,
			DatabaseState: nil,
		})
	for pager.More() {
		nextResult, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range nextResult.Value {
			// TODO: use page item
			_ = v
		}
	}
}
