/*
Identity Security Cloud Beta API

Testing IAIOutliersAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package api_beta

import (
	"context"
	"testing"

	openapiclient "github.com/davidsonjon/golang-sdk/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_api_beta_IAIOutliersAPIService(t *testing.T) {

	configuration := openapiclient.NewDefaultConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test IAIOutliersAPIService ExportOutliersZip", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.API_BETA.IAIOutliersAPI.ExportOutliersZip(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test IAIOutliersAPIService GetIdentityOutlierSnapshots", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.API_BETA.IAIOutliersAPI.GetIdentityOutlierSnapshots(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test IAIOutliersAPIService GetIdentityOutliers", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.API_BETA.IAIOutliersAPI.GetIdentityOutliers(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test IAIOutliersAPIService GetLatestIdentityOutlierSnapshots", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.API_BETA.IAIOutliersAPI.GetLatestIdentityOutlierSnapshots(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test IAIOutliersAPIService GetOutlierContributingFeatureSummary", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var outlierFeatureId string

		resp, httpRes, err := apiClient.API_BETA.IAIOutliersAPI.GetOutlierContributingFeatureSummary(context.Background(), outlierFeatureId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test IAIOutliersAPIService GetPeerGroupOutliersContributingFeatures", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var outlierId string

		resp, httpRes, err := apiClient.API_BETA.IAIOutliersAPI.GetPeerGroupOutliersContributingFeatures(context.Background(), outlierId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test IAIOutliersAPIService IgnoreIdentityOutliers", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.API_BETA.IAIOutliersAPI.IgnoreIdentityOutliers(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test IAIOutliersAPIService ListOutliersContributingFeatureAccessItems", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var outlierId string
		var contributingFeatureName string

		resp, httpRes, err := apiClient.API_BETA.IAIOutliersAPI.ListOutliersContributingFeatureAccessItems(context.Background(), outlierId, contributingFeatureName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test IAIOutliersAPIService UnIgnoreIdentityOutliers", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.API_BETA.IAIOutliersAPI.UnIgnoreIdentityOutliers(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
