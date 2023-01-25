/*
IdentityNow Beta API

Testing MFAConfigurationApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package beta

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func Test_beta_MFAConfigurationApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test MFAConfigurationApiService GetMFAConfig", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var method string

		resp, httpRes, err := apiClient.MFAConfigurationApi.GetMFAConfig(context.Background(), method).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MFAConfigurationApiService SetMFAConfig", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var method string

		resp, httpRes, err := apiClient.MFAConfigurationApi.SetMFAConfig(context.Background(), method).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MFAConfigurationApiService TestMFAConfig", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var method string

		resp, httpRes, err := apiClient.MFAConfigurationApi.TestMFAConfig(context.Background(), method).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
