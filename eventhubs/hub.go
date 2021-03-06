// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package eventhubs

import (
	"context"

	"github.com/Azure-Samples/azure-sdk-for-go-samples/helpers"
	"github.com/Azure-Samples/azure-sdk-for-go-samples/iam"
	"github.com/Azure/azure-sdk-for-go/services/eventhub/mgmt/2017-04-01/eventhub"
	"github.com/Azure/go-autorest/autorest/to"
)

func getHubsClient() eventhub.EventHubsClient {
	hubClient := eventhub.NewEventHubsClient(helpers.SubscriptionID())
	auth, _ := iam.GetResourceManagementAuthorizer(iam.AuthGrantType())
	hubClient.Authorizer = auth
	hubClient.AddToUserAgent(helpers.UserAgent())
	return hubClient
}

// CreateHub creates an Event Hubs hub in a namespace
func CreateHub(ctx context.Context, nsName string, hubName string) (eventhub.Model, error) {
	hubClient := getHubsClient()
	return hubClient.CreateOrUpdate(
		ctx,
		helpers.ResourceGroupName(),
		nsName,
		hubName,
		eventhub.Model{
			Properties: &eventhub.Properties{
				PartitionCount: to.Int64Ptr(4),
			},
		},
	)
}
