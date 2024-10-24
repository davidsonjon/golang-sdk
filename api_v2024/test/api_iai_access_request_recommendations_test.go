/*
Identity Security Cloud V2024 API

Testing IAIAccessRequestRecommendationsAPIService

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

func Test_api_v2024_IAIAccessRequestRecommendationsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test IAIAccessRequestRecommendationsAPIService AddAccessRequestRecommendationsIgnoredItem", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.IAIAccessRequestRecommendationsAPI.AddAccessRequestRecommendationsIgnoredItem(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test IAIAccessRequestRecommendationsAPIService AddAccessRequestRecommendationsRequestedItem", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.IAIAccessRequestRecommendationsAPI.AddAccessRequestRecommendationsRequestedItem(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test IAIAccessRequestRecommendationsAPIService AddAccessRequestRecommendationsViewedItem", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.IAIAccessRequestRecommendationsAPI.AddAccessRequestRecommendationsViewedItem(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test IAIAccessRequestRecommendationsAPIService AddAccessRequestRecommendationsViewedItems", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.IAIAccessRequestRecommendationsAPI.AddAccessRequestRecommendationsViewedItems(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test IAIAccessRequestRecommendationsAPIService GetAccessRequestRecommendations", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.IAIAccessRequestRecommendationsAPI.GetAccessRequestRecommendations(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test IAIAccessRequestRecommendationsAPIService GetAccessRequestRecommendationsIgnoredItems", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.IAIAccessRequestRecommendationsAPI.GetAccessRequestRecommendationsIgnoredItems(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test IAIAccessRequestRecommendationsAPIService GetAccessRequestRecommendationsRequestedItems", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.IAIAccessRequestRecommendationsAPI.GetAccessRequestRecommendationsRequestedItems(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test IAIAccessRequestRecommendationsAPIService GetAccessRequestRecommendationsViewedItems", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.IAIAccessRequestRecommendationsAPI.GetAccessRequestRecommendationsViewedItems(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
