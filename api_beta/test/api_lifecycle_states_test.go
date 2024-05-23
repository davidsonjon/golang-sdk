/*
Identity Security Cloud Beta API

Testing LifecycleStatesAPIService

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

func Test_api_beta_LifecycleStatesAPIService(t *testing.T) {

	configuration := openapiclient.NewDefaultConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test LifecycleStatesAPIService GetLifecycleStates", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var identityProfileId string
		var lifecycleStateId string

		resp, httpRes, err := apiClient.API_BETA.LifecycleStatesAPI.GetLifecycleStates(context.Background(), identityProfileId, lifecycleStateId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LifecycleStatesAPIService UpdateLifecycleStates", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var identityProfileId string
		var lifecycleStateId string

		resp, httpRes, err := apiClient.API_BETA.LifecycleStatesAPI.UpdateLifecycleStates(context.Background(), identityProfileId, lifecycleStateId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
