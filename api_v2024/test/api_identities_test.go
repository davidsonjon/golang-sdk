/*
Identity Security Cloud V2024 API

Testing IdentitiesAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package api_v2024

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/davidsonjon/golang-sdk/v2"
)

func Test_api_v2024_IdentitiesAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test IdentitiesAPIService DeleteIdentity", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		httpRes, err := apiClient.IdentitiesAPI.DeleteIdentity(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test IdentitiesAPIService GetIdentity", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.IdentitiesAPI.GetIdentity(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test IdentitiesAPIService GetIdentityOwnershipDetails", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var identityId string

		resp, httpRes, err := apiClient.IdentitiesAPI.GetIdentityOwnershipDetails(context.Background(), identityId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test IdentitiesAPIService GetRoleAssignment", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var identityId string
		var assignmentId string

		resp, httpRes, err := apiClient.IdentitiesAPI.GetRoleAssignment(context.Background(), identityId, assignmentId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test IdentitiesAPIService GetRoleAssignments", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var identityId string

		resp, httpRes, err := apiClient.IdentitiesAPI.GetRoleAssignments(context.Background(), identityId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test IdentitiesAPIService ListIdentities", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.IdentitiesAPI.ListIdentities(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test IdentitiesAPIService ResetIdentity", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var identityId string

		httpRes, err := apiClient.IdentitiesAPI.ResetIdentity(context.Background(), identityId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test IdentitiesAPIService StartIdentityProcessing", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.IdentitiesAPI.StartIdentityProcessing(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test IdentitiesAPIService SynchronizeAttributesForIdentity", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var identityId string

		resp, httpRes, err := apiClient.IdentitiesAPI.SynchronizeAttributesForIdentity(context.Background(), identityId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
