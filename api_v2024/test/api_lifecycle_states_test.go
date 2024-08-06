/*
Identity Security Cloud V2024 API

Testing LifecycleStatesAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package api_v2024

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/sailpoint-oss/golang-sdk/v2"
)

func Test_api_v2024_LifecycleStatesAPIService(t *testing.T) {

	configuration := openapiclient.NewDefaultConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test LifecycleStatesAPIService CreateLifecycleState", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var identityProfileId string

		resp, httpRes, err := apiClient.API_V2024.LifecycleStatesAPI.CreateLifecycleState(context.Background(), identityProfileId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LifecycleStatesAPIService DeleteLifecycleState", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var identityProfileId string
		var lifecycleStateId string

		resp, httpRes, err := apiClient.API_V2024.LifecycleStatesAPI.DeleteLifecycleState(context.Background(), identityProfileId, lifecycleStateId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LifecycleStatesAPIService GetLifecycleState", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var identityProfileId string
		var lifecycleStateId string

		resp, httpRes, err := apiClient.API_V2024.LifecycleStatesAPI.GetLifecycleState(context.Background(), identityProfileId, lifecycleStateId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LifecycleStatesAPIService GetLifecycleStates", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var identityProfileId string

		resp, httpRes, err := apiClient.API_V2024.LifecycleStatesAPI.GetLifecycleStates(context.Background(), identityProfileId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LifecycleStatesAPIService SetLifecycleState", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var identityId string

		resp, httpRes, err := apiClient.API_V2024.LifecycleStatesAPI.SetLifecycleState(context.Background(), identityId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LifecycleStatesAPIService UpdateLifecycleStates", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var identityProfileId string
		var lifecycleStateId string

		resp, httpRes, err := apiClient.API_V2024.LifecycleStatesAPI.UpdateLifecycleStates(context.Background(), identityProfileId, lifecycleStateId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
