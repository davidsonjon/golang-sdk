/*
IdentityNow Beta API

Testing GovernanceGroupsAPIService

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

func Test_api_beta_GovernanceGroupsAPIService(t *testing.T) {

	configuration := openapiclient.NewDefaultConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test GovernanceGroupsAPIService CreateWorkgroup", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.API_BETA.GovernanceGroupsAPI.CreateWorkgroup(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GovernanceGroupsAPIService DeleteWorkgroup", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		httpRes, err := apiClient.API_BETA.GovernanceGroupsAPI.DeleteWorkgroup(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GovernanceGroupsAPIService DeleteWorkgroupMembers", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var workgroupId string

		resp, httpRes, err := apiClient.API_BETA.GovernanceGroupsAPI.DeleteWorkgroupMembers(context.Background(), workgroupId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GovernanceGroupsAPIService DeleteWorkgroupsInBulk", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.API_BETA.GovernanceGroupsAPI.DeleteWorkgroupsInBulk(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GovernanceGroupsAPIService GetWorkgroup", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.API_BETA.GovernanceGroupsAPI.GetWorkgroup(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GovernanceGroupsAPIService ListConnections", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var workgroupId string

		resp, httpRes, err := apiClient.API_BETA.GovernanceGroupsAPI.ListConnections(context.Background(), workgroupId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GovernanceGroupsAPIService ListWorkgroupMembers", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var workgroupId string

		resp, httpRes, err := apiClient.API_BETA.GovernanceGroupsAPI.ListWorkgroupMembers(context.Background(), workgroupId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GovernanceGroupsAPIService ListWorkgroups", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.API_BETA.GovernanceGroupsAPI.ListWorkgroups(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GovernanceGroupsAPIService PatchWorkgroup", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.API_BETA.GovernanceGroupsAPI.PatchWorkgroup(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GovernanceGroupsAPIService UpdateWorkgroupMembers", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var workgroupId string

		resp, httpRes, err := apiClient.API_BETA.GovernanceGroupsAPI.UpdateWorkgroupMembers(context.Background(), workgroupId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
