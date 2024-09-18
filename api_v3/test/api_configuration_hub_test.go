/*
Identity Security Cloud V3 API

Testing ConfigurationHubAPIService

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

func Test_api_v3_ConfigurationHubAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ConfigurationHubAPIService CreateObjectMapping", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var sourceOrg string

		resp, httpRes, err := apiClient.ConfigurationHubAPI.CreateObjectMapping(context.Background(), sourceOrg).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ConfigurationHubAPIService CreateObjectMappings", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var sourceOrg string

		resp, httpRes, err := apiClient.ConfigurationHubAPI.CreateObjectMappings(context.Background(), sourceOrg).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ConfigurationHubAPIService CreateUploadedConfiguration", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ConfigurationHubAPI.CreateUploadedConfiguration(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ConfigurationHubAPIService DeleteObjectMapping", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var sourceOrg string
		var objectMappingId string

		httpRes, err := apiClient.ConfigurationHubAPI.DeleteObjectMapping(context.Background(), sourceOrg, objectMappingId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ConfigurationHubAPIService DeleteUploadedConfiguration", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		httpRes, err := apiClient.ConfigurationHubAPI.DeleteUploadedConfiguration(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ConfigurationHubAPIService GetObjectMappings", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var sourceOrg string

		resp, httpRes, err := apiClient.ConfigurationHubAPI.GetObjectMappings(context.Background(), sourceOrg).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ConfigurationHubAPIService GetUploadedConfiguration", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.ConfigurationHubAPI.GetUploadedConfiguration(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ConfigurationHubAPIService ListUploadedConfigurations", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ConfigurationHubAPI.ListUploadedConfigurations(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ConfigurationHubAPIService UpdateObjectMappings", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var sourceOrg string

		resp, httpRes, err := apiClient.ConfigurationHubAPI.UpdateObjectMappings(context.Background(), sourceOrg).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
