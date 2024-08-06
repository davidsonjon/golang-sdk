/*
Identity Security Cloud V2024 API

Testing SuggestedEntitlementDescriptionAPIService

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

func Test_api_v2024_SuggestedEntitlementDescriptionAPIService(t *testing.T) {

	configuration := openapiclient.NewDefaultConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test SuggestedEntitlementDescriptionAPIService GetSedBatchStats", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var batchId string

		resp, httpRes, err := apiClient.API_V2024.SuggestedEntitlementDescriptionAPI.GetSedBatchStats(context.Background(), batchId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SuggestedEntitlementDescriptionAPIService GetSedBatches", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.API_V2024.SuggestedEntitlementDescriptionAPI.GetSedBatches(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SuggestedEntitlementDescriptionAPIService ListSeds", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.API_V2024.SuggestedEntitlementDescriptionAPI.ListSeds(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SuggestedEntitlementDescriptionAPIService PatchSed", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.API_V2024.SuggestedEntitlementDescriptionAPI.PatchSed(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SuggestedEntitlementDescriptionAPIService SubmitSedApproval", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.API_V2024.SuggestedEntitlementDescriptionAPI.SubmitSedApproval(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SuggestedEntitlementDescriptionAPIService SubmitSedAssignment", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.API_V2024.SuggestedEntitlementDescriptionAPI.SubmitSedAssignment(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SuggestedEntitlementDescriptionAPIService SubmitSedBatchRequest", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.API_V2024.SuggestedEntitlementDescriptionAPI.SubmitSedBatchRequest(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
