/*
IdentityNow Beta API

Testing RoleInsightsAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package api_beta

import (
	"context"
	"testing"

	openapiclient "github.com/davidsonjon/golang-sdk/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_api_beta_RoleInsightsAPIService(t *testing.T) {

	configuration := openapiclient.NewDefaultConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test RoleInsightsAPIService CreateRoleInsightRequests", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.API_BETA.RoleInsightsAPI.CreateRoleInsightRequests(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RoleInsightsAPIService DownloadRoleInsightsEntitlementsChanges", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var insightId string

		resp, httpRes, err := apiClient.API_BETA.RoleInsightsAPI.DownloadRoleInsightsEntitlementsChanges(context.Background(), insightId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RoleInsightsAPIService GetEntitlementChangesIdentities", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var insightId string
		var entitlementId string

		resp, httpRes, err := apiClient.API_BETA.RoleInsightsAPI.GetEntitlementChangesIdentities(context.Background(), insightId, entitlementId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RoleInsightsAPIService GetRoleInsight", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var insightId string

		resp, httpRes, err := apiClient.API_BETA.RoleInsightsAPI.GetRoleInsight(context.Background(), insightId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RoleInsightsAPIService GetRoleInsights", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.API_BETA.RoleInsightsAPI.GetRoleInsights(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RoleInsightsAPIService GetRoleInsightsCurrentEntitlements", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var insightId string

		resp, httpRes, err := apiClient.API_BETA.RoleInsightsAPI.GetRoleInsightsCurrentEntitlements(context.Background(), insightId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RoleInsightsAPIService GetRoleInsightsEntitlementsChanges", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var insightId string

		resp, httpRes, err := apiClient.API_BETA.RoleInsightsAPI.GetRoleInsightsEntitlementsChanges(context.Background(), insightId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RoleInsightsAPIService GetRoleInsightsRequests", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.API_BETA.RoleInsightsAPI.GetRoleInsightsRequests(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RoleInsightsAPIService GetRoleInsightsSummary", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.API_BETA.RoleInsightsAPI.GetRoleInsightsSummary(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
