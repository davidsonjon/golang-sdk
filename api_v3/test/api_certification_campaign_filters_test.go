/*
Identity Security Cloud V3 API

Testing CertificationCampaignFiltersAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package api_v3

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/davidsonjon/golang-sdk/v2"
)

func Test_api_v3_CertificationCampaignFiltersAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test CertificationCampaignFiltersAPIService CreateCampaignFilter", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.CertificationCampaignFiltersAPI.CreateCampaignFilter(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CertificationCampaignFiltersAPIService DeleteCampaignFilters", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.CertificationCampaignFiltersAPI.DeleteCampaignFilters(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CertificationCampaignFiltersAPIService GetCampaignFilterById", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.CertificationCampaignFiltersAPI.GetCampaignFilterById(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CertificationCampaignFiltersAPIService ListCampaignFilters", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.CertificationCampaignFiltersAPI.ListCampaignFilters(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CertificationCampaignFiltersAPIService UpdateCampaignFilter", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var filterId string

		resp, httpRes, err := apiClient.CertificationCampaignFiltersAPI.UpdateCampaignFilter(context.Background(), filterId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
