/*
IdentityNow V3 API

Testing PasswordConfigurationAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package v3

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/sailpoint-oss/golang-sdk"
)

func Test_v3_PasswordConfigurationAPIService(t *testing.T) {

	configuration := openapiclient.NewDefaultConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test PasswordConfigurationAPIService CreatePasswordOrgConfig", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.V3.PasswordConfigurationAPI.CreatePasswordOrgConfig(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PasswordConfigurationAPIService GetPasswordOrgConfig", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.V3.PasswordConfigurationAPI.GetPasswordOrgConfig(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PasswordConfigurationAPIService PutPasswordOrgConfig", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.V3.PasswordConfigurationAPI.PutPasswordOrgConfig(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
