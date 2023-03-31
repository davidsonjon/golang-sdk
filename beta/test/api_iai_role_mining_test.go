/*
IdentityNow Beta API

Testing IAIRoleMiningApiService

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

func Test_beta_IAIRoleMiningApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test IAIRoleMiningApiService CreatePotentialRoleProvisionRequest", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var sessionId string
		var potentialRoleId string

		resp, httpRes, err := apiClient.IAIRoleMiningApi.CreatePotentialRoleProvisionRequest(context.Background(), sessionId, potentialRoleId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test IAIRoleMiningApiService CreateRoleMiningSessions", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.IAIRoleMiningApi.CreateRoleMiningSessions(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test IAIRoleMiningApiService DownloadRoleMiningPotentialRoleZip", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var sessionId string
		var potentialRoleId string
		var exportId string

		resp, httpRes, err := apiClient.IAIRoleMiningApi.DownloadRoleMiningPotentialRoleZip(context.Background(), sessionId, potentialRoleId, exportId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test IAIRoleMiningApiService ExportRoleMiningPotentialRole", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var sessionId string
		var potentialRoleId string

		resp, httpRes, err := apiClient.IAIRoleMiningApi.ExportRoleMiningPotentialRole(context.Background(), sessionId, potentialRoleId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test IAIRoleMiningApiService ExportRoleMiningPotentialRoleAsync", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var sessionId string
		var potentialRoleId string

		resp, httpRes, err := apiClient.IAIRoleMiningApi.ExportRoleMiningPotentialRoleAsync(context.Background(), sessionId, potentialRoleId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test IAIRoleMiningApiService ExportRoleMiningPotentialRoleStatus", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var sessionId string
		var potentialRoleId string
		var exportId string

		resp, httpRes, err := apiClient.IAIRoleMiningApi.ExportRoleMiningPotentialRoleStatus(context.Background(), sessionId, potentialRoleId, exportId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test IAIRoleMiningApiService GetEntitlementDistributionPotentialRole", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var sessionId string
		var potentialRoleId string

		resp, httpRes, err := apiClient.IAIRoleMiningApi.GetEntitlementDistributionPotentialRole(context.Background(), sessionId, potentialRoleId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test IAIRoleMiningApiService GetEntitlementsPotentialRole", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var sessionId string
		var potentialRoleId string

		resp, httpRes, err := apiClient.IAIRoleMiningApi.GetEntitlementsPotentialRole(context.Background(), sessionId, potentialRoleId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test IAIRoleMiningApiService GetExcludedEntitlementsPotentialRole", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var sessionId string
		var potentialRoleId string

		resp, httpRes, err := apiClient.IAIRoleMiningApi.GetExcludedEntitlementsPotentialRole(context.Background(), sessionId, potentialRoleId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test IAIRoleMiningApiService GetIdentitiesPotentialRole", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var sessionId string
		var potentialRoleId string

		resp, httpRes, err := apiClient.IAIRoleMiningApi.GetIdentitiesPotentialRole(context.Background(), sessionId, potentialRoleId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test IAIRoleMiningApiService GetPotentialRole", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var sessionId string
		var potentialRoleId string

		resp, httpRes, err := apiClient.IAIRoleMiningApi.GetPotentialRole(context.Background(), sessionId, potentialRoleId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test IAIRoleMiningApiService GetPotentialRoleApplications", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var sessionId string
		var potentialRoleId string

		resp, httpRes, err := apiClient.IAIRoleMiningApi.GetPotentialRoleApplications(context.Background(), sessionId, potentialRoleId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test IAIRoleMiningApiService GetPotentialRoleSummaries", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var sessionId string

		resp, httpRes, err := apiClient.IAIRoleMiningApi.GetPotentialRoleSummaries(context.Background(), sessionId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test IAIRoleMiningApiService GetRoleMiningSession", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var sessionId string

		resp, httpRes, err := apiClient.IAIRoleMiningApi.GetRoleMiningSession(context.Background(), sessionId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test IAIRoleMiningApiService GetRoleMiningSessionStatus", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var sessionId string

		resp, httpRes, err := apiClient.IAIRoleMiningApi.GetRoleMiningSessionStatus(context.Background(), sessionId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test IAIRoleMiningApiService GetRoleMiningSessions", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.IAIRoleMiningApi.GetRoleMiningSessions(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test IAIRoleMiningApiService PatchPotentialRole", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var sessionId string
		var potentialRoleId string

		resp, httpRes, err := apiClient.IAIRoleMiningApi.PatchPotentialRole(context.Background(), sessionId, potentialRoleId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test IAIRoleMiningApiService PatchRoleMiningSession", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var sessionId string

		resp, httpRes, err := apiClient.IAIRoleMiningApi.PatchRoleMiningSession(context.Background(), sessionId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test IAIRoleMiningApiService UpdateEntitlementsPotentialRole", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var sessionId string
		var potentialRoleId string

		resp, httpRes, err := apiClient.IAIRoleMiningApi.UpdateEntitlementsPotentialRole(context.Background(), sessionId, potentialRoleId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
