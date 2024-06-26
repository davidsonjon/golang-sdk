/*
Identity Security Cloud V3 API

Testing SourcesAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package api_v3

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/sailpoint-oss/golang-sdk/v2"
)

func Test_api_v3_SourcesAPIService(t *testing.T) {

	configuration := openapiclient.NewDefaultConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test SourcesAPIService CreateProvisioningPolicy", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var sourceId string

		resp, httpRes, err := apiClient.API_V3.SourcesAPI.CreateProvisioningPolicy(context.Background(), sourceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SourcesAPIService CreateSource", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.API_V3.SourcesAPI.CreateSource(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SourcesAPIService CreateSourceSchema", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var sourceId string

		resp, httpRes, err := apiClient.API_V3.SourcesAPI.CreateSourceSchema(context.Background(), sourceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SourcesAPIService DeleteProvisioningPolicy", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var sourceId string
		var usageType UsageType

		httpRes, err := apiClient.API_V3.SourcesAPI.DeleteProvisioningPolicy(context.Background(), sourceId, usageType).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SourcesAPIService DeleteSource", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.API_V3.SourcesAPI.DeleteSource(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SourcesAPIService DeleteSourceSchema", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var sourceId string
		var schemaId string

		httpRes, err := apiClient.API_V3.SourcesAPI.DeleteSourceSchema(context.Background(), sourceId, schemaId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SourcesAPIService GetAccountsSchema", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		httpRes, err := apiClient.API_V3.SourcesAPI.GetAccountsSchema(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SourcesAPIService GetEntitlementsSchema", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		httpRes, err := apiClient.API_V3.SourcesAPI.GetEntitlementsSchema(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SourcesAPIService GetProvisioningPolicy", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var sourceId string
		var usageType UsageType

		resp, httpRes, err := apiClient.API_V3.SourcesAPI.GetProvisioningPolicy(context.Background(), sourceId, usageType).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SourcesAPIService GetSource", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.API_V3.SourcesAPI.GetSource(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SourcesAPIService GetSourceHealth", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var sourceId string

		resp, httpRes, err := apiClient.API_V3.SourcesAPI.GetSourceHealth(context.Background(), sourceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SourcesAPIService GetSourceSchema", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var sourceId string
		var schemaId string

		resp, httpRes, err := apiClient.API_V3.SourcesAPI.GetSourceSchema(context.Background(), sourceId, schemaId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SourcesAPIService GetSourceSchemas", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var sourceId string

		resp, httpRes, err := apiClient.API_V3.SourcesAPI.GetSourceSchemas(context.Background(), sourceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SourcesAPIService ImportAccountsSchema", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.API_V3.SourcesAPI.ImportAccountsSchema(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SourcesAPIService ImportConnectorFile", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var sourceId string

		resp, httpRes, err := apiClient.API_V3.SourcesAPI.ImportConnectorFile(context.Background(), sourceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SourcesAPIService ImportEntitlementsSchema", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.API_V3.SourcesAPI.ImportEntitlementsSchema(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SourcesAPIService ListProvisioningPolicies", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var sourceId string

		resp, httpRes, err := apiClient.API_V3.SourcesAPI.ListProvisioningPolicies(context.Background(), sourceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SourcesAPIService ListSources", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.API_V3.SourcesAPI.ListSources(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SourcesAPIService PutProvisioningPolicy", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var sourceId string
		var usageType UsageType

		resp, httpRes, err := apiClient.API_V3.SourcesAPI.PutProvisioningPolicy(context.Background(), sourceId, usageType).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SourcesAPIService PutSource", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.API_V3.SourcesAPI.PutSource(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SourcesAPIService PutSourceSchema", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var sourceId string
		var schemaId string

		resp, httpRes, err := apiClient.API_V3.SourcesAPI.PutSourceSchema(context.Background(), sourceId, schemaId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SourcesAPIService UpdateProvisioningPoliciesInBulk", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var sourceId string

		resp, httpRes, err := apiClient.API_V3.SourcesAPI.UpdateProvisioningPoliciesInBulk(context.Background(), sourceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SourcesAPIService UpdateProvisioningPolicy", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var sourceId string
		var usageType UsageType

		resp, httpRes, err := apiClient.API_V3.SourcesAPI.UpdateProvisioningPolicy(context.Background(), sourceId, usageType).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SourcesAPIService UpdateSource", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.API_V3.SourcesAPI.UpdateSource(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SourcesAPIService UpdateSourceSchema", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var sourceId string
		var schemaId string

		resp, httpRes, err := apiClient.API_V3.SourcesAPI.UpdateSourceSchema(context.Background(), sourceId, schemaId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
