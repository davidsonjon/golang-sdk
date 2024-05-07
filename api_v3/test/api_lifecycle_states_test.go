/*
Identity Security Cloud V3 API

Testing LifecycleStatesAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package api_v3

import (
	"context"
	"testing"

	openapiclient "github.com/davidsonjon/golang-sdk/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_api_v3_LifecycleStatesAPIService(t *testing.T) {

	configuration := openapiclient.NewDefaultConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test LifecycleStatesAPIService CreateLifecycleState", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var identityProfileId string

		resp, httpRes, err := apiClient.API_V3.LifecycleStatesAPI.CreateLifecycleState(context.Background(), identityProfileId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LifecycleStatesAPIService DeleteLifecycleState", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var identityProfileId string
		var lifecycleStateId string

		resp, httpRes, err := apiClient.API_V3.LifecycleStatesAPI.DeleteLifecycleState(context.Background(), identityProfileId, lifecycleStateId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LifecycleStatesAPIService GetLifecycleState", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var identityProfileId string
		var lifecycleStateId string

		resp, httpRes, err := apiClient.API_V3.LifecycleStatesAPI.GetLifecycleState(context.Background(), identityProfileId, lifecycleStateId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LifecycleStatesAPIService ListLifecycleStates", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var identityProfileId string

		resp, httpRes, err := apiClient.API_V3.LifecycleStatesAPI.ListLifecycleStates(context.Background(), identityProfileId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LifecycleStatesAPIService SetLifecycleState", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var identityId string

		resp, httpRes, err := apiClient.API_V3.LifecycleStatesAPI.SetLifecycleState(context.Background(), identityId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LifecycleStatesAPIService UpdateLifecycleStates", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var identityProfileId string
		var lifecycleStateId string

		resp, httpRes, err := apiClient.API_V3.LifecycleStatesAPI.UpdateLifecycleStates(context.Background(), identityProfileId, lifecycleStateId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
