/*
Identity Security Cloud V3 API

Testing PersonalAccessTokensAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package api_v3

import (
	"context"
	"testing"

	openapiclient "github.com/sailpoint-oss/davidsonjon/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_api_v3_PersonalAccessTokensAPIService(t *testing.T) {

	configuration := openapiclient.NewDefaultConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test PersonalAccessTokensAPIService CreatePersonalAccessToken", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.API_V3.PersonalAccessTokensAPI.CreatePersonalAccessToken(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PersonalAccessTokensAPIService DeletePersonalAccessToken", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		httpRes, err := apiClient.API_V3.PersonalAccessTokensAPI.DeletePersonalAccessToken(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PersonalAccessTokensAPIService ListPersonalAccessTokens", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.API_V3.PersonalAccessTokensAPI.ListPersonalAccessTokens(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PersonalAccessTokensAPIService PatchPersonalAccessToken", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.API_V3.PersonalAccessTokensAPI.PatchPersonalAccessToken(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
