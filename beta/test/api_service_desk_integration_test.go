/*
IdentityNow Beta API

Testing ServiceDeskIntegrationApiService

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

func Test_beta_ServiceDeskIntegrationApiService(t *testing.T) {

	configuration := openapiclient.NewDefaultConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ServiceDeskIntegrationApiService CreateServiceDeskIntegration", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.BETA.ServiceDeskIntegrationApi.CreateServiceDeskIntegration(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ServiceDeskIntegrationApiService DeleteServiceDeskIntegration", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		httpRes, err := apiClient.BETA.ServiceDeskIntegrationApi.DeleteServiceDeskIntegration(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ServiceDeskIntegrationApiService GetServiceDeskIntegration", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.BETA.ServiceDeskIntegrationApi.GetServiceDeskIntegration(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ServiceDeskIntegrationApiService GetServiceDeskIntegrationList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.BETA.ServiceDeskIntegrationApi.GetServiceDeskIntegrationList(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ServiceDeskIntegrationApiService GetServiceDeskIntegrationTemplate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var scriptName string

		resp, httpRes, err := apiClient.BETA.ServiceDeskIntegrationApi.GetServiceDeskIntegrationTemplate(context.Background(), scriptName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ServiceDeskIntegrationApiService GetServiceDeskIntegrationTypes", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.BETA.ServiceDeskIntegrationApi.GetServiceDeskIntegrationTypes(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ServiceDeskIntegrationApiService GetStatusCheckDetails", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.BETA.ServiceDeskIntegrationApi.GetStatusCheckDetails(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ServiceDeskIntegrationApiService PatchServiceDeskIntegration", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.BETA.ServiceDeskIntegrationApi.PatchServiceDeskIntegration(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ServiceDeskIntegrationApiService UpdateManagedClientStatusCheckDetails", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.BETA.ServiceDeskIntegrationApi.UpdateManagedClientStatusCheckDetails(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ServiceDeskIntegrationApiService UpdateServiceDeskIntegration", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.BETA.ServiceDeskIntegrationApi.UpdateServiceDeskIntegration(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
