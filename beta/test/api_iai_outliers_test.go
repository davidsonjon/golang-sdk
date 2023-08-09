/*
IdentityNow Beta API

Testing IAIOutliersApiService

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

func Test_beta_IAIOutliersApiService(t *testing.T) {

	configuration := openapiclient.NewDefaultConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test IAIOutliersApiService ExportOutliersZip", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.BETA.IAIOutliersApi.ExportOutliersZip(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test IAIOutliersApiService GetIdentityOutlierSnapshots", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.BETA.IAIOutliersApi.GetIdentityOutlierSnapshots(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test IAIOutliersApiService GetIdentityOutliers", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.BETA.IAIOutliersApi.GetIdentityOutliers(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test IAIOutliersApiService GetLatestIdentityOutlierSnapshots", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.BETA.IAIOutliersApi.GetLatestIdentityOutlierSnapshots(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test IAIOutliersApiService GetPeerGroupOutliersContributingFeatures", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var outlierId string

		resp, httpRes, err := apiClient.BETA.IAIOutliersApi.GetPeerGroupOutliersContributingFeatures(context.Background(), outlierId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test IAIOutliersApiService IgnoreIdentityOutliers", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.BETA.IAIOutliersApi.IgnoreIdentityOutliers(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test IAIOutliersApiService ListOutliersContributingFeatureAccessItems", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var outlierId string
		var contributingFeatureName string

		resp, httpRes, err := apiClient.BETA.IAIOutliersApi.ListOutliersContributingFeatureAccessItems(context.Background(), outlierId, contributingFeatureName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test IAIOutliersApiService UnIgnoreIdentityOutliers", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.BETA.IAIOutliersApi.UnIgnoreIdentityOutliers(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
