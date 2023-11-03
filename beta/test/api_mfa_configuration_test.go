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
	openapiclient "github.com/sailpoint-oss/golang-sdk"
)

func Test_beta_MFAConfigurationApiService(t *testing.T) {

	configuration := openapiclient.NewDefaultConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test MFAConfigurationApiService DeleteMFAConfig", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var method string

		resp, httpRes, err := apiClient.BETA.MFAConfigurationApi.DeleteMFAConfig(context.Background(), method).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MFAConfigurationApiService GetMFADuoConfig", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.BETA.MFAConfigurationApi.GetMFADuoConfig(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MFAConfigurationApiService GetMFAOktaConfig", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.BETA.MFAConfigurationApi.GetMFAOktaConfig(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MFAConfigurationApiService SetMFADuoConfig", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.BETA.MFAConfigurationApi.SetMFADuoConfig(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MFAConfigurationApiService SetMFAOktaConfig", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.BETA.MFAConfigurationApi.SetMFAOktaConfig(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MFAConfigurationApiService TestMFAConfig", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var method string

		resp, httpRes, err := apiClient.BETA.MFAConfigurationApi.TestMFAConfig(context.Background(), method).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
