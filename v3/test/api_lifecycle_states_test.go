/*
IdentityNow V3 API

Testing LifecycleStatesApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package v3

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func Test_v3_LifecycleStatesApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test LifecycleStatesApiService CreateLifecycleState", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var identityProfileId string

		resp, httpRes, err := apiClient.LifecycleStatesApi.CreateLifecycleState(context.Background(), identityProfileId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LifecycleStatesApiService DeleteLifecycleState", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var identityProfileId string
		var lifecycleStateId string

		resp, httpRes, err := apiClient.LifecycleStatesApi.DeleteLifecycleState(context.Background(), identityProfileId, lifecycleStateId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LifecycleStatesApiService GetLifecycleState", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var identityProfileId string
		var lifecycleStateId string

		resp, httpRes, err := apiClient.LifecycleStatesApi.GetLifecycleState(context.Background(), identityProfileId, lifecycleStateId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LifecycleStatesApiService ListLifecycleStates", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var identityProfileId string

		resp, httpRes, err := apiClient.LifecycleStatesApi.ListLifecycleStates(context.Background(), identityProfileId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LifecycleStatesApiService SetLifecycleState", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var identityId string

		resp, httpRes, err := apiClient.LifecycleStatesApi.SetLifecycleState(context.Background(), identityId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LifecycleStatesApiService UpdateLifecycleStates", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var identityProfileId string
		var lifecycleStateId string

		resp, httpRes, err := apiClient.LifecycleStatesApi.UpdateLifecycleStates(context.Background(), identityProfileId, lifecycleStateId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}