/*
IdentityNow Beta API

Testing SourcesApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package beta

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/sailpoint-oss/golang-sdk"
)

func Test_beta_SourcesApiService(t *testing.T) {

	configuration := openapiclient.NewDefaultConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test SourcesApiService CreateProvisioningPolicy", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var sourceId string

		resp, httpRes, err := apiClient.BETA.SourcesApi.CreateProvisioningPolicy(context.Background(), sourceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SourcesApiService CreateSource", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.BETA.SourcesApi.CreateSource(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SourcesApiService CreateSourceSchema", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var sourceId string

		resp, httpRes, err := apiClient.BETA.SourcesApi.CreateSourceSchema(context.Background(), sourceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SourcesApiService DeleteProvisioningPolicy", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var sourceId string
		var usageType UsageType

		httpRes, err := apiClient.BETA.SourcesApi.DeleteProvisioningPolicy(context.Background(), sourceId, usageType).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SourcesApiService DeleteSource", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.BETA.SourcesApi.DeleteSource(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SourcesApiService DeleteSourceSchema", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var sourceId string
		var schemaId string

		httpRes, err := apiClient.BETA.SourcesApi.DeleteSourceSchema(context.Background(), sourceId, schemaId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SourcesApiService DownloadSourceAccountsSchema", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		httpRes, err := apiClient.BETA.SourcesApi.DownloadSourceAccountsSchema(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SourcesApiService DownloadSourceEntitlementsSchema", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		httpRes, err := apiClient.BETA.SourcesApi.DownloadSourceEntitlementsSchema(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SourcesApiService GetProvisioningPolicy", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var sourceId string
		var usageType UsageType

		resp, httpRes, err := apiClient.BETA.SourcesApi.GetProvisioningPolicy(context.Background(), sourceId, usageType).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SourcesApiService GetSource", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.BETA.SourcesApi.GetSource(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SourcesApiService GetSourceAttrSyncConfig", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.BETA.SourcesApi.GetSourceAttrSyncConfig(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SourcesApiService GetSourceConfig", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.BETA.SourcesApi.GetSourceConfig(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SourcesApiService GetSourceEntitlementRequestConfig", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.BETA.SourcesApi.GetSourceEntitlementRequestConfig(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SourcesApiService GetSourceSchema", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var sourceId string
		var schemaId string

		resp, httpRes, err := apiClient.BETA.SourcesApi.GetSourceSchema(context.Background(), sourceId, schemaId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SourcesApiService ListProvisioningPolicies", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var sourceId string

		resp, httpRes, err := apiClient.BETA.SourcesApi.ListProvisioningPolicies(context.Background(), sourceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SourcesApiService ListSourceSchemas", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var sourceId string

		resp, httpRes, err := apiClient.BETA.SourcesApi.ListSourceSchemas(context.Background(), sourceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SourcesApiService ListSources", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.BETA.SourcesApi.ListSources(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SourcesApiService PeekResourceObjects", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var sourceId string

		resp, httpRes, err := apiClient.BETA.SourcesApi.PeekResourceObjects(context.Background(), sourceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SourcesApiService PingCluster", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var sourceId string

		resp, httpRes, err := apiClient.BETA.SourcesApi.PingCluster(context.Background(), sourceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SourcesApiService PutProvisioningPolicy", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var sourceId string
		var usageType UsageType

		resp, httpRes, err := apiClient.BETA.SourcesApi.PutProvisioningPolicy(context.Background(), sourceId, usageType).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SourcesApiService PutSource", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.BETA.SourcesApi.PutSource(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SourcesApiService PutSourceAttrSyncConfig", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.BETA.SourcesApi.PutSourceAttrSyncConfig(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SourcesApiService PutSourceSchema", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var sourceId string
		var schemaId string

		resp, httpRes, err := apiClient.BETA.SourcesApi.PutSourceSchema(context.Background(), sourceId, schemaId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SourcesApiService SyncAttributesForSource", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.BETA.SourcesApi.SyncAttributesForSource(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SourcesApiService TestSourceConfiguration", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var sourceId string

		resp, httpRes, err := apiClient.BETA.SourcesApi.TestSourceConfiguration(context.Background(), sourceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SourcesApiService TestSourceConnection", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var sourceId string

		resp, httpRes, err := apiClient.BETA.SourcesApi.TestSourceConnection(context.Background(), sourceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SourcesApiService UpdateProvisioningPoliciesInBulk", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var sourceId string

		resp, httpRes, err := apiClient.BETA.SourcesApi.UpdateProvisioningPoliciesInBulk(context.Background(), sourceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SourcesApiService UpdateProvisioningPolicy", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var sourceId string
		var usageType UsageType

		resp, httpRes, err := apiClient.BETA.SourcesApi.UpdateProvisioningPolicy(context.Background(), sourceId, usageType).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SourcesApiService UpdateSource", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.BETA.SourcesApi.UpdateSource(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SourcesApiService UpdateSourceEntitlementRequestConfig", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.BETA.SourcesApi.UpdateSourceEntitlementRequestConfig(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SourcesApiService UpdateSourceSchema", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var sourceId string
		var schemaId string

		resp, httpRes, err := apiClient.BETA.SourcesApi.UpdateSourceSchema(context.Background(), sourceId, schemaId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SourcesApiService UploadSourceAccountsSchema", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.BETA.SourcesApi.UploadSourceAccountsSchema(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SourcesApiService UploadSourceConnectorFile", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var sourceId string

		resp, httpRes, err := apiClient.BETA.SourcesApi.UploadSourceConnectorFile(context.Background(), sourceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SourcesApiService UploadSourceEntitlementsSchema", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.BETA.SourcesApi.UploadSourceEntitlementsSchema(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
