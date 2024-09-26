/*
Identity Security Cloud Beta API

Testing MultiHostIntegrationAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package api_beta

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/sailpoint-oss/golang-sdk/v2"
)

func Test_api_beta_MultiHostIntegrationAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test MultiHostIntegrationAPIService CreateMultiHostIntegration", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.MultiHostIntegrationAPI.CreateMultiHostIntegration(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MultiHostIntegrationAPIService CreateSourcesWithinMultiHost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		httpRes, err := apiClient.MultiHostIntegrationAPI.CreateSourcesWithinMultiHost(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MultiHostIntegrationAPIService DeleteMultiHost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		httpRes, err := apiClient.MultiHostIntegrationAPI.DeleteMultiHost(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MultiHostIntegrationAPIService GetAcctAggregationGroups", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var multiHostId string

		resp, httpRes, err := apiClient.MultiHostIntegrationAPI.GetAcctAggregationGroups(context.Background(), multiHostId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MultiHostIntegrationAPIService GetEntitlementAggregationGroups", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var multiHostId string

		resp, httpRes, err := apiClient.MultiHostIntegrationAPI.GetEntitlementAggregationGroups(context.Background(), multiHostId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MultiHostIntegrationAPIService GetMultiHostIntegrations", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.MultiHostIntegrationAPI.GetMultiHostIntegrations(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MultiHostIntegrationAPIService GetMultiHostIntegrationsList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.MultiHostIntegrationAPI.GetMultiHostIntegrationsList(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MultiHostIntegrationAPIService GetMultiHostSourceCreationErrors", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var multiHostId string

		resp, httpRes, err := apiClient.MultiHostIntegrationAPI.GetMultiHostSourceCreationErrors(context.Background(), multiHostId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MultiHostIntegrationAPIService GetMultihostIntegrationTypes", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.MultiHostIntegrationAPI.GetMultihostIntegrationTypes(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MultiHostIntegrationAPIService GetSourcesWithinMultiHost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.MultiHostIntegrationAPI.GetSourcesWithinMultiHost(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MultiHostIntegrationAPIService TestConnectionMultiHostSources", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var multihostId string

		httpRes, err := apiClient.MultiHostIntegrationAPI.TestConnectionMultiHostSources(context.Background(), multihostId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MultiHostIntegrationAPIService TestSourceConnectionMultihost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var multihostId string
		var sourceId string

		resp, httpRes, err := apiClient.MultiHostIntegrationAPI.TestSourceConnectionMultihost(context.Background(), multihostId, sourceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MultiHostIntegrationAPIService UpdateMultiHostSources", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var multihostId string

		httpRes, err := apiClient.MultiHostIntegrationAPI.UpdateMultiHostSources(context.Background(), multihostId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
