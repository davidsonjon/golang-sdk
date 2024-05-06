/*
Identity Security Cloud V3 API

Testing ServiceDeskIntegrationAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package api_v3

import (
	"context"
	"testing"

	openapiclient "github.com/sailpoint-oss/davidsonjon/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_api_v3_ServiceDeskIntegrationAPIService(t *testing.T) {

	configuration := openapiclient.NewDefaultConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ServiceDeskIntegrationAPIService CreateServiceDeskIntegration", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.API_V3.ServiceDeskIntegrationAPI.CreateServiceDeskIntegration(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ServiceDeskIntegrationAPIService DeleteServiceDeskIntegration", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		httpRes, err := apiClient.API_V3.ServiceDeskIntegrationAPI.DeleteServiceDeskIntegration(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ServiceDeskIntegrationAPIService GetServiceDeskIntegration", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.API_V3.ServiceDeskIntegrationAPI.GetServiceDeskIntegration(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ServiceDeskIntegrationAPIService GetServiceDeskIntegrationTemplate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var scriptName string

		resp, httpRes, err := apiClient.API_V3.ServiceDeskIntegrationAPI.GetServiceDeskIntegrationTemplate(context.Background(), scriptName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ServiceDeskIntegrationAPIService GetServiceDeskIntegrationTypes", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.API_V3.ServiceDeskIntegrationAPI.GetServiceDeskIntegrationTypes(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ServiceDeskIntegrationAPIService GetServiceDeskIntegrations", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.API_V3.ServiceDeskIntegrationAPI.GetServiceDeskIntegrations(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ServiceDeskIntegrationAPIService GetStatusCheckDetails", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.API_V3.ServiceDeskIntegrationAPI.GetStatusCheckDetails(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ServiceDeskIntegrationAPIService PatchServiceDeskIntegration", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.API_V3.ServiceDeskIntegrationAPI.PatchServiceDeskIntegration(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ServiceDeskIntegrationAPIService PutServiceDeskIntegration", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.API_V3.ServiceDeskIntegrationAPI.PutServiceDeskIntegration(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ServiceDeskIntegrationAPIService UpdateStatusCheckDetails", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.API_V3.ServiceDeskIntegrationAPI.UpdateStatusCheckDetails(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
