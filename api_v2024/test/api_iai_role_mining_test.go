/*
Identity Security Cloud V2024 API

Testing IAIRoleMiningAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package api_v2024

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/sailpoint-oss/golang-sdk/v2"
)

func Test_api_v2024_IAIRoleMiningAPIService(t *testing.T) {

	configuration := openapiclient.NewDefaultConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test IAIRoleMiningAPIService CreatePotentialRoleProvisionRequest", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var sessionId string
		var potentialRoleId string

		resp, httpRes, err := apiClient.API_V2024.IAIRoleMiningAPI.CreatePotentialRoleProvisionRequest(context.Background(), sessionId, potentialRoleId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test IAIRoleMiningAPIService CreateRoleMiningSessions", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.API_V2024.IAIRoleMiningAPI.CreateRoleMiningSessions(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test IAIRoleMiningAPIService DownloadRoleMiningPotentialRoleZip", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var sessionId string
		var potentialRoleId string
		var exportId string

		resp, httpRes, err := apiClient.API_V2024.IAIRoleMiningAPI.DownloadRoleMiningPotentialRoleZip(context.Background(), sessionId, potentialRoleId, exportId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test IAIRoleMiningAPIService ExportRoleMiningPotentialRole", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var sessionId string
		var potentialRoleId string

		resp, httpRes, err := apiClient.API_V2024.IAIRoleMiningAPI.ExportRoleMiningPotentialRole(context.Background(), sessionId, potentialRoleId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test IAIRoleMiningAPIService ExportRoleMiningPotentialRoleAsync", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var sessionId string
		var potentialRoleId string

		resp, httpRes, err := apiClient.API_V2024.IAIRoleMiningAPI.ExportRoleMiningPotentialRoleAsync(context.Background(), sessionId, potentialRoleId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test IAIRoleMiningAPIService ExportRoleMiningPotentialRoleStatus", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var sessionId string
		var potentialRoleId string
		var exportId string

		resp, httpRes, err := apiClient.API_V2024.IAIRoleMiningAPI.ExportRoleMiningPotentialRoleStatus(context.Background(), sessionId, potentialRoleId, exportId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test IAIRoleMiningAPIService GetAllPotentialRoleSummaries", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.API_V2024.IAIRoleMiningAPI.GetAllPotentialRoleSummaries(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test IAIRoleMiningAPIService GetEntitlementDistributionPotentialRole", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var sessionId string
		var potentialRoleId string

		resp, httpRes, err := apiClient.API_V2024.IAIRoleMiningAPI.GetEntitlementDistributionPotentialRole(context.Background(), sessionId, potentialRoleId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test IAIRoleMiningAPIService GetEntitlementsPotentialRole", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var sessionId string
		var potentialRoleId string

		resp, httpRes, err := apiClient.API_V2024.IAIRoleMiningAPI.GetEntitlementsPotentialRole(context.Background(), sessionId, potentialRoleId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test IAIRoleMiningAPIService GetExcludedEntitlementsPotentialRole", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var sessionId string
		var potentialRoleId string

		resp, httpRes, err := apiClient.API_V2024.IAIRoleMiningAPI.GetExcludedEntitlementsPotentialRole(context.Background(), sessionId, potentialRoleId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test IAIRoleMiningAPIService GetIdentitiesPotentialRole", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var sessionId string
		var potentialRoleId string

		resp, httpRes, err := apiClient.API_V2024.IAIRoleMiningAPI.GetIdentitiesPotentialRole(context.Background(), sessionId, potentialRoleId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test IAIRoleMiningAPIService GetPotentialRole", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var sessionId string
		var potentialRoleId string

		resp, httpRes, err := apiClient.API_V2024.IAIRoleMiningAPI.GetPotentialRole(context.Background(), sessionId, potentialRoleId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test IAIRoleMiningAPIService GetPotentialRoleApplications", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var sessionId string
		var potentialRoleId string

		resp, httpRes, err := apiClient.API_V2024.IAIRoleMiningAPI.GetPotentialRoleApplications(context.Background(), sessionId, potentialRoleId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test IAIRoleMiningAPIService GetPotentialRoleSourceIdentityUsage", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var potentialRoleId string
		var sourceId string

		resp, httpRes, err := apiClient.API_V2024.IAIRoleMiningAPI.GetPotentialRoleSourceIdentityUsage(context.Background(), potentialRoleId, sourceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test IAIRoleMiningAPIService GetPotentialRoleSummaries", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var sessionId string

		resp, httpRes, err := apiClient.API_V2024.IAIRoleMiningAPI.GetPotentialRoleSummaries(context.Background(), sessionId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test IAIRoleMiningAPIService GetRoleMiningPotentialRole", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var potentialRoleId string

		resp, httpRes, err := apiClient.API_V2024.IAIRoleMiningAPI.GetRoleMiningPotentialRole(context.Background(), potentialRoleId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test IAIRoleMiningAPIService GetRoleMiningSession", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var sessionId string

		resp, httpRes, err := apiClient.API_V2024.IAIRoleMiningAPI.GetRoleMiningSession(context.Background(), sessionId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test IAIRoleMiningAPIService GetRoleMiningSessionStatus", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var sessionId string

		resp, httpRes, err := apiClient.API_V2024.IAIRoleMiningAPI.GetRoleMiningSessionStatus(context.Background(), sessionId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test IAIRoleMiningAPIService GetRoleMiningSessions", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.API_V2024.IAIRoleMiningAPI.GetRoleMiningSessions(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test IAIRoleMiningAPIService GetSavedPotentialRoles", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.API_V2024.IAIRoleMiningAPI.GetSavedPotentialRoles(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test IAIRoleMiningAPIService PatchPotentialRole", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var sessionId string
		var potentialRoleId string

		resp, httpRes, err := apiClient.API_V2024.IAIRoleMiningAPI.PatchPotentialRole(context.Background(), sessionId, potentialRoleId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test IAIRoleMiningAPIService PatchPotentialRole_1", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var sessionId string
		var potentialRoleId string

		resp, httpRes, err := apiClient.API_V2024.IAIRoleMiningAPI.PatchPotentialRole_0(context.Background(), sessionId, potentialRoleId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test IAIRoleMiningAPIService PatchRoleMiningSession", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var sessionId string

		resp, httpRes, err := apiClient.API_V2024.IAIRoleMiningAPI.PatchRoleMiningSession(context.Background(), sessionId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test IAIRoleMiningAPIService UpdateEntitlementsPotentialRole", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var sessionId string
		var potentialRoleId string

		resp, httpRes, err := apiClient.API_V2024.IAIRoleMiningAPI.UpdateEntitlementsPotentialRole(context.Background(), sessionId, potentialRoleId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}