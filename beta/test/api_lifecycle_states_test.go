/*
IdentityNow Beta API

Testing LifecycleStatesApiService

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

func Test_beta_LifecycleStatesApiService(t *testing.T) {

	configuration := openapiclient.NewDefaultConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test LifecycleStatesApiService ListLifecycleStates", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var identityProfileId string
		var lifecycleStateId string

		resp, httpRes, err := apiClient.BETA.LifecycleStatesApi.ListLifecycleStates(context.Background(), identityProfileId, lifecycleStateId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LifecycleStatesApiService UpdateLifecycleStates", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var identityProfileId string
		var lifecycleStateId string

		resp, httpRes, err := apiClient.BETA.LifecycleStatesApi.UpdateLifecycleStates(context.Background(), identityProfileId, lifecycleStateId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
