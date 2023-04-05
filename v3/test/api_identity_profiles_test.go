/*
IdentityNow V3 API

Testing IdentityProfilesApiService

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

func Test_v3_IdentityProfilesApiService(t *testing.T) {

	configuration := openapiclient.NewDefaultConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test IdentityProfilesApiService ExportIdentityProfiles", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.V3.IdentityProfilesApi.ExportIdentityProfiles(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test IdentityProfilesApiService GetDefaultIdentityAttributeConfig", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var identityProfileId string

		resp, httpRes, err := apiClient.V3.IdentityProfilesApi.GetDefaultIdentityAttributeConfig(context.Background(), identityProfileId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test IdentityProfilesApiService GetIdentityProfile", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var identityProfileId string

		resp, httpRes, err := apiClient.V3.IdentityProfilesApi.GetIdentityProfile(context.Background(), identityProfileId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test IdentityProfilesApiService ImportIdentityProfiles", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.V3.IdentityProfilesApi.ImportIdentityProfiles(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test IdentityProfilesApiService ListIdentityProfiles", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.V3.IdentityProfilesApi.ListIdentityProfiles(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
