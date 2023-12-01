/*
IdentityNow Beta API

Testing PasswordManagementAPIService

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

func Test_beta_PasswordManagementAPIService(t *testing.T) {

	configuration := openapiclient.NewDefaultConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test PasswordManagementAPIService GenerateDigitToken", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.BETA.PasswordManagementAPI.GenerateDigitToken(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PasswordManagementAPIService GetIdentityPasswordChangeStatus", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.BETA.PasswordManagementAPI.GetIdentityPasswordChangeStatus(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PasswordManagementAPIService QueryPasswordInfo", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.BETA.PasswordManagementAPI.QueryPasswordInfo(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PasswordManagementAPIService SetIdentityPassword", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.BETA.PasswordManagementAPI.SetIdentityPassword(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
