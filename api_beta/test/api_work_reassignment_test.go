/*
Identity Security Cloud Beta API

Testing WorkReassignmentAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package api_beta

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/davidsonjon/golang-sdk/v2"
)

func Test_api_beta_WorkReassignmentAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test WorkReassignmentAPIService CreateReassignmentConfiguration", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.WorkReassignmentAPI.CreateReassignmentConfiguration(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WorkReassignmentAPIService DeleteReassignmentConfiguration", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var identityId string

		httpRes, err := apiClient.WorkReassignmentAPI.DeleteReassignmentConfiguration(context.Background(), identityId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WorkReassignmentAPIService GetEvaluateReassignmentConfiguration", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var identityId string
		var configType ConfigTypeEnum

		resp, httpRes, err := apiClient.WorkReassignmentAPI.GetEvaluateReassignmentConfiguration(context.Background(), identityId, configType).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WorkReassignmentAPIService GetReassignmentConfigTypes", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.WorkReassignmentAPI.GetReassignmentConfigTypes(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WorkReassignmentAPIService GetReassignmentConfiguration", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var identityId string

		resp, httpRes, err := apiClient.WorkReassignmentAPI.GetReassignmentConfiguration(context.Background(), identityId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WorkReassignmentAPIService GetTenantConfigConfiguration", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.WorkReassignmentAPI.GetTenantConfigConfiguration(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WorkReassignmentAPIService ListReassignmentConfigurations", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.WorkReassignmentAPI.ListReassignmentConfigurations(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WorkReassignmentAPIService PutReassignmentConfig", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var identityId string

		resp, httpRes, err := apiClient.WorkReassignmentAPI.PutReassignmentConfig(context.Background(), identityId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WorkReassignmentAPIService PutTenantConfiguration", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.WorkReassignmentAPI.PutTenantConfiguration(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
