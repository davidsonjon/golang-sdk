/*
Identity Security Cloud V3 API

Testing PasswordSyncGroupsAPIService

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

func Test_api_v3_PasswordSyncGroupsAPIService(t *testing.T) {

	configuration := openapiclient.NewDefaultConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test PasswordSyncGroupsAPIService CreatePasswordSyncGroup", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.API_V3.PasswordSyncGroupsAPI.CreatePasswordSyncGroup(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PasswordSyncGroupsAPIService DeletePasswordSyncGroup", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		httpRes, err := apiClient.API_V3.PasswordSyncGroupsAPI.DeletePasswordSyncGroup(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PasswordSyncGroupsAPIService GetPasswordSyncGroup", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.API_V3.PasswordSyncGroupsAPI.GetPasswordSyncGroup(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PasswordSyncGroupsAPIService GetPasswordSyncGroups", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.API_V3.PasswordSyncGroupsAPI.GetPasswordSyncGroups(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PasswordSyncGroupsAPIService UpdatePasswordSyncGroup", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.API_V3.PasswordSyncGroupsAPI.UpdatePasswordSyncGroup(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
